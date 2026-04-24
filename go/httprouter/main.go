package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

/* ---------- /ping ---------- */
/* GET /ping
   Response: "ok"
   Content-Type: text/plain
*/
func pingHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

/* ---------- /json ---------- */

type StatusResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Timestamp int64  `json:"timestamp"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := StatusResponse{
		Status:    "ok",
		Service:   "go",
		Timestamp: time.Now().Unix(),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

/* ---------- /echo ---------- */

type EchoMessage struct {
	Message string `json:"message"`
}

func echoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "expected application/json", http.StatusBadRequest)
		return
	}

	var payload EchoMessage
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(payload)
}

/* ---------- main ---------- */

func main() {
	router := httprouter.New()

	// Explicit method routing (no method checks inside handlers)
	router.GET("/ping", pingHandler)
	router.GET("/json", jsonHandler)
	router.POST("/echo", echoHandler)

	log.Println("Listening on 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}

