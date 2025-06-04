// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/kendradatasource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference interface {
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
	Condition() KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsConditionOutputReference
	ConditionInput() *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsCondition
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DocumentContentDeletion() interface{}
	SetDocumentContentDeletion(val interface{})
	DocumentContentDeletionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Target() KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference
	TargetInput() *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsTarget
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
	PutCondition(value *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsCondition)
	PutTarget(value *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsTarget)
	ResetCondition()
	ResetDocumentContentDeletion()
	ResetTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference
type jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) Condition() KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsConditionOutputReference {
	var returns KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsConditionOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) ConditionInput() *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsCondition {
	var returns *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsCondition
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) DocumentContentDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentContentDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) DocumentContentDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentContentDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) Target() KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference {
	var returns KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) TargetInput() *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsTarget {
	var returns *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsTarget
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kendraDataSource.KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference_Override(k KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kendraDataSource.KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference)SetDocumentContentDeletion(val interface{}) {
	if err := j.validateSetDocumentContentDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentContentDeletion",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) PutCondition(value *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsCondition) {
	if err := k.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCondition",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) PutTarget(value *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsTarget) {
	if err := k.validatePutTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTarget",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		k,
		"resetCondition",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) ResetDocumentContentDeletion() {
	_jsii_.InvokeVoid(
		k,
		"resetDocumentContentDeletion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		k,
		"resetTarget",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

