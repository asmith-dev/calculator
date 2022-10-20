# Go Calculator
This is a simple calculator that I made in Go.

### Method
The calculator organizes operations using PEMDAS, with some exceptions.

### Operations
* Multiplication `*`
* Division `/`
* Addition `+`
* Subtraction `-` (interpreted as addition of a negative)
* No support for exponents currently

### Numbers
All numbers are interpreted as 64-bit floating point numbers.\
This is admittedly not very memory-efficient.\
Efficiency was not my purpose.

### Common syntax errors, with examples
* No whitespace, i.e. `2 + 2`
* No variables, i.e. `x+2`
* No double operators, i.e. `5+-2`
* No parenthetically implied multiplication, i.e. `5(5+2)` or `(3+2)(7-9)`