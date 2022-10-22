package slognil

import "golang.org/x/exp/slog"

type Handler struct{}

var _ slog.Handler = Handler{}

func (Handler) Enabled(slog.Level) bool {
	return false
}

func (Handler) Handle(r slog.Record) error {
	return nil
}

func (h Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h Handler) WithGroup(name string) slog.Handler {
	return h
}
