FROM rust:1.88-bookworm AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    git clang make ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY Cargo.toml Cargo.lock ./
COPY actix ./actix
COPY axum ./axum
COPY hyper ./hyper
COPY poem ./poem
COPY rocket ./rocket
COPY salvo ./salvo
COPY warp ./warp

RUN cargo build --release --workspace


FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates curl git build-essential procps unzip \
    && git clone https://github.com/wg/wrk.git /tmp/wrk \
    && make -C /tmp/wrk \
    && cp /tmp/wrk/wrk /usr/local/bin/wrk \
    && rm -rf /tmp/wrk \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/target/release /app/bin
COPY bench.sh /app/bench.sh

RUN chmod +x /app/bench.sh

CMD ["/app/bench.sh"]