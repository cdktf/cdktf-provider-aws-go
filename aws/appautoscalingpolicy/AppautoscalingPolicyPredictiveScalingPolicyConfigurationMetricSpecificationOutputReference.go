// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appautoscalingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appautoscalingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference interface {
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
	CustomizedCapacityMetricSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationOutputReference
	CustomizedCapacityMetricSpecificationInput() *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedCapacityMetricSpecification
	CustomizedLoadMetricSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedLoadMetricSpecificationOutputReference
	CustomizedLoadMetricSpecificationInput() *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedLoadMetricSpecification
	CustomizedScalingMetricSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedScalingMetricSpecificationOutputReference
	CustomizedScalingMetricSpecificationInput() *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedScalingMetricSpecification
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PredefinedLoadMetricSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedLoadMetricSpecificationOutputReference
	PredefinedLoadMetricSpecificationInput() *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedLoadMetricSpecification
	PredefinedMetricPairSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedMetricPairSpecificationOutputReference
	PredefinedMetricPairSpecificationInput() *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedMetricPairSpecification
	PredefinedScalingMetricSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedScalingMetricSpecificationOutputReference
	PredefinedScalingMetricSpecificationInput() *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedScalingMetricSpecification
	TargetValue() *string
	SetTargetValue(val *string)
	TargetValueInput() *string
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
	PutCustomizedCapacityMetricSpecification(value *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedCapacityMetricSpecification)
	PutCustomizedLoadMetricSpecification(value *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedLoadMetricSpecification)
	PutCustomizedScalingMetricSpecification(value *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedScalingMetricSpecification)
	PutPredefinedLoadMetricSpecification(value *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedLoadMetricSpecification)
	PutPredefinedMetricPairSpecification(value *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedMetricPairSpecification)
	PutPredefinedScalingMetricSpecification(value *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedScalingMetricSpecification)
	ResetCustomizedCapacityMetricSpecification()
	ResetCustomizedLoadMetricSpecification()
	ResetCustomizedScalingMetricSpecification()
	ResetPredefinedLoadMetricSpecification()
	ResetPredefinedMetricPairSpecification()
	ResetPredefinedScalingMetricSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference
type jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) CustomizedCapacityMetricSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationOutputReference {
	var returns AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"customizedCapacityMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) CustomizedCapacityMetricSpecificationInput() *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedCapacityMetricSpecification {
	var returns *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedCapacityMetricSpecification
	_jsii_.Get(
		j,
		"customizedCapacityMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) CustomizedLoadMetricSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedLoadMetricSpecificationOutputReference {
	var returns AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedLoadMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) CustomizedLoadMetricSpecificationInput() *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedLoadMetricSpecification {
	var returns *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedLoadMetricSpecification
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) CustomizedScalingMetricSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedScalingMetricSpecificationOutputReference {
	var returns AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedScalingMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"customizedScalingMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) CustomizedScalingMetricSpecificationInput() *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedScalingMetricSpecification {
	var returns *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedScalingMetricSpecification
	_jsii_.Get(
		j,
		"customizedScalingMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) PredefinedLoadMetricSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedLoadMetricSpecificationOutputReference {
	var returns AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedLoadMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) PredefinedLoadMetricSpecificationInput() *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedLoadMetricSpecification {
	var returns *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedLoadMetricSpecification
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) PredefinedMetricPairSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedMetricPairSpecificationOutputReference {
	var returns AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedMetricPairSpecificationOutputReference
	_jsii_.Get(
		j,
		"predefinedMetricPairSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) PredefinedMetricPairSpecificationInput() *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedMetricPairSpecification {
	var returns *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedMetricPairSpecification
	_jsii_.Get(
		j,
		"predefinedMetricPairSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) PredefinedScalingMetricSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedScalingMetricSpecificationOutputReference {
	var returns AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedScalingMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"predefinedScalingMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) PredefinedScalingMetricSpecificationInput() *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedScalingMetricSpecification {
	var returns *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedScalingMetricSpecification
	_jsii_.Get(
		j,
		"predefinedScalingMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) TargetValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) TargetValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewAppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appautoscalingPolicy.AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference_Override(a AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appautoscalingPolicy.AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference)SetTargetValue(val *string) {
	if err := j.validateSetTargetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetValue",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) PutCustomizedCapacityMetricSpecification(value *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedCapacityMetricSpecification) {
	if err := a.validatePutCustomizedCapacityMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedCapacityMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) PutCustomizedLoadMetricSpecification(value *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedLoadMetricSpecification) {
	if err := a.validatePutCustomizedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) PutCustomizedScalingMetricSpecification(value *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedScalingMetricSpecification) {
	if err := a.validatePutCustomizedScalingMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedScalingMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) PutPredefinedLoadMetricSpecification(value *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedLoadMetricSpecification) {
	if err := a.validatePutPredefinedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) PutPredefinedMetricPairSpecification(value *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedMetricPairSpecification) {
	if err := a.validatePutPredefinedMetricPairSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedMetricPairSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) PutPredefinedScalingMetricSpecification(value *AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationPredefinedScalingMetricSpecification) {
	if err := a.validatePutPredefinedScalingMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedScalingMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) ResetCustomizedCapacityMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedCapacityMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) ResetCustomizedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) ResetCustomizedScalingMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedScalingMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) ResetPredefinedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) ResetPredefinedMetricPairSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedMetricPairSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) ResetPredefinedScalingMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedScalingMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

