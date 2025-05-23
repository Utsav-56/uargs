/*
Package uargs provides a simple, flexible command-line argument parser for Go applications.

This package allows developers to define, parse, and validate command-line arguments
with features such as:

  - Long (--name) and short (-n) argument formats
  - Required and optional arguments
  - Type validation (string, int, float)
  - Multi-value arguments
  - Conditional requirements
  - Usage help generation

Quick Start

Define your arguments:

	args := []uargs.ArgDef{
		{Name: "input", Short: "i", Usage: "Input file", Required: true, Type: uargs.String},
		{Name: "count", Short: "c", Usage: "Number of iterations", Type: uargs.Int},
		{Name: "verbose", Short: "v", Usage: "Enable verbose mode", Type: uargs.String},
	}

Create a parser and parse the arguments:

	parser := uargs.NewParser(args)
	parsed, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		fmt.Println(parser.Usage())
		os.Exit(1)
	}

Access the parsed values:

	inputFile := parsed["input"].(string)
	
	if count, ok := parsed["count"]; ok {
		iterations := count.(int)
		// Use iterations...
	}
	
	if _, ok := parsed["verbose"]; ok {
		// Verbose mode is enabled
	}

Working with Different Types

String arguments (default type):

	{Name: "file", Short: "f", Usage: "Input file", Type: uargs.String}
	// Accessed as: parsed["file"].(string)

Integer arguments with automatic conversion:

	{Name: "count", Short: "c", Usage: "Count value", Type: uargs.Int}
	// Accessed as: parsed["count"].(int)

Float arguments with automatic conversion:

	{Name: "rate", Short: "r", Usage: "Rate value", Type: uargs.Float}
	// Accessed as: parsed["rate"].(float64)

Multi-value Arguments

For arguments that accept multiple values:

	{Name: "tags", Short: "t", Usage: "Tags", NumArgs: 3, Type: uargs.String}
	// Set NumArgs to the number of values expected
	// Accessed as: parsed["tags"].([]string)

Best Practices

1. Always provide usage descriptions for your arguments
2. Use required flag for mandatory arguments
3. Always handle parsing errors and display usage information
4. Use type assertions cautiously
5. Provide both long and short forms for common arguments

For more examples and detailed documentation, see the examples directory.
*/
package uargs
