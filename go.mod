module github.com/kaytu-io/kaytu-aws-describer

go 1.21

require (
	github.com/aws/aws-lambda-go v1.42.0
	github.com/aws/aws-sdk-go v1.49.10
	github.com/aws/aws-sdk-go-v2 v1.24.1
	github.com/aws/aws-sdk-go-v2/config v1.26.6
	github.com/aws/aws-sdk-go-v2/credentials v1.16.16
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.26.5
	github.com/aws/aws-sdk-go-v2/service/account v1.14.5
	github.com/aws/aws-sdk-go-v2/service/acm v1.22.5
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.25.5
	github.com/aws/aws-sdk-go-v2/service/amp v1.21.5
	github.com/aws/aws-sdk-go-v2/service/amplify v1.18.5
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.21.5
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.18.5
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.26.5
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.25.5
	github.com/aws/aws-sdk-go-v2/service/applicationinsights v1.22.5
	github.com/aws/aws-sdk-go-v2/service/appstream v1.30.0
	github.com/aws/aws-sdk-go-v2/service/athena v1.37.3
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.30.5
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.36.5
	github.com/aws/aws-sdk-go-v2/service/backup v1.31.2
	github.com/aws/aws-sdk-go-v2/service/batch v1.30.5
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.15.5
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.42.4
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.32.5
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v1.20.5
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.35.5
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.32.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.30.0
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.23.5
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.26.5
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.19.5
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.22.1
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.22.5
	github.com/aws/aws-sdk-go-v2/service/codestar v1.19.5
	github.com/aws/aws-sdk-go-v2/service/configservice v1.43.5
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.33.5
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.35.5
	github.com/aws/aws-sdk-go-v2/service/dax v1.17.5
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.22.5
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.22.5
	github.com/aws/aws-sdk-go-v2/service/dlm v1.22.5
	github.com/aws/aws-sdk-go-v2/service/docdb v1.29.5
	github.com/aws/aws-sdk-go-v2/service/drs v1.21.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.26.6
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.18.5
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.141.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.24.5
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.21.5
	github.com/aws/aws-sdk-go-v2/service/ecs v1.35.5
	github.com/aws/aws-sdk-go-v2/service/efs v1.26.5
	github.com/aws/aws-sdk-go-v2/service/eks v1.35.5
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.34.5
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.20.5
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.21.5
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.26.5
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.24.5
	github.com/aws/aws-sdk-go-v2/service/emr v1.35.5
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.26.5
	github.com/aws/aws-sdk-go-v2/service/firehose v1.23.0
	github.com/aws/aws-sdk-go-v2/service/fms v1.29.5
	github.com/aws/aws-sdk-go-v2/service/fsx v1.39.5
	github.com/aws/aws-sdk-go-v2/service/glacier v1.19.5
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.20.5
	github.com/aws/aws-sdk-go-v2/service/glue v1.72.4
	github.com/aws/aws-sdk-go-v2/service/grafana v1.18.5
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.35.5
	github.com/aws/aws-sdk-go-v2/service/health v1.22.5
	github.com/aws/aws-sdk-go-v2/service/iam v1.28.5
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.21.5
	github.com/aws/aws-sdk-go-v2/service/imagebuilder v1.30.0
	github.com/aws/aws-sdk-go-v2/service/inspector v1.19.5
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.20.5
	github.com/aws/aws-sdk-go-v2/service/kafka v1.28.5
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.7.5
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.24.5
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.21.5
	github.com/aws/aws-sdk-go-v2/service/kinesisvideo v1.21.5
	github.com/aws/aws-sdk-go-v2/service/kms v1.27.9
	github.com/aws/aws-sdk-go-v2/service/lambda v1.49.5
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.32.5
	github.com/aws/aws-sdk-go-v2/service/macie2 v1.34.5
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.18.5
	github.com/aws/aws-sdk-go-v2/service/memorydb v1.17.5
	github.com/aws/aws-sdk-go-v2/service/mgn v1.25.5
	github.com/aws/aws-sdk-go-v2/service/mq v1.20.6
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.22.5
	github.com/aws/aws-sdk-go-v2/service/neptune v1.28.0
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.36.5
	github.com/aws/aws-sdk-go-v2/service/oam v1.7.5
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.27.0
	github.com/aws/aws-sdk-go-v2/service/opensearchserverless v1.9.5
	github.com/aws/aws-sdk-go-v2/service/opsworkscm v1.20.5
	github.com/aws/aws-sdk-go-v2/service/organizations v1.23.5
	github.com/aws/aws-sdk-go-v2/service/pinpoint v1.26.6
	github.com/aws/aws-sdk-go-v2/service/pipes v1.9.6
	github.com/aws/aws-sdk-go-v2/service/pricing v1.24.5
	github.com/aws/aws-sdk-go-v2/service/ram v1.23.6
	github.com/aws/aws-sdk-go-v2/service/rds v1.64.6
	github.com/aws/aws-sdk-go-v2/service/redshift v1.39.6
	github.com/aws/aws-sdk-go-v2/service/redshiftserverless v1.15.4
	github.com/aws/aws-sdk-go-v2/service/resourceexplorer2 v1.8.5
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.19.5
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.19.5
	github.com/aws/aws-sdk-go-v2/service/route53 v1.35.5
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.20.5
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.23.5
	github.com/aws/aws-sdk-go-v2/service/s3 v1.47.5
	github.com/aws/aws-sdk-go-v2/service/s3control v1.41.5
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.121.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.25.5
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.44.0
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.10.5
	github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository v1.18.5
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.25.5
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.27.5
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.19.5
	github.com/aws/aws-sdk-go-v2/service/ses v1.19.5
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.24.5
	github.com/aws/aws-sdk-go-v2/service/sfn v1.24.5
	github.com/aws/aws-sdk-go-v2/service/shield v1.23.5
	github.com/aws/aws-sdk-go-v2/service/simspaceweaver v1.8.5
	github.com/aws/aws-sdk-go-v2/service/sns v1.26.5
	github.com/aws/aws-sdk-go-v2/service/sqs v1.29.5
	github.com/aws/aws-sdk-go-v2/service/ssm v1.44.5
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.23.5
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.24.5
	github.com/aws/aws-sdk-go-v2/service/sts v1.26.7
	github.com/aws/aws-sdk-go-v2/service/support v1.19.5
	github.com/aws/aws-sdk-go-v2/service/synthetics v1.22.5
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.23.6
	github.com/aws/aws-sdk-go-v2/service/waf v1.18.5
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.19.5
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.43.5
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.27.5
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.35.6
	github.com/aws/smithy-go v1.19.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-errors/errors v1.4.2
	github.com/gocarina/gocsv v0.0.0-20211203214250-4735fba0c1d9
	github.com/golang-jwt/jwt/v5 v5.0.0
	github.com/golang/protobuf v1.5.3
	github.com/hashicorp/go-hclog v1.6.2
	github.com/kaytu-io/kaytu-util v0.0.0-20240228141953-c0d2f23a5b77
	github.com/manifoldco/promptui v0.9.0
	github.com/spf13/cobra v1.7.0
	github.com/turbot/go-kit v0.9.0
	github.com/turbot/steampipe-plugin-sdk/v5 v5.8.0
	go.uber.org/zap v1.26.0
	golang.org/x/oauth2 v0.15.0
	golang.org/x/tools v0.16.1
	google.golang.org/genproto v0.0.0-20231212172506-995d672761c0
	google.golang.org/grpc v1.61.0
)

