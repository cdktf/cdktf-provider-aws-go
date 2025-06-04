// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdafunction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/lambdafunction/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaFunctionLoggingConfigOutputReference interface {
	cdktf.ComplexObject
	ApplicationLogLevel() *string
	SetApplicationLogLevel(val *string)
	ApplicationLogLevelInput() *string
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
	InternalValue() *LambdaFunctionLoggingConfig
	SetInternalValue(val *LambdaFunctionLoggingConfig)
	LogFormat() *string
	SetLogFormat(val *string)
	LogFormatInput() *string
	LogGroup() *string
	SetLogGroup(val *string)
	LogGroupInput() *string
	SystemLogLevel() *string
	SetSystemLogLevel(val *string)
	SystemLogLevelInput() *string
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
	ResetApplicationLogLevel()
	ResetLogGroup()
	ResetSystemLogLevel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LambdaFunctionLoggingConfigOutputReference
type jsiiProxy_LambdaFunctionLoggingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) ApplicationLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) ApplicationLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) InternalValue() *LambdaFunctionLoggingConfig {
	var returns *LambdaFunctionLoggingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) LogFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) LogFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) LogGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) LogGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) SystemLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) SystemLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLambdaFunctionLoggingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LambdaFunctionLoggingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLambdaFunctionLoggingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunctionLoggingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaFunction.LambdaFunctionLoggingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLambdaFunctionLoggingConfigOutputReference_Override(l LambdaFunctionLoggingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaFunction.LambdaFunctionLoggingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference)SetApplicationLogLevel(val *string) {
	if err := j.validateSetApplicationLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationLogLevel",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference)SetInternalValue(val *LambdaFunctionLoggingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference)SetLogFormat(val *string) {
	if err := j.validateSetLogFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logFormat",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference)SetLogGroup(val *string) {
	if err := j.validateSetLogGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroup",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference)SetSystemLogLevel(val *string) {
	if err := j.validateSetSystemLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemLogLevel",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionLoggingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) ResetApplicationLogLevel() {
	_jsii_.InvokeVoid(
		l,
		"resetApplicationLogLevel",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) ResetLogGroup() {
	_jsii_.InvokeVoid(
		l,
		"resetLogGroup",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) ResetSystemLogLevel() {
	_jsii_.InvokeVoid(
		l,
		"resetSystemLogLevel",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LambdaFunctionLoggingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

