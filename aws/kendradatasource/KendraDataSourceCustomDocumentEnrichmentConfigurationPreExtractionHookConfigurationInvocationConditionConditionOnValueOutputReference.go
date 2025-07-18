// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/kendradatasource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference interface {
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
	DateValue() *string
	SetDateValue(val *string)
	DateValueInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValue
	SetInternalValue(val *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValue)
	LongValue() *float64
	SetLongValue(val *float64)
	LongValueInput() *float64
	StringListValue() *[]*string
	SetStringListValue(val *[]*string)
	StringListValueInput() *[]*string
	StringValue() *string
	SetStringValue(val *string)
	StringValueInput() *string
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
	ResetDateValue()
	ResetLongValue()
	ResetStringListValue()
	ResetStringValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference
type jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) DateValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) DateValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) InternalValue() *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValue {
	var returns *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) LongValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) LongValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) StringListValue() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringListValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) StringListValueInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringListValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) StringValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kendraDataSource.KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference_Override(k KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kendraDataSource.KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetDateValue(val *string) {
	if err := j.validateSetDateValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetInternalValue(val *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetLongValue(val *float64) {
	if err := j.validateSetLongValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetStringListValue(val *[]*string) {
	if err := j.validateSetStringListValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringListValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetStringValue(val *string) {
	if err := j.validateSetStringValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ResetDateValue() {
	_jsii_.InvokeVoid(
		k,
		"resetDateValue",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ResetLongValue() {
	_jsii_.InvokeVoid(
		k,
		"resetLongValue",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ResetStringListValue() {
	_jsii_.InvokeVoid(
		k,
		"resetStringListValue",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ResetStringValue() {
	_jsii_.InvokeVoid(
		k,
		"resetStringValue",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

