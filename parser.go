package uargs

// Package github.com/utsav-56/uargs provides a simple command-line argument parser for Go.
// It allows you to define arguments with their names, short names, usage descriptions,
// type checking, and validation. The package supports string, integer, and floating-point
// argument types, as well as required arguments and conditional requirements.
//
// Basic usage example:
//
//	args := []github.com/utsav-56/uargs.ArgDef{
//		{Name: "input", Short: "i", Usage: "Input file", Required: true, Type: github.com/utsav-56/uargs.String},
//		{Name: "verbose", Short: "v", Usage: "Enable verbose output", Type: github.com/utsav-56/uargs.String},
//	}
//
//	parser := github.com/utsav-56/uargs.NewParser(args)
//	parsed, err := parser.Parse()
//	if err != nil {
//		fmt.Println(err)
//		fmt.Println(parser.Usage())
//		os.Exit(1)
//	}
//
//	inputFile := parsed["input"].(string)

import (
	_ "errors"
	"fmt"
	"os"
	_ "reflect"
	"strconv"
	"strings"
)

// ArgType represents the data type of an argument value
type ArgType string

const (
	// String indicates the argument value should be treated as a string
	String ArgType = "string"
	// Int indicates the argument value should be parsed as an integer
	Int ArgType = "int"
	// Float indicates the argument value should be parsed as a floating-point number
	Float ArgType = "float"
)

// ArgDef defines the properties of a command-line argument
type ArgDef struct {
	// Name is the long name of the argument (used with --)
	Name string
	// Short is the single-character short name of the argument (used with -)
	Short string
	// Usage is a description of the argument for help text
	Usage string
	// NumArgs is the number of values expected for this argument (default: 1)
	NumArgs int
	// Required indicates whether the argument must be provided
	Required bool
	// OptionalIfGiven makes this argument optional if any of the listed arguments are provided
	OptionalIfGiven []string
	// AcceptOverArgs allows accepting more values than specified by NumArgs
	AcceptOverArgs bool
	// Type specifies the data type of the argument value (String, Int, or Float)
	Type ArgType
}

// Parser represents a command-line argument parser
type Parser struct {
	defs        map[string]ArgDef      // Maps argument names to their definitions
	shortToLong map[string]string      // Maps short names to their corresponding long names
	parsed      map[string]interface{} // Stores parsed argument values
}

// NewParser creates a new Parser with the provided argument definitions
//
// Example:
//
//	args := []github.com/utsav-56/uargs.ArgDef{
//		{Name: "config", Short: "c", Usage: "Config file path", Type: github.com/utsav-56/uargs.String},
//	}
//	parser := github.com/utsav-56/uargs.NewParser(args)
func NewParser(args []ArgDef) *Parser {
	defs := make(map[string]ArgDef)
	shortToLong := make(map[string]string)
	for _, arg := range args {
		if arg.NumArgs == 0 {
			arg.NumArgs = 1
		}
		defs[arg.Name] = arg
		if arg.Short != "" {
			shortToLong[arg.Short] = arg.Name
		}
	}
	return &Parser{defs, shortToLong, make(map[string]interface{})}
}

// Parse parses command-line arguments and returns a map of argument names to their values.
// It validates required arguments, checks for duplicates, and handles type conversions.
//
// Example:
//
//	parsed, err := parser.Parse()
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	// Access a string argument
//	inputFile := parsed["input"].(string)
//
//	// Access an integer argument
//	count, ok := parsed["count"]
//	if ok {
//		countValue := count.(int)
//	}
func (p *Parser) Parse() (map[string]interface{}, error) {
	argv := os.Args[1:]
	used := make(map[string]bool)

	for i := 0; i < len(argv); i++ {
		arg := argv[i]
		if strings.HasPrefix(arg, "--") {
			name := arg[2:]
			if def, ok := p.defs[name]; ok {
				if used[name] {
					return nil, fmt.Errorf("duplicate argument --%s", name)
				}
				used[name] = true
				val, err := p.collectArgs(argv, &i, def)
				if err != nil {
					return nil, err
				}
				p.parsed[name] = val
			} else {
				return nil, fmt.Errorf("unknown argument --%s", name)
			}
		} else if strings.HasPrefix(arg, "-") {
			short := arg[1:]
			if len(short) > 1 {
				return nil, fmt.Errorf("invalid short argument usage: -%s", short)
			}
			if name, ok := p.shortToLong[short]; ok {
				if used[name] {
					return nil, fmt.Errorf("duplicate argument -%s/--%s", short, name)
				}
				used[name] = true
				def := p.defs[name]
				val, err := p.collectArgs(argv, &i, def)
				if err != nil {
					return nil, err
				}
				p.parsed[name] = val
			} else {
				return nil, fmt.Errorf("unknown short argument -%s", short)
			}
		} else {
			return nil, fmt.Errorf("unexpected token %s", arg)
		}
	}

	for name, def := range p.defs {
		if def.Required && p.parsed[name] == nil {
			optional := false
			for _, opt := range def.OptionalIfGiven {
				if used[opt] {
					optional = true
					break
				}
			}
			if !optional {
				return nil, fmt.Errorf("missing required argument --%s", name)
			}
		}
	}

	return p.parsed, nil
}

// collectArgs collects argument values from the command-line arguments.
// It handles multi-value arguments and type conversion based on the argument definition.
// This is an internal function used by the Parse method.
func (p *Parser) collectArgs(argv []string, i *int, def ArgDef) (interface{}, error) {
	args := []string{}
	for j := 0; j < def.NumArgs && *i+1 < len(argv); j++ {
		next := argv[*i+1]
		if strings.HasPrefix(next, "-") {
			break
		}
		*i++
		args = append(args, next)
	}
	if !def.AcceptOverArgs && len(args) > def.NumArgs {
		return nil, fmt.Errorf("too many arguments for --%s", def.Name)
	}

	switch def.Type {
	case Int:
		ints := []int{}
		for _, s := range args {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("--%s expects int, got '%s'", def.Name, s)
			}
			ints = append(ints, n)
		}
		if len(ints) == 1 {
			return ints[0], nil
		}
		return ints, nil
	case Float:
		floats := []float64{}
		for _, s := range args {
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return nil, fmt.Errorf("--%s expects float, got '%s'", def.Name, s)
			}
			floats = append(floats, f)
		}
		if len(floats) == 1 {
			return floats[0], nil
		}
		return floats, nil
	default:
		if len(args) == 1 {
			return args[0], nil
		}
		return args, nil
	}
}

// Usage generates a formatted help text showing all defined arguments with their
// names, short options, and usage descriptions. This is helpful for displaying
// to users when invalid arguments are provided or when help is requested.
//
// Example:
//
//	if err != nil {
//		fmt.Println(err)
//		fmt.Println(parser.Usage())
//		os.Exit(1)
//	}
func (p *Parser) Usage() string {
	var b strings.Builder
	b.WriteString("Usage:\n")
	for _, def := range p.defs {
		b.WriteString(fmt.Sprintf("  --%-10s -%s	%s\n", def.Name, def.Short, def.Usage))
	}
	return b.String()
}
