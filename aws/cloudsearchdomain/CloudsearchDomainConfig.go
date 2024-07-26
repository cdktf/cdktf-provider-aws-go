// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudsearchdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudsearchDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudsearch_domain#name CloudsearchDomain#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// endpoint_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudsearch_domain#endpoint_options CloudsearchDomain#endpoint_options}
	EndpointOptions *CloudsearchDomainEndpointOptions `field:"optional" json:"endpointOptions" yaml:"endpointOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudsearch_domain#id CloudsearchDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// index_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudsearch_domain#index_field CloudsearchDomain#index_field}
	IndexField interface{} `field:"optional" json:"indexField" yaml:"indexField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudsearch_domain#multi_az CloudsearchDomain#multi_az}.
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// scaling_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudsearch_domain#scaling_parameters CloudsearchDomain#scaling_parameters}
	ScalingParameters *CloudsearchDomainScalingParameters `field:"optional" json:"scalingParameters" yaml:"scalingParameters"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudsearch_domain#timeouts CloudsearchDomain#timeouts}
	Timeouts *CloudsearchDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

