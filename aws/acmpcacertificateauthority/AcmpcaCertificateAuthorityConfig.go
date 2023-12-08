// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AcmpcaCertificateAuthorityConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// certificate_authority_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/acmpca_certificate_authority#certificate_authority_configuration AcmpcaCertificateAuthority#certificate_authority_configuration}
	CertificateAuthorityConfiguration *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration `field:"required" json:"certificateAuthorityConfiguration" yaml:"certificateAuthorityConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/acmpca_certificate_authority#enabled AcmpcaCertificateAuthority#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/acmpca_certificate_authority#id AcmpcaCertificateAuthority#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/acmpca_certificate_authority#key_storage_security_standard AcmpcaCertificateAuthority#key_storage_security_standard}.
	KeyStorageSecurityStandard *string `field:"optional" json:"keyStorageSecurityStandard" yaml:"keyStorageSecurityStandard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/acmpca_certificate_authority#permanent_deletion_time_in_days AcmpcaCertificateAuthority#permanent_deletion_time_in_days}.
	PermanentDeletionTimeInDays *float64 `field:"optional" json:"permanentDeletionTimeInDays" yaml:"permanentDeletionTimeInDays"`
	// revocation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/acmpca_certificate_authority#revocation_configuration AcmpcaCertificateAuthority#revocation_configuration}
	RevocationConfiguration *AcmpcaCertificateAuthorityRevocationConfiguration `field:"optional" json:"revocationConfiguration" yaml:"revocationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/acmpca_certificate_authority#tags AcmpcaCertificateAuthority#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/acmpca_certificate_authority#tags_all AcmpcaCertificateAuthority#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/acmpca_certificate_authority#timeouts AcmpcaCertificateAuthority#timeouts}
	Timeouts *AcmpcaCertificateAuthorityTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/acmpca_certificate_authority#type AcmpcaCertificateAuthority#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/acmpca_certificate_authority#usage_mode AcmpcaCertificateAuthority#usage_mode}.
	UsageMode *string `field:"optional" json:"usageMode" yaml:"usageMode"`
}

