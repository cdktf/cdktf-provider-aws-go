package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference interface {
	cdktf.ComplexObject
	ColumnDepth() *float64
	SetColumnDepth(val *float64)
	ColumnDepthInput() *float64
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
	IncludeFec() *string
	SetIncludeFec(val *string)
	IncludeFecInput() *string
	InternalValue() *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettings
	SetInternalValue(val *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettings)
	RowLength() *float64
	SetRowLength(val *float64)
	RowLengthInput() *float64
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
	ResetColumnDepth()
	ResetIncludeFec()
	ResetRowLength()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) ColumnDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"columnDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) ColumnDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"columnDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) IncludeFec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeFec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) IncludeFecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeFecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) InternalValue() *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettings {
	var returns *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) RowLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) RowLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference_Override(m MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference)SetColumnDepth(val *float64) {
	if err := j.validateSetColumnDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnDepth",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference)SetIncludeFec(val *string) {
	if err := j.validateSetIncludeFecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeFec",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference)SetRowLength(val *float64) {
	if err := j.validateSetRowLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowLength",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) ResetColumnDepth() {
	_jsii_.InvokeVoid(
		m,
		"resetColumnDepth",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) ResetIncludeFec() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludeFec",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) ResetRowLength() {
	_jsii_.InvokeVoid(
		m,
		"resetRowLength",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

