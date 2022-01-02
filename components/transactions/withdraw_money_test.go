package transactions

import (
	"fmt"
	"testing"

	clientsrepository "github.com/crdev13/moneyprocessing/components/clients/repository"
	"github.com/crdev13/moneyprocessing/components/transactions/dto/input"
	transactionsrepository "github.com/crdev13/moneyprocessing/components/transactions/repository"
)

func TestWithdrawMoney(t *testing.T) {
	clientsRepository := clientsrepository.NewInMemoryClientsRepository()
	transactionsRepository := transactionsrepository.NewInMemoryTransactionsRepository()
	hasErr := false
	for number, testCase := range input.WithdrawMoneyTestCases {
		command, err := NewWithdrawMoney(clientsRepository, transactionsRepository, testCase.Request)
		if (err != nil && !testCase.NeedError) || err == nil && testCase.NeedError {
			hasErr = true
		}
		if err != nil {
			fmt.Printf("Case #%v:\n", (number + 1))
			hasErr = printWithdrawMoneyTestDetails(testCase, err)
			fmt.Printf("Case #%v(end)\n", (number + 1))
		} else {
			hasErr = false
			err := command.Execute()
			if (err != nil && !testCase.NeedError) || err == nil && testCase.NeedError {
				hasErr = true
			}
			fmt.Printf("Case #%v:\n", (number + 1))
			hasErr = printWithdrawMoneyTestDetails(testCase, err)
			fmt.Printf("Case #%v(end)\n", (number + 1))
		}
	}
	if hasErr {
		t.Errorf("TestWithdrawMoney HAS NOT PASSED TESTS ")
		return
	}
	fmt.Printf("TestWithdrawMoney HAS PASSED TESTS ")
}

func printWithdrawMoneyTestDetails(
	testCase *input.WithdrawMoneyTestCase,
	err error,
) bool {
	hasErr := true
	if testCase.NeedError && err != nil {
		if err.Error() != testCase.MsgError {
			fmt.Printf("%v NOT PASS TEST\n", testCase.Name)
			return hasErr
		}
		fmt.Printf("%v PASS TEST => message: %v\n", testCase.Name, testCase.MsgError)
		return !hasErr
	}
	if testCase.NeedError && err == nil {
		fmt.Printf("%v NOT PASS TEST => message: %v\n", testCase.Name, testCase.MsgError)
		return hasErr
	}
	if !testCase.NeedError && err != nil {
		fmt.Printf("%v NOT PASS TEST => message: %v\n", testCase.Name, err)
		return hasErr
	}
	fmt.Printf("%v PASS TEST\n", testCase.Name)
	return !hasErr
}
