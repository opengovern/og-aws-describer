package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSageMakerTrainingJob(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_sagemaker_training_job",
		Description: "AWS SageMaker Training Job",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException", "NotFoundException", "RecordNotFound"}),
			},
			Hydrate: kaytu.GetSageMakerTrainingJob,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSageMakerTrainingJob,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "creation_time", Require: plugin.Optional, Operators: []string{">", ">=", "<", "<="}},
				{Name: "last_modified_time", Require: plugin.Optional, Operators: []string{">", ">=", "<", "<="}},
				{Name: "training_job_status", Require: plugin.Optional, Operators: []string{">", ">=", "<", "<="}},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the training job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrainingJob.TrainingJobName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the training job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrainingJob.TrainingJobArn"),
			},
			{
				Name:        "training_job_status",
				Description: "The status of the training job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrainingJob.TrainingJobStatus")},
			{
				Name:        "auto_ml_job_arn",
				Description: "The Amazon Resource Name (ARN) of an AutoML job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrainingJob.AutoMLJobArn")},
			{
				Name:        "billable_time_in_seconds",
				Description: "The billable time in seconds. Billable time refers to the absolute wall-clock time.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TrainingJob.BillableTimeInSeconds")},
			{
				Name:        "creation_time",
				Description: "A timestamp that shows when the training job was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.TrainingJob.CreationTime")},
			{
				Name:        "enable_managed_spot_training",
				Description: "A Boolean indicating whether managed spot training is enabled or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.TrainingJob.EnableManagedSpotTraining")},
			{
				Name:        "enable_network_isolation",
				Description: "Specifies enable network isolation for training jobs.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.TrainingJob.EnableNetworkIsolation")},
			{
				Name:        "enable_inter_container_traffic_encryption",
				Description: "To encrypt all communications between ML compute instances in distributed training, choose True.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.TrainingJob.EnableInterContainerTrafficEncryption")},
			{
				Name:        "failure_reason",
				Description: "If the training job failed, the reason it failed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrainingJob.FailureReason")},
			{
				Name:        "labeling_job_arn",
				Description: "The Amazon Resource Name (ARN) of the Amazon SageMaker Ground Truth labeling job that created the transform or training job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrainingJob.LabelingJobArn")},
			{
				Name:        "last_modified_time",
				Description: "Timestamp when the training job was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.TrainingJob.LastModifiedTime")},
			{
				Name:        "profiling_status",
				Description: "Profiling status of a training job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrainingJob.ProfilingStatus")},
			{
				Name:        "role_arn",
				Description: "The AWS Identity and Access Management (IAM) role configured for the training job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrainingJob.RoleArn")},
			{
				Name:        "secondary_status",
				Description: "Provides detailed information about the state of the training job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrainingJob.SecondaryStatus")},
			{
				Name:        "training_end_time",
				Description: "A timestamp that shows when the training job ended.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.TrainingJob.TrainingEndTime")},
			{
				Name:        "training_start_time",
				Description: "Indicates the time when the training job starts on training instances.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.TrainingJob.TrainingStartTime")},
			{
				Name:        "training_time_in_seconds",
				Description: "The training time in seconds.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TrainingJob.TrainingTimeInSeconds")},
			{
				Name:        "tuning_job_arn",
				Description: "The Amazon Resource Name (ARN) of the associated hyperparameter tuning job if the training job was launched by a hyperparameter tuning job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrainingJob.TuningJobArn")},
			{
				Name:        "algorithm_specification",
				Description: "Information about the algorithm used for training, and algorithm metadata.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.AlgorithmSpecification")},
			{
				Name:        "checkpoint_config",
				Description: "Contains information about the output location for managed spot training checkpoint data.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.CheckpointConfig")},
			{
				Name:        "debug_hook_config",
				Description: "Configuration information for the Debugger hook parameters, metric and tensor collections, and storage paths.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.DebugHookConfig")},
			{
				Name:        "debug_rule_configurations",
				Description: "Configuration information for Debugger rules for debugging output tensors.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.DebugRuleConfigurations")},
			{
				Name:        "debug_rule_evaluation_statuses",
				Description: "Evaluation status of Debugger rules for debugging on a training job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.DebugRuleEvaluationStatuses")},
			{
				Name:        "environment",
				Description: "The environment variables to set in the Docker container.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.Environment")},
			{
				Name:        "experiment_config",
				Description: "Associates a SageMaker job as a trial component with an experiment and trial.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.ExperimentConfig")},
			{
				Name:        "final_metric_data_list",
				Description: "A collection of MetricData objects that specify the names, values, and dates and times that the training algorithm emitted to Amazon CloudWatch.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.FinalMetricDataList")},
			{
				Name:        "hyper_parameters",
				Description: "Algorithm-specific parameters.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.HyperParameters")},
			{
				Name:        "input_data_config",
				Description: "An array of Channel objects that describes each data input channel.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.InputDataConfig")},
			{
				Name:        "model_artifacts",
				Description: "Information about the Amazon S3 location that is configured for storing model artifacts.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.ModelArtifacts")},
			{
				Name:        "output_data_config",
				Description: "The S3 path where model artifacts that you configured when creating the job are stored.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.OutputDataConfig")},
			{
				Name:        "profiler_config",
				Description: "Configuration information for Debugger system monitoring,framework profiling and storage paths.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.ProfilerConfig")},
			{
				Name:        "profiler_rule_configurations",
				Description: "Configuration information for Debugger rules for profiling system and framework metrics.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.ProfilerRuleConfigurations")},
			{
				Name:        "profiler_rule_evaluation_statuses",
				Description: "Evaluation status of Debugger rules for profiling on a training job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.ProfilerRuleEvaluationStatuses")},
			{
				Name:        "resource_config",
				Description: "Resources, including ML compute instances and ML storage volumes, that are configured for model training.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.ResourceConfig")},
			{
				Name:        "secondary_status_transitions",
				Description: "A history of all of the secondary statuses that the training job has transitioned through.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.SecondaryStatusTransitions")},
			{
				Name:        "stopping_condition",
				Description: "Specifies a limit to how long a model training job can run.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.StoppingCondition")},
			{
				Name:        "tensor_board_output_config",
				Description: "Configuration of storage locations for the Debugger TensorBoard output data.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.TensorBoardOutputConfig")},
			{
				Name:        "vpc_config",
				Description: "A VpcConfig object that specifies the VPC that this training job has access to.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.VpcConfig")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the training job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrainingJob.TrainingJobName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(sageMakerTrainingJobTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrainingJob.TrainingJobArn").Transform(arnToAkas),
			},
		}),
	}
}

func sageMakerTrainingJobTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.SageMakerTrainingJob).Description.Tags

	if tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}

	return nil, nil
}
