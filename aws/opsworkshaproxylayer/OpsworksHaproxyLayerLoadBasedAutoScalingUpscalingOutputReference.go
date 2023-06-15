package opsworkshaproxylayer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v15/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v15/opsworkshaproxylayer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference interface {
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
	InternalValue() *OpsworksHaproxyLayerLoadBasedAutoScalingUpscaling
	SetInternalValue(val *OpsworksHaproxyLayerLoadBasedAutoScalingUpscaling)
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

// The jsii proxy struct for OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference
type jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) Alarms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) AlarmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) CpuThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) CpuThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) IgnoreMetricsTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreMetricsTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) IgnoreMetricsTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreMetricsTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) InternalValue() *OpsworksHaproxyLayerLoadBasedAutoScalingUpscaling {
	var returns *OpsworksHaproxyLayerLoadBasedAutoScalingUpscaling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) LoadThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) LoadThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) MemoryThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) MemoryThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) ThresholdsWaitTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdsWaitTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) ThresholdsWaitTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdsWaitTimeInput",
		&returns,
	)
	return returns
}


func NewOpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference {
	_init_.Initialize()

	if err := validateNewOpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksHaproxyLayer.OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference_Override(o OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksHaproxyLayer.OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference)SetAlarms(val *[]*string) {
	if err := j.validateSetAlarmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarms",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference)SetCpuThreshold(val *float64) {
	if err := j.validateSetCpuThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuThreshold",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference)SetIgnoreMetricsTime(val *float64) {
	if err := j.validateSetIgnoreMetricsTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreMetricsTime",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference)SetInternalValue(val *OpsworksHaproxyLayerLoadBasedAutoScalingUpscaling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference)SetLoadThreshold(val *float64) {
	if err := j.validateSetLoadThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadThreshold",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference)SetMemoryThreshold(val *float64) {
	if err := j.validateSetMemoryThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryThreshold",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference)SetThresholdsWaitTime(val *float64) {
	if err := j.validateSetThresholdsWaitTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdsWaitTime",
		val,
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) ResetAlarms() {
	_jsii_.InvokeVoid(
		o,
		"resetAlarms",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) ResetCpuThreshold() {
	_jsii_.InvokeVoid(
		o,
		"resetCpuThreshold",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) ResetIgnoreMetricsTime() {
	_jsii_.InvokeVoid(
		o,
		"resetIgnoreMetricsTime",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) ResetLoadThreshold() {
	_jsii_.InvokeVoid(
		o,
		"resetLoadThreshold",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) ResetMemoryThreshold() {
	_jsii_.InvokeVoid(
		o,
		"resetMemoryThreshold",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) ResetThresholdsWaitTime() {
	_jsii_.InvokeVoid(
		o,
		"resetThresholdsWaitTime",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OpsworksHaproxyLayerLoadBasedAutoScalingUpscalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

