// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsimagebuilderinfrastructureconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsImagebuilderInfrastructureConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/data-sources/imagebuilder_infrastructure_configuration#arn DataAwsImagebuilderInfrastructureConfiguration#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/data-sources/imagebuilder_infrastructure_configuration#id DataAwsImagebuilderInfrastructureConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/data-sources/imagebuilder_infrastructure_configuration#resource_tags DataAwsImagebuilderInfrastructureConfiguration#resource_tags}.
	ResourceTags *map[string]*string `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/data-sources/imagebuilder_infrastructure_configuration#tags DataAwsImagebuilderInfrastructureConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

