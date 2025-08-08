// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrinstancefleet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrInstanceFleetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/emr_instance_fleet#cluster_id EmrInstanceFleet#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/emr_instance_fleet#id EmrInstanceFleet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instance_type_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/emr_instance_fleet#instance_type_configs EmrInstanceFleet#instance_type_configs}
	InstanceTypeConfigs interface{} `field:"optional" json:"instanceTypeConfigs" yaml:"instanceTypeConfigs"`
	// launch_specifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/emr_instance_fleet#launch_specifications EmrInstanceFleet#launch_specifications}
	LaunchSpecifications *EmrInstanceFleetLaunchSpecifications `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/emr_instance_fleet#name EmrInstanceFleet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/emr_instance_fleet#region EmrInstanceFleet#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/emr_instance_fleet#target_on_demand_capacity EmrInstanceFleet#target_on_demand_capacity}.
	TargetOnDemandCapacity *float64 `field:"optional" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/emr_instance_fleet#target_spot_capacity EmrInstanceFleet#target_spot_capacity}.
	TargetSpotCapacity *float64 `field:"optional" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}

