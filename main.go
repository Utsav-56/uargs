package main

import (
	"fmt"
	"os"
	"uargs/args_parser"
)

func main() {
	args := []args_parser.ArgDef{
		{
			Name: "index", Short: "i", Usage: "Shows item at particular given index",
			NumArgs: 1, Required: false, Type: args_parser.Int,
		},
		{
			Name: "query", Short: "q", Usage: "Query multiple strings",
			NumArgs: 3, Required: false, Type: args_parser.String,
		},
	}

	parser := args_parser.NewParser(args)
	parsed, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		fmt.Println(parser.Usage())
		os.Exit(1)
	}

	fmt.Println(parsed)
}