require (
	cloud.google.com/go v0.111.0 // indirect
	cloud.google.com/go/compute v1.23.3 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v1.1.5 // indirect
	cloud.google.com/go/longrunning v0.5.4 // indirect
	cloud.google.com/go/storage v1.36.0 // indirect
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/allegro/bigcache/v3 v3.1.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.4 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.11 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.2.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.2.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.8.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.16.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.18.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.21.7 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/btubbs/datetime v0.1.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/danwakefield/fnmatch v0.0.0-20160403171240-cbb64ac3d964 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eko/gocache/lib/v4 v4.1.5 // indirect
	github.com/eko/gocache/store/bigcache/v4 v4.2.1 // indirect
	github.com/eko/gocache/store/ristretto/v4 v4.2.1 // indirect
	github.com/elastic/go-elasticsearch/v7 v7.17.10 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/evanphx/json-patch/v5 v5.8.0 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fluxcd/helm-controller/api v0.37.4 // indirect
	github.com/fluxcd/pkg/apis/kustomize v1.3.0 // indirect
	github.com/fluxcd/pkg/apis/meta v1.3.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.20.2 // indirect
	github.com/go-openapi/jsonreference v0.20.4 // indirect
	github.com/go-openapi/swag v0.22.6 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.2.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.1 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-getter v1.7.3 // indirect
	github.com/hashicorp/go-plugin v1.6.0 // indirect
	github.com/hashicorp/go-safetemp v1.0.0 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hcl/v2 v2.19.1 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.2 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/jackc/pgx/v4 v4.18.1 // indirect
	github.com/jackc/pgx/v5 v5.5.3 // indirect
	github.com/jackc/puddle v1.3.0 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
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
	github.com/pganalyze/pg_query_go/v4 v4.2.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rabbitmq/amqp091-go v1.9.0 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/sethvargo/go-retry v0.2.4 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stevenle/topsort v0.2.0 // indirect
	github.com/tkrajina/go-reflector v0.5.6 // indirect
	github.com/ulikunitz/xz v0.5.11 // indirect
	github.com/zclconf/go-cty v1.14.1 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.46.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.46.1 // indirect
	go.opentelemetry.io/otel v1.23.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.44.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.21.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.23.1 // indirect
	go.opentelemetry.io/otel/sdk v1.23.1 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.23.1 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/exp v0.0.0-20231226003508-02704c960a9b // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/term v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	google.golang.org/api v0.154.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231212172506-995d672761c0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231212172506-995d672761c0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.29.1 // indirect
	k8s.io/apiextensions-apiserver v0.29.0 // indirect
	k8s.io/apimachinery v0.29.1 // indirect
	k8s.io/client-go v0.29.0 // indirect
	k8s.io/component-base v0.29.0 // indirect
	k8s.io/klog/v2 v2.110.1 // indirect
	k8s.io/kube-openapi v0.0.0-20231214164306-ab13479f8bf8 // indirect
	k8s.io/utils v0.0.0-20231127182322-b307cd553661 // indirect
	sigs.k8s.io/controller-runtime v0.17.0 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
