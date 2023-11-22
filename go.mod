module github.com/kaytu-io/kaytu-aws-describer

go 1.21

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/aws/aws-sdk-go v1.45.19
	github.com/aws/aws-sdk-go-v2 v1.21.0
	github.com/aws/aws-sdk-go-v2/config v1.18.42
	github.com/aws/aws-sdk-go-v2/credentials v1.13.40
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.21.0
	github.com/aws/aws-sdk-go-v2/service/account v1.11.5
	github.com/aws/aws-sdk-go-v2/service/acm v1.19.0
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.22.7
	github.com/aws/aws-sdk-go-v2/service/amp v1.17.5
	github.com/aws/aws-sdk-go-v2/service/amplify v1.15.0
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.18.0
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.14.5
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.20.0
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.22.5
	github.com/aws/aws-sdk-go-v2/service/applicationinsights v1.19.5
	github.com/aws/aws-sdk-go-v2/service/appstream v1.24.0
	github.com/aws/aws-sdk-go-v2/service/athena v1.31.6
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.26.5
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.30.6
	github.com/aws/aws-sdk-go-v2/service/backup v1.25.0
	github.com/aws/aws-sdk-go-v2/service/batch v1.26.5
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.12.5
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.34.6
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.28.5
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v1.16.0
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.29.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.27.7
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.24.0
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.20.0
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.22.0
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.16.2
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.18.1
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.16.5
	github.com/aws/aws-sdk-go-v2/service/codestar v1.15.0
	github.com/aws/aws-sdk-go-v2/service/configservice v1.36.3
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.28.0
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.31.0
	github.com/aws/aws-sdk-go-v2/service/dax v1.14.0
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.19.5
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.18.5
	github.com/aws/aws-sdk-go-v2/service/dlm v1.16.5
	github.com/aws/aws-sdk-go-v2/service/docdb v1.23.5
	github.com/aws/aws-sdk-go-v2/service/drs v1.16.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.22.0
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.15.5
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.122.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.20.0
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.18.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.30.1
	github.com/aws/aws-sdk-go-v2/service/efs v1.21.7
	github.com/aws/aws-sdk-go-v2/service/eks v1.29.5
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.29.3
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.17.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.17.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.21.4
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.20.6
	github.com/aws/aws-sdk-go-v2/service/emr v1.28.6
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.22.0
	github.com/aws/aws-sdk-go-v2/service/firehose v1.19.0
	github.com/aws/aws-sdk-go-v2/service/fms v1.25.5
	github.com/aws/aws-sdk-go-v2/service/fsx v1.32.6
	github.com/aws/aws-sdk-go-v2/service/glacier v1.16.0
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.17.7
	github.com/aws/aws-sdk-go-v2/service/glue v1.62.0
	github.com/aws/aws-sdk-go-v2/service/grafana v1.15.5
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.28.0
	github.com/aws/aws-sdk-go-v2/service/health v1.19.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.22.5
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.18.2
	github.com/aws/aws-sdk-go-v2/service/imagebuilder v1.24.5
	github.com/aws/aws-sdk-go-v2/service/inspector v1.15.0
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.16.7
	github.com/aws/aws-sdk-go-v2/service/kafka v1.22.6
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.4.5
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.19.0
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.18.6
	github.com/aws/aws-sdk-go-v2/service/kinesisvideo v1.18.4
	github.com/aws/aws-sdk-go-v2/service/kms v1.24.5
	github.com/aws/aws-sdk-go-v2/service/lambda v1.39.5
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.28.5
	github.com/aws/aws-sdk-go-v2/service/macie2 v1.29.7
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.15.0
	github.com/aws/aws-sdk-go-v2/service/memorydb v1.14.5
	github.com/aws/aws-sdk-go-v2/service/mgn v1.20.6
	github.com/aws/aws-sdk-go-v2/service/mq v1.16.5
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.17.5
	github.com/aws/aws-sdk-go-v2/service/neptune v1.22.0
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.30.1
	github.com/aws/aws-sdk-go-v2/service/oam v1.3.0
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.19.6
	github.com/aws/aws-sdk-go-v2/service/opensearchserverless v1.5.0
	github.com/aws/aws-sdk-go-v2/service/opsworkscm v1.17.0
	github.com/aws/aws-sdk-go-v2/service/organizations v1.20.6
	github.com/aws/aws-sdk-go-v2/service/pinpoint v1.22.6
	github.com/aws/aws-sdk-go-v2/service/pipes v1.4.1
	github.com/aws/aws-sdk-go-v2/service/pricing v1.21.7
	github.com/aws/aws-sdk-go-v2/service/ram v1.20.5
	github.com/aws/aws-sdk-go-v2/service/rds v1.54.0
	github.com/aws/aws-sdk-go-v2/service/redshift v1.29.5
	github.com/aws/aws-sdk-go-v2/service/redshiftserverless v1.6.0
	github.com/aws/aws-sdk-go-v2/service/resourceexplorer2 v1.4.0
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.16.0
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.16.0
	github.com/aws/aws-sdk-go-v2/service/route53 v1.29.5
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.17.3
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.20.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.40.0
	github.com/aws/aws-sdk-go-v2/service/s3control v1.33.0
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.108.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.21.3
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.36.2
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.3.6
	github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository v1.14.0
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.21.3
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.24.0
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.16.0
	github.com/aws/aws-sdk-go-v2/service/ses v1.16.8
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.20.1
	github.com/aws/aws-sdk-go-v2/service/sfn v1.19.5
	github.com/aws/aws-sdk-go-v2/service/shield v1.20.0
	github.com/aws/aws-sdk-go-v2/service/simspaceweaver v1.5.1
	github.com/aws/aws-sdk-go-v2/service/sns v1.22.0
	github.com/aws/aws-sdk-go-v2/service/sqs v1.24.5
	github.com/aws/aws-sdk-go-v2/service/ssm v1.38.0
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.18.2
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.20.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.22.0
	github.com/aws/aws-sdk-go-v2/service/support v1.16.5
	github.com/aws/aws-sdk-go-v2/service/synthetics v1.19.0
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.19.0
	github.com/aws/aws-sdk-go-v2/service/waf v1.14.0
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.15.0
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.39.1
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.22.5
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.30.0
	github.com/aws/smithy-go v1.14.2
	github.com/ghodss/yaml v1.0.0
	github.com/go-errors/errors v1.4.2
	github.com/gocarina/gocsv v0.0.0-20211203214250-4735fba0c1d9
	github.com/golang-jwt/jwt/v5 v5.0.0
	github.com/golang/protobuf v1.5.3
	github.com/hashicorp/go-hclog v1.5.0
	github.com/kaytu-io/kaytu-util v0.0.0-20231122143053-67433de2347b
	github.com/manifoldco/promptui v0.9.0
	github.com/spf13/cobra v1.7.0
	github.com/turbot/go-kit v0.8.0-rc.0
	github.com/turbot/steampipe-plugin-sdk/v5 v5.6.2
	go.uber.org/zap v1.25.0
	golang.org/x/oauth2 v0.11.0
	golang.org/x/tools v0.10.0
	google.golang.org/genproto v0.0.0-20230822172742-b8732ec3820d
	google.golang.org/grpc v1.59.0
)

