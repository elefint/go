package main

import "fmt"

func main() {
    a := 21
    b := 42

    fmt.Println(a == b) // false
    fmt.Println(a == a) // true
    fmt.Println(a != b) // true
    fmt.Println(a > b)  // false
    fmt.Println(a < b)  // true
    fmt.Println(a <= b) // true
    fmt.Println(a <= a) // true
    fmt.Println(a >= a) // true
}

/* Operator	Signification
==	equal
!=	not equal
>	greater
>=	greater or equal
<	less
<=	less or equal */

// The comparison operator used to test equality is == and not =. The last is the symbol used for assignment and not for comparison. Replacing == with = is a common mistake, so watch it.