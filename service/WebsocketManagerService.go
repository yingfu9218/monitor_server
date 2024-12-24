package service

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"monitor_server/entity"
	"sync"
)

var WebSocketManagerServ *WebSocketManagerService = NewWebSocketManagerService()

type WebSocketManagerService struct {
	connections map[string]*entity.WebsocketConnection
	mu          sync.Mutex
}

// 构造函数，初始化 WebSocketManagerService
func NewWebSocketManagerService() *WebSocketManagerService {
	return &WebSocketManagerService{
		connections: make(map[string]*entity.WebsocketConnection),
	}
}

func (m *WebSocketManagerService) AddConnection(id string, conn *websocket.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.connections[id] = &entity.WebsocketConnection{
		ID:   id,
		Conn: conn,
	}
}

func (m *WebSocketManagerService) RemoveConnection(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.connections, id)
}

func (m *WebSocketManagerService) GetConnection(id string) *entity.WebsocketConnection {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.connections[id]
}
func (m *WebSocketManagerService) GetConnections() map[string]*entity.WebsocketConnection {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.connections
}

func (m *WebSocketManagerService) CreateSessionId() (connID string) {
	// 为每个连接生成一个唯一的 ID
	connID = uuid.New().String()
	return connID
}
