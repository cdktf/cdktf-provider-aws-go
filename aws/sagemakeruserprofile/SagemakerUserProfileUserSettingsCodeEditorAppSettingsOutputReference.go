// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/sagemakeruserprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference interface {
	cdktf.ComplexObject
	AppLifecycleManagement() SagemakerUserProfileUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference
	AppLifecycleManagementInput() *SagemakerUserProfileUserSettingsCodeEditorAppSettingsAppLifecycleManagement
	BuiltInLifecycleConfigArn() *string
	SetBuiltInLifecycleConfigArn(val *string)
	BuiltInLifecycleConfigArnInput() *string
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
	CustomImage() SagemakerUserProfileUserSettingsCodeEditorAppSettingsCustomImageList
	CustomImageInput() interface{}
	DefaultResourceSpec() SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference
	DefaultResourceSpecInput() *SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpec
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerUserProfileUserSettingsCodeEditorAppSettings
	SetInternalValue(val *SagemakerUserProfileUserSettingsCodeEditorAppSettings)
	LifecycleConfigArns() *[]*string
	SetLifecycleConfigArns(val *[]*string)
	LifecycleConfigArnsInput() *[]*string
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
	PutAppLifecycleManagement(value *SagemakerUserProfileUserSettingsCodeEditorAppSettingsAppLifecycleManagement)
	PutCustomImage(value interface{})
	PutDefaultResourceSpec(value *SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpec)
	ResetAppLifecycleManagement()
	ResetBuiltInLifecycleConfigArn()
	ResetCustomImage()
	ResetDefaultResourceSpec()
	ResetLifecycleConfigArns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference
type jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) AppLifecycleManagement() SagemakerUserProfileUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference {
	var returns SagemakerUserProfileUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference
	_jsii_.Get(
		j,
		"appLifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) AppLifecycleManagementInput() *SagemakerUserProfileUserSettingsCodeEditorAppSettingsAppLifecycleManagement {
	var returns *SagemakerUserProfileUserSettingsCodeEditorAppSettingsAppLifecycleManagement
	_jsii_.Get(
		j,
		"appLifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) BuiltInLifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) BuiltInLifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) CustomImage() SagemakerUserProfileUserSettingsCodeEditorAppSettingsCustomImageList {
	var returns SagemakerUserProfileUserSettingsCodeEditorAppSettingsCustomImageList
	_jsii_.Get(
		j,
		"customImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) CustomImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) DefaultResourceSpec() SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference {
	var returns SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) DefaultResourceSpecInput() *SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpec {
	var returns *SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpec
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) InternalValue() *SagemakerUserProfileUserSettingsCodeEditorAppSettings {
	var returns *SagemakerUserProfileUserSettingsCodeEditorAppSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) LifecycleConfigArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) LifecycleConfigArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerUserProfile.SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference_Override(s SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerUserProfile.SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference)SetBuiltInLifecycleConfigArn(val *string) {
	if err := j.validateSetBuiltInLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtInLifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference)SetInternalValue(val *SagemakerUserProfileUserSettingsCodeEditorAppSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference)SetLifecycleConfigArns(val *[]*string) {
	if err := j.validateSetLifecycleConfigArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArns",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) PutAppLifecycleManagement(value *SagemakerUserProfileUserSettingsCodeEditorAppSettingsAppLifecycleManagement) {
	if err := s.validatePutAppLifecycleManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAppLifecycleManagement",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) PutCustomImage(value interface{}) {
	if err := s.validatePutCustomImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomImage",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) PutDefaultResourceSpec(value *SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpec) {
	if err := s.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) ResetAppLifecycleManagement() {
	_jsii_.InvokeVoid(
		s,
		"resetAppLifecycleManagement",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) ResetBuiltInLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		s,
		"resetBuiltInLifecycleConfigArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) ResetCustomImage() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomImage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) ResetLifecycleConfigArns() {
	_jsii_.InvokeVoid(
		s,
		"resetLifecycleConfigArns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

