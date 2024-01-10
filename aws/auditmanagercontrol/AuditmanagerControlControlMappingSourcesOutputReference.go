// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auditmanagercontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/auditmanagercontrol/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuditmanagerControlControlMappingSourcesOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SourceDescription() *string
	SetSourceDescription(val *string)
	SourceDescriptionInput() *string
	SourceFrequency() *string
	SetSourceFrequency(val *string)
	SourceFrequencyInput() *string
	SourceId() *string
	SourceKeyword() AuditmanagerControlControlMappingSourcesSourceKeywordList
	SourceKeywordInput() interface{}
	SourceName() *string
	SetSourceName(val *string)
	SourceNameInput() *string
	SourceSetUpOption() *string
	SetSourceSetUpOption(val *string)
	SourceSetUpOptionInput() *string
	SourceType() *string
	SetSourceType(val *string)
	SourceTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TroubleshootingText() *string
	SetTroubleshootingText(val *string)
	TroubleshootingTextInput() *string
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
	PutSourceKeyword(value interface{})
	ResetSourceDescription()
	ResetSourceFrequency()
	ResetSourceKeyword()
	ResetTroubleshootingText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AuditmanagerControlControlMappingSourcesOutputReference
type jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) SourceDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) SourceDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) SourceFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) SourceFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) SourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) SourceKeyword() AuditmanagerControlControlMappingSourcesSourceKeywordList {
	var returns AuditmanagerControlControlMappingSourcesSourceKeywordList
	_jsii_.Get(
		j,
		"sourceKeyword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) SourceKeywordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceKeywordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) SourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) SourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) SourceSetUpOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSetUpOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) SourceSetUpOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSetUpOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) SourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) SourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) TroubleshootingText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"troubleshootingText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) TroubleshootingTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"troubleshootingTextInput",
		&returns,
	)
	return returns
}


func NewAuditmanagerControlControlMappingSourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AuditmanagerControlControlMappingSourcesOutputReference {
	_init_.Initialize()

	if err := validateNewAuditmanagerControlControlMappingSourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.auditmanagerControl.AuditmanagerControlControlMappingSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAuditmanagerControlControlMappingSourcesOutputReference_Override(a AuditmanagerControlControlMappingSourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.auditmanagerControl.AuditmanagerControlControlMappingSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference)SetSourceDescription(val *string) {
	if err := j.validateSetSourceDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDescription",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference)SetSourceFrequency(val *string) {
	if err := j.validateSetSourceFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFrequency",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference)SetSourceName(val *string) {
	if err := j.validateSetSourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceName",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference)SetSourceSetUpOption(val *string) {
	if err := j.validateSetSourceSetUpOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSetUpOption",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference)SetSourceType(val *string) {
	if err := j.validateSetSourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceType",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference)SetTroubleshootingText(val *string) {
	if err := j.validateSetTroubleshootingTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"troubleshootingText",
		val,
	)
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) PutSourceKeyword(value interface{}) {
	if err := a.validatePutSourceKeywordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceKeyword",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) ResetSourceDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) ResetSourceFrequency() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceFrequency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) ResetSourceKeyword() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceKeyword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) ResetTroubleshootingText() {
	_jsii_.InvokeVoid(
		a,
		"resetTroubleshootingText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerControlControlMappingSourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

