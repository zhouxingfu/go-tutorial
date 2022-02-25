# Go开发环境搭建

1. 官网下载最新版go（Ubuntu为例），解压到/opt目录，一般把App放置到系统主目录下，而不是user目录
2. 设置GoRoot和GoPath，在/etc/profile中添加上述两项，并添加到系统路径Path。GoRoot是/opt/go, GoPath是/home/bit/go-repo，类似于这样的目录。

```
export GOROOT=/opt/go/ 
export GOPATH=/home/bit/go-repo/
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

3. 在go-repo目录下新建一个目录 src，在src新建 目录 hello，在hello下新建一个main.go

```C++
package main

import "fmt"

func main(){
	fmt.Println("HelloWorld!");
}
```

4. go build会在当前目录下生成hello ，go install会将编译好的二进制剪切到$GOPATH/bin目录下。go run会显式结果，之后会立即删除生成的二进制。


5. gomod管理不同项目间第三方包的方式是： 把所有第三方库（包括不同版本）都放置到pkg目录下，而每个项目需要依赖的第三方包都在每个工程的go.mod文件中记录。  

在go.mod中格式如下
```
module hello

go 1.17

require github.com/gin-gonic/gin v1.5.0
```

go mod download在下载第三方包的同时也会生成一个go.sum文件，这个文件是dependency tree。


6. 关于包的升级，以及由于防火墙的原因导致某些包无法下载，此时我们用其它包来replace的方法请参照 https://studygolang.com/articles/28712

7. import其它文件夹

假设目录结构是 src -- employee
                        -- main.go
                        -- basic
                            -- basic.go
                            -- gross
                                -- gross.go

```go
//main.go
package main
import(
    b "employee/basic" 这里 b类似于namespace，相当于 import numpy as np
    "employ/basic/gross"
    "fmt"
)


//gross.go
package gross
import(
    b "employee/basic"
)
```

从上面的例子中，我们可以得知，__import的路径不是针对当前文件的相对路径，而是以go module root directory开始的相对路径__。


