// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package opsworksiface provides an interface to enable mocking the AWS OpsWorks service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package opsworksiface

import (
	"github.com/stowelly/aws-sdk-go/aws/request"
	"github.com/stowelly/aws-sdk-go/service/opsworks"
)

// OpsWorksAPI provides an interface to enable mocking the
// opsworks.OpsWorks service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS OpsWorks.
//    func myFunc(svc opsworksiface.OpsWorksAPI) bool {
//        // Make svc.AssignInstance request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := opsworks.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockOpsWorksClient struct {
//        opsworksiface.OpsWorksAPI
//    }
//    func (m *mockOpsWorksClient) AssignInstance(input *opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockOpsWorksClient{}
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
type OpsWorksAPI interface {
	AssignInstanceRequest(*opsworks.AssignInstanceInput) (*request.Request, *opsworks.AssignInstanceOutput)

	AssignInstance(*opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error)

	AssignVolumeRequest(*opsworks.AssignVolumeInput) (*request.Request, *opsworks.AssignVolumeOutput)

	AssignVolume(*opsworks.AssignVolumeInput) (*opsworks.AssignVolumeOutput, error)

	AssociateElasticIpRequest(*opsworks.AssociateElasticIpInput) (*request.Request, *opsworks.AssociateElasticIpOutput)

	AssociateElasticIp(*opsworks.AssociateElasticIpInput) (*opsworks.AssociateElasticIpOutput, error)

	AttachElasticLoadBalancerRequest(*opsworks.AttachElasticLoadBalancerInput) (*request.Request, *opsworks.AttachElasticLoadBalancerOutput)

	AttachElasticLoadBalancer(*opsworks.AttachElasticLoadBalancerInput) (*opsworks.AttachElasticLoadBalancerOutput, error)

	CloneStackRequest(*opsworks.CloneStackInput) (*request.Request, *opsworks.CloneStackOutput)

	CloneStack(*opsworks.CloneStackInput) (*opsworks.CloneStackOutput, error)

	CreateAppRequest(*opsworks.CreateAppInput) (*request.Request, *opsworks.CreateAppOutput)

	CreateApp(*opsworks.CreateAppInput) (*opsworks.CreateAppOutput, error)

	CreateDeploymentRequest(*opsworks.CreateDeploymentInput) (*request.Request, *opsworks.CreateDeploymentOutput)

	CreateDeployment(*opsworks.CreateDeploymentInput) (*opsworks.CreateDeploymentOutput, error)

	CreateInstanceRequest(*opsworks.CreateInstanceInput) (*request.Request, *opsworks.CreateInstanceOutput)

	CreateInstance(*opsworks.CreateInstanceInput) (*opsworks.CreateInstanceOutput, error)

	CreateLayerRequest(*opsworks.CreateLayerInput) (*request.Request, *opsworks.CreateLayerOutput)

	CreateLayer(*opsworks.CreateLayerInput) (*opsworks.CreateLayerOutput, error)

	CreateStackRequest(*opsworks.CreateStackInput) (*request.Request, *opsworks.CreateStackOutput)

	CreateStack(*opsworks.CreateStackInput) (*opsworks.CreateStackOutput, error)

	CreateUserProfileRequest(*opsworks.CreateUserProfileInput) (*request.Request, *opsworks.CreateUserProfileOutput)

	CreateUserProfile(*opsworks.CreateUserProfileInput) (*opsworks.CreateUserProfileOutput, error)

	DeleteAppRequest(*opsworks.DeleteAppInput) (*request.Request, *opsworks.DeleteAppOutput)

	DeleteApp(*opsworks.DeleteAppInput) (*opsworks.DeleteAppOutput, error)

	DeleteInstanceRequest(*opsworks.DeleteInstanceInput) (*request.Request, *opsworks.DeleteInstanceOutput)

	DeleteInstance(*opsworks.DeleteInstanceInput) (*opsworks.DeleteInstanceOutput, error)

	DeleteLayerRequest(*opsworks.DeleteLayerInput) (*request.Request, *opsworks.DeleteLayerOutput)

	DeleteLayer(*opsworks.DeleteLayerInput) (*opsworks.DeleteLayerOutput, error)

	DeleteStackRequest(*opsworks.DeleteStackInput) (*request.Request, *opsworks.DeleteStackOutput)

	DeleteStack(*opsworks.DeleteStackInput) (*opsworks.DeleteStackOutput, error)

	DeleteUserProfileRequest(*opsworks.DeleteUserProfileInput) (*request.Request, *opsworks.DeleteUserProfileOutput)

	DeleteUserProfile(*opsworks.DeleteUserProfileInput) (*opsworks.DeleteUserProfileOutput, error)

	DeregisterEcsClusterRequest(*opsworks.DeregisterEcsClusterInput) (*request.Request, *opsworks.DeregisterEcsClusterOutput)

	DeregisterEcsCluster(*opsworks.DeregisterEcsClusterInput) (*opsworks.DeregisterEcsClusterOutput, error)

	DeregisterElasticIpRequest(*opsworks.DeregisterElasticIpInput) (*request.Request, *opsworks.DeregisterElasticIpOutput)

	DeregisterElasticIp(*opsworks.DeregisterElasticIpInput) (*opsworks.DeregisterElasticIpOutput, error)

	DeregisterInstanceRequest(*opsworks.DeregisterInstanceInput) (*request.Request, *opsworks.DeregisterInstanceOutput)

	DeregisterInstance(*opsworks.DeregisterInstanceInput) (*opsworks.DeregisterInstanceOutput, error)

	DeregisterRdsDbInstanceRequest(*opsworks.DeregisterRdsDbInstanceInput) (*request.Request, *opsworks.DeregisterRdsDbInstanceOutput)

	DeregisterRdsDbInstance(*opsworks.DeregisterRdsDbInstanceInput) (*opsworks.DeregisterRdsDbInstanceOutput, error)

	DeregisterVolumeRequest(*opsworks.DeregisterVolumeInput) (*request.Request, *opsworks.DeregisterVolumeOutput)

	DeregisterVolume(*opsworks.DeregisterVolumeInput) (*opsworks.DeregisterVolumeOutput, error)

	DescribeAgentVersionsRequest(*opsworks.DescribeAgentVersionsInput) (*request.Request, *opsworks.DescribeAgentVersionsOutput)

	DescribeAgentVersions(*opsworks.DescribeAgentVersionsInput) (*opsworks.DescribeAgentVersionsOutput, error)

	DescribeAppsRequest(*opsworks.DescribeAppsInput) (*request.Request, *opsworks.DescribeAppsOutput)

	DescribeApps(*opsworks.DescribeAppsInput) (*opsworks.DescribeAppsOutput, error)

	DescribeCommandsRequest(*opsworks.DescribeCommandsInput) (*request.Request, *opsworks.DescribeCommandsOutput)

	DescribeCommands(*opsworks.DescribeCommandsInput) (*opsworks.DescribeCommandsOutput, error)

	DescribeDeploymentsRequest(*opsworks.DescribeDeploymentsInput) (*request.Request, *opsworks.DescribeDeploymentsOutput)

	DescribeDeployments(*opsworks.DescribeDeploymentsInput) (*opsworks.DescribeDeploymentsOutput, error)

	DescribeEcsClustersRequest(*opsworks.DescribeEcsClustersInput) (*request.Request, *opsworks.DescribeEcsClustersOutput)

	DescribeEcsClusters(*opsworks.DescribeEcsClustersInput) (*opsworks.DescribeEcsClustersOutput, error)

	DescribeEcsClustersPages(*opsworks.DescribeEcsClustersInput, func(*opsworks.DescribeEcsClustersOutput, bool) bool) error

	DescribeElasticIpsRequest(*opsworks.DescribeElasticIpsInput) (*request.Request, *opsworks.DescribeElasticIpsOutput)

	DescribeElasticIps(*opsworks.DescribeElasticIpsInput) (*opsworks.DescribeElasticIpsOutput, error)

	DescribeElasticLoadBalancersRequest(*opsworks.DescribeElasticLoadBalancersInput) (*request.Request, *opsworks.DescribeElasticLoadBalancersOutput)

	DescribeElasticLoadBalancers(*opsworks.DescribeElasticLoadBalancersInput) (*opsworks.DescribeElasticLoadBalancersOutput, error)

	DescribeInstancesRequest(*opsworks.DescribeInstancesInput) (*request.Request, *opsworks.DescribeInstancesOutput)

	DescribeInstances(*opsworks.DescribeInstancesInput) (*opsworks.DescribeInstancesOutput, error)

	DescribeLayersRequest(*opsworks.DescribeLayersInput) (*request.Request, *opsworks.DescribeLayersOutput)

	DescribeLayers(*opsworks.DescribeLayersInput) (*opsworks.DescribeLayersOutput, error)

	DescribeLoadBasedAutoScalingRequest(*opsworks.DescribeLoadBasedAutoScalingInput) (*request.Request, *opsworks.DescribeLoadBasedAutoScalingOutput)

	DescribeLoadBasedAutoScaling(*opsworks.DescribeLoadBasedAutoScalingInput) (*opsworks.DescribeLoadBasedAutoScalingOutput, error)

	DescribeMyUserProfileRequest(*opsworks.DescribeMyUserProfileInput) (*request.Request, *opsworks.DescribeMyUserProfileOutput)

	DescribeMyUserProfile(*opsworks.DescribeMyUserProfileInput) (*opsworks.DescribeMyUserProfileOutput, error)

	DescribePermissionsRequest(*opsworks.DescribePermissionsInput) (*request.Request, *opsworks.DescribePermissionsOutput)

	DescribePermissions(*opsworks.DescribePermissionsInput) (*opsworks.DescribePermissionsOutput, error)

	DescribeRaidArraysRequest(*opsworks.DescribeRaidArraysInput) (*request.Request, *opsworks.DescribeRaidArraysOutput)

	DescribeRaidArrays(*opsworks.DescribeRaidArraysInput) (*opsworks.DescribeRaidArraysOutput, error)

	DescribeRdsDbInstancesRequest(*opsworks.DescribeRdsDbInstancesInput) (*request.Request, *opsworks.DescribeRdsDbInstancesOutput)

	DescribeRdsDbInstances(*opsworks.DescribeRdsDbInstancesInput) (*opsworks.DescribeRdsDbInstancesOutput, error)

	DescribeServiceErrorsRequest(*opsworks.DescribeServiceErrorsInput) (*request.Request, *opsworks.DescribeServiceErrorsOutput)

	DescribeServiceErrors(*opsworks.DescribeServiceErrorsInput) (*opsworks.DescribeServiceErrorsOutput, error)

	DescribeStackProvisioningParametersRequest(*opsworks.DescribeStackProvisioningParametersInput) (*request.Request, *opsworks.DescribeStackProvisioningParametersOutput)

	DescribeStackProvisioningParameters(*opsworks.DescribeStackProvisioningParametersInput) (*opsworks.DescribeStackProvisioningParametersOutput, error)

	DescribeStackSummaryRequest(*opsworks.DescribeStackSummaryInput) (*request.Request, *opsworks.DescribeStackSummaryOutput)

	DescribeStackSummary(*opsworks.DescribeStackSummaryInput) (*opsworks.DescribeStackSummaryOutput, error)

	DescribeStacksRequest(*opsworks.DescribeStacksInput) (*request.Request, *opsworks.DescribeStacksOutput)

	DescribeStacks(*opsworks.DescribeStacksInput) (*opsworks.DescribeStacksOutput, error)

	DescribeTimeBasedAutoScalingRequest(*opsworks.DescribeTimeBasedAutoScalingInput) (*request.Request, *opsworks.DescribeTimeBasedAutoScalingOutput)

	DescribeTimeBasedAutoScaling(*opsworks.DescribeTimeBasedAutoScalingInput) (*opsworks.DescribeTimeBasedAutoScalingOutput, error)

	DescribeUserProfilesRequest(*opsworks.DescribeUserProfilesInput) (*request.Request, *opsworks.DescribeUserProfilesOutput)

	DescribeUserProfiles(*opsworks.DescribeUserProfilesInput) (*opsworks.DescribeUserProfilesOutput, error)

	DescribeVolumesRequest(*opsworks.DescribeVolumesInput) (*request.Request, *opsworks.DescribeVolumesOutput)

	DescribeVolumes(*opsworks.DescribeVolumesInput) (*opsworks.DescribeVolumesOutput, error)

	DetachElasticLoadBalancerRequest(*opsworks.DetachElasticLoadBalancerInput) (*request.Request, *opsworks.DetachElasticLoadBalancerOutput)

	DetachElasticLoadBalancer(*opsworks.DetachElasticLoadBalancerInput) (*opsworks.DetachElasticLoadBalancerOutput, error)

	DisassociateElasticIpRequest(*opsworks.DisassociateElasticIpInput) (*request.Request, *opsworks.DisassociateElasticIpOutput)

	DisassociateElasticIp(*opsworks.DisassociateElasticIpInput) (*opsworks.DisassociateElasticIpOutput, error)

	GetHostnameSuggestionRequest(*opsworks.GetHostnameSuggestionInput) (*request.Request, *opsworks.GetHostnameSuggestionOutput)

	GetHostnameSuggestion(*opsworks.GetHostnameSuggestionInput) (*opsworks.GetHostnameSuggestionOutput, error)

	GrantAccessRequest(*opsworks.GrantAccessInput) (*request.Request, *opsworks.GrantAccessOutput)

	GrantAccess(*opsworks.GrantAccessInput) (*opsworks.GrantAccessOutput, error)

	RebootInstanceRequest(*opsworks.RebootInstanceInput) (*request.Request, *opsworks.RebootInstanceOutput)

	RebootInstance(*opsworks.RebootInstanceInput) (*opsworks.RebootInstanceOutput, error)

	RegisterEcsClusterRequest(*opsworks.RegisterEcsClusterInput) (*request.Request, *opsworks.RegisterEcsClusterOutput)

	RegisterEcsCluster(*opsworks.RegisterEcsClusterInput) (*opsworks.RegisterEcsClusterOutput, error)

	RegisterElasticIpRequest(*opsworks.RegisterElasticIpInput) (*request.Request, *opsworks.RegisterElasticIpOutput)

	RegisterElasticIp(*opsworks.RegisterElasticIpInput) (*opsworks.RegisterElasticIpOutput, error)

	RegisterInstanceRequest(*opsworks.RegisterInstanceInput) (*request.Request, *opsworks.RegisterInstanceOutput)

	RegisterInstance(*opsworks.RegisterInstanceInput) (*opsworks.RegisterInstanceOutput, error)

	RegisterRdsDbInstanceRequest(*opsworks.RegisterRdsDbInstanceInput) (*request.Request, *opsworks.RegisterRdsDbInstanceOutput)

	RegisterRdsDbInstance(*opsworks.RegisterRdsDbInstanceInput) (*opsworks.RegisterRdsDbInstanceOutput, error)

	RegisterVolumeRequest(*opsworks.RegisterVolumeInput) (*request.Request, *opsworks.RegisterVolumeOutput)

	RegisterVolume(*opsworks.RegisterVolumeInput) (*opsworks.RegisterVolumeOutput, error)

	SetLoadBasedAutoScalingRequest(*opsworks.SetLoadBasedAutoScalingInput) (*request.Request, *opsworks.SetLoadBasedAutoScalingOutput)

	SetLoadBasedAutoScaling(*opsworks.SetLoadBasedAutoScalingInput) (*opsworks.SetLoadBasedAutoScalingOutput, error)

	SetPermissionRequest(*opsworks.SetPermissionInput) (*request.Request, *opsworks.SetPermissionOutput)

	SetPermission(*opsworks.SetPermissionInput) (*opsworks.SetPermissionOutput, error)

	SetTimeBasedAutoScalingRequest(*opsworks.SetTimeBasedAutoScalingInput) (*request.Request, *opsworks.SetTimeBasedAutoScalingOutput)

	SetTimeBasedAutoScaling(*opsworks.SetTimeBasedAutoScalingInput) (*opsworks.SetTimeBasedAutoScalingOutput, error)

	StartInstanceRequest(*opsworks.StartInstanceInput) (*request.Request, *opsworks.StartInstanceOutput)

	StartInstance(*opsworks.StartInstanceInput) (*opsworks.StartInstanceOutput, error)

	StartStackRequest(*opsworks.StartStackInput) (*request.Request, *opsworks.StartStackOutput)

	StartStack(*opsworks.StartStackInput) (*opsworks.StartStackOutput, error)

	StopInstanceRequest(*opsworks.StopInstanceInput) (*request.Request, *opsworks.StopInstanceOutput)

	StopInstance(*opsworks.StopInstanceInput) (*opsworks.StopInstanceOutput, error)

	StopStackRequest(*opsworks.StopStackInput) (*request.Request, *opsworks.StopStackOutput)

	StopStack(*opsworks.StopStackInput) (*opsworks.StopStackOutput, error)

	UnassignInstanceRequest(*opsworks.UnassignInstanceInput) (*request.Request, *opsworks.UnassignInstanceOutput)

	UnassignInstance(*opsworks.UnassignInstanceInput) (*opsworks.UnassignInstanceOutput, error)

	UnassignVolumeRequest(*opsworks.UnassignVolumeInput) (*request.Request, *opsworks.UnassignVolumeOutput)

	UnassignVolume(*opsworks.UnassignVolumeInput) (*opsworks.UnassignVolumeOutput, error)

	UpdateAppRequest(*opsworks.UpdateAppInput) (*request.Request, *opsworks.UpdateAppOutput)

	UpdateApp(*opsworks.UpdateAppInput) (*opsworks.UpdateAppOutput, error)

	UpdateElasticIpRequest(*opsworks.UpdateElasticIpInput) (*request.Request, *opsworks.UpdateElasticIpOutput)

	UpdateElasticIp(*opsworks.UpdateElasticIpInput) (*opsworks.UpdateElasticIpOutput, error)

	UpdateInstanceRequest(*opsworks.UpdateInstanceInput) (*request.Request, *opsworks.UpdateInstanceOutput)

	UpdateInstance(*opsworks.UpdateInstanceInput) (*opsworks.UpdateInstanceOutput, error)

	UpdateLayerRequest(*opsworks.UpdateLayerInput) (*request.Request, *opsworks.UpdateLayerOutput)

	UpdateLayer(*opsworks.UpdateLayerInput) (*opsworks.UpdateLayerOutput, error)

	UpdateMyUserProfileRequest(*opsworks.UpdateMyUserProfileInput) (*request.Request, *opsworks.UpdateMyUserProfileOutput)

	UpdateMyUserProfile(*opsworks.UpdateMyUserProfileInput) (*opsworks.UpdateMyUserProfileOutput, error)

	UpdateRdsDbInstanceRequest(*opsworks.UpdateRdsDbInstanceInput) (*request.Request, *opsworks.UpdateRdsDbInstanceOutput)

	UpdateRdsDbInstance(*opsworks.UpdateRdsDbInstanceInput) (*opsworks.UpdateRdsDbInstanceOutput, error)

	UpdateStackRequest(*opsworks.UpdateStackInput) (*request.Request, *opsworks.UpdateStackOutput)

	UpdateStack(*opsworks.UpdateStackInput) (*opsworks.UpdateStackOutput, error)

	UpdateUserProfileRequest(*opsworks.UpdateUserProfileInput) (*request.Request, *opsworks.UpdateUserProfileOutput)

	UpdateUserProfile(*opsworks.UpdateUserProfileInput) (*opsworks.UpdateUserProfileOutput, error)

	UpdateVolumeRequest(*opsworks.UpdateVolumeInput) (*request.Request, *opsworks.UpdateVolumeOutput)

	UpdateVolume(*opsworks.UpdateVolumeInput) (*opsworks.UpdateVolumeOutput, error)

	WaitUntilAppExists(*opsworks.DescribeAppsInput) error

	WaitUntilDeploymentSuccessful(*opsworks.DescribeDeploymentsInput) error

	WaitUntilInstanceOnline(*opsworks.DescribeInstancesInput) error

	WaitUntilInstanceRegistered(*opsworks.DescribeInstancesInput) error

	WaitUntilInstanceStopped(*opsworks.DescribeInstancesInput) error

	WaitUntilInstanceTerminated(*opsworks.DescribeInstancesInput) error
}

var _ OpsWorksAPI = (*opsworks.OpsWorks)(nil)
