// operations.go
package operations

import (
	"bankingmod/models"
	"bankingmod/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

var accounts []models.BankAccount
var persons []models.Person
var fullName []string
var accountTypes []models.AccountType

func CreateDefaultAccountTypes() {
	// Create instances of AccountType
	accountTypes = []models.AccountType{
		{Id: uuid.New(), Name: "Conta Corrente"},
		{Id: uuid.New(), Name: "Conta Poupança"},
		{Id: uuid.New(), Name: "Conta Fidelidade"},
	}

	for _, accountType := range accountTypes {
		fmt.Printf("ID: %s, Name: %s\n", accountType.Id.String(), accountType.Name)
	}
}

func AskNameUser() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Digite seu nome completo:")
		nome, _ := reader.ReadString('\n') // Read input until a newline character

		fullName := strings.Fields(nome) // Split the input into fields

		if len(fullName) < 2 {
			fmt.Println("Nome completo inválido.")
			continue
		} else {
			fmt.Printf("Olá %s , bem-vindo ao Banco Xaiking !\n", nome)
			persons = append(persons, models.Person{Name: nome})
			break
		}
	}
}

func findAccount(accountNumber int) *models.BankAccount {
	for _, acc := range accounts {
		if acc.AccountNumber == accountNumber {
			return &acc
		}
	}
	return nil
}

func CreateBankAccount(accountType string, startBalanceZero bool) {
	accb := utils.NextAccountNumber()
	var amount float64 = 0

	findAccountTypeByName := findAccountType(accountTypes, func(a models.AccountType) bool {
		return a.Name == accountType
	})

	if startBalanceZero {
		amount = 0
	} else {
		amount = utils.RandomFloat(10, 1000)
	}

	fmt.Println("Criando sua Conta Bancária!")

	accounts = append(accounts, models.BankAccount{AccountNumber: accb, AccountType: findAccountTypeByName, Balance: amount, Person: persons[0]})
	amountformatted := fmt.Sprintf("%.2f", amount)

	time.Sleep(2 * time.Second)
	fmt.Println("Conta Criada! Aqui está o numero para acessa-la!", accb)
	fmt.Println("Seu primeiro deposito foi de : ", amountformatted)
}

func ShowAccountsCreated() {

	if len(accounts) == 0 {
		fmt.Println("Nenhuma conta criada!")
		return
	}

	// Create a new table
	table := [][]string{
		{"Número", "Tipo de conta", "Saldo"},
	}

	// Iterate over all accounts and add them to the table
	for _, acc := range accounts {
		table = append(table, []string{
			fmt.Sprintf("%d", acc.AccountNumber),
			acc.AccountType.Name,
			fmt.Sprintf("%.2f", acc.Balance),
		})
	}

	// Print the table
	fmt.Println("Listando todas as contas...")
	for _, row := range table {
		fmt.Println("+--------+-------------+-------------------+")
		fmt.Println(strings.Join(row, "|"))
		fmt.Println("+--------+-------------+-------------------+")
	}
}

func showAccountBalance(balance float64, msg string) {
	amountformatted := fmt.Sprintf("%.2f", balance)
	time.Sleep(1 * time.Second)
	fmt.Println(msg, amountformatted)
}

func MovimentAccount(transactionType string) {

	// Realizar crédito em uma determinada conta
	fmt.Println("Digite o número da conta:")
	var acccheck int
	_, err := fmt.Scan(&acccheck)

	if err != nil {
		fmt.Println("Erro ao ler o número da conta:", err)
		time.Sleep(1 * time.Second)
		fmt.Println("Voltando ao Menu Principal!")
		return
	}

	// Buscando a conta
	acc := findAccount(acccheck)
	fmt.Println("Buscando Conta!")
	time.Sleep(1 * time.Second)

	if acc == nil {
		fmt.Println("Conta não encontrada, verifique e tente novamente!")
		time.Sleep(1 * time.Second)
		fmt.Println("Voltando ao Menu Principal!")
		return
	}

	fmt.Println("Conta selecionada")
	time.Sleep(1 * time.Second)

	fmt.Println("Agora digite o valor da operação:")

	var amncheck float64
	_, err = fmt.Scan(&amncheck)
	if err != nil {
		fmt.Println("Erro ao ler o valor:", err)
		return
	}

	if transactionType == "C" {
		acc.Balance += amncheck
		fmt.Println("Aguarde...")
		time.Sleep(1 * time.Second)
		fmt.Println("Depósito Concluido")
	} else {

		if acc.Balance < amncheck {
			fmt.Println("Saldo insuficiente, refaça a operação!")
			return
		}

		acc.Balance -= amncheck
		fmt.Println("Aguarde...")
		time.Sleep(1 * time.Second)
		fmt.Println("Saque Concluido")
	}

	showAccountBalance(acc.Balance, "Saldo atual em conta: ")
}

func CheckAccountBalance(balanceType string) {
	// Consultar o saldo de uma conta
	var acccheck int
	_, err := fmt.Scan(&acccheck)

	if err != nil {
		fmt.Println("Erro ao ler o número da conta:", err)
		time.Sleep(1 * time.Second)
		fmt.Println("Voltando ao Menu Principal!")
		time.Sleep(1 * time.Second)
		return
	}

	// Buscando a conta
	fmt.Println("Buscando Conta!")
	time.Sleep(2 * time.Second)
	acc := findAccount(acccheck)

	if acc == nil {
		fmt.Println("Conta não encontrada, verifique e tente novamente!")
		time.Sleep(1 * time.Second)
		fmt.Println("Voltando ao Menu Principal!")
		time.Sleep(1 * time.Second)
		return
	}

	if balanceType == "balance" {
		showAccountBalance(acc.Balance, "Saldo atual da conta")
	} else {
		showAccountBalance(acc.Bonus, "Bonus atual da conta")
	}
}

func TransferToAnotherAccount(accorg int, accdes int, amncheck float64) {
	// Realizar uma transferência entre duas contas

	accOrigem := findAccount(accorg)
	accDestino := findAccount(accdes)

	if accOrigem == nil {
		fmt.Println("Conta de origem não encontrada, verifique e tente novamente!")
		time.Sleep(1 * time.Second)
		fmt.Println("Voltando ao Menu Principal!")
		time.Sleep(1 * time.Second)
		return
	}

	if accDestino == nil {
		fmt.Println("Conta de destino não encontrada, verifique e tente novamente!")
		time.Sleep(1 * time.Second)
		fmt.Println("Voltando ao Menu Principal!")
		time.Sleep(1 * time.Second)
		return
	}

	// Verifique o saldo da conta de origem antes de transferir
	fmt.Println("Iniciando transferência. Aguarde...")
	time.Sleep(2 * time.Second)

	if accOrigem.Balance < amncheck {
		fmt.Println("Saldo insuficiente na conta de origem.")
		time.Sleep(1 * time.Second)
		fmt.Println("Voltando ao Menu Principal!")
		time.Sleep(1 * time.Second)
		return
	}

	fmt.Println("Transferência Concluida!")
	accOrigem.Balance -= amncheck
	accDestino.Balance += amncheck
	showAccountBalance(accOrigem.Balance, "Saldo atual da conta corrente: ")
}

// Find account type using a custom condition function
func findAccountType(accountTypes []models.AccountType, condition func(models.AccountType) bool) *models.AccountType {
	for _, accountType := range accountTypes {
		if condition(accountType) {
			return &accountType
		}
	}
	return nil
}

func CloseAccount() {

}
