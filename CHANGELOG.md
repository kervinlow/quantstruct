# Change Log
All notable changes to this project will be documented in this file. This change log follows the conventions of [keepachangelog.com](http://keepachangelog.com/).

## [Unreleased]

## 0.0.1 - 2016-06-30
### Added
- Brief usage instruction to _README.md_.
- MIT License clauses to all source files.
- Analytical pricer methods for financial options in the analytical package:
  - `BV2002`: Bos and Vandermark (2002) pricing model.
  - `GK1983`: Garman and Kohlhagen (1983) pricing model.
  - `A1982`: Asay (1982) pricing model.
  - `B1976`: Black (1976) pricing model.
  - `M1973`: Merton (1973) pricing model.
  - `BS1973`: Black and Scholes (1973) pricing model.
  - `GBSM`: Generalized Black Scholes Merton pricing model.
- PricingResult: Struct for storing the results returned from the pricers.
- Equity dividend types, methods, and convenience functions:
  - `DestructDiv`: Function that destructures a Div struct.
  - `MakeDivList`: Function that builds a DivList slice from two slice or array
                   arguments.
  - `AddDiv`: Method for adding a dividend to a DivList slice.
  - `RemainingDivs`: Method for getting the remaining dividends in a DivList slice.
  - `NextDiv`: Method for getting the first dividend in a DivList slice.
  - `DivList`: A slice type that represents a list of discrete dividends.
  - `Div`: A struct type that represents a discrete dividend.
- Financial option types in the options package:
  - `Call`
  - `Put`
- Math functions (essentially wrappers to external math libraries) in the math
  package:
  - `PDF`: Returns the Probability Density Function for the standard Normal
           Distribution.
  - `CDF`: Returns the Cumulative Distribution Function for the standard Normal
           Distribution.

[Unreleased]: https://github.com/kervinlow/quantstruct/compare/0.0.2...HEAD
[0.0.2]: https://github.com/kervinlow/quantstruct/compare/0.0.1...0.0.2
