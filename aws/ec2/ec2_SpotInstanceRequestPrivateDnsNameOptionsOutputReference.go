package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/ec2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpotInstanceRequestPrivateDnsNameOptionsOutputReference interface {
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
	EnableResourceNameDnsAaaaRecord() interface{}
	SetEnableResourceNameDnsAaaaRecord(val interface{})
	EnableResourceNameDnsAaaaRecordInput() interface{}
	EnableResourceNameDnsARecord() interface{}
	SetEnableResourceNameDnsARecord(val interface{})
	EnableResourceNameDnsARecordInput() interface{}
	// Experimental.
	Fqn() *string
	HostnameType() *string
	SetHostnameType(val *string)
	HostnameTypeInput() *string
	InternalValue() *SpotInstanceRequestPrivateDnsNameOptions
	SetInternalValue(val *SpotInstanceRequestPrivateDnsNameOptions)
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
	ResetEnableResourceNameDnsAaaaRecord()
	ResetEnableResourceNameDnsARecord()
	ResetHostnameType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpotInstanceRequestPrivateDnsNameOptionsOutputReference
type jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) EnableResourceNameDnsAaaaRecord() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) EnableResourceNameDnsAaaaRecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) EnableResourceNameDnsARecord() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) EnableResourceNameDnsARecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) HostnameType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) HostnameTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) InternalValue() *SpotInstanceRequestPrivateDnsNameOptions {
	var returns *SpotInstanceRequestPrivateDnsNameOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpotInstanceRequestPrivateDnsNameOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpotInstanceRequestPrivateDnsNameOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewSpotInstanceRequestPrivateDnsNameOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2.SpotInstanceRequestPrivateDnsNameOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpotInstanceRequestPrivateDnsNameOptionsOutputReference_Override(s SpotInstanceRequestPrivateDnsNameOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2.SpotInstanceRequestPrivateDnsNameOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference)SetEnableResourceNameDnsAaaaRecord(val interface{}) {
	if err := j.validateSetEnableResourceNameDnsAaaaRecordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsAaaaRecord",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference)SetEnableResourceNameDnsARecord(val interface{}) {
	if err := j.validateSetEnableResourceNameDnsARecordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsARecord",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference)SetHostnameType(val *string) {
	if err := j.validateSetHostnameTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnameType",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference)SetInternalValue(val *SpotInstanceRequestPrivateDnsNameOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) ResetEnableResourceNameDnsAaaaRecord() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableResourceNameDnsAaaaRecord",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) ResetEnableResourceNameDnsARecord() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableResourceNameDnsARecord",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) ResetHostnameType() {
	_jsii_.InvokeVoid(
		s,
		"resetHostnameType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SpotInstanceRequestPrivateDnsNameOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

