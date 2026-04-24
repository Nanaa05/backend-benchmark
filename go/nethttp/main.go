package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

/* ---------- /ping ---------- */
/* GET /ping
   Response: "ok"
   Content-Type: text/plain
*/
func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp := StatusResponse{
		Status:    "ok",
		Service:   "go",
		Timestamp: time.Now().Unix(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

/* ---------- /echo ---------- */

type EchoMessage struct {
	Message string `json:"message"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
	json.NewEncoder(w).Encode(payload)
}

/* ---------- main ---------- */

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/json", jsonHandler)
	mux.HandleFunc("/echo", echoHandler)

	log.Println("Listening on 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}
