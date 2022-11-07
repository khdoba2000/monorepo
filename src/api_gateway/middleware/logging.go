package middleware

import (
	"log"
	"net/http"
	"time"
)

//"github.com/gorilla/handlers"
//gorillas custom handler, do not used as it prints log without timestamp

// useful for logging with response code
// https://stackoverflow.com/questions/42162211/gorilla-mux-best-way-to-catch-response-codes

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

//newLoggingResponseWriter overrides loggingResponseWriter, adding status code
func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

//Logging logs the time
func Logging(next http.Handler) http.Handler {
	//return handlers.LoggingHandler(os.Stdout, next)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		lrw := newLoggingResponseWriter(w)

		next.ServeHTTP(lrw, req)
		statusCode := lrw.statusCode
		log.Printf("%s %d %s %s", req.Method, statusCode, req.RequestURI, time.Since(start))
		return
	})
}
