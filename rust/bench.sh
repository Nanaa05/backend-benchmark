#!/usr/bin/env bash
set -euo pipefail

PORT=8080
DURATION="5s"
THREAD_LIST="4 8 12 16"
CONNECTION_LIST="64 128 256 512 1024 2048"

RESULTS_DIR="/app/results"
RESULTS_FILE="$RESULTS_DIR/results.csv"
SUMMARY_FILE="$RESULTS_DIR/summary.txt"

mkdir -p "$RESULTS_DIR"

FRAMEWORKS=(
  "actix:actix-bench"
  "axum:axum-bench"
  "hyper:hyper-bench"
  "poem:poem-bench"
  "rocket:rocket-bench"
  "salvo:salvo-bench"
  "warp:warp-bench"
)

cat > /tmp/post.lua <<'EOF'
wrk.method = "POST"
wrk.body = '{"message":"hello benchmark"}'
wrk.headers["Content-Type"] = "application/json"
EOF

echo "framework,endpoint,method,threads,connections,duration,requests_per_sec,avg_latency" > "$RESULTS_FILE"

run_wrk() {
  local framework="$1"
  local endpoint="$2"
  local method="$3"
  local threads="$4"
  local connections="$5"
  local url="http://127.0.0.1:${PORT}${endpoint}"
  local output rps latency

  echo "[$framework] $method $endpoint | threads=$threads connections=$connections"

  if [[ "$method" == "POST" ]]; then
    output="$(wrk -t"$threads" -c"$connections" -d"$DURATION" -s /tmp/post.lua "$url")"
  else
    output="$(wrk -t"$threads" -c"$connections" -d"$DURATION" "$url")"
  fi

  rps="$(echo "$output" | awk '/Requests\/sec:/ {print $2}')"
  latency="$(echo "$output" | awk '/Latency/ {print $2; exit}')"

  echo "$framework,$endpoint,$method,$threads,$connections,$DURATION,$rps,$latency" >> "$RESULTS_FILE"
}

for item in "${FRAMEWORKS[@]}"; do
  IFS=":" read -r framework bin <<< "$item"

  echo
  echo "=================================================="
  echo "Benchmarking $framework"
  echo "=================================================="

  ROCKET_ADDRESS=127.0.0.1 ROCKET_PORT="$PORT" "/app/bin/$bin" &
  SERVER_PID=$!

  trap 'kill "$SERVER_PID" 2>/dev/null || true' EXIT

  for i in {1..100}; do
    if curl -fsS "http://127.0.0.1:${PORT}/ping" >/dev/null 2>&1; then
      break
    fi
    sleep 0.1
  done

  for threads in $THREAD_LIST; do
    for connections in $CONNECTION_LIST; do
      run_wrk "$framework" "/ping" "GET" "$threads" "$connections"
      run_wrk "$framework" "/json" "GET" "$threads" "$connections"
      run_wrk "$framework" "/echo" "POST" "$threads" "$connections"
    done
  done

  kill "$SERVER_PID"
  wait "$SERVER_PID" 2>/dev/null || true
  trap - EXIT

  sleep 1
done

{
  echo "BEST OVERALL"
  tail -n +2 "$RESULTS_FILE" | sort -t, -k7 -nr | head -n 1

  echo
  echo "BEST PER LIBRARY"
  tail -n +2 "$RESULTS_FILE" \
    | sort -t, -k1,1 -k7 -nr \
    | awk -F, '!seen[$1]++'

  echo
  echo "BEST PER ENDPOINT"
  tail -n +2 "$RESULTS_FILE" \
    | sort -t, -k2,2 -k7 -nr \
    | awk -F, '!seen[$2]++'
} | tee "$SUMMARY_FILE"

echo
echo "Results saved to:"
echo "$RESULTS_FILE"
echo "$SUMMARY_FILE"
