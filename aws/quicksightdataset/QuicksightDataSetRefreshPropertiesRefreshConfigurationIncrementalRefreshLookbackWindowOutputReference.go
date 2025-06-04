// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/quicksightdataset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference interface {
	cdktf.ComplexObject
	ColumnName() *string
	SetColumnName(val *string)
	ColumnNameInput() *string
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
	InternalValue() *QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindow
	SetInternalValue(val *QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindow)
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	SizeUnit() *string
	SetSizeUnit(val *string)
	SizeUnitInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference
type jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) ColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) ColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) InternalValue() *QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindow {
	var returns *QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) SizeUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) SizeUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDataSet.QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference_Override(q QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDataSet.QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference)SetColumnName(val *string) {
	if err := j.validateSetColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnName",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference)SetInternalValue(val *QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference)SetSizeUnit(val *string) {
	if err := j.validateSetSizeUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeUnit",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

