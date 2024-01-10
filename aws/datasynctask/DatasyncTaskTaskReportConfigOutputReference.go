// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynctask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/datasynctask/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncTaskTaskReportConfigOutputReference interface {
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
	InternalValue() *DatasyncTaskTaskReportConfig
	SetInternalValue(val *DatasyncTaskTaskReportConfig)
	OutputType() *string
	SetOutputType(val *string)
	OutputTypeInput() *string
	ReportLevel() *string
	SetReportLevel(val *string)
	ReportLevelInput() *string
	ReportOverrides() DatasyncTaskTaskReportConfigReportOverridesOutputReference
	ReportOverridesInput() *DatasyncTaskTaskReportConfigReportOverrides
	S3Destination() DatasyncTaskTaskReportConfigS3DestinationOutputReference
	S3DestinationInput() *DatasyncTaskTaskReportConfigS3Destination
	S3ObjectVersioning() *string
	SetS3ObjectVersioning(val *string)
	S3ObjectVersioningInput() *string
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
	PutReportOverrides(value *DatasyncTaskTaskReportConfigReportOverrides)
	PutS3Destination(value *DatasyncTaskTaskReportConfigS3Destination)
	ResetOutputType()
	ResetReportLevel()
	ResetReportOverrides()
	ResetS3ObjectVersioning()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatasyncTaskTaskReportConfigOutputReference
type jsiiProxy_DatasyncTaskTaskReportConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) InternalValue() *DatasyncTaskTaskReportConfig {
	var returns *DatasyncTaskTaskReportConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) OutputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) OutputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) ReportLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) ReportLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) ReportOverrides() DatasyncTaskTaskReportConfigReportOverridesOutputReference {
	var returns DatasyncTaskTaskReportConfigReportOverridesOutputReference
	_jsii_.Get(
		j,
		"reportOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) ReportOverridesInput() *DatasyncTaskTaskReportConfigReportOverrides {
	var returns *DatasyncTaskTaskReportConfigReportOverrides
	_jsii_.Get(
		j,
		"reportOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) S3Destination() DatasyncTaskTaskReportConfigS3DestinationOutputReference {
	var returns DatasyncTaskTaskReportConfigS3DestinationOutputReference
	_jsii_.Get(
		j,
		"s3Destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) S3DestinationInput() *DatasyncTaskTaskReportConfigS3Destination {
	var returns *DatasyncTaskTaskReportConfigS3Destination
	_jsii_.Get(
		j,
		"s3DestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) S3ObjectVersioning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) S3ObjectVersioningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatasyncTaskTaskReportConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatasyncTaskTaskReportConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDatasyncTaskTaskReportConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatasyncTaskTaskReportConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.datasyncTask.DatasyncTaskTaskReportConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatasyncTaskTaskReportConfigOutputReference_Override(d DatasyncTaskTaskReportConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.datasyncTask.DatasyncTaskTaskReportConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference)SetInternalValue(val *DatasyncTaskTaskReportConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference)SetOutputType(val *string) {
	if err := j.validateSetOutputTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputType",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference)SetReportLevel(val *string) {
	if err := j.validateSetReportLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportLevel",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference)SetS3ObjectVersioning(val *string) {
	if err := j.validateSetS3ObjectVersioningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3ObjectVersioning",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) PutReportOverrides(value *DatasyncTaskTaskReportConfigReportOverrides) {
	if err := d.validatePutReportOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReportOverrides",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) PutS3Destination(value *DatasyncTaskTaskReportConfigS3Destination) {
	if err := d.validatePutS3DestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putS3Destination",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) ResetOutputType() {
	_jsii_.InvokeVoid(
		d,
		"resetOutputType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) ResetReportLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetReportLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) ResetReportOverrides() {
	_jsii_.InvokeVoid(
		d,
		"resetReportOverrides",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) ResetS3ObjectVersioning() {
	_jsii_.InvokeVoid(
		d,
		"resetS3ObjectVersioning",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

