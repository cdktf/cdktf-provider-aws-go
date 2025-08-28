// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector2filter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/inspector2filter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Inspector2FilterFilterCriteriaOutputReference interface {
	cdktf.ComplexObject
	AwsAccountId() Inspector2FilterFilterCriteriaAwsAccountIdList
	AwsAccountIdInput() interface{}
	CodeRepositoryProjectName() Inspector2FilterFilterCriteriaCodeRepositoryProjectNameList
	CodeRepositoryProjectNameInput() interface{}
	CodeRepositoryProviderType() Inspector2FilterFilterCriteriaCodeRepositoryProviderTypeList
	CodeRepositoryProviderTypeInput() interface{}
	CodeVulnerabilityDetectorName() Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorNameList
	CodeVulnerabilityDetectorNameInput() interface{}
	CodeVulnerabilityDetectorTags() Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorTagsList
	CodeVulnerabilityDetectorTagsInput() interface{}
	CodeVulnerabilityFilePath() Inspector2FilterFilterCriteriaCodeVulnerabilityFilePathList
	CodeVulnerabilityFilePathInput() interface{}
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
	ComponentId() Inspector2FilterFilterCriteriaComponentIdList
	ComponentIdInput() interface{}
	ComponentType() Inspector2FilterFilterCriteriaComponentTypeList
	ComponentTypeInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Ec2InstanceImageId() Inspector2FilterFilterCriteriaEc2InstanceImageIdList
	Ec2InstanceImageIdInput() interface{}
	Ec2InstanceSubnetId() Inspector2FilterFilterCriteriaEc2InstanceSubnetIdList
	Ec2InstanceSubnetIdInput() interface{}
	Ec2InstanceVpcId() Inspector2FilterFilterCriteriaEc2InstanceVpcIdList
	Ec2InstanceVpcIdInput() interface{}
	EcrImageArchitecture() Inspector2FilterFilterCriteriaEcrImageArchitectureList
	EcrImageArchitectureInput() interface{}
	EcrImageHash() Inspector2FilterFilterCriteriaEcrImageHashList
	EcrImageHashInput() interface{}
	EcrImageInUseCount() Inspector2FilterFilterCriteriaEcrImageInUseCountList
	EcrImageInUseCountInput() interface{}
	EcrImageLastInUseAt() Inspector2FilterFilterCriteriaEcrImageLastInUseAtList
	EcrImageLastInUseAtInput() interface{}
	EcrImagePushedAt() Inspector2FilterFilterCriteriaEcrImagePushedAtList
	EcrImagePushedAtInput() interface{}
	EcrImageRegistry() Inspector2FilterFilterCriteriaEcrImageRegistryList
	EcrImageRegistryInput() interface{}
	EcrImageRepositoryName() Inspector2FilterFilterCriteriaEcrImageRepositoryNameList
	EcrImageRepositoryNameInput() interface{}
	EcrImageTags() Inspector2FilterFilterCriteriaEcrImageTagsList
	EcrImageTagsInput() interface{}
	EpssScore() Inspector2FilterFilterCriteriaEpssScoreList
	EpssScoreInput() interface{}
	ExploitAvailable() Inspector2FilterFilterCriteriaExploitAvailableList
	ExploitAvailableInput() interface{}
	FindingArn() Inspector2FilterFilterCriteriaFindingArnList
	FindingArnInput() interface{}
	FindingStatus() Inspector2FilterFilterCriteriaFindingStatusList
	FindingStatusInput() interface{}
	FindingType() Inspector2FilterFilterCriteriaFindingTypeList
	FindingTypeInput() interface{}
	FirstObservedAt() Inspector2FilterFilterCriteriaFirstObservedAtList
	FirstObservedAtInput() interface{}
	FixAvailable() Inspector2FilterFilterCriteriaFixAvailableList
	FixAvailableInput() interface{}
	// Experimental.
	Fqn() *string
	InspectorScore() Inspector2FilterFilterCriteriaInspectorScoreList
	InspectorScoreInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LambdaFunctionExecutionRoleArn() Inspector2FilterFilterCriteriaLambdaFunctionExecutionRoleArnList
	LambdaFunctionExecutionRoleArnInput() interface{}
	LambdaFunctionLastModifiedAt() Inspector2FilterFilterCriteriaLambdaFunctionLastModifiedAtList
	LambdaFunctionLastModifiedAtInput() interface{}
	LambdaFunctionLayers() Inspector2FilterFilterCriteriaLambdaFunctionLayersList
	LambdaFunctionLayersInput() interface{}
	LambdaFunctionName() Inspector2FilterFilterCriteriaLambdaFunctionNameList
	LambdaFunctionNameInput() interface{}
	LambdaFunctionRuntime() Inspector2FilterFilterCriteriaLambdaFunctionRuntimeList
	LambdaFunctionRuntimeInput() interface{}
	LastObservedAt() Inspector2FilterFilterCriteriaLastObservedAtList
	LastObservedAtInput() interface{}
	NetworkProtocol() Inspector2FilterFilterCriteriaNetworkProtocolList
	NetworkProtocolInput() interface{}
	PortRange() Inspector2FilterFilterCriteriaPortRangeList
	PortRangeInput() interface{}
	RelatedVulnerabilities() Inspector2FilterFilterCriteriaRelatedVulnerabilitiesList
	RelatedVulnerabilitiesInput() interface{}
	ResourceId() Inspector2FilterFilterCriteriaResourceIdList
	ResourceIdInput() interface{}
	ResourceTags() Inspector2FilterFilterCriteriaResourceTagsList
	ResourceTagsInput() interface{}
	ResourceType() Inspector2FilterFilterCriteriaResourceTypeList
	ResourceTypeInput() interface{}
	Severity() Inspector2FilterFilterCriteriaSeverityList
	SeverityInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() Inspector2FilterFilterCriteriaTitleList
	TitleInput() interface{}
	UpdatedAt() Inspector2FilterFilterCriteriaUpdatedAtList
	UpdatedAtInput() interface{}
	VendorSeverity() Inspector2FilterFilterCriteriaVendorSeverityList
	VendorSeverityInput() interface{}
	VulnerabilityId() Inspector2FilterFilterCriteriaVulnerabilityIdList
	VulnerabilityIdInput() interface{}
	VulnerabilitySource() Inspector2FilterFilterCriteriaVulnerabilitySourceList
	VulnerabilitySourceInput() interface{}
	VulnerablePackages() Inspector2FilterFilterCriteriaVulnerablePackagesList
	VulnerablePackagesInput() interface{}
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
	PutCodeRepositoryProjectName(value interface{})
	PutCodeRepositoryProviderType(value interface{})
	PutCodeVulnerabilityDetectorName(value interface{})
	PutCodeVulnerabilityDetectorTags(value interface{})
	PutCodeVulnerabilityFilePath(value interface{})
	PutComponentId(value interface{})
	PutComponentType(value interface{})
	PutEc2InstanceImageId(value interface{})
	PutEc2InstanceSubnetId(value interface{})
	PutEc2InstanceVpcId(value interface{})
	PutEcrImageArchitecture(value interface{})
	PutEcrImageHash(value interface{})
	PutEcrImageInUseCount(value interface{})
	PutEcrImageLastInUseAt(value interface{})
	PutEcrImagePushedAt(value interface{})
	PutEcrImageRegistry(value interface{})
	PutEcrImageRepositoryName(value interface{})
	PutEcrImageTags(value interface{})
	PutEpssScore(value interface{})
	PutExploitAvailable(value interface{})
	PutFindingArn(value interface{})
	PutFindingStatus(value interface{})
	PutFindingType(value interface{})
	PutFirstObservedAt(value interface{})
	PutFixAvailable(value interface{})
	PutInspectorScore(value interface{})
	PutLambdaFunctionExecutionRoleArn(value interface{})
	PutLambdaFunctionLastModifiedAt(value interface{})
	PutLambdaFunctionLayers(value interface{})
	PutLambdaFunctionName(value interface{})
	PutLambdaFunctionRuntime(value interface{})
	PutLastObservedAt(value interface{})
	PutNetworkProtocol(value interface{})
	PutPortRange(value interface{})
	PutRelatedVulnerabilities(value interface{})
	PutResourceId(value interface{})
	PutResourceTags(value interface{})
	PutResourceType(value interface{})
	PutSeverity(value interface{})
	PutTitle(value interface{})
	PutUpdatedAt(value interface{})
	PutVendorSeverity(value interface{})
	PutVulnerabilityId(value interface{})
	PutVulnerabilitySource(value interface{})
	PutVulnerablePackages(value interface{})
	ResetAwsAccountId()
	ResetCodeRepositoryProjectName()
	ResetCodeRepositoryProviderType()
	ResetCodeVulnerabilityDetectorName()
	ResetCodeVulnerabilityDetectorTags()
	ResetCodeVulnerabilityFilePath()
	ResetComponentId()
	ResetComponentType()
	ResetEc2InstanceImageId()
	ResetEc2InstanceSubnetId()
	ResetEc2InstanceVpcId()
	ResetEcrImageArchitecture()
	ResetEcrImageHash()
	ResetEcrImageInUseCount()
	ResetEcrImageLastInUseAt()
	ResetEcrImagePushedAt()
	ResetEcrImageRegistry()
	ResetEcrImageRepositoryName()
	ResetEcrImageTags()
	ResetEpssScore()
	ResetExploitAvailable()
	ResetFindingArn()
	ResetFindingStatus()
	ResetFindingType()
	ResetFirstObservedAt()
	ResetFixAvailable()
	ResetInspectorScore()
	ResetLambdaFunctionExecutionRoleArn()
	ResetLambdaFunctionLastModifiedAt()
	ResetLambdaFunctionLayers()
	ResetLambdaFunctionName()
	ResetLambdaFunctionRuntime()
	ResetLastObservedAt()
	ResetNetworkProtocol()
	ResetPortRange()
	ResetRelatedVulnerabilities()
	ResetResourceId()
	ResetResourceTags()
	ResetResourceType()
	ResetSeverity()
	ResetTitle()
	ResetUpdatedAt()
	ResetVendorSeverity()
	ResetVulnerabilityId()
	ResetVulnerabilitySource()
	ResetVulnerablePackages()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Inspector2FilterFilterCriteriaOutputReference
