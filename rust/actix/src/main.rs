use actix_web::{
    App, HttpResponse, HttpServer, Responder, get, post,
    web::{Json},
};
use serde::{Deserialize, Serialize};
use std::time::{SystemTime, UNIX_EPOCH};

/// GET /ping
/// Response: "ok"
/// Content-Type: text/plain
#[get("/ping")]
async fn ping() -> impl Responder {
    HttpResponse::Ok().content_type("text/plain").body("ok")
}

/// JSON response structure for /json
#[derive(Serialize)]
struct StatusResponse {
    status: &'static str,
    service: &'static str,
    timestamp: u64,
}

/// GET /json
#[get("/json")]
async fn json() -> impl Responder {
    let timestamp = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_secs();

    let resp = StatusResponse {
        status: "ok",
        service: "actix",
        timestamp,
    };

    HttpResponse::Ok().json(resp)
}

/// Request / response structure for /echo
#[derive(Deserialize, Serialize)]
struct EchoMessage {
    message: String,
}

/// POST /echo
#[post("/echo")]
async fn echo(payload: Json<EchoMessage>) -> impl Responder {
    HttpResponse::Ok().json(payload.into_inner())
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(ping).service(json).service(echo))
        .bind(("127.0.0.1", 8080))?
        .run()
        .await
}
