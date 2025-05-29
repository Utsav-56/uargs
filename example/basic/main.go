// This example demonstrates basic usage of the github.com/utsav-56/uargs library
package main

import (
	"fmt"
	"os" // When used in your own projects, import from the repository:

	// "github.com/yourusername/github.com/utsav-56/uargs"
	//
	// For local development:
	"github.com/utsav-56/uargs"
)

func main() {
	// Define command-line arguments
	args := []uargs.ArgDef{
		{
			Name:     "input", // Long form (--input)
			Short:    "i",     // Short form (-i)
			Usage:    "Input file path",
			Required: true, // This argument must be provided
			Type:     uargs.String,
		},
		{
			Name:  "output",
			Short: "o",
			Usage: "Output file path (optional)",
			Type:  uargs.String, // Default type is also string
		},
		{
			Name:  "verbose",
			Short: "v",
			Usage: "Enable verbose output",
			Type:  uargs.String, // Can be used as a flag with no value
		},
	}

	// Create a new parser
	parser := uargs.NewParser(args)

	// Parse command-line arguments and handle errors
	parsed, err := parser.Parse()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(parser.Usage()) // Print formatted usage information
		os.Exit(1)
	}

	// Access the parsed arguments (with type assertions)
	inputPath := parsed["input"].(string)
	fmt.Printf("Input file: %s\n", inputPath)

	// Check if optional arguments were provided
	if output, ok := parsed["output"]; ok {
		fmt.Printf("Output will be written to: %s\n", output.(string))
	} else {
		fmt.Println("No output path specified, using default")
	}

	// Check for flag argument
	if _, ok := parsed["verbose"]; ok {
		fmt.Println("Verbose mode enabled")
	}
}

/* Example usage:

   # With required input only
   go run main.go --input data.txt

   # With short option format
   go run main.go -i data.txt -o result.txt

   # With all options
   go run main.go --input data.txt --output result.txt --verbose

   # Missing required argument will show an error
   go run main.go --output result.txt
   Error: missing required argument --input
   Usage:
     --input      -i    Input file path
     --output     -o    Output file path (optional)
     --verbose    -v    Enable verbose output
*/