type jsiiProxy_Inspector2FilterFilterCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) AwsAccountId() Inspector2FilterFilterCriteriaAwsAccountIdList {
	var returns Inspector2FilterFilterCriteriaAwsAccountIdList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) AwsAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) CodeRepositoryProjectName() Inspector2FilterFilterCriteriaCodeRepositoryProjectNameList {
	var returns Inspector2FilterFilterCriteriaCodeRepositoryProjectNameList
	_jsii_.Get(
		j,
		"codeRepositoryProjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) CodeRepositoryProjectNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryProjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) CodeRepositoryProviderType() Inspector2FilterFilterCriteriaCodeRepositoryProviderTypeList {
	var returns Inspector2FilterFilterCriteriaCodeRepositoryProviderTypeList
	_jsii_.Get(
		j,
		"codeRepositoryProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) CodeRepositoryProviderTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) CodeVulnerabilityDetectorName() Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorNameList {
	var returns Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorNameList
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) CodeVulnerabilityDetectorNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) CodeVulnerabilityDetectorTags() Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorTagsList {
	var returns Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorTagsList
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) CodeVulnerabilityDetectorTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) CodeVulnerabilityFilePath() Inspector2FilterFilterCriteriaCodeVulnerabilityFilePathList {
	var returns Inspector2FilterFilterCriteriaCodeVulnerabilityFilePathList
	_jsii_.Get(
		j,
		"codeVulnerabilityFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) CodeVulnerabilityFilePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeVulnerabilityFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ComponentId() Inspector2FilterFilterCriteriaComponentIdList {
	var returns Inspector2FilterFilterCriteriaComponentIdList
	_jsii_.Get(
		j,
		"componentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ComponentIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ComponentType() Inspector2FilterFilterCriteriaComponentTypeList {
	var returns Inspector2FilterFilterCriteriaComponentTypeList
	_jsii_.Get(
		j,
		"componentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ComponentTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) Ec2InstanceImageId() Inspector2FilterFilterCriteriaEc2InstanceImageIdList {
	var returns Inspector2FilterFilterCriteriaEc2InstanceImageIdList
	_jsii_.Get(
		j,
		"ec2InstanceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) Ec2InstanceImageIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) Ec2InstanceSubnetId() Inspector2FilterFilterCriteriaEc2InstanceSubnetIdList {
	var returns Inspector2FilterFilterCriteriaEc2InstanceSubnetIdList
	_jsii_.Get(
		j,
		"ec2InstanceSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) Ec2InstanceSubnetIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) Ec2InstanceVpcId() Inspector2FilterFilterCriteriaEc2InstanceVpcIdList {
	var returns Inspector2FilterFilterCriteriaEc2InstanceVpcIdList
	_jsii_.Get(
		j,
		"ec2InstanceVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) Ec2InstanceVpcIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageArchitecture() Inspector2FilterFilterCriteriaEcrImageArchitectureList {
	var returns Inspector2FilterFilterCriteriaEcrImageArchitectureList
	_jsii_.Get(
		j,
		"ecrImageArchitecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageArchitectureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageArchitectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageHash() Inspector2FilterFilterCriteriaEcrImageHashList {
	var returns Inspector2FilterFilterCriteriaEcrImageHashList
	_jsii_.Get(
		j,
		"ecrImageHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageHashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageInUseCount() Inspector2FilterFilterCriteriaEcrImageInUseCountList {
	var returns Inspector2FilterFilterCriteriaEcrImageInUseCountList
	_jsii_.Get(
		j,
		"ecrImageInUseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageInUseCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageInUseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageLastInUseAt() Inspector2FilterFilterCriteriaEcrImageLastInUseAtList {
	var returns Inspector2FilterFilterCriteriaEcrImageLastInUseAtList
	_jsii_.Get(
		j,
		"ecrImageLastInUseAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageLastInUseAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageLastInUseAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImagePushedAt() Inspector2FilterFilterCriteriaEcrImagePushedAtList {
	var returns Inspector2FilterFilterCriteriaEcrImagePushedAtList
	_jsii_.Get(
		j,
		"ecrImagePushedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImagePushedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImagePushedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageRegistry() Inspector2FilterFilterCriteriaEcrImageRegistryList {
	var returns Inspector2FilterFilterCriteriaEcrImageRegistryList
	_jsii_.Get(
		j,
		"ecrImageRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageRegistryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageRepositoryName() Inspector2FilterFilterCriteriaEcrImageRepositoryNameList {
	var returns Inspector2FilterFilterCriteriaEcrImageRepositoryNameList
	_jsii_.Get(
		j,
		"ecrImageRepositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageRepositoryNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageRepositoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageTags() Inspector2FilterFilterCriteriaEcrImageTagsList {
	var returns Inspector2FilterFilterCriteriaEcrImageTagsList
	_jsii_.Get(
		j,
		"ecrImageTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EcrImageTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EpssScore() Inspector2FilterFilterCriteriaEpssScoreList {
	var returns Inspector2FilterFilterCriteriaEpssScoreList
	_jsii_.Get(
		j,
		"epssScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) EpssScoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"epssScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ExploitAvailable() Inspector2FilterFilterCriteriaExploitAvailableList {
	var returns Inspector2FilterFilterCriteriaExploitAvailableList
	_jsii_.Get(
		j,
		"exploitAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ExploitAvailableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exploitAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) FindingArn() Inspector2FilterFilterCriteriaFindingArnList {
	var returns Inspector2FilterFilterCriteriaFindingArnList
	_jsii_.Get(
		j,
		"findingArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) FindingArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) FindingStatus() Inspector2FilterFilterCriteriaFindingStatusList {
	var returns Inspector2FilterFilterCriteriaFindingStatusList
	_jsii_.Get(
		j,
		"findingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) FindingStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) FindingType() Inspector2FilterFilterCriteriaFindingTypeList {
	var returns Inspector2FilterFilterCriteriaFindingTypeList
	_jsii_.Get(
		j,
		"findingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) FindingTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) FirstObservedAt() Inspector2FilterFilterCriteriaFirstObservedAtList {
	var returns Inspector2FilterFilterCriteriaFirstObservedAtList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) FirstObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) FixAvailable() Inspector2FilterFilterCriteriaFixAvailableList {
	var returns Inspector2FilterFilterCriteriaFixAvailableList
	_jsii_.Get(
		j,
		"fixAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) FixAvailableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) InspectorScore() Inspector2FilterFilterCriteriaInspectorScoreList {
	var returns Inspector2FilterFilterCriteriaInspectorScoreList
	_jsii_.Get(
		j,
		"inspectorScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) InspectorScoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inspectorScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) LambdaFunctionExecutionRoleArn() Inspector2FilterFilterCriteriaLambdaFunctionExecutionRoleArnList {
	var returns Inspector2FilterFilterCriteriaLambdaFunctionExecutionRoleArnList
	_jsii_.Get(
		j,
		"lambdaFunctionExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) LambdaFunctionExecutionRoleArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) LambdaFunctionLastModifiedAt() Inspector2FilterFilterCriteriaLambdaFunctionLastModifiedAtList {
	var returns Inspector2FilterFilterCriteriaLambdaFunctionLastModifiedAtList
	_jsii_.Get(
		j,
		"lambdaFunctionLastModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) LambdaFunctionLastModifiedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionLastModifiedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) LambdaFunctionLayers() Inspector2FilterFilterCriteriaLambdaFunctionLayersList {
	var returns Inspector2FilterFilterCriteriaLambdaFunctionLayersList
	_jsii_.Get(
		j,
		"lambdaFunctionLayers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) LambdaFunctionLayersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionLayersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) LambdaFunctionName() Inspector2FilterFilterCriteriaLambdaFunctionNameList {
	var returns Inspector2FilterFilterCriteriaLambdaFunctionNameList
	_jsii_.Get(
		j,
		"lambdaFunctionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) LambdaFunctionNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) LambdaFunctionRuntime() Inspector2FilterFilterCriteriaLambdaFunctionRuntimeList {
	var returns Inspector2FilterFilterCriteriaLambdaFunctionRuntimeList
	_jsii_.Get(
		j,
		"lambdaFunctionRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) LambdaFunctionRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) LastObservedAt() Inspector2FilterFilterCriteriaLastObservedAtList {
	var returns Inspector2FilterFilterCriteriaLastObservedAtList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) LastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) NetworkProtocol() Inspector2FilterFilterCriteriaNetworkProtocolList {
	var returns Inspector2FilterFilterCriteriaNetworkProtocolList
	_jsii_.Get(
		j,
		"networkProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) NetworkProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PortRange() Inspector2FilterFilterCriteriaPortRangeList {
	var returns Inspector2FilterFilterCriteriaPortRangeList
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PortRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) RelatedVulnerabilities() Inspector2FilterFilterCriteriaRelatedVulnerabilitiesList {
	var returns Inspector2FilterFilterCriteriaRelatedVulnerabilitiesList
	_jsii_.Get(
		j,
		"relatedVulnerabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) RelatedVulnerabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedVulnerabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResourceId() Inspector2FilterFilterCriteriaResourceIdList {
	var returns Inspector2FilterFilterCriteriaResourceIdList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResourceTags() Inspector2FilterFilterCriteriaResourceTagsList {
	var returns Inspector2FilterFilterCriteriaResourceTagsList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResourceType() Inspector2FilterFilterCriteriaResourceTypeList {
	var returns Inspector2FilterFilterCriteriaResourceTypeList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) Severity() Inspector2FilterFilterCriteriaSeverityList {
	var returns Inspector2FilterFilterCriteriaSeverityList
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) SeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) Title() Inspector2FilterFilterCriteriaTitleList {
	var returns Inspector2FilterFilterCriteriaTitleList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) TitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) UpdatedAt() Inspector2FilterFilterCriteriaUpdatedAtList {
	var returns Inspector2FilterFilterCriteriaUpdatedAtList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) UpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) VendorSeverity() Inspector2FilterFilterCriteriaVendorSeverityList {
	var returns Inspector2FilterFilterCriteriaVendorSeverityList
	_jsii_.Get(
		j,
		"vendorSeverity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) VendorSeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vendorSeverityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) VulnerabilityId() Inspector2FilterFilterCriteriaVulnerabilityIdList {
	var returns Inspector2FilterFilterCriteriaVulnerabilityIdList
	_jsii_.Get(
		j,
		"vulnerabilityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) VulnerabilityIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerabilityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) VulnerabilitySource() Inspector2FilterFilterCriteriaVulnerabilitySourceList {
	var returns Inspector2FilterFilterCriteriaVulnerabilitySourceList
	_jsii_.Get(
		j,
		"vulnerabilitySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) VulnerabilitySourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerabilitySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) VulnerablePackages() Inspector2FilterFilterCriteriaVulnerablePackagesList {
	var returns Inspector2FilterFilterCriteriaVulnerablePackagesList
	_jsii_.Get(
		j,
		"vulnerablePackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) VulnerablePackagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerablePackagesInput",
		&returns,
	)
	return returns
}


func NewInspector2FilterFilterCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Inspector2FilterFilterCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewInspector2FilterFilterCriteriaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Inspector2FilterFilterCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.inspector2Filter.Inspector2FilterFilterCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewInspector2FilterFilterCriteriaOutputReference_Override(i Inspector2FilterFilterCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.inspector2Filter.Inspector2FilterFilterCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutAwsAccountId(value interface{}) {
	if err := i.validatePutAwsAccountIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAwsAccountId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutCodeRepositoryProjectName(value interface{}) {
	if err := i.validatePutCodeRepositoryProjectNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCodeRepositoryProjectName",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutCodeRepositoryProviderType(value interface{}) {
	if err := i.validatePutCodeRepositoryProviderTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCodeRepositoryProviderType",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutCodeVulnerabilityDetectorName(value interface{}) {
	if err := i.validatePutCodeVulnerabilityDetectorNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCodeVulnerabilityDetectorName",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutCodeVulnerabilityDetectorTags(value interface{}) {
	if err := i.validatePutCodeVulnerabilityDetectorTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCodeVulnerabilityDetectorTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutCodeVulnerabilityFilePath(value interface{}) {
	if err := i.validatePutCodeVulnerabilityFilePathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCodeVulnerabilityFilePath",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutComponentId(value interface{}) {
	if err := i.validatePutComponentIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putComponentId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutComponentType(value interface{}) {
	if err := i.validatePutComponentTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putComponentType",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutEc2InstanceImageId(value interface{}) {
	if err := i.validatePutEc2InstanceImageIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEc2InstanceImageId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutEc2InstanceSubnetId(value interface{}) {
	if err := i.validatePutEc2InstanceSubnetIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEc2InstanceSubnetId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutEc2InstanceVpcId(value interface{}) {
	if err := i.validatePutEc2InstanceVpcIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEc2InstanceVpcId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutEcrImageArchitecture(value interface{}) {
	if err := i.validatePutEcrImageArchitectureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImageArchitecture",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutEcrImageHash(value interface{}) {
	if err := i.validatePutEcrImageHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImageHash",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutEcrImageInUseCount(value interface{}) {
	if err := i.validatePutEcrImageInUseCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImageInUseCount",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutEcrImageLastInUseAt(value interface{}) {
	if err := i.validatePutEcrImageLastInUseAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImageLastInUseAt",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutEcrImagePushedAt(value interface{}) {
	if err := i.validatePutEcrImagePushedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImagePushedAt",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutEcrImageRegistry(value interface{}) {
	if err := i.validatePutEcrImageRegistryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImageRegistry",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutEcrImageRepositoryName(value interface{}) {
	if err := i.validatePutEcrImageRepositoryNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImageRepositoryName",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutEcrImageTags(value interface{}) {
	if err := i.validatePutEcrImageTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImageTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutEpssScore(value interface{}) {
	if err := i.validatePutEpssScoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEpssScore",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutExploitAvailable(value interface{}) {
	if err := i.validatePutExploitAvailableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putExploitAvailable",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutFindingArn(value interface{}) {
	if err := i.validatePutFindingArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFindingArn",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutFindingStatus(value interface{}) {
	if err := i.validatePutFindingStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFindingStatus",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutFindingType(value interface{}) {
	if err := i.validatePutFindingTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFindingType",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutFirstObservedAt(value interface{}) {
	if err := i.validatePutFirstObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFirstObservedAt",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutFixAvailable(value interface{}) {
	if err := i.validatePutFixAvailableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFixAvailable",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutInspectorScore(value interface{}) {
	if err := i.validatePutInspectorScoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putInspectorScore",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutLambdaFunctionExecutionRoleArn(value interface{}) {
	if err := i.validatePutLambdaFunctionExecutionRoleArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambdaFunctionExecutionRoleArn",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutLambdaFunctionLastModifiedAt(value interface{}) {
	if err := i.validatePutLambdaFunctionLastModifiedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambdaFunctionLastModifiedAt",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutLambdaFunctionLayers(value interface{}) {
	if err := i.validatePutLambdaFunctionLayersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambdaFunctionLayers",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutLambdaFunctionName(value interface{}) {
	if err := i.validatePutLambdaFunctionNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambdaFunctionName",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutLambdaFunctionRuntime(value interface{}) {
	if err := i.validatePutLambdaFunctionRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambdaFunctionRuntime",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutLastObservedAt(value interface{}) {
	if err := i.validatePutLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLastObservedAt",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutNetworkProtocol(value interface{}) {
	if err := i.validatePutNetworkProtocolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putNetworkProtocol",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutPortRange(value interface{}) {
	if err := i.validatePutPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPortRange",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutRelatedVulnerabilities(value interface{}) {
	if err := i.validatePutRelatedVulnerabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRelatedVulnerabilities",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutResourceId(value interface{}) {
	if err := i.validatePutResourceIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putResourceId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutResourceTags(value interface{}) {
	if err := i.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutResourceType(value interface{}) {
	if err := i.validatePutResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putResourceType",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutSeverity(value interface{}) {
	if err := i.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSeverity",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutTitle(value interface{}) {
	if err := i.validatePutTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTitle",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutUpdatedAt(value interface{}) {
	if err := i.validatePutUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putUpdatedAt",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutVendorSeverity(value interface{}) {
	if err := i.validatePutVendorSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putVendorSeverity",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutVulnerabilityId(value interface{}) {
	if err := i.validatePutVulnerabilityIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putVulnerabilityId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutVulnerabilitySource(value interface{}) {
	if err := i.validatePutVulnerabilitySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putVulnerabilitySource",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) PutVulnerablePackages(value interface{}) {
	if err := i.validatePutVulnerablePackagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putVulnerablePackages",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		i,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetCodeRepositoryProjectName() {
	_jsii_.InvokeVoid(
		i,
		"resetCodeRepositoryProjectName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetCodeRepositoryProviderType() {
	_jsii_.InvokeVoid(
		i,
		"resetCodeRepositoryProviderType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetCodeVulnerabilityDetectorName() {
	_jsii_.InvokeVoid(
		i,
		"resetCodeVulnerabilityDetectorName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetCodeVulnerabilityDetectorTags() {
	_jsii_.InvokeVoid(
		i,
		"resetCodeVulnerabilityDetectorTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetCodeVulnerabilityFilePath() {
	_jsii_.InvokeVoid(
		i,
		"resetCodeVulnerabilityFilePath",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetComponentId() {
	_jsii_.InvokeVoid(
		i,
		"resetComponentId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetComponentType() {
	_jsii_.InvokeVoid(
		i,
		"resetComponentType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetEc2InstanceImageId() {
	_jsii_.InvokeVoid(
		i,
		"resetEc2InstanceImageId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetEc2InstanceSubnetId() {
	_jsii_.InvokeVoid(
		i,
		"resetEc2InstanceSubnetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetEc2InstanceVpcId() {
	_jsii_.InvokeVoid(
		i,
		"resetEc2InstanceVpcId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetEcrImageArchitecture() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImageArchitecture",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetEcrImageHash() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImageHash",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetEcrImageInUseCount() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImageInUseCount",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetEcrImageLastInUseAt() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImageLastInUseAt",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetEcrImagePushedAt() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImagePushedAt",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetEcrImageRegistry() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImageRegistry",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetEcrImageRepositoryName() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImageRepositoryName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetEcrImageTags() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImageTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetEpssScore() {
	_jsii_.InvokeVoid(
		i,
		"resetEpssScore",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetExploitAvailable() {
	_jsii_.InvokeVoid(
		i,
		"resetExploitAvailable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetFindingArn() {
	_jsii_.InvokeVoid(
		i,
		"resetFindingArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetFindingStatus() {
	_jsii_.InvokeVoid(
		i,
		"resetFindingStatus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetFindingType() {
	_jsii_.InvokeVoid(
		i,
		"resetFindingType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetFirstObservedAt() {
	_jsii_.InvokeVoid(
		i,
		"resetFirstObservedAt",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetFixAvailable() {
	_jsii_.InvokeVoid(
		i,
		"resetFixAvailable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetInspectorScore() {
	_jsii_.InvokeVoid(
		i,
		"resetInspectorScore",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetLambdaFunctionExecutionRoleArn() {
	_jsii_.InvokeVoid(
		i,
		"resetLambdaFunctionExecutionRoleArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetLambdaFunctionLastModifiedAt() {
	_jsii_.InvokeVoid(
		i,
		"resetLambdaFunctionLastModifiedAt",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetLambdaFunctionLayers() {
	_jsii_.InvokeVoid(
		i,
		"resetLambdaFunctionLayers",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetLambdaFunctionName() {
	_jsii_.InvokeVoid(
		i,
		"resetLambdaFunctionName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetLambdaFunctionRuntime() {
	_jsii_.InvokeVoid(
		i,
		"resetLambdaFunctionRuntime",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetLastObservedAt() {
	_jsii_.InvokeVoid(
		i,
		"resetLastObservedAt",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetNetworkProtocol() {
	_jsii_.InvokeVoid(
		i,
		"resetNetworkProtocol",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetPortRange() {
	_jsii_.InvokeVoid(
		i,
		"resetPortRange",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetRelatedVulnerabilities() {
	_jsii_.InvokeVoid(
		i,
		"resetRelatedVulnerabilities",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		i,
		"resetResourceId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		i,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		i,
		"resetResourceType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		i,
		"resetSeverity",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		i,
		"resetTitle",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		i,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetVendorSeverity() {
	_jsii_.InvokeVoid(
		i,
		"resetVendorSeverity",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetVulnerabilityId() {
	_jsii_.InvokeVoid(
		i,
		"resetVulnerabilityId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetVulnerabilitySource() {
	_jsii_.InvokeVoid(
		i,
		"resetVulnerabilitySource",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ResetVulnerablePackages() {
	_jsii_.InvokeVoid(
		i,
		"resetVulnerablePackages",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

