package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/ec2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpotInstanceRequestMetadataOptionsOutputReference interface {
	cdktf.ComplexObject
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
	HttpEndpoint() *string
	SetHttpEndpoint(val *string)
	HttpEndpointInput() *string
	HttpPutResponseHopLimit() *float64
	SetHttpPutResponseHopLimit(val *float64)
	HttpPutResponseHopLimitInput() *float64
	HttpTokens() *string
	SetHttpTokens(val *string)
	HttpTokensInput() *string
	InstanceMetadataTags() *string
	SetInstanceMetadataTags(val *string)
	InstanceMetadataTagsInput() *string
	InternalValue() *SpotInstanceRequestMetadataOptions
	SetInternalValue(val *SpotInstanceRequestMetadataOptions)
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
	ResetHttpEndpoint()
	ResetHttpPutResponseHopLimit()
	ResetHttpTokens()
	ResetInstanceMetadataTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpotInstanceRequestMetadataOptionsOutputReference
type jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) HttpEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) HttpEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) HttpPutResponseHopLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPutResponseHopLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) HttpPutResponseHopLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPutResponseHopLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) HttpTokens() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) HttpTokensInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) InstanceMetadataTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceMetadataTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) InstanceMetadataTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceMetadataTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) InternalValue() *SpotInstanceRequestMetadataOptions {
	var returns *SpotInstanceRequestMetadataOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpotInstanceRequestMetadataOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpotInstanceRequestMetadataOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewSpotInstanceRequestMetadataOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2.SpotInstanceRequestMetadataOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpotInstanceRequestMetadataOptionsOutputReference_Override(s SpotInstanceRequestMetadataOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2.SpotInstanceRequestMetadataOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference)SetHttpEndpoint(val *string) {
	if err := j.validateSetHttpEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpEndpoint",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference)SetHttpPutResponseHopLimit(val *float64) {
	if err := j.validateSetHttpPutResponseHopLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPutResponseHopLimit",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference)SetHttpTokens(val *string) {
	if err := j.validateSetHttpTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpTokens",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference)SetInstanceMetadataTags(val *string) {
	if err := j.validateSetInstanceMetadataTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceMetadataTags",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference)SetInternalValue(val *SpotInstanceRequestMetadataOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) ResetHttpEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) ResetHttpPutResponseHopLimit() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpPutResponseHopLimit",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) ResetHttpTokens() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpTokens",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) ResetInstanceMetadataTags() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceMetadataTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestMetadataOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

