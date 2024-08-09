// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EksClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#name EksCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#role_arn EksCluster#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#vpc_config EksCluster#vpc_config}
	VpcConfig *EksClusterVpcConfig `field:"required" json:"vpcConfig" yaml:"vpcConfig"`
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#access_config EksCluster#access_config}
	AccessConfig *EksClusterAccessConfig `field:"optional" json:"accessConfig" yaml:"accessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#bootstrap_self_managed_addons EksCluster#bootstrap_self_managed_addons}.
	BootstrapSelfManagedAddons interface{} `field:"optional" json:"bootstrapSelfManagedAddons" yaml:"bootstrapSelfManagedAddons"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#enabled_cluster_log_types EksCluster#enabled_cluster_log_types}.
	EnabledClusterLogTypes *[]*string `field:"optional" json:"enabledClusterLogTypes" yaml:"enabledClusterLogTypes"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#encryption_config EksCluster#encryption_config}
	EncryptionConfig *EksClusterEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#id EksCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kubernetes_network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#kubernetes_network_config EksCluster#kubernetes_network_config}
	KubernetesNetworkConfig *EksClusterKubernetesNetworkConfig `field:"optional" json:"kubernetesNetworkConfig" yaml:"kubernetesNetworkConfig"`
	// outpost_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#outpost_config EksCluster#outpost_config}
	OutpostConfig *EksClusterOutpostConfig `field:"optional" json:"outpostConfig" yaml:"outpostConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#tags EksCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#tags_all EksCluster#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#timeouts EksCluster#timeouts}
	Timeouts *EksClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#upgrade_policy EksCluster#upgrade_policy}
	UpgradePolicy *EksClusterUpgradePolicy `field:"optional" json:"upgradePolicy" yaml:"upgradePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_cluster#version EksCluster#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

