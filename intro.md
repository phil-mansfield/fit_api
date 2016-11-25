# A Short Introduciton to the `fit` Package

The `fit` package is meant to serve two types of people: people who want to
think a lot about their curve fits and people who don't want to think about
their curve fits at all.

If you fall into the second camp and don't care at all, here is what using the
fit library will look more or less like this every time:

```go
x, y, yerr := // Data (provided by user).

parabola := func(param []float64, x float64) float64 {
    return param[0] + param[1]*x + param[2]*x*x
}

guess := []fit.ParameterGuess{ {X: 2.0}, {X: 1.0}, {X: -3.0} }
out, err := fit.Curve.YError(parabola, guess, x, y, yerr)
if err != nil { panic("Fit did not converge.") }

fmt.Printf("Values: %.4g\n", out.Values)
fmt.Printf("Errors: %.4g\n", out.Errors)
```

You provide exactly three things: a function, a (potentially very bad)
guess for the correct parameters to that function, and your data.

If you fall into the first camp and care too much, you can manually supply
probability distribution functions to an affine-invariant MCMC sampler and
do whatever analysis you want on the posterior distribution for your parameters:

```go
myComplicatedPDF := func(param []float64) float64 {
    // Bunch of math goes here
}

guesses := []fit.ParameterGuess{ {X: 0.0}, {X: -2.0}, {X: 2.0}, {X: 17.5} }
sampler := NewSampler(myComplicatedPDF, guesses)
samples, err := sampler.Samples()
if err != nil { panic("Sampler did not converge.") }

// Do analysis on your samples
```

If you fall somewhere in between and want to do things that are only moderately
complicated, `fit` will let you do them without much pain. For example, suppose
that you wanted to fit a line to points with both x and y errors, some of your
points exist only as upper bounds, and you know that the y intercept can't be
larger than 7:

```go
x, y, xerr, yerr := // Data (provided by user)
xup, yup := // Locations of upper bounds (provided by user)

guesses := []ParameterGuess{
    {X: 1.0}.UpperLimit(7),
    {X: 3.0},
}
out, err := Line.XYErrors(guesses, x, y, xerr, yerr, fit.YUpperBounds(xup, yup))
if err != nil { panic("Fit did not converge.") }

fmt.Printf("Values: %.4g\n", out.Values)
fmt.Printf("Errors: %.4g\n", out.Errors)
```
