// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaceswebusersettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/workspaceswebusersettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceswebUserSettingsToolbarConfigurationOutputReference interface {
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
	HiddenToolbarItems() *[]*string
	SetHiddenToolbarItems(val *[]*string)
	HiddenToolbarItemsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxDisplayResolution() *string
	SetMaxDisplayResolution(val *string)
	MaxDisplayResolutionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ToolbarType() *string
	SetToolbarType(val *string)
	ToolbarTypeInput() *string
	VisualMode() *string
	SetVisualMode(val *string)
	VisualModeInput() *string
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
	ResetHiddenToolbarItems()
	ResetMaxDisplayResolution()
	ResetToolbarType()
	ResetVisualMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspaceswebUserSettingsToolbarConfigurationOutputReference
type jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) HiddenToolbarItems() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenToolbarItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) HiddenToolbarItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hiddenToolbarItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) MaxDisplayResolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxDisplayResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) MaxDisplayResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxDisplayResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) ToolbarType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolbarType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) ToolbarTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolbarTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) VisualMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) VisualModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualModeInput",
		&returns,
	)
	return returns
}


func NewWorkspaceswebUserSettingsToolbarConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WorkspaceswebUserSettingsToolbarConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspaceswebUserSettingsToolbarConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.workspaceswebUserSettings.WorkspaceswebUserSettingsToolbarConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWorkspaceswebUserSettingsToolbarConfigurationOutputReference_Override(w WorkspaceswebUserSettingsToolbarConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.workspaceswebUserSettings.WorkspaceswebUserSettingsToolbarConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference)SetHiddenToolbarItems(val *[]*string) {
	if err := j.validateSetHiddenToolbarItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiddenToolbarItems",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference)SetMaxDisplayResolution(val *string) {
	if err := j.validateSetMaxDisplayResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDisplayResolution",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference)SetToolbarType(val *string) {
	if err := j.validateSetToolbarTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolbarType",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference)SetVisualMode(val *string) {
	if err := j.validateSetVisualModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualMode",
		val,
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) ResetHiddenToolbarItems() {
	_jsii_.InvokeVoid(
		w,
		"resetHiddenToolbarItems",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) ResetMaxDisplayResolution() {
	_jsii_.InvokeVoid(
		w,
		"resetMaxDisplayResolution",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) ResetToolbarType() {
	_jsii_.InvokeVoid(
		w,
		"resetToolbarType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) ResetVisualMode() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsToolbarConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

