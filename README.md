##Get info on length of functions in a Go package.

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
