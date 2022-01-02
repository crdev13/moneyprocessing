package clients

import (
	"fmt"
	"testing"

	"github.com/crdev13/moneyprocessing/components/clients/dto/input"
	"github.com/crdev13/moneyprocessing/components/clients/dto/output"
	clientsRepository "github.com/crdev13/moneyprocessing/components/clients/repository"
)

func TestCreateAccount(t *testing.T) {
	repository := clientsRepository.NewInMemoryClientsRepository()
	hasErr := false
	for number, testCase := range input.CreateAccountTestCases {
		command, err := NewCreateAccount(repository, testCase.Request)
		if (err != nil && !testCase.NeedError) || err == nil && testCase.NeedError {
			hasErr = true
		}
		if err != nil {
			fmt.Printf("Case #%v:\n", (number + 1))
			hasErr = printCreateAccountTestDetails(nil, testCase, err)
			fmt.Printf("Case #%v(end)\n", (number + 1))
		} else {
			hasErr = false
			response, err := command.Execute()
			if (err != nil && !testCase.NeedError) || err == nil && testCase.NeedError {
				hasErr = true
			}
			fmt.Printf("Case #%v:\n", (number + 1))
			hasErr = printCreateAccountTestDetails(response, testCase, err)
			fmt.Printf("Case #%v(end)\n", (number + 1))
		}
	}
	if hasErr {
		t.Errorf("TestCreateAccount HAS NOT PASSED TESTS ")
		return
	}
	fmt.Printf("TestCreateAccount HAS PASSED TESTS ")
}

func printCreateAccountTestDetails(
	response *output.CreateAccountResponse,
	testCase *input.CreateAccountTestCase,
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
	if response != nil {

		fmt.Printf("%v PASS TEST\n clientID: %v accountID: %v\n", testCase.Name, response.ClientID, response.AccountID)
	} else {
		fmt.Printf("%v PASS TEST\n", testCase.Name)
	}
	return !hasErr
}
