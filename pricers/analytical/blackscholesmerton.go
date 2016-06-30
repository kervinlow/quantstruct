/*
******************************************************************************
MIT License

Copyright (c) 2016 Kervin Low

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
******************************************************************************
*/

package analytical

import (
	"fmt"
	. "github.com/kervinlow/quantstruct/equity"
	. "github.com/kervinlow/quantstruct/math"
	. "github.com/kervinlow/quantstruct/options"
	. "math"
)

/*
===============
Types of Errors
===============
*/

/*
The error ErrPricing is returned when a pricing error has occurred.
*/
type ErrPricing string

func (e ErrPricing) Error() string {
	return fmt.Sprintf("%s", string(e))
}

/*
==========================================================================
Provides the closed-form Black-Scholes-Merton family of pricing models for
valuing finanical options.
==========================================================================
*/

/*
--------------------------------------------------------------------------
GBSM -- Generalized Black Scholes Merton pricing model

Description:
A method that computes the theoretical value and greeks of a financial
option, and saves the computed results in the fields of the ModelOutputs
receiver. It returns the error ErrPricing if a pricing error has occurred;
otherwise, it returns nil.

Usage:
var out analytical.ModelOutputs
err := out.GBSM(ot, s, k, t, v, r, b)

Arguments:
ot option type (either options.Call or options.Put from
   the options package)
s  spot price of the underlying instrument
k  strike price of the option
t  time to expiry of the option
v  volatility of the underlying instrument
r  risk-free rate
b  cost of carry
--------------------------------------------------------------------------
*/
func (out *ModelOutputs) GBSM(ot OptionType, s float64, k float64, t float64, v float64, r float64, b float64) error {
	// Compute d1 and d2 as specified by the pricing model.
	d1 := (Log(s/k) + ((b + v*v/2.0) * t)) / (v * Sqrt(t))
	d2 := d1 - (v * Sqrt(t))
	var ch [6]chan float64 // allocate 6 channels for the channel array ch
	// Create each channel, assign it to the channel array ch, and pass each channel to a Goroutine.
	for n := 0; n < len(ch); n++ {
		ch[n] = make(chan float64)
		switch n {
		case 0:
			go getGBSMValue(ch[n], ot, d1, d2, s, k, t, r, b)
		case 1:
			go getGBSMDelta(ch[n], ot, d1, t, r, b)
		case 2:
			go getGBSMTheta(ch[n], ot, d1, d2, s, k, t, v, r, b)
		case 3:
			go getGBSMRho(ch[n], ot, d2, k, t, r)
		case 4:
			go getGBSMGamma(ch[n], ot, d1, s, t, v, r, b)
		case 5:
			go getGBSMVega(ch[n], ot, d1, s, t, r, b)
		}
	}
	// Receive the computed result from each channel, and store it in the ModelOutputs receiver.
	for i := range ch {
		for result := range ch[i] {
			switch i {
			case 0:
				out.Value = result
			case 1:
				out.Delta = result
			case 2:
				out.Theta = result
			case 3:
				out.Rho = result
			case 4:
				out.Gamma = result
			case 5:
				out.Vega = result
			}
		}
	}
	// Check for pricing error.
	if IsNaN(out.Value) || IsInf(out.Value, 0) || IsNaN(out.Delta) || IsInf(out.Delta, 0) ||
		IsNaN(out.Gamma) || IsInf(out.Gamma, 0) || IsNaN(out.Vega) || IsInf(out.Vega, 0) ||
		IsNaN(out.Theta) || IsInf(out.Theta, 0) || IsNaN(out.Rho) || IsInf(out.Rho, 0) {
		return ErrPricing("Pricing error has occurred.")
	}
	// Scaling some of the Greeks based on market conventions.
	out.Vega = out.Vega / 100.0
	out.Theta = out.Theta / 365.0
	out.Rho = out.Rho / 100.0
	return nil
}

/*
getGBSMValue is a Goroutine and is not accessible outside this package.
It computes the theoretical value of a financial option using the
Generalized Black Scholes Merton pricing model.
*/
func getGBSMValue(c chan float64, ot OptionType, d1 float64, d2 float64, s float64, k float64, t float64, r float64, b float64) {
	switch ot {
	case Call:
		c <- (s * Exp((b-r)*t) * CDF(d1)) - (k * Exp((-r)*t) * CDF(d2))
	case Put:
		c <- (k * Exp((-r)*t) * CDF(-d2)) - (s * Exp((b-r)*t) * CDF(-d1))
	}
	close(c)
}

