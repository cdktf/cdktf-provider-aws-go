// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightspath

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ec2networkinsightspath/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2NetworkInsightsPathFilterAtDestinationOutputReference interface {
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
	DestinationAddress() *string
	SetDestinationAddress(val *string)
	DestinationAddressInput() *string
	DestinationPortRange() Ec2NetworkInsightsPathFilterAtDestinationDestinationPortRangeOutputReference
	DestinationPortRangeInput() *Ec2NetworkInsightsPathFilterAtDestinationDestinationPortRange
	// Experimental.
	Fqn() *string
	InternalValue() *Ec2NetworkInsightsPathFilterAtDestination
	SetInternalValue(val *Ec2NetworkInsightsPathFilterAtDestination)
	SourceAddress() *string
	SetSourceAddress(val *string)
	SourceAddressInput() *string
	SourcePortRange() Ec2NetworkInsightsPathFilterAtDestinationSourcePortRangeOutputReference
	SourcePortRangeInput() *Ec2NetworkInsightsPathFilterAtDestinationSourcePortRange
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutDestinationPortRange(value *Ec2NetworkInsightsPathFilterAtDestinationDestinationPortRange)
	PutSourcePortRange(value *Ec2NetworkInsightsPathFilterAtDestinationSourcePortRange)
	ResetDestinationAddress()
	ResetDestinationPortRange()
	ResetSourceAddress()
	ResetSourcePortRange()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2NetworkInsightsPathFilterAtDestinationOutputReference
type jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) DestinationAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) DestinationAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) DestinationPortRange() Ec2NetworkInsightsPathFilterAtDestinationDestinationPortRangeOutputReference {
	var returns Ec2NetworkInsightsPathFilterAtDestinationDestinationPortRangeOutputReference
	_jsii_.Get(
		j,
		"destinationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) DestinationPortRangeInput() *Ec2NetworkInsightsPathFilterAtDestinationDestinationPortRange {
	var returns *Ec2NetworkInsightsPathFilterAtDestinationDestinationPortRange
	_jsii_.Get(
		j,
		"destinationPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) InternalValue() *Ec2NetworkInsightsPathFilterAtDestination {
	var returns *Ec2NetworkInsightsPathFilterAtDestination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) SourceAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) SourceAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) SourcePortRange() Ec2NetworkInsightsPathFilterAtDestinationSourcePortRangeOutputReference {
	var returns Ec2NetworkInsightsPathFilterAtDestinationSourcePortRangeOutputReference
	_jsii_.Get(
		j,
		"sourcePortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) SourcePortRangeInput() *Ec2NetworkInsightsPathFilterAtDestinationSourcePortRange {
	var returns *Ec2NetworkInsightsPathFilterAtDestinationSourcePortRange
	_jsii_.Get(
		j,
		"sourcePortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2NetworkInsightsPathFilterAtDestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2NetworkInsightsPathFilterAtDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewEc2NetworkInsightsPathFilterAtDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2NetworkInsightsPath.Ec2NetworkInsightsPathFilterAtDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2NetworkInsightsPathFilterAtDestinationOutputReference_Override(e Ec2NetworkInsightsPathFilterAtDestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2NetworkInsightsPath.Ec2NetworkInsightsPathFilterAtDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference)SetDestinationAddress(val *string) {
	if err := j.validateSetDestinationAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddress",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference)SetInternalValue(val *Ec2NetworkInsightsPathFilterAtDestination) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference)SetSourceAddress(val *string) {
	if err := j.validateSetSourceAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddress",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) PutDestinationPortRange(value *Ec2NetworkInsightsPathFilterAtDestinationDestinationPortRange) {
	if err := e.validatePutDestinationPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDestinationPortRange",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) PutSourcePortRange(value *Ec2NetworkInsightsPathFilterAtDestinationSourcePortRange) {
	if err := e.validatePutSourcePortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSourcePortRange",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) ResetDestinationAddress() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationAddress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) ResetDestinationPortRange() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationPortRange",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) ResetSourceAddress() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceAddress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) ResetSourcePortRange() {
	_jsii_.InvokeVoid(
		e,
		"resetSourcePortRange",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsPathFilterAtDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

