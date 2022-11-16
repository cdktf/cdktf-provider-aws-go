package vpcpeeringconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/vpcpeeringconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcPeeringConnectionAccepterOutputReference interface {
	cdktf.ComplexObject
	AllowClassicLinkToRemoteVpc() interface{}
	SetAllowClassicLinkToRemoteVpc(val interface{})
	AllowClassicLinkToRemoteVpcInput() interface{}
	AllowRemoteVpcDnsResolution() interface{}
	SetAllowRemoteVpcDnsResolution(val interface{})
	AllowRemoteVpcDnsResolutionInput() interface{}
	AllowVpcToRemoteClassicLink() interface{}
	SetAllowVpcToRemoteClassicLink(val interface{})
	AllowVpcToRemoteClassicLinkInput() interface{}
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *VpcPeeringConnectionAccepter
	SetInternalValue(val *VpcPeeringConnectionAccepter)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetAllowClassicLinkToRemoteVpc()
	ResetAllowRemoteVpcDnsResolution()
	ResetAllowVpcToRemoteClassicLink()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpcPeeringConnectionAccepterOutputReference
type jsiiProxy_VpcPeeringConnectionAccepterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) AllowClassicLinkToRemoteVpc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClassicLinkToRemoteVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) AllowClassicLinkToRemoteVpcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClassicLinkToRemoteVpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) AllowRemoteVpcDnsResolution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRemoteVpcDnsResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) AllowRemoteVpcDnsResolutionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRemoteVpcDnsResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) AllowVpcToRemoteClassicLink() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVpcToRemoteClassicLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) AllowVpcToRemoteClassicLinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVpcToRemoteClassicLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) InternalValue() *VpcPeeringConnectionAccepter {
	var returns *VpcPeeringConnectionAccepter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpcPeeringConnectionAccepterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpcPeeringConnectionAccepterOutputReference {
	_init_.Initialize()

	if err := validateNewVpcPeeringConnectionAccepterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcPeeringConnectionAccepterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.vpcPeeringConnection.VpcPeeringConnectionAccepterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpcPeeringConnectionAccepterOutputReference_Override(v VpcPeeringConnectionAccepterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.vpcPeeringConnection.VpcPeeringConnectionAccepterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference)SetAllowClassicLinkToRemoteVpc(val interface{}) {
	if err := j.validateSetAllowClassicLinkToRemoteVpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClassicLinkToRemoteVpc",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference)SetAllowRemoteVpcDnsResolution(val interface{}) {
	if err := j.validateSetAllowRemoteVpcDnsResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowRemoteVpcDnsResolution",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference)SetAllowVpcToRemoteClassicLink(val interface{}) {
	if err := j.validateSetAllowVpcToRemoteClassicLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowVpcToRemoteClassicLink",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference)SetInternalValue(val *VpcPeeringConnectionAccepter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) ResetAllowClassicLinkToRemoteVpc() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowClassicLinkToRemoteVpc",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) ResetAllowRemoteVpcDnsResolution() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowRemoteVpcDnsResolution",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) ResetAllowVpcToRemoteClassicLink() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowVpcToRemoteClassicLink",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

