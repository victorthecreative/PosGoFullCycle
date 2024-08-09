package main

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) Simualdor(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {

	NewConta().Simualdor(200)

	conta := Conta{saldo: 100}
	conta.Simualdor(500)
	println(conta.saldo)

}
