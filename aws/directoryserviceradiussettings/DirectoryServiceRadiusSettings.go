// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryserviceradiussettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/directoryserviceradiussettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/directory_service_radius_settings aws_directory_service_radius_settings}.
type DirectoryServiceRadiusSettings interface {
	cdktf.TerraformResource
	AuthenticationProtocol() *string
	SetAuthenticationProtocol(val *string)
	AuthenticationProtocolInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DirectoryId() *string
	SetDirectoryId(val *string)
	DirectoryIdInput() *string
	DisplayLabel() *string
	SetDisplayLabel(val *string)
	DisplayLabelInput() *string
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
	RadiusPort() *float64
	SetRadiusPort(val *float64)
	RadiusPortInput() *float64
	RadiusRetries() *float64
	SetRadiusRetries(val *float64)
	RadiusRetriesInput() *float64
	RadiusServers() *[]*string
	SetRadiusServers(val *[]*string)
	RadiusServersInput() *[]*string
	RadiusTimeout() *float64
	SetRadiusTimeout(val *float64)
	RadiusTimeoutInput() *float64
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SharedSecret() *string
	SetSharedSecret(val *string)
	SharedSecretInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DirectoryServiceRadiusSettingsTimeoutsOutputReference
	TimeoutsInput() interface{}
	UseSameUsername() interface{}
	SetUseSameUsername(val interface{})
	UseSameUsernameInput() interface{}
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
	PutTimeouts(value *DirectoryServiceRadiusSettingsTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetTimeouts()
	ResetUseSameUsername()
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

// The jsii proxy struct for DirectoryServiceRadiusSettings
type jsiiProxy_DirectoryServiceRadiusSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) AuthenticationProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) AuthenticationProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) DisplayLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) DisplayLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) RadiusPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"radiusPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) RadiusPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"radiusPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) RadiusRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"radiusRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) RadiusRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"radiusRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) RadiusServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"radiusServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) RadiusServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"radiusServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) RadiusTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"radiusTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) RadiusTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"radiusTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) SharedSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) SharedSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) Timeouts() DirectoryServiceRadiusSettingsTimeoutsOutputReference {
	var returns DirectoryServiceRadiusSettingsTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) UseSameUsername() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSameUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings) UseSameUsernameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSameUsernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/directory_service_radius_settings aws_directory_service_radius_settings} Resource.
func NewDirectoryServiceRadiusSettings(scope constructs.Construct, id *string, config *DirectoryServiceRadiusSettingsConfig) DirectoryServiceRadiusSettings {
	_init_.Initialize()

	if err := validateNewDirectoryServiceRadiusSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DirectoryServiceRadiusSettings{}

	_jsii_.Create(
		"@cdktf/provider-aws.directoryServiceRadiusSettings.DirectoryServiceRadiusSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/directory_service_radius_settings aws_directory_service_radius_settings} Resource.
func NewDirectoryServiceRadiusSettings_Override(d DirectoryServiceRadiusSettings, scope constructs.Construct, id *string, config *DirectoryServiceRadiusSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.directoryServiceRadiusSettings.DirectoryServiceRadiusSettings",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetAuthenticationProtocol(val *string) {
	if err := j.validateSetAuthenticationProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationProtocol",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetDirectoryId(val *string) {
	if err := j.validateSetDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetDisplayLabel(val *string) {
	if err := j.validateSetDisplayLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayLabel",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetRadiusPort(val *float64) {
	if err := j.validateSetRadiusPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"radiusPort",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetRadiusRetries(val *float64) {
	if err := j.validateSetRadiusRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"radiusRetries",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetRadiusServers(val *[]*string) {
	if err := j.validateSetRadiusServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"radiusServers",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetRadiusTimeout(val *float64) {
	if err := j.validateSetRadiusTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"radiusTimeout",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetSharedSecret(val *string) {
	if err := j.validateSetSharedSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedSecret",
		val,
	)
}

func (j *jsiiProxy_DirectoryServiceRadiusSettings)SetUseSameUsername(val interface{}) {
	if err := j.validateSetUseSameUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSameUsername",
		val,
	)
}

// Generates CDKTF code for importing a DirectoryServiceRadiusSettings resource upon running "cdktf plan <stack-name>".
func DirectoryServiceRadiusSettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDirectoryServiceRadiusSettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.directoryServiceRadiusSettings.DirectoryServiceRadiusSettings",
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
func DirectoryServiceRadiusSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectoryServiceRadiusSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.directoryServiceRadiusSettings.DirectoryServiceRadiusSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DirectoryServiceRadiusSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectoryServiceRadiusSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.directoryServiceRadiusSettings.DirectoryServiceRadiusSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DirectoryServiceRadiusSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectoryServiceRadiusSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.directoryServiceRadiusSettings.DirectoryServiceRadiusSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DirectoryServiceRadiusSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.directoryServiceRadiusSettings.DirectoryServiceRadiusSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) PutTimeouts(value *DirectoryServiceRadiusSettingsTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) ResetUseSameUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUseSameUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryServiceRadiusSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

