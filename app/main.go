package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/dwin/ohioGeranium/app/controller"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

var reqLog *log.Logger

func main() {
	r := mux.NewRouter()

	// Open Request Log
	f, err := os.OpenFile(os.Getenv("LOGDIR")+"/request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer f.Close()
	reqLog = log.New(f, "", 0)

	// Routes
	r.HandleFunc("/status", controller.Status).Methods("GET")
	r.HandleFunc("/search", controller.NewSearch).Methods("POST")

	// Chain Middlewares
	chain := alice.New(logRequest).Then(r)

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":1313", chain))
}

// logRequest middlware function writes all incoming requests to request log
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			reqLog.Printf("Parse user ip (%s) error: %s", r.RemoteAddr, err)
		}

		userIP := net.ParseIP(ip)
		if userIP == nil {
			reqLog.Printf("userip: %q is not in form of ip:port", r.RemoteAddr)
			return
		}

		// This will only be defined when site is accessed via non-anonymous proxy
		// and takes precedence over RemoteAddr
		fwdIP := r.Header.Get("X-Forwarded-For")
		if len(fwdIP) > 0 {
			ip = fwdIP
		}

		reqLog.Printf("%s %s %s %s\n", time.Now().Format(time.RFC3339), ip, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
