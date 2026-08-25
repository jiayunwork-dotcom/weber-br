package drop

import "fmt"

type Fluids struct {
	RhoGas    float64 `json:"rho_g"`
	RhoLiquid float64 `json:"rho_l"`
	MuLiquid  float64 `json:"mu_l"`
	Sigma     float64 `json:"sigma"`
}

type Input struct {
	Fluids
	U float64 `json:"U"`
	D float64 `json:"d"`
}

type Result struct {
	We        float64 `json:"We"`
	Oh        float64 `json:"Oh"`
	WeCrit    float64 `json:"We_crit"`
	Mode      string  `json:"mode"`
	Breakup   bool    `json:"breakup"`
	U         float64 `json:"U"`
	D         float64 `json:"d"`
	CriticalU float64 `json:"U_crit"`
}

func (f Fluids) String() string {
	return fmt.Sprintf("rho_g=%.4g rho_l=%.4g mu_l=%.4g sigma=%.4g",
		f.RhoGas, f.RhoLiquid, f.MuLiquid, f.Sigma)
}

func (i Input) String() string {
	return fmt.Sprintf("%s U=%.4g d=%.4g", i.Fluids, i.U, i.D)
}

func (r Result) IsFinite() bool {
	return !bad(r.We) && !bad(r.Oh) && !bad(r.WeCrit) && !bad(r.CriticalU)
}

func bad(value float64) bool {
	return value != value || value > 1e300 || value < -1e300
}
