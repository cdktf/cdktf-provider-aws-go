// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/kendradatasource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference interface {
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
	ConditionDocumentAttributeKey() *string
	SetConditionDocumentAttributeKey(val *string)
	ConditionDocumentAttributeKeyInput() *string
	ConditionOnValue() KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference
	ConditionOnValueInput() *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValue
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationCondition
	SetInternalValue(val *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationCondition)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
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
	PutConditionOnValue(value *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValue)
	ResetConditionOnValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference
type jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) ConditionDocumentAttributeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionDocumentAttributeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) ConditionDocumentAttributeKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionDocumentAttributeKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) ConditionOnValue() KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference {
	var returns KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference
	_jsii_.Get(
		j,
		"conditionOnValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) ConditionOnValueInput() *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValue {
	var returns *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValue
	_jsii_.Get(
		j,
		"conditionOnValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) InternalValue() *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationCondition {
	var returns *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kendraDataSource.KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference_Override(k KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kendraDataSource.KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference)SetConditionDocumentAttributeKey(val *string) {
	if err := j.validateSetConditionDocumentAttributeKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conditionDocumentAttributeKey",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference)SetInternalValue(val *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationCondition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) PutConditionOnValue(value *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValue) {
	if err := k.validatePutConditionOnValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putConditionOnValue",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) ResetConditionOnValue() {
	_jsii_.InvokeVoid(
		k,
		"resetConditionOnValue",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

