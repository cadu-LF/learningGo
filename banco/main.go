package main

import (
	"encoding/json"
	"fmt"
	"io"
	"learn/clientes"
	"learn/contas"
	"net/http"
)

type Response struct {
	Key string `json:"key"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	myResponse := Response{}

	value, _ := io.ReadAll(r.Body)
	json.Unmarshal(value, &myResponse)

	io.WriteString(w, myResponse.Key)

}

func server() {
	http.HandleFunc("/hello", Hello)

	http.ListenAndServe(":8080", nil)
}

func main() {
	server()
	classes()
}

func classes() {
	conta := contas.ContaCorrente{
		Titular: clientes.Titular{Nome: "", Cpf: "", Profissao: ""},
		Agencia: 0,
		Conta:   0,
	}
	conta2 := new(contas.ContaCorrente)

	conta.Transferir(100, conta2)

	fmt.Println(conta)
	fmt.Println(conta2)
}
