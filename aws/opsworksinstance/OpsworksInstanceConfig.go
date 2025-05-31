// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpsworksInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#layer_ids OpsworksInstance#layer_ids}.
	LayerIds *[]*string `field:"required" json:"layerIds" yaml:"layerIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#stack_id OpsworksInstance#stack_id}.
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#agent_version OpsworksInstance#agent_version}.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#ami_id OpsworksInstance#ami_id}.
	AmiId *string `field:"optional" json:"amiId" yaml:"amiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#architecture OpsworksInstance#architecture}.
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#auto_scaling_type OpsworksInstance#auto_scaling_type}.
	AutoScalingType *string `field:"optional" json:"autoScalingType" yaml:"autoScalingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#availability_zone OpsworksInstance#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#created_at OpsworksInstance#created_at}.
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#delete_ebs OpsworksInstance#delete_ebs}.
	DeleteEbs interface{} `field:"optional" json:"deleteEbs" yaml:"deleteEbs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#delete_eip OpsworksInstance#delete_eip}.
	DeleteEip interface{} `field:"optional" json:"deleteEip" yaml:"deleteEip"`
	// ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#ebs_block_device OpsworksInstance#ebs_block_device}
	EbsBlockDevice interface{} `field:"optional" json:"ebsBlockDevice" yaml:"ebsBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#ebs_optimized OpsworksInstance#ebs_optimized}.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#ecs_cluster_arn OpsworksInstance#ecs_cluster_arn}.
	EcsClusterArn *string `field:"optional" json:"ecsClusterArn" yaml:"ecsClusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#elastic_ip OpsworksInstance#elastic_ip}.
	ElasticIp *string `field:"optional" json:"elasticIp" yaml:"elasticIp"`
	// ephemeral_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#ephemeral_block_device OpsworksInstance#ephemeral_block_device}
	EphemeralBlockDevice interface{} `field:"optional" json:"ephemeralBlockDevice" yaml:"ephemeralBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#hostname OpsworksInstance#hostname}.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#id OpsworksInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#infrastructure_class OpsworksInstance#infrastructure_class}.
	InfrastructureClass *string `field:"optional" json:"infrastructureClass" yaml:"infrastructureClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#install_updates_on_boot OpsworksInstance#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `field:"optional" json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#instance_profile_arn OpsworksInstance#instance_profile_arn}.
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#instance_type OpsworksInstance#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#os OpsworksInstance#os}.
	Os *string `field:"optional" json:"os" yaml:"os"`
	// root_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#root_block_device OpsworksInstance#root_block_device}
	RootBlockDevice interface{} `field:"optional" json:"rootBlockDevice" yaml:"rootBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#root_device_type OpsworksInstance#root_device_type}.
	RootDeviceType *string `field:"optional" json:"rootDeviceType" yaml:"rootDeviceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#security_group_ids OpsworksInstance#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#ssh_key_name OpsworksInstance#ssh_key_name}.
	SshKeyName *string `field:"optional" json:"sshKeyName" yaml:"sshKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#state OpsworksInstance#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#status OpsworksInstance#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#subnet_id OpsworksInstance#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#tenancy OpsworksInstance#tenancy}.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#timeouts OpsworksInstance#timeouts}
	Timeouts *OpsworksInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_instance#virtualization_type OpsworksInstance#virtualization_type}.
	VirtualizationType *string `field:"optional" json:"virtualizationType" yaml:"virtualizationType"`
}

