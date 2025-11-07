// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerfeaturegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/sagemakerfeaturegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerFeatureGroupThroughputConfigOutputReference interface {
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
	InternalValue() *SagemakerFeatureGroupThroughputConfig
	SetInternalValue(val *SagemakerFeatureGroupThroughputConfig)
	ProvisionedReadCapacityUnits() *float64
	SetProvisionedReadCapacityUnits(val *float64)
	ProvisionedReadCapacityUnitsInput() *float64
	ProvisionedWriteCapacityUnits() *float64
	SetProvisionedWriteCapacityUnits(val *float64)
	ProvisionedWriteCapacityUnitsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThroughputMode() *string
	SetThroughputMode(val *string)
	ThroughputModeInput() *string
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
	ResetProvisionedReadCapacityUnits()
	ResetProvisionedWriteCapacityUnits()
	ResetThroughputMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerFeatureGroupThroughputConfigOutputReference
type jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) InternalValue() *SagemakerFeatureGroupThroughputConfig {
	var returns *SagemakerFeatureGroupThroughputConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) ProvisionedReadCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedReadCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) ProvisionedReadCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedReadCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) ProvisionedWriteCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedWriteCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) ProvisionedWriteCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedWriteCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) ThroughputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"throughputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) ThroughputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"throughputModeInput",
		&returns,
	)
	return returns
}


func NewSagemakerFeatureGroupThroughputConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerFeatureGroupThroughputConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerFeatureGroupThroughputConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerFeatureGroup.SagemakerFeatureGroupThroughputConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerFeatureGroupThroughputConfigOutputReference_Override(s SagemakerFeatureGroupThroughputConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerFeatureGroup.SagemakerFeatureGroupThroughputConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference)SetInternalValue(val *SagemakerFeatureGroupThroughputConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference)SetProvisionedReadCapacityUnits(val *float64) {
	if err := j.validateSetProvisionedReadCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedReadCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference)SetProvisionedWriteCapacityUnits(val *float64) {
	if err := j.validateSetProvisionedWriteCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedWriteCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference)SetThroughputMode(val *string) {
	if err := j.validateSetThroughputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughputMode",
		val,
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) ResetProvisionedReadCapacityUnits() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisionedReadCapacityUnits",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) ResetProvisionedWriteCapacityUnits() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisionedWriteCapacityUnits",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) ResetThroughputMode() {
	_jsii_.InvokeVoid(
		s,
		"resetThroughputMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerFeatureGroupThroughputConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

