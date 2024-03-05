// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment


type BatchComputeEnvironmentComputeResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#max_vcpus BatchComputeEnvironment#max_vcpus}.
	MaxVcpus *float64 `field:"required" json:"maxVcpus" yaml:"maxVcpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#subnets BatchComputeEnvironment#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#type BatchComputeEnvironment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#allocation_strategy BatchComputeEnvironment#allocation_strategy}.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#bid_percentage BatchComputeEnvironment#bid_percentage}.
	BidPercentage *float64 `field:"optional" json:"bidPercentage" yaml:"bidPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#desired_vcpus BatchComputeEnvironment#desired_vcpus}.
	DesiredVcpus *float64 `field:"optional" json:"desiredVcpus" yaml:"desiredVcpus"`
	// ec2_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#ec2_configuration BatchComputeEnvironment#ec2_configuration}
	Ec2Configuration interface{} `field:"optional" json:"ec2Configuration" yaml:"ec2Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#ec2_key_pair BatchComputeEnvironment#ec2_key_pair}.
	Ec2KeyPair *string `field:"optional" json:"ec2KeyPair" yaml:"ec2KeyPair"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#image_id BatchComputeEnvironment#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#instance_role BatchComputeEnvironment#instance_role}.
	InstanceRole *string `field:"optional" json:"instanceRole" yaml:"instanceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#instance_type BatchComputeEnvironment#instance_type}.
	InstanceType *[]*string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#launch_template BatchComputeEnvironment#launch_template}
	LaunchTemplate *BatchComputeEnvironmentComputeResourcesLaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#min_vcpus BatchComputeEnvironment#min_vcpus}.
	MinVcpus *float64 `field:"optional" json:"minVcpus" yaml:"minVcpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#placement_group BatchComputeEnvironment#placement_group}.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#security_group_ids BatchComputeEnvironment#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#spot_iam_fleet_role BatchComputeEnvironment#spot_iam_fleet_role}.
	SpotIamFleetRole *string `field:"optional" json:"spotIamFleetRole" yaml:"spotIamFleetRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/batch_compute_environment#tags BatchComputeEnvironment#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

