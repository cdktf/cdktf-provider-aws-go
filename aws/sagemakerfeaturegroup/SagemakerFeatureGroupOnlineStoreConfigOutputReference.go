// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerfeaturegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/sagemakerfeaturegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerFeatureGroupOnlineStoreConfigOutputReference interface {
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
	EnableOnlineStore() interface{}
	SetEnableOnlineStore(val interface{})
	EnableOnlineStoreInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerFeatureGroupOnlineStoreConfig
	SetInternalValue(val *SagemakerFeatureGroupOnlineStoreConfig)
	SecurityConfig() SagemakerFeatureGroupOnlineStoreConfigSecurityConfigOutputReference
	SecurityConfigInput() *SagemakerFeatureGroupOnlineStoreConfigSecurityConfig
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TtlDuration() SagemakerFeatureGroupOnlineStoreConfigTtlDurationOutputReference
	TtlDurationInput() *SagemakerFeatureGroupOnlineStoreConfigTtlDuration
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
	PutSecurityConfig(value *SagemakerFeatureGroupOnlineStoreConfigSecurityConfig)
	PutTtlDuration(value *SagemakerFeatureGroupOnlineStoreConfigTtlDuration)
	ResetEnableOnlineStore()
	ResetSecurityConfig()
	ResetStorageType()
	ResetTtlDuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerFeatureGroupOnlineStoreConfigOutputReference
type jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) EnableOnlineStore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOnlineStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) EnableOnlineStoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOnlineStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) InternalValue() *SagemakerFeatureGroupOnlineStoreConfig {
	var returns *SagemakerFeatureGroupOnlineStoreConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) SecurityConfig() SagemakerFeatureGroupOnlineStoreConfigSecurityConfigOutputReference {
	var returns SagemakerFeatureGroupOnlineStoreConfigSecurityConfigOutputReference
	_jsii_.Get(
		j,
		"securityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) SecurityConfigInput() *SagemakerFeatureGroupOnlineStoreConfigSecurityConfig {
	var returns *SagemakerFeatureGroupOnlineStoreConfigSecurityConfig
	_jsii_.Get(
		j,
		"securityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) TtlDuration() SagemakerFeatureGroupOnlineStoreConfigTtlDurationOutputReference {
	var returns SagemakerFeatureGroupOnlineStoreConfigTtlDurationOutputReference
	_jsii_.Get(
		j,
		"ttlDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) TtlDurationInput() *SagemakerFeatureGroupOnlineStoreConfigTtlDuration {
	var returns *SagemakerFeatureGroupOnlineStoreConfigTtlDuration
	_jsii_.Get(
		j,
		"ttlDurationInput",
		&returns,
	)
	return returns
}


func NewSagemakerFeatureGroupOnlineStoreConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerFeatureGroupOnlineStoreConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerFeatureGroupOnlineStoreConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerFeatureGroup.SagemakerFeatureGroupOnlineStoreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerFeatureGroupOnlineStoreConfigOutputReference_Override(s SagemakerFeatureGroupOnlineStoreConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerFeatureGroup.SagemakerFeatureGroupOnlineStoreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference)SetEnableOnlineStore(val interface{}) {
	if err := j.validateSetEnableOnlineStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableOnlineStore",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference)SetInternalValue(val *SagemakerFeatureGroupOnlineStoreConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) PutSecurityConfig(value *SagemakerFeatureGroupOnlineStoreConfigSecurityConfig) {
	if err := s.validatePutSecurityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSecurityConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) PutTtlDuration(value *SagemakerFeatureGroupOnlineStoreConfigTtlDuration) {
	if err := s.validatePutTtlDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTtlDuration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) ResetEnableOnlineStore() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableOnlineStore",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) ResetSecurityConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) ResetStorageType() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) ResetTtlDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetTtlDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

