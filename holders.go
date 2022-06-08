package holders

import "fmt"

type CurrenteHolder struct {
	Holder        string
	AgencyNumber  int
	AccountNumber int
	BankBalance   float64
}

func (holder *CurrenteHolder) DrawOut(value float64) bool {
	var success bool = true

	if value > 0 {
		if holder.BankBalance < value {
			success = false
			fmt.Println("Insufficient funds")
		} else {
			holder.BankBalance -= value
			fmt.Println("Sucessful drawout")
		}
	} else {
		fmt.Println("Enter a value greater than zero")
	}

	return success
}

func (holder *CurrenteHolder) Deposit(value float64) {
	holder.BankBalance += value

	fmt.Println("Successful deposit")
}

func (holderOwner *CurrenteHolder) TransferValue(value float64, holder *CurrenteHolder) {
	if holderOwner.BankBalance < value {
		fmt.Println("Insufficient funds")
	} else {
		holderOwner.BankBalance -= value
		holder.Deposit(value)
	}
}
