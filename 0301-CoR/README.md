# 职责链

Chain of Responsibility

责任链模式是一种行为设计模式， 允许你将请求沿着处理者链进行发送。 收到请求后， 每个处理者均可对请求进行处理（结束传递）， 或将其传递给链上的下个处理者。

## 使用场景

1. 需要使用不同方式处理不同种类请求， 而且请求类型和顺序预先未知时， 可以使用责任链模式。
2. 当必须按顺序执行多个处理者时， 可以使用该模式
3. 所需处理者及其顺序必须在运行时进行改变， 可以使用责任链模式

## 例子

1. 中间件机制（洋葱模型）
2. js冒泡

## 和装饰器的区别
