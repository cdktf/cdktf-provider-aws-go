// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package amplifydomainassociation


type AmplifyDomainAssociationCertificateSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/amplify_domain_association#type AmplifyDomainAssociation#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/amplify_domain_association#custom_certificate_arn AmplifyDomainAssociation#custom_certificate_arn}.
	CustomCertificateArn *string `field:"optional" json:"customCertificateArn" yaml:"customCertificateArn"`
}

