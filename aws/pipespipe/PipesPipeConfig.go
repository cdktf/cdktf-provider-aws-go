// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipesPipeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#role_arn PipesPipe#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#source PipesPipe#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#target PipesPipe#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#description PipesPipe#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#desired_state PipesPipe#desired_state}.
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#enrichment PipesPipe#enrichment}.
	Enrichment *string `field:"optional" json:"enrichment" yaml:"enrichment"`
	// enrichment_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#enrichment_parameters PipesPipe#enrichment_parameters}
	EnrichmentParameters *PipesPipeEnrichmentParameters `field:"optional" json:"enrichmentParameters" yaml:"enrichmentParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#id PipesPipe#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#kms_key_identifier PipesPipe#kms_key_identifier}.
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#log_configuration PipesPipe#log_configuration}
	LogConfiguration *PipesPipeLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#name PipesPipe#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#name_prefix PipesPipe#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#region PipesPipe#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// source_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#source_parameters PipesPipe#source_parameters}
	SourceParameters *PipesPipeSourceParameters `field:"optional" json:"sourceParameters" yaml:"sourceParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#tags PipesPipe#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#tags_all PipesPipe#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// target_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#target_parameters PipesPipe#target_parameters}
	TargetParameters *PipesPipeTargetParameters `field:"optional" json:"targetParameters" yaml:"targetParameters"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/pipes_pipe#timeouts PipesPipe#timeouts}
	Timeouts *PipesPipeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

