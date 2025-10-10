// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerfeaturegroup


type SagemakerFeatureGroupThroughputConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/sagemaker_feature_group#provisioned_read_capacity_units SagemakerFeatureGroup#provisioned_read_capacity_units}.
	ProvisionedReadCapacityUnits *float64 `field:"optional" json:"provisionedReadCapacityUnits" yaml:"provisionedReadCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/sagemaker_feature_group#provisioned_write_capacity_units SagemakerFeatureGroup#provisioned_write_capacity_units}.
	ProvisionedWriteCapacityUnits *float64 `field:"optional" json:"provisionedWriteCapacityUnits" yaml:"provisionedWriteCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/sagemaker_feature_group#throughput_mode SagemakerFeatureGroup#throughput_mode}.
	ThroughputMode *string `field:"optional" json:"throughputMode" yaml:"throughputMode"`
}

