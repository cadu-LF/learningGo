package contas

import "learn/clientes"

type ContaCorrente struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	saldo   float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	if valorSaque > 0 && valorSaque <= c.saldo {
		c.saldo -= valorSaque
		return "Saldo realizado"
	}
	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valor float64) {
	if valor > 0 {
		c.saldo += valor
	}
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor < c.saldo {
		c.saldo -= valor
		contaDestino.Depositar(valor)
		return true
	}

	return false
}

func (c *ContaCorrente) getSaldo() float64 {
	return c.saldo
}
