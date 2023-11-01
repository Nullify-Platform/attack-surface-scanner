package middleware

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/nullify-platform/logger/pkg/logger"
)

type HTTPRequestMetadata struct {
	Service         string        `json:"service"`
	Host            string        `json:"host"`
	Method          string        `json:"method"`
	URL             string        `json:"url"`
	StatusCode      int           `json:"statusCode"`
	RequestHeaders  []string      `json:"requestHeaders"`
	ResponseHeaders []string      `json:"responseHeaders"`
	Duration        time.Duration `json:"duration"`
}

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (rw *ResponseWriter) Header() http.Header {
	return rw.ResponseWriter.Header()
}

func (rw *ResponseWriter) Write(data []byte) (int, error) {
	return rw.ResponseWriter.Write(data)
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Error("endpoint handler panicked",
					logger.Any("err", err),
					logger.Trace(debug.Stack()),
				)
			}
		}()

		start := time.Now()
		rw := &ResponseWriter{ResponseWriter: w}
		next.ServeHTTP(rw, req)

		reqHeaders := []string{}
		for header, values := range req.Header {
			for _, value := range values {
				reqHeaders = append(reqHeaders, header+": "+value)
			}
		}

		resHeaders := []string{}
		for header, values := range rw.Header() {
			for _, value := range values {
				resHeaders = append(resHeaders, header+": "+value)
			}
		}

		logger.Info(
			"request summary",
			logger.Any("requestSummary", &HTTPRequestMetadata{
				Service:         "testserver",
				Host:            req.Host,
				Method:          req.Method,
				URL:             req.URL.String(),
				StatusCode:      rw.StatusCode,
				RequestHeaders:  reqHeaders,
				ResponseHeaders: resHeaders,
				Duration:        time.Since(start),
			}),
		)
	})
}
