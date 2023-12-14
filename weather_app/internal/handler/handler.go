// handler.go
package handler

import (
	"net"
	"net/http"
	"time"

	"github.com/PunitNaran/weather_app/internal/cache"
	"github.com/sirupsen/logrus" // Import Logrus for structured logging
)

type Handler struct {
	c     *http.Client
	store cache.Cache
	l     *logrus.Logger
}

func (h *Handler) CreateClient(store cache.Cache, l *logrus.Logger) {
	h.c = &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	h.l = l
	h.store = store
}
