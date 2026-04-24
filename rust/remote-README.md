# Remote 2-VM Benchmark

Use this when server and load generator are on different VMs.

## Server VM

Run one framework only:

```bash
cd ~/backend-benchmark/rust
./remote-server.sh actix
```

Available libraries:
```bash
actix
axum
hyper
poem
rocket
salvo
warp
```

## Client VM
install wrk:
```bash
sudo apt update
sudo apt install -y git build-essential unzip curl
git clone https://github.com/wg/wrk.git /tmp/wrk
make -C /tmp/wrk
sudo cp /tmp/wrk/wrk /usr/local/bin/wrk
```

Run benchmark:
```bash
cd ~/backend-benchmark/rust
./remote-client.sh actix <SERVER_PRIVATE_IP>
```

Example:
```bash
./remote-client.sh actix 10.0.0.123
```

Output:
```bash
bench-results/actix.csv
bench-results/actix-summary.txt
```

Make executable:
```bash
chmod +x remote-server.sh remote-client.sh
```