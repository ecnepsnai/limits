# limits

[![Go Report Card](https://goreportcard.com/badge/github.com/ecnepsnai/limits?style=flat-square)](https://goreportcard.com/report/github.com/ecnepsnai/limits)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/ecnepsnai/limits)
[![Releases](https://img.shields.io/github/release/ecnepsnai/limits/all.svg?style=flat-square)](https://github.com/ecnepsnai/limits/releases)
[![LICENSE](https://img.shields.io/github/license/ecnepsnai/limits.svg?style=flat-square)](https://github.com/ecnepsnai/limits/blob/master/LICENSE)

Package limits provides a way to define minimum and maximum length requirements on fields of a struct, then check a instance against those rules.

# Usage

Add the `min:"<int>"` or `max:"<int>"`  tag to any struct field you wish to enforce.
Limits only works for strings and slices (of any types, enforcing the count).

```go
type User struct{
    Username string `min:"1" max:"32"`
    Email string `min:"1" max:"128"`
    Password string `min:"12" max:"256"`
}
```

Then, check an instance of your struct to see if any values violate the limits:

```go
user := User{
    Username: "" // Too short!
    Email: "..." // Too long (pretend)!
    Password: "What are you looking at???" // Just right!
}

err := limits.Check(user) // Also works with references to objects!
if err != nil {
    fmt.Println(err.Error()) // --> Value for 'Username' requires at least 1 character
}
```