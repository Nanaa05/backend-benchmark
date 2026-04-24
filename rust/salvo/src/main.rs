use salvo::prelude::*;
use serde::{Deserialize, Serialize};
use std::time::{SystemTime, UNIX_EPOCH};

/// GET /ping
#[handler]
async fn ping(res: &mut Response) {
    res.add_header("Content-Type", "text/plain", true).unwrap();
    res.render("ok");
}

/// JSON response for /json
#[derive(Serialize)]
struct StatusResponse {
    status: &'static str,
    service: &'static str,
    timestamp: u64,
}

/// GET /json
#[handler]
async fn json(res: &mut Response) {
    let timestamp = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_secs();

    let body = StatusResponse {
        status: "ok",
        service: "salvo",
        timestamp,
    };

    res.render(Json(body));
}

/// JSON payload for /echo
#[derive(Deserialize, Serialize)]
struct EchoMessage {
    message: String,
}

#[handler]
async fn echo(req: &mut Request, res: &mut Response) {
    let payload: EchoMessage = match req.parse_json().await {
        Ok(v) => v,
        Err(_) => {
            res.status_code(StatusCode::BAD_REQUEST);
            return;
        }
    };

    res.render(Json(payload));
}


#[tokio::main]
async fn main() {
    let router = Router::new()
        .push(Router::with_path("ping").get(ping))
        .push(Router::with_path("json").get(json))
        .push(Router::with_path("echo").post(echo));

    let acceptor = TcpListener::new("127.0.0.1:8080").bind().await;

    Server::new(acceptor).serve(router).await;
}
