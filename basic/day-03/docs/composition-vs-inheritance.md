# Composition vs Inheritance in Go

## Composition in Go
- Go does not support traditional class-based inheritance.
- Instead, Go uses composition: you embed one struct type inside another.
- The embedded struct's fields and methods become accessible on the outer struct.

### Example: Composition
```go
type AuditInfo struct {
    ByUserID int
    ByUser   string
    AtTime   time.Time
    Notes    string
}

type Transaction struct {
    Amount    float64
    AuditInfo // Embedded struct (composition)
}

func (a *AuditInfo) Summary() string {
    return a.ByUser + " at " + a.AtTime.Format(time.RFC3339)
}

// Now Transaction has access to AuditInfo's fields and methods:
t := Transaction{}
t.Summary() // Calls AuditInfo's Summary method
```

## Pros of Composition (in Go and in general)
- **Flexible code reuse:** You can embed multiple structs, not just one parent.
- **No tight coupling:** You avoid the fragile base class problem of inheritance.
- **Explicit relationships:** Composition is more explicit and easier to reason about.
- **Favors interfaces:** Go encourages using interfaces for polymorphism, not inheritance.
- **Simpler hierarchy:** No deep inheritance trees.

## Cons of Composition (compared to inheritance)
- **No automatic method overriding:** You can't override embedded methods like in OOP inheritance; you must explicitly define new methods.
- **No is-a relationship:** Composition models has-a, not is-a. Sometimes, inheritance is more natural for true is-a relationships.
- **Manual delegation:** If you want to change embedded behavior, you must write wrapper methods.

## Summary Table
| Aspect         | Inheritance (OOP)         | Composition (Go)           |
|----------------|---------------------------|----------------------------|
| Code reuse     | via parent class          | via embedding/fields       |
| Polymorphism   | via subclassing           | via interfaces             |
| Hierarchy      | Deep, can be complex      | Flat, explicit             |
| Overriding     | Automatic                 | Manual (redefine methods)  |
| Coupling       | Tight                     | Loose                     |
| Flexibility    | Less (single parent)      | More (multiple embeds)     |

## When to Use Composition in Go
- When you want to reuse code from multiple sources.
- When you want to avoid the complexity and pitfalls of inheritance.
- When you want to model has-a relationships.
- When you want to favor composition and interfaces for flexible, decoupled design.

---

**References:**
- [Go Blog: Go's Type System](https://blog.golang.org/types)
- [Go Blog: Embedding](https://go.dev/blog/embedded)
- [Effective Go: Embedding](https://golang.org/doc/effective_go#embedding)
