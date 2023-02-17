package instance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/instance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InstancePrivateDnsNameOptionsOutputReference interface {
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
	InternalValue() *InstancePrivateDnsNameOptions
	SetInternalValue(val *InstancePrivateDnsNameOptions)
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

// The jsii proxy struct for InstancePrivateDnsNameOptionsOutputReference
type jsiiProxy_InstancePrivateDnsNameOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) EnableResourceNameDnsAaaaRecord() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) EnableResourceNameDnsAaaaRecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) EnableResourceNameDnsARecord() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) EnableResourceNameDnsARecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) HostnameType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) HostnameTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) InternalValue() *InstancePrivateDnsNameOptions {
	var returns *InstancePrivateDnsNameOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewInstancePrivateDnsNameOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) InstancePrivateDnsNameOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewInstancePrivateDnsNameOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_InstancePrivateDnsNameOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.instance.InstancePrivateDnsNameOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInstancePrivateDnsNameOptionsOutputReference_Override(i InstancePrivateDnsNameOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.instance.InstancePrivateDnsNameOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference)SetEnableResourceNameDnsAaaaRecord(val interface{}) {
	if err := j.validateSetEnableResourceNameDnsAaaaRecordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsAaaaRecord",
		val,
	)
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference)SetEnableResourceNameDnsARecord(val interface{}) {
	if err := j.validateSetEnableResourceNameDnsARecordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsARecord",
		val,
	)
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference)SetHostnameType(val *string) {
	if err := j.validateSetHostnameTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnameType",
		val,
	)
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference)SetInternalValue(val *InstancePrivateDnsNameOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) ResetEnableResourceNameDnsAaaaRecord() {
	_jsii_.InvokeVoid(
		i,
		"resetEnableResourceNameDnsAaaaRecord",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) ResetEnableResourceNameDnsARecord() {
	_jsii_.InvokeVoid(
		i,
		"resetEnableResourceNameDnsARecord",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) ResetHostnameType() {
	_jsii_.InvokeVoid(
		i,
		"resetHostnameType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePrivateDnsNameOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

