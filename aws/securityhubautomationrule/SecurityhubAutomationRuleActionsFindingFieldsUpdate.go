// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrule


type SecurityhubAutomationRuleActionsFindingFieldsUpdate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securityhub_automation_rule#confidence SecurityhubAutomationRule#confidence}.
	Confidence *float64 `field:"optional" json:"confidence" yaml:"confidence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securityhub_automation_rule#criticality SecurityhubAutomationRule#criticality}.
	Criticality *float64 `field:"optional" json:"criticality" yaml:"criticality"`
	// note block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securityhub_automation_rule#note SecurityhubAutomationRule#note}
	Note interface{} `field:"optional" json:"note" yaml:"note"`
	// related_findings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securityhub_automation_rule#related_findings SecurityhubAutomationRule#related_findings}
	RelatedFindings interface{} `field:"optional" json:"relatedFindings" yaml:"relatedFindings"`
	// severity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securityhub_automation_rule#severity SecurityhubAutomationRule#severity}
	Severity interface{} `field:"optional" json:"severity" yaml:"severity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securityhub_automation_rule#types SecurityhubAutomationRule#types}.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securityhub_automation_rule#user_defined_fields SecurityhubAutomationRule#user_defined_fields}.
	UserDefinedFields *map[string]*string `field:"optional" json:"userDefinedFields" yaml:"userDefinedFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securityhub_automation_rule#verification_state SecurityhubAutomationRule#verification_state}.
	VerificationState *string `field:"optional" json:"verificationState" yaml:"verificationState"`
	// workflow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securityhub_automation_rule#workflow SecurityhubAutomationRule#workflow}
	Workflow interface{} `field:"optional" json:"workflow" yaml:"workflow"`
}

