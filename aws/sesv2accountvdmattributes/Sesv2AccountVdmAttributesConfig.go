// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2accountvdmattributes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Sesv2AccountVdmAttributesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/sesv2_account_vdm_attributes#vdm_enabled Sesv2AccountVdmAttributes#vdm_enabled}.
	VdmEnabled *string `field:"required" json:"vdmEnabled" yaml:"vdmEnabled"`
	// dashboard_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/sesv2_account_vdm_attributes#dashboard_attributes Sesv2AccountVdmAttributes#dashboard_attributes}
	DashboardAttributes *Sesv2AccountVdmAttributesDashboardAttributes `field:"optional" json:"dashboardAttributes" yaml:"dashboardAttributes"`
	// guardian_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/sesv2_account_vdm_attributes#guardian_attributes Sesv2AccountVdmAttributes#guardian_attributes}
	GuardianAttributes *Sesv2AccountVdmAttributesGuardianAttributes `field:"optional" json:"guardianAttributes" yaml:"guardianAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/sesv2_account_vdm_attributes#id Sesv2AccountVdmAttributes#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

