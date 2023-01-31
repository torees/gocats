package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Hello world, the web server
	APP_PORT := os.Getenv("APP_PORT")
	APP_MSG := os.Getenv("APP_MSG")
	LOG_LEVEL := os.Getenv("LOG_LEVEL")
	f, _ := os.ReadFile("catlist.txt")
	catList := strings.Split(string(f), "\n")

	//[]string{"Max", "George", "Floyd", "Sir Bark"}

	// HELLO endpoint, responds with custom string
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Cache-Control", "max-age=2592000")
		io.WriteString(w, ""+APP_MSG+"\n")
		logRequest(req)
		log.Println("Returned app msg: " + APP_MSG)
	}

	// Main cat endpoints, responds with a random cat name
	catHandler := func(w http.ResponseWriter, req *http.Request) {
		thisCat := catList[rand.Intn(len(catList))]
		io.WriteString(w, ""+thisCat+"\n")
		logRequest(req)
		opsProcessed.Inc()
		if LOG_LEVEL == "CATS" {
			log.Println("Returned cat: " + thisCat)
		}
	}
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", catHandler)
	log.Println("Starting CAT API, listening for requests on port " + APP_PORT)
	log.Fatal(http.ListenAndServe(":"+APP_PORT+"", nil))
}

func logRequest(r *http.Request) {
	//uri := r.URL.String()
	method := r.Method
	client := r.RemoteAddr
	log.Println(method + " client: " + client + "")
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gocats_processed_cats_total",
		Help: "The total number of returned cat names",
	})
)

func catLen(catList []string) bool {
	return (len(catList) > 10)
}
