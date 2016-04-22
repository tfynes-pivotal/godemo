package main

import (
	"fmt"
	"log"
	"log/syslog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	buildstamp string
	githash    string
	rlog       *syslog.Writer
)

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "4000"
	}

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/info", infoHandler)
	log.Printf(fmt.Sprintf("Listening at %s", port))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello swisscom cloud!")
}

func infoHandler(w http.ResponseWriter, req *http.Request) {
	r := "Binary INFO:\n"
	r += fmt.Sprintf("buildstamp %s\n", buildstamp)
	r += fmt.Sprintf("githash %s\n", githash)
	r += fmt.Sprintf("\n\nENV Variables\n")
	for _, e := range os.Environ() {
		r += fmt.Sprintf("%s\n", e)
	}
	fmt.Fprintln(w, r)

	return
}
