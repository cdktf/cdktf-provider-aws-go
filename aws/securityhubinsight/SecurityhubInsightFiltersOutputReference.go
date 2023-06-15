package securityhubinsight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v15/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v15/securityhubinsight/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubInsightFiltersOutputReference interface {
	cdktf.ComplexObject
	AwsAccountId() SecurityhubInsightFiltersAwsAccountIdList
	AwsAccountIdInput() interface{}
	CompanyName() SecurityhubInsightFiltersCompanyNameList
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
	ComplianceStatus() SecurityhubInsightFiltersComplianceStatusList
	ComplianceStatusInput() interface{}
	Confidence() SecurityhubInsightFiltersConfidenceList
	ConfidenceInput() interface{}
	CreatedAt() SecurityhubInsightFiltersCreatedAtList
	CreatedAtInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Criticality() SecurityhubInsightFiltersCriticalityList
	CriticalityInput() interface{}
	Description() SecurityhubInsightFiltersDescriptionList
	DescriptionInput() interface{}
	FindingProviderFieldsConfidence() SecurityhubInsightFiltersFindingProviderFieldsConfidenceList
	FindingProviderFieldsConfidenceInput() interface{}
	FindingProviderFieldsCriticality() SecurityhubInsightFiltersFindingProviderFieldsCriticalityList
	FindingProviderFieldsCriticalityInput() interface{}
	FindingProviderFieldsRelatedFindingsId() SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsIdList
	FindingProviderFieldsRelatedFindingsIdInput() interface{}
	FindingProviderFieldsRelatedFindingsProductArn() SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArnList
	FindingProviderFieldsRelatedFindingsProductArnInput() interface{}
	FindingProviderFieldsSeverityLabel() SecurityhubInsightFiltersFindingProviderFieldsSeverityLabelList
	FindingProviderFieldsSeverityLabelInput() interface{}
	FindingProviderFieldsSeverityOriginal() SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginalList
	FindingProviderFieldsSeverityOriginalInput() interface{}
	FindingProviderFieldsTypes() SecurityhubInsightFiltersFindingProviderFieldsTypesList
	FindingProviderFieldsTypesInput() interface{}
	FirstObservedAt() SecurityhubInsightFiltersFirstObservedAtList
	FirstObservedAtInput() interface{}
	// Experimental.
	Fqn() *string
	GeneratorId() SecurityhubInsightFiltersGeneratorIdList
	GeneratorIdInput() interface{}
	Id() SecurityhubInsightFiltersIdList
	IdInput() interface{}
	InternalValue() *SecurityhubInsightFilters
	SetInternalValue(val *SecurityhubInsightFilters)
	Keyword() SecurityhubInsightFiltersKeywordList
	KeywordInput() interface{}
	LastObservedAt() SecurityhubInsightFiltersLastObservedAtList
	LastObservedAtInput() interface{}
	MalwareName() SecurityhubInsightFiltersMalwareNameList
	MalwareNameInput() interface{}
	MalwarePath() SecurityhubInsightFiltersMalwarePathList
	MalwarePathInput() interface{}
	MalwareState() SecurityhubInsightFiltersMalwareStateList
	MalwareStateInput() interface{}
	MalwareType() SecurityhubInsightFiltersMalwareTypeList
	MalwareTypeInput() interface{}
	NetworkDestinationDomain() SecurityhubInsightFiltersNetworkDestinationDomainList
	NetworkDestinationDomainInput() interface{}
	NetworkDestinationIpv4() SecurityhubInsightFiltersNetworkDestinationIpv4List
	NetworkDestinationIpv4Input() interface{}
	NetworkDestinationIpv6() SecurityhubInsightFiltersNetworkDestinationIpv6List
	NetworkDestinationIpv6Input() interface{}
	NetworkDestinationPort() SecurityhubInsightFiltersNetworkDestinationPortList
	NetworkDestinationPortInput() interface{}
	NetworkDirection() SecurityhubInsightFiltersNetworkDirectionList
	NetworkDirectionInput() interface{}
	NetworkProtocol() SecurityhubInsightFiltersNetworkProtocolList
	NetworkProtocolInput() interface{}
	NetworkSourceDomain() SecurityhubInsightFiltersNetworkSourceDomainList
	NetworkSourceDomainInput() interface{}
	NetworkSourceIpv4() SecurityhubInsightFiltersNetworkSourceIpv4List
	NetworkSourceIpv4Input() interface{}
	NetworkSourceIpv6() SecurityhubInsightFiltersNetworkSourceIpv6List
	NetworkSourceIpv6Input() interface{}
	NetworkSourceMac() SecurityhubInsightFiltersNetworkSourceMacList
	NetworkSourceMacInput() interface{}
	NetworkSourcePort() SecurityhubInsightFiltersNetworkSourcePortList
	NetworkSourcePortInput() interface{}
	NoteText() SecurityhubInsightFiltersNoteTextList
	NoteTextInput() interface{}
	NoteUpdatedAt() SecurityhubInsightFiltersNoteUpdatedAtList
	NoteUpdatedAtInput() interface{}
	NoteUpdatedBy() SecurityhubInsightFiltersNoteUpdatedByList
	NoteUpdatedByInput() interface{}
	ProcessLaunchedAt() SecurityhubInsightFiltersProcessLaunchedAtList
	ProcessLaunchedAtInput() interface{}
	ProcessName() SecurityhubInsightFiltersProcessNameList
	ProcessNameInput() interface{}
	ProcessParentPid() SecurityhubInsightFiltersProcessParentPidList
	ProcessParentPidInput() interface{}
	ProcessPath() SecurityhubInsightFiltersProcessPathList
	ProcessPathInput() interface{}
	ProcessPid() SecurityhubInsightFiltersProcessPidList
	ProcessPidInput() interface{}
	ProcessTerminatedAt() SecurityhubInsightFiltersProcessTerminatedAtList
	ProcessTerminatedAtInput() interface{}
	ProductArn() SecurityhubInsightFiltersProductArnList
	ProductArnInput() interface{}
	ProductFields() SecurityhubInsightFiltersProductFieldsList
	ProductFieldsInput() interface{}
	ProductName() SecurityhubInsightFiltersProductNameList
	ProductNameInput() interface{}
	RecommendationText() SecurityhubInsightFiltersRecommendationTextList
	RecommendationTextInput() interface{}
	RecordState() SecurityhubInsightFiltersRecordStateList
	RecordStateInput() interface{}
	RelatedFindingsId() SecurityhubInsightFiltersRelatedFindingsIdList
	RelatedFindingsIdInput() interface{}
	RelatedFindingsProductArn() SecurityhubInsightFiltersRelatedFindingsProductArnList
	RelatedFindingsProductArnInput() interface{}
	ResourceAwsEc2InstanceIamInstanceProfileArn() SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArnList
	ResourceAwsEc2InstanceIamInstanceProfileArnInput() interface{}
	ResourceAwsEc2InstanceImageId() SecurityhubInsightFiltersResourceAwsEc2InstanceImageIdList
	ResourceAwsEc2InstanceImageIdInput() interface{}
	ResourceAwsEc2InstanceIpv4Addresses() SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4AddressesList
	ResourceAwsEc2InstanceIpv4AddressesInput() interface{}
	ResourceAwsEc2InstanceIpv6Addresses() SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6AddressesList
	ResourceAwsEc2InstanceIpv6AddressesInput() interface{}
	ResourceAwsEc2InstanceKeyName() SecurityhubInsightFiltersResourceAwsEc2InstanceKeyNameList
	ResourceAwsEc2InstanceKeyNameInput() interface{}
	ResourceAwsEc2InstanceLaunchedAt() SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtList
	ResourceAwsEc2InstanceLaunchedAtInput() interface{}
	ResourceAwsEc2InstanceSubnetId() SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetIdList
	ResourceAwsEc2InstanceSubnetIdInput() interface{}
	ResourceAwsEc2InstanceType() SecurityhubInsightFiltersResourceAwsEc2InstanceTypeList
	ResourceAwsEc2InstanceTypeInput() interface{}
	ResourceAwsEc2InstanceVpcId() SecurityhubInsightFiltersResourceAwsEc2InstanceVpcIdList
	ResourceAwsEc2InstanceVpcIdInput() interface{}
	ResourceAwsIamAccessKeyCreatedAt() SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtList
	ResourceAwsIamAccessKeyCreatedAtInput() interface{}
	ResourceAwsIamAccessKeyStatus() SecurityhubInsightFiltersResourceAwsIamAccessKeyStatusList
	ResourceAwsIamAccessKeyStatusInput() interface{}
	ResourceAwsIamAccessKeyUserName() SecurityhubInsightFiltersResourceAwsIamAccessKeyUserNameList
	ResourceAwsIamAccessKeyUserNameInput() interface{}
	ResourceAwsS3BucketOwnerId() SecurityhubInsightFiltersResourceAwsS3BucketOwnerIdList
	ResourceAwsS3BucketOwnerIdInput() interface{}
	ResourceAwsS3BucketOwnerName() SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList
	ResourceAwsS3BucketOwnerNameInput() interface{}
	ResourceContainerImageId() SecurityhubInsightFiltersResourceContainerImageIdList
	ResourceContainerImageIdInput() interface{}
	ResourceContainerImageName() SecurityhubInsightFiltersResourceContainerImageNameList
	ResourceContainerImageNameInput() interface{}
	ResourceContainerLaunchedAt() SecurityhubInsightFiltersResourceContainerLaunchedAtList
	ResourceContainerLaunchedAtInput() interface{}
	ResourceContainerName() SecurityhubInsightFiltersResourceContainerNameList
	ResourceContainerNameInput() interface{}
	ResourceDetailsOther() SecurityhubInsightFiltersResourceDetailsOtherList
	ResourceDetailsOtherInput() interface{}
	ResourceId() SecurityhubInsightFiltersResourceIdList
	ResourceIdInput() interface{}
	ResourcePartition() SecurityhubInsightFiltersResourcePartitionList
	ResourcePartitionInput() interface{}
	ResourceRegion() SecurityhubInsightFiltersResourceRegionList
	ResourceRegionInput() interface{}
	ResourceTags() SecurityhubInsightFiltersResourceTagsList
	ResourceTagsInput() interface{}
	ResourceType() SecurityhubInsightFiltersResourceTypeList
	ResourceTypeInput() interface{}
	SeverityLabel() SecurityhubInsightFiltersSeverityLabelList
	SeverityLabelInput() interface{}
	SourceUrl() SecurityhubInsightFiltersSourceUrlList
	SourceUrlInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThreatIntelIndicatorCategory() SecurityhubInsightFiltersThreatIntelIndicatorCategoryList
	ThreatIntelIndicatorCategoryInput() interface{}
	ThreatIntelIndicatorLastObservedAt() SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtList
	ThreatIntelIndicatorLastObservedAtInput() interface{}
	ThreatIntelIndicatorSource() SecurityhubInsightFiltersThreatIntelIndicatorSourceList
	ThreatIntelIndicatorSourceInput() interface{}
	ThreatIntelIndicatorSourceUrl() SecurityhubInsightFiltersThreatIntelIndicatorSourceUrlList
	ThreatIntelIndicatorSourceUrlInput() interface{}
	ThreatIntelIndicatorType() SecurityhubInsightFiltersThreatIntelIndicatorTypeList
	ThreatIntelIndicatorTypeInput() interface{}
	ThreatIntelIndicatorValue() SecurityhubInsightFiltersThreatIntelIndicatorValueList
	ThreatIntelIndicatorValueInput() interface{}
	Title() SecurityhubInsightFiltersTitleList
	TitleInput() interface{}
	Type() SecurityhubInsightFiltersTypeList
	TypeInput() interface{}
	UpdatedAt() SecurityhubInsightFiltersUpdatedAtList
	UpdatedAtInput() interface{}
	UserDefinedValues() SecurityhubInsightFiltersUserDefinedValuesList
	UserDefinedValuesInput() interface{}
	VerificationState() SecurityhubInsightFiltersVerificationStateList
	VerificationStateInput() interface{}
	WorkflowStatus() SecurityhubInsightFiltersWorkflowStatusList
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
	PutCompanyName(value interface{})
	PutComplianceStatus(value interface{})
	PutConfidence(value interface{})
	PutCreatedAt(value interface{})
	PutCriticality(value interface{})
	PutDescription(value interface{})
	PutFindingProviderFieldsConfidence(value interface{})
	PutFindingProviderFieldsCriticality(value interface{})
	PutFindingProviderFieldsRelatedFindingsId(value interface{})
	PutFindingProviderFieldsRelatedFindingsProductArn(value interface{})
	PutFindingProviderFieldsSeverityLabel(value interface{})
	PutFindingProviderFieldsSeverityOriginal(value interface{})
	PutFindingProviderFieldsTypes(value interface{})
	PutFirstObservedAt(value interface{})
	PutGeneratorId(value interface{})
	PutId(value interface{})
	PutKeyword(value interface{})
	PutLastObservedAt(value interface{})
	PutMalwareName(value interface{})
	PutMalwarePath(value interface{})
	PutMalwareState(value interface{})
	PutMalwareType(value interface{})
	PutNetworkDestinationDomain(value interface{})
	PutNetworkDestinationIpv4(value interface{})
	PutNetworkDestinationIpv6(value interface{})
	PutNetworkDestinationPort(value interface{})
	PutNetworkDirection(value interface{})
	PutNetworkProtocol(value interface{})
	PutNetworkSourceDomain(value interface{})
	PutNetworkSourceIpv4(value interface{})
	PutNetworkSourceIpv6(value interface{})
	PutNetworkSourceMac(value interface{})
	PutNetworkSourcePort(value interface{})
	PutNoteText(value interface{})
	PutNoteUpdatedAt(value interface{})
	PutNoteUpdatedBy(value interface{})
	PutProcessLaunchedAt(value interface{})
	PutProcessName(value interface{})
	PutProcessParentPid(value interface{})
	PutProcessPath(value interface{})
	PutProcessPid(value interface{})
	PutProcessTerminatedAt(value interface{})
	PutProductArn(value interface{})
	PutProductFields(value interface{})
	PutProductName(value interface{})
	PutRecommendationText(value interface{})
	PutRecordState(value interface{})
	PutRelatedFindingsId(value interface{})
	PutRelatedFindingsProductArn(value interface{})
	PutResourceAwsEc2InstanceIamInstanceProfileArn(value interface{})
	PutResourceAwsEc2InstanceImageId(value interface{})
	PutResourceAwsEc2InstanceIpv4Addresses(value interface{})
	PutResourceAwsEc2InstanceIpv6Addresses(value interface{})
	PutResourceAwsEc2InstanceKeyName(value interface{})
	PutResourceAwsEc2InstanceLaunchedAt(value interface{})
	PutResourceAwsEc2InstanceSubnetId(value interface{})
	PutResourceAwsEc2InstanceType(value interface{})
	PutResourceAwsEc2InstanceVpcId(value interface{})
	PutResourceAwsIamAccessKeyCreatedAt(value interface{})
	PutResourceAwsIamAccessKeyStatus(value interface{})
	PutResourceAwsIamAccessKeyUserName(value interface{})
	PutResourceAwsS3BucketOwnerId(value interface{})
	PutResourceAwsS3BucketOwnerName(value interface{})
	PutResourceContainerImageId(value interface{})
	PutResourceContainerImageName(value interface{})
	PutResourceContainerLaunchedAt(value interface{})
	PutResourceContainerName(value interface{})
	PutResourceDetailsOther(value interface{})
	PutResourceId(value interface{})
	PutResourcePartition(value interface{})
	PutResourceRegion(value interface{})
	PutResourceTags(value interface{})
	PutResourceType(value interface{})
	PutSeverityLabel(value interface{})
	PutSourceUrl(value interface{})
	PutThreatIntelIndicatorCategory(value interface{})
	PutThreatIntelIndicatorLastObservedAt(value interface{})
	PutThreatIntelIndicatorSource(value interface{})
	PutThreatIntelIndicatorSourceUrl(value interface{})
	PutThreatIntelIndicatorType(value interface{})
	PutThreatIntelIndicatorValue(value interface{})
	PutTitle(value interface{})
	PutType(value interface{})
	PutUpdatedAt(value interface{})
	PutUserDefinedValues(value interface{})
	PutVerificationState(value interface{})
	PutWorkflowStatus(value interface{})
	ResetAwsAccountId()
	ResetCompanyName()
	ResetComplianceStatus()
	ResetConfidence()
	ResetCreatedAt()
	ResetCriticality()
	ResetDescription()
	ResetFindingProviderFieldsConfidence()
	ResetFindingProviderFieldsCriticality()
	ResetFindingProviderFieldsRelatedFindingsId()
	ResetFindingProviderFieldsRelatedFindingsProductArn()
	ResetFindingProviderFieldsSeverityLabel()
	ResetFindingProviderFieldsSeverityOriginal()
	ResetFindingProviderFieldsTypes()
	ResetFirstObservedAt()
	ResetGeneratorId()
	ResetId()
	ResetKeyword()
	ResetLastObservedAt()
	ResetMalwareName()
	ResetMalwarePath()
	ResetMalwareState()
	ResetMalwareType()
	ResetNetworkDestinationDomain()
	ResetNetworkDestinationIpv4()
	ResetNetworkDestinationIpv6()
	ResetNetworkDestinationPort()
	ResetNetworkDirection()
	ResetNetworkProtocol()
	ResetNetworkSourceDomain()
	ResetNetworkSourceIpv4()
	ResetNetworkSourceIpv6()
	ResetNetworkSourceMac()
	ResetNetworkSourcePort()
	ResetNoteText()
	ResetNoteUpdatedAt()
	ResetNoteUpdatedBy()
	ResetProcessLaunchedAt()
	ResetProcessName()
	ResetProcessParentPid()
	ResetProcessPath()
	ResetProcessPid()
	ResetProcessTerminatedAt()
	ResetProductArn()
	ResetProductFields()
	ResetProductName()
	ResetRecommendationText()
	ResetRecordState()
	ResetRelatedFindingsId()
	ResetRelatedFindingsProductArn()
	ResetResourceAwsEc2InstanceIamInstanceProfileArn()
	ResetResourceAwsEc2InstanceImageId()
	ResetResourceAwsEc2InstanceIpv4Addresses()
	ResetResourceAwsEc2InstanceIpv6Addresses()
	ResetResourceAwsEc2InstanceKeyName()
	ResetResourceAwsEc2InstanceLaunchedAt()
	ResetResourceAwsEc2InstanceSubnetId()
	ResetResourceAwsEc2InstanceType()
	ResetResourceAwsEc2InstanceVpcId()
	ResetResourceAwsIamAccessKeyCreatedAt()
	ResetResourceAwsIamAccessKeyStatus()
	ResetResourceAwsIamAccessKeyUserName()
	ResetResourceAwsS3BucketOwnerId()
	ResetResourceAwsS3BucketOwnerName()
	ResetResourceContainerImageId()
	ResetResourceContainerImageName()
	ResetResourceContainerLaunchedAt()
	ResetResourceContainerName()
	ResetResourceDetailsOther()
	ResetResourceId()
	ResetResourcePartition()
	ResetResourceRegion()
	ResetResourceTags()
	ResetResourceType()
	ResetSeverityLabel()
	ResetSourceUrl()
	ResetThreatIntelIndicatorCategory()
	ResetThreatIntelIndicatorLastObservedAt()
	ResetThreatIntelIndicatorSource()
	ResetThreatIntelIndicatorSourceUrl()
	ResetThreatIntelIndicatorType()
	ResetThreatIntelIndicatorValue()
	ResetTitle()
	ResetType()
	ResetUpdatedAt()
	ResetUserDefinedValues()
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

// The jsii proxy struct for SecurityhubInsightFiltersOutputReference
type jsiiProxy_SecurityhubInsightFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) AwsAccountId() SecurityhubInsightFiltersAwsAccountIdList {
	var returns SecurityhubInsightFiltersAwsAccountIdList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) AwsAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CompanyName() SecurityhubInsightFiltersCompanyNameList {
	var returns SecurityhubInsightFiltersCompanyNameList
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CompanyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"companyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ComplianceStatus() SecurityhubInsightFiltersComplianceStatusList {
	var returns SecurityhubInsightFiltersComplianceStatusList
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ComplianceStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Confidence() SecurityhubInsightFiltersConfidenceList {
	var returns SecurityhubInsightFiltersConfidenceList
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ConfidenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CreatedAt() SecurityhubInsightFiltersCreatedAtList {
	var returns SecurityhubInsightFiltersCreatedAtList
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CreatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Criticality() SecurityhubInsightFiltersCriticalityList {
	var returns SecurityhubInsightFiltersCriticalityList
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CriticalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Description() SecurityhubInsightFiltersDescriptionList {
	var returns SecurityhubInsightFiltersDescriptionList
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) DescriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsConfidence() SecurityhubInsightFiltersFindingProviderFieldsConfidenceList {
	var returns SecurityhubInsightFiltersFindingProviderFieldsConfidenceList
	_jsii_.Get(
		j,
		"findingProviderFieldsConfidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsConfidenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsConfidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsCriticality() SecurityhubInsightFiltersFindingProviderFieldsCriticalityList {
	var returns SecurityhubInsightFiltersFindingProviderFieldsCriticalityList
	_jsii_.Get(
		j,
		"findingProviderFieldsCriticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsCriticalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsCriticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsId() SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsIdList {
	var returns SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsIdList
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsProductArn() SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArnList {
	var returns SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArnList
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityLabel() SecurityhubInsightFiltersFindingProviderFieldsSeverityLabelList {
	var returns SecurityhubInsightFiltersFindingProviderFieldsSeverityLabelList
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityOriginal() SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginalList {
	var returns SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginalList
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityOriginal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityOriginalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityOriginalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsTypes() SecurityhubInsightFiltersFindingProviderFieldsTypesList {
	var returns SecurityhubInsightFiltersFindingProviderFieldsTypesList
	_jsii_.Get(
		j,
		"findingProviderFieldsTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FirstObservedAt() SecurityhubInsightFiltersFirstObservedAtList {
	var returns SecurityhubInsightFiltersFirstObservedAtList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FirstObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) GeneratorId() SecurityhubInsightFiltersGeneratorIdList {
	var returns SecurityhubInsightFiltersGeneratorIdList
	_jsii_.Get(
		j,
		"generatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) GeneratorIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generatorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Id() SecurityhubInsightFiltersIdList {
	var returns SecurityhubInsightFiltersIdList
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) IdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) InternalValue() *SecurityhubInsightFilters {
	var returns *SecurityhubInsightFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Keyword() SecurityhubInsightFiltersKeywordList {
	var returns SecurityhubInsightFiltersKeywordList
	_jsii_.Get(
		j,
		"keyword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) KeywordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keywordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) LastObservedAt() SecurityhubInsightFiltersLastObservedAtList {
	var returns SecurityhubInsightFiltersLastObservedAtList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) LastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareName() SecurityhubInsightFiltersMalwareNameList {
	var returns SecurityhubInsightFiltersMalwareNameList
	_jsii_.Get(
		j,
		"malwareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwarePath() SecurityhubInsightFiltersMalwarePathList {
	var returns SecurityhubInsightFiltersMalwarePathList
	_jsii_.Get(
		j,
		"malwarePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwarePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwarePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareState() SecurityhubInsightFiltersMalwareStateList {
	var returns SecurityhubInsightFiltersMalwareStateList
	_jsii_.Get(
		j,
		"malwareState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareType() SecurityhubInsightFiltersMalwareTypeList {
	var returns SecurityhubInsightFiltersMalwareTypeList
	_jsii_.Get(
		j,
		"malwareType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationDomain() SecurityhubInsightFiltersNetworkDestinationDomainList {
	var returns SecurityhubInsightFiltersNetworkDestinationDomainList
	_jsii_.Get(
		j,
		"networkDestinationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationIpv4() SecurityhubInsightFiltersNetworkDestinationIpv4List {
	var returns SecurityhubInsightFiltersNetworkDestinationIpv4List
	_jsii_.Get(
		j,
		"networkDestinationIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationIpv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationIpv6() SecurityhubInsightFiltersNetworkDestinationIpv6List {
	var returns SecurityhubInsightFiltersNetworkDestinationIpv6List
	_jsii_.Get(
		j,
		"networkDestinationIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationPort() SecurityhubInsightFiltersNetworkDestinationPortList {
	var returns SecurityhubInsightFiltersNetworkDestinationPortList
	_jsii_.Get(
		j,
		"networkDestinationPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDirection() SecurityhubInsightFiltersNetworkDirectionList {
	var returns SecurityhubInsightFiltersNetworkDirectionList
	_jsii_.Get(
		j,
		"networkDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDirectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkProtocol() SecurityhubInsightFiltersNetworkProtocolList {
	var returns SecurityhubInsightFiltersNetworkProtocolList
	_jsii_.Get(
		j,
		"networkProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceDomain() SecurityhubInsightFiltersNetworkSourceDomainList {
	var returns SecurityhubInsightFiltersNetworkSourceDomainList
	_jsii_.Get(
		j,
		"networkSourceDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceIpv4() SecurityhubInsightFiltersNetworkSourceIpv4List {
	var returns SecurityhubInsightFiltersNetworkSourceIpv4List
	_jsii_.Get(
		j,
		"networkSourceIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceIpv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceIpv6() SecurityhubInsightFiltersNetworkSourceIpv6List {
	var returns SecurityhubInsightFiltersNetworkSourceIpv6List
	_jsii_.Get(
		j,
		"networkSourceIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceMac() SecurityhubInsightFiltersNetworkSourceMacList {
	var returns SecurityhubInsightFiltersNetworkSourceMacList
	_jsii_.Get(
		j,
		"networkSourceMac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceMacInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceMacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourcePort() SecurityhubInsightFiltersNetworkSourcePortList {
	var returns SecurityhubInsightFiltersNetworkSourcePortList
	_jsii_.Get(
		j,
		"networkSourcePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourcePortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourcePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteText() SecurityhubInsightFiltersNoteTextList {
	var returns SecurityhubInsightFiltersNoteTextList
	_jsii_.Get(
		j,
		"noteText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteUpdatedAt() SecurityhubInsightFiltersNoteUpdatedAtList {
	var returns SecurityhubInsightFiltersNoteUpdatedAtList
	_jsii_.Get(
		j,
		"noteUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteUpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteUpdatedBy() SecurityhubInsightFiltersNoteUpdatedByList {
	var returns SecurityhubInsightFiltersNoteUpdatedByList
	_jsii_.Get(
		j,
		"noteUpdatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteUpdatedByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessLaunchedAt() SecurityhubInsightFiltersProcessLaunchedAtList {
	var returns SecurityhubInsightFiltersProcessLaunchedAtList
	_jsii_.Get(
		j,
		"processLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessLaunchedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessName() SecurityhubInsightFiltersProcessNameList {
	var returns SecurityhubInsightFiltersProcessNameList
	_jsii_.Get(
		j,
		"processName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessParentPid() SecurityhubInsightFiltersProcessParentPidList {
	var returns SecurityhubInsightFiltersProcessParentPidList
	_jsii_.Get(
		j,
		"processParentPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessParentPidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processParentPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessPath() SecurityhubInsightFiltersProcessPathList {
	var returns SecurityhubInsightFiltersProcessPathList
	_jsii_.Get(
		j,
		"processPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessPid() SecurityhubInsightFiltersProcessPidList {
	var returns SecurityhubInsightFiltersProcessPidList
	_jsii_.Get(
		j,
		"processPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessPidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessTerminatedAt() SecurityhubInsightFiltersProcessTerminatedAtList {
	var returns SecurityhubInsightFiltersProcessTerminatedAtList
	_jsii_.Get(
		j,
		"processTerminatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessTerminatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processTerminatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductArn() SecurityhubInsightFiltersProductArnList {
	var returns SecurityhubInsightFiltersProductArnList
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductFields() SecurityhubInsightFiltersProductFieldsList {
	var returns SecurityhubInsightFiltersProductFieldsList
	_jsii_.Get(
		j,
		"productFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductName() SecurityhubInsightFiltersProductNameList {
	var returns SecurityhubInsightFiltersProductNameList
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RecommendationText() SecurityhubInsightFiltersRecommendationTextList {
	var returns SecurityhubInsightFiltersRecommendationTextList
	_jsii_.Get(
		j,
		"recommendationText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RecommendationTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recommendationTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RecordState() SecurityhubInsightFiltersRecordStateList {
	var returns SecurityhubInsightFiltersRecordStateList
	_jsii_.Get(
		j,
		"recordState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RecordStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RelatedFindingsId() SecurityhubInsightFiltersRelatedFindingsIdList {
	var returns SecurityhubInsightFiltersRelatedFindingsIdList
	_jsii_.Get(
		j,
		"relatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RelatedFindingsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RelatedFindingsProductArn() SecurityhubInsightFiltersRelatedFindingsProductArnList {
	var returns SecurityhubInsightFiltersRelatedFindingsProductArnList
	_jsii_.Get(
		j,
		"relatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RelatedFindingsProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIamInstanceProfileArn() SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArnList {
	var returns SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArnList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIamInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIamInstanceProfileArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIamInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceImageId() SecurityhubInsightFiltersResourceAwsEc2InstanceImageIdList {
	var returns SecurityhubInsightFiltersResourceAwsEc2InstanceImageIdList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceImageIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpv4Addresses() SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4AddressesList {
	var returns SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4AddressesList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpv4AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv4AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpv6Addresses() SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6AddressesList {
	var returns SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6AddressesList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpv6AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceKeyName() SecurityhubInsightFiltersResourceAwsEc2InstanceKeyNameList {
	var returns SecurityhubInsightFiltersResourceAwsEc2InstanceKeyNameList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceKeyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceLaunchedAt() SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtList {
	var returns SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceLaunchedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceSubnetId() SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetIdList {
	var returns SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetIdList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceSubnetIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceType() SecurityhubInsightFiltersResourceAwsEc2InstanceTypeList {
	var returns SecurityhubInsightFiltersResourceAwsEc2InstanceTypeList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceVpcId() SecurityhubInsightFiltersResourceAwsEc2InstanceVpcIdList {
	var returns SecurityhubInsightFiltersResourceAwsEc2InstanceVpcIdList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceVpcIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyCreatedAt() SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtList {
	var returns SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtList
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyCreatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyCreatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyStatus() SecurityhubInsightFiltersResourceAwsIamAccessKeyStatusList {
	var returns SecurityhubInsightFiltersResourceAwsIamAccessKeyStatusList
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyUserName() SecurityhubInsightFiltersResourceAwsIamAccessKeyUserNameList {
	var returns SecurityhubInsightFiltersResourceAwsIamAccessKeyUserNameList
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyUserNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerId() SecurityhubInsightFiltersResourceAwsS3BucketOwnerIdList {
	var returns SecurityhubInsightFiltersResourceAwsS3BucketOwnerIdList
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerName() SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList {
	var returns SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerImageId() SecurityhubInsightFiltersResourceContainerImageIdList {
	var returns SecurityhubInsightFiltersResourceContainerImageIdList
	_jsii_.Get(
		j,
		"resourceContainerImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerImageIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerImageName() SecurityhubInsightFiltersResourceContainerImageNameList {
	var returns SecurityhubInsightFiltersResourceContainerImageNameList
	_jsii_.Get(
		j,
		"resourceContainerImageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerImageNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerImageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerLaunchedAt() SecurityhubInsightFiltersResourceContainerLaunchedAtList {
	var returns SecurityhubInsightFiltersResourceContainerLaunchedAtList
	_jsii_.Get(
		j,
		"resourceContainerLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerLaunchedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerName() SecurityhubInsightFiltersResourceContainerNameList {
	var returns SecurityhubInsightFiltersResourceContainerNameList
	_jsii_.Get(
		j,
		"resourceContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceDetailsOther() SecurityhubInsightFiltersResourceDetailsOtherList {
	var returns SecurityhubInsightFiltersResourceDetailsOtherList
	_jsii_.Get(
		j,
		"resourceDetailsOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceDetailsOtherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceDetailsOtherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceId() SecurityhubInsightFiltersResourceIdList {
	var returns SecurityhubInsightFiltersResourceIdList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourcePartition() SecurityhubInsightFiltersResourcePartitionList {
	var returns SecurityhubInsightFiltersResourcePartitionList
	_jsii_.Get(
		j,
		"resourcePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourcePartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcePartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceRegion() SecurityhubInsightFiltersResourceRegionList {
	var returns SecurityhubInsightFiltersResourceRegionList
	_jsii_.Get(
		j,
		"resourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceRegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceTags() SecurityhubInsightFiltersResourceTagsList {
	var returns SecurityhubInsightFiltersResourceTagsList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceType() SecurityhubInsightFiltersResourceTypeList {
	var returns SecurityhubInsightFiltersResourceTypeList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SeverityLabel() SecurityhubInsightFiltersSeverityLabelList {
	var returns SecurityhubInsightFiltersSeverityLabelList
	_jsii_.Get(
		j,
		"severityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SeverityLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SourceUrl() SecurityhubInsightFiltersSourceUrlList {
	var returns SecurityhubInsightFiltersSourceUrlList
	_jsii_.Get(
		j,
		"sourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SourceUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorCategory() SecurityhubInsightFiltersThreatIntelIndicatorCategoryList {
	var returns SecurityhubInsightFiltersThreatIntelIndicatorCategoryList
	_jsii_.Get(
		j,
		"threatIntelIndicatorCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorCategoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorLastObservedAt() SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtList {
	var returns SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtList
	_jsii_.Get(
		j,
		"threatIntelIndicatorLastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorLastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorLastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSource() SecurityhubInsightFiltersThreatIntelIndicatorSourceList {
	var returns SecurityhubInsightFiltersThreatIntelIndicatorSourceList
	_jsii_.Get(
		j,
		"threatIntelIndicatorSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSourceUrl() SecurityhubInsightFiltersThreatIntelIndicatorSourceUrlList {
	var returns SecurityhubInsightFiltersThreatIntelIndicatorSourceUrlList
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSourceUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorType() SecurityhubInsightFiltersThreatIntelIndicatorTypeList {
	var returns SecurityhubInsightFiltersThreatIntelIndicatorTypeList
	_jsii_.Get(
		j,
		"threatIntelIndicatorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorValue() SecurityhubInsightFiltersThreatIntelIndicatorValueList {
	var returns SecurityhubInsightFiltersThreatIntelIndicatorValueList
	_jsii_.Get(
		j,
		"threatIntelIndicatorValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Title() SecurityhubInsightFiltersTitleList {
	var returns SecurityhubInsightFiltersTitleList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) TitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Type() SecurityhubInsightFiltersTypeList {
	var returns SecurityhubInsightFiltersTypeList
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) TypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) UpdatedAt() SecurityhubInsightFiltersUpdatedAtList {
	var returns SecurityhubInsightFiltersUpdatedAtList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) UpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) UserDefinedValues() SecurityhubInsightFiltersUserDefinedValuesList {
	var returns SecurityhubInsightFiltersUserDefinedValuesList
	_jsii_.Get(
		j,
		"userDefinedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) UserDefinedValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefinedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) VerificationState() SecurityhubInsightFiltersVerificationStateList {
	var returns SecurityhubInsightFiltersVerificationStateList
	_jsii_.Get(
		j,
		"verificationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) VerificationStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verificationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) WorkflowStatus() SecurityhubInsightFiltersWorkflowStatusList {
	var returns SecurityhubInsightFiltersWorkflowStatusList
	_jsii_.Get(
		j,
		"workflowStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) WorkflowStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowStatusInput",
		&returns,
	)
	return returns
}


func NewSecurityhubInsightFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SecurityhubInsightFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityhubInsightFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityhubInsightFiltersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubInsight.SecurityhubInsightFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersOutputReference_Override(s SecurityhubInsightFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubInsight.SecurityhubInsightFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference)SetInternalValue(val *SecurityhubInsightFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutAwsAccountId(value interface{}) {
	if err := s.validatePutAwsAccountIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAwsAccountId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutCompanyName(value interface{}) {
	if err := s.validatePutCompanyNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCompanyName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutComplianceStatus(value interface{}) {
	if err := s.validatePutComplianceStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putComplianceStatus",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutConfidence(value interface{}) {
	if err := s.validatePutConfidenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putConfidence",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutCreatedAt(value interface{}) {
	if err := s.validatePutCreatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCreatedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutCriticality(value interface{}) {
	if err := s.validatePutCriticalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCriticality",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutDescription(value interface{}) {
	if err := s.validatePutDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDescription",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutFindingProviderFieldsConfidence(value interface{}) {
	if err := s.validatePutFindingProviderFieldsConfidenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFindingProviderFieldsConfidence",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutFindingProviderFieldsCriticality(value interface{}) {
	if err := s.validatePutFindingProviderFieldsCriticalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFindingProviderFieldsCriticality",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutFindingProviderFieldsRelatedFindingsId(value interface{}) {
	if err := s.validatePutFindingProviderFieldsRelatedFindingsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFindingProviderFieldsRelatedFindingsId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutFindingProviderFieldsRelatedFindingsProductArn(value interface{}) {
	if err := s.validatePutFindingProviderFieldsRelatedFindingsProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFindingProviderFieldsRelatedFindingsProductArn",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutFindingProviderFieldsSeverityLabel(value interface{}) {
	if err := s.validatePutFindingProviderFieldsSeverityLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFindingProviderFieldsSeverityLabel",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutFindingProviderFieldsSeverityOriginal(value interface{}) {
	if err := s.validatePutFindingProviderFieldsSeverityOriginalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFindingProviderFieldsSeverityOriginal",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutFindingProviderFieldsTypes(value interface{}) {
	if err := s.validatePutFindingProviderFieldsTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFindingProviderFieldsTypes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutFirstObservedAt(value interface{}) {
	if err := s.validatePutFirstObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFirstObservedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutGeneratorId(value interface{}) {
	if err := s.validatePutGeneratorIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGeneratorId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutId(value interface{}) {
	if err := s.validatePutIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutKeyword(value interface{}) {
	if err := s.validatePutKeywordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKeyword",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutLastObservedAt(value interface{}) {
	if err := s.validatePutLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLastObservedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutMalwareName(value interface{}) {
	if err := s.validatePutMalwareNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMalwareName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutMalwarePath(value interface{}) {
	if err := s.validatePutMalwarePathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMalwarePath",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutMalwareState(value interface{}) {
	if err := s.validatePutMalwareStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMalwareState",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutMalwareType(value interface{}) {
	if err := s.validatePutMalwareTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMalwareType",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNetworkDestinationDomain(value interface{}) {
	if err := s.validatePutNetworkDestinationDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkDestinationDomain",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNetworkDestinationIpv4(value interface{}) {
	if err := s.validatePutNetworkDestinationIpv4Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkDestinationIpv4",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNetworkDestinationIpv6(value interface{}) {
	if err := s.validatePutNetworkDestinationIpv6Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkDestinationIpv6",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNetworkDestinationPort(value interface{}) {
	if err := s.validatePutNetworkDestinationPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkDestinationPort",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNetworkDirection(value interface{}) {
	if err := s.validatePutNetworkDirectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkDirection",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNetworkProtocol(value interface{}) {
	if err := s.validatePutNetworkProtocolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkProtocol",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNetworkSourceDomain(value interface{}) {
	if err := s.validatePutNetworkSourceDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkSourceDomain",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNetworkSourceIpv4(value interface{}) {
	if err := s.validatePutNetworkSourceIpv4Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkSourceIpv4",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNetworkSourceIpv6(value interface{}) {
	if err := s.validatePutNetworkSourceIpv6Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkSourceIpv6",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNetworkSourceMac(value interface{}) {
	if err := s.validatePutNetworkSourceMacParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkSourceMac",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNetworkSourcePort(value interface{}) {
	if err := s.validatePutNetworkSourcePortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkSourcePort",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNoteText(value interface{}) {
	if err := s.validatePutNoteTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNoteText",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNoteUpdatedAt(value interface{}) {
	if err := s.validatePutNoteUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNoteUpdatedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutNoteUpdatedBy(value interface{}) {
	if err := s.validatePutNoteUpdatedByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNoteUpdatedBy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutProcessLaunchedAt(value interface{}) {
	if err := s.validatePutProcessLaunchedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProcessLaunchedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutProcessName(value interface{}) {
	if err := s.validatePutProcessNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProcessName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutProcessParentPid(value interface{}) {
	if err := s.validatePutProcessParentPidParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProcessParentPid",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutProcessPath(value interface{}) {
	if err := s.validatePutProcessPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProcessPath",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutProcessPid(value interface{}) {
	if err := s.validatePutProcessPidParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProcessPid",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutProcessTerminatedAt(value interface{}) {
	if err := s.validatePutProcessTerminatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProcessTerminatedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutProductArn(value interface{}) {
	if err := s.validatePutProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProductArn",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutProductFields(value interface{}) {
	if err := s.validatePutProductFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProductFields",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutProductName(value interface{}) {
	if err := s.validatePutProductNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProductName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutRecommendationText(value interface{}) {
	if err := s.validatePutRecommendationTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRecommendationText",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutRecordState(value interface{}) {
	if err := s.validatePutRecordStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRecordState",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutRelatedFindingsId(value interface{}) {
	if err := s.validatePutRelatedFindingsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRelatedFindingsId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutRelatedFindingsProductArn(value interface{}) {
	if err := s.validatePutRelatedFindingsProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRelatedFindingsProductArn",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsEc2InstanceIamInstanceProfileArn(value interface{}) {
	if err := s.validatePutResourceAwsEc2InstanceIamInstanceProfileArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsEc2InstanceIamInstanceProfileArn",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsEc2InstanceImageId(value interface{}) {
	if err := s.validatePutResourceAwsEc2InstanceImageIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsEc2InstanceImageId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsEc2InstanceIpv4Addresses(value interface{}) {
	if err := s.validatePutResourceAwsEc2InstanceIpv4AddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsEc2InstanceIpv4Addresses",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsEc2InstanceIpv6Addresses(value interface{}) {
	if err := s.validatePutResourceAwsEc2InstanceIpv6AddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsEc2InstanceIpv6Addresses",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsEc2InstanceKeyName(value interface{}) {
	if err := s.validatePutResourceAwsEc2InstanceKeyNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsEc2InstanceKeyName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsEc2InstanceLaunchedAt(value interface{}) {
	if err := s.validatePutResourceAwsEc2InstanceLaunchedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsEc2InstanceLaunchedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsEc2InstanceSubnetId(value interface{}) {
	if err := s.validatePutResourceAwsEc2InstanceSubnetIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsEc2InstanceSubnetId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsEc2InstanceType(value interface{}) {
	if err := s.validatePutResourceAwsEc2InstanceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsEc2InstanceType",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsEc2InstanceVpcId(value interface{}) {
	if err := s.validatePutResourceAwsEc2InstanceVpcIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsEc2InstanceVpcId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsIamAccessKeyCreatedAt(value interface{}) {
	if err := s.validatePutResourceAwsIamAccessKeyCreatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsIamAccessKeyCreatedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsIamAccessKeyStatus(value interface{}) {
	if err := s.validatePutResourceAwsIamAccessKeyStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsIamAccessKeyStatus",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsIamAccessKeyUserName(value interface{}) {
	if err := s.validatePutResourceAwsIamAccessKeyUserNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsIamAccessKeyUserName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsS3BucketOwnerId(value interface{}) {
	if err := s.validatePutResourceAwsS3BucketOwnerIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsS3BucketOwnerId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceAwsS3BucketOwnerName(value interface{}) {
	if err := s.validatePutResourceAwsS3BucketOwnerNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceAwsS3BucketOwnerName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceContainerImageId(value interface{}) {
	if err := s.validatePutResourceContainerImageIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceContainerImageId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceContainerImageName(value interface{}) {
	if err := s.validatePutResourceContainerImageNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceContainerImageName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceContainerLaunchedAt(value interface{}) {
	if err := s.validatePutResourceContainerLaunchedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceContainerLaunchedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceContainerName(value interface{}) {
	if err := s.validatePutResourceContainerNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceContainerName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceDetailsOther(value interface{}) {
	if err := s.validatePutResourceDetailsOtherParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceDetailsOther",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceId(value interface{}) {
	if err := s.validatePutResourceIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceId",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourcePartition(value interface{}) {
	if err := s.validatePutResourcePartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourcePartition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceRegion(value interface{}) {
	if err := s.validatePutResourceRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceRegion",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceTags(value interface{}) {
	if err := s.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutResourceType(value interface{}) {
	if err := s.validatePutResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceType",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutSeverityLabel(value interface{}) {
	if err := s.validatePutSeverityLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSeverityLabel",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutSourceUrl(value interface{}) {
	if err := s.validatePutSourceUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSourceUrl",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutThreatIntelIndicatorCategory(value interface{}) {
	if err := s.validatePutThreatIntelIndicatorCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThreatIntelIndicatorCategory",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutThreatIntelIndicatorLastObservedAt(value interface{}) {
	if err := s.validatePutThreatIntelIndicatorLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThreatIntelIndicatorLastObservedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutThreatIntelIndicatorSource(value interface{}) {
	if err := s.validatePutThreatIntelIndicatorSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThreatIntelIndicatorSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutThreatIntelIndicatorSourceUrl(value interface{}) {
	if err := s.validatePutThreatIntelIndicatorSourceUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThreatIntelIndicatorSourceUrl",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutThreatIntelIndicatorType(value interface{}) {
	if err := s.validatePutThreatIntelIndicatorTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThreatIntelIndicatorType",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutThreatIntelIndicatorValue(value interface{}) {
	if err := s.validatePutThreatIntelIndicatorValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThreatIntelIndicatorValue",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutTitle(value interface{}) {
	if err := s.validatePutTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTitle",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutType(value interface{}) {
	if err := s.validatePutTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putType",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutUpdatedAt(value interface{}) {
	if err := s.validatePutUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUpdatedAt",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutUserDefinedValues(value interface{}) {
	if err := s.validatePutUserDefinedValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUserDefinedValues",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutVerificationState(value interface{}) {
	if err := s.validatePutVerificationStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVerificationState",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) PutWorkflowStatus(value interface{}) {
	if err := s.validatePutWorkflowStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWorkflowStatus",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetCompanyName() {
	_jsii_.InvokeVoid(
		s,
		"resetCompanyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetComplianceStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetComplianceStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetConfidence() {
	_jsii_.InvokeVoid(
		s,
		"resetConfidence",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		s,
		"resetCriticality",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsConfidence() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsConfidence",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsCriticality() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsCriticality",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsRelatedFindingsId() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsRelatedFindingsId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsRelatedFindingsProductArn() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsRelatedFindingsProductArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsSeverityLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsSeverityLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsSeverityOriginal() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsSeverityOriginal",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFirstObservedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetFirstObservedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetGeneratorId() {
	_jsii_.InvokeVoid(
		s,
		"resetGeneratorId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetKeyword() {
	_jsii_.InvokeVoid(
		s,
		"resetKeyword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetLastObservedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetLastObservedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetMalwareName() {
	_jsii_.InvokeVoid(
		s,
		"resetMalwareName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetMalwarePath() {
	_jsii_.InvokeVoid(
		s,
		"resetMalwarePath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetMalwareState() {
	_jsii_.InvokeVoid(
		s,
		"resetMalwareState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetMalwareType() {
	_jsii_.InvokeVoid(
		s,
		"resetMalwareType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkDestinationDomain() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkDestinationDomain",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkDestinationIpv4() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkDestinationIpv4",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkDestinationIpv6() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkDestinationIpv6",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkDestinationPort() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkDestinationPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkDirection() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkDirection",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkProtocol() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkProtocol",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkSourceDomain() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkSourceDomain",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkSourceIpv4() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkSourceIpv4",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkSourceIpv6() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkSourceIpv6",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkSourceMac() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkSourceMac",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkSourcePort() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkSourcePort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNoteText() {
	_jsii_.InvokeVoid(
		s,
		"resetNoteText",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNoteUpdatedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetNoteUpdatedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNoteUpdatedBy() {
	_jsii_.InvokeVoid(
		s,
		"resetNoteUpdatedBy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProcessLaunchedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessLaunchedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProcessName() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProcessParentPid() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessParentPid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProcessPath() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProcessPid() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessPid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProcessTerminatedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessTerminatedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProductArn() {
	_jsii_.InvokeVoid(
		s,
		"resetProductArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProductFields() {
	_jsii_.InvokeVoid(
		s,
		"resetProductFields",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProductName() {
	_jsii_.InvokeVoid(
		s,
		"resetProductName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetRecommendationText() {
	_jsii_.InvokeVoid(
		s,
		"resetRecommendationText",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetRecordState() {
	_jsii_.InvokeVoid(
		s,
		"resetRecordState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetRelatedFindingsId() {
	_jsii_.InvokeVoid(
		s,
		"resetRelatedFindingsId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetRelatedFindingsProductArn() {
	_jsii_.InvokeVoid(
		s,
		"resetRelatedFindingsProductArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceIamInstanceProfileArn() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceIamInstanceProfileArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceImageId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceImageId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceIpv4Addresses() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceIpv4Addresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceIpv6Addresses() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceIpv6Addresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceKeyName() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceKeyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceLaunchedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceLaunchedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceSubnetId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceSubnetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceType() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceVpcId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceVpcId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsIamAccessKeyCreatedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsIamAccessKeyCreatedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsIamAccessKeyStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsIamAccessKeyStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsIamAccessKeyUserName() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsIamAccessKeyUserName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsS3BucketOwnerId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsS3BucketOwnerId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsS3BucketOwnerName() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsS3BucketOwnerName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceContainerImageId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceContainerImageId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceContainerImageName() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceContainerImageName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceContainerLaunchedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceContainerLaunchedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceContainerName() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceContainerName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceDetailsOther() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceDetailsOther",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourcePartition() {
	_jsii_.InvokeVoid(
		s,
		"resetResourcePartition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetSeverityLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetSeverityLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetSourceUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetThreatIntelIndicatorCategory() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatIntelIndicatorCategory",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetThreatIntelIndicatorLastObservedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatIntelIndicatorLastObservedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetThreatIntelIndicatorSource() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatIntelIndicatorSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetThreatIntelIndicatorSourceUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatIntelIndicatorSourceUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetThreatIntelIndicatorType() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatIntelIndicatorType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetThreatIntelIndicatorValue() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatIntelIndicatorValue",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		s,
		"resetTitle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetUserDefinedValues() {
	_jsii_.InvokeVoid(
		s,
		"resetUserDefinedValues",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetVerificationState() {
	_jsii_.InvokeVoid(
		s,
		"resetVerificationState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetWorkflowStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkflowStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

