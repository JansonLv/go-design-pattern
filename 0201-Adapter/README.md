# 适配器

封装器模式、Wrapper、Adapter

适配器是一种结构型设计模式， 它能使不兼容的对象能够相互合作。

适配器可担任两个对象间的封装器， 它会接收对于一个对象的调用， 并将其转换为另一个对象可识别的格式和接口。

## 适配器结构模式

对象适配器：适配器实现了其中一个对象的接口， 并对另一个对象进行封装。

类适配器：使用了继承（多重继承）机制，同时继承两个对象的接口。

在go中，使用了对象适配器，两个对象同时实现对象的接口

## 概念示例

这里有一段客户端代码， 用于接收一个对象 （Lightning 接口） 的部分功能， 不过我们还有另一个名为 adaptee 的对象 （Windows 笔记本）， 可通过不同的接口 （USB 接口） 实现相同的功能

这就是适配器模式发挥作用的场景。 我们可以创建这样一个名为 adapter 的结构体：

1. 遵循符合客户端期望的相同接口 （Lightning 接口）。

2. 可以适合被适配对象的方式对来自客户端的请求进行 “翻译”。 适配器能够接受来自 Lightning 连接器的信息， 并将其转换成 USB 格式的信号， 同时将信号传递给 Windows 笔记本的 USB 接口。

## 弊端

使用适配器模式本身就可能是个问题，因为一个好的系统内部不应该做任何侨界，模型应该保持一致性。只有在如下情况才考虑使用适配器模式：

1. 新老系统接替，改造成本非常高。
2. 三方包适配。
3. 新旧 API 兼容。
4. 统一多个类的接口。一般可以结合工厂方法使用。
5. 数据库ORM
