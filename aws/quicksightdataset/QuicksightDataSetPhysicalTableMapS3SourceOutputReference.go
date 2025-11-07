// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/quicksightdataset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightDataSetPhysicalTableMapS3SourceOutputReference interface {
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
	DataSourceArn() *string
	SetDataSourceArn(val *string)
	DataSourceArnInput() *string
	// Experimental.
	Fqn() *string
	InputColumns() QuicksightDataSetPhysicalTableMapS3SourceInputColumnsList
	InputColumnsInput() interface{}
	InternalValue() *QuicksightDataSetPhysicalTableMapS3Source
	SetInternalValue(val *QuicksightDataSetPhysicalTableMapS3Source)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UploadSettings() QuicksightDataSetPhysicalTableMapS3SourceUploadSettingsOutputReference
	UploadSettingsInput() *QuicksightDataSetPhysicalTableMapS3SourceUploadSettings
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
	PutInputColumns(value interface{})
	PutUploadSettings(value *QuicksightDataSetPhysicalTableMapS3SourceUploadSettings)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDataSetPhysicalTableMapS3SourceOutputReference
type jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) DataSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) DataSourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) InputColumns() QuicksightDataSetPhysicalTableMapS3SourceInputColumnsList {
	var returns QuicksightDataSetPhysicalTableMapS3SourceInputColumnsList
	_jsii_.Get(
		j,
		"inputColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) InputColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) InternalValue() *QuicksightDataSetPhysicalTableMapS3Source {
	var returns *QuicksightDataSetPhysicalTableMapS3Source
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) UploadSettings() QuicksightDataSetPhysicalTableMapS3SourceUploadSettingsOutputReference {
	var returns QuicksightDataSetPhysicalTableMapS3SourceUploadSettingsOutputReference
	_jsii_.Get(
		j,
		"uploadSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) UploadSettingsInput() *QuicksightDataSetPhysicalTableMapS3SourceUploadSettings {
	var returns *QuicksightDataSetPhysicalTableMapS3SourceUploadSettings
	_jsii_.Get(
		j,
		"uploadSettingsInput",
		&returns,
	)
	return returns
}


func NewQuicksightDataSetPhysicalTableMapS3SourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QuicksightDataSetPhysicalTableMapS3SourceOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSetPhysicalTableMapS3SourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDataSet.QuicksightDataSetPhysicalTableMapS3SourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDataSetPhysicalTableMapS3SourceOutputReference_Override(q QuicksightDataSetPhysicalTableMapS3SourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDataSet.QuicksightDataSetPhysicalTableMapS3SourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference)SetDataSourceArn(val *string) {
	if err := j.validateSetDataSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceArn",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference)SetInternalValue(val *QuicksightDataSetPhysicalTableMapS3Source) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) PutInputColumns(value interface{}) {
	if err := q.validatePutInputColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putInputColumns",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) PutUploadSettings(value *QuicksightDataSetPhysicalTableMapS3SourceUploadSettings) {
	if err := q.validatePutUploadSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putUploadSettings",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapS3SourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

