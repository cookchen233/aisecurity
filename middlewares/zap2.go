package middlewares

import (
	"aisecurity/utils"
	http2 "aisecurity/utils/http"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"hash/crc32"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggerMiddleware ref: https://github.com/gin-gonic/gin/blob/v1.9.0/logger.go#L182
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// record trace information
		s := sessions.Default(c)
		fmt.Println("s.id", s.ID())
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
		utils.Logger.L = utils.Logger.With(zap.String("traceid", traceid))
		c.Next()

		// Stop timer
		stop := time.Now()
		latency := stop.Sub(start)
		if latency > time.Minute {
			latency = latency.Truncate(time.Second)
		}

		stackTraces := make([]errors.StackTrace, 0, len(c.Errors))
		for _, e := range c.Errors {
			if err, ok := e.Err.(interface{ StackTrace() errors.StackTrace }); ok {
				stackTraces = append(stackTraces, err.StackTrace())
			}
		}
		zfs := []zap.Field{
			zap.String("Start", start.Format(time.RFC3339)),
			zap.Int("Status", c.Writer.Status()),
			zap.String("Latency", fmt.Sprintf("%s", latency)),
			zap.String("Method", c.Request.Method),
			zap.String("Path", c.Request.URL.Path),
			zap.String("RawQuery", c.Request.URL.RawQuery),
			zap.String("ClientIP", c.ClientIP()),
			zap.String("RemoteIP", c.RemoteIP()),
			zap.String("userAgent", c.Request.UserAgent()),
		}
		if len(c.Errors) > 0 {
			if c.Writer.Status() >= 500 {
				zfs = append(
					zfs,
					zap.String("Error", c.Errors.ByType(gin.ErrorTypePrivate).String()),
					zap.Any("Stack", stackTraces),
				)
				utils.Logger.Error("GIN request", zfs...)
			} else {
				utils.Logger.Warn("GIN request", zfs...)
			}
		} else {
			utils.Logger.Info("GIN request", zfs...)
		}

		// clear the with fields
		utils.Logger.L = oriLoggerL

	}
}

// RecoveryMiddleware ref: https://github.com/gin-gonic/gin/blob/v1.9.0/recovery.go#L33
func RecoveryMiddleware2() gin.HandlerFunc {
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
				c.Error(err.(error))
				if brokenPipe {
					utils.Logger.Error(c.Request.URL.String(),
						zap.Any("err", err),
						zap.String("headers", headersToStr),
						zap.Stack("stack"),
					)
					// If the connection is dead, we can't write a status to it.

					c.Abort()
				} else {
					utils.Logger.Panic(c.Request.URL.String(),
						zap.Any("err", err),
						zap.String("headers", headersToStr),
						zap.Stack("stack"),
						zap.String("panicRecoveredTime", time.Now().Format(time.RFC3339)),
					)
					http2.Error(c, err.(error), http.StatusInternalServerError)
				}
			}
		}()
		c.Next()
	}
}

// GinRecovery 用于替换gin框架的Recovery中间件，因为传入参数，再包一层
func RecoveryMiddleware() gin.HandlerFunc {
	logger := zap.L()
	return func(c *gin.Context) {
		defer func() {
			// defer 延迟调用，出了异常，处理并恢复异常，记录日志
			if err := recover(); err != nil {
				//  这个不必须，检查是否存在断开的连接(broken pipe或者connection reset by peer)---------开始--------
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				//httputil包预先准备好的DumpRequest方法
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// 如果连接已断开，我们无法向其写入状态
					c.Error(err.(error))
					c.Abort()
					return
				}
				//  这个不必须，检查是否存在断开的连接(broken pipe或者connection reset by peer)---------结束--------

				// 是否打印堆栈信息，使用的是debug.Stack()，传入false，在日志中就没有堆栈信息

				logger.Error("[Recovery from panic]",
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())),
				)
				// 有错误，直接返回给前端错误，前端直接报错
				//c.AbortWithStatus(http.StatusInternalServerError)
				// 该方式前端不报错
				c.String(200, "访问出错了")
			}
		}()
		c.Next()
	}
}
