package logger

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type (
	loggedResponse struct {
		w    http.ResponseWriter
		size int
	}
)

func (r *loggedResponse) Header() http.Header {
	return r.w.Header()
}

func (r *loggedResponse) WriteHeader(statusCode int) {
	r.w.WriteHeader(statusCode)
}

func (r *loggedResponse) Write(p []byte) (int, error) {
	size, err := r.w.Write(p)
	r.size = size
	return size, err
}

func InitLogger(level string) error {
	log.SetFormatter(&log.JSONFormatter{})
	logLevel, err := log.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("failed on parsing log level: %w", err)
	}
	log.SetLevel(logLevel)

	return nil
}

func LogMiddleware(next http.Handler) http.Handler {
	fh := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		mw := &loggedResponse{
			w:    w,
			size: 0,
		}
		next.ServeHTTP(mw, r)

		requestTime := time.Since(start)
		log.WithFields(log.Fields{
			"uri":      r.RequestURI,
			"method":   r.Method,
			"duration": requestTime,
		}).Info("incoming request > ")

		log.WithFields(log.Fields{
			"size": mw.size,
		}).Info("outcoming request < ")
	}

	return http.HandlerFunc(fh)
}
