# 知识点

## 知识点一
这是因为 Go 语言面向对象编程不像 PHP、Java 那样支持隐式的 this 指针，所有的东西都是显式声明的，在 GetXXX 方法中，由于不需要对类的成员变量进行修改，所以不需要传入指针，而 SetXXX 方法需要在函数内部修改成员变量的值，并且作用到该函数作用域以外，所以需要传入指针类型（结构体是值类型，不是引用类型，所以需要显式传入指针）。

## 知识点二
按照 Go 语言的约定，Integer 类型实现了 IntNumber 接口。然后我们可以这样将 Integer 类型对应的对象实例赋值给 IntNumber 接口：

```
var a Integer = 1
var b IntNumber = &a
```  
注意到上述赋值语句中，我们将对象实例 a 的指针引用赋值给了接口变量，为什么要这么做呢？因为 Go 语言会根据类似下面这样的非指针成员方法：

```
func (a Integer) Equal(i Integer) bool
```

自动生成一个新的与之对应的指针成员方法：

```
func (a *Integer) Equal(i Integer) bool { 
    return (*a).Equal(i) 
}
```

这样一来，类型 *Integer 就存在所有 IntNumber 接口中声明的方法，而 Integer 类型不包含指针方法 Increase 和 Decrease（关于这一点我们前面已经介绍过），所以严格来说，只有 *Integer 类型实现了 IntNumber 接口。如果我们贸然将 a 的值引用赋值给 b，编译时会报错：

```
cannot use a (type Integer) as type IntNumber in assignment:
    Integer does not implement IntNumber (Decrease method has pointer receiver)
```