// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/quicksightdashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightDashboardParametersOutputReference interface {
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
	DateTimeParameters() QuicksightDashboardParametersDateTimeParametersList
	DateTimeParametersInput() interface{}
	DecimalParameters() QuicksightDashboardParametersDecimalParametersList
	DecimalParametersInput() interface{}
	// Experimental.
	Fqn() *string
	IntegerParameters() QuicksightDashboardParametersIntegerParametersList
	IntegerParametersInput() interface{}
	InternalValue() *QuicksightDashboardParameters
	SetInternalValue(val *QuicksightDashboardParameters)
	StringParameters() QuicksightDashboardParametersStringParametersList
	StringParametersInput() interface{}
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
	PutDateTimeParameters(value interface{})
	PutDecimalParameters(value interface{})
	PutIntegerParameters(value interface{})
	PutStringParameters(value interface{})
	ResetDateTimeParameters()
	ResetDecimalParameters()
	ResetIntegerParameters()
	ResetStringParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDashboardParametersOutputReference
type jsiiProxy_QuicksightDashboardParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) DateTimeParameters() QuicksightDashboardParametersDateTimeParametersList {
	var returns QuicksightDashboardParametersDateTimeParametersList
	_jsii_.Get(
		j,
		"dateTimeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) DateTimeParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dateTimeParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) DecimalParameters() QuicksightDashboardParametersDecimalParametersList {
	var returns QuicksightDashboardParametersDecimalParametersList
	_jsii_.Get(
		j,
		"decimalParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) DecimalParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decimalParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) IntegerParameters() QuicksightDashboardParametersIntegerParametersList {
	var returns QuicksightDashboardParametersIntegerParametersList
	_jsii_.Get(
		j,
		"integerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) IntegerParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) InternalValue() *QuicksightDashboardParameters {
	var returns *QuicksightDashboardParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) StringParameters() QuicksightDashboardParametersStringParametersList {
	var returns QuicksightDashboardParametersStringParametersList
	_jsii_.Get(
		j,
		"stringParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) StringParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightDashboardParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QuicksightDashboardParametersOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDashboardParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDashboardParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDashboard.QuicksightDashboardParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDashboardParametersOutputReference_Override(q QuicksightDashboardParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDashboard.QuicksightDashboardParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference)SetInternalValue(val *QuicksightDashboardParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) PutDateTimeParameters(value interface{}) {
	if err := q.validatePutDateTimeParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDateTimeParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) PutDecimalParameters(value interface{}) {
	if err := q.validatePutDecimalParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDecimalParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) PutIntegerParameters(value interface{}) {
	if err := q.validatePutIntegerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putIntegerParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) PutStringParameters(value interface{}) {
	if err := q.validatePutStringParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putStringParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) ResetDateTimeParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetDateTimeParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) ResetDecimalParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetDecimalParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) ResetIntegerParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetIntegerParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) ResetStringParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetStringParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

