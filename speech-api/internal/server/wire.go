//go:build wireinject
// +build wireinject

package server

import (
	"net/http"
	"speech-api/internal/handler"
	"speech-api/internal/pkg/httpclient"
	"speech-api/internal/service/transcribe"
	"speech-api/internal/websocket"

	"github.com/google/wire"
)

func InitializeServer() *http.Server {
	wire.Build(
		websocket.NewWebSocketPool,
		httpclient.NewHttpClient,
		transcribe.NewTranscribeService,
		handler.NewStream,
		NewServer,
	)
	return &http.Server{}
}
