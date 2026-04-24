import json
import time

async def app(scope, receive, send):
    if scope["type"] != "http":
        return

    method = scope["method"]
    path = scope["path"]

    # ---------- /ping ----------
    if method == "GET" and path == "/ping":
        await send({
            "type": "http.response.start",
            "status": 200,
            "headers": [
                (b"content-type", b"text/plain"),
            ],
        })
        await send({
            "type": "http.response.body",
            "body": b"ok",
        })
        return

    # ---------- /json ----------
    if method == "GET" and path == "/json":
        body = json.dumps({
            "status": "ok",
            "service": "asgi",
            "timestamp": int(time.time()),
        }).encode()

        await send({
            "type": "http.response.start",
            "status": 200,
            "headers": [
                (b"content-type", b"application/json"),
            ],
        })
        await send({
            "type": "http.response.body",
            "body": body,
        })
        return

    # ---------- /echo ----------
    if method == "POST" and path == "/echo":
        body_bytes = b""
        while True:
            event = await receive()
            body_bytes += event.get("body", b"")
            if not event.get("more_body", False):
                break

        await send({
            "type": "http.response.start",
            "status": 200,
            "headers": [
                (b"content-type", b"application/json"),
            ],
        })
        await send({
            "type": "http.response.body",
            "body": body_bytes,
        })
        return

    # ---------- 404 ----------
    await send({
        "type": "http.response.start",
        "status": 404,
        "headers": [],
    })
    await send({
        "type": "http.response.body",
        "body": b"not found",
    })

