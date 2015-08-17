# goisimg
Check if filepath is an Image

### How to install
```go
go get github.com/shamsher31/goisimg
```

### How to use
```go
package main

import (
	"fmt"
	"github.com/shamsher31/goisimg"
)

func main() {
    fmt.Println(image.Is("./golang.jpg"))
    // true
}
```
###Aliasing Imports
If you already have package name ```image``` you can use following.

```go
package main

import (
	"fmt"
	pckImage "github.com/shamsher31/goisimg"
)

func main() {
    fmt.Println(pckImage.Is("./golang.jpg"))
    // true
}
```
### Related
[goimgext](https://github.com/shamsher31/goimgext)<br>
[go-binary](https://github.com/ferhatelmas/go-binary)<br>
[go-archive](https://github.com/ferhatelmas/go-archive)<br>

### Why
This package is inspired by [is-image](https://www.npmjs.com/package/is-image) npm module.

# License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
