// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2apikey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2ApiKeyConfig struct {
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
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application.
	//
	// Valid values are CLOUDFRONT or REGIONAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/wafv2_api_key#scope Wafv2ApiKey#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// The domains that you want to be able to use the API key with, for example example.com. Maximum of 5 domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/wafv2_api_key#token_domains Wafv2ApiKey#token_domains}
	TokenDomains *[]*string `field:"required" json:"tokenDomains" yaml:"tokenDomains"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/wafv2_api_key#region Wafv2ApiKey#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

