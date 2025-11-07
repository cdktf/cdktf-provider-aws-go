// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appmeshvirtualnode/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference interface {
	cdktf.ComplexObject
	BaseEjectionDuration() AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDurationOutputReference
	BaseEjectionDurationInput() *AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDuration
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
	InternalValue() *AppmeshVirtualNodeSpecListenerOutlierDetection
	SetInternalValue(val *AppmeshVirtualNodeSpecListenerOutlierDetection)
	Interval() AppmeshVirtualNodeSpecListenerOutlierDetectionIntervalOutputReference
	IntervalInput() *AppmeshVirtualNodeSpecListenerOutlierDetectionInterval
	MaxEjectionPercent() *float64
	SetMaxEjectionPercent(val *float64)
	MaxEjectionPercentInput() *float64
	MaxServerErrors() *float64
	SetMaxServerErrors(val *float64)
	MaxServerErrorsInput() *float64
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
	PutBaseEjectionDuration(value *AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDuration)
	PutInterval(value *AppmeshVirtualNodeSpecListenerOutlierDetectionInterval)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference
type jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) BaseEjectionDuration() AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDurationOutputReference {
	var returns AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDurationOutputReference
	_jsii_.Get(
		j,
		"baseEjectionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) BaseEjectionDurationInput() *AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDuration {
	var returns *AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDuration
	_jsii_.Get(
		j,
		"baseEjectionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) InternalValue() *AppmeshVirtualNodeSpecListenerOutlierDetection {
	var returns *AppmeshVirtualNodeSpecListenerOutlierDetection
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) Interval() AppmeshVirtualNodeSpecListenerOutlierDetectionIntervalOutputReference {
	var returns AppmeshVirtualNodeSpecListenerOutlierDetectionIntervalOutputReference
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) IntervalInput() *AppmeshVirtualNodeSpecListenerOutlierDetectionInterval {
	var returns *AppmeshVirtualNodeSpecListenerOutlierDetectionInterval
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) MaxEjectionPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEjectionPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) MaxEjectionPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEjectionPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) MaxServerErrors() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxServerErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) MaxServerErrorsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxServerErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference {
	_init_.Initialize()

	if err := validateNewAppmeshVirtualNodeSpecListenerOutlierDetectionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualNode.AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference_Override(a AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualNode.AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference)SetInternalValue(val *AppmeshVirtualNodeSpecListenerOutlierDetection) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference)SetMaxEjectionPercent(val *float64) {
	if err := j.validateSetMaxEjectionPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxEjectionPercent",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference)SetMaxServerErrors(val *float64) {
	if err := j.validateSetMaxServerErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxServerErrors",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) PutBaseEjectionDuration(value *AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDuration) {
	if err := a.validatePutBaseEjectionDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBaseEjectionDuration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) PutInterval(value *AppmeshVirtualNodeSpecListenerOutlierDetectionInterval) {
	if err := a.validatePutIntervalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInterval",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

