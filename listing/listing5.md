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

Так как до вызова функции ```test()``` переменная ```err``` не имела конкретного типа и значения, 
при сравнении с ```nil``` мы бы получили ```err == nil (true)```.
Из функции нам вернулся конкретный тип ```*customError```, и теперь, хоть значение ```err``` и ```nil```,
но тип уже ```*customError```, таким образом сравнение ```err != nil```, даёт ```true```.