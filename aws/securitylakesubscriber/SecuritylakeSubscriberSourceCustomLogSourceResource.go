// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitylakesubscriber


type SecuritylakeSubscriberSourceCustomLogSourceResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securitylake_subscriber#source_name SecuritylakeSubscriber#source_name}.
	SourceName *string `field:"required" json:"sourceName" yaml:"sourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securitylake_subscriber#source_version SecuritylakeSubscriber#source_version}.
	SourceVersion *string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
}

