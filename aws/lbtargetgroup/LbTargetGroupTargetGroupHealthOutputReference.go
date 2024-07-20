// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbtargetgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lbtargetgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbTargetGroupTargetGroupHealthOutputReference interface {
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
	DnsFailover() LbTargetGroupTargetGroupHealthDnsFailoverOutputReference
	DnsFailoverInput() *LbTargetGroupTargetGroupHealthDnsFailover
	// Experimental.
	Fqn() *string
	InternalValue() *LbTargetGroupTargetGroupHealth
	SetInternalValue(val *LbTargetGroupTargetGroupHealth)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnhealthyStateRouting() LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference
	UnhealthyStateRoutingInput() *LbTargetGroupTargetGroupHealthUnhealthyStateRouting
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
	PutDnsFailover(value *LbTargetGroupTargetGroupHealthDnsFailover)
	PutUnhealthyStateRouting(value *LbTargetGroupTargetGroupHealthUnhealthyStateRouting)
	ResetDnsFailover()
	ResetUnhealthyStateRouting()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbTargetGroupTargetGroupHealthOutputReference
type jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) DnsFailover() LbTargetGroupTargetGroupHealthDnsFailoverOutputReference {
	var returns LbTargetGroupTargetGroupHealthDnsFailoverOutputReference
	_jsii_.Get(
		j,
		"dnsFailover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) DnsFailoverInput() *LbTargetGroupTargetGroupHealthDnsFailover {
	var returns *LbTargetGroupTargetGroupHealthDnsFailover
	_jsii_.Get(
		j,
		"dnsFailoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) InternalValue() *LbTargetGroupTargetGroupHealth {
	var returns *LbTargetGroupTargetGroupHealth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) UnhealthyStateRouting() LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference {
	var returns LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference
	_jsii_.Get(
		j,
		"unhealthyStateRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) UnhealthyStateRoutingInput() *LbTargetGroupTargetGroupHealthUnhealthyStateRouting {
	var returns *LbTargetGroupTargetGroupHealthUnhealthyStateRouting
	_jsii_.Get(
		j,
		"unhealthyStateRoutingInput",
		&returns,
	)
	return returns
}


func NewLbTargetGroupTargetGroupHealthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LbTargetGroupTargetGroupHealthOutputReference {
	_init_.Initialize()

	if err := validateNewLbTargetGroupTargetGroupHealthOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lbTargetGroup.LbTargetGroupTargetGroupHealthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLbTargetGroupTargetGroupHealthOutputReference_Override(l LbTargetGroupTargetGroupHealthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lbTargetGroup.LbTargetGroupTargetGroupHealthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference)SetInternalValue(val *LbTargetGroupTargetGroupHealth) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) PutDnsFailover(value *LbTargetGroupTargetGroupHealthDnsFailover) {
	if err := l.validatePutDnsFailoverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDnsFailover",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) PutUnhealthyStateRouting(value *LbTargetGroupTargetGroupHealthUnhealthyStateRouting) {
	if err := l.validatePutUnhealthyStateRoutingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putUnhealthyStateRouting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) ResetDnsFailover() {
	_jsii_.InvokeVoid(
		l,
		"resetDnsFailover",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) ResetUnhealthyStateRouting() {
	_jsii_.InvokeVoid(
		l,
		"resetUnhealthyStateRouting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

