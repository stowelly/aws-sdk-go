// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package efsiface provides an interface to enable mocking the Amazon Elastic File System service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package efsiface

import (
	"github.com/stowelly/aws-sdk-go/aws/request"
	"github.com/stowelly/aws-sdk-go/service/efs"
)

// EFSAPI provides an interface to enable mocking the
// efs.EFS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Elastic File System.
//    func myFunc(svc efsiface.EFSAPI) bool {
//        // Make svc.CreateFileSystem request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := efs.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockEFSClient struct {
//        efsiface.EFSAPI
//    }
//    func (m *mockEFSClient) CreateFileSystem(input *efs.CreateFileSystemInput) (*efs.FileSystemDescription, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockEFSClient{}
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
type EFSAPI interface {
	CreateFileSystemRequest(*efs.CreateFileSystemInput) (*request.Request, *efs.FileSystemDescription)

	CreateFileSystem(*efs.CreateFileSystemInput) (*efs.FileSystemDescription, error)

	CreateMountTargetRequest(*efs.CreateMountTargetInput) (*request.Request, *efs.MountTargetDescription)

	CreateMountTarget(*efs.CreateMountTargetInput) (*efs.MountTargetDescription, error)

	CreateTagsRequest(*efs.CreateTagsInput) (*request.Request, *efs.CreateTagsOutput)

	CreateTags(*efs.CreateTagsInput) (*efs.CreateTagsOutput, error)

	DeleteFileSystemRequest(*efs.DeleteFileSystemInput) (*request.Request, *efs.DeleteFileSystemOutput)

	DeleteFileSystem(*efs.DeleteFileSystemInput) (*efs.DeleteFileSystemOutput, error)

	DeleteMountTargetRequest(*efs.DeleteMountTargetInput) (*request.Request, *efs.DeleteMountTargetOutput)

	DeleteMountTarget(*efs.DeleteMountTargetInput) (*efs.DeleteMountTargetOutput, error)

	DeleteTagsRequest(*efs.DeleteTagsInput) (*request.Request, *efs.DeleteTagsOutput)

	DeleteTags(*efs.DeleteTagsInput) (*efs.DeleteTagsOutput, error)

	DescribeFileSystemsRequest(*efs.DescribeFileSystemsInput) (*request.Request, *efs.DescribeFileSystemsOutput)

	DescribeFileSystems(*efs.DescribeFileSystemsInput) (*efs.DescribeFileSystemsOutput, error)

	DescribeMountTargetSecurityGroupsRequest(*efs.DescribeMountTargetSecurityGroupsInput) (*request.Request, *efs.DescribeMountTargetSecurityGroupsOutput)

	DescribeMountTargetSecurityGroups(*efs.DescribeMountTargetSecurityGroupsInput) (*efs.DescribeMountTargetSecurityGroupsOutput, error)

	DescribeMountTargetsRequest(*efs.DescribeMountTargetsInput) (*request.Request, *efs.DescribeMountTargetsOutput)

	DescribeMountTargets(*efs.DescribeMountTargetsInput) (*efs.DescribeMountTargetsOutput, error)

	DescribeTagsRequest(*efs.DescribeTagsInput) (*request.Request, *efs.DescribeTagsOutput)

	DescribeTags(*efs.DescribeTagsInput) (*efs.DescribeTagsOutput, error)

	ModifyMountTargetSecurityGroupsRequest(*efs.ModifyMountTargetSecurityGroupsInput) (*request.Request, *efs.ModifyMountTargetSecurityGroupsOutput)

	ModifyMountTargetSecurityGroups(*efs.ModifyMountTargetSecurityGroupsInput) (*efs.ModifyMountTargetSecurityGroupsOutput, error)
}

var _ EFSAPI = (*efs.EFS)(nil)
