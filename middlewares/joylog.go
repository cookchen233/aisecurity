package middlewares

import (
	"aisecurity/utils"
	http2 "aisecurity/utils/http"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var LOGFILEBASE = "./logs/request/"
var _log *log.Logger
var _f *os.File
var _today = time.Now()

func init() {
	var err error
	if _, err := os.Stat(LOGFILEBASE); os.IsNotExist(err) {
		// The directory does not exist, create it
		err := os.MkdirAll(LOGFILEBASE, os.ModePerm)
		if err != nil {
			fmt.Printf("failed to create directory: %v", err)
		}
		fmt.Printf("Directory created: %s\n", LOGFILEBASE)
	}
	var infoLogFile = LOGFILEBASE + time.Now().Format("2006-01-02") + ".log"

	_f, err = os.OpenFile(infoLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panicf("error opening file: %v", err)
	}
	//wr := io.MultiWriter(_f, os.Stdout)
	//_log = log.New(wr, "INFO ", log.LstdFlags|log.Lmicroseconds)
	_log = log.New(_f, "INFO ", log.LstdFlags|log.Lmicroseconds)
	_log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func JoyRequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// this if block is where important things happen for rotation
		// changing output file for logger
		if !dateEqual(_today, time.Now()) {
			_today = time.Now()

			dailyLogFile := LOGFILEBASE + time.Now().Format("2006-01-02") + ".log"
			newF, err := os.OpenFile(dailyLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Panicf("error opening file: %v", err)
			}
			//wr := io.MultiWriter(newF, os.Stdout)
			//_log.SetOutput(wr)
			_log.SetOutput(newF)
		}

		var bodyString string
		if c.ContentType() == "multipart/form-data" {
			form, err := c.MultipartForm()
			bodyString = fmt.Sprintf("multipart/form-data: %v", form)
			if err != nil {
				bodyString = fmt.Sprintf("multipart/form-data error: %v, form: %v", err, form)
			}
		} else if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			body, _ := io.ReadAll(c.Request.Body)

			processedData, err := utils.ProcessJSON(body)
			if err != nil {
				bodyString = string(body)
			} else {
				prettyJSON, _ := json.MarshalIndent(processedData, "", "    ")
				bodyString = string(prettyJSON)
			}
			c.Request.Body = io.NopCloser(bytes.NewReader(body))
		}

		// better if you have a user in the context
		var bs = "non debug mode"
		if gin.Mode() == "debug" {
			bs = bodyString
		}
		go _log.Printf("JOY REQUEST:\n traceid: %v\n ip: %v\n ip2: %v\n method: %v\n url: %v\n body: %v", c.GetString("traceid"), c.ClientIP(), c.Request.RemoteAddr, c.Request.Method, c.Request.RequestURI, bs)
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		var p http2.Payload
		payload, ex := c.Get("payload")
		if !ex {
			utils.Logger.Error("payload not set")
		} else {
			p = payload.(http2.Payload)
		}
		var respBody = "non debug mode"
		if gin.Mode() == "debug" {
			bd, err := io.ReadAll(blw.body)
			if err != nil {
				utils.Logger.Error("Failed to read response body", zap.Error(err))
			}
			respBody = string(bd)
			bb := []rune(bodyString)
			bs = string(bb[:min(len(bb), 256)])
		}
		go _log.Printf("JOY RESPONSE:\n traceid: %v\n status: %v\n payload: %v\n respBody: %v\n reqBody: %v\n%v\n\n\n", c.GetString("traceid"), blw.Status(), p, respBody, bs, strings.Repeat("=", 128))
	}
}

func dateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
