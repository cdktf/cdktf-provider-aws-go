// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/sagemakerendpointconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfig
	SetInternalValue(val *SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfig)
	MaxConcurrency() *float64
	SetMaxConcurrency(val *float64)
	MaxConcurrencyInput() *float64
	MemorySizeInMb() *float64
	SetMemorySizeInMb(val *float64)
	MemorySizeInMbInput() *float64
	ProvisionedConcurrency() *float64
	SetProvisionedConcurrency(val *float64)
	ProvisionedConcurrencyInput() *float64
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
	ResetProvisionedConcurrency()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference
type jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) InternalValue() *SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfig {
	var returns *SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) MaxConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) MaxConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) MemorySizeInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) MemorySizeInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) ProvisionedConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) ProvisionedConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerEndpointConfiguration.SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference_Override(s SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerEndpointConfiguration.SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference)SetInternalValue(val *SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference)SetMaxConcurrency(val *float64) {
	if err := j.validateSetMaxConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrency",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference)SetMemorySizeInMb(val *float64) {
	if err := j.validateSetMemorySizeInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySizeInMb",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference)SetProvisionedConcurrency(val *float64) {
	if err := j.validateSetProvisionedConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedConcurrency",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) ResetProvisionedConcurrency() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisionedConcurrency",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

