use poem::{Route, Server, get, handler, listener::TcpListener, post, web::Json};
use serde::{Deserialize, Serialize};
use std::time::{SystemTime, UNIX_EPOCH};

/// GET /ping
/// Response: "ok"
/// Content-Type: text/plain
#[handler]
fn ping() -> &'static str {
    "ok"
}

/// JSON response structure for /json
#[derive(Serialize)]
struct StatusResponse {
    status: &'static str,
    service: &'static str,
    timestamp: u64,
}

/// GET /json
#[handler]
fn json() -> Json<StatusResponse> {
    let timestamp = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_secs();

    Json(StatusResponse {
        status: "ok",
        service: "poem",
        timestamp,
    })
}

/// Request / response structure for /echo
#[derive(Deserialize, Serialize)]
struct EchoMessage {
    message: String,
}

/// POST /echo
#[handler]
fn echo(Json(payload): Json<EchoMessage>) -> Json<EchoMessage> {
    Json(payload)
}

#[tokio::main]
async fn main() -> std::io::Result<()> {
    let app = Route::new()
        .at("/ping", get(ping))
        .at("/json", get(json))
        .at("/echo", post(echo));

    Server::new(TcpListener::bind("127.0.0.1:8080"))
        .run(app)
        .await
}
