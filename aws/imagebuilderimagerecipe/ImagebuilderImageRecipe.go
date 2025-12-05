// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagerecipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/imagebuilderimagerecipe/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/imagebuilder_image_recipe aws_imagebuilder_image_recipe}.
type ImagebuilderImageRecipe interface {
	cdktf.TerraformResource
	AmiTags() *map[string]*string
	SetAmiTags(val *map[string]*string)
	AmiTagsInput() *map[string]*string
	Arn() *string
	BlockDeviceMapping() ImagebuilderImageRecipeBlockDeviceMappingList
	BlockDeviceMappingInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Component() ImagebuilderImageRecipeComponentList
	ComponentInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DateCreated() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Owner() *string
	ParentImage() *string
	SetParentImage(val *string)
	ParentImageInput() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SystemsManagerAgent() ImagebuilderImageRecipeSystemsManagerAgentOutputReference
	SystemsManagerAgentInput() *ImagebuilderImageRecipeSystemsManagerAgent
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
	UserDataBase64() *string
	SetUserDataBase64(val *string)
	UserDataBase64Input() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	WorkingDirectory() *string
	SetWorkingDirectory(val *string)
	WorkingDirectoryInput() *string
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
	PutBlockDeviceMapping(value interface{})
	PutComponent(value interface{})
	PutSystemsManagerAgent(value *ImagebuilderImageRecipeSystemsManagerAgent)
	ResetAmiTags()
	ResetBlockDeviceMapping()
	ResetDescription()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSystemsManagerAgent()
	ResetTags()
	ResetTagsAll()
	ResetUserDataBase64()
	ResetWorkingDirectory()
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

// The jsii proxy struct for ImagebuilderImageRecipe
type jsiiProxy_ImagebuilderImageRecipe struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ImagebuilderImageRecipe) AmiTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"amiTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) AmiTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"amiTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) BlockDeviceMapping() ImagebuilderImageRecipeBlockDeviceMappingList {
	var returns ImagebuilderImageRecipeBlockDeviceMappingList
	_jsii_.Get(
		j,
		"blockDeviceMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) BlockDeviceMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Component() ImagebuilderImageRecipeComponentList {
	var returns ImagebuilderImageRecipeComponentList
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) ComponentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) DateCreated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateCreated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) ParentImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) ParentImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) SystemsManagerAgent() ImagebuilderImageRecipeSystemsManagerAgentOutputReference {
	var returns ImagebuilderImageRecipeSystemsManagerAgentOutputReference
	_jsii_.Get(
		j,
		"systemsManagerAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) SystemsManagerAgentInput() *ImagebuilderImageRecipeSystemsManagerAgent {
	var returns *ImagebuilderImageRecipeSystemsManagerAgent
	_jsii_.Get(
		j,
		"systemsManagerAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) UserDataBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) UserDataBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageRecipe) WorkingDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectoryInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/imagebuilder_image_recipe aws_imagebuilder_image_recipe} Resource.
func NewImagebuilderImageRecipe(scope constructs.Construct, id *string, config *ImagebuilderImageRecipeConfig) ImagebuilderImageRecipe {
	_init_.Initialize()

	if err := validateNewImagebuilderImageRecipeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImagebuilderImageRecipe{}

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilderImageRecipe.ImagebuilderImageRecipe",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/imagebuilder_image_recipe aws_imagebuilder_image_recipe} Resource.
func NewImagebuilderImageRecipe_Override(i ImagebuilderImageRecipe, scope constructs.Construct, id *string, config *ImagebuilderImageRecipeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilderImageRecipe.ImagebuilderImageRecipe",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetAmiTags(val *map[string]*string) {
	if err := j.validateSetAmiTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amiTags",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetParentImage(val *string) {
	if err := j.validateSetParentImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentImage",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetUserDataBase64(val *string) {
	if err := j.validateSetUserDataBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDataBase64",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageRecipe)SetWorkingDirectory(val *string) {
	if err := j.validateSetWorkingDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDirectory",
		val,
	)
}

// Generates CDKTF code for importing a ImagebuilderImageRecipe resource upon running "cdktf plan <stack-name>".
func ImagebuilderImageRecipe_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateImagebuilderImageRecipe_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.imagebuilderImageRecipe.ImagebuilderImageRecipe",
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
func ImagebuilderImageRecipe_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImagebuilderImageRecipe_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.imagebuilderImageRecipe.ImagebuilderImageRecipe",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ImagebuilderImageRecipe_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImagebuilderImageRecipe_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.imagebuilderImageRecipe.ImagebuilderImageRecipe",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ImagebuilderImageRecipe_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImagebuilderImageRecipe_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.imagebuilderImageRecipe.ImagebuilderImageRecipe",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ImagebuilderImageRecipe_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.imagebuilderImageRecipe.ImagebuilderImageRecipe",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_ImagebuilderImageRecipe) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_ImagebuilderImageRecipe) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_ImagebuilderImageRecipe) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_ImagebuilderImageRecipe) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_ImagebuilderImageRecipe) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_ImagebuilderImageRecipe) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_ImagebuilderImageRecipe) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_ImagebuilderImageRecipe) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_ImagebuilderImageRecipe) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_ImagebuilderImageRecipe) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_ImagebuilderImageRecipe) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) PutBlockDeviceMapping(value interface{}) {
	if err := i.validatePutBlockDeviceMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putBlockDeviceMapping",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) PutComponent(value interface{}) {
	if err := i.validatePutComponentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putComponent",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) PutSystemsManagerAgent(value *ImagebuilderImageRecipeSystemsManagerAgent) {
	if err := i.validatePutSystemsManagerAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSystemsManagerAgent",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ResetAmiTags() {
	_jsii_.InvokeVoid(
		i,
		"resetAmiTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ResetBlockDeviceMapping() {
	_jsii_.InvokeVoid(
		i,
		"resetBlockDeviceMapping",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ResetRegion() {
	_jsii_.InvokeVoid(
		i,
		"resetRegion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ResetSystemsManagerAgent() {
	_jsii_.InvokeVoid(
		i,
		"resetSystemsManagerAgent",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ResetTagsAll() {
	_jsii_.InvokeVoid(
		i,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ResetUserDataBase64() {
	_jsii_.InvokeVoid(
		i,
		"resetUserDataBase64",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ResetWorkingDirectory() {
	_jsii_.InvokeVoid(
		i,
		"resetWorkingDirectory",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImageRecipe) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageRecipe) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageRecipe) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

