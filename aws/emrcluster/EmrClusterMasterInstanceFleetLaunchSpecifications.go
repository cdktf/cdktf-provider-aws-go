// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcluster


type EmrClusterMasterInstanceFleetLaunchSpecifications struct {
	// on_demand_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/emr_cluster#on_demand_specification EmrCluster#on_demand_specification}
	OnDemandSpecification interface{} `field:"optional" json:"onDemandSpecification" yaml:"onDemandSpecification"`
	// spot_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/emr_cluster#spot_specification EmrCluster#spot_specification}
	SpotSpecification interface{} `field:"optional" json:"spotSpecification" yaml:"spotSpecification"`
}

