package service

import (
	"context"
)

// SpeechRecognitionService is the interface for speech recognition
type SpeechRecognitionService interface {
	ProcessAudio(ctx context.Context, filePath string) (string, error)
}

// SpeechRecognitionServiceImpl implements SpeechRecognitionService
type SpeechRecognitionServiceImpl struct{}

// NewSpeechRecognitionService creates a new speech recognition service
func NewSpeechRecognitionService() SpeechRecognitionService {
	return &SpeechRecognitionServiceImpl{}
}

// ProcessAudio processes the audio file and returns the recognized text
func (s *SpeechRecognitionServiceImpl) ProcessAudio(ctx context.Context, filePath string) (string, error) {
	return "", nil
}
