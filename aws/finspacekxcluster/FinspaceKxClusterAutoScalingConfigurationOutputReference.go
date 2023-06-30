package finspacekxcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/finspacekxcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FinspaceKxClusterAutoScalingConfigurationOutputReference interface {
	cdktf.ComplexObject
	AutoScalingMetric() *string
	SetAutoScalingMetric(val *string)
	AutoScalingMetricInput() *string
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
	InternalValue() *FinspaceKxClusterAutoScalingConfiguration
	SetInternalValue(val *FinspaceKxClusterAutoScalingConfiguration)
	MaxNodeCount() *float64
	SetMaxNodeCount(val *float64)
	MaxNodeCountInput() *float64
	MetricTarget() *float64
	SetMetricTarget(val *float64)
	MetricTargetInput() *float64
	MinNodeCount() *float64
	SetMinNodeCount(val *float64)
	MinNodeCountInput() *float64
	ScaleInCooldownSeconds() *float64
	SetScaleInCooldownSeconds(val *float64)
	ScaleInCooldownSecondsInput() *float64
	ScaleOutCooldownSeconds() *float64
	SetScaleOutCooldownSeconds(val *float64)
	ScaleOutCooldownSecondsInput() *float64
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

// The jsii proxy struct for FinspaceKxClusterAutoScalingConfigurationOutputReference
type jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) AutoScalingMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) AutoScalingMetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) InternalValue() *FinspaceKxClusterAutoScalingConfiguration {
	var returns *FinspaceKxClusterAutoScalingConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) MaxNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) MaxNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) MetricTarget() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) MetricTargetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) MinNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) MinNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) ScaleInCooldownSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldownSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) ScaleInCooldownSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldownSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) ScaleOutCooldownSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldownSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) ScaleOutCooldownSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldownSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFinspaceKxClusterAutoScalingConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FinspaceKxClusterAutoScalingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFinspaceKxClusterAutoScalingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.finspaceKxCluster.FinspaceKxClusterAutoScalingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFinspaceKxClusterAutoScalingConfigurationOutputReference_Override(f FinspaceKxClusterAutoScalingConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.finspaceKxCluster.FinspaceKxClusterAutoScalingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference)SetAutoScalingMetric(val *string) {
	if err := j.validateSetAutoScalingMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingMetric",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference)SetInternalValue(val *FinspaceKxClusterAutoScalingConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference)SetMaxNodeCount(val *float64) {
	if err := j.validateSetMaxNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNodeCount",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference)SetMetricTarget(val *float64) {
	if err := j.validateSetMetricTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricTarget",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference)SetMinNodeCount(val *float64) {
	if err := j.validateSetMinNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodeCount",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference)SetScaleInCooldownSeconds(val *float64) {
	if err := j.validateSetScaleInCooldownSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleInCooldownSeconds",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference)SetScaleOutCooldownSeconds(val *float64) {
	if err := j.validateSetScaleOutCooldownSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleOutCooldownSeconds",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxClusterAutoScalingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

