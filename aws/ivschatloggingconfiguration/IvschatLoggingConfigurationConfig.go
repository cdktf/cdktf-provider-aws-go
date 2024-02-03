// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ivschatloggingconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IvschatLoggingConfigurationConfig struct {
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
	// destination_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ivschat_logging_configuration#destination_configuration IvschatLoggingConfiguration#destination_configuration}
	DestinationConfiguration *IvschatLoggingConfigurationDestinationConfiguration `field:"optional" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ivschat_logging_configuration#id IvschatLoggingConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ivschat_logging_configuration#name IvschatLoggingConfiguration#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ivschat_logging_configuration#tags IvschatLoggingConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ivschat_logging_configuration#tags_all IvschatLoggingConfiguration#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ivschat_logging_configuration#timeouts IvschatLoggingConfiguration#timeouts}
	Timeouts *IvschatLoggingConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

