package handlers

import (
	"aisecurity/utils"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

// ToycarHandler Toycar handles
type ToycarHandler struct {
	msgs     []byte
	isOnline bool
}

func NewToycarHandler() *ToycarHandler {
	return &ToycarHandler{}
}

var client1MsgCh = make(chan []byte) // Channel to send message from wsHandler1 to wsHandler2

var mu sync.Mutex

func (toycar *ToycarHandler) getConnsByGroupId(groupID string, pool *utils.ConnPool) (error, []*websocket.Conn) {
	keyConns, ok := pool.GetByGroupID(groupID)
	if !ok {
		err := errors.New("not found connection group")
		return err, nil
	}
	var conns []*websocket.Conn
	for connId, _ := range keyConns {
		conn, ok := pool.GetByConnectionID(connId, groupID)
		if !ok {
			log.Println("not found connection id " + connId)
		}
		conns = append(conns, conn)
	}
	return nil, conns
}

// ControlPanel 控制面板
func (toycar *ToycarHandler) ControlPanel(receivedMsg []byte, conn *websocket.Conn, pool *utils.ConnPool) (returnMsg []byte) {
	log.Printf("Received on ControlPanel: %s", receivedMsg)
	go func() {
		for {
			select {
			case msg := <-client1MsgCh:
				err, conns := toycar.getConnsByGroupId("accept-control", pool)
				if err != nil {
					log.Println(err)
					continue
				}
				for _, conn := range conns {
					mu.Lock()
					err := conn.WriteMessage(websocket.TextMessage, msg)
					mu.Unlock()
					if err != nil {
						log.Println("Failed to write message:", err)
					}
				}
			}
		}
	}()
	go func() {
		for {
			err, acceptConns := toycar.getConnsByGroupId("accept-control", pool)
			if err != nil {
				if err.Error() != "not found connection group" {
					log.Println(err)
					time.Sleep(10 * time.Second)
					continue
				}
			}
			err, controlConns := toycar.getConnsByGroupId("control-panel", pool)
			if err != nil {
				log.Println(err)
				time.Sleep(10 * time.Second)
				continue
			}
			log.Println("acceptConns count:", len(acceptConns))
			online := "the-toy-car-is-offline"
			if len(acceptConns) > 0 {
				online = "the-toy-car-is-online"
			}
			for _, controlConn := range controlConns {
				mu.Lock()
				err := controlConn.WriteMessage(websocket.TextMessage, []byte(online))
				mu.Unlock()
				if err != nil {
					log.Println("Failed to write message:", err)
				}
			}
			time.Sleep(3 * time.Second)
		}
	}()

	mu.Lock()
	client1MsgCh <- receivedMsg // Send message to wsHandler2
	mu.Unlock()

	return nil
}

// AcceptControl 小车接受控制
func (toycar *ToycarHandler) AcceptControl(receivedMsg []byte, conn *websocket.Conn, poll *utils.ConnPool) (returnMsg []byte) {
	log.Printf("Received on AcceptControl: %s", receivedMsg)
	//return []byte("Handled message on ws1: " + string(receivedMsg))

	toycar.isOnline = true

	return nil
}
