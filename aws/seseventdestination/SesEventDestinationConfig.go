// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package seseventdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SesEventDestinationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/ses_event_destination#configuration_set_name SesEventDestination#configuration_set_name}.
	ConfigurationSetName *string `field:"required" json:"configurationSetName" yaml:"configurationSetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/ses_event_destination#matching_types SesEventDestination#matching_types}.
	MatchingTypes *[]*string `field:"required" json:"matchingTypes" yaml:"matchingTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/ses_event_destination#name SesEventDestination#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// cloudwatch_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/ses_event_destination#cloudwatch_destination SesEventDestination#cloudwatch_destination}
	CloudwatchDestination interface{} `field:"optional" json:"cloudwatchDestination" yaml:"cloudwatchDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/ses_event_destination#enabled SesEventDestination#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/ses_event_destination#id SesEventDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kinesis_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/ses_event_destination#kinesis_destination SesEventDestination#kinesis_destination}
	KinesisDestination *SesEventDestinationKinesisDestination `field:"optional" json:"kinesisDestination" yaml:"kinesisDestination"`
	// sns_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/ses_event_destination#sns_destination SesEventDestination#sns_destination}
	SnsDestination *SesEventDestinationSnsDestination `field:"optional" json:"snsDestination" yaml:"snsDestination"`
}

