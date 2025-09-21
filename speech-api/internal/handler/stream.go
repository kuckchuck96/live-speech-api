package handler

import (
	"log"
	"net/http"
	"speech-api/internal/service/transcribe"

	ws "speech-api/internal/websocket"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type (
	Stream interface {
		Listen(*gin.Context)
	}

	stream struct {
		upgrader          websocket.Upgrader
		websocketPool     ws.WebSocketPool
		transcribeService transcribe.TranscribeService
	}
)

func NewStream(transcribeService transcribe.TranscribeService, websocketPool ws.WebSocketPool) Stream {
	return &stream{
		upgrader:          websocket.Upgrader{},
		transcribeService: transcribeService,
		websocketPool:     websocketPool,
	}
}

func (s *stream) Listen(ctx *gin.Context) {
	c, err := s.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to upgrade connection")
		return
	}
	defer c.Close()
	log.Println("WebSocket connected")
	s.websocketPool.Add(c)
	defer func() {
		s.websocketPool.Remove(c)
		if err := c.Close(); err != nil {
			log.Fatalln("websocket close error:", err)
		}
	}()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		res, err := s.transcribeService.SpeechToText(message)
		if err != nil {
			log.Fatal("transcription error:", err)
			break
		}

		err = c.WriteMessage(mt, []byte(*res))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
