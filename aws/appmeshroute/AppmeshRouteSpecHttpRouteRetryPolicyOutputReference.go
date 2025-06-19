// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appmeshroute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppmeshRouteSpecHttpRouteRetryPolicyOutputReference interface {
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
	HttpRetryEvents() *[]*string
	SetHttpRetryEvents(val *[]*string)
	HttpRetryEventsInput() *[]*string
	InternalValue() *AppmeshRouteSpecHttpRouteRetryPolicy
	SetInternalValue(val *AppmeshRouteSpecHttpRouteRetryPolicy)
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	PerRetryTimeout() AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeoutOutputReference
	PerRetryTimeoutInput() *AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeout
	TcpRetryEvents() *[]*string
	SetTcpRetryEvents(val *[]*string)
	TcpRetryEventsInput() *[]*string
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
	PutPerRetryTimeout(value *AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeout)
	ResetHttpRetryEvents()
	ResetTcpRetryEvents()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppmeshRouteSpecHttpRouteRetryPolicyOutputReference
type jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) HttpRetryEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpRetryEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) HttpRetryEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpRetryEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) InternalValue() *AppmeshRouteSpecHttpRouteRetryPolicy {
	var returns *AppmeshRouteSpecHttpRouteRetryPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) PerRetryTimeout() AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeoutOutputReference {
	var returns AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeoutOutputReference
	_jsii_.Get(
		j,
		"perRetryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) PerRetryTimeoutInput() *AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeout {
	var returns *AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeout
	_jsii_.Get(
		j,
		"perRetryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) TcpRetryEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tcpRetryEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) TcpRetryEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tcpRetryEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppmeshRouteSpecHttpRouteRetryPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppmeshRouteSpecHttpRouteRetryPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewAppmeshRouteSpecHttpRouteRetryPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshRoute.AppmeshRouteSpecHttpRouteRetryPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppmeshRouteSpecHttpRouteRetryPolicyOutputReference_Override(a AppmeshRouteSpecHttpRouteRetryPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshRoute.AppmeshRouteSpecHttpRouteRetryPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference)SetHttpRetryEvents(val *[]*string) {
	if err := j.validateSetHttpRetryEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpRetryEvents",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference)SetInternalValue(val *AppmeshRouteSpecHttpRouteRetryPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference)SetTcpRetryEvents(val *[]*string) {
	if err := j.validateSetTcpRetryEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpRetryEvents",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) PutPerRetryTimeout(value *AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeout) {
	if err := a.validatePutPerRetryTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPerRetryTimeout",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) ResetHttpRetryEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpRetryEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) ResetTcpRetryEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetTcpRetryEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

