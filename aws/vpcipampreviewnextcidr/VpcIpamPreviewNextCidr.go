// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcipampreviewnextcidr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpcipampreviewnextcidr/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/vpc_ipam_preview_next_cidr aws_vpc_ipam_preview_next_cidr}.
type VpcIpamPreviewNextCidr interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cidr() *string
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
	DisallowedCidrs() *[]*string
	SetDisallowedCidrs(val *[]*string)
	DisallowedCidrsInput() *[]*string
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
	IpamPoolId() *string
	SetIpamPoolId(val *string)
	IpamPoolIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NetmaskLength() *float64
	SetNetmaskLength(val *float64)
	NetmaskLengthInput() *float64
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetDisallowedCidrs()
	ResetId()
	ResetNetmaskLength()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VpcIpamPreviewNextCidr
type jsiiProxy_VpcIpamPreviewNextCidr struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) DisallowedCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disallowedCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) DisallowedCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disallowedCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) IpamPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) NetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) NetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/vpc_ipam_preview_next_cidr aws_vpc_ipam_preview_next_cidr} Resource.
func NewVpcIpamPreviewNextCidr(scope constructs.Construct, id *string, config *VpcIpamPreviewNextCidrConfig) VpcIpamPreviewNextCidr {
	_init_.Initialize()

	if err := validateNewVpcIpamPreviewNextCidrParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcIpamPreviewNextCidr{}

	_jsii_.Create(
		"@cdktf/provider-aws.vpcIpamPreviewNextCidr.VpcIpamPreviewNextCidr",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/vpc_ipam_preview_next_cidr aws_vpc_ipam_preview_next_cidr} Resource.
func NewVpcIpamPreviewNextCidr_Override(v VpcIpamPreviewNextCidr, scope constructs.Construct, id *string, config *VpcIpamPreviewNextCidrConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.vpcIpamPreviewNextCidr.VpcIpamPreviewNextCidr",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr)SetDisallowedCidrs(val *[]*string) {
	if err := j.validateSetDisallowedCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disallowedCidrs",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr)SetIpamPoolId(val *string) {
	if err := j.validateSetIpamPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipamPoolId",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr)SetNetmaskLength(val *float64) {
	if err := j.validateSetNetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netmaskLength",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPreviewNextCidr)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a VpcIpamPreviewNextCidr resource upon running "cdktf plan <stack-name>".
func VpcIpamPreviewNextCidr_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVpcIpamPreviewNextCidr_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcIpamPreviewNextCidr.VpcIpamPreviewNextCidr",
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
func VpcIpamPreviewNextCidr_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcIpamPreviewNextCidr_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcIpamPreviewNextCidr.VpcIpamPreviewNextCidr",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcIpamPreviewNextCidr_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcIpamPreviewNextCidr_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcIpamPreviewNextCidr.VpcIpamPreviewNextCidr",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcIpamPreviewNextCidr_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcIpamPreviewNextCidr_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcIpamPreviewNextCidr.VpcIpamPreviewNextCidr",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpcIpamPreviewNextCidr_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.vpcIpamPreviewNextCidr.VpcIpamPreviewNextCidr",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) ResetDisallowedCidrs() {
	_jsii_.InvokeVoid(
		v,
		"resetDisallowedCidrs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) ResetNetmaskLength() {
	_jsii_.InvokeVoid(
		v,
		"resetNetmaskLength",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPreviewNextCidr) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

