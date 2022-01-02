package clients

import (
	"fmt"
	"testing"

	"github.com/crdev13/moneyprocessing/components/clients/dto/input"
	"github.com/crdev13/moneyprocessing/components/clients/dto/output"
	clientsRepository "github.com/crdev13/moneyprocessing/components/clients/repository"
)

func TestCreateClient(t *testing.T) {
	repository := clientsRepository.NewInMemoryClientsRepository()
	hasErr := false
	for number, testCase := range input.CreateClientTestCases {
		command, err := NewCreateClient(repository, testCase.Request)
		if (err != nil && !testCase.NeedError) || err == nil && testCase.NeedError {
			hasErr = true
		}
		if err != nil {
			fmt.Printf("Case #%v:\n", (number + 1))
			hasErr = printCreateClientTestDetails(nil, testCase, err)
			fmt.Printf("Case #%v(end)\n", (number + 1))
		} else {
			hasErr = false
			response, err := command.Execute()
			if (err != nil && !testCase.NeedError) || err == nil && testCase.NeedError {
				hasErr = true
			}
			fmt.Printf("Case #%v:\n", (number + 1))
			hasErr = printCreateClientTestDetails(response, testCase, err)
			fmt.Printf("Case #%v(end)\n", (number + 1))
		}
	}
	if hasErr {
		t.Errorf("TestCreateClient HAS NOT PASSED TESTS ")
		return
	}
	fmt.Printf("TestCreateClient HAS PASSED TESTS ")
}

func printCreateClientTestDetails(
	response *output.CreateClientResponse,
	testCase *input.CreateClientTestCase,
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

		fmt.Printf("%v PASS TEST\n clientID: %v\n", testCase.Name, response.ClientID)
	} else {
		fmt.Printf("%v PASS TEST\n", testCase.Name)
	}
	return !hasErr
}
