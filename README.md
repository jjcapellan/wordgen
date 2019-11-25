# WORDGEN  
Simple golang random words generator.  
  
## Usage  
Pass the length of the desired word to one of the four available functions.  
Example:
```go
package main

import (
    "fmt"
    "github.com/jjcapellan/wordgen"
)

func main(){  

    strA := wordgen.New(10)
    fmt.Println(strA)
    // Output -> "6=COp@0E50"

    strB := wordgen.OnlyLetters(15)
    fmt.Println(strB)
    // Output -> "XOjXsBcIWRswCuk"

    strC := wordgen.OnlyNumbers(12)
    fmt.Println(strC)
    // Output -> "084243421703"

    strD := wordgen.NotSymbols(20)
    fmt.Println(strD)
    // Output -> "iAP33k102y2m2v3bA5S4"

}
```