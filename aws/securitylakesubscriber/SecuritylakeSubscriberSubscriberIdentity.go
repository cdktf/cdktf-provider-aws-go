// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitylakesubscriber


type SecuritylakeSubscriberSubscriberIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/securitylake_subscriber#external_id SecuritylakeSubscriber#external_id}.
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/securitylake_subscriber#principal SecuritylakeSubscriber#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

