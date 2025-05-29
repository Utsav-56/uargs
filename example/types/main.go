// This example demonstrates using different argument types in github.com/utsav-56/uargs
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
	// Define arguments with different types
	args := []uargs.ArgDef{
		{
			Name:  "text",
			Short: "t",
			Usage: "Text input (string type)",
			Type:  uargs.String,
		},
		{
			Name:  "count",
			Short: "c",
			Usage: "Count value (integer type)",
			Type:  uargs.Int,
		},
		{
			Name:  "rate",
			Short: "r",
			Usage: "Rate value (float type)",
			Type:  uargs.Float,
		},
	}

	parser := uargs.NewParser(args)
	parsed, err := parser.Parse()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(parser.Usage())
		os.Exit(1)
	}

	// Access each type with proper type assertion
	if text, ok := parsed["text"]; ok {
		// String type
		textValue := text.(string)
		fmt.Printf("Text: %s (type: %T)\n", textValue, textValue)
	}

	if count, ok := parsed["count"]; ok {
		// Integer type - automatically converted from string by the parser
		countValue := count.(int)
		fmt.Printf("Count: %d (type: %T)\n", countValue, countValue)

		// The parser will return an error if this argument is provided with a non-integer value
	}

	if rate, ok := parsed["rate"]; ok {
		// Float type - automatically converted from string by the parser
		rateValue := rate.(float64)
		fmt.Printf("Rate: %.2f (type: %T)\n", rateValue, rateValue)

		// The parser will return an error if this argument is provided with a non-float value
	}
}

/* Example usage:

   # Providing values of different types
   go run main.go --text hello --count 42 --rate 3.14

   # If a non-integer is provided for the count argument, an error will be shown
   go run main.go --count abc
   Error: --count expects int, got 'abc'

   # If a non-float is provided for the rate argument, an error will be shown
   go run main.go --rate xyz
   Error: --rate expects float, got 'xyz'
*/
