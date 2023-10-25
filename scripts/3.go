//Array

var tarefas [5]string

//Definindo Slice

var listaTarefas []string

//Adicionando elementos ao Slice

listaTarefas = append(listaTarefas, "Comprar pão")

#### Passo 2: Mapas
1. **Definindo um Mapa:** Um mapa é uma coleção de pares chave-valor. Por exemplo, um mapa para rastrear o status das tarefas:
```go
var statusTarefas map[string]string
statusTarefas = make(map[string]string)
```

//Adicionando e Acessando Itens

statusTarefas["Comprar pão"] = "pendente"
status := statusTarefas["Comprar pão"]

//Concatenando Strings
mensagem := "Tarefa" + "Comprar pão"






