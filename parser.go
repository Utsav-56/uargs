package uargs

// Package uargs provides a simple command-line argument parser for Go.
// It allows you to define arguments with their names, short names, usage descriptions,

import (
	_ "errors"
	"fmt"
	"os"
	_ "reflect"
	"strconv"
	"strings"
)

type ArgType string

const (
	String ArgType = "string"
	Int    ArgType = "int"
	Float  ArgType = "float"
)

type ArgDef struct {
	Name            string
	Short           string
	Usage           string
	NumArgs         int
	Required        bool
	OptionalIfGiven []string
	AcceptOverArgs  bool
	Type            ArgType
}

type Parser struct {
	defs        map[string]ArgDef
	shortToLong map[string]string
	parsed      map[string]interface{}
}

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

func (p *Parser) Usage() string {
	var b strings.Builder
	b.WriteString("Usage:\n")
	for _, def := range p.defs {
		b.WriteString(fmt.Sprintf("  --%-10s -%s	%s\n", def.Name, def.Short, def.Usage))
	}
	return b.String()
}
