# Arrays vs Slices in Go

## Arrays
- An array is a fixed-length sequence of elements of the same type.
- The length of an array is part of its type and cannot be changed after creation.
- Arrays are value types: assigning one array to another copies all elements.
- Example:
  ```go
  var arr [3]int = [3]int{1, 2, 3}
  arr[0] = 10
  ```

## Slices
- A slice is a dynamically-sized, flexible view into the elements of an array.
- Slices are reference types: assigning one slice to another points both to the same underlying array.
- Slices have both a length and a capacity, and can grow or shrink as needed.
- Example:
  ```go
  s := []int{1, 2, 3}
  s = append(s, 4)
  ```

## Key Differences
| Feature         | Array                        | Slice                        |
|----------------|------------------------------|------------------------------|
| Length         | Fixed                        | Dynamic                      |
| Type           | [N]T (e.g., [3]int)          | []T (e.g., []int)            |
| Value/Ref      | Value type (copies on assign)| Reference type (shares data)  |
| Memory         | Single block                 | Backed by array, flexible    |
| Can grow/shrink| No                           | Yes (with append)            |

## When to Use Each
- **Use Arrays:**
  - When you need a fixed-size collection known at compile time.
  - For low-level programming, performance-critical code, or when working with C interop.
  - When you want value semantics (copy on assignment).
- **Use Slices:**
  - For most cases in Go, especially when the size can change or is not known in advance.
  - When you want to pass collections to functions without copying all elements.
  - For idiomatic Go code, slices are preferred for lists, buffers, and collections.

## Example Use Cases
- **Array:**
  - Storing a fixed number of configuration values.
  - Implementing a ring buffer with a known size.
- **Slice:**
  - Managing a list of user inputs.
  - Building up a collection of results from a loop or function.

---

**Summary:**
- Use arrays for fixed-size, value-type collections.
- Use slices for flexible, dynamic, and idiomatic Go code.
