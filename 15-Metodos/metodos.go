package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p pessoa) getIdade() uint8 {
	return p.idade
}

func (p *pessoa) setNome(novoNome string) {
	p.nome = novoNome
}

func (p *pessoa) setSobrenome(novoSobrenome string) {
	p.sobrenome = novoSobrenome
}

func (p *pessoa) setIdade(novaIdade uint8) {
	p.idade = novaIdade
}

func (p pessoa) maiorDeIdade() bool {
	return p.idade >= 18
}

func (p *pessoa) fazerAniversario() {
	p.idade++
}

func main() {
	p1 := pessoa{"João", "Silva", 18}
	fmt.Println(p1.getNomeCompleto())
	fmt.Println(p1.getIdade())
	fmt.Printf("Olá, meu nome é %s e tenho %d anos", p1.getNomeCompleto(), p1.getIdade())
	p1.setNome("Pedro")
	p1.setSobrenome("Santos")
	p1.setIdade(20)
	fmt.Println(" ")

	fmt.Println(p1.getNomeCompleto())
	fmt.Println(p1.getIdade())
	fmt.Println(p1.maiorDeIdade())
	p1.fazerAniversario()
	fmt.Println(p1.getIdade())
	fmt.Printf("Olá, meu nome é %s e tenho %d anos", p1.getNomeCompleto(), p1.getIdade())
}
