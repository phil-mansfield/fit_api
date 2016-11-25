package fit_api

type Output struct {
	Values, Errors []float64
	Covariance [][]float64
}

var (
	Line Model
	PowerLaw Model
)

type Func func(x float64, param []float64) (y float64)

func Curve(f Func) Model { return nil }

type Model interface {
	UnknownErrors(p []Parameter, x, y []float64) (Output, error)

	YErrors(p []Parameter, x, y, yerr []float64) (Output, error)
	XErrors(p []Parameter, x, y, xerr []float64) (Output, error)
 	XYErrors(p []Parameter, x, y, xerr, yerr []float64) (Output, error)

	YErrorsAndScatter(p []Parameter, x, y, yerr []float64) (Output, error)
	XErrorsAndScatter(p []Parameter, x, y, xerr []float64) (Output, error)
 	XYErrorsAndScatter(p []Parameter, x, y, xerr, yerr []float64) (Output, error)
}

type Parameter struct {
	V, S float64
	logPrior func(float64) float64
}

func (p Parameter) Freeze() Parameter { return p }
func (p Parameter) LowerLimit(lim float64) Parameter { return p }
func (p Parameter) UpperLimit(lim float64) Parameter { return p }
func (p Parameter) Limits(lower, upper float64) Parameter { return p }
func (p Parameter) LogPrior(pr func(float64) float64) Parameter { return p }

