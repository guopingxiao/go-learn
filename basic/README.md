## 1. 变量定义

使用var关键字

```
  var a b c bool
  var s1, s2 string = "hello", "world"
```

可以放在函数内，或直接放在包里，作为包的一个变量

使用`var()`集合来定义变量

## 内置变量

- bool, string

go没有long类型，需要几位，声明几位就可以，int的长度是根据操作系统来的；

- (u)int , (u)int8, (u)int8, (u)int16, (u)int32, (u)int864, uintptr 

- byte（8位）, rune(类似于char，但是在面对多国语言时，一字节的char，有坑多， rune是32位)

- float32 , float64(没哟double), complex64, complex128 (复数，实部和虚部分别为64位)


 