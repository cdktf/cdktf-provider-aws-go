// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/securityhubautomationrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubAutomationRuleCriteriaOutputReference interface {
	cdktf.ComplexObject
	AwsAccountId() SecurityhubAutomationRuleCriteriaAwsAccountIdList
	AwsAccountIdInput() interface{}
	AwsAccountName() SecurityhubAutomationRuleCriteriaAwsAccountNameList
	AwsAccountNameInput() interface{}
	CompanyName() SecurityhubAutomationRuleCriteriaCompanyNameList
	CompanyNameInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ComplianceAssociatedStandardsId() SecurityhubAutomationRuleCriteriaComplianceAssociatedStandardsIdList
	ComplianceAssociatedStandardsIdInput() interface{}
	ComplianceSecurityControlId() SecurityhubAutomationRuleCriteriaComplianceSecurityControlIdList
	ComplianceSecurityControlIdInput() interface{}
	ComplianceStatus() SecurityhubAutomationRuleCriteriaComplianceStatusList
	ComplianceStatusInput() interface{}
	Confidence() SecurityhubAutomationRuleCriteriaConfidenceList
	ConfidenceInput() interface{}
	CreatedAt() SecurityhubAutomationRuleCriteriaCreatedAtList
	CreatedAtInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Criticality() SecurityhubAutomationRuleCriteriaCriticalityList
	CriticalityInput() interface{}
	Description() SecurityhubAutomationRuleCriteriaDescriptionList
	DescriptionInput() interface{}
	FirstObservedAt() SecurityhubAutomationRuleCriteriaFirstObservedAtList
	FirstObservedAtInput() interface{}
	// Experimental.
	Fqn() *string
	GeneratorId() SecurityhubAutomationRuleCriteriaGeneratorIdList
	GeneratorIdInput() interface{}
	Id() SecurityhubAutomationRuleCriteriaIdList
	IdInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LastObservedAt() SecurityhubAutomationRuleCriteriaLastObservedAtList
	LastObservedAtInput() interface{}
	NoteText() SecurityhubAutomationRuleCriteriaNoteTextList
	NoteTextInput() interface{}
	NoteUpdatedAt() SecurityhubAutomationRuleCriteriaNoteUpdatedAtList
	NoteUpdatedAtInput() interface{}
	NoteUpdatedBy() SecurityhubAutomationRuleCriteriaNoteUpdatedByList
	NoteUpdatedByInput() interface{}
	ProductArn() SecurityhubAutomationRuleCriteriaProductArnList
	ProductArnInput() interface{}
	ProductName() SecurityhubAutomationRuleCriteriaProductNameList
	ProductNameInput() interface{}
	RecordState() SecurityhubAutomationRuleCriteriaRecordStateList
	RecordStateInput() interface{}
	RelatedFindingsId() SecurityhubAutomationRuleCriteriaRelatedFindingsIdList
	RelatedFindingsIdInput() interface{}
	RelatedFindingsProductArn() SecurityhubAutomationRuleCriteriaRelatedFindingsProductArnList
	RelatedFindingsProductArnInput() interface{}
	ResourceApplicationArn() SecurityhubAutomationRuleCriteriaResourceApplicationArnList
	ResourceApplicationArnInput() interface{}
	ResourceApplicationName() SecurityhubAutomationRuleCriteriaResourceApplicationNameList
	ResourceApplicationNameInput() interface{}
	ResourceDetailsOther() SecurityhubAutomationRuleCriteriaResourceDetailsOtherList
	ResourceDetailsOtherInput() interface{}
	ResourceId() SecurityhubAutomationRuleCriteriaResourceIdList
	ResourceIdInput() interface{}
	ResourcePartition() SecurityhubAutomationRuleCriteriaResourcePartitionList
	ResourcePartitionInput() interface{}
	ResourceRegion() SecurityhubAutomationRuleCriteriaResourceRegionList
	ResourceRegionInput() interface{}
	ResourceTags() SecurityhubAutomationRuleCriteriaResourceTagsList
	ResourceTagsInput() interface{}
	ResourceType() SecurityhubAutomationRuleCriteriaResourceTypeList
	ResourceTypeInput() interface{}
	SeverityLabel() SecurityhubAutomationRuleCriteriaSeverityLabelList
	SeverityLabelInput() interface{}
	SourceUrl() SecurityhubAutomationRuleCriteriaSourceUrlList
	SourceUrlInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() SecurityhubAutomationRuleCriteriaTitleList
	TitleInput() interface{}
	Type() SecurityhubAutomationRuleCriteriaTypeList
	TypeInput() interface{}
	UpdatedAt() SecurityhubAutomationRuleCriteriaUpdatedAtList
	UpdatedAtInput() interface{}
	UserDefinedFields() SecurityhubAutomationRuleCriteriaUserDefinedFieldsList
	UserDefinedFieldsInput() interface{}
	VerificationState() SecurityhubAutomationRuleCriteriaVerificationStateList
	VerificationStateInput() interface{}
	WorkflowStatus() SecurityhubAutomationRuleCriteriaWorkflowStatusList
	WorkflowStatusInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAwsAccountId(value interface{})
	PutAwsAccountName(value interface{})
	PutCompanyName(value interface{})
	PutComplianceAssociatedStandardsId(value interface{})
	PutComplianceSecurityControlId(value interface{})
	PutComplianceStatus(value interface{})
	PutConfidence(value interface{})
	PutCreatedAt(value interface{})
	PutCriticality(value interface{})
	PutDescription(value interface{})
	PutFirstObservedAt(value interface{})
	PutGeneratorId(value interface{})
	PutId(value interface{})
	PutLastObservedAt(value interface{})
	PutNoteText(value interface{})
	PutNoteUpdatedAt(value interface{})
	PutNoteUpdatedBy(value interface{})
	PutProductArn(value interface{})
	PutProductName(value interface{})
	PutRecordState(value interface{})
	PutRelatedFindingsId(value interface{})
	PutRelatedFindingsProductArn(value interface{})
	PutResourceApplicationArn(value interface{})
	PutResourceApplicationName(value interface{})
	PutResourceDetailsOther(value interface{})
	PutResourceId(value interface{})
	PutResourcePartition(value interface{})
	PutResourceRegion(value interface{})
	PutResourceTags(value interface{})
	PutResourceType(value interface{})
	PutSeverityLabel(value interface{})
	PutSourceUrl(value interface{})
	PutTitle(value interface{})
	PutType(value interface{})
	PutUpdatedAt(value interface{})
	PutUserDefinedFields(value interface{})
	PutVerificationState(value interface{})
	PutWorkflowStatus(value interface{})
	ResetAwsAccountId()
	ResetAwsAccountName()
	ResetCompanyName()
	ResetComplianceAssociatedStandardsId()
	ResetComplianceSecurityControlId()
	ResetComplianceStatus()
	ResetConfidence()
	ResetCreatedAt()
	ResetCriticality()
	ResetDescription()
	ResetFirstObservedAt()
	ResetGeneratorId()
	ResetId()
	ResetLastObservedAt()
	ResetNoteText()
	ResetNoteUpdatedAt()
	ResetNoteUpdatedBy()
	ResetProductArn()
	ResetProductName()
	ResetRecordState()
	ResetRelatedFindingsId()
	ResetRelatedFindingsProductArn()
	ResetResourceApplicationArn()
	ResetResourceApplicationName()
	ResetResourceDetailsOther()
	ResetResourceId()
	ResetResourcePartition()
	ResetResourceRegion()
	ResetResourceTags()
	ResetResourceType()
	ResetSeverityLabel()
	ResetSourceUrl()
	ResetTitle()
	ResetType()
	ResetUpdatedAt()
	ResetUserDefinedFields()
	ResetVerificationState()
	ResetWorkflowStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityhubAutomationRuleCriteriaOutputReference
type jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) AwsAccountId() SecurityhubAutomationRuleCriteriaAwsAccountIdList {
	var returns SecurityhubAutomationRuleCriteriaAwsAccountIdList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) AwsAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) AwsAccountName() SecurityhubAutomationRuleCriteriaAwsAccountNameList {
	var returns SecurityhubAutomationRuleCriteriaAwsAccountNameList
	_jsii_.Get(
		j,
		"awsAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) AwsAccountNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) CompanyName() SecurityhubAutomationRuleCriteriaCompanyNameList {
	var returns SecurityhubAutomationRuleCriteriaCompanyNameList
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) CompanyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"companyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ComplianceAssociatedStandardsId() SecurityhubAutomationRuleCriteriaComplianceAssociatedStandardsIdList {
	var returns SecurityhubAutomationRuleCriteriaComplianceAssociatedStandardsIdList
	_jsii_.Get(
		j,
		"complianceAssociatedStandardsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ComplianceAssociatedStandardsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceAssociatedStandardsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ComplianceSecurityControlId() SecurityhubAutomationRuleCriteriaComplianceSecurityControlIdList {
	var returns SecurityhubAutomationRuleCriteriaComplianceSecurityControlIdList
	_jsii_.Get(
		j,
		"complianceSecurityControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ComplianceSecurityControlIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceSecurityControlIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ComplianceStatus() SecurityhubAutomationRuleCriteriaComplianceStatusList {
	var returns SecurityhubAutomationRuleCriteriaComplianceStatusList
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ComplianceStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) Confidence() SecurityhubAutomationRuleCriteriaConfidenceList {
	var returns SecurityhubAutomationRuleCriteriaConfidenceList
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ConfidenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) CreatedAt() SecurityhubAutomationRuleCriteriaCreatedAtList {
	var returns SecurityhubAutomationRuleCriteriaCreatedAtList
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) CreatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) Criticality() SecurityhubAutomationRuleCriteriaCriticalityList {
	var returns SecurityhubAutomationRuleCriteriaCriticalityList
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) CriticalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) Description() SecurityhubAutomationRuleCriteriaDescriptionList {
	var returns SecurityhubAutomationRuleCriteriaDescriptionList
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) DescriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) FirstObservedAt() SecurityhubAutomationRuleCriteriaFirstObservedAtList {
	var returns SecurityhubAutomationRuleCriteriaFirstObservedAtList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) FirstObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) GeneratorId() SecurityhubAutomationRuleCriteriaGeneratorIdList {
	var returns SecurityhubAutomationRuleCriteriaGeneratorIdList
	_jsii_.Get(
		j,
		"generatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) GeneratorIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generatorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) Id() SecurityhubAutomationRuleCriteriaIdList {
	var returns SecurityhubAutomationRuleCriteriaIdList
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) IdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) LastObservedAt() SecurityhubAutomationRuleCriteriaLastObservedAtList {
	var returns SecurityhubAutomationRuleCriteriaLastObservedAtList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) LastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) NoteText() SecurityhubAutomationRuleCriteriaNoteTextList {
	var returns SecurityhubAutomationRuleCriteriaNoteTextList
	_jsii_.Get(
		j,
		"noteText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) NoteTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) NoteUpdatedAt() SecurityhubAutomationRuleCriteriaNoteUpdatedAtList {
	var returns SecurityhubAutomationRuleCriteriaNoteUpdatedAtList
	_jsii_.Get(
		j,
		"noteUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) NoteUpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) NoteUpdatedBy() SecurityhubAutomationRuleCriteriaNoteUpdatedByList {
	var returns SecurityhubAutomationRuleCriteriaNoteUpdatedByList
	_jsii_.Get(
		j,
		"noteUpdatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) NoteUpdatedByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ProductArn() SecurityhubAutomationRuleCriteriaProductArnList {
	var returns SecurityhubAutomationRuleCriteriaProductArnList
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ProductName() SecurityhubAutomationRuleCriteriaProductNameList {
	var returns SecurityhubAutomationRuleCriteriaProductNameList
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ProductNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) RecordState() SecurityhubAutomationRuleCriteriaRecordStateList {
	var returns SecurityhubAutomationRuleCriteriaRecordStateList
	_jsii_.Get(
		j,
		"recordState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) RecordStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) RelatedFindingsId() SecurityhubAutomationRuleCriteriaRelatedFindingsIdList {
	var returns SecurityhubAutomationRuleCriteriaRelatedFindingsIdList
	_jsii_.Get(
		j,
		"relatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) RelatedFindingsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) RelatedFindingsProductArn() SecurityhubAutomationRuleCriteriaRelatedFindingsProductArnList {
	var returns SecurityhubAutomationRuleCriteriaRelatedFindingsProductArnList
	_jsii_.Get(
		j,
		"relatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) RelatedFindingsProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceApplicationArn() SecurityhubAutomationRuleCriteriaResourceApplicationArnList {
	var returns SecurityhubAutomationRuleCriteriaResourceApplicationArnList
	_jsii_.Get(
		j,
		"resourceApplicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceApplicationArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceApplicationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceApplicationName() SecurityhubAutomationRuleCriteriaResourceApplicationNameList {
	var returns SecurityhubAutomationRuleCriteriaResourceApplicationNameList
	_jsii_.Get(
		j,
		"resourceApplicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceApplicationNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceApplicationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceDetailsOther() SecurityhubAutomationRuleCriteriaResourceDetailsOtherList {
	var returns SecurityhubAutomationRuleCriteriaResourceDetailsOtherList
	_jsii_.Get(
		j,
		"resourceDetailsOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceDetailsOtherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceDetailsOtherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceId() SecurityhubAutomationRuleCriteriaResourceIdList {
	var returns SecurityhubAutomationRuleCriteriaResourceIdList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourcePartition() SecurityhubAutomationRuleCriteriaResourcePartitionList {
	var returns SecurityhubAutomationRuleCriteriaResourcePartitionList
	_jsii_.Get(
		j,
		"resourcePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourcePartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcePartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceRegion() SecurityhubAutomationRuleCriteriaResourceRegionList {
	var returns SecurityhubAutomationRuleCriteriaResourceRegionList
	_jsii_.Get(
		j,
		"resourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceRegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceTags() SecurityhubAutomationRuleCriteriaResourceTagsList {
	var returns SecurityhubAutomationRuleCriteriaResourceTagsList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceType() SecurityhubAutomationRuleCriteriaResourceTypeList {
	var returns SecurityhubAutomationRuleCriteriaResourceTypeList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) SeverityLabel() SecurityhubAutomationRuleCriteriaSeverityLabelList {
	var returns SecurityhubAutomationRuleCriteriaSeverityLabelList
	_jsii_.Get(
		j,
		"severityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) SeverityLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) SourceUrl() SecurityhubAutomationRuleCriteriaSourceUrlList {
	var returns SecurityhubAutomationRuleCriteriaSourceUrlList
	_jsii_.Get(
		j,
		"sourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) SourceUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) Title() SecurityhubAutomationRuleCriteriaTitleList {
	var returns SecurityhubAutomationRuleCriteriaTitleList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) TitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) Type() SecurityhubAutomationRuleCriteriaTypeList {
	var returns SecurityhubAutomationRuleCriteriaTypeList
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) TypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) UpdatedAt() SecurityhubAutomationRuleCriteriaUpdatedAtList {
	var returns SecurityhubAutomationRuleCriteriaUpdatedAtList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) UpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) UserDefinedFields() SecurityhubAutomationRuleCriteriaUserDefinedFieldsList {
	var returns SecurityhubAutomationRuleCriteriaUserDefinedFieldsList
	_jsii_.Get(
		j,
		"userDefinedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) UserDefinedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefinedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) VerificationState() SecurityhubAutomationRuleCriteriaVerificationStateList {
	var returns SecurityhubAutomationRuleCriteriaVerificationStateList
	_jsii_.Get(
		j,
		"verificationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) VerificationStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verificationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) WorkflowStatus() SecurityhubAutomationRuleCriteriaWorkflowStatusList {
	var returns SecurityhubAutomationRuleCriteriaWorkflowStatusList
	_jsii_.Get(
		j,
		"workflowStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) WorkflowStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowStatusInput",
		&returns,
	)
	return returns
}


func NewSecurityhubAutomationRuleCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SecurityhubAutomationRuleCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityhubAutomationRuleCriteriaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubAutomationRule.SecurityhubAutomationRuleCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSecurityhubAutomationRuleCriteriaOutputReference_Override(s SecurityhubAutomationRuleCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubAutomationRule.SecurityhubAutomationRuleCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutAwsAccountId(value interface{}) {
	if err := s.validatePutAwsAccountIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAwsAccountId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutAwsAccountName(value interface{}) {
	if err := s.validatePutAwsAccountNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAwsAccountName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutCompanyName(value interface{}) {
	if err := s.validatePutCompanyNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCompanyName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutComplianceAssociatedStandardsId(value interface{}) {
	if err := s.validatePutComplianceAssociatedStandardsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putComplianceAssociatedStandardsId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutComplianceSecurityControlId(value interface{}) {
	if err := s.validatePutComplianceSecurityControlIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putComplianceSecurityControlId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutComplianceStatus(value interface{}) {
	if err := s.validatePutComplianceStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putComplianceStatus",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutConfidence(value interface{}) {
	if err := s.validatePutConfidenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putConfidence",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutCreatedAt(value interface{}) {
	if err := s.validatePutCreatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCreatedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutCriticality(value interface{}) {
	if err := s.validatePutCriticalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCriticality",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutDescription(value interface{}) {
	if err := s.validatePutDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDescription",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutFirstObservedAt(value interface{}) {
	if err := s.validatePutFirstObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFirstObservedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutGeneratorId(value interface{}) {
	if err := s.validatePutGeneratorIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGeneratorId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutId(value interface{}) {
	if err := s.validatePutIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutLastObservedAt(value interface{}) {
	if err := s.validatePutLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLastObservedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutNoteText(value interface{}) {
	if err := s.validatePutNoteTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNoteText",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutNoteUpdatedAt(value interface{}) {
	if err := s.validatePutNoteUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNoteUpdatedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutNoteUpdatedBy(value interface{}) {
	if err := s.validatePutNoteUpdatedByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNoteUpdatedBy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutProductArn(value interface{}) {
	if err := s.validatePutProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProductArn",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutProductName(value interface{}) {
	if err := s.validatePutProductNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProductName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutRecordState(value interface{}) {
	if err := s.validatePutRecordStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRecordState",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutRelatedFindingsId(value interface{}) {
	if err := s.validatePutRelatedFindingsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRelatedFindingsId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutRelatedFindingsProductArn(value interface{}) {
	if err := s.validatePutRelatedFindingsProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRelatedFindingsProductArn",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutResourceApplicationArn(value interface{}) {
	if err := s.validatePutResourceApplicationArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceApplicationArn",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutResourceApplicationName(value interface{}) {
	if err := s.validatePutResourceApplicationNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceApplicationName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutResourceDetailsOther(value interface{}) {
	if err := s.validatePutResourceDetailsOtherParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceDetailsOther",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutResourceId(value interface{}) {
	if err := s.validatePutResourceIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutResourcePartition(value interface{}) {
	if err := s.validatePutResourcePartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourcePartition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutResourceRegion(value interface{}) {
	if err := s.validatePutResourceRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceRegion",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutResourceTags(value interface{}) {
	if err := s.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutResourceType(value interface{}) {
	if err := s.validatePutResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceType",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutSeverityLabel(value interface{}) {
	if err := s.validatePutSeverityLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSeverityLabel",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutSourceUrl(value interface{}) {
	if err := s.validatePutSourceUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSourceUrl",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutTitle(value interface{}) {
	if err := s.validatePutTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTitle",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutType(value interface{}) {
	if err := s.validatePutTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putType",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutUpdatedAt(value interface{}) {
	if err := s.validatePutUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUpdatedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutUserDefinedFields(value interface{}) {
	if err := s.validatePutUserDefinedFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUserDefinedFields",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutVerificationState(value interface{}) {
	if err := s.validatePutVerificationStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVerificationState",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) PutWorkflowStatus(value interface{}) {
	if err := s.validatePutWorkflowStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWorkflowStatus",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetAwsAccountName() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsAccountName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetCompanyName() {
	_jsii_.InvokeVoid(
		s,
		"resetCompanyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetComplianceAssociatedStandardsId() {
	_jsii_.InvokeVoid(
		s,
		"resetComplianceAssociatedStandardsId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetComplianceSecurityControlId() {
	_jsii_.InvokeVoid(
		s,
		"resetComplianceSecurityControlId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetComplianceStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetComplianceStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetConfidence() {
	_jsii_.InvokeVoid(
		s,
		"resetConfidence",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		s,
		"resetCriticality",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetFirstObservedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetFirstObservedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetGeneratorId() {
	_jsii_.InvokeVoid(
		s,
		"resetGeneratorId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetLastObservedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetLastObservedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetNoteText() {
	_jsii_.InvokeVoid(
		s,
		"resetNoteText",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetNoteUpdatedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetNoteUpdatedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetNoteUpdatedBy() {
	_jsii_.InvokeVoid(
		s,
		"resetNoteUpdatedBy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetProductArn() {
	_jsii_.InvokeVoid(
		s,
		"resetProductArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetProductName() {
	_jsii_.InvokeVoid(
		s,
		"resetProductName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetRecordState() {
	_jsii_.InvokeVoid(
		s,
		"resetRecordState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetRelatedFindingsId() {
	_jsii_.InvokeVoid(
		s,
		"resetRelatedFindingsId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetRelatedFindingsProductArn() {
	_jsii_.InvokeVoid(
		s,
		"resetRelatedFindingsProductArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetResourceApplicationArn() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceApplicationArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetResourceApplicationName() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceApplicationName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetResourceDetailsOther() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceDetailsOther",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetResourcePartition() {
	_jsii_.InvokeVoid(
		s,
		"resetResourcePartition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetResourceRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetSeverityLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetSeverityLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetSourceUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		s,
		"resetTitle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetUserDefinedFields() {
	_jsii_.InvokeVoid(
		s,
		"resetUserDefinedFields",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetVerificationState() {
	_jsii_.InvokeVoid(
		s,
		"resetVerificationState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ResetWorkflowStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkflowStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

