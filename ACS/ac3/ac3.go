package main

import "fmt"

type Contato struct {
	nome string
	email string
}

func (c *Contato) AlterarEmail(novoEmail string){
	c.email = novoEmail
}

func adicionarContato(contato Contato, listaContatos []*Contato) {
	for i, c := range listaContatos {
		if c == nil {
			listaContatos[i] = &contato
			fmt.Println("Contato adicionado com sucesso!")
			return
		}
	}
	fmt.Println("O array de contatos está cheio. Não foi possível adicionar o contato.")
}

func excluirContato(listaContatos []*Contato) {
	for i := len(listaContatos) - 1; i >= 0; i-- {
		if listaContatos[i] != nil {
			listaContatos[i] = nil
			fmt.Println("Contato excluído com sucesso!")
			return
		}
	}
	fmt.Println("Não há contatos para excluir.")
}

// Função para listar os contatos cadastrados
func listarContatos(listaContatos []*Contato) {
	fmt.Println("\nLista de Contatos:")
	for i := 0; i < len(listaContatos); i++ {
		contato := listaContatos[i]
		if contato != nil {
			fmt.Println("Nome:", contato.nome, "- Email:", contato.email)
		}
	}
}


func main() {
	listaContatos := make([]*Contato, 5)

	for {
		fmt.Println("\nOpções:")
		fmt.Println("1. Adicionar Contato")
		fmt.Println("2. Excluir Último Contato")
		fmt.Println("3. Sair")
		fmt.Println("4. para visualizar contatos")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			var nome, email string
			fmt.Print("Digite o nome do contato: ")
			fmt.Scanln(&nome)
			fmt.Print("Digite o e-mail do contato: ")
			fmt.Scanln(&email)

			novoContato := Contato{
				nome:  nome,
				email: email,
			}

			adicionarContato(novoContato, listaContatos)

		case 2:
			excluirContato(listaContatos)

		case 3:
			fmt.Println("Saindo...")
			return

		case 4:
			listarContatos(listaContatos)

		default:
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}