require (
	cloud.google.com/go v0.110.7 // indirect
	cloud.google.com/go/compute v1.23.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v1.1.1 // indirect
	cloud.google.com/go/longrunning v0.5.1 // indirect
	cloud.google.com/go/storage v1.30.1 // indirect
	github.com/XiaoMi/pegasus-go-client v0.0.0-20210427083443-f3b6b08bc4c2 // indirect
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/allegro/bigcache/v3 v3.1.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.13 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.11 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.41 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.35 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.43 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.1.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.36 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.35 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.35 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.15.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.14.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.17.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/bradfitz/gomemcache v0.0.0-20221031212613-62deef7fc822 // indirect
	github.com/btubbs/datetime v0.1.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/confluentinc/confluent-kafka-go/v2 v2.1.1 // indirect
	github.com/danwakefield/fnmatch v0.0.0-20160403171240-cbb64ac3d964 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eko/gocache/v3 v3.1.2 // indirect
	github.com/elastic/go-elasticsearch/v7 v7.17.10 // indirect
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/fluxcd/helm-controller/api v0.36.1 // indirect
	github.com/fluxcd/pkg/apis/kustomize v1.1.1 // indirect
	github.com/fluxcd/pkg/apis/meta v1.1.2 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.1.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/s2a-go v0.1.4 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/googleapis/gax-go/v2 v2.11.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-getter v1.7.2 // indirect
	github.com/hashicorp/go-plugin v1.5.2 // indirect
	github.com/hashicorp/go-safetemp v1.0.0 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hcl/v2 v2.18.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20211028200310-0bc27b27de87 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jackc/puddle v1.3.0 // indirect
	github.com/jackc/puddle/v2 v2.2.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.11 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/opensearch-project/opensearch-go/v2 v2.3.0 // indirect
	github.com/pegasus-kv/thrift v0.13.0 // indirect
	github.com/pganalyze/pg_query_go/v4 v4.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.16.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/rabbitmq/amqp091-go v1.9.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/sethvargo/go-retry v0.2.4 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stevenle/topsort v0.2.0 // indirect
	github.com/tkrajina/go-reflector v0.5.6 // indirect
	github.com/ulikunitz/xz v0.5.11 // indirect
	github.com/zclconf/go-cty v1.14.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/otel v1.17.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.17.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.40.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.40.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.16.0 // indirect
	go.opentelemetry.io/otel/metric v1.17.0 // indirect
	go.opentelemetry.io/otel/sdk v1.17.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.40.0 // indirect
	go.opentelemetry.io/otel/trace v1.17.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/term v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	google.golang.org/api v0.126.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.28.1 // indirect
	k8s.io/apiextensions-apiserver v0.28.0 // indirect
	k8s.io/apimachinery v0.28.1 // indirect
	k8s.io/client-go v0.28.1 // indirect
	k8s.io/component-base v0.28.1 // indirect
	k8s.io/klog/v2 v2.100.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230717233707-2695361300d9 // indirect
	k8s.io/utils v0.0.0-20230406110748-d93618cff8a2 // indirect
	sigs.k8s.io/controller-runtime v0.16.2 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
