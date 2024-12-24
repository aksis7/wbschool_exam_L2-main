Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```
<nil>
false
Ответ:
```
Интерфейс error хранит тип и значение. В Foo() тип – *os.PathError, а значение – nil.
Сравнение с nil проверяет оба поля, и так как тип не nil, результат – false.

```
