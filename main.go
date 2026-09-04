package main

import (
	"GOCups/utils"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var logger = utils.NewPrint()

func banner() {
	cb := color.New(color.FgCyan, color.Bold).SprintFunc()
	c := color.New(color.FgCyan).SprintFunc()

	version := "2026.1.1"
	name := "Enumeración de CUPS"
	author := "Daniel Vargas"

	fmt.Printf("%s: %s\n", cb(name), c(version))
	fmt.Printf("%s: %s\n\n", cb("Autor"), c(author))
}

func main() {
	banner()

	host := ""
	port := 631
	help := false

	flag.StringVar(&host, "host", "", "IP del host")
	flag.IntVar(&port, "port", 631, "Puerto de CUPS.")
	flag.BoolVar(&help, "help", false, "Imprime la ayuda")

	flag.Usage = func() {
		logger.Info("[Uso]: %s IP PORT\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if help == true {
		flag.PrintDefaults()
	}
	if host == "" {
		flag.PrintDefaults()
		return
	}

	cups := NewClient()

	err := cups.EnumPrinters(host, port)

	if err != nil {
		logger.Error(err)
	}
}
