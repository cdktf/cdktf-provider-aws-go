// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaceswebdataprotectionsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/workspaceswebdataprotectionsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference interface {
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
	GlobalConfidenceLevel() *float64
	SetGlobalConfidenceLevel(val *float64)
	GlobalConfidenceLevelInput() *float64
	GlobalEnforcedUrls() *[]*string
	SetGlobalEnforcedUrls(val *[]*string)
	GlobalEnforcedUrlsInput() *[]*string
	GlobalExemptUrls() *[]*string
	SetGlobalExemptUrls(val *[]*string)
	GlobalExemptUrlsInput() *[]*string
	InlineRedactionPattern() WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternList
	InlineRedactionPatternInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutInlineRedactionPattern(value interface{})
	ResetGlobalConfidenceLevel()
	ResetGlobalEnforcedUrls()
	ResetGlobalExemptUrls()
	ResetInlineRedactionPattern()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference
type jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GlobalConfidenceLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"globalConfidenceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GlobalConfidenceLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"globalConfidenceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GlobalEnforcedUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"globalEnforcedUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GlobalEnforcedUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"globalEnforcedUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GlobalExemptUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"globalExemptUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GlobalExemptUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"globalExemptUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) InlineRedactionPattern() WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternList {
	var returns WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternList
	_jsii_.Get(
		j,
		"inlineRedactionPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) InlineRedactionPatternInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineRedactionPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.workspaceswebDataProtectionSettings.WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference_Override(w WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.workspaceswebDataProtectionSettings.WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference)SetGlobalConfidenceLevel(val *float64) {
	if err := j.validateSetGlobalConfidenceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalConfidenceLevel",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference)SetGlobalEnforcedUrls(val *[]*string) {
	if err := j.validateSetGlobalEnforcedUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalEnforcedUrls",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference)SetGlobalExemptUrls(val *[]*string) {
	if err := j.validateSetGlobalExemptUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalExemptUrls",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) PutInlineRedactionPattern(value interface{}) {
	if err := w.validatePutInlineRedactionPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putInlineRedactionPattern",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) ResetGlobalConfidenceLevel() {
	_jsii_.InvokeVoid(
		w,
		"resetGlobalConfidenceLevel",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) ResetGlobalEnforcedUrls() {
	_jsii_.InvokeVoid(
		w,
		"resetGlobalEnforcedUrls",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) ResetGlobalExemptUrls() {
	_jsii_.InvokeVoid(
		w,
		"resetGlobalExemptUrls",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) ResetInlineRedactionPattern() {
	_jsii_.InvokeVoid(
		w,
		"resetInlineRedactionPattern",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

