Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и порядок их вызовов.

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

output: 
2
1

Так как первая функция возвращает именованный результат - функция завершится => дефер изменит Х, и именно его значение 
вернётся из функции. тогда как во втором случае функция завершится => дефер изменит локальный Х, но вернётся уже его значение,
которое на момент завершения функции было равно 1.