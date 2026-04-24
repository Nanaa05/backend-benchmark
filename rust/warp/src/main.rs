use serde::{Deserialize, Serialize};
use std::time::{SystemTime, UNIX_EPOCH};
use warp::{Filter};

#[tokio::main]
async fn main() {
    // GET /ping
    let ping = warp::path("ping")
        .and(warp::get())
        .map(|| warp::reply::with_header("ok", "Content-Type", "text/plain"));

    // GET /json
    let json = warp::path("json").and(warp::get()).map(|| {
        let timestamp = SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .unwrap()
            .as_secs();

        let resp = StatusResponse {
            status: "ok",
            service: "warp",
            timestamp,
        };

        warp::reply::json(&resp)
    });

    // POST /echo
    let echo = warp::path("echo")
        .and(warp::post())
        .and(warp::body::json())
        .map(|payload: EchoMessage| warp::reply::json(&payload));

    // Combine routes
    let routes = ping.or(json).or(echo);

    // Run server
    warp::serve(routes).run(([127, 0, 0, 1], 8080)).await;
}

/// JSON response for /json
#[derive(Serialize)]
struct StatusResponse {
    status: &'static str,
    service: &'static str,
    timestamp: u64,
}

/// JSON payload for /echo
#[derive(Deserialize, Serialize)]
struct EchoMessage {
    message: String,
}
