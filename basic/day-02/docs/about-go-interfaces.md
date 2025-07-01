# Go Interfaces: Concepts and Usage

## What is an Interface?
- An interface in Go is a type that specifies a set of method signatures.
- Any type that implements those methods satisfies the interface—no explicit declaration is needed.

## Declaring and Implementing Interfaces
```go
type Transaction interface {
    Execute() (bool, error)
}

type Transfer struct{}

func (t Transfer) Execute() (bool, error) {
    // Implementation
    return true, nil
}
```
- Here, `Transfer` implements the `Transaction` interface by defining the `Execute` method.

## Interface Satisfaction
- Go uses implicit interface satisfaction: a type implements an interface simply by having the required methods.
- No `implements` or `extends` keyword is needed.

## Interfaces and Files
- The interface and its implementing types do not need to be in the same file—just in the same package (or imported appropriately).

## Multiple Interfaces with Overlapping Methods
- If two interfaces have methods with the same signature, a type only needs to implement that method once to satisfy both interfaces.
```go
type A interface { Foo() error }
type B interface { Foo() error; Bar() }

type MyType struct{}
func (m MyType) Foo() error { return nil }
func (m MyType) Bar()       {}
// MyType implements both A and B
```

## Assigning to Interface Variables
- If a type implements all methods of an interface, you can assign a value of that type to a variable of the interface type:
```go
var t Transaction = Transfer{}
```
- If a type implements multiple interfaces, it can be assigned to variables of any of those interface types.

## Compile-Time Checking
- If you try to assign a type to an interface variable, but the type does not implement all required methods, Go will give a compile-time error.

## Interface Example with Multiple Interfaces
```go
type A interface { Foo() error }
type B interface { Foo() error; Bar() }

type MyType struct{}
func (m MyType) Foo() error { return nil }
// func (m MyType) Bar() {} // Uncomment to satisfy B

var a A = MyType{} // OK
// var b B = MyType{} // Compile error: missing Bar()
```

## Value vs Pointer Receiver in Interface Implementation

In Go, you can implement interface methods using either a value receiver or a pointer receiver:

### Value Receiver
```go
func (t Transfer) Execute() (bool, error) { ... }
```
- The method can be called on both `Transfer` and `*Transfer`.
- The method receives a copy of the struct, so changes to fields inside the method do not affect the original struct.
- Use value receivers when the method does not need to modify the struct or when the struct is small and copying is cheap.

### Pointer Receiver
```go
func (t *Transfer) Execute() (bool, error) { ... }
```
- The method can only be called on `*Transfer` (a pointer to Transfer).
- The method can modify the struct’s fields, and those changes will be visible to the caller.
- Use pointer receivers when the method needs to modify the struct, or when the struct is large and you want to avoid copying.

### Summary Table of Receivers
| Receiver Type   | Can Modify Fields | Usable On      | Common Use Case                        |
|----------------|------------------|---------------|----------------------------------------|
| Value          | No               | value & ptr   | Read-only methods, small structs       |
| Pointer        | Yes              | pointer only  | Methods that modify struct, large data |

**Tip:**
If any method of an interface is implemented with a pointer receiver, only a pointer to the struct implements the interface. If all methods use value receivers, both the value and pointer types implement the interface.

Use value receivers for immutable/read-only operations, and pointer receivers when you need to mutate the struct or avoid copying large structs.

## Summary Table of Interface
| Feature                        | Go Interface Behavior                         |
|--------------------------------|-----------------------------------------------|
| Declaration                    | `type MyInterface interface { ... }`          |
| Implementation                 | Implicit (by method set)                      |
| Multiple interfaces, same method| One method satisfies all matching signatures  |
| Assignment                     | Only if all methods are implemented           |
| Compile-time check             | Yes, for interface satisfaction               |
| File/package location          | Can be in different files, same package       |

---

**References:**
- [Go Blog: Interfaces](https://go.dev/blog/interfaces)
- [Go Tour: Interfaces](https://tour.golang.org/methods/9)
