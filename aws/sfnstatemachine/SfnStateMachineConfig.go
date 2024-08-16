// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sfnstatemachine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SfnStateMachineConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/sfn_state_machine#definition SfnStateMachine#definition}.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/sfn_state_machine#role_arn SfnStateMachine#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/sfn_state_machine#encryption_configuration SfnStateMachine#encryption_configuration}
	EncryptionConfiguration *SfnStateMachineEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/sfn_state_machine#id SfnStateMachine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/sfn_state_machine#logging_configuration SfnStateMachine#logging_configuration}
	LoggingConfiguration *SfnStateMachineLoggingConfiguration `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/sfn_state_machine#name SfnStateMachine#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/sfn_state_machine#name_prefix SfnStateMachine#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/sfn_state_machine#publish SfnStateMachine#publish}.
	Publish interface{} `field:"optional" json:"publish" yaml:"publish"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/sfn_state_machine#tags SfnStateMachine#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/sfn_state_machine#tags_all SfnStateMachine#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/sfn_state_machine#timeouts SfnStateMachine#timeouts}
	Timeouts *SfnStateMachineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// tracing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/sfn_state_machine#tracing_configuration SfnStateMachine#tracing_configuration}
	TracingConfiguration *SfnStateMachineTracingConfiguration `field:"optional" json:"tracingConfiguration" yaml:"tracingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/sfn_state_machine#type SfnStateMachine#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

