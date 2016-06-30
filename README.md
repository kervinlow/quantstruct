# quantstruct

A Golang library for pricing financial instruments.

**Important Note**: This is a work-in-progress.

Please refer to the [Change Log]
(https://github.com/kervinlow/quantstruct/blob/master/CHANGELOG.md)
for more information on the changes made to the library.

## Usage

Generally, you will need to import the library's packages in your program as follows:
```go
import (
    "github.com/kervinlow/quantstruct/options"
    "github.com/kervinlow/quantstruct/pricers/analytical"
)
```

If you want to use the analytical option pricers that can handle discrete dividends, you will also need to import an additional package as follows:
```go
import (
    "github.com/kervinlow/quantstruct/options"
    "github.com/kervinlow/quantstruct/pricers/analytical"
    "github.com/kervinlow/quantstruct/equity"
)
```

Please refer to the comments in the library's source files (e.g. 
[blackscholesmerton.go]
(https://github.com/kervinlow/quantstruct/blob/master/pricers/analytical/blackscholesmerton.go)
 etc.) for more usage instructions.

## License

```
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
```
