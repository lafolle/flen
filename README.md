##Get info on length of functions in a Go package.

Given package is searched in  directories provided by envs in following order:  
1. GOPATH  
2. GOROOT  
AST is generated only for Go files present in package path, ie, `flen -pkg crypto` shall only parse `crypto.go`. For parsing `sha1`, full package path needs to be provided, ie `flen -pkg crypto/sha1`.

###Install
`go get github.com/lafolle/flen`

###Usage  
####Simple usage 
```
$ flen strings
$

If not pkg is provided, flen will consider cwd to be a pkg, and shall give results for it.
```  
