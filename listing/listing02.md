Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
В test() используется именнованный возвращаемый параметр, который можно изменить внутри deferred функции
Когда выполняется test():
1. x=1
2. x++ внутри defer меняет значение x на 2
3. возвращаем x=2

anotherTest() показывает, что мы не можем изменить неименнованные возвращаемые параметры внутри deferred функции
anotherTest():
1. var x int: x=0
2. x=1
3. возвращаем x=1

defer выполняется после return или panic. Поэтому может быть опасно вставлять defer в цикл, иногда лучше вызывать функции напрямую.
Deferred функции заносятся в стек. Когда функция завершает выполнение, deferred функции выполняются в LIFO-стиле.

```
