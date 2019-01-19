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
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func Log(msg string) {
	log.Println("INFO - ", msg)
}

func Warn(msg string) {
	log.Println("---------------------------")
	log.Println(fmt.Sprintf("WARN: %s", msg))
	log.Println("---------------------------")
}

