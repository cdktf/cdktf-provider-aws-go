// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksnodegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EksNodeGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#cluster_name EksNodeGroup#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#node_role_arn EksNodeGroup#node_role_arn}.
	NodeRoleArn *string `field:"required" json:"nodeRoleArn" yaml:"nodeRoleArn"`
	// scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#scaling_config EksNodeGroup#scaling_config}
	ScalingConfig *EksNodeGroupScalingConfig `field:"required" json:"scalingConfig" yaml:"scalingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#subnet_ids EksNodeGroup#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#ami_type EksNodeGroup#ami_type}.
	AmiType *string `field:"optional" json:"amiType" yaml:"amiType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#capacity_type EksNodeGroup#capacity_type}.
	CapacityType *string `field:"optional" json:"capacityType" yaml:"capacityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#disk_size EksNodeGroup#disk_size}.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#force_update_version EksNodeGroup#force_update_version}.
	ForceUpdateVersion interface{} `field:"optional" json:"forceUpdateVersion" yaml:"forceUpdateVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#id EksNodeGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#instance_types EksNodeGroup#instance_types}.
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#labels EksNodeGroup#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#launch_template EksNodeGroup#launch_template}
	LaunchTemplate *EksNodeGroupLaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#node_group_name EksNodeGroup#node_group_name}.
	NodeGroupName *string `field:"optional" json:"nodeGroupName" yaml:"nodeGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#node_group_name_prefix EksNodeGroup#node_group_name_prefix}.
	NodeGroupNamePrefix *string `field:"optional" json:"nodeGroupNamePrefix" yaml:"nodeGroupNamePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#release_version EksNodeGroup#release_version}.
	ReleaseVersion *string `field:"optional" json:"releaseVersion" yaml:"releaseVersion"`
	// remote_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#remote_access EksNodeGroup#remote_access}
	RemoteAccess *EksNodeGroupRemoteAccess `field:"optional" json:"remoteAccess" yaml:"remoteAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#tags EksNodeGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#tags_all EksNodeGroup#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// taint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#taint EksNodeGroup#taint}
	Taint interface{} `field:"optional" json:"taint" yaml:"taint"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#timeouts EksNodeGroup#timeouts}
	Timeouts *EksNodeGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// update_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#update_config EksNodeGroup#update_config}
	UpdateConfig *EksNodeGroupUpdateConfig `field:"optional" json:"updateConfig" yaml:"updateConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/eks_node_group#version EksNodeGroup#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

