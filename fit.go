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
	UnknownErrors(p0 []Parameter, x, y []float64) (Output, error)

	YErrors(p0 []Parameter, x, y, yerr []float64) (Output, error)
	XErrors(p0 []Parameter, x, y, xerr []float64) (Output, error)
 	XYErrors(p0 []Parameter, x, y, xerr, yerr []float64) (Output, error)

	YErrorsAndScatter(p0 []Parameter, x, y, yerr []float64) (Output, error)
	XErrorsAndScatter(p0 []Parameter, x, y, xerr []float64) (Output, error)
 	XYErrorsAndScatter(p0 []Parameter, x, y, xerr, yerr []float64) (Output, error)
}

var (
	LinePDF ModelPDF
	PowerLawPDF ModelPDF
)

func CurvePDF(f Func) ModelPDF { return nil }

type ModelPDF interface {
	UnknownErrors(p0 []Parameter, x, y []float64) LogPDF

	YErrors(p0 []Parameter, x, y, yerr []float64) LogPDF
	XErrors(p0 []Parameter, x, y, xerr []float64) LogPDF
 	XYErrors(p0 []Parameter, x, y, xerr, yerr []float64) LogPDF

	YErrorsAndScatter(p0 []Parameter, x, y, yerr []float64) LogPDF
	XErrorsAndScatter(p0 []Parameter, x, y, xerr []float64) LogPDF
 	XYErrorsAndScatter(p0 []Parameter, x, y, xerr, yerr []float64) LogPDF
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

type LogPDF func(param []float64) float64

type Sampler struct { }

type SamplerOption samplerOption
type samplerConfig struct { }
type samplerOption func(*samplerConfig)

func NewSampler(pdf LogPDF, p0 []Parameter, opts ...SamplerOption) *Sampler {
	return nil
}

func Walkers(n int) SamplerOption { return nil }
func Stretch(a float64) SamplerOption { return nil }
func Threads(n int) SamplerOption { return nil }
func StepGranularity(n int) SamplerOption { return nil }

type ConvergenceTest convergenceTest
type convergenceTest func(chains [][][]float64) bool

func (sampler *Sampler) Samples(tests ...ConvergenceTest) ([][]float64, error) {
	return nil, nil
}

func (sampler *Sampler) RawSamples(steps int) ([][]float64, error) {
	return nil, nil
}

func (sampler *Sampler) AcceptanceRatio() float64 { return 0.0 }
func (sampler *Sampler) MostLikelyParameters() []float64 { return nil }
