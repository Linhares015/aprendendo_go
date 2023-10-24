var nome string = "João"
var idade int = 10
var versao float32 = 1.1
var ehLegal bool = true

//Operadores Aritméticos
// + - * / %

//Operadores de Atribuição
// = += -= *= /= %=

//Operadores de Comparação
// == != > < >= <=

//Operadores Lógicos
// && || !

//Operadores Bit a Bit
// & | ^ << >>

// Condicionais
// if else else if
if idade >= 18 {
	fmt.PrintIn("Maior de idade")
} else {
	fmt.PrintIn("Menor de idade")
}

// switch case
switch dia {
case "segunda":
	fmt.PrintIn("Hoje é segunda")
case "terça":
	fmt.PrintIn("Hoje é terça")
default:
	fmt.PrintIn("Hoje é outro dia")
}

//Loops
// for
for i := 0; i < 10; i++ {
	fmt.PrintIn(i)
}

//Funções
func saudacao(nome string) {
	fmt.PrintIn("Olá, " + nome)
}

//Chamando a função
saudacao("João")




