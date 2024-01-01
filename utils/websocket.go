package utils

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	maxConnections = 10000             // Maximum number of connections
	idleTimeout    = 600 * time.Second // Idle timeout for connections
)

type ConnInfo struct {
	conn *websocket.Conn
	time time.Time
}

type ConnPool struct {
	mu          sync.RWMutex
	connections map[string]map[string]*ConnInfo // Connection groups mapping to connection information
}

func NewConnPool() *ConnPool {
	pool := &ConnPool{
		connections: make(map[string]map[string]*ConnInfo),
	}
	go func() {
		// Periodically close idle connections
		ticker := time.NewTicker(idleTimeout)
		defer ticker.Stop()

		for {
			<-ticker.C
			pool.CloseIdleConnections()
		}
	}()
	return pool
}

func (cp *ConnPool) Add(connID string, conn *websocket.Conn, groupID string) bool {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	// Check if the maximum connections limit has been reached
	if len(cp.connections[groupID]) >= maxConnections {
		return false
	}

	// Create group map if it doesn't exist
	if _, ok := cp.connections[groupID]; !ok {
		cp.connections[groupID] = make(map[string]*ConnInfo)
	}
	cp.connections[groupID][connID] = &ConnInfo{conn: conn, time: time.Now()}
	return true
}

func (cp *ConnPool) RemoveConnection(connID string, groupID string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	if conns, ok := cp.connections[groupID]; ok {
		if _, exists := conns[connID]; exists {
			// Close the connection and remove from the group
			err := conns[connID].conn.Close()
			if err != nil {
				log.Println("Failed to close websocket connection:", err)
				return
			}
			delete(conns, connID)
		}
		// If there are no more connections, remove the group
		if len(conns) == 0 {
			delete(cp.connections, groupID)
		}
	}
}

func (cp *ConnPool) GetByConnectionID(connID string, groupID string) (*websocket.Conn, bool) {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	if conns, ok := cp.connections[groupID]; ok {
		if connInfo, exists := conns[connID]; exists {
			return connInfo.conn, true
		}
	}
	return nil, false
}

func (cp *ConnPool) GetByGroupID(groupID string) (map[string]*ConnInfo, bool) {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	conns, ok := cp.connections[groupID]
	return conns, ok
}

func (cp *ConnPool) CloseIdleConnections() {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	currentTime := time.Now()
	for groupID, conns := range cp.connections {
		for connID, connInfo := range conns {
			// Check if the connection has been idle for longer than the idleTimeout
			if currentTime.Sub(connInfo.time) > idleTimeout {
				// Close the connection and remove it from the group
				err := connInfo.conn.Close()
				if err != nil {
					log.Println("Failed to close websocket connection:", err)
					return
				}
				delete(conns, connID)
			}
		}
		// If there are no more connections in the group, remove the group
		if len(conns) == 0 {
			delete(cp.connections, groupID)
		}
	}
}

func HandleWebSocket(messageHandler func([]byte, *websocket.Conn, *ConnPool) []byte, pool *ConnPool, groupID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Failed to upgrade to WebSocket connection:", err)
			c.Abort()
			return
		}
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				log.Println("Failed to close websocket connection:", err)
			}
		}(conn)

		if groupID == "" {
			groupID = strconv.FormatInt(time.Now().UnixNano(), 10)
		}
		connID := fmt.Sprintf("%p", conn)

		pool.Add(connID, conn, groupID)
		defer pool.RemoveConnection(connID, groupID)

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				if !websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("Failed to read message:", err)
				}
				break
			}
			response := messageHandler(msg, conn, pool)
			if response != nil {
				err := conn.WriteMessage(websocket.TextMessage, response)
				if err != nil {
					log.Println("Failed to write message:", err)
					break
				}
			}
		}
	}
}
