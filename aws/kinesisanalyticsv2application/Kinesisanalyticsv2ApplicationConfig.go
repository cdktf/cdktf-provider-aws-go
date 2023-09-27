// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Kinesisanalyticsv2ApplicationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#name Kinesisanalyticsv2Application#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#runtime_environment Kinesisanalyticsv2Application#runtime_environment}.
	RuntimeEnvironment *string `field:"required" json:"runtimeEnvironment" yaml:"runtimeEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#service_execution_role Kinesisanalyticsv2Application#service_execution_role}.
	ServiceExecutionRole *string `field:"required" json:"serviceExecutionRole" yaml:"serviceExecutionRole"`
	// application_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#application_configuration Kinesisanalyticsv2Application#application_configuration}
	ApplicationConfiguration *Kinesisanalyticsv2ApplicationApplicationConfiguration `field:"optional" json:"applicationConfiguration" yaml:"applicationConfiguration"`
	// cloudwatch_logging_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#cloudwatch_logging_options Kinesisanalyticsv2Application#cloudwatch_logging_options}
	CloudwatchLoggingOptions *Kinesisanalyticsv2ApplicationCloudwatchLoggingOptions `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#description Kinesisanalyticsv2Application#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#force_stop Kinesisanalyticsv2Application#force_stop}.
	ForceStop interface{} `field:"optional" json:"forceStop" yaml:"forceStop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#id Kinesisanalyticsv2Application#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#start_application Kinesisanalyticsv2Application#start_application}.
	StartApplication interface{} `field:"optional" json:"startApplication" yaml:"startApplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#tags Kinesisanalyticsv2Application#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#tags_all Kinesisanalyticsv2Application#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#timeouts Kinesisanalyticsv2Application#timeouts}
	Timeouts *Kinesisanalyticsv2ApplicationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

