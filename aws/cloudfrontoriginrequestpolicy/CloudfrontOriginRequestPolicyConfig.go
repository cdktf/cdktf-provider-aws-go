// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontoriginrequestpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontOriginRequestPolicyConfig struct {
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
	// cookies_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/cloudfront_origin_request_policy#cookies_config CloudfrontOriginRequestPolicy#cookies_config}
	CookiesConfig *CloudfrontOriginRequestPolicyCookiesConfig `field:"required" json:"cookiesConfig" yaml:"cookiesConfig"`
	// headers_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/cloudfront_origin_request_policy#headers_config CloudfrontOriginRequestPolicy#headers_config}
	HeadersConfig *CloudfrontOriginRequestPolicyHeadersConfig `field:"required" json:"headersConfig" yaml:"headersConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/cloudfront_origin_request_policy#name CloudfrontOriginRequestPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// query_strings_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/cloudfront_origin_request_policy#query_strings_config CloudfrontOriginRequestPolicy#query_strings_config}
	QueryStringsConfig *CloudfrontOriginRequestPolicyQueryStringsConfig `field:"required" json:"queryStringsConfig" yaml:"queryStringsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/cloudfront_origin_request_policy#comment CloudfrontOriginRequestPolicy#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/cloudfront_origin_request_policy#id CloudfrontOriginRequestPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

