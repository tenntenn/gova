gova
===============

It provides validation functions using struct tags.


How to use?
---------------

```
import "github.com/tenntenn/gova"

type MyType struct {
	FirstName string `length:"20" pattern:"[A-Z][a-z]*"`
	LastName  string `length:"20" pattern:"[A-Z][a-z]*"`
	Email     string `email:"-"`
}

v1 := &MyType{"Hoge", "Fuga", "hogefuga@gmail.com"}
fmt.Println(gova.Validate(v1)) // return nil

v2 := &MyType{"Hoge", "Fuga", "hogefuga"}
fmt.Println(gova.Validate(v2)) // return error
```
