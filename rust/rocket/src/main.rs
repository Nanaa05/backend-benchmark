#[macro_use]
extern crate rocket;

use rocket::http::ContentType;
use rocket::serde::{Deserialize, Serialize, json::Json};
use std::time::{SystemTime, UNIX_EPOCH};

/// GET /ping
#[get("/ping")]
fn ping() -> (ContentType, &'static str) {
    (ContentType::Plain, "ok")
}

/// JSON response for /json
#[derive(Serialize)]
#[serde(crate = "rocket::serde")]
struct StatusResponse {
    status: &'static str,
    service: &'static str,
    timestamp: u64,
}

/// GET /json
#[get("/json")]
fn json() -> Json<StatusResponse> {
    let timestamp = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_secs();

    Json(StatusResponse {
        status: "ok",
        service: "rocket",
        timestamp,
    })
}

/// JSON payload for /echo
#[derive(Deserialize, Serialize)]
#[serde(crate = "rocket::serde")]
struct EchoMessage {
    message: String,
}

/// POST /echo
#[post("/echo", format = "json", data = "<payload>")]
fn echo(payload: Json<EchoMessage>) -> Json<EchoMessage> {
    Json(payload.into_inner())
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![ping, json, echo])
}
