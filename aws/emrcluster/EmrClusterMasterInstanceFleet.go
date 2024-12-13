// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcluster


type EmrClusterMasterInstanceFleet struct {
	// instance_type_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/emr_cluster#instance_type_configs EmrCluster#instance_type_configs}
	InstanceTypeConfigs interface{} `field:"optional" json:"instanceTypeConfigs" yaml:"instanceTypeConfigs"`
	// launch_specifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/emr_cluster#launch_specifications EmrCluster#launch_specifications}
	LaunchSpecifications *EmrClusterMasterInstanceFleetLaunchSpecifications `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/emr_cluster#name EmrCluster#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/emr_cluster#target_on_demand_capacity EmrCluster#target_on_demand_capacity}.
	TargetOnDemandCapacity *float64 `field:"optional" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/emr_cluster#target_spot_capacity EmrCluster#target_spot_capacity}.
	TargetSpotCapacity *float64 `field:"optional" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}

