// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package msksinglescramsecretassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MskSingleScramSecretAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/msk_single_scram_secret_association#cluster_arn MskSingleScramSecretAssociation#cluster_arn}.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/msk_single_scram_secret_association#secret_arn MskSingleScramSecretAssociation#secret_arn}.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}

