package vpcpeeringconnectionoptions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v15/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v15/vpcpeeringconnectionoptions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcPeeringConnectionOptionsRequesterOutputReference interface {
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
	InternalValue() *VpcPeeringConnectionOptionsRequester
	SetInternalValue(val *VpcPeeringConnectionOptionsRequester)
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

// The jsii proxy struct for VpcPeeringConnectionOptionsRequesterOutputReference
type jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) AllowClassicLinkToRemoteVpc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClassicLinkToRemoteVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) AllowClassicLinkToRemoteVpcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClassicLinkToRemoteVpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) AllowRemoteVpcDnsResolution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRemoteVpcDnsResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) AllowRemoteVpcDnsResolutionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRemoteVpcDnsResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) AllowVpcToRemoteClassicLink() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVpcToRemoteClassicLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) AllowVpcToRemoteClassicLinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVpcToRemoteClassicLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) InternalValue() *VpcPeeringConnectionOptionsRequester {
	var returns *VpcPeeringConnectionOptionsRequester
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpcPeeringConnectionOptionsRequesterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpcPeeringConnectionOptionsRequesterOutputReference {
	_init_.Initialize()

	if err := validateNewVpcPeeringConnectionOptionsRequesterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.vpcPeeringConnectionOptions.VpcPeeringConnectionOptionsRequesterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpcPeeringConnectionOptionsRequesterOutputReference_Override(v VpcPeeringConnectionOptionsRequesterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.vpcPeeringConnectionOptions.VpcPeeringConnectionOptionsRequesterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference)SetAllowClassicLinkToRemoteVpc(val interface{}) {
	if err := j.validateSetAllowClassicLinkToRemoteVpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClassicLinkToRemoteVpc",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference)SetAllowRemoteVpcDnsResolution(val interface{}) {
	if err := j.validateSetAllowRemoteVpcDnsResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowRemoteVpcDnsResolution",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference)SetAllowVpcToRemoteClassicLink(val interface{}) {
	if err := j.validateSetAllowVpcToRemoteClassicLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowVpcToRemoteClassicLink",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference)SetInternalValue(val *VpcPeeringConnectionOptionsRequester) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) ResetAllowClassicLinkToRemoteVpc() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowClassicLinkToRemoteVpc",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) ResetAllowRemoteVpcDnsResolution() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowRemoteVpcDnsResolution",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) ResetAllowVpcToRemoteClassicLink() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowVpcToRemoteClassicLink",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VpcPeeringConnectionOptionsRequesterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

