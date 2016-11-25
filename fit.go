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
	UnknownErrors(x, y []float64) (Output, error)

	YErrors(x, y, yerr []float64) (Output, error)
	XErrors(x, y, xerr []float64) (Output, error)
 	XYErrors(x, y, xerr, yerr []float64) (Output, error)

	YErrorsAndScatter(x, y, yerr []float64) (Output, error)
	XErrorsAndScatter(x, y, xerr []float64) (Output, error)
 	XYErrorsAndScatter(x, y, xerr, yerr []float64) (Output, error)
}
