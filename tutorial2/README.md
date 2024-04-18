# Constants, Variables and Basic Data Types in Go

## Constants

In Go, constants are declared like variables, but with the `const` keyword. Constants can be character, string, boolean, or numeric values. Constants cannot be declared using the `:=` syntax.

```go
const Pi = 3.14
```


## Variables

Variables in Go are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.

```go
var x int = 1
y := 2 // Short variable declarations, infer the type based on the assigned value
```

## Basic Data Types

Go's basic types are:

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32, represents a Unicode code point

float32 float64

complex64 complex128
```

The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.

## Zero Values

Variables declared without an explicit initial value are given their zero value. The zero value is:

- `0` for numeric types,
- `false` for the boolean type, and
- `""` (the empty string) for strings.

### Checkout the code

 - [main.go](main.go)