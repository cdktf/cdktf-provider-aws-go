// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2allowedimagessettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ec2allowedimagessettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2AllowedImagesSettingsImageCriterionOutputReference interface {
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
	CreationDateCondition() Ec2AllowedImagesSettingsImageCriterionCreationDateConditionList
	CreationDateConditionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeprecationTimeCondition() Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList
	DeprecationTimeConditionInput() interface{}
	// Experimental.
	Fqn() *string
	ImageNames() *[]*string
	SetImageNames(val *[]*string)
	ImageNamesInput() *[]*string
	ImageProviders() *[]*string
	SetImageProviders(val *[]*string)
	ImageProvidersInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MarketplaceProductCodes() *[]*string
	SetMarketplaceProductCodes(val *[]*string)
	MarketplaceProductCodesInput() *[]*string
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
	PutCreationDateCondition(value interface{})
	PutDeprecationTimeCondition(value interface{})
	ResetCreationDateCondition()
	ResetDeprecationTimeCondition()
	ResetImageNames()
	ResetImageProviders()
	ResetMarketplaceProductCodes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2AllowedImagesSettingsImageCriterionOutputReference
type jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) CreationDateCondition() Ec2AllowedImagesSettingsImageCriterionCreationDateConditionList {
	var returns Ec2AllowedImagesSettingsImageCriterionCreationDateConditionList
	_jsii_.Get(
		j,
		"creationDateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) CreationDateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"creationDateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) DeprecationTimeCondition() Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList {
	var returns Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList
	_jsii_.Get(
		j,
		"deprecationTimeCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) DeprecationTimeConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deprecationTimeConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) ImageNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imageNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) ImageNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imageNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) ImageProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imageProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) ImageProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imageProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) MarketplaceProductCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"marketplaceProductCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) MarketplaceProductCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"marketplaceProductCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2AllowedImagesSettingsImageCriterionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2AllowedImagesSettingsImageCriterionOutputReference {
	_init_.Initialize()

	if err := validateNewEc2AllowedImagesSettingsImageCriterionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2AllowedImagesSettings.Ec2AllowedImagesSettingsImageCriterionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2AllowedImagesSettingsImageCriterionOutputReference_Override(e Ec2AllowedImagesSettingsImageCriterionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2AllowedImagesSettings.Ec2AllowedImagesSettingsImageCriterionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference)SetImageNames(val *[]*string) {
	if err := j.validateSetImageNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageNames",
		val,
	)
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference)SetImageProviders(val *[]*string) {
	if err := j.validateSetImageProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageProviders",
		val,
	)
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference)SetMarketplaceProductCodes(val *[]*string) {
	if err := j.validateSetMarketplaceProductCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"marketplaceProductCodes",
		val,
	)
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) PutCreationDateCondition(value interface{}) {
	if err := e.validatePutCreationDateConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCreationDateCondition",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) PutDeprecationTimeCondition(value interface{}) {
	if err := e.validatePutDeprecationTimeConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeprecationTimeCondition",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) ResetCreationDateCondition() {
	_jsii_.InvokeVoid(
		e,
		"resetCreationDateCondition",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) ResetDeprecationTimeCondition() {
	_jsii_.InvokeVoid(
		e,
		"resetDeprecationTimeCondition",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) ResetImageNames() {
	_jsii_.InvokeVoid(
		e,
		"resetImageNames",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) ResetImageProviders() {
	_jsii_.InvokeVoid(
		e,
		"resetImageProviders",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) ResetMarketplaceProductCodes() {
	_jsii_.InvokeVoid(
		e,
		"resetMarketplaceProductCodes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

