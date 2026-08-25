# weber-br

weber-br is a Go droplet breakup calculator. It computes the gas-side Weber
number `We = rho_g*U^2*d/sigma`, the Ohnesorge number
`Oh = mu_l/sqrt(rho_l*sigma*d)`, an Oh-corrected breakup threshold, the breakup
regime, and the critical relative velocity for a given droplet. The service is
available through HTTP JSON endpoints and CLI subcommands with no web page.

## Usage

Run the HTTP server:

```bash
go run . serve -addr :8080
```

Evaluate from the command line:

```bash
go run . weber -rho_g 1.2 -rho_l 1000 -mu_l 0.001 -sigma 0.072 -U 20 -d 0.002
go run . u-crit -rho_g 1.2 -rho_l 1000 -mu_l 0.001 -sigma 0.072 -d 0.002
```

Run the rain-drop example:

```bash
go run . example -file example/rain-drop.json
```

The 2 mm rain droplet at 20 m/s crosses the bag breakup threshold.

## HTTP API

```text
POST /api/weber   {"rho_g":1.2,"rho_l":1000,"mu_l":0.001,"sigma":0.072,"U":20,"d":0.002}
POST /api/u-crit  {"rho_g":1.2,"rho_l":1000,"mu_l":0.001,"sigma":0.072,"d":0.002}
GET  /health
```

Non-positive densities, diameter, or surface tension, or negative velocity
return an error body with HTTP 400.

## Regimes

The base thresholds are bag 12, multi-mode 50, sheet stripping 100, and
catastrophic 350. The bag threshold is raised linearly with Ohnesorge number.

## Code Layout

```text
internal/drop     Weber, Ohnesorge, regime, critical velocity
internal/spray    droplet geometry, scans, sweeps
internal/scale    velocity scaling tables and crossings
internal/table    result tables and CSV
internal/verify   U^2, sigma, diameter and mode checks
internal/server   HTTP handlers and JSON responses
internal/cli      subcommand parsing and terminal output
example/          offline scenario JSON files
```

## Build and Test

```bash
export GOTOOLCHAIN=local CGO_ENABLED=0
go build ./...
go test ./...
```

The Dockerfile builds the server binary and starts it on port 8080.
