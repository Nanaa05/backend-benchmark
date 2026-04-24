use axum::{
    Json, Router,
    http::{HeaderMap, StatusCode},
    routing::{get, post},
};
use serde::{Deserialize, Serialize};
use std::time::{SystemTime, UNIX_EPOCH};

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/ping", get(ping))
        .route("/json", get(json))
        .route("/echo", post(echo));

    let listener = tokio::net::TcpListener::bind("127.0.0.1:8080")
        .await
        .unwrap();

    axum::serve(listener, app).await.unwrap();
}

/// GET /ping
async fn ping() -> (HeaderMap, &'static str) {
    let mut headers = HeaderMap::new();
    headers.insert("Content-Type", "text/plain".parse().unwrap());
    (headers, "ok")
}

/// JSON response
#[derive(Serialize)]
struct StatusResponse {
    status: &'static str,
    service: &'static str,
    timestamp: u64,
}

/// GET /json
async fn json() -> Json<StatusResponse> {
    let timestamp = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_secs();

    Json(StatusResponse {
        status: "ok",
        service: "axum",
        timestamp,
    })
}

/// Echo payload
#[derive(Deserialize, Serialize)]
struct EchoMessage {
    message: String,
}

/// POST /echo
async fn echo(Json(payload): Json<EchoMessage>) -> (StatusCode, Json<EchoMessage>) {
    (StatusCode::OK, Json(payload))
}
