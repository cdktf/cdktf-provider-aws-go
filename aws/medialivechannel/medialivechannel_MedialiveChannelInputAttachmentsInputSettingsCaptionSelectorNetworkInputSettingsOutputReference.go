package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference interface {
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
	HlsInputSettings() MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference
	HlsInputSettingsInput() *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettings
	InternalValue() *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettings
	SetInternalValue(val *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettings)
	ServerValidation() *string
	SetServerValidation(val *string)
	ServerValidationInput() *string
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
	PutHlsInputSettings(value *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettings)
	ResetHlsInputSettings()
	ResetServerValidation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference
type jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) HlsInputSettings() MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference {
	var returns MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettingsOutputReference
	_jsii_.Get(
		j,
		"hlsInputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) HlsInputSettingsInput() *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettings {
	var returns *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettings
	_jsii_.Get(
		j,
		"hlsInputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) InternalValue() *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettings {
	var returns *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) ServerValidation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) ServerValidationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference_Override(m MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference)SetInternalValue(val *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference)SetServerValidation(val *string) {
	if err := j.validateSetServerValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverValidation",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) PutHlsInputSettings(value *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsHlsInputSettings) {
	if err := m.validatePutHlsInputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putHlsInputSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) ResetHlsInputSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetHlsInputSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) ResetServerValidation() {
	_jsii_.InvokeVoid(
		m,
		"resetServerValidation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

