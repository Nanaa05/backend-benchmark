#![deny(warnings)]

use bytes::Bytes;
use http_body_util::{BodyExt, Full};
use hyper::{
    Method, Request, Response, StatusCode, body::Incoming, server::conn::http1, service::service_fn,
};
use hyper_util::rt::TokioIo;
use serde::{Deserialize, Serialize};
use std::{
    convert::Infallible,
    net::SocketAddr,
    time::{SystemTime, UNIX_EPOCH},
};
use tokio::net::TcpListener;

// -------------------- data types --------------------

#[derive(Serialize)]
struct StatusResponse {
    status: &'static str,
    service: &'static str,
    timestamp: u64,
}

#[derive(Deserialize, Serialize)]
struct EchoMessage {
    message: String,
}

// -------------------- handler --------------------

async fn handler(req: Request<Incoming>) -> Result<Response<Full<Bytes>>, Infallible> {
    match (req.method(), req.uri().path()) {
        // GET /ping
        (&Method::GET, "/ping") => Ok(Response::builder()
            .status(StatusCode::OK)
            .header("Content-Type", "text/plain")
            .body(Full::new(Bytes::from_static(b"ok")))
            .unwrap()),

        // GET /json
        (&Method::GET, "/json") => {
            let timestamp = SystemTime::now()
                .duration_since(UNIX_EPOCH)
                .unwrap()
                .as_secs();

            let resp = StatusResponse {
                status: "ok",
                service: "hyper",
                timestamp,
            };

            let body = serde_json::to_vec(&resp).unwrap();

            Ok(Response::builder()
                .status(StatusCode::OK)
                .header("Content-Type", "application/json")
                .body(Full::new(Bytes::from(body)))
                .unwrap())
        }

        // POST /echo
        (&Method::POST, "/echo") => {
            // Hyper 1.x: collect the body via http-body-util's BodyExt
            let collected = req.into_body().collect().await.unwrap();
            let body_bytes = collected.to_bytes();

            let payload: EchoMessage = serde_json::from_slice(&body_bytes).unwrap();
            let body = serde_json::to_vec(&payload).unwrap();

            Ok(Response::builder()
                .status(StatusCode::OK)
                .header("Content-Type", "application/json")
                .body(Full::new(Bytes::from(body)))
                .unwrap())
        }

        // 404
        _ => Ok(Response::builder()
            .status(StatusCode::NOT_FOUND)
            .header("Content-Type", "text/plain")
            .body(Full::new(Bytes::from_static(b"Not Found")))
            .unwrap()),
    }
}

// -------------------- server bootstrap --------------------

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = SocketAddr::from(([127, 0, 0, 1], 8080));
    let listener = TcpListener::bind(addr).await?;
    println!("Listening on http://{}", addr);

    loop {
        let (stream, _) = listener.accept().await?;
        let io = TokioIo::new(stream);

        tokio::task::spawn(async move {
            if let Err(err) = http1::Builder::new()
                .serve_connection(io, service_fn(handler))
                .await
            {
                eprintln!("Connection error: {err}");
            }
        });
    }
}
