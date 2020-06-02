package interfaces

// Definindo tipos para exemplo
type Autor struct {
	Nome  string
	Idade int32
}

type Computador struct {
	Modelo string
	Marca  string
}

// Atribuindo os metodos aos tipos (structs) criados
func (c Computador) Escreve() string {
	return "Texto escrito pelo PC " + c.Marca + " " + c.Modelo
}

func (a Autor) Escreve() string {
	return "Livro escrito por " + a.Nome
}

// A interface livro ira conter os metodos chamados "escreve" e que retornem uma string
type Livro interface {
	Escreve() string
}

// Uma interface pode ser composta por outras interfaces
type Biblioteca interface {
	Livro
}

// func main() {
// 	dono := Autor{
// 		Nome:  "Joao da Silva",
// 		Idade: 30,
// 	}

// 	maquina := Computador{
// 		Modelo: "Inspiron",
// 		Marca:  "Dell",
// 	}

// 	texto := caneta(dono)
// 	textoMaq := Digitar(maquina)

// 	fmt.Println(texto)
// 	fmt.Println(textoMaq)
// }

// Os metodos abaixo implementam as interfaces criadas (No Go nao existe classes)
func caneta(l Livro) string {
	return l.Escreve()
}

func Digitar(b Biblioteca) string {
	return b.Escreve()
}
