// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auditmanageraccountregistration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuditmanagerAccountRegistrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/auditmanager_account_registration#delegated_admin_account AuditmanagerAccountRegistration#delegated_admin_account}.
	DelegatedAdminAccount *string `field:"optional" json:"delegatedAdminAccount" yaml:"delegatedAdminAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/auditmanager_account_registration#deregister_on_destroy AuditmanagerAccountRegistration#deregister_on_destroy}.
	DeregisterOnDestroy interface{} `field:"optional" json:"deregisterOnDestroy" yaml:"deregisterOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/auditmanager_account_registration#kms_key AuditmanagerAccountRegistration#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/auditmanager_account_registration#region AuditmanagerAccountRegistration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

