package cli

import (
	"fmt"
	"os"

	"weber-br/internal/drop"
)

func Run(args []string) int {
	if len(args) == 0 {
		return runServe([]string{})
	}
	switch args[0] {
	case "weber":
		return runWeber(args[1:])
	case "u-crit":
		return runUCrit(args[1:])
	case "example":
		return runExample(args[1:])
	case "serve":
		return runServe(args[1:])
	case "help", "-h", "--help":
		printHelp()
		return 0
	case "version":
		fmt.Println("weber-br 1.0.0")
		return 0
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", args[0])
		printHelp()
		return 2
	}
}

func printHelp() {
	fmt.Println(`weber-br: droplet Weber breakup calculator

Usage:
  weber-br                          start HTTP server on :8080
  weber-br weber -rho_g 1.2 -rho_l 1000 -mu_l 0.001 -sigma 0.072 -U 12 -d 0.002
  weber-br u-crit -rho_g 1.2 -rho_l 1000 -mu_l 0.001 -sigma 0.072 -d 0.002
  weber-br example -file example/rain-drop.json
  weber-br serve -addr :8080

HTTP:
  POST /api/weber   {"rho_g":1.2,"rho_l":1000,"mu_l":0.001,"sigma":0.072,"U":12,"d":0.002}
  POST /api/u-crit  {"rho_g":1.2,"rho_l":1000,"mu_l":0.001,"sigma":0.072,"d":0.002}`)
}

func fail(err error) int {
	fmt.Fprintln(os.Stderr, "error:", err)
	return 1
}

func runWeber(args []string) int {
	fs := flagSet("weber")
	rhoG := fs.Float64("rho_g", 1.2, "gas density")
	rhoL := fs.Float64("rho_l", 1000, "liquid density")
	muL := fs.Float64("mu_l", 0.001, "liquid viscosity")
	sigma := fs.Float64("sigma", 0.072, "surface tension")
	u := fs.Float64("U", 12, "relative velocity")
	d := fs.Float64("d", 0.002, "diameter")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	in := drop.Input{
		Fluids: drop.Fluids{RhoGas: *rhoG, RhoLiquid: *rhoL, MuLiquid: *muL, Sigma: *sigma},
		U:      *u, D: *d,
	}
	result, err := drop.Evaluate(in)
	if err != nil {
		return fail(err)
	}
	return printJSON(result)
}

func runUCrit(args []string) int {
	fs := flagSet("u-crit")
	rhoG := fs.Float64("rho_g", 1.2, "gas density")
	rhoL := fs.Float64("rho_l", 1000, "liquid density")
	muL := fs.Float64("mu_l", 0.001, "liquid viscosity")
	sigma := fs.Float64("sigma", 0.072, "surface tension")
	d := fs.Float64("d", 0.002, "diameter")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	f := drop.Fluids{RhoGas: *rhoG, RhoLiquid: *rhoL, MuLiquid: *muL, Sigma: *sigma}
	oh, err := drop.OhnesorgeNumber(f, *d)
	if err != nil {
		return fail(err)
	}
	criticalU, err := drop.CriticalVelocity(f, *d, drop.CriticalWeber(oh))
	if err != nil {
		return fail(err)
	}
	return printJSON(map[string]interface{}{"U_crit": criticalU, "We_crit": drop.CriticalWeber(oh), "Oh": oh})
}

func runExample(args []string) int {
	fs := flagSet("example")
	file := fs.String("file", "example/rain-drop.json", "scenario JSON file")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	scenario, err := drop.LoadScenario(*file)
	if err != nil {
		return fail(err)
	}
	result, err := drop.RunScenario(scenario)
	if err != nil {
		return fail(err)
	}
	return printJSON(result)
}

func loadExample(path string) (drop.Scenario, error) {
	return drop.LoadScenario(path)
}

func ExamplePaths() []string {
	return drop.ScenarioPaths()
}
