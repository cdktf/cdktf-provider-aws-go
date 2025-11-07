// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ssmincidentsresponseplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsmincidentsResponsePlanActionSsmAutomationOutputReference interface {
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
	DocumentName() *string
	SetDocumentName(val *string)
	DocumentNameInput() *string
	DocumentVersion() *string
	SetDocumentVersion(val *string)
	DocumentVersionInput() *string
	DynamicParameters() *map[string]*string
	SetDynamicParameters(val *map[string]*string)
	DynamicParametersInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Parameter() SsmincidentsResponsePlanActionSsmAutomationParameterList
	ParameterInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	TargetAccount() *string
	SetTargetAccount(val *string)
	TargetAccountInput() *string
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
	PutParameter(value interface{})
	ResetDocumentVersion()
	ResetDynamicParameters()
	ResetParameter()
	ResetTargetAccount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SsmincidentsResponsePlanActionSsmAutomationOutputReference
type jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) DocumentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) DocumentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) DocumentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) DocumentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) DynamicParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dynamicParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) DynamicParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dynamicParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) Parameter() SsmincidentsResponsePlanActionSsmAutomationParameterList {
	var returns SsmincidentsResponsePlanActionSsmAutomationParameterList
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) ParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) TargetAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) TargetAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSsmincidentsResponsePlanActionSsmAutomationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SsmincidentsResponsePlanActionSsmAutomationOutputReference {
	_init_.Initialize()

	if err := validateNewSsmincidentsResponsePlanActionSsmAutomationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ssmincidentsResponsePlan.SsmincidentsResponsePlanActionSsmAutomationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSsmincidentsResponsePlanActionSsmAutomationOutputReference_Override(s SsmincidentsResponsePlanActionSsmAutomationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ssmincidentsResponsePlan.SsmincidentsResponsePlanActionSsmAutomationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference)SetDocumentName(val *string) {
	if err := j.validateSetDocumentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentName",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference)SetDocumentVersion(val *string) {
	if err := j.validateSetDocumentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentVersion",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference)SetDynamicParameters(val *map[string]*string) {
	if err := j.validateSetDynamicParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicParameters",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference)SetTargetAccount(val *string) {
	if err := j.validateSetTargetAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetAccount",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) PutParameter(value interface{}) {
	if err := s.validatePutParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putParameter",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) ResetDocumentVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetDocumentVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) ResetDynamicParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetDynamicParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) ResetParameter() {
	_jsii_.InvokeVoid(
		s,
		"resetParameter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) ResetTargetAccount() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetAccount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SsmincidentsResponsePlanActionSsmAutomationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

