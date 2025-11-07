// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dlmlifecyclepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DlmLifecyclePolicyPolicyDetailsOutputReference interface {
	cdktf.ComplexObject
	Action() DlmLifecyclePolicyPolicyDetailsActionOutputReference
	ActionInput() *DlmLifecyclePolicyPolicyDetailsAction
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
	CopyTags() interface{}
	SetCopyTags(val interface{})
	CopyTagsInput() interface{}
	CreateInterval() *float64
	SetCreateInterval(val *float64)
	CreateIntervalInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EventSource() DlmLifecyclePolicyPolicyDetailsEventSourceOutputReference
	EventSourceInput() *DlmLifecyclePolicyPolicyDetailsEventSource
	Exclusions() DlmLifecyclePolicyPolicyDetailsExclusionsOutputReference
	ExclusionsInput() *DlmLifecyclePolicyPolicyDetailsExclusions
	ExtendDeletion() interface{}
	SetExtendDeletion(val interface{})
	ExtendDeletionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DlmLifecyclePolicyPolicyDetails
	SetInternalValue(val *DlmLifecyclePolicyPolicyDetails)
	Parameters() DlmLifecyclePolicyPolicyDetailsParametersOutputReference
	ParametersInput() *DlmLifecyclePolicyPolicyDetailsParameters
	PolicyLanguage() *string
	SetPolicyLanguage(val *string)
	PolicyLanguageInput() *string
	PolicyType() *string
	SetPolicyType(val *string)
	PolicyTypeInput() *string
	ResourceLocations() *[]*string
	SetResourceLocations(val *[]*string)
	ResourceLocationsInput() *[]*string
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	ResourceTypes() *[]*string
	SetResourceTypes(val *[]*string)
	ResourceTypesInput() *[]*string
	RetainInterval() *float64
	SetRetainInterval(val *float64)
	RetainIntervalInput() *float64
	Schedule() DlmLifecyclePolicyPolicyDetailsScheduleList
	ScheduleInput() interface{}
	TargetTags() *map[string]*string
	SetTargetTags(val *map[string]*string)
	TargetTagsInput() *map[string]*string
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
	PutAction(value *DlmLifecyclePolicyPolicyDetailsAction)
	PutEventSource(value *DlmLifecyclePolicyPolicyDetailsEventSource)
	PutExclusions(value *DlmLifecyclePolicyPolicyDetailsExclusions)
	PutParameters(value *DlmLifecyclePolicyPolicyDetailsParameters)
	PutSchedule(value interface{})
	ResetAction()
	ResetCopyTags()
	ResetCreateInterval()
	ResetEventSource()
	ResetExclusions()
	ResetExtendDeletion()
	ResetParameters()
	ResetPolicyLanguage()
	ResetPolicyType()
	ResetResourceLocations()
	ResetResourceType()
	ResetResourceTypes()
	ResetRetainInterval()
	ResetSchedule()
	ResetTargetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DlmLifecyclePolicyPolicyDetailsOutputReference
type jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) Action() DlmLifecyclePolicyPolicyDetailsActionOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsActionOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ActionInput() *DlmLifecyclePolicyPolicyDetailsAction {
	var returns *DlmLifecyclePolicyPolicyDetailsAction
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) CopyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) CopyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) CreateInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) CreateIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) EventSource() DlmLifecyclePolicyPolicyDetailsEventSourceOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsEventSourceOutputReference
	_jsii_.Get(
		j,
		"eventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) EventSourceInput() *DlmLifecyclePolicyPolicyDetailsEventSource {
	var returns *DlmLifecyclePolicyPolicyDetailsEventSource
	_jsii_.Get(
		j,
		"eventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) Exclusions() DlmLifecyclePolicyPolicyDetailsExclusionsOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsExclusionsOutputReference
	_jsii_.Get(
		j,
		"exclusions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ExclusionsInput() *DlmLifecyclePolicyPolicyDetailsExclusions {
	var returns *DlmLifecyclePolicyPolicyDetailsExclusions
	_jsii_.Get(
		j,
		"exclusionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ExtendDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ExtendDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) InternalValue() *DlmLifecyclePolicyPolicyDetails {
	var returns *DlmLifecyclePolicyPolicyDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) Parameters() DlmLifecyclePolicyPolicyDetailsParametersOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsParametersOutputReference
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ParametersInput() *DlmLifecyclePolicyPolicyDetailsParameters {
	var returns *DlmLifecyclePolicyPolicyDetailsParameters
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) PolicyLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) PolicyLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) PolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResourceLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResourceLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) RetainInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retainInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) RetainIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retainIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) Schedule() DlmLifecyclePolicyPolicyDetailsScheduleList {
	var returns DlmLifecyclePolicyPolicyDetailsScheduleList
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) TargetTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) TargetTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDlmLifecyclePolicyPolicyDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DlmLifecyclePolicyPolicyDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewDlmLifecyclePolicyPolicyDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dlmLifecyclePolicy.DlmLifecyclePolicyPolicyDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDlmLifecyclePolicyPolicyDetailsOutputReference_Override(d DlmLifecyclePolicyPolicyDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dlmLifecyclePolicy.DlmLifecyclePolicyPolicyDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetCopyTags(val interface{}) {
	if err := j.validateSetCopyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTags",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetCreateInterval(val *float64) {
	if err := j.validateSetCreateIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createInterval",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetExtendDeletion(val interface{}) {
	if err := j.validateSetExtendDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendDeletion",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetInternalValue(val *DlmLifecyclePolicyPolicyDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetPolicyLanguage(val *string) {
	if err := j.validateSetPolicyLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyLanguage",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetPolicyType(val *string) {
	if err := j.validateSetPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetResourceLocations(val *[]*string) {
	if err := j.validateSetResourceLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceLocations",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetResourceTypes(val *[]*string) {
	if err := j.validateSetResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypes",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetRetainInterval(val *float64) {
	if err := j.validateSetRetainIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainInterval",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetTargetTags(val *map[string]*string) {
	if err := j.validateSetTargetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetTags",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) PutAction(value *DlmLifecyclePolicyPolicyDetailsAction) {
	if err := d.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) PutEventSource(value *DlmLifecyclePolicyPolicyDetailsEventSource) {
	if err := d.validatePutEventSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) PutExclusions(value *DlmLifecyclePolicyPolicyDetailsExclusions) {
	if err := d.validatePutExclusionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExclusions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) PutParameters(value *DlmLifecyclePolicyPolicyDetailsParameters) {
	if err := d.validatePutParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putParameters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) PutSchedule(value interface{}) {
	if err := d.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchedule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		d,
		"resetAction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetCopyTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCopyTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetCreateInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetEventSource() {
	_jsii_.InvokeVoid(
		d,
		"resetEventSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetExclusions() {
	_jsii_.InvokeVoid(
		d,
		"resetExclusions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetExtendDeletion() {
	_jsii_.InvokeVoid(
		d,
		"resetExtendDeletion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetPolicyLanguage() {
	_jsii_.InvokeVoid(
		d,
		"resetPolicyLanguage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetPolicyType() {
	_jsii_.InvokeVoid(
		d,
		"resetPolicyType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetResourceLocations() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceLocations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetResourceTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetRetainInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetRetainInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResetTargetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

