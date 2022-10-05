package defaultroutetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/defaultroutetable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DefaultRouteTableRouteOutputReference interface {
	cdktf.ComplexObject
	CidrBlock() *string
	SetCidrBlock(val *string)
	CidrBlockInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CoreNetworkArn() *string
	SetCoreNetworkArn(val *string)
	CoreNetworkArnInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationPrefixListId() *string
	SetDestinationPrefixListId(val *string)
	DestinationPrefixListIdInput() *string
	EgressOnlyGatewayId() *string
	SetEgressOnlyGatewayId(val *string)
	EgressOnlyGatewayIdInput() *string
	// Experimental.
	Fqn() *string
	GatewayId() *string
	SetGatewayId(val *string)
	GatewayIdInput() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv6CidrBlock() *string
	SetIpv6CidrBlock(val *string)
	Ipv6CidrBlockInput() *string
	NatGatewayId() *string
	SetNatGatewayId(val *string)
	NatGatewayIdInput() *string
	NetworkInterfaceId() *string
	SetNetworkInterfaceId(val *string)
	NetworkInterfaceIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransitGatewayId() *string
	SetTransitGatewayId(val *string)
	TransitGatewayIdInput() *string
	VpcEndpointId() *string
	SetVpcEndpointId(val *string)
	VpcEndpointIdInput() *string
	VpcPeeringConnectionId() *string
	SetVpcPeeringConnectionId(val *string)
	VpcPeeringConnectionIdInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCidrBlock()
	ResetCoreNetworkArn()
	ResetDestinationPrefixListId()
	ResetEgressOnlyGatewayId()
	ResetGatewayId()
	ResetInstanceId()
	ResetIpv6CidrBlock()
	ResetNatGatewayId()
	ResetNetworkInterfaceId()
	ResetTransitGatewayId()
	ResetVpcEndpointId()
	ResetVpcPeeringConnectionId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DefaultRouteTableRouteOutputReference
type jsiiProxy_DefaultRouteTableRouteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) CoreNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) CoreNetworkArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) DestinationPrefixListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPrefixListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) DestinationPrefixListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPrefixListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) EgressOnlyGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressOnlyGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) EgressOnlyGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressOnlyGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) GatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) Ipv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) Ipv6CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) NatGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) NatGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) NetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) TransitGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) VpcEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) VpcEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) VpcPeeringConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference) VpcPeeringConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringConnectionIdInput",
		&returns,
	)
	return returns
}


func NewDefaultRouteTableRouteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DefaultRouteTableRouteOutputReference {
	_init_.Initialize()

	if err := validateNewDefaultRouteTableRouteOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DefaultRouteTableRouteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.defaultRouteTable.DefaultRouteTableRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDefaultRouteTableRouteOutputReference_Override(d DefaultRouteTableRouteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.defaultRouteTable.DefaultRouteTableRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetCidrBlock(val *string) {
	if err := j.validateSetCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrBlock",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetCoreNetworkArn(val *string) {
	if err := j.validateSetCoreNetworkArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreNetworkArn",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetDestinationPrefixListId(val *string) {
	if err := j.validateSetDestinationPrefixListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPrefixListId",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetEgressOnlyGatewayId(val *string) {
	if err := j.validateSetEgressOnlyGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressOnlyGatewayId",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetGatewayId(val *string) {
	if err := j.validateSetGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayId",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetIpv6CidrBlock(val *string) {
	if err := j.validateSetIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetNatGatewayId(val *string) {
	if err := j.validateSetNatGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natGatewayId",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetNetworkInterfaceId(val *string) {
	if err := j.validateSetNetworkInterfaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetTransitGatewayId(val *string) {
	if err := j.validateSetTransitGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayId",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetVpcEndpointId(val *string) {
	if err := j.validateSetVpcEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcEndpointId",
		val,
	)
}

func (j *jsiiProxy_DefaultRouteTableRouteOutputReference)SetVpcPeeringConnectionId(val *string) {
	if err := j.validateSetVpcPeeringConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcPeeringConnectionId",
		val,
	)
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ResetCidrBlock() {
	_jsii_.InvokeVoid(
		d,
		"resetCidrBlock",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ResetCoreNetworkArn() {
	_jsii_.InvokeVoid(
		d,
		"resetCoreNetworkArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ResetDestinationPrefixListId() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationPrefixListId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ResetEgressOnlyGatewayId() {
	_jsii_.InvokeVoid(
		d,
		"resetEgressOnlyGatewayId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ResetGatewayId() {
	_jsii_.InvokeVoid(
		d,
		"resetGatewayId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ResetInstanceId() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ResetIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6CidrBlock",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ResetNatGatewayId() {
	_jsii_.InvokeVoid(
		d,
		"resetNatGatewayId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ResetNetworkInterfaceId() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkInterfaceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ResetTransitGatewayId() {
	_jsii_.InvokeVoid(
		d,
		"resetTransitGatewayId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ResetVpcEndpointId() {
	_jsii_.InvokeVoid(
		d,
		"resetVpcEndpointId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ResetVpcPeeringConnectionId() {
	_jsii_.InvokeVoid(
		d,
		"resetVpcPeeringConnectionId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultRouteTableRouteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

