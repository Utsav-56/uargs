# uargs - Simple Command-Line Argument Parser for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/utsav56/uargs.svg)](https://pkg.go.dev/github.com/utsav56/uargs)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`uargs` is a lightweight, intuitive command-line argument parsing library for Go. It simplifies the process of defining, parsing, and validating command-line arguments with features like required arguments, type checking, and usage help generation.

## Table of Contents

-   [Installation](#installation)
-   [Quick Start](#quick-start)
-   [Core Concepts](#core-concepts)
    -   [Argument Types](#argument-types)
    -   [Argument Definition](#argument-definition)
    -   [Parser](#parser)
-   [Examples](#examples)
    -   [Basic Usage](#basic-usage)
    -   [Required Arguments](#required-arguments)
    -   [Conditionally Required Arguments](#conditionally-required-arguments)
    -   [Multiple Arguments](#multiple-arguments)
    -   [Type Validation](#type-validation)
-   [API Reference](#api-reference)
    -   [ArgDef Struct](#argdef-struct)
    -   [Parser Methods](#parser-methods)
-   [Best Practices](#best-practices)
-   [Contributing](#contributing)
-   [License](#license)

## Installation

To install the `uargs` library, use `go get`:

```bash
go get github.com/utsav-56/github.com/utsav-56/uargs
```

## Quick Start

Here's a simple example to get you started:

```go
package main

import (
    "fmt"
    "os"

    "github.com/utsav56/uargs"
)

func main() {
    // Define your command-line arguments
    args := []uargs.ArgDef{
        {
            Name: "name",
            Short: "n",
            Usage: "Your name",
            Required: true,
            Type: uargs.String,
        },
        {
            Name: "age",
            Short: "a",
            Usage: "Your age",
            Type: uargs.Int,
        },
    }

    // Create a new parser with the defined arguments
    parser := uargs.NewParser(args)

    // Parse the command-line arguments
    parsed, err := parser.Parse()
    if err != nil {
        fmt.Println(err)
        fmt.Println(parser.Usage())
        os.Exit(1)
    }

    // Access the parsed arguments
    name := parsed["name"].(string)
    fmt.Printf("Hello, %s!\n", name)

    if age, ok := parsed["age"]; ok {
        fmt.Printf("You are %d years old.\n", age.(int))
    }
}
```

Run the program with:

```bash
go run main.go --name John --age 30
# Or using short options:
go run main.go -n John -a 30
```

## Core Concepts

### Argument Types

`uargs` supports the following argument types:

-   `String` - Text values (default)
-   `Int` - Integer values
-   `Float` - Floating-point values

### Argument Definition

Arguments are defined using the `ArgDef` struct, which includes:

-   `Name` - The long name of the argument (used with `--`)
-   `Short` - The short name of the argument (used with `-`)
-   `Usage` - Description of the argument for help text
-   `NumArgs` - Number of values expected (default: 1)
-   `Required` - Whether the argument is required
-   `OptionalIfGiven` - Makes the argument optional if specified arguments are provided
-   `AcceptOverArgs` - Whether to accept more arguments than specified by NumArgs
-   `Type` - The type of the argument (String, Int, Float)

### Parser

The `Parser` struct provides methods to:

-   Create a new parser with argument definitions
-   Parse command-line arguments
-   Generate usage help text

## Examples

### Basic Usage

```go
args := []uargs.ArgDef{
    {
        Name: "output",
        Short: "o",
        Usage: "Output file path",
        Type: uargs.String,
    },
}

parser := uargs.NewParser(args)
parsed, err := parser.Parse()
if err != nil {
    fmt.Println(err)
    fmt.Println(parser.Usage())
    os.Exit(1)
}

// Access the parsed output path if provided
if outputPath, ok := parsed["output"]; ok {
    fmt.Printf("Output will be written to: %s\n", outputPath.(string))
}
```

### Required Arguments

```go
args := []uargs.ArgDef{
    {
        Name: "input",
        Short: "i",
        Usage: "Input file (required)",
        Required: true,
        Type: uargs.String,
    },
}

// The parser will return an error if --input is not provided
```

### Conditionally Required Arguments

```go
args := []uargs.ArgDef{
    {
        Name: "format",
        Short: "f",
        Usage: "Output format",
        Type: uargs.String,
    },
    {
        Name: "template",
        Short: "t",
        Usage: "Template file (required when using --format)",
        Required: true,
        OptionalIfGiven: []string{"format"},
        Type: uargs.String,
    },
}

// --template is only required if --format is provided
```

### Multiple Arguments

```go
args := []uargs.ArgDef{
    {
        Name: "tags",
        Short: "t",
        Usage: "Specify up to 3 tags",
        NumArgs: 3,
        Type: uargs.String,
    },
}

// Can be used like: --tags tag1 tag2 tag3
// Access with: parsed["tags"].([]string)
```

### Type Validation

```go
args := []uargs.ArgDef{
    {
        Name: "count",
        Short: "c",
        Usage: "Number of iterations",
        Type: uargs.Int,
    },
    {
        Name: "threshold",
        Short: "t",
        Usage: "Threshold value",
        Type: uargs.Float,
    },
}

// The parser will validate types and return errors for invalid values
```

## API Reference

### ArgDef Struct

```go
type ArgDef struct {
    Name            string   // Long name (used with --)
    Short           string   // Short name (used with -)
    Usage           string   // Help text description
    NumArgs         int      // Number of values (default: 1)
    Required        bool     // Whether argument is required
    OptionalIfGiven []string // Makes argument optional if these args are given
    AcceptOverArgs  bool     // Accept more values than NumArgs
    Type            ArgType  // String, Int, or Float
}
```

### Parser Methods

#### NewParser

```go
func NewParser(args []ArgDef) *Parser
```

Creates a new argument parser with the specified argument definitions.

#### Parse

```go
func (p *Parser) Parse() (map[string]interface{}, error)
```

Parses the command-line arguments and returns a map of argument names to their values.

#### Usage

```go
func (p *Parser) Usage() string
```

Generates a formatted usage help text string.

## Best Practices

-   **Use Descriptive Names**: Choose clear, descriptive names for arguments.
-   **Provide Short Options**: Always provide short options for common arguments.
-   **Include Usage Descriptions**: Write helpful usage descriptions for all arguments.
-   **Validate Required Arguments**: Use the `Required` flag for mandatory arguments.
-   **Handle Errors Gracefully**: Always check for parsing errors and display the usage when errors occur.
-   **Type Safety**: Use type assertions carefully when accessing parsed values.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
