package entity

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	account := &Account{
		ID:            1,
		Name:          "John Doe",
		Balance:       100.0,
		Status:        "Active",
		AccountNumber: "123456789",
	}

	account.Withdraw(50.0)
	if account.Balance != 50.0 {
		t.Errorf("Expected balance to be 50.0, got %f", account.Balance)
	}
	if account.Status != "Active" {
		t.Errorf("Expected status to be 'Active', got '%s'", account.Status)
	}

	account.Withdraw(50.0)
	if account.Balance != 0.0 {
		t.Errorf("Expected balance to be 0.0, got %f", account.Balance)
	}
	if account.Status != "Inactive" {
		t.Errorf("Expected status to be 'Inactive', got '%s'", account.Status)
	}
}

func TestDeposit(t *testing.T) {
	account := &Account{
		ID:            1,
		Name:          "John Doe",
		Balance:       100.0,
		Status:        "Active",
		AccountNumber: "123456789",
	}

	account.Deposit(50.0)
	if account.Balance != 150.0 {
		t.Errorf("Expected balance to be 150.0, got %f", account.Balance)
	}
	if account.Status != "Active" {
		t.Errorf("Expected status to be 'Active', got '%s'", account.Status)
	}
}

// create failed Transfer test
func TestTransfer(t *testing.T) {
	account1 := &Account{
		ID:            1,
		Name:          "John Doe",
		Balance:       100.0,
		Status:        "Active",
		AccountNumber: "123456789",
	}
	account2 := &Account{
		ID:            2,
		Name:          "Jane Doe",
		Balance:       50.0,
		Status:        "Active",
		AccountNumber: "987654321",
	}
	err := account1.Transfer(30.0, account2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if account1.Balance != 70.0 {
		t.Errorf("Expected account1 balance to be 70.0, got %f", account1.Balance)
	}
	if account2.Balance != 80.0 {
		t.Errorf("Expected account2 balance to be 80.0, got %f", account2.Balance)
	}
	if account1.Status != "Active" {
		t.Errorf("Expected account1 status to be 'Active', got '%s'", account1.Status)
	}
	if account2.Status != "Active" {
		t.Errorf("Expected account2 status to be 'Active', got '%s'", account2.Status)
	}
	// Test insufficient balance for transfer
	err = account1.Transfer(100.0, account2)
	if err == nil {
		t.Error("Expected error for insufficient balance, got nil")
	}
	if account1.Balance != 70.0 {
		t.Errorf("Expected account1 balance to remain 70.0, got %f", account1.Balance)
	}
	if account2.Balance != 80.0 {
		t.Errorf("Expected account2 balance to remain 80.0, got %f", account2.Balance)
	}
	if account1.Status != "Active" {
		t.Errorf("Expected account1 status to remain 'Active', got '%s'", account1.Status)
	}
}

// Benchmark: Compare passing Account by value vs by reference (pointer)

func BenchmarkProcessBalanceByValue(b *testing.B) {
	account := Account{
		ID:            1,
		Name:          "John Doe",
		Balance:       100.0,
		Status:        "Active",
		AccountNumber: "123456789",
	}
	for i := 0; i < b.N; i++ {
		ProcessBalanceByValue(account, 50.0)
	}
}

func BenchmarkProcessBalanceByReference(b *testing.B) {
	account := &Account{
		ID:            1,
		Name:          "John Doe",
		Balance:       100.0,
		Status:        "Active",
		AccountNumber: "123456789",
	}
	for i := 0; i < b.N; i++ {
		ProcessBalanceByReference(account, 50.0)
	}
}
