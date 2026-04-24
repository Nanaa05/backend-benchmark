from flask import Flask, request, jsonify, Response
import time

app = Flask(__name__)

# ---------- /ping ----------
# GET /ping
# Response: "ok"
# Content-Type: text/plain
@app.route("/ping", methods=["GET"])
def ping():
    return Response("ok", content_type="text/plain")


# ---------- /json ----------
# GET /json
@app.route("/json", methods=["GET"])
def json_endpoint():
    return jsonify({
        "status": "ok",
        "service": "flask",
        "timestamp": int(time.time()),
    })


# ---------- /echo ----------
# POST /echo
@app.route("/echo", methods=["POST"])
def echo():
    if request.content_type != "application/json":
        return jsonify({"error": "expected application/json"}), 400

    return jsonify(request.get_json())


# ---------- main ----------
if __name__ == "__main__":
    app.run(
        host="127.0.0.1",
        port=8080,
        debug=False,
        use_reloader=False,
    )

