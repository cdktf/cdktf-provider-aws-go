// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventtarget


type CloudwatchEventTargetRedshiftTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/cloudwatch_event_target#database CloudwatchEventTarget#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/cloudwatch_event_target#db_user CloudwatchEventTarget#db_user}.
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/cloudwatch_event_target#secrets_manager_arn CloudwatchEventTarget#secrets_manager_arn}.
	SecretsManagerArn *string `field:"optional" json:"secretsManagerArn" yaml:"secretsManagerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/cloudwatch_event_target#sql CloudwatchEventTarget#sql}.
	Sql *string `field:"optional" json:"sql" yaml:"sql"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/cloudwatch_event_target#statement_name CloudwatchEventTarget#statement_name}.
	StatementName *string `field:"optional" json:"statementName" yaml:"statementName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/cloudwatch_event_target#with_event CloudwatchEventTarget#with_event}.
	WithEvent interface{} `field:"optional" json:"withEvent" yaml:"withEvent"`
}

