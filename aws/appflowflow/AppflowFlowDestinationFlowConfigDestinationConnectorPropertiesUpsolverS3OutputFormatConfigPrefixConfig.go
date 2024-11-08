// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/appflow_flow#prefix_type AppflowFlow#prefix_type}.
	PrefixType *string `field:"required" json:"prefixType" yaml:"prefixType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/appflow_flow#prefix_format AppflowFlow#prefix_format}.
	PrefixFormat *string `field:"optional" json:"prefixFormat" yaml:"prefixFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/appflow_flow#prefix_hierarchy AppflowFlow#prefix_hierarchy}.
	PrefixHierarchy *[]*string `field:"optional" json:"prefixHierarchy" yaml:"prefixHierarchy"`
}

