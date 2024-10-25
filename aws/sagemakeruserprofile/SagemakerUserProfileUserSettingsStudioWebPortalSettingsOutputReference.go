// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/sagemakeruserprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference interface {
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
	HiddenAppTypes() *[]*string
	SetHiddenAppTypes(val *[]*string)
	HiddenAppTypesInput() *[]*string
	HiddenInstanceTypes() *[]*string
	SetHiddenInstanceTypes(val *[]*string)
	HiddenInstanceTypesInput() *[]*string
	HiddenMlTools() *[]*string
	SetHiddenMlTools(val *[]*string)
	HiddenMlToolsInput() *[]*string
	InternalValue() *SagemakerUserProfileUserSettingsStudioWebPortalSettings
	SetInternalValue(val *SagemakerUserProfileUserSettingsStudioWebPortalSettings)
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
	ResetHiddenAppTypes()
	ResetHiddenInstanceTypes()
	ResetHiddenMlTools()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference
type jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) HiddenAppTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenAppTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) HiddenAppTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenAppTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) HiddenInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) HiddenInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) HiddenMlTools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenMlTools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) HiddenMlToolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenMlToolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) InternalValue() *SagemakerUserProfileUserSettingsStudioWebPortalSettings {
	var returns *SagemakerUserProfileUserSettingsStudioWebPortalSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerUserProfile.SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference_Override(s SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerUserProfile.SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference)SetHiddenAppTypes(val *[]*string) {
	if err := j.validateSetHiddenAppTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiddenAppTypes",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference)SetHiddenInstanceTypes(val *[]*string) {
	if err := j.validateSetHiddenInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiddenInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference)SetHiddenMlTools(val *[]*string) {
	if err := j.validateSetHiddenMlToolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiddenMlTools",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference)SetInternalValue(val *SagemakerUserProfileUserSettingsStudioWebPortalSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) ResetHiddenAppTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetHiddenAppTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) ResetHiddenInstanceTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetHiddenInstanceTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) ResetHiddenMlTools() {
	_jsii_.InvokeVoid(
		s,
		"resetHiddenMlTools",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

