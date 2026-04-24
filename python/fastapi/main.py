from fastapi import FastAPI, Request, HTTPException
from pydantic import BaseModel
import time

app = FastAPI()


# ---------- /ping ----------
# GET /ping
# Response: "ok"
# Content-Type: text/plain
@app.get("/ping", response_class=None)
async def ping():
    return "ok"


# ---------- /json ----------
class StatusResponse(BaseModel):
    status: str
    service: str
    timestamp: int


@app.get("/json", response_model=StatusResponse)
async def json_endpoint():
    return StatusResponse(
        status="ok",
        service="fastapi",
        timestamp=int(time.time()),
    )


# ---------- /echo ----------
class EchoMessage(BaseModel):
    message: str


@app.post("/echo", response_model=EchoMessage)
async def echo(payload: EchoMessage):
    return payload

