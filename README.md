##Get info on length of functions in a Go package.

Given package is searched in  directories provided by envs in following order:  
1. GOPATH  
2. GOROOT  
AST is generated only for Go files present in package path, ie, `flen -pkg crypto` shall only parse `crypto.go`. For parsing `sha1`, full package path needs to be provided, ie `flen -pkg crypto/sha1`.

###Install
`go get github.com/lafolle/flen`

###Usage  
`flen -pkg <pkg name>`  
eg `flen -pkg encoding/json`

###TODO:
1. Test histogram func  
2. Better ways to show data  
3. Also include anon funcs  
4. Show only stats of funclengths  
5. Optionally show histogram(?)  
6. Some way to show, which funcs are of really great length.
