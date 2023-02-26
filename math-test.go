package main

import "testing"

func TesteSoma(t *testing.T) {
	total := Soma(5, 5)
	if total != 10 {
		t.Errorf("Resultado da soma é inválido, resultado: %d. Esperado: %d.", total, 10)
	}
}
