package opsworksmemcachedlayer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/opsworksmemcachedlayer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference interface {
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

// The jsii proxy struct for OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference
type jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) BatchCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) BatchCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) BufferDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) BufferDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) DatetimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datetimeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) DatetimeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datetimeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) File() *string {
	var returns *string
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) FileFingerprintLines() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileFingerprintLines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) FileFingerprintLinesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileFingerprintLinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) FileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) InitialPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) InitialPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) LogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) MultilineStartPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multilineStartPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) MultilineStartPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multilineStartPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewOpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference {
	_init_.Initialize()

	if err := validateNewOpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksMemcachedLayer.OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference_Override(o OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksMemcachedLayer.OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetBatchCount(val *float64) {
	if err := j.validateSetBatchCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchCount",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetBatchSize(val *float64) {
	if err := j.validateSetBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetBufferDuration(val *float64) {
	if err := j.validateSetBufferDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferDuration",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetDatetimeFormat(val *string) {
	if err := j.validateSetDatetimeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datetimeFormat",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetFile(val *string) {
	if err := j.validateSetFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"file",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetFileFingerprintLines(val *string) {
	if err := j.validateSetFileFingerprintLinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileFingerprintLines",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetInitialPosition(val *string) {
	if err := j.validateSetInitialPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialPosition",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetLogGroupName(val *string) {
	if err := j.validateSetLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetMultilineStartPattern(val *string) {
	if err := j.validateSetMultilineStartPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multilineStartPattern",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) ResetBatchCount() {
	_jsii_.InvokeVoid(
		o,
		"resetBatchCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) ResetBatchSize() {
	_jsii_.InvokeVoid(
		o,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) ResetBufferDuration() {
	_jsii_.InvokeVoid(
		o,
		"resetBufferDuration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) ResetDatetimeFormat() {
	_jsii_.InvokeVoid(
		o,
		"resetDatetimeFormat",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		o,
		"resetEncoding",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) ResetFileFingerprintLines() {
	_jsii_.InvokeVoid(
		o,
		"resetFileFingerprintLines",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) ResetInitialPosition() {
	_jsii_.InvokeVoid(
		o,
		"resetInitialPosition",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) ResetMultilineStartPattern() {
	_jsii_.InvokeVoid(
		o,
		"resetMultilineStartPattern",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationLogStreamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

