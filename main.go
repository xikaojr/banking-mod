package main

import (
	"bankingmod/operations"
	"fmt"
	"strings"
	"time"
)

func mostrarMenu() {
	fmt.Println("------------------------------------------------")
	fmt.Println("** Menu principal **")
	fmt.Println("------------------------------------------------")

	for {
		fmt.Println("1. Criar uma conta corrente")
		fmt.Println("2. Criar uma conta poupança")
		fmt.Println("3. Criar uma conta fidelidade")
		fmt.Println("4. Depósito")
		fmt.Println("5. Saque")
		fmt.Println("6. Consultar o saldo de uma conta")
		fmt.Println("7. Consultar o bônus de uma conta fidelidade")
		fmt.Println("8. Realizar uma transferência")
		fmt.Println("9. Cancelar sua conta")
		fmt.Println("10. Imprimir número e saldo de todas as contas cadastradas")
		fmt.Println("0. Sair")

		fmt.Print("Opção: ")
		var option string

		fmt.Scanf("%s", &option)
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			// Criar uma conta bancária
			operations.CreateBankAccount("Conta Corrente", true)
		case "2":
			// Criar uma conta poupança
			operations.CreateBankAccount("Conta Poupança", false)
		case "3":
			// Criar uma conta fidelidade
			operations.CreateBankAccount("Conta Fidelidade", false)
		case "4":
			// Depósito
			operations.MovimentAccount("C")
		case "5":
			// Saque
			operations.MovimentAccount("D")
		case "6":
			// Consultar o saldo de uma conta
			operations.CheckAccountBalance("balance")
		case "7":
			// Consultar o bônus de uma conta fidelidade
			operations.CheckAccountBalance("bonus")
		case "8":
			// Realizar uma transferência
			var accorg int
			var accdes int
			var amncheck float64

			fmt.Println("Digite o numero da conta de saída ?")
			fmt.Scan(&accorg)
			fmt.Println("Digite o numero da conta de entrada ?")
			fmt.Scan(&accdes)
			fmt.Println("Digite o valor que deseja transferir:")
			fmt.Scan(&amncheck)

			operations.TransferToAnotherAccount(accorg, accdes, amncheck)
		case "9":
			// Cancelar sua conta
			operations.CloseAccount()
		case "10":
			// Imprimir número e saldo de todas as contas cadastradas
			operations.ShowAccountsCreated()
		case "0":
			// Sair
			break
		default:
			fmt.Println("Opção inválida.")
		}
	}
}

func main() {
	fmt.Println("Iniciando Sistema...")
	time.Sleep(500 * time.Millisecond)

	operations.CreateDefaultAccountTypes()

	operations.AskNameUser()

	fmt.Println("Carregando Painel...")
	time.Sleep(2 * time.Second)

	mostrarMenu()
}
