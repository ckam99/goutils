
Example

```go
package main

import (
	"fmt"
	//  "time"
  "main/try"
)



func main() {
  
  try.Catch(func(e error){
    fmt.Println(e)
  }).Do(func() {
      fmt.Println("I tried: catch/try")
      panic("error occured!")
  })

  try.Do(func(){
    fmt.Println("I tried: try/catch")
    panic("error occured!")
  }).Catch(func(e error) {
      fmt.Println(e)
  })

  
}

```
