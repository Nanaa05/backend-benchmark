#!/usr/bin/env bash
set -euo pipefail

LIB="${1:?usage: ./remote-server.sh <actix|axum|hyper|poem|rocket|salvo|warp>}"

case "$LIB" in
  actix) PKG="actix-bench" ;;
  axum) PKG="axum-bench" ;;
  hyper) PKG="hyper-bench" ;;
  poem) PKG="poem-bench" ;;
  rocket) PKG="rocket-bench" ;;
  salvo) PKG="salvo-bench" ;;
  warp) PKG="warp-bench" ;;
  *) echo "Unknown library: $LIB"; exit 1 ;;
esac

cargo build --release -p "$PKG"

echo "Running $LIB server on 0.0.0.0:8080"
ROCKET_ADDRESS=0.0.0.0 ROCKET_PORT=8080 "./target/release/$PKG"
