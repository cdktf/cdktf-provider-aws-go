package vpcpeeringconnectionaccepter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/vpcpeeringconnectionaccepter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/vpc_peering_connection_accepter aws_vpc_peering_connection_accepter}.
type VpcPeeringConnectionAccepterA interface {
	cdktf.TerraformResource
	Accepter() VpcPeeringConnectionAccepterAccepterOutputReference
	AccepterInput() *VpcPeeringConnectionAccepterAccepter
	AcceptStatus() *string
	AutoAccept() interface{}
	SetAutoAccept(val interface{})
	AutoAcceptInput() interface{}
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
	PeerOwnerId() *string
	PeerRegion() *string
	PeerVpcId() *string
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
	Requester() VpcPeeringConnectionAccepterRequesterOutputReference
	RequesterInput() *VpcPeeringConnectionAccepterRequester
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
	Timeouts() VpcPeeringConnectionAccepterTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcId() *string
	VpcPeeringConnectionId() *string
	SetVpcPeeringConnectionId(val *string)
	VpcPeeringConnectionIdInput() *string
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
	PutAccepter(value *VpcPeeringConnectionAccepterAccepter)
	PutRequester(value *VpcPeeringConnectionAccepterRequester)
	PutTimeouts(value *VpcPeeringConnectionAccepterTimeouts)
	ResetAccepter()
	ResetAutoAccept()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequester()
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

// The jsii proxy struct for VpcPeeringConnectionAccepterA
type jsiiProxy_VpcPeeringConnectionAccepterA struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) Accepter() VpcPeeringConnectionAccepterAccepterOutputReference {
	var returns VpcPeeringConnectionAccepterAccepterOutputReference
	_jsii_.Get(
		j,
		"accepter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) AccepterInput() *VpcPeeringConnectionAccepterAccepter {
	var returns *VpcPeeringConnectionAccepterAccepter
	_jsii_.Get(
		j,
		"accepterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) AcceptStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) AutoAccept() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAccept",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) AutoAcceptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAcceptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) PeerOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) PeerRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) PeerVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) Requester() VpcPeeringConnectionAccepterRequesterOutputReference {
	var returns VpcPeeringConnectionAccepterRequesterOutputReference
	_jsii_.Get(
		j,
		"requester",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) RequesterInput() *VpcPeeringConnectionAccepterRequester {
	var returns *VpcPeeringConnectionAccepterRequester
	_jsii_.Get(
		j,
		"requesterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) Timeouts() VpcPeeringConnectionAccepterTimeoutsOutputReference {
	var returns VpcPeeringConnectionAccepterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) VpcPeeringConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA) VpcPeeringConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringConnectionIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/vpc_peering_connection_accepter aws_vpc_peering_connection_accepter} Resource.
func NewVpcPeeringConnectionAccepterA(scope constructs.Construct, id *string, config *VpcPeeringConnectionAccepterAConfig) VpcPeeringConnectionAccepterA {
	_init_.Initialize()

	if err := validateNewVpcPeeringConnectionAccepterAParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcPeeringConnectionAccepterA{}

	_jsii_.Create(
		"@cdktf/provider-aws.vpcPeeringConnectionAccepter.VpcPeeringConnectionAccepterA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/vpc_peering_connection_accepter aws_vpc_peering_connection_accepter} Resource.
func NewVpcPeeringConnectionAccepterA_Override(v VpcPeeringConnectionAccepterA, scope constructs.Construct, id *string, config *VpcPeeringConnectionAccepterAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.vpcPeeringConnectionAccepter.VpcPeeringConnectionAccepterA",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA)SetAutoAccept(val interface{}) {
	if err := j.validateSetAutoAcceptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAccept",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterA)SetVpcPeeringConnectionId(val *string) {
	if err := j.validateSetVpcPeeringConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcPeeringConnectionId",
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
func VpcPeeringConnectionAccepterA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcPeeringConnectionAccepterA_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcPeeringConnectionAccepter.VpcPeeringConnectionAccepterA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcPeeringConnectionAccepterA_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcPeeringConnectionAccepterA_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcPeeringConnectionAccepter.VpcPeeringConnectionAccepterA",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcPeeringConnectionAccepterA_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcPeeringConnectionAccepterA_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcPeeringConnectionAccepter.VpcPeeringConnectionAccepterA",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpcPeeringConnectionAccepterA_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.vpcPeeringConnectionAccepter.VpcPeeringConnectionAccepterA",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) PutAccepter(value *VpcPeeringConnectionAccepterAccepter) {
	if err := v.validatePutAccepterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAccepter",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) PutRequester(value *VpcPeeringConnectionAccepterRequester) {
	if err := v.validatePutRequesterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRequester",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) PutTimeouts(value *VpcPeeringConnectionAccepterTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) ResetAccepter() {
	_jsii_.InvokeVoid(
		v,
		"resetAccepter",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) ResetAutoAccept() {
	_jsii_.InvokeVoid(
		v,
		"resetAutoAccept",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) ResetRequester() {
	_jsii_.InvokeVoid(
		v,
		"resetRequester",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) ResetTagsAll() {
	_jsii_.InvokeVoid(
		v,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterA) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

