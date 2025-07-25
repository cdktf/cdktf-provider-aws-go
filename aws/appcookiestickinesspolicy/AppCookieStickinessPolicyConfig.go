// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appcookiestickinesspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppCookieStickinessPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/app_cookie_stickiness_policy#cookie_name AppCookieStickinessPolicy#cookie_name}.
	CookieName *string `field:"required" json:"cookieName" yaml:"cookieName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/app_cookie_stickiness_policy#lb_port AppCookieStickinessPolicy#lb_port}.
	LbPort *float64 `field:"required" json:"lbPort" yaml:"lbPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/app_cookie_stickiness_policy#load_balancer AppCookieStickinessPolicy#load_balancer}.
	LoadBalancer *string `field:"required" json:"loadBalancer" yaml:"loadBalancer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/app_cookie_stickiness_policy#name AppCookieStickinessPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/app_cookie_stickiness_policy#id AppCookieStickinessPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/app_cookie_stickiness_policy#region AppCookieStickinessPolicy#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

