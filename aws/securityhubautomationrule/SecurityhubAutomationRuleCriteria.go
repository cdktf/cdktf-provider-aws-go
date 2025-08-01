// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrule


type SecurityhubAutomationRuleCriteria struct {
	// aws_account_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#aws_account_id SecurityhubAutomationRule#aws_account_id}
	AwsAccountId interface{} `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// aws_account_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#aws_account_name SecurityhubAutomationRule#aws_account_name}
	AwsAccountName interface{} `field:"optional" json:"awsAccountName" yaml:"awsAccountName"`
	// company_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#company_name SecurityhubAutomationRule#company_name}
	CompanyName interface{} `field:"optional" json:"companyName" yaml:"companyName"`
	// compliance_associated_standards_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#compliance_associated_standards_id SecurityhubAutomationRule#compliance_associated_standards_id}
	ComplianceAssociatedStandardsId interface{} `field:"optional" json:"complianceAssociatedStandardsId" yaml:"complianceAssociatedStandardsId"`
	// compliance_security_control_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#compliance_security_control_id SecurityhubAutomationRule#compliance_security_control_id}
	ComplianceSecurityControlId interface{} `field:"optional" json:"complianceSecurityControlId" yaml:"complianceSecurityControlId"`
	// compliance_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#compliance_status SecurityhubAutomationRule#compliance_status}
	ComplianceStatus interface{} `field:"optional" json:"complianceStatus" yaml:"complianceStatus"`
	// confidence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#confidence SecurityhubAutomationRule#confidence}
	Confidence interface{} `field:"optional" json:"confidence" yaml:"confidence"`
	// created_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#created_at SecurityhubAutomationRule#created_at}
	CreatedAt interface{} `field:"optional" json:"createdAt" yaml:"createdAt"`
	// criticality block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#criticality SecurityhubAutomationRule#criticality}
	Criticality interface{} `field:"optional" json:"criticality" yaml:"criticality"`
	// description block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#description SecurityhubAutomationRule#description}
	Description interface{} `field:"optional" json:"description" yaml:"description"`
	// first_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#first_observed_at SecurityhubAutomationRule#first_observed_at}
	FirstObservedAt interface{} `field:"optional" json:"firstObservedAt" yaml:"firstObservedAt"`
	// generator_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#generator_id SecurityhubAutomationRule#generator_id}
	GeneratorId interface{} `field:"optional" json:"generatorId" yaml:"generatorId"`
	// id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#id SecurityhubAutomationRule#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id interface{} `field:"optional" json:"id" yaml:"id"`
	// last_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#last_observed_at SecurityhubAutomationRule#last_observed_at}
	LastObservedAt interface{} `field:"optional" json:"lastObservedAt" yaml:"lastObservedAt"`
	// note_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#note_text SecurityhubAutomationRule#note_text}
	NoteText interface{} `field:"optional" json:"noteText" yaml:"noteText"`
	// note_updated_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#note_updated_at SecurityhubAutomationRule#note_updated_at}
	NoteUpdatedAt interface{} `field:"optional" json:"noteUpdatedAt" yaml:"noteUpdatedAt"`
	// note_updated_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#note_updated_by SecurityhubAutomationRule#note_updated_by}
	NoteUpdatedBy interface{} `field:"optional" json:"noteUpdatedBy" yaml:"noteUpdatedBy"`
	// product_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#product_arn SecurityhubAutomationRule#product_arn}
	ProductArn interface{} `field:"optional" json:"productArn" yaml:"productArn"`
	// product_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#product_name SecurityhubAutomationRule#product_name}
	ProductName interface{} `field:"optional" json:"productName" yaml:"productName"`
	// record_state block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#record_state SecurityhubAutomationRule#record_state}
	RecordState interface{} `field:"optional" json:"recordState" yaml:"recordState"`
	// related_findings_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#related_findings_id SecurityhubAutomationRule#related_findings_id}
	RelatedFindingsId interface{} `field:"optional" json:"relatedFindingsId" yaml:"relatedFindingsId"`
	// related_findings_product_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#related_findings_product_arn SecurityhubAutomationRule#related_findings_product_arn}
	RelatedFindingsProductArn interface{} `field:"optional" json:"relatedFindingsProductArn" yaml:"relatedFindingsProductArn"`
	// resource_application_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#resource_application_arn SecurityhubAutomationRule#resource_application_arn}
	ResourceApplicationArn interface{} `field:"optional" json:"resourceApplicationArn" yaml:"resourceApplicationArn"`
	// resource_application_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#resource_application_name SecurityhubAutomationRule#resource_application_name}
	ResourceApplicationName interface{} `field:"optional" json:"resourceApplicationName" yaml:"resourceApplicationName"`
	// resource_details_other block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#resource_details_other SecurityhubAutomationRule#resource_details_other}
	ResourceDetailsOther interface{} `field:"optional" json:"resourceDetailsOther" yaml:"resourceDetailsOther"`
	// resource_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#resource_id SecurityhubAutomationRule#resource_id}
	ResourceId interface{} `field:"optional" json:"resourceId" yaml:"resourceId"`
	// resource_partition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#resource_partition SecurityhubAutomationRule#resource_partition}
	ResourcePartition interface{} `field:"optional" json:"resourcePartition" yaml:"resourcePartition"`
	// resource_region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#resource_region SecurityhubAutomationRule#resource_region}
	ResourceRegion interface{} `field:"optional" json:"resourceRegion" yaml:"resourceRegion"`
	// resource_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#resource_tags SecurityhubAutomationRule#resource_tags}
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// resource_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#resource_type SecurityhubAutomationRule#resource_type}
	ResourceType interface{} `field:"optional" json:"resourceType" yaml:"resourceType"`
	// severity_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#severity_label SecurityhubAutomationRule#severity_label}
	SeverityLabel interface{} `field:"optional" json:"severityLabel" yaml:"severityLabel"`
	// source_url block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#source_url SecurityhubAutomationRule#source_url}
	SourceUrl interface{} `field:"optional" json:"sourceUrl" yaml:"sourceUrl"`
	// title block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#title SecurityhubAutomationRule#title}
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#type SecurityhubAutomationRule#type}
	Type interface{} `field:"optional" json:"type" yaml:"type"`
	// updated_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#updated_at SecurityhubAutomationRule#updated_at}
	UpdatedAt interface{} `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// user_defined_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#user_defined_fields SecurityhubAutomationRule#user_defined_fields}
	UserDefinedFields interface{} `field:"optional" json:"userDefinedFields" yaml:"userDefinedFields"`
	// verification_state block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#verification_state SecurityhubAutomationRule#verification_state}
	VerificationState interface{} `field:"optional" json:"verificationState" yaml:"verificationState"`
	// workflow_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_automation_rule#workflow_status SecurityhubAutomationRule#workflow_status}
	WorkflowStatus interface{} `field:"optional" json:"workflowStatus" yaml:"workflowStatus"`
}

