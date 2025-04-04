// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockprovisionedmodelthroughput


type BedrockProvisionedModelThroughputTimeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/bedrock_provisioned_model_throughput#create BedrockProvisionedModelThroughput#create}
	Create *string `field:"optional" json:"create" yaml:"create"`
}

