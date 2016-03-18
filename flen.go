package flen

import (
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
	sentinel = 1000000
)

var (
	goroot        string
	pkgpath       string
	gopath        string = os.Getenv("GOPATH")
	pathSeparator string
	envSeparator  string
	opts          Options
)

// init sets platform dependent vars.
func init() {
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

type Options struct {
	IncludeTests bool
	BucketSize   int
	StreakChar   string // TODO: can be a rune?
}

type funcLen struct {
	Name           string
	Size           int
	Filepath       string
	Lbrace, Rbrace int
}

type zeroLenFunc struct {
	funcLen
}

type FuncLens []funcLen

// createHistogram computes and returns slice of histogram data points.
func (flens *FuncLens) computeHistogram() []int {
	var hg []int
	if len(*flens) == 0 {
		return nil
	}
	// find max func len
	var maxFlen int = -sentinel
	for _, flen := range *flens {
		if flen.Size > maxFlen {
			maxFlen = flen.Size
		}
	}
	hglen := maxFlen / opts.BucketSize
	hg = make([]int, hglen+1)
	for _, v := range *flens {
		hg[v.Size/opts.BucketSize]++
	}
	return hg
}

// Display histogram data points
func (flens *FuncLens) DisplayHistogram() {
	var (
		start int   = 0
		hg    []int = flens.computeHistogram()
		hglen int   = len(hg)
	)
	tabw := new(tabwriter.Writer)
	defer tabw.Flush()
	tabw.Init(os.Stdout, 0, 4, 0, '\t', 0)
	fmt.Fprintln(tabw)
	for i := 0; i < hglen; i++ {
		bucketrange := fmt.Sprintf("[%d-%d)", start, start+opts.BucketSize)
		streak := ""
		for j := 0; j < hg[i]; j++ {
			streak = fmt.Sprintf("%s%s", streak, opts.StreakChar)
		}
		fmt.Fprintln(tabw, fmt.Sprintf("%s\t-\t%s", bucketrange, streak))
		start += opts.BucketSize
	}
}

// Stats generate statstics on length of functions in a package.
func (flens *FuncLens) Stats() {
}

// GenerateFuncLens generates FuncLens for the given package. If options.InclTests is true,
// functions in tests are also evaluated.
func GenerateFuncLens(pkg string, options Options) (FuncLens, []zeroLenFunc, error) {
	opts = options
	pkgpath, err := getPkgPath(pkg)
	if err != nil {
		return nil, nil, err
	}
	println("Full path of pkg: ", pkgpath)

	tabw := new(tabwriter.Writer)
	tabw.Init(os.Stdout, 0, 4, 0, '\t', 0)
	zeroLenFuncs := make([]zeroLenFunc, 0)
	fset := token.NewFileSet()
	pkgs, ferr := parser.ParseDir(fset, pkgpath, func(f os.FileInfo) bool {
		if opts.IncludeTests {
			return true
		}
		return !strings.HasSuffix(f.Name(), "_test.go")
	}, parser.AllErrors)
	if ferr != nil {
		panic(ferr)
	}
	flens := make([]funcLen, 0)
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
							zeroLenFuncs = append(zeroLenFuncs, zeroLenFunc{
								funcLen{
									Name:     funcname,
									Size:     0,
									Filepath: fname,
									Lbrace:   lln,
									Rbrace:   rln,
								},
							})
							break
						}
						flens = append(flens, funcLen{
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
	return flens, zeroLenFuncs, nil
}

// getPkgPath tries to get path of pkg. Path is platform dependent.
// First pkg is checked in GOPATH, then in GOROOT, then err.
func getPkgPath(pkgname string) (string, *os.PathError) {
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
		return "", err.(*os.PathError)
	}
	return ppath, nil
}
