# Method Overriding in Go (with Composition)

## Does Go Support Method Overriding?
- Go does not have traditional class-based inheritance, so it does not support method overriding in the OOP sense.
- However, you can achieve similar behavior using composition and method redefinition.

## How It Works
- When you embed a struct (either by value or by field), you can define a method with the same name on the outer struct.
- The outer struct's method will "override" (shadow) the embedded struct's method when called on the outer struct.
- You can still call the embedded struct's method explicitly if needed.

## Example
```go
type BaseHandler struct {}

func (b BaseHandler) Log(message string) {
    fmt.Println("Base Log:", message)
}

type SecureHandler struct {
    BaseHandler BaseHandler
}

// This method shadows BaseHandler.Log when called on SecureHandler
func (s SecureHandler) Log(message string) {
    s.BaseHandler.Log(message) // Call the embedded method
    fmt.Println("Secure Log:", message)
}

func main() {
    s := SecureHandler{}
    s.Log("Hello")
    // Output:
    // Base Log: Hello
    // Secure Log: Hello
}
```

## Key Points
- This is not true overriding (as in OOP), but method shadowing via composition.
- You must call the embedded method explicitly if you want to chain or extend its behavior.
- If you embed a struct anonymously (e.g., `type SecureHandler struct { BaseHandler }`), you can call `s.Log()` directly, and Go will use the outermost method.

## Pros and Cons
| Aspect         | Go Composition (Shadowing) | OOP Overriding         |
|----------------|---------------------------|-----------------------|
| Inheritance    | No                        | Yes                   |
| Overriding     | Manual (shadowing)        | Automatic             |
| Super call     | Explicit (call embedded)  | `super` keyword       |
| Flexibility    | More explicit             | More automatic        |

## Summary
- Go does not have automatic method overriding, but you can shadow and extend methods using composition.
- This approach is explicit and flexible, but requires manual delegation if you want to call the embedded method.

---

**References:**
- [Effective Go: Embedding](https://golang.org/doc/effective_go#embedding)
- [Go Blog: Embedding](https://go.dev/blog/embedded)
