package main

// Reference: https://medium.com/wearetheledger/about-smart-contract-testing-fbf7b576bb9f
// go get github.com/stretchr/testify/assert
// go get -u --tags nopkcs11 github.com/hyperledger/fabric/core/chaincode/shim
// go test --tags nopkcs11

//$ echo $GOPATH
// /Users/olivieri/go:/Users/olivieri/git/cc-inner-folder

import (
	"fmt"
	"myproject/mylib"
	"testing"

	"github.com/stretchr/testify/assert"
	// "encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestXYZ(t *testing.T) {
	var a string = "Hello"
	var b string = "Hello"
	assert.Equal(t, a, b, "The two words should be the same.")
}

func TestInnerFolder(t *testing.T) {
	user := mylib.User{Email: "john_doe@us.ibm.com", Firstname: "John", Lastname: "Doe"}
	assert.Equal(t, user.Firstname, "John", "The two words should be the same.")
	fmt.Println("Finished TestInnerFolder...")
}

func TestCreateNewUser(t *testing.T) {
	// ARRANGE
	simpleAsset := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleAsset)
	fmt.Println(mockStub)
	// testStub := TestAPIStub{data: make(map[string][]byte)}
	txId := "mockTxID"
	internalId := "00001"
	firstName := "John"

	// args := []string {
	// 	internalId,
	// 	firstName,
	// 	"Doe",
	// 	"john.doek@somewhere.com",
	// }

	args := []string{
		internalId,
		firstName,
	}

	// ACT
	mockStub.MockTransactionStart(txId)
	response := simpleAsset.createUser(mockStub, args)
	mockStub.MockTransactionEnd(txId)

	// ASSERT
	if s := response.GetStatus(); s != 200 {
		t.Errorf("The response status is %d, instead of 200", s)
		t.Errorf("The message: %s", response.Message)
	}

	byteArray := mockStub.State[internalId]
	fName := string(byteArray[:])
	assert.Equal(t, firstName, fName, "The two first names should be the same.")

}
