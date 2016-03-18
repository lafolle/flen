/*
	Program to compute length of functions and analyze them.
*/
package main

import (
	"flag"
	"fmt"
	"github.com/lafolle/flen"
	"os"
	"text/tabwriter"
)

var (
	pkg        string
	bucketSize int
	inclTests  bool
	streakChar string
)

// init sets clas and flag package.
func init() {
	flag.StringVar(&pkg, "pkg", "", "pkg to parse")
	flag.BoolVar(&inclTests, "t", false, "include tests")
	flag.IntVar(&bucketSize, "bs", 5, "bucket size (natural number)")
	flag.StringVar(&streakChar, "c", "|", "streak char used to display histogram")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage: flen -pkg <package> [options]\n")
		flag.PrintDefaults()
	}
}

func main() {

	flag.Parse()
	if pkg == "" {
		flag.Usage()
		return
	}

	flenOptions := flen.Options{
		IncludeTests: inclTests,
		BucketSize:   bucketSize,
		StreakChar:   streakChar,
	}
	flens, zeroLenFuncs, err := flen.GenerateFuncLens(pkg, flenOptions)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(zeroLenFuncs) != 0 {
		fmt.Printf("\nWe have %d zero len funcs!", len(zeroLenFuncs))
		zeroTabw := new(tabwriter.Writer)
		zeroTabw.Init(os.Stdout, 0, 4, 0, '\t', 0)
		fmt.Fprintln(zeroTabw)
		for _, zlf := range zeroLenFuncs {
			fmt.Fprintln(zeroTabw, fmt.Sprintf("%s\t%s", zlf.Name, zlf.Filepath))
		}
		zeroTabw.Flush()
	}

	flens.DisplayHistogram()
	os.Exit(0)
}
