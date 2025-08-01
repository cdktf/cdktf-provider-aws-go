// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontresponseheaderspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontResponseHeadersPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cloudfront_response_headers_policy#name CloudfrontResponseHeadersPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cloudfront_response_headers_policy#comment CloudfrontResponseHeadersPolicy#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// cors_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cloudfront_response_headers_policy#cors_config CloudfrontResponseHeadersPolicy#cors_config}
	CorsConfig *CloudfrontResponseHeadersPolicyCorsConfig `field:"optional" json:"corsConfig" yaml:"corsConfig"`
	// custom_headers_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cloudfront_response_headers_policy#custom_headers_config CloudfrontResponseHeadersPolicy#custom_headers_config}
	CustomHeadersConfig *CloudfrontResponseHeadersPolicyCustomHeadersConfig `field:"optional" json:"customHeadersConfig" yaml:"customHeadersConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cloudfront_response_headers_policy#id CloudfrontResponseHeadersPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// remove_headers_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cloudfront_response_headers_policy#remove_headers_config CloudfrontResponseHeadersPolicy#remove_headers_config}
	RemoveHeadersConfig *CloudfrontResponseHeadersPolicyRemoveHeadersConfig `field:"optional" json:"removeHeadersConfig" yaml:"removeHeadersConfig"`
	// security_headers_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cloudfront_response_headers_policy#security_headers_config CloudfrontResponseHeadersPolicy#security_headers_config}
	SecurityHeadersConfig *CloudfrontResponseHeadersPolicySecurityHeadersConfig `field:"optional" json:"securityHeadersConfig" yaml:"securityHeadersConfig"`
	// server_timing_headers_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cloudfront_response_headers_policy#server_timing_headers_config CloudfrontResponseHeadersPolicy#server_timing_headers_config}
	ServerTimingHeadersConfig *CloudfrontResponseHeadersPolicyServerTimingHeadersConfig `field:"optional" json:"serverTimingHeadersConfig" yaml:"serverTimingHeadersConfig"`
}

