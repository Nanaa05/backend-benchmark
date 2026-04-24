from sanic import Sanic, text, json
from sanic.request import Request
import time

app = Sanic("benchmark-sanic")


# --------------------
# /ping
# --------------------
@app.get("/ping")
async def ping(request: Request):
    return text("ok", content_type="text/plain")


# --------------------
# /json
# --------------------
@app.get("/json")
async def json_endpoint(request: Request):
    return json({
        "status": "ok",
        "service": "sanic",
        "timestamp": int(time.time()),
    })


# --------------------
# /echo
# --------------------
@app.post("/echo")
async def echo(request: Request):
    if request.content_type != "application/json":
        return json({"error": "expected application/json"}, status=400)

    try:
        body = request.json
    except Exception:
        return json({"error": "invalid json"}, status=400)

    return json(body)


# --------------------
# Entry point
# --------------------
if __name__ == "__main__":
    app.run(
        host="127.0.0.1",
        port=8080,
        single_process=True,  # IMPORTANT: match benchmark fairness
        access_log=False,
        auto_reload=False,
    )

