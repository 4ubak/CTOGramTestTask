package test

import (
	"testing"

	"github.com/4ubak/CTOGramTestTask/cmd"
	"github.com/4ubak/CTOGramTestTask/internal/adapters/db/pg"
)

var db *sql.DB

func Test_test(t *testing.T) {
	value := cmd.Execute()
	if value == "Successfully connected!" {
		t.Logf("Success")
	} else {
		t.Errorf("Failed %s", value)
	}
	pg.Show(db)
}
