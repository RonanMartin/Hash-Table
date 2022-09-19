package main

import (
	"fmt"
	"sync"
)

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

type HashTable struct {
	items map[int][]Pessoa
	lock  sync.RWMutex
}

func hash(nome string) (key int) {
	for _, letra := range nome {
		key = 31*key + int(letra)
	}

	return
}

func (ht *HashTable) Put(pessoa Pessoa) int {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	key := hash(pessoa.Nome)
	if ht.items == nil {
		ht.items = make(map[int][]Pessoa)
	}
	ht.items[key] = append(ht.items[key], pessoa)
	return key
}

func (ht *HashTable) Remove(nome string) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	key := hash(nome)
	delete(ht.items, key)
}

func (ht *HashTable) Get(key int) []Pessoa {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return ht.items[key]
}

func (ht *HashTable) Search(nome string) []Pessoa {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	key := hash(nome)
	return ht.items[key]
}
