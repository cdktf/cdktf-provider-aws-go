// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcachepolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontCachePolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cloudfront_cache_policy#name CloudfrontCachePolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// parameters_in_cache_key_and_forwarded_to_origin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cloudfront_cache_policy#parameters_in_cache_key_and_forwarded_to_origin CloudfrontCachePolicy#parameters_in_cache_key_and_forwarded_to_origin}
	ParametersInCacheKeyAndForwardedToOrigin *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin `field:"required" json:"parametersInCacheKeyAndForwardedToOrigin" yaml:"parametersInCacheKeyAndForwardedToOrigin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cloudfront_cache_policy#comment CloudfrontCachePolicy#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cloudfront_cache_policy#default_ttl CloudfrontCachePolicy#default_ttl}.
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cloudfront_cache_policy#id CloudfrontCachePolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cloudfront_cache_policy#max_ttl CloudfrontCachePolicy#max_ttl}.
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cloudfront_cache_policy#min_ttl CloudfrontCachePolicy#min_ttl}.
	MinTtl *float64 `field:"optional" json:"minTtl" yaml:"minTtl"`
}

