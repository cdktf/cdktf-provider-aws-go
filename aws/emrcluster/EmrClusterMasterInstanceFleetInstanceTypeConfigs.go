// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcluster


type EmrClusterMasterInstanceFleetInstanceTypeConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/emr_cluster#instance_type EmrCluster#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/emr_cluster#bid_price EmrCluster#bid_price}.
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/emr_cluster#bid_price_as_percentage_of_on_demand_price EmrCluster#bid_price_as_percentage_of_on_demand_price}.
	BidPriceAsPercentageOfOnDemandPrice *float64 `field:"optional" json:"bidPriceAsPercentageOfOnDemandPrice" yaml:"bidPriceAsPercentageOfOnDemandPrice"`
	// configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/emr_cluster#configurations EmrCluster#configurations}
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// ebs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/emr_cluster#ebs_config EmrCluster#ebs_config}
	EbsConfig interface{} `field:"optional" json:"ebsConfig" yaml:"ebsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/emr_cluster#weighted_capacity EmrCluster#weighted_capacity}.
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

