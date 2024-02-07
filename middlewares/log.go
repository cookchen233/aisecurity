package middlewares

import (
	"aisecurity/utils"
	http2 "aisecurity/utils/http"
	"bytes"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"hash/crc32"
	"io"
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RequestLog ref: https://github.com/gin-gonic/gin/blob/v1.9.0/logger.go#L182
func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// record trace information
		s := sessions.Default(c)
		sessid := s.Get("anonymous")
		if sessid == nil || sessid == "" {
			sessid = fmt.Sprintf("%08x", crc32.ChecksumIEEE([]byte(uuid.NewString())))
			s.Set("anonymous", sessid)
			err := s.Save()
			if err != nil {
				utils.Logger.Error("save session error", zap.Error(err))
			}
		}

		traceid := sessid.(string) + "_" + fmt.Sprintf("%08x", crc32.ChecksumIEEE([]byte(uuid.NewString())))
		// set the context information
		c.Set("traceid", traceid)
		var oriLoggerL = utils.Logger.L
		utils.Logger.L = utils.Logger.With(
			zap.String("traceid", traceid),
			zap.String("url", c.Request.URL.String()))
		zfs := []zap.Field{
			zap.String("start", start.Format(time.StampMicro)),
			zap.String("method", c.Request.Method),
			zap.String("content_type", c.ContentType()),
			zap.String("client_ip", c.ClientIP()),
			zap.String("remote_ip", c.RemoteIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		}
		var bodyBytes []byte
		var err error
		if c.ContentType() == "multipart/form-data" {
			form, err := c.MultipartForm()
			var bs = fmt.Sprintf("multipart/form-data: %v", form)
			if err != nil {
				bs = fmt.Sprintf("multipart/form-data error: %v, form: %v", err, form)
			}
			bodyBytes = []byte(bs)
		} else if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			// Create a buffer to hold the body data
			var buf bytes.Buffer
			tee := io.TeeReader(c.Request.Body, &buf)
			bodyBytes, err = io.ReadAll(tee)
			if err != nil {
				utils.Logger.Error("Failed to read request body")
			}
			// Replace the request body so it can be read again
			c.Request.Body = io.NopCloser(&buf)
		}

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		// Stop timer
		stop := time.Now()
		latency := stop.Sub(start)
		if latency > time.Minute {
			latency = latency.Truncate(time.Second)
		}

		//stackTraces := make([]errors.StackTrace, 0, len(c.Errors))
		//for _, e := range c.Errors {
		//	if err, ok := e.Err.(interface{ StackTrace() errors.StackTrace }); ok {
		//		stackTraces = append(stackTraces, err.StackTrace())
		//	}
		//}
		stackTraces := make([][]string, 0, len(c.Errors))
		for _, e := range c.Errors {
			stackTraces = append(stackTraces, strings.Split(strings.Replace(fmt.Sprintf("%+v", e.Err), "\t", "", -1), "\n"))
		}
		zfs = append(zfs, zap.String("latency", fmt.Sprintf("%s", latency)))
		zfs = append(zfs, zap.Int("status", c.Writer.Status()))
		if len(c.Errors) > 0 {
			resp, err := io.ReadAll(blw.body)
			if err != nil {
				utils.Logger.Error("Failed to read response body")
			}
			zfs = append(zfs, zap.ByteString("body", bodyBytes))
			zfs = append(zfs, zap.ByteString("resp", resp))
			zfs = append(zfs, zap.String("error", c.Errors.ByType(gin.ErrorTypePrivate).String()))
			if c.Writer.Status() >= 500 {
				zfs = append(zfs, zap.Any("stack", stackTraces))
				utils.Logger.Error("Request", zfs...)
			} else {
				if gin.Mode() == "debug" {
					utils.Logger.Warn("Request", zfs...)
				}
			}
		} else {
			if gin.Mode() == "debug" {
				utils.Logger.Info("Request", zfs...)
			}
		}

		// clear the with fields
		utils.Logger.L = oriLoggerL

	}
}

// Recovery ref: https://github.com/gin-gonic/gin/blob/v1.9.0/recovery.go#L33
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				headers := strings.Split(string(httpRequest), "\r\n")
				for idx, header := range headers {
					current := strings.Split(header, ":")
					if current[0] == "Authorization" {
						headers[idx] = current[0] + ": *"
					}
				}
				headersToStr := strings.Join(headers, "\r\n")
				var panicErr error
				panicErr, ok := err.(error)
				if !ok {
					// If r is indeed an error, use it.
					panicErr = fmt.Errorf("panic: %v", err)
				}
				c.Error(panicErr)
				traceid := c.GetString("traceid")
				if brokenPipe {
					utils.Logger.Fatal("Panic",
						zap.String("traceid", traceid),
						zap.String("url", c.Request.URL.String()),
						zap.String("method", c.Request.Method),
						zap.String("content_type", c.ContentType()),
						zap.String("client_ip", c.ClientIP()),
						zap.String("remote_ip", c.RemoteIP()),
						zap.String("user_agent", c.Request.UserAgent()),
						zap.String("headers", headersToStr),
						zap.Any("err", err),
						zap.Stack("stack"),
					)
				} else {
					utils.Logger.Fatal("Panic",
						zap.String("traceid", traceid),
						zap.String("url", c.Request.URL.String()),
						zap.String("method", c.Request.Method),
						zap.String("content_type", c.ContentType()),
						zap.String("client_ip", c.ClientIP()),
						zap.String("remote_ip", c.RemoteIP()),
						zap.String("user_agent", c.Request.UserAgent()),
						zap.String("headers", headersToStr),
						zap.Any("err", err),
						zap.Stack("stack"),
						zap.String("panicRecoveredTime", time.Now().Format(time.RFC3339)),
					)
				}
				http2.Error(c, panicErr)
				return
			}
		}()
		c.Next()
	}
}
