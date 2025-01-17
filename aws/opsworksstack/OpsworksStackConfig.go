// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksstack

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpsworksStackConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#default_instance_profile_arn OpsworksStack#default_instance_profile_arn}.
	DefaultInstanceProfileArn *string `field:"required" json:"defaultInstanceProfileArn" yaml:"defaultInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#name OpsworksStack#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#region OpsworksStack#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#service_role_arn OpsworksStack#service_role_arn}.
	ServiceRoleArn *string `field:"required" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#agent_version OpsworksStack#agent_version}.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#berkshelf_version OpsworksStack#berkshelf_version}.
	BerkshelfVersion *string `field:"optional" json:"berkshelfVersion" yaml:"berkshelfVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#color OpsworksStack#color}.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#configuration_manager_name OpsworksStack#configuration_manager_name}.
	ConfigurationManagerName *string `field:"optional" json:"configurationManagerName" yaml:"configurationManagerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#configuration_manager_version OpsworksStack#configuration_manager_version}.
	ConfigurationManagerVersion *string `field:"optional" json:"configurationManagerVersion" yaml:"configurationManagerVersion"`
	// custom_cookbooks_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#custom_cookbooks_source OpsworksStack#custom_cookbooks_source}
	CustomCookbooksSource *OpsworksStackCustomCookbooksSource `field:"optional" json:"customCookbooksSource" yaml:"customCookbooksSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#custom_json OpsworksStack#custom_json}.
	CustomJson *string `field:"optional" json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#default_availability_zone OpsworksStack#default_availability_zone}.
	DefaultAvailabilityZone *string `field:"optional" json:"defaultAvailabilityZone" yaml:"defaultAvailabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#default_os OpsworksStack#default_os}.
	DefaultOs *string `field:"optional" json:"defaultOs" yaml:"defaultOs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#default_root_device_type OpsworksStack#default_root_device_type}.
	DefaultRootDeviceType *string `field:"optional" json:"defaultRootDeviceType" yaml:"defaultRootDeviceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#default_ssh_key_name OpsworksStack#default_ssh_key_name}.
	DefaultSshKeyName *string `field:"optional" json:"defaultSshKeyName" yaml:"defaultSshKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#default_subnet_id OpsworksStack#default_subnet_id}.
	DefaultSubnetId *string `field:"optional" json:"defaultSubnetId" yaml:"defaultSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#hostname_theme OpsworksStack#hostname_theme}.
	HostnameTheme *string `field:"optional" json:"hostnameTheme" yaml:"hostnameTheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#id OpsworksStack#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#manage_berkshelf OpsworksStack#manage_berkshelf}.
	ManageBerkshelf interface{} `field:"optional" json:"manageBerkshelf" yaml:"manageBerkshelf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#tags OpsworksStack#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#tags_all OpsworksStack#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#timeouts OpsworksStack#timeouts}
	Timeouts *OpsworksStackTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#use_custom_cookbooks OpsworksStack#use_custom_cookbooks}.
	UseCustomCookbooks interface{} `field:"optional" json:"useCustomCookbooks" yaml:"useCustomCookbooks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#use_opsworks_security_groups OpsworksStack#use_opsworks_security_groups}.
	UseOpsworksSecurityGroups interface{} `field:"optional" json:"useOpsworksSecurityGroups" yaml:"useOpsworksSecurityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opsworks_stack#vpc_id OpsworksStack#vpc_id}.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

