// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksecsclusterlayer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/opsworksecsclusterlayer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference interface {
	cdktf.ComplexObject
	Alarms() *[]*string
	SetAlarms(val *[]*string)
	AlarmsInput() *[]*string
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
	CpuThreshold() *float64
	SetCpuThreshold(val *float64)
	CpuThresholdInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IgnoreMetricsTime() *float64
	SetIgnoreMetricsTime(val *float64)
	IgnoreMetricsTimeInput() *float64
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InternalValue() *OpsworksEcsClusterLayerLoadBasedAutoScalingUpscaling
	SetInternalValue(val *OpsworksEcsClusterLayerLoadBasedAutoScalingUpscaling)
	LoadThreshold() *float64
	SetLoadThreshold(val *float64)
	LoadThresholdInput() *float64
	MemoryThreshold() *float64
	SetMemoryThreshold(val *float64)
	MemoryThresholdInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThresholdsWaitTime() *float64
	SetThresholdsWaitTime(val *float64)
	ThresholdsWaitTimeInput() *float64
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
	ResetAlarms()
	ResetCpuThreshold()
	ResetIgnoreMetricsTime()
	ResetInstanceCount()
	ResetLoadThreshold()
	ResetMemoryThreshold()
	ResetThresholdsWaitTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference
type jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) Alarms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) AlarmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) CpuThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) CpuThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) IgnoreMetricsTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreMetricsTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) IgnoreMetricsTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreMetricsTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) InternalValue() *OpsworksEcsClusterLayerLoadBasedAutoScalingUpscaling {
	var returns *OpsworksEcsClusterLayerLoadBasedAutoScalingUpscaling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) LoadThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) LoadThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) MemoryThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) MemoryThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) ThresholdsWaitTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdsWaitTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) ThresholdsWaitTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdsWaitTimeInput",
		&returns,
	)
	return returns
}


func NewOpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference {
	_init_.Initialize()

	if err := validateNewOpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksEcsClusterLayer.OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference_Override(o OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksEcsClusterLayer.OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference)SetAlarms(val *[]*string) {
	if err := j.validateSetAlarmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarms",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference)SetCpuThreshold(val *float64) {
	if err := j.validateSetCpuThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuThreshold",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference)SetIgnoreMetricsTime(val *float64) {
	if err := j.validateSetIgnoreMetricsTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreMetricsTime",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference)SetInternalValue(val *OpsworksEcsClusterLayerLoadBasedAutoScalingUpscaling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference)SetLoadThreshold(val *float64) {
	if err := j.validateSetLoadThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadThreshold",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference)SetMemoryThreshold(val *float64) {
	if err := j.validateSetMemoryThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryThreshold",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference)SetThresholdsWaitTime(val *float64) {
	if err := j.validateSetThresholdsWaitTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdsWaitTime",
		val,
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) ResetAlarms() {
	_jsii_.InvokeVoid(
		o,
		"resetAlarms",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) ResetCpuThreshold() {
	_jsii_.InvokeVoid(
		o,
		"resetCpuThreshold",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) ResetIgnoreMetricsTime() {
	_jsii_.InvokeVoid(
		o,
		"resetIgnoreMetricsTime",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) ResetLoadThreshold() {
	_jsii_.InvokeVoid(
		o,
		"resetLoadThreshold",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) ResetMemoryThreshold() {
	_jsii_.InvokeVoid(
		o,
		"resetMemoryThreshold",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) ResetThresholdsWaitTime() {
	_jsii_.InvokeVoid(
		o,
		"resetThresholdsWaitTime",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayerLoadBasedAutoScalingUpscalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

