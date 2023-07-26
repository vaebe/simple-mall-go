package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/thoas/go-funk"
	"go.uber.org/zap"
	"net/http"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

var clients = make(map[string]*Client)

// SendMessage 向客户端发送消息格式
type SendMessage struct {
	Type string `json:"type"`
	Code int    `json:"code"`
	Data any    `json:"data"`
}

// SendMessageToClient 向客户端发送消息
func SendMessageToClient(clientID string, message interface{}) error {
	client, found := clients[clientID]
	if !found {
		return fmt.Errorf("client not found with ID: %s", clientID)
	}

	return client.Conn.WriteJSON(message)
}

var origins = []string{"http://localhost:5173", "https://mall.vaebe.top"}
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		var origin = r.Header.Get("origin")
		zap.S().Debug("ws 请求连接地址", origin)

		return funk.Contains(origins, origin)
	},
}

// HandleWebSocket
//
//	@Summary		处理 ws 请求
//	@Description	处理 ws 请求
//	@Tags				ws
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	false	"连接id"
//	@Success		200		{object}	string
//	@Failure		500		{object}	string
//	@Router			/ws [get]
func HandleWebSocket(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		zap.S().Debug(err.Error())
		return
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
	}(conn)

	// 获取连接 id
	clientID := ctx.Query("id")
	zap.S().Debug(clientID)

	// 保存连接
	client := &Client{ID: clientID, Conn: conn}
	clients[clientID] = client
	defer delete(clients, clientID)

	zap.S().Debug("ws连接信息", clients, "客户端id", clientID)

	err = SendMessageToClient(clientID, &SendMessage{Type: "success", Code: 0, Data: "连接成功！"})
	if err != nil {
		return
	}

	// 在此处继续处理WebSocket连接，保持长连接状态
	for {
		// 读取客户端发送的消息
		_, message, err := conn.ReadMessage()
		if err != nil {
			zap.S().Debug("读取消息错误：", err)
			break // 跳出循环，关闭连接
		}

		zap.S().Debug("客户端消息: ", message, "客户端id", clientID)
	}
}
