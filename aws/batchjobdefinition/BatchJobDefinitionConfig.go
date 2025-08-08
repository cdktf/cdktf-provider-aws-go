// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchJobDefinitionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#name BatchJobDefinition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#type BatchJobDefinition#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#container_properties BatchJobDefinition#container_properties}.
	ContainerProperties *string `field:"optional" json:"containerProperties" yaml:"containerProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#deregister_on_new_revision BatchJobDefinition#deregister_on_new_revision}.
	DeregisterOnNewRevision interface{} `field:"optional" json:"deregisterOnNewRevision" yaml:"deregisterOnNewRevision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#ecs_properties BatchJobDefinition#ecs_properties}.
	EcsProperties *string `field:"optional" json:"ecsProperties" yaml:"ecsProperties"`
	// eks_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#eks_properties BatchJobDefinition#eks_properties}
	EksProperties *BatchJobDefinitionEksProperties `field:"optional" json:"eksProperties" yaml:"eksProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#id BatchJobDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#node_properties BatchJobDefinition#node_properties}.
	NodeProperties *string `field:"optional" json:"nodeProperties" yaml:"nodeProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#parameters BatchJobDefinition#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#platform_capabilities BatchJobDefinition#platform_capabilities}.
	PlatformCapabilities *[]*string `field:"optional" json:"platformCapabilities" yaml:"platformCapabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#propagate_tags BatchJobDefinition#propagate_tags}.
	PropagateTags interface{} `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#region BatchJobDefinition#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// retry_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#retry_strategy BatchJobDefinition#retry_strategy}
	RetryStrategy *BatchJobDefinitionRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#scheduling_priority BatchJobDefinition#scheduling_priority}.
	SchedulingPriority *float64 `field:"optional" json:"schedulingPriority" yaml:"schedulingPriority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#tags BatchJobDefinition#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#tags_all BatchJobDefinition#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/batch_job_definition#timeout BatchJobDefinition#timeout}
	Timeout *BatchJobDefinitionTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}

