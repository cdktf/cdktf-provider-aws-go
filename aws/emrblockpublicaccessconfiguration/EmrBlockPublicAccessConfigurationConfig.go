// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrblockpublicaccessconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrBlockPublicAccessConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/emr_block_public_access_configuration#block_public_security_group_rules EmrBlockPublicAccessConfiguration#block_public_security_group_rules}.
	BlockPublicSecurityGroupRules interface{} `field:"required" json:"blockPublicSecurityGroupRules" yaml:"blockPublicSecurityGroupRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/emr_block_public_access_configuration#id EmrBlockPublicAccessConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// permitted_public_security_group_rule_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/emr_block_public_access_configuration#permitted_public_security_group_rule_range EmrBlockPublicAccessConfiguration#permitted_public_security_group_rule_range}
	PermittedPublicSecurityGroupRuleRange interface{} `field:"optional" json:"permittedPublicSecurityGroupRuleRange" yaml:"permittedPublicSecurityGroupRuleRange"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/emr_block_public_access_configuration#region EmrBlockPublicAccessConfiguration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

