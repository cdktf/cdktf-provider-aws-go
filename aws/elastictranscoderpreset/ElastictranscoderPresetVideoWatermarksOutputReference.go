package elastictranscoderpreset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/elastictranscoderpreset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastictranscoderPresetVideoWatermarksOutputReference interface {
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
	HorizontalAlign() *string
	SetHorizontalAlign(val *string)
	HorizontalAlignInput() *string
	HorizontalOffset() *string
	SetHorizontalOffset(val *string)
	HorizontalOffsetInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxHeight() *string
	SetMaxHeight(val *string)
	MaxHeightInput() *string
	MaxWidth() *string
	SetMaxWidth(val *string)
	MaxWidthInput() *string
	Opacity() *string
	SetOpacity(val *string)
	OpacityInput() *string
	SizingPolicy() *string
	SetSizingPolicy(val *string)
	SizingPolicyInput() *string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VerticalAlign() *string
	SetVerticalAlign(val *string)
	VerticalAlignInput() *string
	VerticalOffset() *string
	SetVerticalOffset(val *string)
	VerticalOffsetInput() *string
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
	ResetHorizontalAlign()
	ResetHorizontalOffset()
	ResetId()
	ResetMaxHeight()
	ResetMaxWidth()
	ResetOpacity()
	ResetSizingPolicy()
	ResetTarget()
	ResetVerticalAlign()
	ResetVerticalOffset()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastictranscoderPresetVideoWatermarksOutputReference
type jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) HorizontalAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"horizontalAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) HorizontalAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"horizontalAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) HorizontalOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"horizontalOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) HorizontalOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"horizontalOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) MaxHeight() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) MaxHeightInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) MaxWidth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) MaxWidthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) Opacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) OpacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) SizingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) SizingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) VerticalAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) VerticalAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) VerticalOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) VerticalOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalOffsetInput",
		&returns,
	)
	return returns
}


func NewElastictranscoderPresetVideoWatermarksOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ElastictranscoderPresetVideoWatermarksOutputReference {
	_init_.Initialize()

	if err := validateNewElastictranscoderPresetVideoWatermarksOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.elastictranscoderPreset.ElastictranscoderPresetVideoWatermarksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElastictranscoderPresetVideoWatermarksOutputReference_Override(e ElastictranscoderPresetVideoWatermarksOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.elastictranscoderPreset.ElastictranscoderPresetVideoWatermarksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetHorizontalAlign(val *string) {
	if err := j.validateSetHorizontalAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"horizontalAlign",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetHorizontalOffset(val *string) {
	if err := j.validateSetHorizontalOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"horizontalOffset",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetMaxHeight(val *string) {
	if err := j.validateSetMaxHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxHeight",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetMaxWidth(val *string) {
	if err := j.validateSetMaxWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWidth",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetOpacity(val *string) {
	if err := j.validateSetOpacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opacity",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetSizingPolicy(val *string) {
	if err := j.validateSetSizingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizingPolicy",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetVerticalAlign(val *string) {
	if err := j.validateSetVerticalAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verticalAlign",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference)SetVerticalOffset(val *string) {
	if err := j.validateSetVerticalOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verticalOffset",
		val,
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ResetHorizontalAlign() {
	_jsii_.InvokeVoid(
		e,
		"resetHorizontalAlign",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ResetHorizontalOffset() {
	_jsii_.InvokeVoid(
		e,
		"resetHorizontalOffset",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ResetMaxHeight() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxHeight",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ResetMaxWidth() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxWidth",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ResetOpacity() {
	_jsii_.InvokeVoid(
		e,
		"resetOpacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ResetSizingPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetSizingPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		e,
		"resetTarget",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ResetVerticalAlign() {
	_jsii_.InvokeVoid(
		e,
		"resetVerticalAlign",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ResetVerticalOffset() {
	_jsii_.InvokeVoid(
		e,
		"resetVerticalOffset",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastictranscoderPresetVideoWatermarksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

