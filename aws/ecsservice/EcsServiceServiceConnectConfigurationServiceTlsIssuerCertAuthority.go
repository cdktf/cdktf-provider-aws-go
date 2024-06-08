// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceConnectConfigurationServiceTlsIssuerCertAuthority struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/ecs_service#aws_pca_authority_arn EcsService#aws_pca_authority_arn}.
	AwsPcaAuthorityArn *string `field:"required" json:"awsPcaAuthorityArn" yaml:"awsPcaAuthorityArn"`
}

