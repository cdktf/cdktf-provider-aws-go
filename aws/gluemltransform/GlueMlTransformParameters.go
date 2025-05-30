// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluemltransform


type GlueMlTransformParameters struct {
	// find_matches_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/glue_ml_transform#find_matches_parameters GlueMlTransform#find_matches_parameters}
	FindMatchesParameters *GlueMlTransformParametersFindMatchesParameters `field:"required" json:"findMatchesParameters" yaml:"findMatchesParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/glue_ml_transform#transform_type GlueMlTransform#transform_type}.
	TransformType *string `field:"required" json:"transformType" yaml:"transformType"`
}

