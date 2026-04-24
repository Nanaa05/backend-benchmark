const express = require("express");
const app = express();

// Built-in JSON body parser (needed for /echo)
app.use(express.json());

/* ---------- /ping ---------- */
/* GET /ping
   Response: "ok"
   Content-Type: text/plain
*/
app.get("/ping", (req, res) => {
  res.type("text/plain").send("ok");
});

/* ---------- /json ---------- */

app.get("/json", (req, res) => {
  const timestamp = Math.floor(Date.now() / 1000);

  res.json({
    status: "ok",
    service: "express",
    timestamp,
  });
});

/* ---------- /echo ---------- */

app.post("/echo", (req, res) => {
  if (!req.is("application/json")) {
    return res.status(400).json({ error: "expected application/json" });
  }

  res.json(req.body);
});

/* ---------- main ---------- */

const PORT = 8080;
app.listen(PORT, "127.0.0.1", () => {
  console.log(`Listening on http://127.0.0.1:${PORT}`);
});
