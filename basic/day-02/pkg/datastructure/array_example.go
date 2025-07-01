package datastructure

import "fmt"

type Nasabah struct {
	ID    int
	Name  string
	Saldo *float64
}

func ArrayExample() {
	// declare an array of strings using the ellipsis (...) syntax means the compiler will infer the length of the array
	// the length of the array is fixed and cannot be changed
	// without ellipsis, the length of the array must be specified
	// otherwise, it will be a slice
	var bankProduct = [...]string{"Checking Account", "Savings Account", "Credit Card", "Loan", "Mortgage"}
	for i, product := range bankProduct {
		if i == 0 {
			println("Welcome to our bank! Here are our products:")
		}
		println(i+1, "-", product)
	}

	var transactionStatus = [3]string{"Pending", "Completed", "Failed"}
	for i, status := range transactionStatus {
		if i == 0 {
			println("\nTransaction Status:")
		}
		println(i+1, "-", status)
	}

	var p *[3]string = &transactionStatus // pointer length must match the array length
	fmt.Printf("\nTransaction Status (using pointer): %v\n", *p)

	saldo1, saldo2 := 1000.0, 2000.0
	// when delcare an array of pointers, the value must be a pointer and cannot be a float64 directly
	var saldoArray = [...]*float64{&saldo1, &saldo2}

	*saldoArray[0] += 500.0
	*saldoArray[1] -= 200.0

	fmt.Printf("\nSaldo after transactions: %v, %v\n", saldo1, saldo2)

	saldo := 1500.0
	nasabah := Nasabah{
		ID:    1,
		Name:  "John Doe",
		Saldo: &saldo,
	}
	fmt.Printf("\nNasabah ID: %d, Name: %s, Saldo: %.2f\n", nasabah.ID, nasabah.Name, *nasabah.Saldo)

	// multidimensional array [rows][columns]
	kurs := [2][2]float64{
		{1.0, 16300.0},
		{0.00006, 1.0},
	}
	fmt.Println("Kurs USD to IDR:", kurs[0][1])
	fmt.Println("Kurs IDR to USD:", kurs[1][0])

}
