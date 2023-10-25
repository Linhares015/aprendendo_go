package main

import "fmt"

type Tarefa struct {
	ID     int
	Nome   string
	Status string
}

var tarefas []Tarefa

func adicionarTarefa(nome string) {
	id := len(tarefas) + 1
	tarefa := Tarefa{
		ID:     id,
		Nome:   nome,
		Status: "Pendente",
	}
	tarefas = append(tarefas, tarefa)
}

func listarTarefas() {
	for _, tarefa := range tarefas {
		fmt.Println("ID:", tarefa.ID, "Nome:", tarefa.Nome, "Status:", tarefa.Status)
	}
}

func atualizarTarefa(id int, novoNome string) {
	for i, tarefa := range tarefas {
		if tarefa.ID == id {
			tarefas[i].Nome = novoNome
		}
	}
}

func removerTarefa(id int) {
	for i, tarefa := range tarefas {
		if tarefa.ID == id {
			tarefas = append(tarefas[:i], tarefas[i+1:]...)
		}
	}
}

func main() {
	var opcao int
	for {
		fmt.Println("1. Adicionar Tarefa")
		fmt.Println("2. Listar Tarefa")
		fmt.Println("3. Atualizar Tarefa")
		fmt.Println("4. Remover Tarefa")
		fmt.Println("5. Sair")
		fmt.Println("Escolha uma opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			var nome string
			fmt.Println("Nome da Tarefa: ")
			fmt.Scan(&nome)
			adicionarTarefa(nome)
		case 2:
			listarTarefas()
		case 3:
			var id int
			var novoNome string
			fmt.Println("ID da Tarefa: ")
			fmt.Scan(&id)
			fmt.Print("Novo Nome: ")
			fmt.Scan(&novoNome)
			atualizarTarefa(id, novoNome)
		case 4:
			var id int
			fmt.Print("ID da Tarefa: ")
			fmt.Scan(&id)
			removerTarefa(id)
		case 5:
			return
		default:
			fmt.Println("Opção inválida!")
		}

	}
}
