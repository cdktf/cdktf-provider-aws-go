package appmeshvirtualgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/appmeshvirtualgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference interface {
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
	Exact() *[]*string
	SetExact(val *[]*string)
	ExactInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch
	SetInternalValue(val *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference
type jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) Exact() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) ExactInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) InternalValue() *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch {
	var returns *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference {
	_init_.Initialize()

	if err := validateNewAppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualGateway.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference_Override(a AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualGateway.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference)SetExact(val *[]*string) {
	if err := j.validateSetExactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exact",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference)SetInternalValue(val *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

