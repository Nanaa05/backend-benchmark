```
cargo run -p axum-bench --release
cargo run -p actix-bench --release
cargo run -p warp-bench --release
cargo run -p rocket-bench --release
cargo run -p salvo-bench --release
```

docker
```
mkdir -p bench-results
docker compose build
docker compose run --rm rust-web-bench
cat bench-results/summary.txt
```