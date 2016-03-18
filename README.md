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
$ flen -pkg strings
Full path of pkg:  /usr/local/go/src/strings
IndexByte	=> implemented in package runtime

[0-5)	-	|||||||||||||||||||||||||||||||||||||
[5-10)	-	||||||||||
[10-15)	-	|||||||||||||
[15-20)	-	||||
[20-25)	-	||||
[25-30)	-	||||
[30-35)	-	|||||
[35-40)	-	||
[40-45)	-	|||
[45-50)	-	
[50-55)	-	|
[55-60)	-	
[60-65)	-	|
$
```  
####Use `-bs` to specify bucket size of histogram  
```
$ flen -bs 2 -pkg strings
Full path of pkg:  /usr/local/go/src/strings
IndexByte	=> implemented in package runtime

[0-2)	-	|||||||||||||||||||||||||||
[2-4)	-	|||
[4-6)	-	|||||||||
[6-8)	-	||||||
[8-10)	-	||
[10-12)	-	|||
[12-14)	-	||||||
[14-16)	-	|||||
[16-18)	-	||
[18-20)	-	|
[20-22)	-	||
[22-24)	-	||
[24-26)	-	|
[26-28)	-	||
[28-30)	-	|
[30-32)	-	||
[32-34)	-	|||
[34-36)	-	
[36-38)	-	|
[38-40)	-	|
[40-42)	-	|
[42-44)	-	|
[44-46)	-	|
[46-48)	-	
[48-50)	-	
[50-52)	-	|
[52-54)	-	
[54-56)	-	
[56-58)	-	
[58-60)	-	
[60-62)	-	|
$
```
####Use `-t` to include test files of package.
```
$ flen -t -pkg strings
Full path of pkg:  /usr/local/go/src/strings
IndexByte	=> implemented in package runtime

[0-5)		-	|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
[5-10)		-	||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
[10-15)		-	|||||||||||||||||||||||
[15-20)		-	||||||||||
[20-25)		-	||||||
[25-30)		-	|||||||
[30-35)		-	|||||
[35-40)		-	|||
[40-45)		-	|||||||
[45-50)		-	|
[50-55)		-	|
[55-60)		-	|
[60-65)		-	|
[65-70)		-	|
[70-75)		-	
[75-80)		-	
[80-85)		-	
[85-90)		-	
[90-95)		-	
[95-100)	-	
[100-105)	-	
[105-110)	-	
[110-115)	-	
[115-120)	-	
[120-125)	-	
[125-130)	-	
[130-135)	-	
[135-140)	-	
[140-145)	-	
[145-150)	-	
[150-155)	-	
[155-160)	-	
[160-165)	-	
[165-170)	-	
[170-175)	-	
[175-180)	-	
[180-185)	-	
[185-190)	-	
[190-195)	-	
[195-200)	-	
[200-205)	-	
[205-210)	-	
[210-215)	-	
[215-220)	-	
[220-225)	-	
[225-230)	-	
[230-235)	-	
[235-240)	-	
[240-245)	-	
[245-250)	-	
[250-255)	-	
[255-260)	-	
[260-265)	-	
[265-270)	-	|
```  
#### Use `-c` to specify streak char of histogram
```
$ flen -c \# -pkg strings
Full path of pkg:  /usr/local/go/src/strings
IndexByte	=> implemented in package runtime

[0-5)	-	#####################################
[5-10)	-	##########
[10-15)	-	#############
[15-20)	-	####
[20-25)	-	####
[25-30)	-	####
[30-35)	-	#####
[35-40)	-	##
[40-45)	-	###
[45-50)	-	
[50-55)	-	#
[55-60)	-	
[60-65)	-	#
$
```
