package logger

//@ NOTE Shyft logs responses in terminal
import (
	"log"
	"net/http"
	"time"
	"fmt"
)

//Logger logs responses to terminal
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"\033[32mINFO:\033[0m  %s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// Log logs string message prefixed with INFO
func Log(msg string) {
	log.Println("\033[32mINFO:\033[0m ", msg)
}

// Warn logs string message prefixed with WARN
func Warn(msg string) {
	log.Println("---------------------------")
	log.Println(fmt.Sprintf("\033[31mWARN:\033[0m  %s", msg))
	log.Println("---------------------------")
}

// WriteLogger returns and handles error message
func WriteLogger(n int, err error) {
	if err != nil {
		Warn("Write failed: %v"+  err.Error())
	}
}

