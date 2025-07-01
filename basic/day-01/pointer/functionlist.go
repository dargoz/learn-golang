package pointer

func WithoutPointer(a int) {
	// This function does not use pointers.
	a = a + 100
}

func WithPointer(x *int) {
	// This function uses pointers.
	*x += 20
	// Now x value is plus 20
}
