package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const logFile string = "development.log"

func main() {

	httpPort := 4000

	openLogFile(logFile)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", rootHandler)
	http.HandleFunc("/file", openFile)

	fmt.Printf("listening on %v\n", httpPort)
	fmt.Printf("Logging to %v\n", logFile)

	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1><div>Welcome to realtime logger</div>")
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s URL[%s] -  Agent:[%s]\n", r.RemoteAddr, r.Method, r.URL, r.UserAgent())
		fmt.Printf("%s %s %s Agent:[%s]\n", r.RemoteAddr, r.Method, r.URL, r.UserAgent())
		handler.ServeHTTP(w, r)
	})
}

func openFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	var lastShowed []byte
	for {
		lastLine := readLastLine(logFile)
		if !bytes.Equal(lastLine, lastShowed) {
			w.Write(lastLine)
			lastShowed = lastLine
			w.(http.Flusher).Flush()
		}
		time.Sleep(time.Second)
	}
}

func openLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}

		log.SetOutput(lf)
	}
}

func readLastLine(filePath string) []byte {
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer f.Close()
	var lastLine []byte
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lastLine = scanner.Bytes()
	}
	return append(lastLine, []byte("\n")...)
}
