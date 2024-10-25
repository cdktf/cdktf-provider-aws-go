// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceConnectConfigurationServiceTls struct {
	// issuer_cert_authority block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/ecs_service#issuer_cert_authority EcsService#issuer_cert_authority}
	IssuerCertAuthority *EcsServiceServiceConnectConfigurationServiceTlsIssuerCertAuthority `field:"required" json:"issuerCertAuthority" yaml:"issuerCertAuthority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/ecs_service#kms_key EcsService#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/ecs_service#role_arn EcsService#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

