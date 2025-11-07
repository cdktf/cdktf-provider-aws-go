// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ssmincidentsresponseplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsmincidentsResponsePlanIncidentTemplateOutputReference interface {
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
	DedupeString() *string
	SetDedupeString(val *string)
	DedupeStringInput() *string
	// Experimental.
	Fqn() *string
	Impact() *float64
	SetImpact(val *float64)
	ImpactInput() *float64
	IncidentTags() *map[string]*string
	SetIncidentTags(val *map[string]*string)
	IncidentTagsInput() *map[string]*string
	InternalValue() *SsmincidentsResponsePlanIncidentTemplate
	SetInternalValue(val *SsmincidentsResponsePlanIncidentTemplate)
	NotificationTarget() SsmincidentsResponsePlanIncidentTemplateNotificationTargetList
	NotificationTargetInput() interface{}
	Summary() *string
	SetSummary(val *string)
	SummaryInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
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
	PutNotificationTarget(value interface{})
	ResetDedupeString()
	ResetIncidentTags()
	ResetNotificationTarget()
	ResetSummary()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SsmincidentsResponsePlanIncidentTemplateOutputReference
type jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) DedupeString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedupeString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) DedupeStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedupeStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) Impact() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"impact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) ImpactInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"impactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) IncidentTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"incidentTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) IncidentTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"incidentTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) InternalValue() *SsmincidentsResponsePlanIncidentTemplate {
	var returns *SsmincidentsResponsePlanIncidentTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) NotificationTarget() SsmincidentsResponsePlanIncidentTemplateNotificationTargetList {
	var returns SsmincidentsResponsePlanIncidentTemplateNotificationTargetList
	_jsii_.Get(
		j,
		"notificationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) NotificationTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) Summary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) SummaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}


func NewSsmincidentsResponsePlanIncidentTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SsmincidentsResponsePlanIncidentTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewSsmincidentsResponsePlanIncidentTemplateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ssmincidentsResponsePlan.SsmincidentsResponsePlanIncidentTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSsmincidentsResponsePlanIncidentTemplateOutputReference_Override(s SsmincidentsResponsePlanIncidentTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ssmincidentsResponsePlan.SsmincidentsResponsePlanIncidentTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference)SetDedupeString(val *string) {
	if err := j.validateSetDedupeStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedupeString",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference)SetImpact(val *float64) {
	if err := j.validateSetImpactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"impact",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference)SetIncidentTags(val *map[string]*string) {
	if err := j.validateSetIncidentTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incidentTags",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference)SetInternalValue(val *SsmincidentsResponsePlanIncidentTemplate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference)SetSummary(val *string) {
	if err := j.validateSetSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summary",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) PutNotificationTarget(value interface{}) {
	if err := s.validatePutNotificationTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNotificationTarget",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) ResetDedupeString() {
	_jsii_.InvokeVoid(
		s,
		"resetDedupeString",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) ResetIncidentTags() {
	_jsii_.InvokeVoid(
		s,
		"resetIncidentTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) ResetNotificationTarget() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationTarget",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) ResetSummary() {
	_jsii_.InvokeVoid(
		s,
		"resetSummary",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SsmincidentsResponsePlanIncidentTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

