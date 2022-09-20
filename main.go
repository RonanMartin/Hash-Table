package main

import (
	"fmt"
)

// Como está em mais de um arquivo, para funcionar é necessário fazer o comando: go run .
// Desta forma Go vai ler todos os arquivos da pasta e pegar também o arquivo hash_table.go

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Sexo      string
}

func main() {
	pessoas := []Pessoa{
		{"João", "Batista", 50, "M"},
		{"João", "dos Santos", 45, "M"},
		{"Paulo", "De Tarso", 35, "M"},
		{"Maria", "Madalena", 40, "F"},
		{"Ronan", "Martin", 36, "M"},
	}

	table := HashTable{}

	keys := make([]int, len(pessoas))
	for i, pessoa := range pessoas {
		keys[i] = table.Put(pessoa)
	}

	for _, key := range keys {
		ps := table.Get(key)
		for _, p := range ps {
			fmt.Println(p.Nome, p.Sobrenome)
		}
	}

	joão := table.Search("João")
	fmt.Println(joão)

	table.Remove("João")

	fmt.Println("\n", "Após o Remove e com o Put:", "\n")

	table.Put(Pessoa{"Ronan", "de Lima", 36, "M"})

	for _, key := range keys {
		ps := table.Get(key)
		for _, p := range ps {
			fmt.Println(p.Nome, p.Sobrenome)
		}
	}

	ronan := table.Search("Ronan")
	fmt.Println(ronan)

}
