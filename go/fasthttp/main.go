package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

/* ---------- /ping ---------- */
/* GET /ping
   Response: "ok"
   Content-Type: text/plain
*/
func pingHandler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodGet {
		ctx.Error("method not allowed", fasthttp.StatusMethodNotAllowed)
		return
	}

	ctx.SetContentType("text/plain")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString("ok")
}

/* ---------- /json ---------- */

type StatusResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Timestamp int64  `json:"timestamp"`
}

func jsonHandler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodGet {
		ctx.Error("method not allowed", fasthttp.StatusMethodNotAllowed)
		return
	}

	resp := StatusResponse{
		Status:    "ok",
		Service:   "go",
		Timestamp: time.Now().Unix(),
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)

	// zero-copy write via encoder
	_ = json.NewEncoder(ctx).Encode(resp)
}

/* ---------- /echo ---------- */

type EchoMessage struct {
	Message string `json:"message"`
}

func echoHandler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodPost {
		ctx.Error("method not allowed", fasthttp.StatusMethodNotAllowed)
		return
	}

	if string(ctx.Request.Header.ContentType()) != "application/json" {
		ctx.Error("expected application/json", fasthttp.StatusBadRequest)
		return
	}

	var payload EchoMessage
	if err := json.Unmarshal(ctx.PostBody(), &payload); err != nil {
		ctx.Error("invalid json", fasthttp.StatusBadRequest)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)

	_ = json.NewEncoder(ctx).Encode(payload)
}

/* ---------- main ---------- */

func main() {
	handler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/ping":
			pingHandler(ctx)
		case "/json":
			jsonHandler(ctx)
		case "/echo":
			echoHandler(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	log.Println("Listening on 127.0.0.1:8080")
	log.Fatal(fasthttp.ListenAndServe("127.0.0.1:8080", handler))
}

