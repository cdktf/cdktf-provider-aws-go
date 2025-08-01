// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrpullthroughcacherule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcrPullThroughCacheRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ecr_pull_through_cache_rule#ecr_repository_prefix EcrPullThroughCacheRule#ecr_repository_prefix}.
	EcrRepositoryPrefix *string `field:"required" json:"ecrRepositoryPrefix" yaml:"ecrRepositoryPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ecr_pull_through_cache_rule#upstream_registry_url EcrPullThroughCacheRule#upstream_registry_url}.
	UpstreamRegistryUrl *string `field:"required" json:"upstreamRegistryUrl" yaml:"upstreamRegistryUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ecr_pull_through_cache_rule#credential_arn EcrPullThroughCacheRule#credential_arn}.
	CredentialArn *string `field:"optional" json:"credentialArn" yaml:"credentialArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ecr_pull_through_cache_rule#custom_role_arn EcrPullThroughCacheRule#custom_role_arn}.
	CustomRoleArn *string `field:"optional" json:"customRoleArn" yaml:"customRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ecr_pull_through_cache_rule#id EcrPullThroughCacheRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ecr_pull_through_cache_rule#region EcrPullThroughCacheRule#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ecr_pull_through_cache_rule#upstream_repository_prefix EcrPullThroughCacheRule#upstream_repository_prefix}.
	UpstreamRepositoryPrefix *string `field:"optional" json:"upstreamRepositoryPrefix" yaml:"upstreamRepositoryPrefix"`
}

