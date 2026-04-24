asgi
```
export PYTHONHASHSEED=0
export PYTHONDONTWRITEBYTECODE=1

gunicorn main:app \
  -k uvicorn.workers.UvicornWorker \
  -w 4 \
  --env UVICORN_LOOP=uvloop \
  --env UVICORN_HTTP=httptools \
  -b 127.0.0.1:8080 \
  --log-level critical \
  --access-logfile /dev/null \
  --error-logfile /dev/null
```

django
```
gunicorn main:application \
  --bind 127.0.0.1:8080 \
  --workers 4 \
  --threads 1 \
  --worker-class sync \
  --log-level critical \
  --access-logfile /dev/null \
  --error-logfile /dev/null
```

falcon
```
gunicorn main:app \
  --bind 127.0.0.1:8080 \
  --workers 4 \
  --threads 1 \
  --worker-class sync \
  --log-level critical \
  --access-logfile /dev/null \
  --error-logfile /dev/null
```
fastapi
```
uvicorn main:app \
  --host 127.0.0.1 \
  --port 8080 \
  --log-level critical \
  --no-access-log
```
fastapiuv
```
gunicorn main:app \
  -k uvicorn.workers.UvicornWorker \
  -w 4 \
  --env UVICORN_LOOP=uvloop \
  --env UVICORN_HTTP=httptools \
  -b 127.0.0.1:8080 \
  --log-level critical \
  --access-logfile /dev/null \
  --error-logfile /dev/null
```
flask
```
gunicorn main:app \
  --bind 127.0.0.1:8080 \
  --workers 4 \
  --threads 1 \
  --worker-class sync \
  --log-level critical \
  --access-logfile /dev/null \
  --error-logfile /dev/null
```
litestar
```
export PYTHONHASHSEED=0
export PYTHONDONTWRITEBYTECODE=1
export PYTHONWARNINGS=ignore
export LITESTAR_WARN_IMPLICIT_SYNC_TO_THREAD=0

gunicorn main:app \
  -k uvicorn.workers.UvicornWorker \
  -w 4 \
  --env UVICORN_LOOP=uvloop \
  --env UVICORN_HTTP=httptools \
  -b 127.0.0.1:8080 \
  --log-level critical \
  --access-logfile /dev/null \
  --error-logfile /dev/null
```
sanic
```
export PYTHONHASHSEED=0
export PYTHONDONTWRITEBYTECODE=1

sanic main.app \
  --host=127.0.0.1 \
  --port=8080 \
  --workers=4 \
  --no-access-log
```
