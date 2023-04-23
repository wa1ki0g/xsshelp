package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"xsshelp/lib"
)

var (
	u string
	f string
	t int
	h bool
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&u, "u", "", "a target url(Please add http or https)")
	flag.IntVar(&t, "t", runtime.NumCPU(), "thread Num")

	// 改变默认的 Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `xsshelp version: 1.0.0
Usage:  [-ut] [-u url] [-t thread] [-h help]

Options:
`)
	flag.PrintDefaults()
}

func main() {
	lib.Logo()
	flag.Parse()

	if h {
		flag.Usage()
	}

	res1 := strings.Contains(u, "http://")
	res2 := strings.Contains(u, "https://")

	if u == "" {

	} else {
		if !res1 && !res2 {
			fmt.Println(lib.Red("[-] Please add http or https for url !!!"))
			os.Exit(0)
		} else {
			fmt.Println(lib.Yellow("target: " + u))

			lib.Getscan(u, t)


		}
	}
}
