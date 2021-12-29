package main

type Bank struct {
	Balance int
}

func Deposit (ba *Bank, amount int) {
	ba.Balance += amount
}

func Withdraw (ba *Bank, amount int) {
	ba.Balance -= amount
}

func main() {
	ba := &Bank{Balance: 0}
	var commands []func()
	commands = append(commands, func() {Deposit(ba, 100)})
	commands = append(commands, func() {Withdraw(ba, 50)})

	for _, cmd := range commands {
		cmd()
	}
}