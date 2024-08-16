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

Ответ:
```
При проверке if err != nil, условие выполнится, так как err содержит интерфейсное значение, у которого часть (*customError) не равна nil.
Программа выведет error.
Интерфейсное значение err не является nil, хотя оно содержит nil как значение типа *customError. Интерфейс включает информацию о типе, эта информация не является nil.
```
