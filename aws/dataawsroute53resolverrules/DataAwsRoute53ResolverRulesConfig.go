// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsroute53resolverrules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsRoute53ResolverRulesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/route53_resolver_rules#id DataAwsRoute53ResolverRules#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/route53_resolver_rules#name_regex DataAwsRoute53ResolverRules#name_regex}.
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/route53_resolver_rules#owner_id DataAwsRoute53ResolverRules#owner_id}.
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/route53_resolver_rules#resolver_endpoint_id DataAwsRoute53ResolverRules#resolver_endpoint_id}.
	ResolverEndpointId *string `field:"optional" json:"resolverEndpointId" yaml:"resolverEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/route53_resolver_rules#rule_type DataAwsRoute53ResolverRules#rule_type}.
	RuleType *string `field:"optional" json:"ruleType" yaml:"ruleType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/route53_resolver_rules#share_status DataAwsRoute53ResolverRules#share_status}.
	ShareStatus *string `field:"optional" json:"shareStatus" yaml:"shareStatus"`
}