/*
getGBSMDelta is a Goroutine and is not accessible outside this package.
It computes the Delta of a financial option using the Generalized Black
Scholes Merton pricing model.
*/
func getGBSMDelta(c chan float64, ot OptionType, d1 float64, t float64, r float64, b float64) {
	switch ot {
	case Call:
		c <- Exp((b-r)*t) * CDF(d1)
	case Put:
		c <- Exp((b-r)*t) * (CDF(d1) - 1.0)
	}
	close(c)
}

/*
getGBSMTheta is a Goroutine and is not accessible outside this package.
It computes the Theta of a financial option using the Generalized Black
Scholes Merton pricing model.
*/
func getGBSMTheta(c chan float64, ot OptionType, d1 float64, d2 float64, s float64, k float64, t float64, v float64, r float64, b float64) {
	switch ot {
	case Call:
		c <- (((-s) * Exp((b-r)*t) * PDF(d1) * v) / (2.0 * Sqrt(t))) -
			((b - r) * s * Exp((b-r)*t) * CDF(d1)) -
			(r * k * Exp((-r)*t) * CDF(d2))
	case Put:
		c <- (((-s) * Exp((b-r)*t) * PDF(d1) * v) / (2.0 * Sqrt(t))) +
			((b - r) * s * Exp((b-r)*t) * CDF(-d1)) +
			(r * k * Exp((-r)*t) * CDF(-d2))
	}
	close(c)
}

/*
getGBSMRho is a Goroutine and is not accessible outside this package.
It computes the Rho of a financial option using the Generalized Black
Scholes Merton pricing model.
*/
func getGBSMRho(c chan float64, ot OptionType, d2 float64, k float64, t float64, r float64) {
	switch ot {
	case Call:
		c <- t * k * Exp((-r)*t) * CDF(d2)
	case Put:
		c <- (-t) * k * Exp((-r)*t) * CDF(-d2)
	}
	close(c)
}

/*
getGBSMGamma is a Goroutine and is not accessible outside this package.
It computes the Gamma of a financial option using the Generalized Black
Scholes Merton pricing model.
*/
func getGBSMGamma(c chan float64, ot OptionType, d1 float64, s float64, t float64, v float64, r float64, b float64) {
	c <- (PDF(d1) * Exp((b-r)*t)) / (s * v * Sqrt(t))
	close(c)
}

/*
getGBSMVega is a Goroutine and is not accessible outside this package.
It computes the Vega of a financial option using the Generalized Black
Scholes Merton pricing model.
*/
func getGBSMVega(c chan float64, ot OptionType, d1 float64, s float64, t float64, r float64, b float64) {
	c <- s * Exp((b-r)*t) * PDF(d1) * Sqrt(t)
	close(c)
}

/*
----------------------------------------------------------------------
BS1973 -- Black and Scholes (1973) pricing model

Description:
A method that computes the theoretical value and greeks of an option
on a stock that does not pay dividend, and saves the computed results
in the fields of the ModelOutputs receiver. It returns the error
ErrPricing if a pricing error has occurred; otherwise, it returns nil.

Usage:
var out analytical.ModelOutputs
err := out.BS1973(ot, s, k, t, v, r)

Arguments:
ot option type (either options.Call or options.Put from
   the options package)
s  spot price of the underlying instrument
k  strike price of the option
t  time to expiry of the option
v  volatility of the underlying instrument
r  risk-free rate
----------------------------------------------------------------------
*/
func (out *ModelOutputs) BS1973(ot OptionType, s float64, k float64, t float64, v float64, r float64) error {
	err := out.GBSM(ot, s, k, t, v, r, r)
	if err != nil {
		return err
	}
	return nil
}

/*
-------------------------------------------------------------------------
M1973 -- Merton (1973) pricing model

Description:
A method that computes the theoretical value and greeks of an option on
a stock (or stock index) that pays a known continuous dividend yield, and
saves the computed results in the fields of the ModelOutputs receiver. It
returns the error ErrPricing if a pricing error has occurred; otherwise,
it returns nil.

Usage:
var out analytical.ModelOutputs
err := out.M1973(ot, s, k, t, v, r, q)

Arguments:
ot option type (either options.Call or options.Put from
   the options package)
s  spot price of the underlying instrument
k  strike price of the option
t  time to expiry of the option
v  volatility of the underlying instrument
r  risk-free rate
q  continuous dividend yield
-------------------------------------------------------------------------
*/
func (out *ModelOutputs) M1973(ot OptionType, s float64, k float64, t float64, v float64, r float64, q float64) error {
	err := out.GBSM(ot, s, k, t, v, r, r-q)
	if err != nil {
		return err
	}
	return nil
}

