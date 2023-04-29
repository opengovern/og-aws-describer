module github.com/kaytu-io/kaytu-aws-describer

go 1.19

require (
	github.com/aws/aws-lambda-go v1.13.3
	github.com/aws/aws-sdk-go v1.44.153
	github.com/aws/aws-sdk-go-v2 v1.17.3
	github.com/aws/aws-sdk-go-v2/config v1.17.8
	github.com/aws/aws-sdk-go-v2/credentials v1.12.21
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.16.0
	github.com/aws/aws-sdk-go-v2/service/account v1.7.8
	github.com/aws/aws-sdk-go-v2/service/acm v1.14.8
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.21.0
	github.com/aws/aws-sdk-go-v2/service/amp v1.15.7
	github.com/aws/aws-sdk-go-v2/service/amplify v1.11.18
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.15.10
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.12.8
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.13.7
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.15.18
	github.com/aws/aws-sdk-go-v2/service/applicationinsights v1.8.0
	github.com/aws/aws-sdk-go-v2/service/appstream v1.17.13
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.20.4
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.23.10
	github.com/aws/aws-sdk-go-v2/service/backup v1.18.0
	github.com/aws/aws-sdk-go-v2/service/batch v1.19.0
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.10.13
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.23.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.20.7
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v1.13.19
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.21.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.21.6
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.15.14
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.13.11
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.19.13
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.13.19
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.15.2
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.13.19
	github.com/aws/aws-sdk-go-v2/service/codestar v1.12.3
	github.com/aws/aws-sdk-go-v2/service/configservice v1.28.0
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.19.2
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.21.10
	github.com/aws/aws-sdk-go-v2/service/dax v1.11.15
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.18.0
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.15.3
	github.com/aws/aws-sdk-go-v2/service/dlm v1.12.4
	github.com/aws/aws-sdk-go-v2/service/docdb v1.19.11
	github.com/aws/aws-sdk-go-v2/service/drs v1.9.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.15.9
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.13.22
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.74.1
	github.com/aws/aws-sdk-go-v2/service/ecr v1.17.16
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.13.15
	github.com/aws/aws-sdk-go-v2/service/ecs v1.21.0
	github.com/aws/aws-sdk-go-v2/service/efs v1.17.15
	github.com/aws/aws-sdk-go-v2/service/eks v1.26.0
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.22.10
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.14.18
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.14.12
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.18.12
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.16.10
	github.com/aws/aws-sdk-go-v2/service/emr v1.20.11
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.16.17
	github.com/aws/aws-sdk-go-v2/service/fms v1.22.0
	github.com/aws/aws-sdk-go-v2/service/fsx v1.24.14
	github.com/aws/aws-sdk-go-v2/service/glacier v1.13.19
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.15.2
	github.com/aws/aws-sdk-go-v2/service/glue v1.38.1
	github.com/aws/aws-sdk-go-v2/service/grafana v1.9.15
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.15.9
	github.com/aws/aws-sdk-go-v2/service/health v1.15.22
	github.com/aws/aws-sdk-go-v2/service/iam v1.18.9
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.15.5
	github.com/aws/aws-sdk-go-v2/service/imagebuilder v1.20.3
	github.com/aws/aws-sdk-go-v2/service/inspector v1.12.15
	github.com/aws/aws-sdk-go-v2/service/kafka v1.18.0
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.1.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.15.21
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.14.18
	github.com/aws/aws-sdk-go-v2/service/kms v1.18.11
	github.com/aws/aws-sdk-go-v2/service/lambda v1.26.0
	github.com/aws/aws-sdk-go-v2/service/memorydb v1.12.0
	github.com/aws/aws-sdk-go-v2/service/mq v1.13.15
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.13.12
	github.com/aws/aws-sdk-go-v2/service/neptune v1.18.0
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.20.3
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.10.12
	github.com/aws/aws-sdk-go-v2/service/opsworkscm v1.15.0
	github.com/aws/aws-sdk-go-v2/service/organizations v1.16.8
	github.com/aws/aws-sdk-go-v2/service/rds v1.26.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.26.10
	github.com/aws/aws-sdk-go-v2/service/redshiftserverless v1.2.13
	github.com/aws/aws-sdk-go-v2/service/route53 v1.24.0
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.15.19
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.1
	github.com/aws/aws-sdk-go-v2/service/s3control v1.21.9
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.48.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.16.2
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.26.0
	github.com/aws/aws-sdk-go-v2/service/ses v1.14.18
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.16.0
	github.com/aws/aws-sdk-go-v2/service/shield v1.18.0
	github.com/aws/aws-sdk-go-v2/service/sns v1.17.9
	github.com/aws/aws-sdk-go-v2/service/sqs v1.19.10
	github.com/aws/aws-sdk-go-v2/service/ssm v1.30.0
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.15.13
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.18.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.19
	github.com/aws/aws-sdk-go-v2/service/support v1.11.0
	github.com/aws/aws-sdk-go-v2/service/synthetics v1.9.1
	github.com/aws/aws-sdk-go-v2/service/waf v1.11.19
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.12.18
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.22.9
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.23.0
	github.com/aws/smithy-go v1.13.5
	github.com/go-errors/errors v1.4.2
	github.com/gocarina/gocsv v0.0.0-20211203214250-4735fba0c1d9
	github.com/spf13/cobra v1.6.1
	github.com/turbot/go-kit v0.4.0
	go.uber.org/zap v1.21.0
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.9 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.17 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.27 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.21 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.24 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.20 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.13.6 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	github.com/tkrajina/go-reflector v0.5.4 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/goleak v1.1.12 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20220407144326-9054f6ed7bac // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
