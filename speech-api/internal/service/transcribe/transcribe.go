package transcribe

import (
	"log"
	"speech-api/internal/pkg/httpclient"
)

type (
	TranscribeService interface {
		SpeechToText([]byte) (*string, error)
	}

	transcribeService struct {
		httpClient httpclient.HttpClient
	}
)

func NewTranscribeService(httpClient httpclient.HttpClient) TranscribeService {
	return &transcribeService{
		httpClient: httpClient,
	}
}

func (t *transcribeService) SpeechToText(data []byte) (*string, error) {
	log.Println("Starting transcription service...")
	if res, err := t.httpClient.Execute(data); err != nil {
		return nil, err
	} else {
		str := string(res)
		return &str, nil
	}
}
