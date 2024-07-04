# inject

`inject` is a small dependency injection library for Go.

[![GoDoc](https://godoc.org/github.com/infogulch/inject?status.svg)](https://godoc.org/github.com/infogulch/inject)

```shell
go get -u github.com/infogulch/inject
```

## Why?

### Dependency Injection?

Dependency Injection (DI) is a software design pattern that promotes loose
coupling between components by separating the creation and management of object
dependencies from the objects themselves. The promise of this pattern is that it
helps authors write more modular and testable code.

### DI in Go?

In the wider programming world, dependency injection typically revolves around
language level constructs like constructors. Go does not have constructors in
that sense, only conventions like the common `New()` function. Consequently,
usage of `inject` revolves around providing and requesting functions and types.

## How to use

1. Define your dependencies as structs, interfaces, functions that accept and
   return structs or interfaces, or slices of the previous.
2. Call `inject.New()`, providing all of your dependencies and constructor funcs
   as arguments, to get a `var scope *Scope` which implements `Injector`.
3. Request types with an `Injector`: (Note that Go does not support generic
   methods, so these must be regular functions.)
    a. Call the generic `inject.Get` like so: `t, err := inject.Get[T](scope)`
       to recursively evaluate the constructors to return a `t`, or an error if
       there was an issue.
    b. Call `inject.Inject[T](scope, myfunc)` to call the func with arguments
       sourced from the scope and return the func's return value. The func must
       return one value, and optionally an error.

## How it works

Build an injection scope by calling `inject.New` with values, functions, and
slices. Then call `inject.Get` and `inject.Inject` to get values and call
functions with the types available in the scope.

Provide an object `T` and it will be returned when requested. Provide any
function that returns a `U` (or a `(U, error)`), then that function will be
called to produce a `U` when requested. If a function requires inputs its
arguments will be extracted from the injection scope then called, recursively,
until the originally requested type is returned.

Each type can only have one provider, that is, you can't have two functions that
both return `T`, or provide a `T` directly and a `func() T`, etc. The exception
to this is slices, which are concatenated together if more than one are
provided. Like all types, only one of any specific primitive type is allowed. To
get around this, name your type like `type MyInt int`.

The exception to the "one type one provider" rule is slices. Adding multiple
slices of the same type (or slices of functions that return the same type)
results the scope containing one slice that is the concatenation of the provided
slices.

### Example

See the [example](./example/main.go) for a working example http server
application that demonstrates middleware, repository, and handler patterns.
