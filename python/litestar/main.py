from litestar import Litestar, get, post
from pydantic import BaseModel
import time


# ---------- /ping ----------
@get("/ping", sync_to_thread=False)
def ping() -> str:
    return "ok"


# ---------- /json ----------
class StatusResponse(BaseModel):
    status: str
    service: str
    timestamp: int


@get("/json", sync_to_thread=False)
def json_endpoint() -> StatusResponse:
    return StatusResponse(
        status="ok",
        service="litestar",
        timestamp=int(time.time()),
    )


# ---------- /echo ----------
class EchoMessage(BaseModel):
    message: str


@post("/echo", sync_to_thread=False)
def echo(data: EchoMessage) -> EchoMessage:
    return data


# ---------- app ----------
app = Litestar(
    route_handlers=[ping, json_endpoint, echo],
)

