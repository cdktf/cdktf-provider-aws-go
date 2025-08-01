// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerimageversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/sagemakerimageversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/sagemaker_image_version aws_sagemaker_image_version}.
type SagemakerImageVersion interface {
	cdktf.TerraformResource
	Aliases() *[]*string
	SetAliases(val *[]*string)
	AliasesInput() *[]*string
	Arn() *string
	BaseImage() *string
	SetBaseImage(val *string)
	BaseImageInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerImage() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Horovod() interface{}
	SetHorovod(val interface{})
	HorovodInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageArn() *string
	ImageName() *string
	SetImageName(val *string)
	ImageNameInput() *string
	JobType() *string
	SetJobType(val *string)
	JobTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MlFramework() *string
	SetMlFramework(val *string)
	MlFrameworkInput() *string
	// The tree node.
	Node() constructs.Node
	Processor() *string
	SetProcessor(val *string)
	ProcessorInput() *string
	ProgrammingLang() *string
	SetProgrammingLang(val *string)
	ProgrammingLangInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReleaseNotes() *string
	SetReleaseNotes(val *string)
	ReleaseNotesInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VendorGuidance() *string
	SetVendorGuidance(val *string)
	VendorGuidanceInput() *string
	Version() *float64
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAliases()
	ResetHorovod()
	ResetId()
	ResetJobType()
	ResetMlFramework()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProcessor()
	ResetProgrammingLang()
	ResetRegion()
	ResetReleaseNotes()
	ResetVendorGuidance()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SagemakerImageVersion
type jsiiProxy_SagemakerImageVersion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SagemakerImageVersion) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) AliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) BaseImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) BaseImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) ContainerImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) Horovod() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"horovod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) HorovodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"horovodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) ImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) JobType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) JobTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) MlFramework() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlFramework",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) MlFrameworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlFrameworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) Processor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) ProcessorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) ProgrammingLang() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programmingLang",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) ProgrammingLangInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programmingLangInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) ReleaseNotes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseNotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) ReleaseNotesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseNotesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) VendorGuidance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vendorGuidance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) VendorGuidanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vendorGuidanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerImageVersion) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/sagemaker_image_version aws_sagemaker_image_version} Resource.
func NewSagemakerImageVersion(scope constructs.Construct, id *string, config *SagemakerImageVersionConfig) SagemakerImageVersion {
	_init_.Initialize()

	if err := validateNewSagemakerImageVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerImageVersion{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerImageVersion.SagemakerImageVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/sagemaker_image_version aws_sagemaker_image_version} Resource.
func NewSagemakerImageVersion_Override(s SagemakerImageVersion, scope constructs.Construct, id *string, config *SagemakerImageVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerImageVersion.SagemakerImageVersion",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetAliases(val *[]*string) {
	if err := j.validateSetAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliases",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetBaseImage(val *string) {
	if err := j.validateSetBaseImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseImage",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetHorovod(val interface{}) {
	if err := j.validateSetHorovodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"horovod",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetImageName(val *string) {
	if err := j.validateSetImageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetJobType(val *string) {
	if err := j.validateSetJobTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobType",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetMlFramework(val *string) {
	if err := j.validateSetMlFrameworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mlFramework",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetProcessor(val *string) {
	if err := j.validateSetProcessorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetProgrammingLang(val *string) {
	if err := j.validateSetProgrammingLangParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programmingLang",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetReleaseNotes(val *string) {
	if err := j.validateSetReleaseNotesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseNotes",
		val,
	)
}

func (j *jsiiProxy_SagemakerImageVersion)SetVendorGuidance(val *string) {
	if err := j.validateSetVendorGuidanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vendorGuidance",
		val,
	)
}

// Generates CDKTF code for importing a SagemakerImageVersion resource upon running "cdktf plan <stack-name>".
func SagemakerImageVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSagemakerImageVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sagemakerImageVersion.SagemakerImageVersion",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SagemakerImageVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerImageVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sagemakerImageVersion.SagemakerImageVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerImageVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerImageVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sagemakerImageVersion.SagemakerImageVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerImageVersion_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerImageVersion_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sagemakerImageVersion.SagemakerImageVersion",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SagemakerImageVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.sagemakerImageVersion.SagemakerImageVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SagemakerImageVersion) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SagemakerImageVersion) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SagemakerImageVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerImageVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerImageVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerImageVersion) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerImageVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerImageVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerImageVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerImageVersion) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerImageVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerImageVersion) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerImageVersion) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SagemakerImageVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerImageVersion) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SagemakerImageVersion) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SagemakerImageVersion) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SagemakerImageVersion) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SagemakerImageVersion) ResetAliases() {
	_jsii_.InvokeVoid(
		s,
		"resetAliases",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerImageVersion) ResetHorovod() {
	_jsii_.InvokeVoid(
		s,
		"resetHorovod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerImageVersion) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerImageVersion) ResetJobType() {
	_jsii_.InvokeVoid(
		s,
		"resetJobType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerImageVersion) ResetMlFramework() {
	_jsii_.InvokeVoid(
		s,
		"resetMlFramework",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerImageVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerImageVersion) ResetProcessor() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessor",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerImageVersion) ResetProgrammingLang() {
	_jsii_.InvokeVoid(
		s,
		"resetProgrammingLang",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerImageVersion) ResetRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerImageVersion) ResetReleaseNotes() {
	_jsii_.InvokeVoid(
		s,
		"resetReleaseNotes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerImageVersion) ResetVendorGuidance() {
	_jsii_.InvokeVoid(
		s,
		"resetVendorGuidance",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerImageVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerImageVersion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerImageVersion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerImageVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerImageVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerImageVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

