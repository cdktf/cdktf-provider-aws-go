// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotfleetrequest


type SpotFleetRequestLaunchSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#ami SpotFleetRequest#ami}.
	Ami *string `field:"required" json:"ami" yaml:"ami"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#instance_type SpotFleetRequest#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#associate_public_ip_address SpotFleetRequest#associate_public_ip_address}.
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#availability_zone SpotFleetRequest#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#ebs_block_device SpotFleetRequest#ebs_block_device}
	EbsBlockDevice interface{} `field:"optional" json:"ebsBlockDevice" yaml:"ebsBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#ebs_optimized SpotFleetRequest#ebs_optimized}.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// ephemeral_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#ephemeral_block_device SpotFleetRequest#ephemeral_block_device}
	EphemeralBlockDevice interface{} `field:"optional" json:"ephemeralBlockDevice" yaml:"ephemeralBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#iam_instance_profile SpotFleetRequest#iam_instance_profile}.
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#iam_instance_profile_arn SpotFleetRequest#iam_instance_profile_arn}.
	IamInstanceProfileArn *string `field:"optional" json:"iamInstanceProfileArn" yaml:"iamInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#key_name SpotFleetRequest#key_name}.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#monitoring SpotFleetRequest#monitoring}.
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#placement_group SpotFleetRequest#placement_group}.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#placement_tenancy SpotFleetRequest#placement_tenancy}.
	PlacementTenancy *string `field:"optional" json:"placementTenancy" yaml:"placementTenancy"`
	// root_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#root_block_device SpotFleetRequest#root_block_device}
	RootBlockDevice interface{} `field:"optional" json:"rootBlockDevice" yaml:"rootBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#spot_price SpotFleetRequest#spot_price}.
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#subnet_id SpotFleetRequest#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#tags SpotFleetRequest#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#user_data SpotFleetRequest#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#vpc_security_group_ids SpotFleetRequest#vpc_security_group_ids}.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/spot_fleet_request#weighted_capacity SpotFleetRequest#weighted_capacity}.
	WeightedCapacity *string `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

