// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/sagemakeruserprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference interface {
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
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() *SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpec
	SetInternalValue(val *SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpec)
	LifecycleConfigArn() *string
	SetLifecycleConfigArn(val *string)
	LifecycleConfigArnInput() *string
	SagemakerImageArn() *string
	SetSagemakerImageArn(val *string)
	SagemakerImageArnInput() *string
	SagemakerImageVersionAlias() *string
	SetSagemakerImageVersionAlias(val *string)
	SagemakerImageVersionAliasInput() *string
	SagemakerImageVersionArn() *string
	SetSagemakerImageVersionArn(val *string)
	SagemakerImageVersionArnInput() *string
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
	ResetInstanceType()
	ResetLifecycleConfigArn()
	ResetSagemakerImageArn()
	ResetSagemakerImageVersionAlias()
	ResetSagemakerImageVersionArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference
type jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) InternalValue() *SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpec {
	var returns *SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) LifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) LifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) SagemakerImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) SagemakerImageArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) SagemakerImageVersionAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) SagemakerImageVersionAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) SagemakerImageVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) SagemakerImageVersionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerUserProfile.SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference_Override(s SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerUserProfile.SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference)SetInternalValue(val *SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference)SetLifecycleConfigArn(val *string) {
	if err := j.validateSetLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference)SetSagemakerImageArn(val *string) {
	if err := j.validateSetSagemakerImageArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sagemakerImageArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference)SetSagemakerImageVersionAlias(val *string) {
	if err := j.validateSetSagemakerImageVersionAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sagemakerImageVersionAlias",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference)SetSagemakerImageVersionArn(val *string) {
	if err := j.validateSetSagemakerImageVersionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sagemakerImageVersionArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) ResetLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		s,
		"resetLifecycleConfigArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) ResetSagemakerImageArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSagemakerImageArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) ResetSagemakerImageVersionAlias() {
	_jsii_.InvokeVoid(
		s,
		"resetSagemakerImageVersionAlias",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) ResetSagemakerImageVersionArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSagemakerImageVersionArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

