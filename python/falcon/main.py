import falcon
import json
import time


class Ping:
    def on_get(self, req, resp):
        resp.status = falcon.HTTP_200
        resp.content_type = "text/plain"
        resp.text = "ok"


class JsonEndpoint:
    def on_get(self, req, resp):
        resp.status = falcon.HTTP_200
        resp.media = {
            "status": "ok",
            "service": "falcon",
            "timestamp": int(time.time()),
        }


class Echo:
    def on_post(self, req, resp):
        if req.content_type != "application/json":
            resp.status = falcon.HTTP_400
            resp.media = {"error": "expected application/json"}
            return

        data = req.media
        resp.status = falcon.HTTP_200
        resp.media = data


app = falcon.App()

app.add_route("/ping", Ping())
app.add_route("/json", JsonEndpoint())
app.add_route("/echo", Echo())

