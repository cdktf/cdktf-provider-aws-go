package dxgatewayassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/dxgatewayassociation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/dx_gateway_association aws_dx_gateway_association}.
type DxGatewayAssociation interface {
	cdktf.TerraformResource
	AllowedPrefixes() *[]*string
	SetAllowedPrefixes(val *[]*string)
	AllowedPrefixesInput() *[]*string
	AssociatedGatewayId() *string
	SetAssociatedGatewayId(val *string)
	AssociatedGatewayIdInput() *string
	AssociatedGatewayOwnerAccountId() *string
	SetAssociatedGatewayOwnerAccountId(val *string)
	AssociatedGatewayOwnerAccountIdInput() *string
	AssociatedGatewayType() *string
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
	DxGatewayAssociationId() *string
	DxGatewayId() *string
	SetDxGatewayId(val *string)
	DxGatewayIdInput() *string
	DxGatewayOwnerAccountId() *string
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
	ProposalId() *string
	SetProposalId(val *string)
	ProposalIdInput() *string
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
	Timeouts() DxGatewayAssociationTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpnGatewayId() *string
	SetVpnGatewayId(val *string)
	VpnGatewayIdInput() *string
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
	PutTimeouts(value *DxGatewayAssociationTimeouts)
	ResetAllowedPrefixes()
	ResetAssociatedGatewayId()
	ResetAssociatedGatewayOwnerAccountId()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProposalId()
	ResetTimeouts()
	ResetVpnGatewayId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DxGatewayAssociation
type jsiiProxy_DxGatewayAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DxGatewayAssociation) AllowedPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) AllowedPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) AssociatedGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) AssociatedGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) AssociatedGatewayOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedGatewayOwnerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) AssociatedGatewayOwnerAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedGatewayOwnerAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) AssociatedGatewayType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedGatewayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) DxGatewayAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dxGatewayAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) DxGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dxGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) DxGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dxGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) DxGatewayOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dxGatewayOwnerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) ProposalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proposalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) ProposalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proposalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) Timeouts() DxGatewayAssociationTimeoutsOutputReference {
	var returns DxGatewayAssociationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) VpnGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DxGatewayAssociation) VpnGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/dx_gateway_association aws_dx_gateway_association} Resource.
func NewDxGatewayAssociation(scope constructs.Construct, id *string, config *DxGatewayAssociationConfig) DxGatewayAssociation {
	_init_.Initialize()

	if err := validateNewDxGatewayAssociationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DxGatewayAssociation{}

	_jsii_.Create(
		"@cdktf/provider-aws.dxGatewayAssociation.DxGatewayAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/dx_gateway_association aws_dx_gateway_association} Resource.
func NewDxGatewayAssociation_Override(d DxGatewayAssociation, scope constructs.Construct, id *string, config *DxGatewayAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dxGatewayAssociation.DxGatewayAssociation",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetAllowedPrefixes(val *[]*string) {
	if err := j.validateSetAllowedPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedPrefixes",
		val,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetAssociatedGatewayId(val *string) {
	if err := j.validateSetAssociatedGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatedGatewayId",
		val,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetAssociatedGatewayOwnerAccountId(val *string) {
	if err := j.validateSetAssociatedGatewayOwnerAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatedGatewayOwnerAccountId",
		val,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetDxGatewayId(val *string) {
	if err := j.validateSetDxGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dxGatewayId",
		val,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetProposalId(val *string) {
	if err := j.validateSetProposalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proposalId",
		val,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DxGatewayAssociation)SetVpnGatewayId(val *string) {
	if err := j.validateSetVpnGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnGatewayId",
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
func DxGatewayAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDxGatewayAssociation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dxGatewayAssociation.DxGatewayAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DxGatewayAssociation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDxGatewayAssociation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dxGatewayAssociation.DxGatewayAssociation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DxGatewayAssociation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDxGatewayAssociation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dxGatewayAssociation.DxGatewayAssociation",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DxGatewayAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dxGatewayAssociation.DxGatewayAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DxGatewayAssociation) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DxGatewayAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DxGatewayAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DxGatewayAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DxGatewayAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DxGatewayAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DxGatewayAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DxGatewayAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DxGatewayAssociation) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DxGatewayAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DxGatewayAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DxGatewayAssociation) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DxGatewayAssociation) PutTimeouts(value *DxGatewayAssociationTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DxGatewayAssociation) ResetAllowedPrefixes() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedPrefixes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxGatewayAssociation) ResetAssociatedGatewayId() {
	_jsii_.InvokeVoid(
		d,
		"resetAssociatedGatewayId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxGatewayAssociation) ResetAssociatedGatewayOwnerAccountId() {
	_jsii_.InvokeVoid(
		d,
		"resetAssociatedGatewayOwnerAccountId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxGatewayAssociation) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxGatewayAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxGatewayAssociation) ResetProposalId() {
	_jsii_.InvokeVoid(
		d,
		"resetProposalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxGatewayAssociation) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxGatewayAssociation) ResetVpnGatewayId() {
	_jsii_.InvokeVoid(
		d,
		"resetVpnGatewayId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DxGatewayAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DxGatewayAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DxGatewayAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DxGatewayAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

