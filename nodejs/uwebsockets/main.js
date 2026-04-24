const uWS = require("uWebSockets.js");

const PORT = 8080;

uWS
  .App()

  /* ---------- /ping ---------- */
  .get("/ping", (res, req) => {
    res.writeHeader("Content-Type", "text/plain");
    res.end("ok");
  })

  /* ---------- /json ---------- */
  .get("/json", (res, req) => {
    const timestamp = Math.floor(Date.now() / 1000);

    const body = JSON.stringify({
      status: "ok",
      service: "uwebsockets",
      timestamp,
    });

    res.writeHeader("Content-Type", "application/json");
    res.end(body);
  })

  /* ---------- /echo ---------- */
  .post("/echo", (res, req) => {
    let contentType = req.getHeader("content-type") || "";

    if (!contentType.includes("application/json")) {
      res.writeStatus("400 Bad Request");
      res.writeHeader("Content-Type", "application/json");
      res.end(JSON.stringify({ error: "expected application/json" }));
      return;
    }

    let buffer = Buffer.alloc(0);

    res.onData((ab, isLast) => {
      const chunk = Buffer.from(ab);
      buffer = Buffer.concat([buffer, chunk]);

      if (isLast) {
        try {
          const json = JSON.parse(buffer.toString());

          res.writeHeader("Content-Type", "application/json");
          res.end(JSON.stringify(json));
        } catch {
          res.writeStatus("400 Bad Request");
          res.writeHeader("Content-Type", "application/json");
          res.end(JSON.stringify({ error: "invalid json" }));
        }
      }
    });

    res.onAborted(() => {
      buffer = null;
    });
  })

  .listen("127.0.0.1", PORT, (token) => {
    if (token) {
      console.log(`Listening on http://127.0.0.1:${PORT}`);
    } else {
      console.log("Failed to listen");
    }
  });
