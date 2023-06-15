package elastictranscoderpreset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v15/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v15/elastictranscoderpreset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastictranscoderPresetVideoOutputReference interface {
	cdktf.ComplexObject
	AspectRatio() *string
	SetAspectRatio(val *string)
	AspectRatioInput() *string
	BitRate() *string
	SetBitRate(val *string)
	BitRateInput() *string
	Codec() *string
	SetCodec(val *string)
	CodecInput() *string
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
	DisplayAspectRatio() *string
	SetDisplayAspectRatio(val *string)
	DisplayAspectRatioInput() *string
	FixedGop() *string
	SetFixedGop(val *string)
	FixedGopInput() *string
	// Experimental.
	Fqn() *string
	FrameRate() *string
	SetFrameRate(val *string)
	FrameRateInput() *string
	InternalValue() *ElastictranscoderPresetVideo
	SetInternalValue(val *ElastictranscoderPresetVideo)
	KeyframesMaxDist() *string
	SetKeyframesMaxDist(val *string)
	KeyframesMaxDistInput() *string
	MaxFrameRate() *string
	SetMaxFrameRate(val *string)
	MaxFrameRateInput() *string
	MaxHeight() *string
	SetMaxHeight(val *string)
	MaxHeightInput() *string
	MaxWidth() *string
	SetMaxWidth(val *string)
	MaxWidthInput() *string
	PaddingPolicy() *string
	SetPaddingPolicy(val *string)
	PaddingPolicyInput() *string
	Resolution() *string
	SetResolution(val *string)
	ResolutionInput() *string
	SizingPolicy() *string
	SetSizingPolicy(val *string)
	SizingPolicyInput() *string
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
	ResetAspectRatio()
	ResetBitRate()
	ResetCodec()
	ResetDisplayAspectRatio()
	ResetFixedGop()
	ResetFrameRate()
	ResetKeyframesMaxDist()
	ResetMaxFrameRate()
	ResetMaxHeight()
	ResetMaxWidth()
	ResetPaddingPolicy()
	ResetResolution()
	ResetSizingPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastictranscoderPresetVideoOutputReference
type jsiiProxy_ElastictranscoderPresetVideoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) AspectRatio() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aspectRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) AspectRatioInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aspectRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) BitRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) BitRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) Codec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) CodecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) DisplayAspectRatio() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayAspectRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) DisplayAspectRatioInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayAspectRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) FixedGop() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedGop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) FixedGopInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedGopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) FrameRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) FrameRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) InternalValue() *ElastictranscoderPresetVideo {
	var returns *ElastictranscoderPresetVideo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) KeyframesMaxDist() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyframesMaxDist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) KeyframesMaxDistInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyframesMaxDistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) MaxFrameRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxFrameRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) MaxFrameRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxFrameRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) MaxHeight() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) MaxHeightInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) MaxWidth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) MaxWidthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) PaddingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paddingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) PaddingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paddingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) Resolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SizingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SizingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastictranscoderPresetVideoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastictranscoderPresetVideoOutputReference {
	_init_.Initialize()

	if err := validateNewElastictranscoderPresetVideoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastictranscoderPresetVideoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.elastictranscoderPreset.ElastictranscoderPresetVideoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastictranscoderPresetVideoOutputReference_Override(e ElastictranscoderPresetVideoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.elastictranscoderPreset.ElastictranscoderPresetVideoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetAspectRatio(val *string) {
	if err := j.validateSetAspectRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aspectRatio",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetBitRate(val *string) {
	if err := j.validateSetBitRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitRate",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetCodec(val *string) {
	if err := j.validateSetCodecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codec",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetDisplayAspectRatio(val *string) {
	if err := j.validateSetDisplayAspectRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayAspectRatio",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetFixedGop(val *string) {
	if err := j.validateSetFixedGopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedGop",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetFrameRate(val *string) {
	if err := j.validateSetFrameRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frameRate",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetInternalValue(val *ElastictranscoderPresetVideo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetKeyframesMaxDist(val *string) {
	if err := j.validateSetKeyframesMaxDistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyframesMaxDist",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetMaxFrameRate(val *string) {
	if err := j.validateSetMaxFrameRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFrameRate",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetMaxHeight(val *string) {
	if err := j.validateSetMaxHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxHeight",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetMaxWidth(val *string) {
	if err := j.validateSetMaxWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWidth",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetPaddingPolicy(val *string) {
	if err := j.validateSetPaddingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paddingPolicy",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetResolution(val *string) {
	if err := j.validateSetResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolution",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetSizingPolicy(val *string) {
	if err := j.validateSetSizingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizingPolicy",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetAspectRatio() {
	_jsii_.InvokeVoid(
		e,
		"resetAspectRatio",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetBitRate() {
	_jsii_.InvokeVoid(
		e,
		"resetBitRate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetCodec() {
	_jsii_.InvokeVoid(
		e,
		"resetCodec",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetDisplayAspectRatio() {
	_jsii_.InvokeVoid(
		e,
		"resetDisplayAspectRatio",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetFixedGop() {
	_jsii_.InvokeVoid(
		e,
		"resetFixedGop",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetFrameRate() {
	_jsii_.InvokeVoid(
		e,
		"resetFrameRate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetKeyframesMaxDist() {
	_jsii_.InvokeVoid(
		e,
		"resetKeyframesMaxDist",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetMaxFrameRate() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxFrameRate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetMaxHeight() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxHeight",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetMaxWidth() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxWidth",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetPaddingPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetPaddingPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetResolution() {
	_jsii_.InvokeVoid(
		e,
		"resetResolution",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetSizingPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetSizingPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

