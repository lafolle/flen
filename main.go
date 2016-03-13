/*
	Program to compute length of functions in go package.

	Usage:
	flen -pkg <package name>

	eg.
	flen -pkg go/parser
*/
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"runtime"
	"strings"
	"text/tabwriter"
)

const (
	SENTINEL = 1000000
)

var (
	goroot    string
	pkgpath   string
	pkg       string
	inclTests bool
	gopath    string = os.Getenv("GOPATH")

	defaultBucketSize = 5
	bucketSize        int
	pathSeparator     string
	envSeparator      string
)

type FuncLen struct {
	Name           string
	Size           int
	Filepath       string
	Lbrace, Rbrace int
}

type ZeroLenFunc struct {
	FuncLen
}

// init() sets clas and set platform dependent vars.
func init() {
	flag.StringVar(&pkg, "pkg", "", "pkg to parse")
	flag.BoolVar(&inclTests, "t", false, "include tests")

	platform := runtime.GOOS
	switch platform {
	case "linux", "darwin", "freebsd":
		goroot = "/usr/local/go"
		pathSeparator = "/"
		envSeparator = ":"
	case "windows":
		goroot = "c:\\Go"
		pathSeparator = "\\"
		envSeparator = ";"
	default:
		panic(fmt.Sprintf("platform not supported: ", platform))
	}
}

// getPkgPath tries to get path of pkg. Path is platform dependent.
// First pkg is checked in GOPATH, then in GOROOT, then err
func getPkgPath(pkgname string) (string, error) {
	var ppath string
	if gopath != "" {
		for _, godir := range strings.Split(gopath, envSeparator) {
			ppath = strings.Join([]string{godir, "src", pkgname}, pathSeparator)
			_, err := os.Stat(ppath)
			if err != nil {
				continue
			}
			return ppath, nil
		}
	}
	ppath = strings.Join([]string{goroot, "src", pkgname}, pathSeparator)
	_, err := os.Stat(ppath)
	if err != nil {
		return "", err
	}
	return ppath, nil
}

// histogram computes and returns slice of histogram data points.
// If bucketSize is 0, default bucket size of defaultBucketSize is chosen.
func createHistogram(flens []FuncLen) []int {
	var hg []int
	if len(flens) == 0 {
		return nil
	}
	if bucketSize == 0 {
		bucketSize = defaultBucketSize
	}
	// find max func len
	var maxFlen int = -SENTINEL
	for _, flen := range flens {
		if flen.Size > maxFlen {
			maxFlen = flen.Size
		}
	}
	hglen := maxFlen / bucketSize
	hg = make([]int, hglen+1)
	for _, v := range flens {
		hg[v.Size/bucketSize]++
	}
	return hg
}

// Display histogram data points
func displayHistogram(hg []int) {
	var (
		start int = 0
		hglen int = len(hg)
	)
	tabw := new(tabwriter.Writer)
	tabw.Init(os.Stdout, 0, 4, 0, '\t', 0)
	fmt.Fprintln(tabw)
	for i := 0; i < hglen; i++ {
		bucketrange := fmt.Sprintf("[%d-%d)", start, start+bucketSize)
		streak := ""
		for j := 0; j < hg[i]; j++ {
			streak = fmt.Sprintf("%s#", streak)
		}
		start += bucketSize
		fmt.Fprintln(tabw, fmt.Sprintf("%s\t-\t%s", bucketrange, streak))
	}
	tabw.Flush()
}

func generateFuncLens(pkgpath string) ([]FuncLen, []ZeroLenFunc) {
	tabw := new(tabwriter.Writer)
	tabw.Init(os.Stdout, 0, 4, 0, '\t', 0)
	zeroLenFuncs := make([]ZeroLenFunc, 0)
	// create AST by parsing src
	fset := token.NewFileSet()
	pkgs, ferr := parser.ParseDir(fset, pkgpath, func(f os.FileInfo) bool {
		if inclTests {
			return true
		}
		return !strings.HasSuffix(f.Name(), "_test.go")
	}, parser.AllErrors)
	if ferr != nil {
		panic(ferr)
	}
	flens := make([]FuncLen, 0)
	for _, v := range pkgs {
		for fname, astf := range v.Files {
			for _, decl := range astf.Decls {
				ast.Inspect(decl, func(node ast.Node) bool {
					var (
						funcname string
						diff     int
					)

					switch x := node.(type) {
					case *ast.FuncDecl:
						funcname = x.Name.Name
						if x.Body == nil {
							fmt.Fprintln(tabw, fmt.Sprintf("%s\t%s", funcname, "=> implemented in package runtime"))
							return false
						}
						lb := x.Body.Lbrace
						rb := x.Body.Rbrace
						if !lb.IsValid() || !rb.IsValid() {
							return false
						}
						rln := fset.Position(rb).Line
						lln := fset.Position(lb).Line
						diff = rln - lln - 1
						if diff == -1 {
							diff = 1 // single line func
						}
						if diff == 0 {
							zeroLenFuncs = append(zeroLenFuncs, ZeroLenFunc{
								FuncLen{
									Name:     funcname,
									Size:     0,
									Filepath: fname,
									Lbrace:   lln,
									Rbrace:   rln,
								},
							})
							break
						}
						flens = append(flens, FuncLen{
							Name:     funcname,
							Size:     diff,
							Filepath: fname,
							Lbrace:   lln,
							Rbrace:   rln,
						})
					}
					return false
				})
			}
		}
	}
	tabw.Flush()
	return flens, zeroLenFuncs
}

func main() {
	var err error

	flag.Parse()

	if pkg == "" {
		fmt.Println("pkgname cannot be empty")
		return
	} else {
		pkgpath, err = getPkgPath(pkg)
		if err != nil {
			fmt.Printf("pkg %s cannot be found.\n", pkg)
			return
		}
		println("Full path of pkg: ", pkgpath)
	}

	flens, zeroLenFuncs := generateFuncLens(pkgpath)

	if len(zeroLenFuncs) != 0 {
		fmt.Printf("We have %d 0 len funcs!\n", len(zeroLenFuncs))
		zeroTabw := new(tabwriter.Writer)
		zeroTabw.Init(os.Stdout, 0, 4, 0, '\t', 0)
		fmt.Fprintln(zeroTabw)
		for _, zlf := range zeroLenFuncs {
			fmt.Fprintln(zeroTabw, fmt.Sprintf("%s\t%s", zlf.Name, zlf.Filepath))
		}
		zeroTabw.Flush()
	}

	hg := createHistogram(flens)
	if hg != nil {
		displayHistogram(hg)
	} else {
		fmt.Println("No functions found in pkg", pkg)
	}
}
