// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codecatalystdevenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodecatalystDevEnvironmentConfig struct {
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
	// ides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/codecatalyst_dev_environment#ides CodecatalystDevEnvironment#ides}
	Ides *CodecatalystDevEnvironmentIdes `field:"required" json:"ides" yaml:"ides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/codecatalyst_dev_environment#instance_type CodecatalystDevEnvironment#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// persistent_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/codecatalyst_dev_environment#persistent_storage CodecatalystDevEnvironment#persistent_storage}
	PersistentStorage *CodecatalystDevEnvironmentPersistentStorage `field:"required" json:"persistentStorage" yaml:"persistentStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/codecatalyst_dev_environment#project_name CodecatalystDevEnvironment#project_name}.
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/codecatalyst_dev_environment#space_name CodecatalystDevEnvironment#space_name}.
	SpaceName *string `field:"required" json:"spaceName" yaml:"spaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/codecatalyst_dev_environment#alias CodecatalystDevEnvironment#alias}.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/codecatalyst_dev_environment#id CodecatalystDevEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/codecatalyst_dev_environment#inactivity_timeout_minutes CodecatalystDevEnvironment#inactivity_timeout_minutes}.
	InactivityTimeoutMinutes *float64 `field:"optional" json:"inactivityTimeoutMinutes" yaml:"inactivityTimeoutMinutes"`
	// repositories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/codecatalyst_dev_environment#repositories CodecatalystDevEnvironment#repositories}
	Repositories interface{} `field:"optional" json:"repositories" yaml:"repositories"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/codecatalyst_dev_environment#timeouts CodecatalystDevEnvironment#timeouts}
	Timeouts *CodecatalystDevEnvironmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

