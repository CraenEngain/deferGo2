Что выведет программа? Объяснить вывод программы.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
Создается массив из 5 элементов
После чего создается срез со ссылкой на 3 элемента из массива а с элементами с 1 до 3 индекса (77, 78, 79) 
После этого выводится содержимое среза
```
