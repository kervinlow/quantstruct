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
Package math provides the mathematical functions required by the
Quantstruct Library.
*/
package math

import "github.com/datastream/probab/dst"

/*
=================
Wrapper Functions
=================
*/

/*
CDF returns the Cumulative Distribution Function of the standard
Normal Distribution at x.
*/
func CDF(x float64) float64 {
	return dst.NormalCDFAt(0.0, 1.0, x)
}

/*
PDF returns the Probability Density Function of the standard
Normal Distribution at x.
*/
func PDF(x float64) float64 {
	return dst.NormalPDFAt(0.0, 1.0, x)
}
