// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appstreamstack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/appstreamstack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/appstream_stack aws_appstream_stack}.
type AppstreamStack interface {
	cdktf.TerraformResource
	AccessEndpoints() AppstreamStackAccessEndpointsList
	AccessEndpointsInput() interface{}
	ApplicationSettings() AppstreamStackApplicationSettingsOutputReference
	ApplicationSettingsInput() *AppstreamStackApplicationSettings
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreatedTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EmbedHostDomains() *[]*string
	SetEmbedHostDomains(val *[]*string)
	EmbedHostDomainsInput() *[]*string
	FeedbackUrl() *string
	SetFeedbackUrl(val *string)
	FeedbackUrlInput() *string
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
	RedirectUrl() *string
	SetRedirectUrl(val *string)
	RedirectUrlInput() *string
	StorageConnectors() AppstreamStackStorageConnectorsList
	StorageConnectorsInput() interface{}
	StreamingExperienceSettings() AppstreamStackStreamingExperienceSettingsOutputReference
	StreamingExperienceSettingsInput() *AppstreamStackStreamingExperienceSettings
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
	UserSettings() AppstreamStackUserSettingsList
	UserSettingsInput() interface{}
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
	PutAccessEndpoints(value interface{})
	PutApplicationSettings(value *AppstreamStackApplicationSettings)
	PutStorageConnectors(value interface{})
	PutStreamingExperienceSettings(value *AppstreamStackStreamingExperienceSettings)
	PutUserSettings(value interface{})
	ResetAccessEndpoints()
	ResetApplicationSettings()
	ResetDescription()
	ResetDisplayName()
	ResetEmbedHostDomains()
	ResetFeedbackUrl()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRedirectUrl()
	ResetStorageConnectors()
	ResetStreamingExperienceSettings()
	ResetTags()
	ResetTagsAll()
	ResetUserSettings()
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

// The jsii proxy struct for AppstreamStack
type jsiiProxy_AppstreamStack struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppstreamStack) AccessEndpoints() AppstreamStackAccessEndpointsList {
	var returns AppstreamStackAccessEndpointsList
	_jsii_.Get(
		j,
		"accessEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) AccessEndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) ApplicationSettings() AppstreamStackApplicationSettingsOutputReference {
	var returns AppstreamStackApplicationSettingsOutputReference
	_jsii_.Get(
		j,
		"applicationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) ApplicationSettingsInput() *AppstreamStackApplicationSettings {
	var returns *AppstreamStackApplicationSettings
	_jsii_.Get(
		j,
		"applicationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) EmbedHostDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"embedHostDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) EmbedHostDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"embedHostDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) FeedbackUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedbackUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) FeedbackUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedbackUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) RedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) RedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) StorageConnectors() AppstreamStackStorageConnectorsList {
	var returns AppstreamStackStorageConnectorsList
	_jsii_.Get(
		j,
		"storageConnectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) StorageConnectorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageConnectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) StreamingExperienceSettings() AppstreamStackStreamingExperienceSettingsOutputReference {
	var returns AppstreamStackStreamingExperienceSettingsOutputReference
	_jsii_.Get(
		j,
		"streamingExperienceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) StreamingExperienceSettingsInput() *AppstreamStackStreamingExperienceSettings {
	var returns *AppstreamStackStreamingExperienceSettings
	_jsii_.Get(
		j,
		"streamingExperienceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) UserSettings() AppstreamStackUserSettingsList {
	var returns AppstreamStackUserSettingsList
	_jsii_.Get(
		j,
		"userSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) UserSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userSettingsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/appstream_stack aws_appstream_stack} Resource.
func NewAppstreamStack(scope constructs.Construct, id *string, config *AppstreamStackConfig) AppstreamStack {
	_init_.Initialize()

	if err := validateNewAppstreamStackParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppstreamStack{}

	_jsii_.Create(
		"@cdktf/provider-aws.appstreamStack.AppstreamStack",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/appstream_stack aws_appstream_stack} Resource.
func NewAppstreamStack_Override(a AppstreamStack, scope constructs.Construct, id *string, config *AppstreamStackConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appstreamStack.AppstreamStack",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppstreamStack)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetEmbedHostDomains(val *[]*string) {
	if err := j.validateSetEmbedHostDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embedHostDomains",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetFeedbackUrl(val *string) {
	if err := j.validateSetFeedbackUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"feedbackUrl",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetRedirectUrl(val *string) {
	if err := j.validateSetRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a AppstreamStack resource upon running "cdktf plan <stack-name>".
func AppstreamStack_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAppstreamStack_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appstreamStack.AppstreamStack",
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
func AppstreamStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppstreamStack_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appstreamStack.AppstreamStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppstreamStack_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppstreamStack_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appstreamStack.AppstreamStack",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppstreamStack_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppstreamStack_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appstreamStack.AppstreamStack",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppstreamStack_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.appstreamStack.AppstreamStack",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppstreamStack) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppstreamStack) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppstreamStack) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppstreamStack) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppstreamStack) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppstreamStack) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppstreamStack) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppstreamStack) PutAccessEndpoints(value interface{}) {
	if err := a.validatePutAccessEndpointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessEndpoints",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamStack) PutApplicationSettings(value *AppstreamStackApplicationSettings) {
	if err := a.validatePutApplicationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApplicationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamStack) PutStorageConnectors(value interface{}) {
	if err := a.validatePutStorageConnectorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStorageConnectors",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamStack) PutStreamingExperienceSettings(value *AppstreamStackStreamingExperienceSettings) {
	if err := a.validatePutStreamingExperienceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStreamingExperienceSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamStack) PutUserSettings(value interface{}) {
	if err := a.validatePutUserSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamStack) ResetAccessEndpoints() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessEndpoints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetApplicationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetEmbedHostDomains() {
	_jsii_.InvokeVoid(
		a,
		"resetEmbedHostDomains",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetFeedbackUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetFeedbackUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetStorageConnectors() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageConnectors",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetStreamingExperienceSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamingExperienceSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetUserSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetUserSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

