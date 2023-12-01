// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/imagebuilderimage/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/imagebuilder_image aws_imagebuilder_image}.
type ImagebuilderImage interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerRecipeArn() *string
	SetContainerRecipeArn(val *string)
	ContainerRecipeArnInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DateCreated() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DistributionConfigurationArn() *string
	SetDistributionConfigurationArn(val *string)
	DistributionConfigurationArnInput() *string
	EnhancedImageMetadataEnabled() interface{}
	SetEnhancedImageMetadataEnabled(val interface{})
	EnhancedImageMetadataEnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageRecipeArn() *string
	SetImageRecipeArn(val *string)
	ImageRecipeArnInput() *string
	ImageScanningConfiguration() ImagebuilderImageImageScanningConfigurationOutputReference
	ImageScanningConfigurationInput() *ImagebuilderImageImageScanningConfiguration
	ImageTestsConfiguration() ImagebuilderImageImageTestsConfigurationOutputReference
	ImageTestsConfigurationInput() *ImagebuilderImageImageTestsConfiguration
	InfrastructureConfigurationArn() *string
	SetInfrastructureConfigurationArn(val *string)
	InfrastructureConfigurationArnInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	OsVersion() *string
	OutputResources() ImagebuilderImageOutputResourcesList
	Platform() *string
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ImagebuilderImageTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutImageScanningConfiguration(value *ImagebuilderImageImageScanningConfiguration)
	PutImageTestsConfiguration(value *ImagebuilderImageImageTestsConfiguration)
	PutTimeouts(value *ImagebuilderImageTimeouts)
	ResetContainerRecipeArn()
	ResetDistributionConfigurationArn()
	ResetEnhancedImageMetadataEnabled()
	ResetId()
	ResetImageRecipeArn()
	ResetImageScanningConfiguration()
	ResetImageTestsConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ImagebuilderImage
type jsiiProxy_ImagebuilderImage struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ImagebuilderImage) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) ContainerRecipeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRecipeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) ContainerRecipeArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRecipeArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) DateCreated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateCreated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) DistributionConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) DistributionConfigurationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionConfigurationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) EnhancedImageMetadataEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedImageMetadataEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) EnhancedImageMetadataEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedImageMetadataEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) ImageRecipeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRecipeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) ImageRecipeArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRecipeArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) ImageScanningConfiguration() ImagebuilderImageImageScanningConfigurationOutputReference {
	var returns ImagebuilderImageImageScanningConfigurationOutputReference
	_jsii_.Get(
		j,
		"imageScanningConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) ImageScanningConfigurationInput() *ImagebuilderImageImageScanningConfiguration {
	var returns *ImagebuilderImageImageScanningConfiguration
	_jsii_.Get(
		j,
		"imageScanningConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) ImageTestsConfiguration() ImagebuilderImageImageTestsConfigurationOutputReference {
	var returns ImagebuilderImageImageTestsConfigurationOutputReference
	_jsii_.Get(
		j,
		"imageTestsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) ImageTestsConfigurationInput() *ImagebuilderImageImageTestsConfiguration {
	var returns *ImagebuilderImageImageTestsConfiguration
	_jsii_.Get(
		j,
		"imageTestsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) InfrastructureConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) InfrastructureConfigurationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureConfigurationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) OsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) OutputResources() ImagebuilderImageOutputResourcesList {
	var returns ImagebuilderImageOutputResourcesList
	_jsii_.Get(
		j,
		"outputResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) Timeouts() ImagebuilderImageTimeoutsOutputReference {
	var returns ImagebuilderImageTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImage) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/imagebuilder_image aws_imagebuilder_image} Resource.
func NewImagebuilderImage(scope constructs.Construct, id *string, config *ImagebuilderImageConfig) ImagebuilderImage {
	_init_.Initialize()

	if err := validateNewImagebuilderImageParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImagebuilderImage{}

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilderImage.ImagebuilderImage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/imagebuilder_image aws_imagebuilder_image} Resource.
func NewImagebuilderImage_Override(i ImagebuilderImage, scope constructs.Construct, id *string, config *ImagebuilderImageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilderImage.ImagebuilderImage",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetContainerRecipeArn(val *string) {
	if err := j.validateSetContainerRecipeArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRecipeArn",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetDistributionConfigurationArn(val *string) {
	if err := j.validateSetDistributionConfigurationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distributionConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetEnhancedImageMetadataEnabled(val interface{}) {
	if err := j.validateSetEnhancedImageMetadataEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedImageMetadataEnabled",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetImageRecipeArn(val *string) {
	if err := j.validateSetImageRecipeArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageRecipeArn",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetInfrastructureConfigurationArn(val *string) {
	if err := j.validateSetInfrastructureConfigurationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImage)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a ImagebuilderImage resource upon running "cdktf plan <stack-name>".
func ImagebuilderImage_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateImagebuilderImage_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.imagebuilderImage.ImagebuilderImage",
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
func ImagebuilderImage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImagebuilderImage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.imagebuilderImage.ImagebuilderImage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ImagebuilderImage_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImagebuilderImage_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.imagebuilderImage.ImagebuilderImage",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ImagebuilderImage_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImagebuilderImage_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.imagebuilderImage.ImagebuilderImage",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ImagebuilderImage_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.imagebuilderImage.ImagebuilderImage",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_ImagebuilderImage) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_ImagebuilderImage) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_ImagebuilderImage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_ImagebuilderImage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_ImagebuilderImage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_ImagebuilderImage) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_ImagebuilderImage) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_ImagebuilderImage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_ImagebuilderImage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_ImagebuilderImage) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_ImagebuilderImage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_ImagebuilderImage) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_ImagebuilderImage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImage) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_ImagebuilderImage) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_ImagebuilderImage) PutImageScanningConfiguration(value *ImagebuilderImageImageScanningConfiguration) {
	if err := i.validatePutImageScanningConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putImageScanningConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderImage) PutImageTestsConfiguration(value *ImagebuilderImageImageTestsConfiguration) {
	if err := i.validatePutImageTestsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putImageTestsConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderImage) PutTimeouts(value *ImagebuilderImageTimeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderImage) ResetContainerRecipeArn() {
	_jsii_.InvokeVoid(
		i,
		"resetContainerRecipeArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImage) ResetDistributionConfigurationArn() {
	_jsii_.InvokeVoid(
		i,
		"resetDistributionConfigurationArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImage) ResetEnhancedImageMetadataEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetEnhancedImageMetadataEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImage) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImage) ResetImageRecipeArn() {
	_jsii_.InvokeVoid(
		i,
		"resetImageRecipeArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImage) ResetImageScanningConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetImageScanningConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImage) ResetImageTestsConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetImageTestsConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImage) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImage) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImage) ResetTagsAll() {
	_jsii_.InvokeVoid(
		i,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImage) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImage) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImage) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImage) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

