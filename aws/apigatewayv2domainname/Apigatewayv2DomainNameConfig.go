// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2domainname

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Apigatewayv2DomainNameConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/apigatewayv2_domain_name#domain_name Apigatewayv2DomainName#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// domain_name_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/apigatewayv2_domain_name#domain_name_configuration Apigatewayv2DomainName#domain_name_configuration}
	DomainNameConfiguration *Apigatewayv2DomainNameDomainNameConfiguration `field:"required" json:"domainNameConfiguration" yaml:"domainNameConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/apigatewayv2_domain_name#id Apigatewayv2DomainName#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// mutual_tls_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/apigatewayv2_domain_name#mutual_tls_authentication Apigatewayv2DomainName#mutual_tls_authentication}
	MutualTlsAuthentication *Apigatewayv2DomainNameMutualTlsAuthentication `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/apigatewayv2_domain_name#tags Apigatewayv2DomainName#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/apigatewayv2_domain_name#tags_all Apigatewayv2DomainName#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/apigatewayv2_domain_name#timeouts Apigatewayv2DomainName#timeouts}
	Timeouts *Apigatewayv2DomainNameTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

