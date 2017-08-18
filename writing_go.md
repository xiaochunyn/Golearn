编写地道的Go代码

最地道的Go代码就是Go的标准库的代码，有空的时候可以多看看Google的工程师是如何实现的。

## 1. 注释
可以通过/* ... */或者//增加注释， //之后应该有个空格
如果想在每个文件的头部加上注释，需要在版权注释和Package前面加一个空行，否则版权注释会作为package的注释
```
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package net provides a portable interface for network I/O, including
TCP/IP, UDP, domain name resolution, and Unix domain sockets.
......
*/

package net
......

```
注：注释应该用一个完整的句子，注释的第一个单词应该是要注释的指示符，以便在godoc中容易查找;
   注释应该以 `.` 结尾；

## 2. 声明slice
使用下面这种方式声明slice:
``` 
var s []string
```
而不是下面这种格式
``` 
t := []string{}
```
注：前者声明了一个`nil`slice, 而后者声明了一个长度为0的非`nil`slice

## 3. 字符串的大小写
错误字符串不应该大写，应写成：
``` 
fmt.Errorf("failed to write data.")
```
而不是写成：
``` 
fmt.Errorf("Failed to write data")
```
因为，这些字符串可能和其他字符串相连接，组合后的字符串如果中间有大写字母开头的单词很突兀，除非这些首字母大写单词是固定使用的单词。

注：缩写词必须保持一致，比如都大写URL或者小写url；
常亮一般声明为MaxLength，而不是以下划线分割MAX_LENGTH或者MAXLENGTH；


## 4.处理error而不是panic或者忽略
为了代码的强健性，不要使用_忽略错误，而是要处理每一个错误，尽管代码写起来有些繁琐也不要忽略错误；

尽量不要使用panic;

## 5. 一些名称

包名应该使用单数形式，比如util,model,而不是utils,models;

Receiver的名称应该缩写，一般使用一个或两个字符作为Receiver的名称，如：
``` 
func (f foo) method {

    ...

}
```

有些单词可能有多种写法，在项目中应保持一致，比如Golang采用的写法：
``` 
// marshaling
// unmarshaling
// canceling
// cancelation

```
而不是：
``` 
// marshalling
// unmarshalling
// cancelling
// cancellation

```

## 6.空字符串检查
正确方式：
``` 
if s == "" {
    ...
}
```
而不是：
```go
if len(s) == 0 {
    ...
}
```
更不是：
```go
if s == nil || s == ""{
    ...
}
```
## 7.非空slice检查
正确方式：
```go
if len(s) > 0 {
    ...
}
```
而不是：
```go

if s != nil && len(s) > 0 {
    ...
}
```

## 8. 直接使用bool值
对于bool类型的变量var b bool, 直接使用它作为判断，而不是使用它和true/false进行比较
正确方式：
```go
if b {
    ...
}
if !b {
    ...
}
```

而不是：
```go
if b == true {
    ...
}
if b == false {
    ...
}
```

## 9. byte/slice/string相等性比较
```go
var s1 []byte
var s2 []byte

    ...
bytes.Equal(s1, s2) == 0
bytes.Equal(s1, s2) != 0

```
而不是：
```go
var s1 []byte
var s2 []byte

    ...
bytes.Compare(s1, s2) == 0    
bytes.Compare(s1, s2) != 0
```

## 10. 检查是否包含子字符串
应使用strings.ContainesRune, strings.ContainesAny, strings.Contains

## 11. 复制slice
使用内建函数copy，而不是遍历slice逐个复制
正确方式
```go
var b1, b2 []byte
copy(b2, b1)
```

## 12. 尽量缩短if
正确方式：
``` 
  var a, b int
  ...
  return a > b
```
而不是：
```go
    var a, b int
    ...
    if a > b {
        return true
    } else {
        return false
    }
```

## 13.简化range
正确方式：
```go
    for range m {
        ...
    }
```
而不是：
```go
       var m map[string]int
       for _ = range m { 
    }
    for _, _ = range m {
    }

```

## 14.使用strings.TrimPrefix / strings.TrimSuffix
正确方式：
```go
    var s1 = "a string value"
    var s2 = "a "
    var s3 = strings.TrimPrefix(s1, s2)

```

而不是：
```go
       var s1 = "a string value"
       var s2 = "a "
       var s3 string
    if strings.HasPrefix(s1, s2) { 
        s3 = s1[len(s2):]
    }
```

## 15.append slice 
正确方式：
```go
    var a, b []byte
    a = append(b, a...)

```

而不是：
```go
    var a, b []byte
    for _,v range a {
        append(b, v)
    }
```

## 参考文档

http://colobu.com/2017/02/07/write-idiomatic-golang-codes/

[effective go](https://go-zh.org/doc/effective_go.html)