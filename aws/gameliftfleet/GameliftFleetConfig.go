// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GameliftFleetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#ec2_instance_type GameliftFleet#ec2_instance_type}.
	Ec2InstanceType *string `field:"required" json:"ec2InstanceType" yaml:"ec2InstanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#name GameliftFleet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#build_id GameliftFleet#build_id}.
	BuildId *string `field:"optional" json:"buildId" yaml:"buildId"`
	// certificate_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#certificate_configuration GameliftFleet#certificate_configuration}
	CertificateConfiguration *GameliftFleetCertificateConfiguration `field:"optional" json:"certificateConfiguration" yaml:"certificateConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#description GameliftFleet#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// ec2_inbound_permission block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#ec2_inbound_permission GameliftFleet#ec2_inbound_permission}
	Ec2InboundPermission interface{} `field:"optional" json:"ec2InboundPermission" yaml:"ec2InboundPermission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#fleet_type GameliftFleet#fleet_type}.
	FleetType *string `field:"optional" json:"fleetType" yaml:"fleetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#id GameliftFleet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#instance_role_arn GameliftFleet#instance_role_arn}.
	InstanceRoleArn *string `field:"optional" json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#metric_groups GameliftFleet#metric_groups}.
	MetricGroups *[]*string `field:"optional" json:"metricGroups" yaml:"metricGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#new_game_session_protection_policy GameliftFleet#new_game_session_protection_policy}.
	NewGameSessionProtectionPolicy *string `field:"optional" json:"newGameSessionProtectionPolicy" yaml:"newGameSessionProtectionPolicy"`
	// resource_creation_limit_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#resource_creation_limit_policy GameliftFleet#resource_creation_limit_policy}
	ResourceCreationLimitPolicy *GameliftFleetResourceCreationLimitPolicy `field:"optional" json:"resourceCreationLimitPolicy" yaml:"resourceCreationLimitPolicy"`
	// runtime_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#runtime_configuration GameliftFleet#runtime_configuration}
	RuntimeConfiguration *GameliftFleetRuntimeConfiguration `field:"optional" json:"runtimeConfiguration" yaml:"runtimeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#script_id GameliftFleet#script_id}.
	ScriptId *string `field:"optional" json:"scriptId" yaml:"scriptId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#tags GameliftFleet#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#tags_all GameliftFleet#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/gamelift_fleet#timeouts GameliftFleet#timeouts}
	Timeouts *GameliftFleetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

