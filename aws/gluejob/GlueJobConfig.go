// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluejob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueJobConfig struct {
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
	// command block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#command GlueJob#command}
	Command *GlueJobCommand `field:"required" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#name GlueJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#role_arn GlueJob#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#connections GlueJob#connections}.
	Connections *[]*string `field:"optional" json:"connections" yaml:"connections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#default_arguments GlueJob#default_arguments}.
	DefaultArguments *map[string]*string `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#description GlueJob#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#execution_class GlueJob#execution_class}.
	ExecutionClass *string `field:"optional" json:"executionClass" yaml:"executionClass"`
	// execution_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#execution_property GlueJob#execution_property}
	ExecutionProperty *GlueJobExecutionProperty `field:"optional" json:"executionProperty" yaml:"executionProperty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#glue_version GlueJob#glue_version}.
	GlueVersion *string `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#id GlueJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#max_capacity GlueJob#max_capacity}.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#max_retries GlueJob#max_retries}.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#non_overridable_arguments GlueJob#non_overridable_arguments}.
	NonOverridableArguments *map[string]*string `field:"optional" json:"nonOverridableArguments" yaml:"nonOverridableArguments"`
	// notification_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#notification_property GlueJob#notification_property}
	NotificationProperty *GlueJobNotificationProperty `field:"optional" json:"notificationProperty" yaml:"notificationProperty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#number_of_workers GlueJob#number_of_workers}.
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#security_configuration GlueJob#security_configuration}.
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#tags GlueJob#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#tags_all GlueJob#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#timeout GlueJob#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/glue_job#worker_type GlueJob#worker_type}.
	WorkerType *string `field:"optional" json:"workerType" yaml:"workerType"`
}

