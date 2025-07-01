package intf

type Transaction interface {
	Execute() (bool, error)
}

type Transfer struct {
}

func (t Transfer) Execute() (bool, error) {
	// Implementation of the transfer logic
	return true, nil
}
