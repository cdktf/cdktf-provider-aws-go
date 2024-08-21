// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcendpointserviceprivatednsverification


type VpcEndpointServicePrivateDnsVerificationTimeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/vpc_endpoint_service_private_dns_verification#create VpcEndpointServicePrivateDnsVerification#create}
	Create *string `field:"optional" json:"create" yaml:"create"`
}

