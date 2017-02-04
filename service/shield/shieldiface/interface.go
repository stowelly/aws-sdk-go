// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package shieldiface provides an interface to enable mocking the AWS Shield service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package shieldiface

import (
	"github.com/stowelly/aws-sdk-go/aws/request"
	"github.com/stowelly/aws-sdk-go/service/shield"
)

// ShieldAPI provides an interface to enable mocking the
// shield.Shield service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Shield.
//    func myFunc(svc shieldiface.ShieldAPI) bool {
//        // Make svc.CreateProtection request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := shield.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockShieldClient struct {
//        shieldiface.ShieldAPI
//    }
//    func (m *mockShieldClient) CreateProtection(input *shield.CreateProtectionInput) (*shield.CreateProtectionOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockShieldClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ShieldAPI interface {
	CreateProtectionRequest(*shield.CreateProtectionInput) (*request.Request, *shield.CreateProtectionOutput)

	CreateProtection(*shield.CreateProtectionInput) (*shield.CreateProtectionOutput, error)

	CreateSubscriptionRequest(*shield.CreateSubscriptionInput) (*request.Request, *shield.CreateSubscriptionOutput)

	CreateSubscription(*shield.CreateSubscriptionInput) (*shield.CreateSubscriptionOutput, error)

	DeleteProtectionRequest(*shield.DeleteProtectionInput) (*request.Request, *shield.DeleteProtectionOutput)

	DeleteProtection(*shield.DeleteProtectionInput) (*shield.DeleteProtectionOutput, error)

	DeleteSubscriptionRequest(*shield.DeleteSubscriptionInput) (*request.Request, *shield.DeleteSubscriptionOutput)

	DeleteSubscription(*shield.DeleteSubscriptionInput) (*shield.DeleteSubscriptionOutput, error)

	DescribeAttackRequest(*shield.DescribeAttackInput) (*request.Request, *shield.DescribeAttackOutput)

	DescribeAttack(*shield.DescribeAttackInput) (*shield.DescribeAttackOutput, error)

	DescribeProtectionRequest(*shield.DescribeProtectionInput) (*request.Request, *shield.DescribeProtectionOutput)

	DescribeProtection(*shield.DescribeProtectionInput) (*shield.DescribeProtectionOutput, error)

	DescribeSubscriptionRequest(*shield.DescribeSubscriptionInput) (*request.Request, *shield.DescribeSubscriptionOutput)

	DescribeSubscription(*shield.DescribeSubscriptionInput) (*shield.DescribeSubscriptionOutput, error)

	ListAttacksRequest(*shield.ListAttacksInput) (*request.Request, *shield.ListAttacksOutput)

	ListAttacks(*shield.ListAttacksInput) (*shield.ListAttacksOutput, error)

	ListProtectionsRequest(*shield.ListProtectionsInput) (*request.Request, *shield.ListProtectionsOutput)

	ListProtections(*shield.ListProtectionsInput) (*shield.ListProtectionsOutput, error)
}

var _ ShieldAPI = (*shield.Shield)(nil)
