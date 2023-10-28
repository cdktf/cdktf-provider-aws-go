// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guarddutyorganizationconfiguration


type GuarddutyOrganizationConfigurationDatasources struct {
	// kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/guardduty_organization_configuration#kubernetes GuarddutyOrganizationConfiguration#kubernetes}
	Kubernetes *GuarddutyOrganizationConfigurationDatasourcesKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// malware_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/guardduty_organization_configuration#malware_protection GuarddutyOrganizationConfiguration#malware_protection}
	MalwareProtection *GuarddutyOrganizationConfigurationDatasourcesMalwareProtection `field:"optional" json:"malwareProtection" yaml:"malwareProtection"`
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/guardduty_organization_configuration#s3_logs GuarddutyOrganizationConfiguration#s3_logs}
	S3Logs *GuarddutyOrganizationConfigurationDatasourcesS3Logs `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

