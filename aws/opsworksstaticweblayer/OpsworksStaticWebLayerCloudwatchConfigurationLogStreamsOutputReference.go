// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksstaticweblayer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/opsworksstaticweblayer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference interface {
	cdktf.ComplexObject
	BatchCount() *float64
	SetBatchCount(val *float64)
	BatchCountInput() *float64
	BatchSize() *float64
	SetBatchSize(val *float64)
	BatchSizeInput() *float64
	BufferDuration() *float64
	SetBufferDuration(val *float64)
	BufferDurationInput() *float64
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
	DatetimeFormat() *string
	SetDatetimeFormat(val *string)
	DatetimeFormatInput() *string
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	File() *string
	SetFile(val *string)
	FileFingerprintLines() *string
	SetFileFingerprintLines(val *string)
	FileFingerprintLinesInput() *string
	FileInput() *string
	// Experimental.
	Fqn() *string
	InitialPosition() *string
	SetInitialPosition(val *string)
	InitialPositionInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogGroupName() *string
	SetLogGroupName(val *string)
	LogGroupNameInput() *string
	MultilineStartPattern() *string
	SetMultilineStartPattern(val *string)
	MultilineStartPatternInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	ResetBatchCount()
	ResetBatchSize()
	ResetBufferDuration()
	ResetDatetimeFormat()
	ResetEncoding()
	ResetFileFingerprintLines()
	ResetInitialPosition()
	ResetMultilineStartPattern()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference
type jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) BatchCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) BatchCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) BufferDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) BufferDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) DatetimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datetimeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) DatetimeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datetimeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) File() *string {
	var returns *string
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) FileFingerprintLines() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileFingerprintLines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) FileFingerprintLinesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileFingerprintLinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) FileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) InitialPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) InitialPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) LogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) MultilineStartPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multilineStartPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) MultilineStartPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multilineStartPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewOpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference {
	_init_.Initialize()

	if err := validateNewOpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksStaticWebLayer.OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference_Override(o OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksStaticWebLayer.OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetBatchCount(val *float64) {
	if err := j.validateSetBatchCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchCount",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetBatchSize(val *float64) {
	if err := j.validateSetBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetBufferDuration(val *float64) {
	if err := j.validateSetBufferDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferDuration",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetDatetimeFormat(val *string) {
	if err := j.validateSetDatetimeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datetimeFormat",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetFile(val *string) {
	if err := j.validateSetFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"file",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetFileFingerprintLines(val *string) {
	if err := j.validateSetFileFingerprintLinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileFingerprintLines",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetInitialPosition(val *string) {
	if err := j.validateSetInitialPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialPosition",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetLogGroupName(val *string) {
	if err := j.validateSetLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetMultilineStartPattern(val *string) {
	if err := j.validateSetMultilineStartPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multilineStartPattern",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) ResetBatchCount() {
	_jsii_.InvokeVoid(
		o,
		"resetBatchCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) ResetBatchSize() {
	_jsii_.InvokeVoid(
		o,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) ResetBufferDuration() {
	_jsii_.InvokeVoid(
		o,
		"resetBufferDuration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) ResetDatetimeFormat() {
	_jsii_.InvokeVoid(
		o,
		"resetDatetimeFormat",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		o,
		"resetEncoding",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) ResetFileFingerprintLines() {
	_jsii_.InvokeVoid(
		o,
		"resetFileFingerprintLines",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) ResetInitialPosition() {
	_jsii_.InvokeVoid(
		o,
		"resetInitialPosition",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) ResetMultilineStartPattern() {
	_jsii_.InvokeVoid(
		o,
		"resetMultilineStartPattern",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationLogStreamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

