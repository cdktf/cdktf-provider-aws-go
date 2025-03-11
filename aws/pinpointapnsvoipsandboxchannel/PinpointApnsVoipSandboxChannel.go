// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pinpointapnsvoipsandboxchannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/pinpointapnsvoipsandboxchannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/pinpoint_apns_voip_sandbox_channel aws_pinpoint_apns_voip_sandbox_channel}.
type PinpointApnsVoipSandboxChannel interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	BundleId() *string
	SetBundleId(val *string)
	BundleIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
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
	DefaultAuthenticationMethod() *string
	SetDefaultAuthenticationMethod(val *string)
	DefaultAuthenticationMethodInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
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
	TeamId() *string
	SetTeamId(val *string)
	TeamIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenKey() *string
	SetTokenKey(val *string)
	TokenKeyId() *string
	SetTokenKeyId(val *string)
	TokenKeyIdInput() *string
	TokenKeyInput() *string
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
	ResetBundleId()
	ResetCertificate()
	ResetDefaultAuthenticationMethod()
	ResetEnabled()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateKey()
	ResetTeamId()
	ResetTokenKey()
	ResetTokenKeyId()
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

// The jsii proxy struct for PinpointApnsVoipSandboxChannel
type jsiiProxy_PinpointApnsVoipSandboxChannel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) BundleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) DefaultAuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) DefaultAuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TokenKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TokenKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TokenKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/pinpoint_apns_voip_sandbox_channel aws_pinpoint_apns_voip_sandbox_channel} Resource.
func NewPinpointApnsVoipSandboxChannel(scope constructs.Construct, id *string, config *PinpointApnsVoipSandboxChannelConfig) PinpointApnsVoipSandboxChannel {
	_init_.Initialize()

	if err := validateNewPinpointApnsVoipSandboxChannelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PinpointApnsVoipSandboxChannel{}

	_jsii_.Create(
		"@cdktf/provider-aws.pinpointApnsVoipSandboxChannel.PinpointApnsVoipSandboxChannel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/pinpoint_apns_voip_sandbox_channel aws_pinpoint_apns_voip_sandbox_channel} Resource.
func NewPinpointApnsVoipSandboxChannel_Override(p PinpointApnsVoipSandboxChannel, scope constructs.Construct, id *string, config *PinpointApnsVoipSandboxChannelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pinpointApnsVoipSandboxChannel.PinpointApnsVoipSandboxChannel",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetApplicationId(val *string) {
	if err := j.validateSetApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetBundleId(val *string) {
	if err := j.validateSetBundleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetDefaultAuthenticationMethod(val *string) {
	if err := j.validateSetDefaultAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAuthenticationMethod",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetPrivateKey(val *string) {
	if err := j.validateSetPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetTeamId(val *string) {
	if err := j.validateSetTeamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetTokenKey(val *string) {
	if err := j.validateSetTokenKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenKey",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel)SetTokenKeyId(val *string) {
	if err := j.validateSetTokenKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenKeyId",
		val,
	)
}

// Generates CDKTF code for importing a PinpointApnsVoipSandboxChannel resource upon running "cdktf plan <stack-name>".
func PinpointApnsVoipSandboxChannel_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePinpointApnsVoipSandboxChannel_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.pinpointApnsVoipSandboxChannel.PinpointApnsVoipSandboxChannel",
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
func PinpointApnsVoipSandboxChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePinpointApnsVoipSandboxChannel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.pinpointApnsVoipSandboxChannel.PinpointApnsVoipSandboxChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PinpointApnsVoipSandboxChannel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePinpointApnsVoipSandboxChannel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.pinpointApnsVoipSandboxChannel.PinpointApnsVoipSandboxChannel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PinpointApnsVoipSandboxChannel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePinpointApnsVoipSandboxChannel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.pinpointApnsVoipSandboxChannel.PinpointApnsVoipSandboxChannel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PinpointApnsVoipSandboxChannel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.pinpointApnsVoipSandboxChannel.PinpointApnsVoipSandboxChannel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetBundleId() {
	_jsii_.InvokeVoid(
		p,
		"resetBundleId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetCertificate() {
	_jsii_.InvokeVoid(
		p,
		"resetCertificate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetDefaultAuthenticationMethod() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultAuthenticationMethod",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetTeamId() {
	_jsii_.InvokeVoid(
		p,
		"resetTeamId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetTokenKey() {
	_jsii_.InvokeVoid(
		p,
		"resetTokenKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetTokenKeyId() {
	_jsii_.InvokeVoid(
		p,
		"resetTokenKeyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

