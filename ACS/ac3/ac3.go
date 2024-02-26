package main

import "fmt"

type Contato struct {
    nome  string
    email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
    c.email = novoEmail
}

func adicionarContato(contato *Contato, listaContatos *[5]*Contato) {
    for i, c := range listaContatos {
        if c == nil {
            listaContatos[i] = contato
            fmt.Println("Contato adicionado com sucesso!")
            return
        }
    }
    fmt.Println("O array de contatos está cheio. Não foi possível adicionar o contato.")
}

func excluirContato(listaContatos *[5]*Contato) {
    for i := len(listaContatos) - 1; i >= 0; i-- {
        if listaContatos[i] != nil {
            listaContatos[i] = nil
            fmt.Println("Contato excluído com sucesso!")
            return
        }
    }
    fmt.Println("Não há contatos para excluir.")
}

func listarContatos(listaContatos *[5]*Contato) {
    fmt.Println("\nLista de Contatos:")
    for _, contato := range listaContatos {
        if contato != nil {
            fmt.Println("Nome:", contato.nome, "- Email:", contato.email)
        }
    }
}

func main() {
    listaContatos := [5]*Contato{}

    for {
        fmt.Println("\nOpções:")
        fmt.Println("1. Adicionar Contato")
        fmt.Println("2. Excluir Último Contato")
        fmt.Println("3. Listar Contatos")
        fmt.Println("4. Sair")
        fmt.Print("Escolha uma opção: ")

        var opcao int
        fmt.Scanln(&opcao)

        switch opcao {
        case 1:
		if listaContatos[4] == nil {
			var nome, email string
			fmt.Print("Digite o nome do contato: ")
			fmt.Scanln(&nome)
			fmt.Print("Digite o e-mail do contato: ")
			fmt.Scanln(&email)
	
			novoContato := &Contato{
				nome:  nome,
				email: email,
			}
	
			adicionarContato(novoContato, &listaContatos)
		} else {
	                fmt.Println("Número máximo de contatos atingido.")
	        }

        case 2:
            excluirContato(&listaContatos)

        case 3:
            listarContatos(&listaContatos)

        case 4:
            fmt.Println("Saindo...")
            return

        default:
            fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
        }
    }
}
