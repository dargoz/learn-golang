package composition

import "fmt"

type BaseHandler struct {
}

func (b BaseHandler) Log(message string) {
	fmt.Println("Base Log:", message)
}

type SecureHandler struct {
	BaseHandler BaseHandler
}

// SecureHandler overrides the Log method of BaseHandler
// to provide additional functionality, such as logging securely.
func (s SecureHandler) Log(message string) {
	s.BaseHandler.Log(message)
	fmt.Println("Secure Log:", message)
}
