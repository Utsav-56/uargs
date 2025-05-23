# uargs Examples

This directory contains examples demonstrating different features of the `uargs` library:

## Basic Example

Located in `basic/main.go` - Shows the fundamental usage of the library with required and optional arguments.

Key features demonstrated:

-   Creating basic argument definitions
-   Required vs. optional arguments
-   String-type arguments
-   Checking for flag arguments
-   Generating usage help

To run:

```bash
cd basic
go run main.go --input example.txt --output result.txt --verbose
# Or with short options:
go run main.go -i example.txt -o result.txt -v
```

## Type Validation Example

Located in `types/main.go` - Shows how to use different argument types and automatic type validation.

Key features demonstrated:

-   String arguments
-   Integer arguments with automatic conversion and validation
-   Float arguments with automatic conversion and validation
-   Type-safe access to parsed arguments

To run:

```bash
cd types
go run main.go --text hello --count 42 --rate 3.14
```

## Advanced Features Example

Located in `advanced/main.go` - Shows more advanced features of the library.

Key features demonstrated:

-   Multi-value arguments (accepting multiple values for a single argument)
-   Conditional requirements (arguments only required under certain conditions)
-   Working with arrays of values

To run:

```bash
cd advanced
go run main.go --tags red blue green --coords 10.5 20.3 --file data.txt
```

## Running with Errors

Try running the examples with invalid inputs to see how error handling works:

```bash
# Missing a required argument
cd basic
go run main.go --output result.txt

# Invalid type (non-integer value for int type)
cd types
go run main.go --count hello

# Wrong number of argument values
cd advanced
go run main.go --coords 1.0
```
