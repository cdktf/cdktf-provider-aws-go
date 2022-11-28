package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference interface {
	cdktf.ComplexObject
	Bitrate() *float64
	SetBitrate(val *float64)
	BitrateInput() *float64
	CodingMode() *string
	SetCodingMode(val *string)
	CodingModeInput() *string
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
	Dialnorm() *float64
	SetDialnorm(val *float64)
	DialnormInput() *float64
	DrcLine() *string
	SetDrcLine(val *string)
	DrcLineInput() *string
	DrcRf() *string
	SetDrcRf(val *string)
	DrcRfInput() *string
	// Experimental.
	Fqn() *string
	HeightTrim() *float64
	SetHeightTrim(val *float64)
	HeightTrimInput() *float64
	InternalValue() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettings
	SetInternalValue(val *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettings)
	SurroundTrim() *float64
	SetSurroundTrim(val *float64)
	SurroundTrimInput() *float64
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
	ResetBitrate()
	ResetCodingMode()
	ResetDialnorm()
	ResetDrcLine()
	ResetDrcRf()
	ResetHeightTrim()
	ResetSurroundTrim()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) Bitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) BitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) CodingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) CodingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) Dialnorm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dialnorm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) DialnormInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dialnormInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) DrcLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) DrcLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) DrcRf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcRf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) DrcRfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcRfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) HeightTrim() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightTrim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) HeightTrimInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightTrimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) InternalValue() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettings {
	var returns *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) SurroundTrim() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"surroundTrim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) SurroundTrimInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"surroundTrimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference_Override(m MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference)SetBitrate(val *float64) {
	if err := j.validateSetBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrate",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference)SetCodingMode(val *string) {
	if err := j.validateSetCodingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codingMode",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference)SetDialnorm(val *float64) {
	if err := j.validateSetDialnormParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dialnorm",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference)SetDrcLine(val *string) {
	if err := j.validateSetDrcLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drcLine",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference)SetDrcRf(val *string) {
	if err := j.validateSetDrcRfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drcRf",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference)SetHeightTrim(val *float64) {
	if err := j.validateSetHeightTrimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heightTrim",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference)SetSurroundTrim(val *float64) {
	if err := j.validateSetSurroundTrimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"surroundTrim",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) ResetBitrate() {
	_jsii_.InvokeVoid(
		m,
		"resetBitrate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) ResetCodingMode() {
	_jsii_.InvokeVoid(
		m,
		"resetCodingMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) ResetDialnorm() {
	_jsii_.InvokeVoid(
		m,
		"resetDialnorm",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) ResetDrcLine() {
	_jsii_.InvokeVoid(
		m,
		"resetDrcLine",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) ResetDrcRf() {
	_jsii_.InvokeVoid(
		m,
		"resetDrcRf",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) ResetHeightTrim() {
	_jsii_.InvokeVoid(
		m,
		"resetHeightTrim",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) ResetSurroundTrim() {
	_jsii_.InvokeVoid(
		m,
		"resetSurroundTrim",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

