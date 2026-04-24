import json
import time
import os
from django.conf import settings
from django.core.handlers.wsgi import WSGIHandler
from django.http import HttpResponse, JsonResponse
from django.urls import path
from django.core.servers.basehttp import run


# --------------------
# Views
# --------------------

def ping(request):
    return HttpResponse("ok", content_type="text/plain")


def json_endpoint(request):
    return JsonResponse({
        "status": "ok",
        "service": "django",
        "timestamp": int(time.time()),
    })


def echo(request):
    if request.method != "POST":
        return JsonResponse({"error": "method not allowed"}, status=405)

    if request.content_type != "application/json":
        return JsonResponse({"error": "expected application/json"}, status=400)

    try:
        body = json.loads(request.body)
    except json.JSONDecodeError:
        return JsonResponse({"error": "invalid json"}, status=400)

    return JsonResponse(body)


# --------------------
# URL routing
# --------------------

urlpatterns = [
    path("ping", ping),
    path("json", json_endpoint),
    path("echo", echo),
]


# --------------------
# Minimal Django setup
# --------------------

if not settings.configured:
    settings.configure(
        DEBUG=False,
        SECRET_KEY="benchmark",
        ROOT_URLCONF=__name__,
        ALLOWED_HOSTS=["127.0.0.1"],
        MIDDLEWARE=[],          # IMPORTANT: no middleware
        INSTALLED_APPS=[],      # IMPORTANT: no apps
        USE_TZ=False,
    )


application = WSGIHandler()


# --------------------
# Entry point
# --------------------

if __name__ == "__main__":
    os.environ.setdefault("DJANGO_SETTINGS_MODULE", "__main__")

    run(
        addr="127.0.0.1",
        port=8080,
        wsgi_handler=application,
        ipv6=False,
        threading=True,   # Django’s built-in threaded server
    )

