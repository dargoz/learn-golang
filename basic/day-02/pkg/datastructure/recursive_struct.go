package datastructure

// A recursive struct is a struct that contains a field that is of the same type as the struct itself.
// This allows the struct to reference itself, creating a recursive relationship.
// Recursive structs are useful for representing hierarchical data structures, such as trees or linked lists.
type Branch struct {
	Name          string
	Location      string
	Contact       string
	ChildBranches []*Branch
	ParentBranch  *Branch
}
