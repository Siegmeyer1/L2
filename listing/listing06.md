Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
Несмотря на то, что слайс содержит указатель на массив, он всё равно является структурой, содержащей также len и cap -- он не является исключительно указателем.
И мы передаём в функцию только копию слайса.

Мы передаём в modifySlice этот дескриптор слайса по значению. Он содержит указатель на массив, поэтому и оригинальный слайс, и его копия в функции ссылаются на один массив.
Содержимое массива может измениться в функции, но len, cap и сам указатель -- нет, так как функция имеет дело только с копией.

Следовательно, i[0] = 3 присвоило первому элементу оригинального слайса значение 3, а дальше копии i присвоили другой слайс, который никак не может повлиять на оригинальный.

```
