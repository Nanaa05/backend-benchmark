// nodejs/http/main.js
const http = require("http");

/* ---------- helpers ---------- */

function sendText(res, status, body) {
  res.writeHead(status, {
    "Content-Type": "text/plain",
    "Content-Length": Buffer.byteLength(body),
  });
  res.end(body);
}

function sendJSON(res, status, obj) {
  const body = JSON.stringify(obj);
  res.writeHead(status, {
    "Content-Type": "application/json",
    "Content-Length": Buffer.byteLength(body),
  });
  res.end(body);
}

/* ---------- server ---------- */

const server = http.createServer((req, res) => {
  const { method, url, headers } = req;

  /* ---------- /ping ---------- */
  if (method === "GET" && url === "/ping") {
    return sendText(res, 200, "ok");
  }

  /* ---------- /json ---------- */
  if (method === "GET" && url === "/json") {
    return sendJSON(res, 200, {
      status: "ok",
      service: "node-http",
      timestamp: Math.floor(Date.now() / 1000),
    });
  }

  /* ---------- /echo ---------- */
  if (method === "POST" && url === "/echo") {
    if (headers["content-type"] !== "application/json") {
      return sendJSON(res, 400, { error: "expected application/json" });
    }

    let body = "";
    req.on("data", (chunk) => {
      body += chunk;
    });

    req.on("end", () => {
      try {
        const payload = JSON.parse(body);
        sendJSON(res, 200, payload);
      } catch {
        sendJSON(res, 400, { error: "invalid json" });
      }
    });

    return;
  }

  /* ---------- fallback ---------- */
  res.writeHead(404);
  res.end();
});

/* ---------- main ---------- */

const PORT = 8080;
server.listen(PORT, "127.0.0.1", () => {
  console.log(`Listening on http://127.0.0.1:${PORT}`);
});
