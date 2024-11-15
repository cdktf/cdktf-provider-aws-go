// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtask


type SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters struct {
	// cloudwatch_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ssm_maintenance_window_task#cloudwatch_config SsmMaintenanceWindowTask#cloudwatch_config}
	CloudwatchConfig *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig `field:"optional" json:"cloudwatchConfig" yaml:"cloudwatchConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ssm_maintenance_window_task#comment SsmMaintenanceWindowTask#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ssm_maintenance_window_task#document_hash SsmMaintenanceWindowTask#document_hash}.
	DocumentHash *string `field:"optional" json:"documentHash" yaml:"documentHash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ssm_maintenance_window_task#document_hash_type SsmMaintenanceWindowTask#document_hash_type}.
	DocumentHashType *string `field:"optional" json:"documentHashType" yaml:"documentHashType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ssm_maintenance_window_task#document_version SsmMaintenanceWindowTask#document_version}.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ssm_maintenance_window_task#notification_config SsmMaintenanceWindowTask#notification_config}
	NotificationConfig *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ssm_maintenance_window_task#output_s3_bucket SsmMaintenanceWindowTask#output_s3_bucket}.
	OutputS3Bucket *string `field:"optional" json:"outputS3Bucket" yaml:"outputS3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ssm_maintenance_window_task#output_s3_key_prefix SsmMaintenanceWindowTask#output_s3_key_prefix}.
	OutputS3KeyPrefix *string `field:"optional" json:"outputS3KeyPrefix" yaml:"outputS3KeyPrefix"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ssm_maintenance_window_task#parameter SsmMaintenanceWindowTask#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ssm_maintenance_window_task#service_role_arn SsmMaintenanceWindowTask#service_role_arn}.
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ssm_maintenance_window_task#timeout_seconds SsmMaintenanceWindowTask#timeout_seconds}.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

