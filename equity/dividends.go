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
Package equity provides the representations of the equity instrument and its
corporate actions.
*/
package equity

import "fmt"

/*
===============
Types of Errors
===============
*/

/*
The error ErrEmptyList is returned when a list is empty.
*/
type ErrEmptyList string

func (e ErrEmptyList) Error() string {
	return fmt.Sprintf("%s", string(e))
}

/*
The error ErrSlicesDiffLength is returned when the input slices
are of different lengths.
*/
type ErrSlicesDiffLength string

func (e ErrSlicesDiffLength) Error() string {
	return fmt.Sprintf("%s", string(e))
}

/*
-----------------
Discrete Dividend
-----------------
*/

/*
Div represents a discrete dividend and is expressed in the
form of (Time to the Dividend, Dividend Amount).

Usage (example):
var d = equity.Div{0.5, 4.0}
*/
type Div struct {
	TimeToDividend float64 // expressed as a year fraction
	Amount         float64
}

/*
--------------------------
List of Discrete Dividends
--------------------------
*/

/*
DivList represents a list of discrete dividends.

Usage (example):
var dl = equity.DivList(append(make([]equity.Div, 0), equity.Div{0.5, 4.0}, equity.Div{1.5, 4.0}))
*/
type DivList []Div

/*
-----------------------------------
Methods of a Discrete Dividend List
-----------------------------------
*/

/*
NextDiv returns the head (first element) of the discrete dividend list.

Usage (example):
var dl = equity.DivList(append(make([]equity.Div, 0), equity.Div{0.5, 4.0}, equity.Div{1.5, 4.0}))
var d, e = dl.NextDiv()
*/
func (dl DivList) NextDiv() (Div, error) {
	if len(dl) == 0 {
		return Div{0.0, 0.0}, ErrEmptyList("List is empty.")
	}
	return dl[0], nil
}

/*
RemainingDivs drops the first dividend in the discrete dividend list, and
returns the list of the remaining dividends.

Usage (example):
var dl1 = equity.DivList(append(make([]equity.Div, 0), equity.Div{0.5, 4.0}, equity.Div{1.5, 4.0}))
var dl2, e = dl1.RemainingDivs()
*/
func (dl DivList) RemainingDivs() (DivList, error) {
	if len(dl) == 0 {
		return DivList(nil), ErrEmptyList("List is empty.")
	}
	return dl[1:], nil
}

/*
AddDiv appends a dividend to the discrete dividend list.

Usage (example):
var dl1 = equity.DivList(append(make([]equity.Div, 0), equity.Div{0.5, 4.0}, equity.Div{1.5, 4.0}))
var dl2 = dl1.AddDiv(2.5, 4.0)
*/
func (dl DivList) AddDiv(t float64, a float64) DivList {
	return append(dl, Div{t, a})
}

/*
---------------------
Convenience Functions
---------------------
*/

/*
MakeDivList creates a discrete dividend list from two arrays representing
the times to the dividends and the corresponding dividend amounts.

Usage (example):
var dl, e = equity.MakeDivList([]float64{0.5, 1.5}, []float64{4.0, 4.0})
*/
func MakeDivList(t []float64, d []float64) (DivList, error) {
	dl := make(DivList, 0)
	if len(t) == len(d) {
		for i, v := range t {
			dl = dl.AddDiv(v, d[i])
		}
		return dl, nil
	}
	return dl, ErrSlicesDiffLength("The two slices given are of different length.")
}

/*
DestructDiv destructures a dividend into its two components:
the time to the dividend and the dividend amount.

Usage (example):
var t, d = DestructDiv(div)

where div is of the type equity.Div.
*/
func DestructDiv(d Div) (float64, float64) {
	return d.TimeToDividend, d.Amount
}
