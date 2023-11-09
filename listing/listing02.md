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
2
1

Defer функций пушит их в стек. При возвращении из функции, в которой они добавлены, defer вызовы срабатывают в порядке Last In First Out.

Так как в функции test переменная x инициализируется в стеке вызывающей функции, она будет инкрементирована. В anothertest x инициализирована на стеке исполняемой функции, поэтому после ее завершения инкрементирована не будет.

```