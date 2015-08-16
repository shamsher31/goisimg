# goisimg
Check if filepath is an Image

# How to install
<pre>
go get github.com/shamsher31/goisimg
</pre>

# How to use
<pre>
package main

import (
	"fmt"
	"github.com/shamsher31/goisimg"
)

func main() {
	fmt.Println(image.Is("./golang.jpg"))
  // true
}
</pre>

# Related
[goimgext](https://github.com/shamsher31/goimgext)
[go-binary](https://github.com/ferhatelmas/go-binary)
[go-archive](https://github.com/ferhatelmas/go-archive)

# Why
This package is inspired by [sindresorhus](https://www.npmjs.com/package/is-image) npm module to check if filepath is an Image.

# License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
