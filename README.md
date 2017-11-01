# go-tcnum
Turkish Identification Number Validation

## Usage

```
go get github.com/netinternet/gotcnum
```

## Exemple

```go
package main

import (
  "fmt"
  "github.com/netinternet/gotcnum"
)

func main() {
  if valid := gotcnum.tcnum("10148343778"); valid {
    fmt.Println("TC Number Is Valid")
  } else {
    fmt.Println("TC Number Is Not Valid")
  }
}


```
