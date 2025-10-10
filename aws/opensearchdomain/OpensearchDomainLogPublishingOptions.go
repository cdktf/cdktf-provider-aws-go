// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainLogPublishingOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/opensearch_domain#cloudwatch_log_group_arn OpensearchDomain#cloudwatch_log_group_arn}.
	CloudwatchLogGroupArn *string `field:"required" json:"cloudwatchLogGroupArn" yaml:"cloudwatchLogGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/opensearch_domain#log_type OpensearchDomain#log_type}.
	LogType *string `field:"required" json:"logType" yaml:"logType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/opensearch_domain#enabled OpensearchDomain#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

