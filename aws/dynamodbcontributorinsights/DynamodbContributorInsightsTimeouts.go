// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbcontributorinsights


type DynamodbContributorInsightsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/dynamodb_contributor_insights#create DynamodbContributorInsights#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/dynamodb_contributor_insights#delete DynamodbContributorInsights#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

