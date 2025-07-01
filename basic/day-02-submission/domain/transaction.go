package domain

import (
	"github.com/dargoz/simplebank/model"
)

type Transaction interface {
	Apply(acc *model.Account) error
	Description() string
}
