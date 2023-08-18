package pinpointapnssandboxchannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/pinpointapnssandboxchannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/pinpoint_apns_sandbox_channel aws_pinpoint_apns_sandbox_channel}.
type PinpointApnsSandboxChannel interface {
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
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
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PinpointApnsSandboxChannel
type jsiiProxy_PinpointApnsSandboxChannel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) BundleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) DefaultAuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) DefaultAuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TokenKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TokenKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TokenKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/pinpoint_apns_sandbox_channel aws_pinpoint_apns_sandbox_channel} Resource.
func NewPinpointApnsSandboxChannel(scope constructs.Construct, id *string, config *PinpointApnsSandboxChannelConfig) PinpointApnsSandboxChannel {
	_init_.Initialize()

	if err := validateNewPinpointApnsSandboxChannelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PinpointApnsSandboxChannel{}

	_jsii_.Create(
		"@cdktf/provider-aws.pinpointApnsSandboxChannel.PinpointApnsSandboxChannel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/pinpoint_apns_sandbox_channel aws_pinpoint_apns_sandbox_channel} Resource.
func NewPinpointApnsSandboxChannel_Override(p PinpointApnsSandboxChannel, scope constructs.Construct, id *string, config *PinpointApnsSandboxChannelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pinpointApnsSandboxChannel.PinpointApnsSandboxChannel",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetApplicationId(val *string) {
	if err := j.validateSetApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetBundleId(val *string) {
	if err := j.validateSetBundleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetDefaultAuthenticationMethod(val *string) {
	if err := j.validateSetDefaultAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAuthenticationMethod",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetPrivateKey(val *string) {
	if err := j.validateSetPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetTeamId(val *string) {
	if err := j.validateSetTeamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetTokenKey(val *string) {
	if err := j.validateSetTokenKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenKey",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel)SetTokenKeyId(val *string) {
	if err := j.validateSetTokenKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenKeyId",
		val,
	)
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
func PinpointApnsSandboxChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePinpointApnsSandboxChannel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.pinpointApnsSandboxChannel.PinpointApnsSandboxChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PinpointApnsSandboxChannel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePinpointApnsSandboxChannel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.pinpointApnsSandboxChannel.PinpointApnsSandboxChannel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PinpointApnsSandboxChannel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePinpointApnsSandboxChannel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.pinpointApnsSandboxChannel.PinpointApnsSandboxChannel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PinpointApnsSandboxChannel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.pinpointApnsSandboxChannel.PinpointApnsSandboxChannel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PinpointApnsSandboxChannel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PinpointApnsSandboxChannel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PinpointApnsSandboxChannel) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PinpointApnsSandboxChannel) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PinpointApnsSandboxChannel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PinpointApnsSandboxChannel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PinpointApnsSandboxChannel) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PinpointApnsSandboxChannel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PinpointApnsSandboxChannel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PinpointApnsSandboxChannel) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetBundleId() {
	_jsii_.InvokeVoid(
		p,
		"resetBundleId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetCertificate() {
	_jsii_.InvokeVoid(
		p,
		"resetCertificate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetDefaultAuthenticationMethod() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultAuthenticationMethod",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetTeamId() {
	_jsii_.InvokeVoid(
		p,
		"resetTeamId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetTokenKey() {
	_jsii_.InvokeVoid(
		p,
		"resetTokenKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetTokenKeyId() {
	_jsii_.InvokeVoid(
		p,
		"resetTokenKeyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

