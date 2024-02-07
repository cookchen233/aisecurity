package middlewares

import (
	http2 "aisecurity/utils/http"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Custom writer to buffer the response
type bufferedWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (b *bufferedWriter) Write(data []byte) (int, error) {
	return b.body.Write(data) // Only write to the buffer, not to the ResponseWriter
}

// Define a structure that represents the expected Data structure
type FuyouVideoReponse struct {
	VideoId string `json:"VideoId"`
}

func FuyouResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		buffer := bytes.NewBufferString("")

		// Replace the current writer with our bufferedWriter
		originalWriter := c.Writer
		c.Writer = &bufferedWriter{body: buffer, ResponseWriter: originalWriter}

		// Process the request
		c.Next()

		// Retrieve the payload set by the handler
		payloadValue, exists := c.Get("payload")
		if !exists {
			c.Writer = originalWriter
			c.JSON(http.StatusInternalServerError, gin.H{"error": "payload not set"})
			return
		}

		// Cast the payload to the Payload struct
		payload, ok := payloadValue.(http2.Payload)
		if !ok {
			c.Writer = originalWriter
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid payload type"})
			return
		}

		var modifiedResponse []byte
		var err error
		if data, ok := payload.Data.(FuyouVideoReponse); ok {
			// Define the new response structure
			var response struct {
				Result struct {
					Code int    `json:"Code"`
					Desc string `json:"Desc"`
				} `json:"Result"`
				VideoId string `json:"VideoId"`
			}
			response.Result.Desc = payload.Message
			response.Result.Code = payload.Code
			if payload.Code == 200 {
				response.Result.Code = 0
			}
			response.VideoId = data.VideoId
			modifiedResponse, err = json.Marshal(response)
		} else {
			var response2 struct {
				Result struct {
					Code int    `json:"Code"`
					Desc string `json:"Desc"`
				} `json:"Result"`
			}
			modifiedResponse, err = json.Marshal(response2)
		}

		if err != nil {
			c.Writer = originalWriter
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.Writer = originalWriter
		c.Writer.Write(modifiedResponse)
	}
}
