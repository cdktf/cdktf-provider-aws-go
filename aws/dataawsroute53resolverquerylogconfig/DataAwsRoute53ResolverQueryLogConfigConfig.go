// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsroute53resolverquerylogconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsRoute53ResolverQueryLogConfigConfig struct {
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
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/route53_resolver_query_log_config#filter DataAwsRoute53ResolverQueryLogConfig#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/route53_resolver_query_log_config#id DataAwsRoute53ResolverQueryLogConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/route53_resolver_query_log_config#name DataAwsRoute53ResolverQueryLogConfig#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/route53_resolver_query_log_config#region DataAwsRoute53ResolverQueryLogConfig#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/route53_resolver_query_log_config#resolver_query_log_config_id DataAwsRoute53ResolverQueryLogConfig#resolver_query_log_config_id}.
	ResolverQueryLogConfigId *string `field:"optional" json:"resolverQueryLogConfigId" yaml:"resolverQueryLogConfigId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/route53_resolver_query_log_config#tags DataAwsRoute53ResolverQueryLogConfig#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

