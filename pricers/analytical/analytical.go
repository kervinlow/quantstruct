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

/*
Package analytical provides the analytical pricers that can be used to
value financial instruments and their risks.

This is a multi-file package and is made up of the following source files:
  analytical.go          provides the common definitions that are used by
                         the other source files in the package;
  blackscholesmerton.go  provides the analytical pricers that belong to the
                         Black-Scholes-Merton family of pricing models.
*/
package analytical

/*
==================
Common Definitions
==================
*/

/*
ModelOutputs is the structure that holds the results returned by the pricing
methods defined in the package.
*/
type ModelOutputs struct {
	Value float64
	Delta float64
	Gamma float64
	Vega  float64
	Theta float64
	Rho   float64
}
