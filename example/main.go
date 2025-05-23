package main

import (
	"fmt"
	"os"

	-"github.com/utsav56/uargs"
)

func main() {
	args := []uargs.ArgDef{
		{
			Name: "index", Short: "i", Usage: "Shows item at particular given index",
			NumArgs: 1, Required: false, Type: args_parser.Int,
		},
		{
			Name: "query", Short: "q", Usage: "Query multiple strings",
			NumArgs: 3, Required: false, Type: args_parser.String,
		},
	}

	parser := uargs.NewParser(args)
	parsed, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		fmt.Println(parser.Usage())
		os.Exit(1)
	}

	fmt.Println(parsed)
}
