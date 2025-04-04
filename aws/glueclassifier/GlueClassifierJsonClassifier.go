// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package glueclassifier


type GlueClassifierJsonClassifier struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/glue_classifier#json_path GlueClassifier#json_path}.
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
}

