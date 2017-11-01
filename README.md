# go-tcnum
Turkish Identification Number Validation

## Usage

```
go get github.com/netinternet/go-tcnum
```

## Exemple

```go
package main

import (
    "fmt"
    "github.com/netinternet/go-tcnum"
)

func main() {
    if valid := VerifyTcnum("10148343778"); valid {
      fmt.Println("TC Number Is Valid")
    } else {
      fmt.Println("TC Number Is Not Valid")
    }
}

```
