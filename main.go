package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/kataras/muxie"
)

type qjson struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
	Time   string `json:"time"`
}

type dataStore struct {
	totalQuotes int
	quotes      []qdata
}

var database dataStore

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
	srvInfo := ParseEnvVars()

	database.quotes = parseJSON("data/data.json")
	database.totalQuotes = len(database.quotes)
	fmt.Println("Total quotes: ", database.totalQuotes)
	//get the address and the port
	addr := JoinAddr(srvInfo.addr, srvInfo.port)

	fmt.Printf("Serving on......\n")
	fmt.Printf("Address: %s\n", addr)

	//create a mux
	mux := muxie.NewMux()
	mux.HandleFunc("/", indexHandler)
	srv := &http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a hit from :", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	randoNum := random(0, database.totalQuotes)
	curtime := time.Now()
	q := qjson{
		Quote:  database.quotes[randoNum].quote,
		Author: database.quotes[randoNum].author,
		Time:   curtime.Format("02-01-2006 15:04:05"),
	}
	qqjson, err := json.Marshal(q)
	if err != nil {
		fmt.Println("Error!")
		fmt.Fprintf(w, "error")
	}
	w.Write(qqjson)
}
