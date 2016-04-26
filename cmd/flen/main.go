package main

import (
	"flag"
	"fmt"
	"github.com/lafolle/flen"
	"os"
)

var (
	pkg                          string
	bucketSize                   int
	inclTests                    bool
	lenLowerLimit, lenUpperLimit int
	max                          int
)

// init sets clas and flag package.
func init() {
	flag.BoolVar(&inclTests, "t", false, "include tests files")
	flag.IntVar(&bucketSize, "bs", 5, "bucket size (natural number)")
	flag.IntVar(&lenLowerLimit, "l", 0, "min length (inclusive)")
	flag.IntVar(&lenUpperLimit, "u", flen.Sentinel, "max length (exclusive)")
	flag.IntVar(&max, "m", 0, "error if any function longer than m")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage: flen [options] <pkg>\n")
		flag.PrintDefaults()
	}
}

func rangeAsked(ll, ul int) bool { return ll != 0 || ul != flen.Sentinel }

func main() {

	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}

	pkg = flag.Args()[0]

	if pkg == "" {
		flag.Usage()
		return
	}

	flenOptions := &flen.Options{
		IncludeTests: inclTests,
		BucketSize:   bucketSize,
	}
	flens, pkgpath, err := flen.GenerateFuncLens(pkg, flenOptions)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Full pkg path: %s\n\n", pkgpath)

	if max > 0 {
		for _, y := range flens {
			if y.Size > max {
				fmt.Printf("Function %s exceeds limit %d: %d", y.Name, max, y.Size)
				os.Exit(-1)
			}
		}
	}
	if !rangeAsked(lenLowerLimit, lenUpperLimit) {
		zeroLenFuncs := flens.GetZeroLenFuncs()
		if len(zeroLenFuncs) > 0 {
			fmt.Println("0 len funcs")
			zeroLenFuncs.Print()
			fmt.Println()
		}

		extImplFuncs := flens.GetExternallyImplementedFuncs()
		if len(extImplFuncs) > 0 {
			fmt.Println("Externally implemented funcs")
			extImplFuncs.Print()
			fmt.Println()
		}
	} else {
		rangeFlens := flens.Query(lenLowerLimit, lenUpperLimit)
		if len(rangeFlens) > 0 {
			fmt.Printf("Functions with length in range [%d, %d)\n", lenLowerLimit, lenUpperLimit)
			rangeFlens.Print()
			fmt.Println()
		}
	}
	fmt.Println("Histogram")
	flens.DisplayHistogram()

	os.Exit(0)
}
