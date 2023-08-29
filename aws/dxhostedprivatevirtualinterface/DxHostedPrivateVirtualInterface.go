// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dxhostedprivatevirtualinterface

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/dxhostedprivatevirtualinterface/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/dx_hosted_private_virtual_interface aws_dx_hosted_private_virtual_interface}.
type DxHostedPrivateVirtualInterface interface {
	cdktf.TerraformResource
	AddressFamily() *string
	SetAddressFamily(val *string)
	AddressFamilyInput() *string
	AmazonAddress() *string
	SetAmazonAddress(val *string)
	AmazonAddressInput() *string
	AmazonSideAsn() *string
	Arn() *string
	AwsDevice() *string
	BgpAsn() *float64
	SetBgpAsn(val *float64)
	BgpAsnInput() *float64
	BgpAuthKey() *string
	SetBgpAuthKey(val *string)
	BgpAuthKeyInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionId() *string
	SetConnectionId(val *string)
	ConnectionIdInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomerAddress() *string
	SetCustomerAddress(val *string)
	CustomerAddressInput() *string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	JumboFrameCapable() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Mtu() *float64
	SetMtu(val *float64)
	MtuInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OwnerAccountId() *string
	SetOwnerAccountId(val *string)
	OwnerAccountIdInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DxHostedPrivateVirtualInterfaceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Vlan() *float64
	SetVlan(val *float64)
	VlanInput() *float64
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
	PutTimeouts(value *DxHostedPrivateVirtualInterfaceTimeouts)
	ResetAmazonAddress()
	ResetBgpAuthKey()
	ResetCustomerAddress()
	ResetId()
	ResetMtu()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DxHostedPrivateVirtualInterface
type jsiiProxy_DxHostedPrivateVirtualInterface struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) AddressFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) AddressFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) AmazonAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) AmazonAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) AmazonSideAsn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonSideAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) AwsDevice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) BgpAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgpAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) BgpAsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgpAsnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) BgpAuthKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpAuthKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) BgpAuthKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpAuthKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) ConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) CustomerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) CustomerAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) JumboFrameCapable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"jumboFrameCapable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) Mtu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) MtuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) OwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) OwnerAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) Timeouts() DxHostedPrivateVirtualInterfaceTimeoutsOutputReference {
	var returns DxHostedPrivateVirtualInterfaceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) Vlan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface) VlanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/dx_hosted_private_virtual_interface aws_dx_hosted_private_virtual_interface} Resource.
func NewDxHostedPrivateVirtualInterface(scope constructs.Construct, id *string, config *DxHostedPrivateVirtualInterfaceConfig) DxHostedPrivateVirtualInterface {
	_init_.Initialize()

	if err := validateNewDxHostedPrivateVirtualInterfaceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DxHostedPrivateVirtualInterface{}

	_jsii_.Create(
		"@cdktf/provider-aws.dxHostedPrivateVirtualInterface.DxHostedPrivateVirtualInterface",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/dx_hosted_private_virtual_interface aws_dx_hosted_private_virtual_interface} Resource.
func NewDxHostedPrivateVirtualInterface_Override(d DxHostedPrivateVirtualInterface, scope constructs.Construct, id *string, config *DxHostedPrivateVirtualInterfaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dxHostedPrivateVirtualInterface.DxHostedPrivateVirtualInterface",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetAddressFamily(val *string) {
	if err := j.validateSetAddressFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressFamily",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetAmazonAddress(val *string) {
	if err := j.validateSetAmazonAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amazonAddress",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetBgpAsn(val *float64) {
	if err := j.validateSetBgpAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpAsn",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetBgpAuthKey(val *string) {
	if err := j.validateSetBgpAuthKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpAuthKey",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetConnectionId(val *string) {
	if err := j.validateSetConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetCustomerAddress(val *string) {
	if err := j.validateSetCustomerAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerAddress",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetMtu(val *float64) {
	if err := j.validateSetMtuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtu",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetOwnerAccountId(val *string) {
	if err := j.validateSetOwnerAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownerAccountId",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DxHostedPrivateVirtualInterface)SetVlan(val *float64) {
	if err := j.validateSetVlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vlan",
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
func DxHostedPrivateVirtualInterface_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDxHostedPrivateVirtualInterface_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dxHostedPrivateVirtualInterface.DxHostedPrivateVirtualInterface",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DxHostedPrivateVirtualInterface_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDxHostedPrivateVirtualInterface_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dxHostedPrivateVirtualInterface.DxHostedPrivateVirtualInterface",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DxHostedPrivateVirtualInterface_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDxHostedPrivateVirtualInterface_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dxHostedPrivateVirtualInterface.DxHostedPrivateVirtualInterface",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DxHostedPrivateVirtualInterface_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dxHostedPrivateVirtualInterface.DxHostedPrivateVirtualInterface",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) PutTimeouts(value *DxHostedPrivateVirtualInterfaceTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) ResetAmazonAddress() {
	_jsii_.InvokeVoid(
		d,
		"resetAmazonAddress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) ResetBgpAuthKey() {
	_jsii_.InvokeVoid(
		d,
		"resetBgpAuthKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) ResetCustomerAddress() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomerAddress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) ResetMtu() {
	_jsii_.InvokeVoid(
		d,
		"resetMtu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DxHostedPrivateVirtualInterface) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

