##Get info on length of functions in a Go package.

Given package is searched in  directories provided by envs in following order:  
1. GOPATH  
2. GOROOT  
AST is generated only for Go files present in package path, ie, `flen crypto` shall only parse `crypto.go`. For parsing `sha1`, full package path needs to be provided, ie `flen crypto/sha1`.

###Install
`go get github.com/lafolle/flen`

###Usage  
###Simple Usage
```
Usage: flen <pkg> [options]
  -bs int
        bucket size (natural number) (default 5)
  -l int
        min length (inclusive)
  -t    include tests files
  -u int
        max length (exclusive) (default 1000000)	
```
###Example usage  
```
$ flen strings
Full path of pkg:  /usr/local/go/src/strings
Externally implemented funcs
+-------+-----------+-------------------------------------------+---------+------+
| INDEX |   NAME    |                 FILEPATH                  | LINE NO | SIZE |
+-------+-----------+-------------------------------------------+---------+------+
|     0 | IndexByte | /usr/local/go/src/strings/strings_decl.go |       0 |    0 |
+-------+-----------+-------------------------------------------+---------+------+

[1-6)   -       ∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
[6-11)  -       ∎∎∎∎∎∎∎∎∎∎
[11-16) -       ∎∎∎∎∎∎∎∎∎∎∎∎
[16-21) -       ∎∎∎∎∎
[21-26) -       ∎∎∎
[26-31) -       ∎∎∎∎
[31-36) -       ∎∎∎∎
[36-41) -       ∎∎∎
[41-46) -       ∎∎
[46-51) -       ∎
[51-56) -
[56-61) -       ∎
[61-66) -
$
```  
