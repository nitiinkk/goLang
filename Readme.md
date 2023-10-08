## The Go Programming Language
### Chapter-1
__________________
- Go is a compiled language
- to run the go program we do ```go run filename.go``` and if we want to compile result and use for later we do a build ```go build filename.go``` it creates a file ```filename and ./filename runs the executable file without further processing.```
- Go code is organize d into packages, which are similar to libraries or modules in other languages.
- slice in Go : Go uses half-open intervals s[m:n] , where last indx is omitted, eg ```os.Args[1:] or os.Args[1 :len(os.Args)]```
- Semicolon is not required in go.
- If variables is not explicitly initialized, it is implicitly initialized to the zero value for its type, which is 0 for numer ic types and the empty string "" for strings.
- := short variable declaration, assigns value and types.
- Go doesn't permit unused local variables and throws compilation error.
- Functions and other package levels entities may be declared in any order.
- Go Routine :: Is a concurrent function execution. 
    - A channel is a communication mechanism that allows one go routine to pass values of a sepecified type to another go routine.
- The function main runs in a go routine and go statement create additional go routines.
- 