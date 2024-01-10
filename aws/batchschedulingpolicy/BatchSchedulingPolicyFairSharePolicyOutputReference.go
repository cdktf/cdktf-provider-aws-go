// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchschedulingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/batchschedulingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchSchedulingPolicyFairSharePolicyOutputReference interface {
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
	ComputeReservation() *float64
	SetComputeReservation(val *float64)
	ComputeReservationInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *BatchSchedulingPolicyFairSharePolicy
	SetInternalValue(val *BatchSchedulingPolicyFairSharePolicy)
	ShareDecaySeconds() *float64
	SetShareDecaySeconds(val *float64)
	ShareDecaySecondsInput() *float64
	ShareDistribution() BatchSchedulingPolicyFairSharePolicyShareDistributionList
	ShareDistributionInput() interface{}
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
	PutShareDistribution(value interface{})
	ResetComputeReservation()
	ResetShareDecaySeconds()
	ResetShareDistribution()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchSchedulingPolicyFairSharePolicyOutputReference
type jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ComputeReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ComputeReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) InternalValue() *BatchSchedulingPolicyFairSharePolicy {
	var returns *BatchSchedulingPolicyFairSharePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ShareDecaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shareDecaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ShareDecaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shareDecaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ShareDistribution() BatchSchedulingPolicyFairSharePolicyShareDistributionList {
	var returns BatchSchedulingPolicyFairSharePolicyShareDistributionList
	_jsii_.Get(
		j,
		"shareDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ShareDistributionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareDistributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchSchedulingPolicyFairSharePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchSchedulingPolicyFairSharePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewBatchSchedulingPolicyFairSharePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.batchSchedulingPolicy.BatchSchedulingPolicyFairSharePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchSchedulingPolicyFairSharePolicyOutputReference_Override(b BatchSchedulingPolicyFairSharePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.batchSchedulingPolicy.BatchSchedulingPolicyFairSharePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference)SetComputeReservation(val *float64) {
	if err := j.validateSetComputeReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeReservation",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference)SetInternalValue(val *BatchSchedulingPolicyFairSharePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference)SetShareDecaySeconds(val *float64) {
	if err := j.validateSetShareDecaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDecaySeconds",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) PutShareDistribution(value interface{}) {
	if err := b.validatePutShareDistributionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putShareDistribution",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ResetComputeReservation() {
	_jsii_.InvokeVoid(
		b,
		"resetComputeReservation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ResetShareDecaySeconds() {
	_jsii_.InvokeVoid(
		b,
		"resetShareDecaySeconds",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ResetShareDistribution() {
	_jsii_.InvokeVoid(
		b,
		"resetShareDistribution",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

