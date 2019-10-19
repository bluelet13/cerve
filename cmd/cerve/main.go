package main

import (
	"flag"
	"fmt"
	"ngen.co.jp/cermo/pkg/cerve"
	"os"
)

func main() {
	var option = cerve.Option{}
	flag.Parse()

	c, err := cerve.NewCerve(flag.Args(), option)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	v, err := c.Verify()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(v.JSON())
}
