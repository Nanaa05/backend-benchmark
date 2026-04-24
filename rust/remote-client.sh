#!/usr/bin/env bash
set -euo pipefail

LIB="${1:?usage: ./remote-client.sh <library> <server-private-ip>}"
SERVER_IP="${2:?usage: ./remote-client.sh <library> <server-private-ip>}"

PORT=8080
DURATION="5s"
THREAD_LIST="4 8 12 16"
CONNECTION_LIST="64 128 256 512 1024 2048"

RESULTS_DIR="./bench-results"
RESULTS_FILE="$RESULTS_DIR/${LIB}.csv"
SUMMARY_FILE="$RESULTS_DIR/${LIB}-summary.txt"

mkdir -p "$RESULTS_DIR"

cat > /tmp/post.lua <<'EOF'
wrk.method = "POST"
wrk.body = '{"message":"hello benchmark"}'
wrk.headers["Content-Type"] = "application/json"
EOF

echo "library,endpoint,method,threads,connections,duration,requests_per_sec,avg_latency" > "$RESULTS_FILE"

run_wrk() {
  local endpoint="$1"
  local method="$2"
  local threads="$3"
  local connections="$4"
  local url="http://${SERVER_IP}:${PORT}${endpoint}"

  echo "[$LIB] $method $endpoint | threads=$threads connections=$connections"

  if [[ "$method" == "POST" ]]; then
    output="$(wrk -t"$threads" -c"$connections" -d"$DURATION" -s /tmp/post.lua "$url")"
  else
    output="$(wrk -t"$threads" -c"$connections" -d"$DURATION" "$url")"
  fi

  rps="$(echo "$output" | awk '/Requests\/sec:/ {print $2}')"
  latency="$(echo "$output" | awk '/Latency/ {print $2; exit}')"

  echo "$LIB,$endpoint,$method,$threads,$connections,$DURATION,$rps,$latency" >> "$RESULTS_FILE"
}

curl -fsS "http://${SERVER_IP}:${PORT}/ping" >/dev/null

for threads in $THREAD_LIST; do
  for connections in $CONNECTION_LIST; do
    run_wrk "/ping" "GET" "$threads" "$connections"
    run_wrk "/json" "GET" "$threads" "$connections"
    run_wrk "/echo" "POST" "$threads" "$connections"
  done
done

{
  echo "BEST OVERALL FOR $LIB"
  tail -n +2 "$RESULTS_FILE" | sort -t, -k7 -nr | head -n 1

  echo
  echo "BEST PER ENDPOINT FOR $LIB"
  tail -n +2 "$RESULTS_FILE" \
    | sort -t, -k2,2 -k7 -nr \
    | awk -F, '!seen[$2]++'
} | tee "$SUMMARY_FILE"

echo
echo "Saved:"
echo "$RESULTS_FILE"
echo "$SUMMARY_FILE"
