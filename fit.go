package fit_api

// Output contains the output of a fitting routine.
type Output struct {
	// Values contains the best fit parameter values. Formally, they correspond
	// to the 50th percentile of the marginalized posterior distribution for
	// each parameter.
	Values []float64
	// Errors contains the errors on the best fit parameter values. Formally,
	// they correspond to half the width of the 68% contours of the marginalized
	// posterior distribution for each parameter.
	Errors []float64
	// Covariance is the covaraince matrix for the fit parameter. Formally,
	// Covariance[i][j] = <(<p_i> - p_i)(<p_j> - p_j)>, meaning that
	Covariance [][]float64
}

var (
	// Line is a Model which fits against straight lines with the form
	// y = p[0] + p[1]*x. All error bars provided to Line methods are assumed
	// to represent the standard deviaiton of normal distributions. Line is
	// provided as a special case over the more general Curve Model to provide
	// access to Model methods with X error bars.
	//
	// Example usage:
	// out, err := Line.XYErrors(p0, x, y, xerr, yerr)
	Line Model
	// PowerLaw is a model which fits against power laws with the form
	// p[0]*x^p[1]. All error bars provided to to PowerLaw methods are assumed
	// to represent the logarithmic standard deviation of (base 10) log-
	// normal distributions (If you do not have log-normal error distributions,
	// you'll have to use Curve). PowerLaw is provided as a special case over
	// the more general Curve Model to provide access to Model methods with X
	// error bars.
	//
	// Example usage:
	// out, err := PowerLaw.XErrors(p0, x, y, xerr)
	PowerLaw Model
)

// Func is a one dimensional fitting function. x is the point that the function
// is being evaluated at, and param is the current values of the fitting
// parameters.
type Func func(x float64, param []float64) float64

// Curve returns a Model which fits against an arbitrary 1D function. All error
// bars to Curve methods are assumed to be the standard deviation of normal
// distributions. Curve does not implement any methods with X error bars and
// calling them will result in a runtime panic.
//
// Example usage:
// out, err := Curve(f).YErrors(p0, x, y, yerr)
func Curve(f Func) Model { return nil }

// Model is a collection of fitting functions for a particular functional form.
// It has methods of two types: fitting methods and modification methods. All
// fitting methods take an initial parameter, x and y points, and error bars,
// if appropriate. Modification methods can be used to add information to the
// model, like limits.
//
// Example usage:
//
// Line.XYErrors(x, y, xerr, yerr)
//
// Curve.YLowerLimits(xLower, yLower).YErrors(x, y, yerr)
type Model interface {
	// UnknownErrors will return an additional parameter representing the
	// best-fit intrisic scatter.
	UnknownErrors(p0 []Parameter, x, y []float64) (Output, error)

	YErrors(p0 []Parameter, x, y, yerr []float64) (Output, error)
	XErrors(p0 []Parameter, x, y, xerr []float64) (Output, error)
 	XYErrors(p0 []Parameter, x, y, xerr, yerr []float64) (Output, error)

	// YErrorsAndScatter will return an additional parameter representing the
	// best-fit intrisic scatter. Other "AndScatter" methods behave analogously.
	YErrorsAndScatter(p0 []Parameter, x, y, yerr []float64) (Output, error)
	XErrorsAndScatter(p0 []Parameter, x, y, xerr []float64) (Output, error)
 	XYErrorsAndScatter(p0 []Parameter, x, y, xerr, yerr []float64) (Output, error)
 
	// YUpperLimits creates a new Model which contains upper limits in the
	// y direction at the specified points. Other "Limits" methods behave
	// analogously.
	YUpperLimits(x, y []float64) Model
	YLowerLimits(x, y []float64) Model
	XUpperLimits(x, y []float64) Model
	XLowerLimits(x, y []float64) Model
}

var (
	LinePDF ModelPDF
	PowerLawPDF ModelPDF
)

func CurvePDF(f Func) ModelPDF { return nil }

// ModelPDF is directly analogous to Model, except its methods return LogPDF
// instances that can be passed to NewSampler.
//
// (I'll write  more complete documentation once the Model API is stable).
type ModelPDF interface {
	UnknownErrors(p0 []Parameter, x, y []float64) LogPDF

	YErrors(p0 []Parameter, x, y, yerr []float64) LogPDF
	XErrors(p0 []Parameter, x, y, xerr []float64) LogPDF
 	XYErrors(p0 []Parameter, x, y, xerr, yerr []float64) LogPDF

	YErrorsAndScatter(p0 []Parameter, x, y, yerr []float64) LogPDF
	XErrorsAndScatter(p0 []Parameter, x, y, xerr []float64) LogPDF
 	XYErrorsAndScatter(p0 []Parameter, x, y, xerr, yerr []float64) LogPDF

	YUpperLimits(x, y []float64) ModelPDF
	YLowerLimits(x, y []float64) ModelPDF
	XUpperLimits(x, y []float64) ModelPDF
	XLowerLimits(x, y []float64) ModelPDF
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

func GelmanRubin(lim float64) ConvergenceTest
func Steps(n int) ConvergenceTest

func (sampler *Sampler) Samples(tests ...ConvergenceTest) ([][]float64, error) {
	return nil, nil
}

func (sampler *Sampler) RawSamples(steps int) ([][]float64, error) {
	return nil, nil
}

func (sampler *Sampler) AcceptanceRatio() float64 { return 0.0 }
func (sampler *Sampler) MostLikelyParameters() []float64 { return nil }