/*
--------------------------------------------------------------------------
B1976 -- Black (1976) pricing model

Description:
A method that computes the theoretical value and greeks of an option on a
forward or futures contract, and saves the computed results in the fields
of the ModelOutputs receiver. It returns the error ErrPricing if a pricing
error has occurred; otherwise, it returns nil.

Usage:
var out analytical.ModelOutputs
err := out.B1976(ot, f, k, t, v, r)

Arguments:
ot option type (either options.Call or options.Put from
   the options package)
f  forward price of the underlying instrument
k  strike price of the option
t  time to expiry of the option
v  volatility of the underlying instrument
r  risk-free rate
--------------------------------------------------------------------------
*/
func (out *ModelOutputs) B1976(ot OptionType, f float64, k float64, t float64, v float64, r float64) error {
	err := out.GBSM(ot, f, k, t, v, r, 0.0)
	if err != nil {
		return err
	}
	return nil
}

/*
------------------------------------------------------------------------
A1982 -- Asay (1982) pricing model

Description:
A method that computes the theoretical value and greeks of an option
on a futures contract where the premium is fully margined, and saves
the computed results in the fields of the ModelOutputs receiver. It
returns the error ErrPricing if a pricing error has occurred; otherwise,
it returns nil.

Usage:
var out analytical.ModelOutputs
err := out.A1982(ot, f, k, t, v)

Arguments:
ot option type (either options.Call or options.Put from
   the options package)
f  forward price of the underlying instrument
k  strike price of the option
t  time to expiry of the option
v  volatility of the underlying instrument
------------------------------------------------------------------------
*/
func (out *ModelOutputs) A1982(ot OptionType, f float64, k float64, t float64, v float64) error {
	err := out.GBSM(ot, f, k, t, v, 0.0, 0.0)
	if err != nil {
		return err
	}
	return nil
}

/*
--------------------------------------------------------------------------
GK1983 -- Garman and Kohlhagen (1983)

Description:
A method that computes the theoretical value and greeks of a currency
option, and saves the computed results in the fields of the ModelOutputs
receiver. It returns the error ErrPricing if a pricing error has occurred;
otherwise, it returns nil.

Usage:
var out analytical.ModelOutputs
err := out.GK1983(ot, s, k, t, v, rd, rf)

Arguments:
ot option type (either options.Call or options.Put from
   the options package)
s  spot price of the underlying instrument
k  strike price of the option
t  time to expiry of the option
v  volatility of the underlying instrument
rd domestic risk-free rate (base currency)
rf foreign risk-free rate (quoted currency)
--------------------------------------------------------------------------
*/
func (out *ModelOutputs) GK1983(ot OptionType, s float64, k float64, t float64, v float64, rd float64, rf float64) error {
	err := out.GBSM(ot, s, k, t, v, rd, rd-rf)
	if err != nil {
		return err
	}
	return nil
}

/*
--------------------------------------------------------------------
BV2002 -- Bos and Vandermark (2002) pricing model

Description:
A method that computes the theoretical value and greeks of an option
on a stock that pays discrete cash dividends, and saves the computed
results in the fields of the ModelOutputs receiver. It returns the
error ErrPricing if a pricing error has occurred; otherwise, it
returns nil.

Usage:
var out analytical.ModelOutputs
err := out.BV2002(ot, s, k, t, v, r, dl)

Arguments:
ot option type (either options.Call or options.Put from
   the options package)
s  spot price of the underlying instrument
k  strike price of the option
t  time to expiry of the option
v  volatility of the underlying instrument
r  risk-free rate
dl discrete dividend list (the equity.DivList type
   in the equity package)
--------------------------------------------------------------------
*/
func (out *ModelOutputs) BV2002(ot OptionType, s float64, k float64, t float64, v float64, r float64, dl DivList) error {
	adjS := s - divNear(r, t, dl)
	adjK := k + (divFar(r, t, dl) * Exp(r*t))
	err := out.BS1973(ot, adjS, adjK, t, v, r)
	if err != nil {
		return err
	}
	return nil
}

/*
divNear is an unexported function that returns the weighted
present value of 'near' dividends.
*/
func divNear(r float64, tte float64, dl DivList) float64 {
	t, d, dn := 0.0, 0.0, 0.0
	for _, v := range dl {
		if t, d = DestructDiv(v); t <= tte {
			dn += (tte - t) / tte * d * Exp((-r)*t)
		}
	}
	return dn
}

/*
divFar is an unexported function that returns the weighted
present value of 'far' dividends.
*/
func divFar(r float64, tte float64, dl DivList) float64 {
	t, d, df := 0.0, 0.0, 0.0
	for _, v := range dl {
		if t, d = DestructDiv(v); t <= tte {
			df += t / tte * d * Exp((-r)*t)
		}
	}
	return df
}
