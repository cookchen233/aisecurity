package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
	"sync"
)

var (
	store       = sessions.NewFilesystemStore("/path/to/session/files", []byte("secret-key"))
	sessionList = make(map[string]bool)
	mutex       = &sync.Mutex{}
)

// Middleware to track active sessions
func sessionTrackerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "session-name")
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		mutex.Lock()
		if session.IsNew {
			session.ID = generateSessionID() // Replace with your method to generate a unique session ID
			sessionList[session.ID] = true
			session.Save(c.Request, c.Writer)
		}
		mutex.Unlock()

		c.Next()
	}
}

// Handler to get the count of active sessions
func activeUserCountHandler(c *gin.Context) {
	mutex.Lock()
	count := len(sessionList)
	mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{"active_user_count": count})
}

func main() {
	router := gin.Default()
	router.Use(sessionTrackerMiddleware())

	router.GET("/active-user-count", activeUserCountHandler)

	router.Run(":8080")
}

func generateSessionID() string {
	// Implement a function to generate a unique session ID
	return "some_unique_session_id"
}
