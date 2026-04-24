import { Hono } from 'hono'
import { serve } from '@hono/node-server'

const app = new Hono()

/* ---------- /ping ---------- */
/* GET /ping
   Response: "ok"
   Content-Type: text/plain
*/
app.get('/ping', (c) => {
  return c.text('ok')
})

/* ---------- /json ---------- */

app.get('/json', (c) => {
  const timestamp = Math.floor(Date.now() / 1000)

  return c.json({
    status: 'ok',
    service: 'hono',
    timestamp,
  })
})

/* ---------- /echo ---------- */

app.post('/echo', async (c) => {
  const contentType = c.req.header('content-type') || ''
  if (!contentType.includes('application/json')) {
    return c.json({ error: 'expected application/json' }, 400)
  }

  const body = await c.req.json()
  return c.json(body)
})

/* ---------- main ---------- */

const PORT = 8080

serve({
  fetch: app.fetch,
  hostname: '127.0.0.1',
  port: PORT,
})

console.log(`Listening on http://127.0.0.1:${PORT}`)

