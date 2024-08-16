// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksaccessentry

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EksAccessEntryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/eks_access_entry#cluster_name EksAccessEntry#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/eks_access_entry#principal_arn EksAccessEntry#principal_arn}.
	PrincipalArn *string `field:"required" json:"principalArn" yaml:"principalArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/eks_access_entry#id EksAccessEntry#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/eks_access_entry#kubernetes_groups EksAccessEntry#kubernetes_groups}.
	KubernetesGroups *[]*string `field:"optional" json:"kubernetesGroups" yaml:"kubernetesGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/eks_access_entry#tags EksAccessEntry#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/eks_access_entry#tags_all EksAccessEntry#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/eks_access_entry#timeouts EksAccessEntry#timeouts}
	Timeouts *EksAccessEntryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/eks_access_entry#type EksAccessEntry#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/eks_access_entry#user_name EksAccessEntry#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

