// +build integration

package performance

import (
	"reflect"

	"github.com/stowelly/aws-sdk-go/service/acm"
	"github.com/stowelly/aws-sdk-go/service/apigateway"
	"github.com/stowelly/aws-sdk-go/service/autoscaling"
	"github.com/stowelly/aws-sdk-go/service/cloudformation"
	"github.com/stowelly/aws-sdk-go/service/cloudfront"
	"github.com/stowelly/aws-sdk-go/service/cloudhsm"
	"github.com/stowelly/aws-sdk-go/service/cloudsearch"
	"github.com/stowelly/aws-sdk-go/service/cloudsearchdomain"
	"github.com/stowelly/aws-sdk-go/service/cloudtrail"
	"github.com/stowelly/aws-sdk-go/service/cloudwatch"
	"github.com/stowelly/aws-sdk-go/service/cloudwatchevents"
	"github.com/stowelly/aws-sdk-go/service/cloudwatchlogs"
	"github.com/stowelly/aws-sdk-go/service/codecommit"
	"github.com/stowelly/aws-sdk-go/service/codedeploy"
	"github.com/stowelly/aws-sdk-go/service/codepipeline"
	"github.com/stowelly/aws-sdk-go/service/cognitoidentity"
	"github.com/stowelly/aws-sdk-go/service/cognitosync"
	"github.com/stowelly/aws-sdk-go/service/configservice"
	"github.com/stowelly/aws-sdk-go/service/datapipeline"
	"github.com/stowelly/aws-sdk-go/service/devicefarm"
	"github.com/stowelly/aws-sdk-go/service/directconnect"
	"github.com/stowelly/aws-sdk-go/service/directoryservice"
	"github.com/stowelly/aws-sdk-go/service/dynamodb"
	"github.com/stowelly/aws-sdk-go/service/dynamodbstreams"
	"github.com/stowelly/aws-sdk-go/service/ec2"
	"github.com/stowelly/aws-sdk-go/service/ecr"
	"github.com/stowelly/aws-sdk-go/service/ecs"
	"github.com/stowelly/aws-sdk-go/service/efs"
	"github.com/stowelly/aws-sdk-go/service/elasticache"
	"github.com/stowelly/aws-sdk-go/service/elasticbeanstalk"
	"github.com/stowelly/aws-sdk-go/service/elasticsearchservice"
	"github.com/stowelly/aws-sdk-go/service/elastictranscoder"
	"github.com/stowelly/aws-sdk-go/service/elb"
	"github.com/stowelly/aws-sdk-go/service/emr"
	"github.com/stowelly/aws-sdk-go/service/firehose"
	"github.com/stowelly/aws-sdk-go/service/glacier"
	"github.com/stowelly/aws-sdk-go/service/iam"
	"github.com/stowelly/aws-sdk-go/service/inspector"
	"github.com/stowelly/aws-sdk-go/service/iot"
	"github.com/stowelly/aws-sdk-go/service/iotdataplane"
	"github.com/stowelly/aws-sdk-go/service/kinesis"
	"github.com/stowelly/aws-sdk-go/service/kms"
	"github.com/stowelly/aws-sdk-go/service/lambda"
	"github.com/stowelly/aws-sdk-go/service/machinelearning"
	"github.com/stowelly/aws-sdk-go/service/marketplacecommerceanalytics"
	"github.com/stowelly/aws-sdk-go/service/mobileanalytics"
	"github.com/stowelly/aws-sdk-go/service/opsworks"
	"github.com/stowelly/aws-sdk-go/service/rds"
	"github.com/stowelly/aws-sdk-go/service/redshift"
	"github.com/stowelly/aws-sdk-go/service/route53"
	"github.com/stowelly/aws-sdk-go/service/route53domains"
	"github.com/stowelly/aws-sdk-go/service/s3"
	"github.com/stowelly/aws-sdk-go/service/ses"
	"github.com/stowelly/aws-sdk-go/service/simpledb"
	"github.com/stowelly/aws-sdk-go/service/sns"
	"github.com/stowelly/aws-sdk-go/service/sqs"
	"github.com/stowelly/aws-sdk-go/service/ssm"
	"github.com/stowelly/aws-sdk-go/service/storagegateway"
	"github.com/stowelly/aws-sdk-go/service/sts"
	"github.com/stowelly/aws-sdk-go/service/support"
	"github.com/stowelly/aws-sdk-go/service/swf"
	"github.com/stowelly/aws-sdk-go/service/waf"
	"github.com/stowelly/aws-sdk-go/service/workspaces"
)

var clients = []reflect.Value{
	reflect.ValueOf(acm.New),
	reflect.ValueOf(apigateway.New),
	reflect.ValueOf(autoscaling.New),
	reflect.ValueOf(cloudformation.New),
	reflect.ValueOf(cloudfront.New),
	reflect.ValueOf(cloudhsm.New),
	reflect.ValueOf(cloudsearch.New),
	reflect.ValueOf(cloudsearchdomain.New),
	reflect.ValueOf(cloudtrail.New),
	reflect.ValueOf(cloudwatch.New),
	reflect.ValueOf(cloudwatchevents.New),
	reflect.ValueOf(cloudwatchlogs.New),
	reflect.ValueOf(codecommit.New),
	reflect.ValueOf(codedeploy.New),
	reflect.ValueOf(codepipeline.New),
	reflect.ValueOf(cognitoidentity.New),
	reflect.ValueOf(cognitosync.New),
	reflect.ValueOf(configservice.New),
	reflect.ValueOf(datapipeline.New),
	reflect.ValueOf(devicefarm.New),
	reflect.ValueOf(directconnect.New),
	reflect.ValueOf(directoryservice.New),
	reflect.ValueOf(dynamodb.New),
	reflect.ValueOf(dynamodbstreams.New),
	reflect.ValueOf(ec2.New),
	reflect.ValueOf(ecr.New),
	reflect.ValueOf(ecs.New),
	reflect.ValueOf(efs.New),
	reflect.ValueOf(elasticache.New),
	reflect.ValueOf(elasticbeanstalk.New),
	reflect.ValueOf(elasticsearchservice.New),
	reflect.ValueOf(elastictranscoder.New),
	reflect.ValueOf(elb.New),
	reflect.ValueOf(emr.New),
	reflect.ValueOf(firehose.New),
	reflect.ValueOf(glacier.New),
	reflect.ValueOf(iam.New),
	reflect.ValueOf(inspector.New),
	reflect.ValueOf(iot.New),
	reflect.ValueOf(iotdataplane.New),
	reflect.ValueOf(kinesis.New),
	reflect.ValueOf(kms.New),
	reflect.ValueOf(lambda.New),
	reflect.ValueOf(machinelearning.New),
	reflect.ValueOf(marketplacecommerceanalytics.New),
	reflect.ValueOf(mobileanalytics.New),
	reflect.ValueOf(opsworks.New),
	reflect.ValueOf(rds.New),
	reflect.ValueOf(redshift.New),
	reflect.ValueOf(route53.New),
	reflect.ValueOf(route53domains.New),
	reflect.ValueOf(s3.New),
	reflect.ValueOf(ses.New),
	reflect.ValueOf(simpledb.New),
	reflect.ValueOf(sns.New),
	reflect.ValueOf(sqs.New),
	reflect.ValueOf(ssm.New),
	reflect.ValueOf(storagegateway.New),
	reflect.ValueOf(sts.New),
	reflect.ValueOf(support.New),
	reflect.ValueOf(swf.New),
	reflect.ValueOf(waf.New),
	reflect.ValueOf(workspaces.New),
}
