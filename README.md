# images

##### Description

图片处理库

##### Installation

```
go get gitee.com/eshax/images
```

##### Usage

- go.mod
```
module image.test

go 1.17

require gitee.com/eshax/images v0.5.0
```

- main.go
```
package main

import (
    "os"
    "gitee.com/eshax/images"
)

func main() {

    // 单图瓦片
    images.Single("dist/100x/0.jpg", "dist/100x/0/dzi", 256)

    // 连续视野瓦片
    images.Matrix("dist/lower", "dist/lower/dzi", 16384, 256)
}
```