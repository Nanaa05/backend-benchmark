// nodejs/fastify/main.js
"use strict";

const fastify = require("fastify")({
  logger: false, // IMPORTANT: disable logger for benchmarking
});

/* ---------- /ping ---------- */
/* GET /ping
   Response: "ok"
   Content-Type: text/plain
*/
fastify.get("/ping", async (request, reply) => {
  reply.type("text/plain").send("ok");
});

/* ---------- /json ---------- */

fastify.get("/json", async (request, reply) => {
  reply.send({
    status: "ok",
    service: "fastify",
    timestamp: Math.floor(Date.now() / 1000),
  });
});

/* ---------- /echo ---------- */

fastify.post("/echo", async (request, reply) => {
  const contentType = request.headers["content-type"];

  if (contentType !== "application/json") {
    reply.code(400);
    return { error: "expected application/json" };
  }

  // Fastify parses JSON natively
  return request.body;
});

/* ---------- main ---------- */

const start = async () => {
  try {
    await fastify.listen({ port: 8080, host: "127.0.0.1" });
    console.log("Listening on http://127.0.0.1:8080");
  } catch (err) {
    fastify.log.error(err);
    process.exit(1);
  }
};

start();
