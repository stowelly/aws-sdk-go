// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elasticache_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/stowelly/aws-sdk-go/aws"
	"github.com/stowelly/aws-sdk-go/aws/session"
	"github.com/stowelly/aws-sdk-go/service/elasticache"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleElastiCache_AddTagsToResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.AddTagsToResourceInput{
		ResourceName: aws.String("String"), // Required
		Tags: []*elasticache.Tag{ // Required
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.AddTagsToResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_AuthorizeCacheSecurityGroupIngress() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.AuthorizeCacheSecurityGroupIngressInput{
		CacheSecurityGroupName:  aws.String("String"), // Required
		EC2SecurityGroupName:    aws.String("String"), // Required
		EC2SecurityGroupOwnerId: aws.String("String"), // Required
	}
	resp, err := svc.AuthorizeCacheSecurityGroupIngress(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_CopySnapshot() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.CopySnapshotInput{
		SourceSnapshotName: aws.String("String"), // Required
		TargetSnapshotName: aws.String("String"), // Required
		TargetBucket:       aws.String("String"),
	}
	resp, err := svc.CopySnapshot(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_CreateCacheCluster() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.CreateCacheClusterInput{
		CacheClusterId:          aws.String("String"), // Required
		AZMode:                  aws.String("AZMode"),
		AuthToken:               aws.String("String"),
		AutoMinorVersionUpgrade: aws.Bool(true),
		CacheNodeType:           aws.String("String"),
		CacheParameterGroupName: aws.String("String"),
		CacheSecurityGroupNames: []*string{
			aws.String("String"), // Required
			// More values...
		},
		CacheSubnetGroupName: aws.String("String"),
		Engine:               aws.String("String"),
		EngineVersion:        aws.String("String"),
		NotificationTopicArn: aws.String("String"),
		NumCacheNodes:        aws.Int64(1),
		Port:                 aws.Int64(1),
		PreferredAvailabilityZone: aws.String("String"),
		PreferredAvailabilityZones: []*string{
			aws.String("String"), // Required
			// More values...
		},
		PreferredMaintenanceWindow: aws.String("String"),
		ReplicationGroupId:         aws.String("String"),
		SecurityGroupIds: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SnapshotArns: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SnapshotName:           aws.String("String"),
		SnapshotRetentionLimit: aws.Int64(1),
		SnapshotWindow:         aws.String("String"),
		Tags: []*elasticache.Tag{
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateCacheCluster(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_CreateCacheParameterGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.CreateCacheParameterGroupInput{
		CacheParameterGroupFamily: aws.String("String"), // Required
		CacheParameterGroupName:   aws.String("String"), // Required
		Description:               aws.String("String"), // Required
	}
	resp, err := svc.CreateCacheParameterGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_CreateCacheSecurityGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.CreateCacheSecurityGroupInput{
		CacheSecurityGroupName: aws.String("String"), // Required
		Description:            aws.String("String"), // Required
	}
	resp, err := svc.CreateCacheSecurityGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_CreateCacheSubnetGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.CreateCacheSubnetGroupInput{
		CacheSubnetGroupDescription: aws.String("String"), // Required
		CacheSubnetGroupName:        aws.String("String"), // Required
		SubnetIds: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateCacheSubnetGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_CreateReplicationGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.CreateReplicationGroupInput{
		ReplicationGroupDescription: aws.String("String"), // Required
		ReplicationGroupId:          aws.String("String"), // Required
		AuthToken:                   aws.String("String"),
		AutoMinorVersionUpgrade:     aws.Bool(true),
		AutomaticFailoverEnabled:    aws.Bool(true),
		CacheNodeType:               aws.String("String"),
		CacheParameterGroupName:     aws.String("String"),
		CacheSecurityGroupNames: []*string{
			aws.String("String"), // Required
			// More values...
		},
		CacheSubnetGroupName: aws.String("String"),
		Engine:               aws.String("String"),
		EngineVersion:        aws.String("String"),
		NodeGroupConfiguration: []*elasticache.NodeGroupConfiguration{
			{ // Required
				PrimaryAvailabilityZone: aws.String("String"),
				ReplicaAvailabilityZones: []*string{
					aws.String("String"), // Required
					// More values...
				},
				ReplicaCount: aws.Int64(1),
				Slots:        aws.String("String"),
			},
			// More values...
		},
		NotificationTopicArn: aws.String("String"),
		NumCacheClusters:     aws.Int64(1),
		NumNodeGroups:        aws.Int64(1),
		Port:                 aws.Int64(1),
		PreferredCacheClusterAZs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		PreferredMaintenanceWindow: aws.String("String"),
		PrimaryClusterId:           aws.String("String"),
		ReplicasPerNodeGroup:       aws.Int64(1),
		SecurityGroupIds: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SnapshotArns: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SnapshotName:           aws.String("String"),
		SnapshotRetentionLimit: aws.Int64(1),
		SnapshotWindow:         aws.String("String"),
		Tags: []*elasticache.Tag{
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateReplicationGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_CreateSnapshot() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.CreateSnapshotInput{
		SnapshotName:       aws.String("String"), // Required
		CacheClusterId:     aws.String("String"),
		ReplicationGroupId: aws.String("String"),
	}
	resp, err := svc.CreateSnapshot(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DeleteCacheCluster() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DeleteCacheClusterInput{
		CacheClusterId:          aws.String("String"), // Required
		FinalSnapshotIdentifier: aws.String("String"),
	}
	resp, err := svc.DeleteCacheCluster(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DeleteCacheParameterGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DeleteCacheParameterGroupInput{
		CacheParameterGroupName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteCacheParameterGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DeleteCacheSecurityGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DeleteCacheSecurityGroupInput{
		CacheSecurityGroupName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteCacheSecurityGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DeleteCacheSubnetGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DeleteCacheSubnetGroupInput{
		CacheSubnetGroupName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteCacheSubnetGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DeleteReplicationGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DeleteReplicationGroupInput{
		ReplicationGroupId:      aws.String("String"), // Required
		FinalSnapshotIdentifier: aws.String("String"),
		RetainPrimaryCluster:    aws.Bool(true),
	}
	resp, err := svc.DeleteReplicationGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DeleteSnapshot() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DeleteSnapshotInput{
		SnapshotName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteSnapshot(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DescribeCacheClusters() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DescribeCacheClustersInput{
		CacheClusterId:    aws.String("String"),
		Marker:            aws.String("String"),
		MaxRecords:        aws.Int64(1),
		ShowCacheNodeInfo: aws.Bool(true),
	}
	resp, err := svc.DescribeCacheClusters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DescribeCacheEngineVersions() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DescribeCacheEngineVersionsInput{
		CacheParameterGroupFamily: aws.String("String"),
		DefaultOnly:               aws.Bool(true),
		Engine:                    aws.String("String"),
		EngineVersion:             aws.String("String"),
		Marker:                    aws.String("String"),
		MaxRecords:                aws.Int64(1),
	}
	resp, err := svc.DescribeCacheEngineVersions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DescribeCacheParameterGroups() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DescribeCacheParameterGroupsInput{
		CacheParameterGroupName: aws.String("String"),
		Marker:                  aws.String("String"),
		MaxRecords:              aws.Int64(1),
	}
	resp, err := svc.DescribeCacheParameterGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DescribeCacheParameters() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DescribeCacheParametersInput{
		CacheParameterGroupName: aws.String("String"), // Required
		Marker:                  aws.String("String"),
		MaxRecords:              aws.Int64(1),
		Source:                  aws.String("String"),
	}
	resp, err := svc.DescribeCacheParameters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DescribeCacheSecurityGroups() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DescribeCacheSecurityGroupsInput{
		CacheSecurityGroupName: aws.String("String"),
		Marker:                 aws.String("String"),
		MaxRecords:             aws.Int64(1),
	}
	resp, err := svc.DescribeCacheSecurityGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DescribeCacheSubnetGroups() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DescribeCacheSubnetGroupsInput{
		CacheSubnetGroupName: aws.String("String"),
		Marker:               aws.String("String"),
		MaxRecords:           aws.Int64(1),
	}
	resp, err := svc.DescribeCacheSubnetGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DescribeEngineDefaultParameters() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DescribeEngineDefaultParametersInput{
		CacheParameterGroupFamily: aws.String("String"), // Required
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeEngineDefaultParameters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DescribeEvents() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DescribeEventsInput{
		Duration:         aws.Int64(1),
		EndTime:          aws.Time(time.Now()),
		Marker:           aws.String("String"),
		MaxRecords:       aws.Int64(1),
		SourceIdentifier: aws.String("String"),
		SourceType:       aws.String("SourceType"),
		StartTime:        aws.Time(time.Now()),
	}
	resp, err := svc.DescribeEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DescribeReplicationGroups() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DescribeReplicationGroupsInput{
		Marker:             aws.String("String"),
		MaxRecords:         aws.Int64(1),
		ReplicationGroupId: aws.String("String"),
	}
	resp, err := svc.DescribeReplicationGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DescribeReservedCacheNodes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DescribeReservedCacheNodesInput{
		CacheNodeType:                aws.String("String"),
		Duration:                     aws.String("String"),
		Marker:                       aws.String("String"),
		MaxRecords:                   aws.Int64(1),
		OfferingType:                 aws.String("String"),
		ProductDescription:           aws.String("String"),
		ReservedCacheNodeId:          aws.String("String"),
		ReservedCacheNodesOfferingId: aws.String("String"),
	}
	resp, err := svc.DescribeReservedCacheNodes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DescribeReservedCacheNodesOfferings() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DescribeReservedCacheNodesOfferingsInput{
		CacheNodeType:                aws.String("String"),
		Duration:                     aws.String("String"),
		Marker:                       aws.String("String"),
		MaxRecords:                   aws.Int64(1),
		OfferingType:                 aws.String("String"),
		ProductDescription:           aws.String("String"),
		ReservedCacheNodesOfferingId: aws.String("String"),
	}
	resp, err := svc.DescribeReservedCacheNodesOfferings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_DescribeSnapshots() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.DescribeSnapshotsInput{
		CacheClusterId:      aws.String("String"),
		Marker:              aws.String("String"),
		MaxRecords:          aws.Int64(1),
		ReplicationGroupId:  aws.String("String"),
		ShowNodeGroupConfig: aws.Bool(true),
		SnapshotName:        aws.String("String"),
		SnapshotSource:      aws.String("String"),
	}
	resp, err := svc.DescribeSnapshots(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_ListAllowedNodeTypeModifications() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.ListAllowedNodeTypeModificationsInput{
		CacheClusterId:     aws.String("String"),
		ReplicationGroupId: aws.String("String"),
	}
	resp, err := svc.ListAllowedNodeTypeModifications(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_ListTagsForResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.ListTagsForResourceInput{
		ResourceName: aws.String("String"), // Required
	}
	resp, err := svc.ListTagsForResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_ModifyCacheCluster() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.ModifyCacheClusterInput{
		CacheClusterId:          aws.String("String"), // Required
		AZMode:                  aws.String("AZMode"),
		ApplyImmediately:        aws.Bool(true),
		AutoMinorVersionUpgrade: aws.Bool(true),
		CacheNodeIdsToRemove: []*string{
			aws.String("String"), // Required
			// More values...
		},
		CacheNodeType:           aws.String("String"),
		CacheParameterGroupName: aws.String("String"),
		CacheSecurityGroupNames: []*string{
			aws.String("String"), // Required
			// More values...
		},
		EngineVersion: aws.String("String"),
		NewAvailabilityZones: []*string{
			aws.String("String"), // Required
			// More values...
		},
		NotificationTopicArn:       aws.String("String"),
		NotificationTopicStatus:    aws.String("String"),
		NumCacheNodes:              aws.Int64(1),
		PreferredMaintenanceWindow: aws.String("String"),
		SecurityGroupIds: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SnapshotRetentionLimit: aws.Int64(1),
		SnapshotWindow:         aws.String("String"),
	}
	resp, err := svc.ModifyCacheCluster(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_ModifyCacheParameterGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.ModifyCacheParameterGroupInput{
		CacheParameterGroupName: aws.String("String"), // Required
		ParameterNameValues: []*elasticache.ParameterNameValue{ // Required
			{ // Required
				ParameterName:  aws.String("String"),
				ParameterValue: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.ModifyCacheParameterGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_ModifyCacheSubnetGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.ModifyCacheSubnetGroupInput{
		CacheSubnetGroupName:        aws.String("String"), // Required
		CacheSubnetGroupDescription: aws.String("String"),
		SubnetIds: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.ModifyCacheSubnetGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_ModifyReplicationGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.ModifyReplicationGroupInput{
		ReplicationGroupId:       aws.String("String"), // Required
		ApplyImmediately:         aws.Bool(true),
		AutoMinorVersionUpgrade:  aws.Bool(true),
		AutomaticFailoverEnabled: aws.Bool(true),
		CacheNodeType:            aws.String("String"),
		CacheParameterGroupName:  aws.String("String"),
		CacheSecurityGroupNames: []*string{
			aws.String("String"), // Required
			// More values...
		},
		EngineVersion:               aws.String("String"),
		NotificationTopicArn:        aws.String("String"),
		NotificationTopicStatus:     aws.String("String"),
		PreferredMaintenanceWindow:  aws.String("String"),
		PrimaryClusterId:            aws.String("String"),
		ReplicationGroupDescription: aws.String("String"),
		SecurityGroupIds: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SnapshotRetentionLimit: aws.Int64(1),
		SnapshotWindow:         aws.String("String"),
		SnapshottingClusterId:  aws.String("String"),
	}
	resp, err := svc.ModifyReplicationGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_PurchaseReservedCacheNodesOffering() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.PurchaseReservedCacheNodesOfferingInput{
		ReservedCacheNodesOfferingId: aws.String("String"), // Required
		CacheNodeCount:               aws.Int64(1),
		ReservedCacheNodeId:          aws.String("String"),
	}
	resp, err := svc.PurchaseReservedCacheNodesOffering(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_RebootCacheCluster() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.RebootCacheClusterInput{
		CacheClusterId: aws.String("String"), // Required
		CacheNodeIdsToReboot: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.RebootCacheCluster(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_RemoveTagsFromResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.RemoveTagsFromResourceInput{
		ResourceName: aws.String("String"), // Required
		TagKeys: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTagsFromResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_ResetCacheParameterGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.ResetCacheParameterGroupInput{
		CacheParameterGroupName: aws.String("String"), // Required
		ParameterNameValues: []*elasticache.ParameterNameValue{
			{ // Required
				ParameterName:  aws.String("String"),
				ParameterValue: aws.String("String"),
			},
			// More values...
		},
		ResetAllParameters: aws.Bool(true),
	}
	resp, err := svc.ResetCacheParameterGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElastiCache_RevokeCacheSecurityGroupIngress() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elasticache.New(sess)

	params := &elasticache.RevokeCacheSecurityGroupIngressInput{
		CacheSecurityGroupName:  aws.String("String"), // Required
		EC2SecurityGroupName:    aws.String("String"), // Required
		EC2SecurityGroupOwnerId: aws.String("String"), // Required
	}
	resp, err := svc.RevokeCacheSecurityGroupIngress(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
