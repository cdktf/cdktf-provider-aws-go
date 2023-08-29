// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendraindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/kendraindex/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference interface {
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
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	// Experimental.
	Fqn() *string
	Freshness() interface{}
	SetFreshness(val interface{})
	FreshnessInput() interface{}
	Importance() *float64
	SetImportance(val *float64)
	ImportanceInput() *float64
	InternalValue() *KendraIndexDocumentMetadataConfigurationUpdatesRelevance
	SetInternalValue(val *KendraIndexDocumentMetadataConfigurationUpdatesRelevance)
	RankOrder() *string
	SetRankOrder(val *string)
	RankOrderInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValuesImportanceMap() *map[string]*float64
	SetValuesImportanceMap(val *map[string]*float64)
	ValuesImportanceMapInput() *map[string]*float64
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
	ResetDuration()
	ResetFreshness()
	ResetImportance()
	ResetRankOrder()
	ResetValuesImportanceMap()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference
type jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) Freshness() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freshness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) FreshnessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freshnessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) Importance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) ImportanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) InternalValue() *KendraIndexDocumentMetadataConfigurationUpdatesRelevance {
	var returns *KendraIndexDocumentMetadataConfigurationUpdatesRelevance
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) RankOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rankOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) RankOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rankOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) ValuesImportanceMap() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"valuesImportanceMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) ValuesImportanceMapInput() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"valuesImportanceMapInput",
		&returns,
	)
	return returns
}


func NewKendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference {
	_init_.Initialize()

	if err := validateNewKendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kendraIndex.KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference_Override(k KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kendraIndex.KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference)SetFreshness(val interface{}) {
	if err := j.validateSetFreshnessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"freshness",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference)SetImportance(val *float64) {
	if err := j.validateSetImportanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importance",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference)SetInternalValue(val *KendraIndexDocumentMetadataConfigurationUpdatesRelevance) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference)SetRankOrder(val *string) {
	if err := j.validateSetRankOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rankOrder",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference)SetValuesImportanceMap(val *map[string]*float64) {
	if err := j.validateSetValuesImportanceMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valuesImportanceMap",
		val,
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) ResetDuration() {
	_jsii_.InvokeVoid(
		k,
		"resetDuration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) ResetFreshness() {
	_jsii_.InvokeVoid(
		k,
		"resetFreshness",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) ResetImportance() {
	_jsii_.InvokeVoid(
		k,
		"resetImportance",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) ResetRankOrder() {
	_jsii_.InvokeVoid(
		k,
		"resetRankOrder",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) ResetValuesImportanceMap() {
	_jsii_.InvokeVoid(
		k,
		"resetValuesImportanceMap",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationUpdatesRelevanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

