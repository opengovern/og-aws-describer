module github.com/kaytu-io/kaytu-aws-describer

go 1.19

require (
	github.com/aws/aws-lambda-go v1.13.3
	github.com/aws/aws-sdk-go v1.44.153
	github.com/aws/aws-sdk-go-v2 v1.18.0
	github.com/aws/aws-sdk-go-v2/config v1.18.25
	github.com/aws/aws-sdk-go-v2/credentials v1.13.24
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.16.0
	github.com/aws/aws-sdk-go-v2/service/account v1.10.6
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
	github.com/aws/aws-sdk-go-v2/service/dlm v1.15.3
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
	github.com/aws/aws-sdk-go-v2/service/firehose v1.16.12
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
	github.com/aws/aws-sdk-go-v2/service/imagebuilder v1.23.4
	github.com/aws/aws-sdk-go-v2/service/inspector v1.12.15
	github.com/aws/aws-sdk-go-v2/service/kafka v1.18.0
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.1.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.15.21
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.14.18
	github.com/aws/aws-sdk-go-v2/service/kinesisvideo v1.15.11
	github.com/aws/aws-sdk-go-v2/service/kms v1.21.1
	github.com/aws/aws-sdk-go-v2/service/lambda v1.26.0
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.26.6
	github.com/aws/aws-sdk-go-v2/service/macie2 v1.27.7
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.13.10
	github.com/aws/aws-sdk-go-v2/service/memorydb v1.12.0
	github.com/aws/aws-sdk-go-v2/service/mgn v1.18.5
	github.com/aws/aws-sdk-go-v2/service/mq v1.13.15
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.13.12
	github.com/aws/aws-sdk-go-v2/service/neptune v1.18.0
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.20.3
	github.com/aws/aws-sdk-go-v2/service/oam v1.1.11
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.10.12
	github.com/aws/aws-sdk-go-v2/service/opsworkscm v1.15.0
	github.com/aws/aws-sdk-go-v2/service/organizations v1.16.8
	github.com/aws/aws-sdk-go-v2/service/pinpoint v1.19.1
	github.com/aws/aws-sdk-go-v2/service/pipes v1.2.6
	github.com/aws/aws-sdk-go-v2/service/ram v1.18.2
	github.com/aws/aws-sdk-go-v2/service/rds v1.26.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.26.10
	github.com/aws/aws-sdk-go-v2/service/redshiftserverless v1.2.13
	github.com/aws/aws-sdk-go-v2/service/resourceexplorer2 v1.2.13
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.14.11
	github.com/aws/aws-sdk-go-v2/service/route53 v1.24.0
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.14.10
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.15.19
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.1
	github.com/aws/aws-sdk-go-v2/service/s3control v1.21.9
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.48.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.16.2
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.26.0
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.3.6
	github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository v1.12.11
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.14.12
	github.com/aws/aws-sdk-go-v2/service/ses v1.14.18
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.16.0
	github.com/aws/aws-sdk-go-v2/service/sfn v1.17.11
	github.com/aws/aws-sdk-go-v2/service/shield v1.18.0
	github.com/aws/aws-sdk-go-v2/service/simspaceweaver v1.2.1
	github.com/aws/aws-sdk-go-v2/service/sns v1.17.9
	github.com/aws/aws-sdk-go-v2/service/sqs v1.19.10
	github.com/aws/aws-sdk-go-v2/service/ssm v1.36.4
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.15.13
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.18.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.19.0
	github.com/aws/aws-sdk-go-v2/service/support v1.11.0
	github.com/aws/aws-sdk-go-v2/service/synthetics v1.9.1
	github.com/aws/aws-sdk-go-v2/service/waf v1.11.19
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.12.18
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.22.9
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.20.1
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.23.0
	github.com/aws/smithy-go v1.13.5
	github.com/go-errors/errors v1.4.2
	github.com/gocarina/gocsv v0.0.0-20211203214250-4735fba0c1d9
	github.com/golang-jwt/jwt/v5 v5.0.0
	github.com/hashicorp/go-hclog v1.2.2
	github.com/kaytu-io/kaytu-util v0.0.0-20230605084433-73e3af414279
	github.com/spf13/cobra v1.7.0
	github.com/turbot/go-kit v0.4.0
	github.com/turbot/steampipe-plugin-sdk/v4 v4.1.13
	gitlab.com/keibiengine/steampipe-plugin-aws v0.0.0-20230606154334-9ac06ff8fb56
	go.uber.org/zap v1.24.0
	golang.org/x/oauth2 v0.6.0
	google.golang.org/grpc v1.54.0
)

require (
	cloud.google.com/go/compute v1.19.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/XiaoMi/pegasus-go-client v0.0.0-20210427083443-f3b6b08bc4c2 // indirect
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d // indirect
	github.com/agext/levenshtein v1.2.2 // indirect
	github.com/allegro/bigcache/v3 v3.0.2 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.9 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.33 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.27 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.27 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/pricing v1.16.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.10 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bradfitz/gomemcache v0.0.0-20220106215444-fb4bf637b56d // indirect
	github.com/btubbs/datetime v0.1.1 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/eko/gocache/v3 v3.1.1 // indirect
	github.com/elastic/go-elasticsearch/v7 v7.17.10 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3 // indirect
	github.com/hashicorp/go-plugin v1.4.4 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20211028200310-0bc27b27de87 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jackc/puddle v1.3.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pegasus-kv/thrift v0.13.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.15.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/sethvargo/go-retry v0.1.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stevenle/topsort v0.0.0-20130922064739-8130c1d7596b // indirect
	github.com/streadway/amqp v1.0.0 // indirect
	github.com/tkrajina/go-reflector v0.5.4 // indirect
	github.com/zclconf/go-cty v1.10.0 // indirect
	go.opentelemetry.io/otel v1.8.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.30.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.30.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.7.0 // indirect
	go.opentelemetry.io/otel/metric v0.30.0 // indirect
	go.opentelemetry.io/otel/sdk v1.8.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.8.0 // indirect
	go.opentelemetry.io/proto/otlp v0.16.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/exp v0.0.0-20220827204233-334a2380cb91 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230331144136-dcfb400f0633 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apimachinery v0.24.2 // indirect
)
