package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)
	Desativacao(wesley)
	wesley.Address.Cidade = "São Paulo"

	fmt.Printf("Nome:%s Idade:%d Ativo: %t\n", wesley.Nome, wesley.Idade, wesley.Ativo)
}
