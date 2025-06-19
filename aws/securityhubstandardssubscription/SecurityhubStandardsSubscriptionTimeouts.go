// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubstandardssubscription


type SecurityhubStandardsSubscriptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/securityhub_standards_subscription#create SecurityhubStandardsSubscription#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/securityhub_standards_subscription#delete SecurityhubStandardsSubscription#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

