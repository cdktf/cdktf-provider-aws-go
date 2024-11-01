// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instance


type InstanceCreditSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/instance#cpu_credits Instance#cpu_credits}.
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
}

