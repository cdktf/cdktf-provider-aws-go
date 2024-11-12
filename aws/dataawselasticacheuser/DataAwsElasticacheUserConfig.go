// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawselasticacheuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsElasticacheUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/elasticache_user#user_id DataAwsElasticacheUser#user_id}.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/elasticache_user#access_string DataAwsElasticacheUser#access_string}.
	AccessString *string `field:"optional" json:"accessString" yaml:"accessString"`
	// authentication_mode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/elasticache_user#authentication_mode DataAwsElasticacheUser#authentication_mode}
	AuthenticationMode interface{} `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/elasticache_user#engine DataAwsElasticacheUser#engine}.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/elasticache_user#id DataAwsElasticacheUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/elasticache_user#no_password_required DataAwsElasticacheUser#no_password_required}.
	NoPasswordRequired interface{} `field:"optional" json:"noPasswordRequired" yaml:"noPasswordRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/elasticache_user#passwords DataAwsElasticacheUser#passwords}.
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/elasticache_user#user_name DataAwsElasticacheUser#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

