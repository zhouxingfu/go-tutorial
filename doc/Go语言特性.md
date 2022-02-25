# Go 语言特性


# 1. for

go的loop只有for，而且相当于C语言中的while和for的结合体。Go的for后面可以接 0个条件，1个条件，或者3个条件。

```
    for{} //一直执行
``` 

# 2. array 和 slice

~~固定大小的叫array，非固定大小的叫slice。~~ 这么说也不对，因为var a = make([]string, 3) 此时a的大小就是确定的
  
  关于array和slice的区别 https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html


  对上面这篇文章做一个简要总结

  array是只能传value的固定大小的元素集合，声明方式如下

  ```Go
  [N]Type
  [N]Type(value1, value2, value3, ... , valueN)
  [...]Type(value1, value2, ... , valueN)
  ```

  而slice是可以resize的，最初我们可以不设定大小，即使给定了一个大小，我们也可以改变它。 在这一点上类似于vector。不过还有明显的一点不同是，slice传的是引用。  

而且在go中很严格，slice和array形参不能转换。  

比如下面的代码

```Go
package hello


func printArray(a [5]int){ //形参是array
    fmt.Printf("address of parameter : %p\n", &a)
}
func printSlice(a []int) { //形参是slice
    fmt.Printf("address of parameter : %p\n", &a)
}

func main(){
    a := [5]int{1,2,3,4,5} //array

    printArray(a) //OK
    printArray(a[0:2]) //error, 实参是slice，形参是array

    printSlice(a) //error，实参是array，形参是slice
    printSlice(a[0:2]) //OK
}

```

slice的定义方式有如下几种
```Go
make([]Type, length, capacity)
make([]Type, length)
[]Type{}
[]Type{value1, value2, ..., valueN} //注意这里和array声明的区别

```

__多维array和多维slice__ 


```Go
var a = [3][2]int{
    [2]int{1,2},
    [2]int{2, 3},
}

	var a = [3][2]int{
		[2]int{},
		[2]int{},
	}

   var a10 = [...][2]int{
        {6,7},
        {8,9},
        {0,1},
    }
```

多维数组遍历方式

```Go
package main

import "fmt"

func main()  {
    var a10 = [...][2]int{
        {6,7},
        {8,9},
        {0,1},
    }
    // 方式一
    for i := 0; i < len(a10); i++ {
        fmt.Print(a10[i])  // [6 7][8 9][0 1]
    }
    // 方式二
    for _, i := range a10{
        fmt.Print(i)  // [6 7][8 9][0 1]
    }
}
```


多维slice的创建必须确定第一维度的大小，否则第二维度甚至第N维度此时的size没有意义。  


```Go
    a := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            a[i][j] = i + j
        }
    }
    fmt.Println("2d: ", a)
```

# 3. map 

map跟其它语言的map dict是一样的数据结构。  

```Go
    a := make(map[string]int) 
    b := make(map[string]int){"a":1, "b":2}
```



range类似于python中的作用，可以iterator array slice map等，返回两个值，index value或者 key value。

# 4. functions

go中函数返回值可以是多个数据的组合，有点类似于tuple。

还有可能返回的是一个函数体，这里叫闭包closure，他们说闭包是为了扩展生成周期，但实际上有问题的