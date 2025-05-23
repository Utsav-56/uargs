// Example demonstrates how to use the uargs library for parsing command-line arguments
package main

import (
	"fmt"
	"os"

	// When used in your own projects, import from the repository:
	// "github.com/yourusername/uargs"
	// 
	// For local development:
	"uargs"
)

func main() {
	// Define command-line arguments with various options and types
	args := []uargs.ArgDef{
		{
			Name:     "input",
			Short:    "i",
			Usage:    "Input file path",
			Required: true,
			Type:     uargs.String,
		},
		{
			Name:     "count",
			Short:    "c",
			Usage:    "Number of iterations",
			Type:     uargs.Int,
		},
		{
			Name:     "tags",
			Short:    "t",
			Usage:    "Up to 3 tags for categorization",
			NumArgs:  3,
			Type:     uargs.String,
		},
		{
			Name:            "threshold",
			Short:           "th",
			Usage:           "Threshold value (only required if count is provided)",
			Required:        true,
			OptionalIfGiven: []string{"count"},
			Type:            uargs.Float,
		},
		{
			Name:     "verbose",
			Short:    "v",
			Usage:    "Enable verbose output",
			Type:     uargs.String,
		},
	}

	// Create a parser with the defined arguments
	parser := uargs.NewParser(args)
	
	// Parse command-line arguments
	parsed, err := parser.Parse()
	if err != nil {
		// Show error and usage information if parsing fails
		fmt.Println(err)
		fmt.Println(parser.Usage())
		os.Exit(1)
	}

	// Access and use the parsed arguments with proper type assertions
	
	// String argument (required)
	inputFile := parsed["input"].(string)
	fmt.Printf("Input file: %s\n", inputFile)
	
	// Integer argument (optional)
	if count, ok := parsed["count"]; ok {
		countValue := count.(int)
		fmt.Printf("Will perform %d iterations\n", countValue)
	}
	
	// Multiple string arguments
	if tags, ok := parsed["tags"]; ok {
		tagList := tags.([]string)
		fmt.Println("Tags:", tagList)
	}
	
	// Float argument
	if threshold, ok := parsed["threshold"]; ok {
		thresholdValue := threshold.(float64)
		fmt.Printf("Using threshold: %.2f\n", thresholdValue)
	}
	
	// Flag argument
	if _, ok := parsed["verbose"]; ok {
		fmt.Println("Verbose mode enabled")
	}
	
	fmt.Println("\nAll parsed arguments:")
	fmt.Println(parsed)
}
