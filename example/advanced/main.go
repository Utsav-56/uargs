// This example demonstrates using multi-value arguments and conditional requirements
package main

import (
	"fmt"
	"os" // When used in your own projects, import from the repository:

	// "github.com/yourusername/uargs"
	//
	// For local development:
	"uargs"
)

func main() {
	// Define arguments with multiple values and conditional requirements
	args := []uargs.ArgDef{
		{
			Name:  "file",
			Short: "f",
			Usage: "Input file",
			Type:  uargs.String,
		},
		{
			Name:    "tags",
			Short:   "t",
			Usage:   "Specify up to 3 tags for categorization",
			NumArgs: 3, // This argument expects up to 3 values
			Type:    uargs.String,
		},
		{
			Name:    "coords",
			Short:   "c",
			Usage:   "X and Y coordinates (requires 2 float values)",
			NumArgs: 2, // This argument expects exactly 2 values
			Type:    uargs.Float,
		},
		{
			Name:            "template",
			Short:           "tpl",
			Usage:           "Template file (required only when --format is provided)",
			Required:        true,
			OptionalIfGiven: []string{"file"}, // Only required if "file" is not given
			Type:            uargs.String,
		},
	}

	parser := uargs.NewParser(args)
	parsed, err := parser.Parse()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(parser.Usage())
		os.Exit(1)
	}

	// Access a standard single-value argument
	if file, ok := parsed["file"]; ok {
		fmt.Printf("File: %s\n", file.(string))
	}

	// Access a multi-value string argument
	if tags, ok := parsed["tags"]; ok {
		// Multi-value arguments are returned as slices
		tagsArr := tags.([]string)
		fmt.Printf("Tags (%d provided): %v\n", len(tagsArr), tagsArr)

		// Access individual tag values
		for i, tag := range tagsArr {
			fmt.Printf("  Tag %d: %s\n", i+1, tag)
		}
	}

	// Access a multi-value float argument
	if coords, ok := parsed["coords"]; ok {
		// Multi-value float arguments are returned as []float64
		coordsArr := coords.([]float64)
		fmt.Printf("Coordinates: X=%.2f, Y=%.2f\n", coordsArr[0], coordsArr[1])
	}

	// Access conditionally required argument
	if template, ok := parsed["template"]; ok {
		fmt.Printf("Template: %s\n", template.(string))
	}
}

/* Example usage:

   # Providing multiple values for the "tags" argument
   go run main.go --tags red blue green

   # Providing coordinates (must be valid floats)
   go run main.go --coords 10.5 20.3

   # Using file argument makes template optional
   go run main.go --file data.txt

   # Not using file argument means template is required
   go run main.go
   Error: missing required argument --template

   # You can use the short format for all arguments
   go run main.go -f data.txt -t red blue green -c 10.5 20.3
*/
