// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccesstrustprovider


type VerifiedaccessTrustProviderTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/verifiedaccess_trust_provider#create VerifiedaccessTrustProvider#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/verifiedaccess_trust_provider#delete VerifiedaccessTrustProvider#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/verifiedaccess_trust_provider#update VerifiedaccessTrustProvider#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

