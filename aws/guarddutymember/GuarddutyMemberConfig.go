// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guarddutymember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GuarddutyMemberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_member#account_id GuarddutyMember#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_member#detector_id GuarddutyMember#detector_id}.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_member#email GuarddutyMember#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_member#disable_email_notification GuarddutyMember#disable_email_notification}.
	DisableEmailNotification interface{} `field:"optional" json:"disableEmailNotification" yaml:"disableEmailNotification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_member#id GuarddutyMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_member#invitation_message GuarddutyMember#invitation_message}.
	InvitationMessage *string `field:"optional" json:"invitationMessage" yaml:"invitationMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_member#invite GuarddutyMember#invite}.
	Invite interface{} `field:"optional" json:"invite" yaml:"invite"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_member#region GuarddutyMember#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_member#timeouts GuarddutyMember#timeouts}
	Timeouts *GuarddutyMemberTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

