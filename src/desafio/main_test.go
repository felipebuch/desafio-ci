package main
import "testing"

func TestMain( t *testing.T) {
	var result int
	result = soma(5, 5)
	if result != 10 {
		t.Errorf("ERRO - A soma de 5+5 deve ser igual a 10")
	} else {
		t.Logf("SUCESSO - A soma de 5+5 deve ser igual a 10")
	}

	result = soma(10, 7)
	if result != 17 {
		t.Errorf("ERRO - A soma de 10+7 deve ser igual a 17")
	} else {
		t.Logf("SUCESSO - A soma de 10+7 deve ser igual a 17")
	}
}