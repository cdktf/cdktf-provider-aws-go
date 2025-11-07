// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaceswebusersettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/workspaceswebusersettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/workspacesweb_user_settings aws_workspacesweb_user_settings}.
type WorkspaceswebUserSettings interface {
	cdktf.TerraformResource
	AdditionalEncryptionContext() *map[string]*string
	SetAdditionalEncryptionContext(val *map[string]*string)
	AdditionalEncryptionContextInput() *map[string]*string
	AssociatedPortalArns() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CookieSynchronizationConfiguration() WorkspaceswebUserSettingsCookieSynchronizationConfigurationList
	CookieSynchronizationConfigurationInput() interface{}
	CopyAllowed() *string
	SetCopyAllowed(val *string)
	CopyAllowedInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomerManagedKey() *string
	SetCustomerManagedKey(val *string)
	CustomerManagedKeyInput() *string
	DeepLinkAllowed() *string
	SetDeepLinkAllowed(val *string)
	DeepLinkAllowedInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisconnectTimeoutInMinutes() *float64
	SetDisconnectTimeoutInMinutes(val *float64)
	DisconnectTimeoutInMinutesInput() *float64
	DownloadAllowed() *string
	SetDownloadAllowed(val *string)
	DownloadAllowedInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IdleDisconnectTimeoutInMinutes() *float64
	SetIdleDisconnectTimeoutInMinutes(val *float64)
	IdleDisconnectTimeoutInMinutesInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PasteAllowed() *string
	SetPasteAllowed(val *string)
	PasteAllowedInput() *string
	PrintAllowed() *string
	SetPrintAllowed(val *string)
	PrintAllowedInput() *string
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktf.StringMap
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ToolbarConfiguration() WorkspaceswebUserSettingsToolbarConfigurationList
	ToolbarConfigurationInput() interface{}
	UploadAllowed() *string
	SetUploadAllowed(val *string)
	UploadAllowedInput() *string
	UserSettingsArn() *string
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
	PutCookieSynchronizationConfiguration(value interface{})
	PutToolbarConfiguration(value interface{})
	ResetAdditionalEncryptionContext()
	ResetCookieSynchronizationConfiguration()
	ResetCustomerManagedKey()
	ResetDeepLinkAllowed()
	ResetDisconnectTimeoutInMinutes()
	ResetIdleDisconnectTimeoutInMinutes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetTags()
	ResetToolbarConfiguration()
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

// The jsii proxy struct for WorkspaceswebUserSettings
type jsiiProxy_WorkspaceswebUserSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WorkspaceswebUserSettings) AdditionalEncryptionContext() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalEncryptionContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) AdditionalEncryptionContextInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalEncryptionContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) AssociatedPortalArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"associatedPortalArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CookieSynchronizationConfiguration() WorkspaceswebUserSettingsCookieSynchronizationConfigurationList {
	var returns WorkspaceswebUserSettingsCookieSynchronizationConfigurationList
	_jsii_.Get(
		j,
		"cookieSynchronizationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CookieSynchronizationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookieSynchronizationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CopyAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CopyAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CustomerManagedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerManagedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CustomerManagedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerManagedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DeepLinkAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deepLinkAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DeepLinkAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deepLinkAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DisconnectTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DisconnectTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DownloadAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"downloadAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DownloadAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"downloadAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) IdleDisconnectTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) IdleDisconnectTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) PasteAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pasteAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) PasteAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pasteAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) PrintAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) PrintAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) ToolbarConfiguration() WorkspaceswebUserSettingsToolbarConfigurationList {
	var returns WorkspaceswebUserSettingsToolbarConfigurationList
	_jsii_.Get(
		j,
		"toolbarConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) ToolbarConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"toolbarConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) UploadAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) UploadAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) UserSettingsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSettingsArn",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/workspacesweb_user_settings aws_workspacesweb_user_settings} Resource.
func NewWorkspaceswebUserSettings(scope constructs.Construct, id *string, config *WorkspaceswebUserSettingsConfig) WorkspaceswebUserSettings {
	_init_.Initialize()

	if err := validateNewWorkspaceswebUserSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceswebUserSettings{}

	_jsii_.Create(
		"@cdktf/provider-aws.workspaceswebUserSettings.WorkspaceswebUserSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/workspacesweb_user_settings aws_workspacesweb_user_settings} Resource.
func NewWorkspaceswebUserSettings_Override(w WorkspaceswebUserSettings, scope constructs.Construct, id *string, config *WorkspaceswebUserSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.workspaceswebUserSettings.WorkspaceswebUserSettings",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetAdditionalEncryptionContext(val *map[string]*string) {
	if err := j.validateSetAdditionalEncryptionContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalEncryptionContext",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetCopyAllowed(val *string) {
	if err := j.validateSetCopyAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyAllowed",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetCustomerManagedKey(val *string) {
	if err := j.validateSetCustomerManagedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerManagedKey",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetDeepLinkAllowed(val *string) {
	if err := j.validateSetDeepLinkAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deepLinkAllowed",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetDisconnectTimeoutInMinutes(val *float64) {
	if err := j.validateSetDisconnectTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disconnectTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetDownloadAllowed(val *string) {
	if err := j.validateSetDownloadAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"downloadAllowed",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetIdleDisconnectTimeoutInMinutes(val *float64) {
	if err := j.validateSetIdleDisconnectTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleDisconnectTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetPasteAllowed(val *string) {
	if err := j.validateSetPasteAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pasteAllowed",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetPrintAllowed(val *string) {
	if err := j.validateSetPrintAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"printAllowed",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetUploadAllowed(val *string) {
	if err := j.validateSetUploadAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadAllowed",
		val,
	)
}

// Generates CDKTF code for importing a WorkspaceswebUserSettings resource upon running "cdktf plan <stack-name>".
func WorkspaceswebUserSettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWorkspaceswebUserSettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.workspaceswebUserSettings.WorkspaceswebUserSettings",
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
func WorkspaceswebUserSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkspaceswebUserSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.workspaceswebUserSettings.WorkspaceswebUserSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkspaceswebUserSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkspaceswebUserSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.workspaceswebUserSettings.WorkspaceswebUserSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkspaceswebUserSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkspaceswebUserSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.workspaceswebUserSettings.WorkspaceswebUserSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorkspaceswebUserSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.workspaceswebUserSettings.WorkspaceswebUserSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) PutCookieSynchronizationConfiguration(value interface{}) {
	if err := w.validatePutCookieSynchronizationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookieSynchronizationConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) PutToolbarConfiguration(value interface{}) {
	if err := w.validatePutToolbarConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putToolbarConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetAdditionalEncryptionContext() {
	_jsii_.InvokeVoid(
		w,
		"resetAdditionalEncryptionContext",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetCookieSynchronizationConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetCookieSynchronizationConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetCustomerManagedKey() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomerManagedKey",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetDeepLinkAllowed() {
	_jsii_.InvokeVoid(
		w,
		"resetDeepLinkAllowed",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetDisconnectTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		w,
		"resetDisconnectTimeoutInMinutes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetIdleDisconnectTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		w,
		"resetIdleDisconnectTimeoutInMinutes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetRegion() {
	_jsii_.InvokeVoid(
		w,
		"resetRegion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetToolbarConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetToolbarConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

