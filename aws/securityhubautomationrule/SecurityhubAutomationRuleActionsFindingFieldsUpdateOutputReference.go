// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/securityhubautomationrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference interface {
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
	Confidence() *float64
	SetConfidence(val *float64)
	ConfidenceInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Criticality() *float64
	SetCriticality(val *float64)
	CriticalityInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Note() SecurityhubAutomationRuleActionsFindingFieldsUpdateNoteList
	NoteInput() interface{}
	RelatedFindings() SecurityhubAutomationRuleActionsFindingFieldsUpdateRelatedFindingsList
	RelatedFindingsInput() interface{}
	Severity() SecurityhubAutomationRuleActionsFindingFieldsUpdateSeverityList
	SeverityInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Types() *[]*string
	SetTypes(val *[]*string)
	TypesInput() *[]*string
	UserDefinedFields() *map[string]*string
	SetUserDefinedFields(val *map[string]*string)
	UserDefinedFieldsInput() *map[string]*string
	VerificationState() *string
	SetVerificationState(val *string)
	VerificationStateInput() *string
	Workflow() SecurityhubAutomationRuleActionsFindingFieldsUpdateWorkflowList
	WorkflowInput() interface{}
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
	PutNote(value interface{})
	PutRelatedFindings(value interface{})
	PutSeverity(value interface{})
	PutWorkflow(value interface{})
	ResetConfidence()
	ResetCriticality()
	ResetNote()
	ResetRelatedFindings()
	ResetSeverity()
	ResetTypes()
	ResetUserDefinedFields()
	ResetVerificationState()
	ResetWorkflow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference
type jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) Confidence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ConfidenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) Criticality() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) CriticalityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) Note() SecurityhubAutomationRuleActionsFindingFieldsUpdateNoteList {
	var returns SecurityhubAutomationRuleActionsFindingFieldsUpdateNoteList
	_jsii_.Get(
		j,
		"note",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) NoteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) RelatedFindings() SecurityhubAutomationRuleActionsFindingFieldsUpdateRelatedFindingsList {
	var returns SecurityhubAutomationRuleActionsFindingFieldsUpdateRelatedFindingsList
	_jsii_.Get(
		j,
		"relatedFindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) RelatedFindingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) Severity() SecurityhubAutomationRuleActionsFindingFieldsUpdateSeverityList {
	var returns SecurityhubAutomationRuleActionsFindingFieldsUpdateSeverityList
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) SeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) Types() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) TypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"typesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) UserDefinedFields() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userDefinedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) UserDefinedFieldsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userDefinedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) VerificationState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verificationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) VerificationStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verificationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) Workflow() SecurityhubAutomationRuleActionsFindingFieldsUpdateWorkflowList {
	var returns SecurityhubAutomationRuleActionsFindingFieldsUpdateWorkflowList
	_jsii_.Get(
		j,
		"workflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) WorkflowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowInput",
		&returns,
	)
	return returns
}


func NewSecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubAutomationRule.SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference_Override(s SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubAutomationRule.SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference)SetConfidence(val *float64) {
	if err := j.validateSetConfidenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidence",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference)SetCriticality(val *float64) {
	if err := j.validateSetCriticalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"criticality",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference)SetTypes(val *[]*string) {
	if err := j.validateSetTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"types",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference)SetUserDefinedFields(val *map[string]*string) {
	if err := j.validateSetUserDefinedFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDefinedFields",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference)SetVerificationState(val *string) {
	if err := j.validateSetVerificationStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verificationState",
		val,
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) PutNote(value interface{}) {
	if err := s.validatePutNoteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNote",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) PutRelatedFindings(value interface{}) {
	if err := s.validatePutRelatedFindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRelatedFindings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) PutSeverity(value interface{}) {
	if err := s.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSeverity",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) PutWorkflow(value interface{}) {
	if err := s.validatePutWorkflowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWorkflow",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ResetConfidence() {
	_jsii_.InvokeVoid(
		s,
		"resetConfidence",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		s,
		"resetCriticality",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ResetNote() {
	_jsii_.InvokeVoid(
		s,
		"resetNote",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ResetRelatedFindings() {
	_jsii_.InvokeVoid(
		s,
		"resetRelatedFindings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		s,
		"resetSeverity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ResetTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ResetUserDefinedFields() {
	_jsii_.InvokeVoid(
		s,
		"resetUserDefinedFields",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ResetVerificationState() {
	_jsii_.InvokeVoid(
		s,
		"resetVerificationState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ResetWorkflow() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkflow",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityhubAutomationRuleActionsFindingFieldsUpdateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

