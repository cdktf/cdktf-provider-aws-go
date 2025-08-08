// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerfeaturegroup


type SagemakerFeatureGroupOnlineStoreConfigTtlDuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/sagemaker_feature_group#unit SagemakerFeatureGroup#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/sagemaker_feature_group#value SagemakerFeatureGroup#value}.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

