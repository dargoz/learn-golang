package calculator

import "testing"

// generate test for all calculator functions
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}
func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}
func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}
func TestDivide(t *testing.T) {
	result, err := Divide(6, 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}

	_, err = Divide(6, 0)
	if err == nil {
		t.Error("Expected error for division by zero, got nil")
	}
}
func TestAddAndSubtract(t *testing.T) {
	add, subtract := AddAndSubtract(5, 3)
	if add != 8 {
		t.Errorf("Expected 8, got %d", add)
	}
	if subtract != 2 {
		t.Errorf("Expected 2, got %d", subtract)
	}
}
func TestAddAndSubtractNegative(t *testing.T) {
	add, subtract := AddAndSubtract(-5, -3)
	if add != -8 {
		t.Errorf("Expected -8, got %d", add)
	}
	if subtract != -2 {
		t.Errorf("Expected -2, got %d", subtract)
	}
}
func TestAddAndSubtractZero(t *testing.T) {
	add, subtract := AddAndSubtract(0, 0)
	if add != 0 {
		t.Errorf("Expected 0, got %d", add)
	}
	if subtract != 0 {
		t.Errorf("Expected 0, got %d", subtract)
	}
}
func TestAddAndSubtractMixed(t *testing.T) {
	add, subtract := AddAndSubtract(5, -3)
	if add != 2 {
		t.Errorf("Expected 2, got %d", add)
	}
	if subtract != 8 {
		t.Errorf("Expected 8, got %d", subtract)
	}
}
func TestAddAndSubtractMixedNegative(t *testing.T) {
	add, subtract := AddAndSubtract(-5, 3)
	if add != -2 {
		t.Errorf("Expected -2, got %d", add)
	}
	if subtract != -8 {
		t.Errorf("Expected -8, got %d", subtract)
	}
}
