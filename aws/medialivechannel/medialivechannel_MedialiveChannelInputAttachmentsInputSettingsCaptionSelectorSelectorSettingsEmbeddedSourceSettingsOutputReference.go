package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference interface {
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
	Convert608To708() *string
	SetConvert608To708(val *string)
	Convert608To708Input() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettings
	SetInternalValue(val *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettings)
	Scte20Detection() *string
	SetScte20Detection(val *string)
	Scte20DetectionInput() *string
	Source608ChannelNumber() *float64
	SetSource608ChannelNumber(val *float64)
	Source608ChannelNumberInput() *float64
	Source608TrackNumber() *float64
	SetSource608TrackNumber(val *float64)
	Source608TrackNumberInput() *float64
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
	ResetConvert608To708()
	ResetScte20Detection()
	ResetSource608ChannelNumber()
	ResetSource608TrackNumber()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference
type jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) Convert608To708() *string {
	var returns *string
	_jsii_.Get(
		j,
		"convert608To708",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) Convert608To708Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"convert608To708Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) InternalValue() *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettings {
	var returns *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) Scte20Detection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte20Detection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) Scte20DetectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte20DetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) Source608ChannelNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"source608ChannelNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) Source608ChannelNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"source608ChannelNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) Source608TrackNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"source608TrackNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) Source608TrackNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"source608TrackNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference_Override(m MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference)SetConvert608To708(val *string) {
	if err := j.validateSetConvert608To708Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"convert608To708",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference)SetInternalValue(val *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference)SetScte20Detection(val *string) {
	if err := j.validateSetScte20DetectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scte20Detection",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference)SetSource608ChannelNumber(val *float64) {
	if err := j.validateSetSource608ChannelNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source608ChannelNumber",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference)SetSource608TrackNumber(val *float64) {
	if err := j.validateSetSource608TrackNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source608TrackNumber",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) ResetConvert608To708() {
	_jsii_.InvokeVoid(
		m,
		"resetConvert608To708",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) ResetScte20Detection() {
	_jsii_.InvokeVoid(
		m,
		"resetScte20Detection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) ResetSource608ChannelNumber() {
	_jsii_.InvokeVoid(
		m,
		"resetSource608ChannelNumber",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) ResetSource608TrackNumber() {
	_jsii_.InvokeVoid(
		m,
		"resetSource608TrackNumber",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

