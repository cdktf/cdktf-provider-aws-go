// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority


type AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/acmpca_certificate_authority#custom_cname AcmpcaCertificateAuthority#custom_cname}.
	CustomCname *string `field:"optional" json:"customCname" yaml:"customCname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/acmpca_certificate_authority#enabled AcmpcaCertificateAuthority#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/acmpca_certificate_authority#expiration_in_days AcmpcaCertificateAuthority#expiration_in_days}.
	ExpirationInDays *float64 `field:"optional" json:"expirationInDays" yaml:"expirationInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/acmpca_certificate_authority#s3_bucket_name AcmpcaCertificateAuthority#s3_bucket_name}.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/acmpca_certificate_authority#s3_object_acl AcmpcaCertificateAuthority#s3_object_acl}.
	S3ObjectAcl *string `field:"optional" json:"s3ObjectAcl" yaml:"s3ObjectAcl"`
}

