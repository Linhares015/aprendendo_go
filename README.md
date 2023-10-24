# aprendendo_go

Introdução ao Go: Guia Passo a Passo
Passo 1: Instalação do Go

    Acesse o site oficial do Go em golang.org.
    Navegue até a seção de downloads e escolha a versão adequada para o seu sistema operacional (Windows, macOS ou Linux).
    Baixe o instalador e siga as instruções na tela para instalar o Go em seu computador.
    Após a instalação, abra o terminal ou prompt de comando e digite go version para verificar se o Go foi instalado corretamente. Você deve ver a versão do Go que instalou.

Passo 2: Estrutura Básica de um Programa Go

    Abra seu editor de código de preferência.
    Crie um novo arquivo chamado hello.go.
    No arquivo, escreva o seguinte código:

```go

package main

import "fmt"

func main() {
    fmt.Println("Olá, Mundo!")
}
```

Este é um programa Go simples que imprime "Olá, Mundo!" na tela. Aqui está o que cada parte faz:

    package main: Define o pacote principal onde o programa começa a execução.
    import "fmt": Importa o pacote "fmt", que contém funções para formatar e imprimir informações.
    func main(): A função principal onde o programa começa a execução.
    fmt.Println(...): Uma função que imprime a string fornecida na tela.

Passo 3: Compilação e Execução do Programa

    Salve o arquivo hello.go.
    Abra o terminal ou prompt de comando e navegue até o diretório onde o arquivo hello.go está localizado.
    Digite o comando go run hello.go e pressione Enter.
    Você deve ver a mensagem "Olá, Mundo!" impressa na tela. Isso significa que você compilou e executou com sucesso seu primeiro programa Go!

2. Fundamentos da Linguagem Go: Guia Passo a Passo
Passo 1: Variáveis e Tipos de Dados

    Definindo uma Variável: Em Go, você pode definir uma variável usando a palavra-chave var. Por exemplo:

go

var nome string = "João"

    Tipos de Dados Básicos: Go tem vários tipos de dados básicos, incluindo:
        int: para números inteiros.
        float64: para números de ponto flutuante.
        bool: para valores booleanos (verdadeiro ou falso).
        string: para cadeias de caracteres.

    Declarando Variáveis de Forma Curta: Em Go, você também pode declarar e inicializar uma variável de forma curta usando :=. Por exemplo:

go

idade := 25

Passo 2: Operadores

    Operadores Aritméticos: Go suporta operadores aritméticos básicos como +, -, *, / e %.
    Operadores de Comparação: Você pode comparar duas variáveis usando operadores como ==, !=, <, >, <= e >=.

Passo 3: Estruturas de Controle

    Condicional IF: Em Go, você pode usar a estrutura if para testar condições. Por exemplo:

go

if idade >= 18 {
    fmt.Println("Você é maior de idade.")
} else {
    fmt.Println("Você é menor de idade.")
}

    Switch Case: Go também suporta a estrutura switch para testar múltiplas condições. Por exemplo:

go

switch dia {
case "segunda":
    fmt.Println("Hoje é segunda-feira!")
case "terça":
    fmt.Println("Hoje é terça-feira!")
default:
    fmt.Println("Não é nem segunda nem terça.")
}

    Loops: Go tem uma única palavra-chave para loops, que é for. Por exemplo:

go

for i := 0; i < 10; i++ {
    fmt.Println(i)
}

Passo 4: Funções

    Definindo uma Função: Em Go, você pode definir uma função usando a palavra-chave func. Por exemplo:

go

func saudacao(nome string) {
    fmt.Println("Olá,", nome)
}

    Chamando uma Função: Para chamar a função que você definiu, use o nome da função seguido de parênteses. Por exemplo:

go

saudacao("Maria")
