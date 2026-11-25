package logger

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
)

const (
	HandlerText = "text"
	HandlerJSON = "json"
)

var ErrUnsupportedHandler = errors.New("unsupported log handler")

func New(level slog.Level, handlerName string, writer io.Writer) (*slog.Logger, error) {
	switch handlerName {
	case HandlerText:
		return slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{Level: level})), nil
	case HandlerJSON:
		return slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: level})), nil
	default:
		return nil, fmt.Errorf("%w: %q", ErrUnsupportedHandler, handlerName)
	}
}
