## Go Introduction

### Compiled or Intepreted

- Compiled
- Statically typed
- The God tool can run a file without precompiling
- Executables are OS specific 
- No external VM is needed (like a java JVM)


### Is it object orientated?

Kind of

- You can define custom interfaces
- Custom types can implement one or more interfaces
- Custom types can have member methods
- Custom structs (data structures) can have member fields
- Feels like classes in Java
- No classes so no type inheritence

But

- No type inheritence as no classes
- No method or operator overloading
- Structured exception handling (try catch finally) - err objects returned by functions
- No implicit numeric conversions

Designed to be easier to read and concise

### Essential Syntax Rules

- Go is case sensitive
- Variable and package names are lowercase and mixed case
- Names of public fields have initial uppercase characters
- Initial uppercase character means the symbol is exported
- Semicolons usually not needed (so it is sensitive to whitespace as the end marker)

#### Built-In Functions

- len(string) - returns length of string
- panic(error) - stops execution; displays err msg
- recover() - manages behaviour of a panicking goroutine
- [Full list](https://pkg.go.dev/builtin)


 #### Memory Allocation

- Go runtime is statically linked into application
- Memory is allocated and deallocated automatically
- new() - allocates but doesn't initialize - zeroed storage but returns a memory addres - use for value types e.g. struct where you want a pointer to an initialised structure
- make() allocates and initializes memory - allocates non-zeroes and returns a memory address - use for reference types e.g. to create

 #### Memory Deallocation

 - Memory˙ is deallocated by garbage collector (GC)
 - Objects out of scopr or set to nil are eligible
 - GC was rebuilt in Go 1.5 for very low latency

 




