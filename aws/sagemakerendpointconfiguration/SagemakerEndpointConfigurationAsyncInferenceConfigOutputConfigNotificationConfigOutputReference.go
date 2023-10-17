// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/sagemakerendpointconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference interface {
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
	ErrorTopic() *string
	SetErrorTopic(val *string)
	ErrorTopicInput() *string
	// Experimental.
	Fqn() *string
	IncludeInferenceResponseIn() *[]*string
	SetIncludeInferenceResponseIn(val *[]*string)
	IncludeInferenceResponseInInput() *[]*string
	InternalValue() *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig
	SetInternalValue(val *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig)
	SuccessTopic() *string
	SetSuccessTopic(val *string)
	SuccessTopicInput() *string
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
	ResetErrorTopic()
	ResetIncludeInferenceResponseIn()
	ResetSuccessTopic()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference
type jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) ErrorTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) ErrorTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) IncludeInferenceResponseIn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeInferenceResponseIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) IncludeInferenceResponseInInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeInferenceResponseInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) InternalValue() *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig {
	var returns *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) SuccessTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) SuccessTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerEndpointConfiguration.SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference_Override(s SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerEndpointConfiguration.SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference)SetErrorTopic(val *string) {
	if err := j.validateSetErrorTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorTopic",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference)SetIncludeInferenceResponseIn(val *[]*string) {
	if err := j.validateSetIncludeInferenceResponseInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeInferenceResponseIn",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference)SetInternalValue(val *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference)SetSuccessTopic(val *string) {
	if err := j.validateSetSuccessTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successTopic",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) ResetErrorTopic() {
	_jsii_.InvokeVoid(
		s,
		"resetErrorTopic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) ResetIncludeInferenceResponseIn() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeInferenceResponseIn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) ResetSuccessTopic() {
	_jsii_.InvokeVoid(
		s,
		"resetSuccessTopic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

