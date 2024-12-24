Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```
error
Ответ:
```


Интерфейс error хранит тип (*customError) и значение (nil). При сравнении с nil проверяются оба поля, и так как тип не nil, условие err != nil возвращает true.

```
