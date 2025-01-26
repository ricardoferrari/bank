package main

// Bank is a bank with a number of accounts.
import (
	"encoding/json"
	"fmt"
	"os"
)

type Bank struct {
	accounts map[string]int
}

func (b *Bank) writeBankToFile() {
	// Open file for writing
	file, err := os.Create("bank.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Converts bank object to json string
	bankJson, err := json.Marshal(b.accounts)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bankJson))
	// Save bank object to file
	_, err = file.WriteString(string(bankJson))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}

func readBankFromFile() *Bank {
	// Open file for reading
	file, err := os.Open("bank.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// Read bank object from file
	var bank Bank
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&bank.accounts)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// _, err = fmt.Fscanf(file, "%v\n", &bank)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil
	// }
	defer file.Close()
	return &bank
}

// NewBank creates a new bank with no accounts.
func NewBank() *Bank {
	return &Bank{accounts: make(map[string]int)}
}

// CreateAccount creates a new account with the given name and a balance of 0.
func (b *Bank) CreateAccount(name string) {
	b.accounts[name] = 0
}

// Deposit deposits the given amount into the account with the given name.
func (b *Bank) Deposit(name string, amount int) {
	b.accounts[name] += amount
}

// Withdraw withdraws the given amount from the account with the given name.
func (b *Bank) Withdraw(name string, amount int) {
	if b.accounts[name] < amount {
		fmt.Println("Insufficient balance")
		return
	}
	b.accounts[name] -= amount
}

// Transfer transfers the given amount from the account with name1 to the account with name2.
func (b *Bank) Transfer(name1, name2 string, amount int) {
	if b.accounts[name1] < amount {
		fmt.Println("Insufficient balance")
		return
	}
	b.accounts[name1] -= amount
	b.accounts[name2] += amount
}

// Balance prints the balance of the account with the given name.
func (b *Bank) Balance(name string) {
	fmt.Println(b.accounts[name])
}

// PrintAccounts prints the names and balances of all accounts.
func (b *Bank) PrintAccounts() {
	for name, balance := range b.accounts {
		fmt.Printf("%s: %d\n", name, balance)
	}
}

func main() {
	bank := NewBank()

	var exitFlag bool = false
	var choice int

	for !exitFlag {
		// Clear screen before printing the menu
		fmt.Print("\033[H\033[2J")
		fmt.Println("Choose the option:")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Transfer")
		fmt.Println("5. Balance")
		fmt.Println("6. Print Accounts")
		fmt.Println("7. Save Data")
		fmt.Println("8. Recover Data")
		fmt.Println("9. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("Enter the name of the account holder:")
			var name string
			fmt.Scanln(&name)
			bank.CreateAccount(name)
		case 2:
			fmt.Println("Enter the name of the account holder to deposit:")
			var name string
			fmt.Scanln(&name)
			fmt.Println("Enter the amount to deposit:")
			var amount int
			fmt.Scanln(&amount)
			bank.Deposit(name, amount)
		case 3:
			fmt.Println("Enter the name of the account holder to withdraw:")
			var name string
			fmt.Scanln(&name)
			fmt.Println("Enter the amount to withdraw:")
			var amount int
			fmt.Scanln(&amount)
			bank.Withdraw(name, amount)
		case 4:
			fmt.Println("Enter the name of the account holder from which you want to transfer:")
			var name1 string
			fmt.Scanln(&name1)
			fmt.Println("Enter the name of the account holder to which you want to transfer:")
			var name2 string
			fmt.Scanln(&name2)
			fmt.Println("Enter the amount to transfer:")
			var amount int
			fmt.Scanln(&amount)
			bank.Transfer(name1, name2, amount)
		case 5:
			fmt.Println("Enter the name to verify the balance:")
			var name string
			fmt.Scanln(&name)
			bank.Balance(name)
			// Press enter to continue
			fmt.Println("Press enter to continue...")
			fmt.Scanln()
		case 6:
			bank.PrintAccounts()
			// Press enter to continue
			fmt.Println("Press enter to continue...")
			fmt.Scanln()
		case 7:
			bank.writeBankToFile()
			fmt.Println("Saved enter to continue...")
			fmt.Scanln()
		case 8:
			bank = readBankFromFile()
		case 9:
			fmt.Println("Thanks for choosing our bank!")
			exitFlag = true
		}
	}

}
