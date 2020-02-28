package core

import "github.com/4ubak/CTOGramTestTask/internal/interfaces"

type St struct {
	db interfaces.Db
}

func NewSt(db interfaces.Db) *St {
	return &St{
		db: db,
	}
}