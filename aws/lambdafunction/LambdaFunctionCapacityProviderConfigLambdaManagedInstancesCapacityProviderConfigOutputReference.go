// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdafunction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/lambdafunction/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference interface {
	cdktf.ComplexObject
	CapacityProviderArn() *string
	SetCapacityProviderArn(val *string)
	CapacityProviderArnInput() *string
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
	ExecutionEnvironmentMemoryGibPerVcpu() *float64
	SetExecutionEnvironmentMemoryGibPerVcpu(val *float64)
	ExecutionEnvironmentMemoryGibPerVcpuInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfig
	SetInternalValue(val *LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfig)
	PerExecutionEnvironmentMaxConcurrency() *float64
	SetPerExecutionEnvironmentMaxConcurrency(val *float64)
	PerExecutionEnvironmentMaxConcurrencyInput() *float64
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
	ResetExecutionEnvironmentMemoryGibPerVcpu()
	ResetPerExecutionEnvironmentMaxConcurrency()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference
type jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) CapacityProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) CapacityProviderArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProviderArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) ExecutionEnvironmentMemoryGibPerVcpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionEnvironmentMemoryGibPerVcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) ExecutionEnvironmentMemoryGibPerVcpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionEnvironmentMemoryGibPerVcpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) InternalValue() *LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfig {
	var returns *LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) PerExecutionEnvironmentMaxConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perExecutionEnvironmentMaxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) PerExecutionEnvironmentMaxConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perExecutionEnvironmentMaxConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaFunction.LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference_Override(l LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaFunction.LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference)SetCapacityProviderArn(val *string) {
	if err := j.validateSetCapacityProviderArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityProviderArn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference)SetExecutionEnvironmentMemoryGibPerVcpu(val *float64) {
	if err := j.validateSetExecutionEnvironmentMemoryGibPerVcpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionEnvironmentMemoryGibPerVcpu",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference)SetInternalValue(val *LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference)SetPerExecutionEnvironmentMaxConcurrency(val *float64) {
	if err := j.validateSetPerExecutionEnvironmentMaxConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perExecutionEnvironmentMaxConcurrency",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) ResetExecutionEnvironmentMemoryGibPerVcpu() {
	_jsii_.InvokeVoid(
		l,
		"resetExecutionEnvironmentMemoryGibPerVcpu",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) ResetPerExecutionEnvironmentMaxConcurrency() {
	_jsii_.InvokeVoid(
		l,
		"resetPerExecutionEnvironmentMaxConcurrency",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

