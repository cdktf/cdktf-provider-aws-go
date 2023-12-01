// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate


type FisExperimentTemplateLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/fis_experiment_template#log_schema_version FisExperimentTemplate#log_schema_version}.
	LogSchemaVersion *float64 `field:"required" json:"logSchemaVersion" yaml:"logSchemaVersion"`
	// cloudwatch_logs_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/fis_experiment_template#cloudwatch_logs_configuration FisExperimentTemplate#cloudwatch_logs_configuration}
	CloudwatchLogsConfiguration *FisExperimentTemplateLogConfigurationCloudwatchLogsConfiguration `field:"optional" json:"cloudwatchLogsConfiguration" yaml:"cloudwatchLogsConfiguration"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/fis_experiment_template#s3_configuration FisExperimentTemplate#s3_configuration}
	S3Configuration *FisExperimentTemplateLogConfigurationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

