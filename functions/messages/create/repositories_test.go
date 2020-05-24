package main

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRepositories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repositories Suite")
}

var _ = Describe("Save Message", func() {
	var message 	Message
	var dbSession 	stubDynamoDBSuccess
	var err 		error

	BeforeEach(func() {
		dbSession = stubDynamoDBSuccess{}
		message = buildMessage()
		_, err = SaveMessage(&dbSession, message)
	})

	It("returns no errors", func() {
		Expect(err).To(BeNil())
	})
})

const Boom = "something went wrong"

type stubDynamoDBSuccess struct {
	dynamodbiface.DynamoDBAPI
}

type stubDynamoDBFailure struct {
	dynamodbiface.DynamoDBAPI
}

func (m *stubDynamoDBSuccess) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{}, nil
}

func (m *stubDynamoDBFailure) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{}, errors.New(Boom)
}

//func TestSaveMessageSuccess(t *testing.T) {
//	dbSession	:= &stubDynamoDBSuccess{}
//	message 	:= buildMessage()
//	_, err 		:= SaveMessage(dbSession, message)
//
//	if err != nil {
//		t.Error("There should be no error")
//	}
//}
//
//func TestSaveMessageFailure(t *testing.T) {
//	dbSession	:= &stubDynamoDBFailure{}
//	message 	:= buildMessage()
//	_, err 		:= SaveMessage(dbSession, message)
//
//	if err.Error() != Boom {
//		t.Errorf("Reported error should be %s", Boom)
//	}
//}
