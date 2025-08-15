// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkgroupConfigurationIdentityCenterConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/athena_workgroup#enable_identity_center AthenaWorkgroup#enable_identity_center}.
	EnableIdentityCenter interface{} `field:"optional" json:"enableIdentityCenter" yaml:"enableIdentityCenter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/athena_workgroup#identity_center_instance_arn AthenaWorkgroup#identity_center_instance_arn}.
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
}

