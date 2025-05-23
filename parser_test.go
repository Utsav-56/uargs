package uargs_test

import (
	"os"
	"testing"
	
	"uargs"
)

// Example_basic demonstrates basic usage of the uargs library
func Example_basic() {
	// Save original args and restore after test
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	
	// Simulate command-line arguments
	os.Args = []string{"app", "--input", "file.txt", "--verbose"}
	
	// Define argument definitions
	args := []uargs.ArgDef{
		{Name: "input", Short: "i", Usage: "Input file", Type: uargs.String},
		{Name: "output", Short: "o", Usage: "Output file", Type: uargs.String},
		{Name: "verbose", Short: "v", Usage: "Enable verbose mode", Type: uargs.String},
	}
	
	// Create a new parser and parse args
	parser := uargs.NewParser(args)
	parsed, err := parser.Parse()
	if err != nil {
		panic(err)
	}
	
	// Access the parsed arguments
	inputFile := parsed["input"].(string)
	
	// Check for optional arguments
	var outputFile string
	if output, ok := parsed["output"]; ok {
		outputFile = output.(string)
	} else {
		outputFile = "default.out"
	}
	
	// Check for flag argument
	verbose := false
	if _, ok := parsed["verbose"]; ok {
		verbose = true
	}
	
	// Output: file.txt
	// default.out
	// true
	println(inputFile)
	println(outputFile)
	println(verbose)
}

// Example_types demonstrates using different argument types
func Example_types() {
	// Save original args and restore after test
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	
	// Simulate command-line arguments
	os.Args = []string{"app", "--count", "42", "--rate", "3.14"}
	
	// Define arguments with different types
	args := []uargs.ArgDef{
		{Name: "count", Short: "c", Usage: "Count value", Type: uargs.Int},
		{Name: "rate", Short: "r", Usage: "Rate value", Type: uargs.Float},
	}
	
	// Create a new parser and parse args
	parser := uargs.NewParser(args)
	parsed, err := parser.Parse()
	if err != nil {
		panic(err)
	}
	
	// Access the integer argument with type assertion
	count := parsed["count"].(int)
	
	// Access the float argument with type assertion
	rate := parsed["rate"].(float64)
	
	// Output: 42
	// 3.14
	println(count)
	println(rate)
}

// Example_multiValue demonstrates using multi-value arguments
func Example_multiValue() {
	// Save original args and restore after test
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	
	// Simulate command-line arguments
	os.Args = []string{"app", "--tags", "red", "green", "blue"}
	
	// Define a multi-value argument
	args := []uargs.ArgDef{
		{Name: "tags", Short: "t", Usage: "Color tags", NumArgs: 3, Type: uargs.String},
	}
	
	// Create a new parser and parse args
	parser := uargs.NewParser(args)
	parsed, err := parser.Parse()
	if err != nil {
		panic(err)
	}
	
	// Access the multi-value argument (returns a slice)
	tags := parsed["tags"].([]string)
	
	// Output: [red green blue]
	// red
	println(tags)
	println(tags[0])
}

// TestParser tests the core functionality of the Parser
func TestParser(t *testing.T) {
	// Save original args and restore after test
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	
	// Test case 1: Basic argument parsing
	os.Args = []string{"app", "--input", "test.txt", "--count", "42"}
	
	args := []uargs.ArgDef{
		{Name: "input", Short: "i", Usage: "Input file", Type: uargs.String},
		{Name: "count", Short: "c", Usage: "Count value", Type: uargs.Int},
	}
	
	parser := uargs.NewParser(args)
	parsed, err := parser.Parse()
	
	if err != nil {
		t.Fatalf("Failed to parse valid arguments: %v", err)
	}
	
	// Verify string argument
	input, ok := parsed["input"]
	if !ok {
		t.Fatal("Missing 'input' argument in parsed results")
	}
	if input.(string) != "test.txt" {
		t.Errorf("Expected input='test.txt', got '%s'", input)
	}
	
	// Verify int argument
	count, ok := parsed["count"]
	if !ok {
		t.Fatal("Missing 'count' argument in parsed results")
	}
	if count.(int) != 42 {
		t.Errorf("Expected count=42, got %d", count)
	}
	
	// Test case 2: Required arguments
	os.Args = []string{"app", "--optional", "value"}
	
	args = []uargs.ArgDef{
		{Name: "required", Short: "r", Usage: "Required arg", Required: true, Type: uargs.String},
		{Name: "optional", Short: "o", Usage: "Optional arg", Type: uargs.String},
	}
	
	parser = uargs.NewParser(args)
	_, err = parser.Parse()
	
	if err == nil {
		t.Error("Expected error due to missing required argument, got nil")
	}
	
	// Test case 3: Type validation
	os.Args = []string{"app", "--number", "not-a-number"}
	
	args = []uargs.ArgDef{
		{Name: "number", Short: "n", Usage: "Number value", Type: uargs.Int},
	}
	
	parser = uargs.NewParser(args)
	_, err = parser.Parse()
	
	if err == nil {
		t.Error("Expected error due to invalid number format, got nil")
	}
}
