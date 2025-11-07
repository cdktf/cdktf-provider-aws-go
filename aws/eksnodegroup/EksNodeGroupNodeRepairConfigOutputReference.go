// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksnodegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/eksnodegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EksNodeGroupNodeRepairConfigOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *EksNodeGroupNodeRepairConfig
	SetInternalValue(val *EksNodeGroupNodeRepairConfig)
	MaxParallelNodesRepairedCount() *float64
	SetMaxParallelNodesRepairedCount(val *float64)
	MaxParallelNodesRepairedCountInput() *float64
	MaxParallelNodesRepairedPercentage() *float64
	SetMaxParallelNodesRepairedPercentage(val *float64)
	MaxParallelNodesRepairedPercentageInput() *float64
	MaxUnhealthyNodeThresholdCount() *float64
	SetMaxUnhealthyNodeThresholdCount(val *float64)
	MaxUnhealthyNodeThresholdCountInput() *float64
	MaxUnhealthyNodeThresholdPercentage() *float64
	SetMaxUnhealthyNodeThresholdPercentage(val *float64)
	MaxUnhealthyNodeThresholdPercentageInput() *float64
	NodeRepairConfigOverrides() EksNodeGroupNodeRepairConfigNodeRepairConfigOverridesList
	NodeRepairConfigOverridesInput() interface{}
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
	PutNodeRepairConfigOverrides(value interface{})
	ResetEnabled()
	ResetMaxParallelNodesRepairedCount()
	ResetMaxParallelNodesRepairedPercentage()
	ResetMaxUnhealthyNodeThresholdCount()
	ResetMaxUnhealthyNodeThresholdPercentage()
	ResetNodeRepairConfigOverrides()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EksNodeGroupNodeRepairConfigOutputReference
type jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) InternalValue() *EksNodeGroupNodeRepairConfig {
	var returns *EksNodeGroupNodeRepairConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) MaxParallelNodesRepairedCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelNodesRepairedCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) MaxParallelNodesRepairedCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelNodesRepairedCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) MaxParallelNodesRepairedPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelNodesRepairedPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) MaxParallelNodesRepairedPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelNodesRepairedPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) MaxUnhealthyNodeThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodeThresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) MaxUnhealthyNodeThresholdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodeThresholdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) MaxUnhealthyNodeThresholdPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodeThresholdPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) MaxUnhealthyNodeThresholdPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodeThresholdPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) NodeRepairConfigOverrides() EksNodeGroupNodeRepairConfigNodeRepairConfigOverridesList {
	var returns EksNodeGroupNodeRepairConfigNodeRepairConfigOverridesList
	_jsii_.Get(
		j,
		"nodeRepairConfigOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) NodeRepairConfigOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeRepairConfigOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEksNodeGroupNodeRepairConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EksNodeGroupNodeRepairConfigOutputReference {
	_init_.Initialize()

	if err := validateNewEksNodeGroupNodeRepairConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.eksNodeGroup.EksNodeGroupNodeRepairConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEksNodeGroupNodeRepairConfigOutputReference_Override(e EksNodeGroupNodeRepairConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.eksNodeGroup.EksNodeGroupNodeRepairConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference)SetInternalValue(val *EksNodeGroupNodeRepairConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference)SetMaxParallelNodesRepairedCount(val *float64) {
	if err := j.validateSetMaxParallelNodesRepairedCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelNodesRepairedCount",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference)SetMaxParallelNodesRepairedPercentage(val *float64) {
	if err := j.validateSetMaxParallelNodesRepairedPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelNodesRepairedPercentage",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference)SetMaxUnhealthyNodeThresholdCount(val *float64) {
	if err := j.validateSetMaxUnhealthyNodeThresholdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnhealthyNodeThresholdCount",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference)SetMaxUnhealthyNodeThresholdPercentage(val *float64) {
	if err := j.validateSetMaxUnhealthyNodeThresholdPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnhealthyNodeThresholdPercentage",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) PutNodeRepairConfigOverrides(value interface{}) {
	if err := e.validatePutNodeRepairConfigOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNodeRepairConfigOverrides",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) ResetMaxParallelNodesRepairedCount() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxParallelNodesRepairedCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) ResetMaxParallelNodesRepairedPercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxParallelNodesRepairedPercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) ResetMaxUnhealthyNodeThresholdCount() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxUnhealthyNodeThresholdCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) ResetMaxUnhealthyNodeThresholdPercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxUnhealthyNodeThresholdPercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) ResetNodeRepairConfigOverrides() {
	_jsii_.InvokeVoid(
		e,
		"resetNodeRepairConfigOverrides",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EksNodeGroupNodeRepairConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

