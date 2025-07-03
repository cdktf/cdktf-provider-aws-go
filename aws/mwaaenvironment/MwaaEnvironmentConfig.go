// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwaaenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwaaEnvironmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#dag_s3_path MwaaEnvironment#dag_s3_path}.
	DagS3Path *string `field:"required" json:"dagS3Path" yaml:"dagS3Path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#execution_role_arn MwaaEnvironment#execution_role_arn}.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#name MwaaEnvironment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#network_configuration MwaaEnvironment#network_configuration}
	NetworkConfiguration *MwaaEnvironmentNetworkConfiguration `field:"required" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#source_bucket_arn MwaaEnvironment#source_bucket_arn}.
	SourceBucketArn *string `field:"required" json:"sourceBucketArn" yaml:"sourceBucketArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#airflow_configuration_options MwaaEnvironment#airflow_configuration_options}.
	AirflowConfigurationOptions *map[string]*string `field:"optional" json:"airflowConfigurationOptions" yaml:"airflowConfigurationOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#airflow_version MwaaEnvironment#airflow_version}.
	AirflowVersion *string `field:"optional" json:"airflowVersion" yaml:"airflowVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#endpoint_management MwaaEnvironment#endpoint_management}.
	EndpointManagement *string `field:"optional" json:"endpointManagement" yaml:"endpointManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#environment_class MwaaEnvironment#environment_class}.
	EnvironmentClass *string `field:"optional" json:"environmentClass" yaml:"environmentClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#id MwaaEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#kms_key MwaaEnvironment#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#logging_configuration MwaaEnvironment#logging_configuration}
	LoggingConfiguration *MwaaEnvironmentLoggingConfiguration `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#max_webservers MwaaEnvironment#max_webservers}.
	MaxWebservers *float64 `field:"optional" json:"maxWebservers" yaml:"maxWebservers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#max_workers MwaaEnvironment#max_workers}.
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#min_webservers MwaaEnvironment#min_webservers}.
	MinWebservers *float64 `field:"optional" json:"minWebservers" yaml:"minWebservers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#min_workers MwaaEnvironment#min_workers}.
	MinWorkers *float64 `field:"optional" json:"minWorkers" yaml:"minWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#plugins_s3_object_version MwaaEnvironment#plugins_s3_object_version}.
	PluginsS3ObjectVersion *string `field:"optional" json:"pluginsS3ObjectVersion" yaml:"pluginsS3ObjectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#plugins_s3_path MwaaEnvironment#plugins_s3_path}.
	PluginsS3Path *string `field:"optional" json:"pluginsS3Path" yaml:"pluginsS3Path"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#region MwaaEnvironment#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#requirements_s3_object_version MwaaEnvironment#requirements_s3_object_version}.
	RequirementsS3ObjectVersion *string `field:"optional" json:"requirementsS3ObjectVersion" yaml:"requirementsS3ObjectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#requirements_s3_path MwaaEnvironment#requirements_s3_path}.
	RequirementsS3Path *string `field:"optional" json:"requirementsS3Path" yaml:"requirementsS3Path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#schedulers MwaaEnvironment#schedulers}.
	Schedulers *float64 `field:"optional" json:"schedulers" yaml:"schedulers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#startup_script_s3_object_version MwaaEnvironment#startup_script_s3_object_version}.
	StartupScriptS3ObjectVersion *string `field:"optional" json:"startupScriptS3ObjectVersion" yaml:"startupScriptS3ObjectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#startup_script_s3_path MwaaEnvironment#startup_script_s3_path}.
	StartupScriptS3Path *string `field:"optional" json:"startupScriptS3Path" yaml:"startupScriptS3Path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#tags MwaaEnvironment#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#tags_all MwaaEnvironment#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#timeouts MwaaEnvironment#timeouts}
	Timeouts *MwaaEnvironmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#webserver_access_mode MwaaEnvironment#webserver_access_mode}.
	WebserverAccessMode *string `field:"optional" json:"webserverAccessMode" yaml:"webserverAccessMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/mwaa_environment#weekly_maintenance_window_start MwaaEnvironment#weekly_maintenance_window_start}.
	WeeklyMaintenanceWindowStart *string `field:"optional" json:"weeklyMaintenanceWindowStart" yaml:"weeklyMaintenanceWindowStart"`
}

