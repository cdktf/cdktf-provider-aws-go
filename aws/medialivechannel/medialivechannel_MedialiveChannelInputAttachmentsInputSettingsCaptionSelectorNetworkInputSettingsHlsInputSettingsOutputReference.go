package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference interface {
	cdktf.ComplexObject
	Bandwidth() *float64
	SetBandwidth(val *float64)
	BandwidthInput() *float64
	BufferSegments() *float64
	SetBufferSegments(val *float64)
	BufferSegmentsInput() *float64
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
	InternalValue() *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettings
	SetInternalValue(val *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettings)
	Retries() *float64
	SetRetries(val *float64)
	RetriesInput() *float64
	RetryInterval() *float64
	SetRetryInterval(val *float64)
	RetryIntervalInput() *float64
	Scte35Source() *string
	SetScte35Source(val *string)
	Scte35SourceInput() *string
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
	ResetBandwidth()
	ResetBufferSegments()
	ResetRetries()
	ResetRetryInterval()
	ResetScte35Source()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference
type jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) Bandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) BandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) BufferSegments() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferSegments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) BufferSegmentsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferSegmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) InternalValue() *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettings {
	var returns *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) Retries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) RetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) RetryInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) RetryIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) Scte35Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte35Source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) Scte35SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte35SourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference_Override(m MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference)SetBandwidth(val *float64) {
	if err := j.validateSetBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidth",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference)SetBufferSegments(val *float64) {
	if err := j.validateSetBufferSegmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferSegments",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference)SetInternalValue(val *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference)SetRetries(val *float64) {
	if err := j.validateSetRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retries",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference)SetRetryInterval(val *float64) {
	if err := j.validateSetRetryIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryInterval",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference)SetScte35Source(val *string) {
	if err := j.validateSetScte35SourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scte35Source",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) ResetBandwidth() {
	_jsii_.InvokeVoid(
		m,
		"resetBandwidth",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) ResetBufferSegments() {
	_jsii_.InvokeVoid(
		m,
		"resetBufferSegments",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) ResetRetries() {
	_jsii_.InvokeVoid(
		m,
		"resetRetries",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) ResetRetryInterval() {
	_jsii_.InvokeVoid(
		m,
		"resetRetryInterval",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) ResetScte35Source() {
	_jsii_.InvokeVoid(
		m,
		"resetScte35Source",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

