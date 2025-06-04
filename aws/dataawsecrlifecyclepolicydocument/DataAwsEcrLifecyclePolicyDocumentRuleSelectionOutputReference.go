// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsecrlifecyclepolicydocument

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/dataawsecrlifecyclepolicydocument/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference interface {
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
	CountNumber() *float64
	SetCountNumber(val *float64)
	CountNumberInput() *float64
	CountType() *string
	SetCountType(val *string)
	CountTypeInput() *string
	CountUnit() *string
	SetCountUnit(val *string)
	CountUnitInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	TagPatternList() *[]*string
	SetTagPatternList(val *[]*string)
	TagPatternListInput() *[]*string
	TagPrefixList() *[]*string
	SetTagPrefixList(val *[]*string)
	TagPrefixListInput() *[]*string
	TagStatus() *string
	SetTagStatus(val *string)
	TagStatusInput() *string
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
	ResetCountUnit()
	ResetTagPatternList()
	ResetTagPrefixList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference
type jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) CountNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) CountNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) CountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) CountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) CountUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) CountUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) TagPatternList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagPatternList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) TagPatternListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagPatternListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) TagPrefixList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagPrefixList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) TagPrefixListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagPrefixListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) TagStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) TagStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEcrLifecyclePolicyDocument.DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference_Override(d DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEcrLifecyclePolicyDocument.DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference)SetCountNumber(val *float64) {
	if err := j.validateSetCountNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countNumber",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference)SetCountType(val *string) {
	if err := j.validateSetCountTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countType",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference)SetCountUnit(val *string) {
	if err := j.validateSetCountUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countUnit",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference)SetTagPatternList(val *[]*string) {
	if err := j.validateSetTagPatternListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPatternList",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference)SetTagPrefixList(val *[]*string) {
	if err := j.validateSetTagPrefixListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPrefixList",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference)SetTagStatus(val *string) {
	if err := j.validateSetTagStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagStatus",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) ResetCountUnit() {
	_jsii_.InvokeVoid(
		d,
		"resetCountUnit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) ResetTagPatternList() {
	_jsii_.InvokeVoid(
		d,
		"resetTagPatternList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) ResetTagPrefixList() {
	_jsii_.InvokeVoid(
		d,
		"resetTagPrefixList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocumentRuleSelectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

