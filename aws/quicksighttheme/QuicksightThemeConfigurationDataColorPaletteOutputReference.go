package quicksighttheme

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/quicksighttheme/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightThemeConfigurationDataColorPaletteOutputReference interface {
	cdktf.ComplexObject
	Colors() *[]*string
	SetColors(val *[]*string)
	ColorsInput() *[]*string
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
	EmptyFillColor() *string
	SetEmptyFillColor(val *string)
	EmptyFillColorInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *QuicksightThemeConfigurationDataColorPalette
	SetInternalValue(val *QuicksightThemeConfigurationDataColorPalette)
	MinMaxGradient() *[]*string
	SetMinMaxGradient(val *[]*string)
	MinMaxGradientInput() *[]*string
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
	ResetColors()
	ResetEmptyFillColor()
	ResetMinMaxGradient()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightThemeConfigurationDataColorPaletteOutputReference
type jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) Colors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"colors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) ColorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"colorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) EmptyFillColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyFillColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) EmptyFillColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyFillColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) InternalValue() *QuicksightThemeConfigurationDataColorPalette {
	var returns *QuicksightThemeConfigurationDataColorPalette
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) MinMaxGradient() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"minMaxGradient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) MinMaxGradientInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"minMaxGradientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightThemeConfigurationDataColorPaletteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QuicksightThemeConfigurationDataColorPaletteOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightThemeConfigurationDataColorPaletteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightTheme.QuicksightThemeConfigurationDataColorPaletteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightThemeConfigurationDataColorPaletteOutputReference_Override(q QuicksightThemeConfigurationDataColorPaletteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightTheme.QuicksightThemeConfigurationDataColorPaletteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference)SetColors(val *[]*string) {
	if err := j.validateSetColorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colors",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference)SetEmptyFillColor(val *string) {
	if err := j.validateSetEmptyFillColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emptyFillColor",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference)SetInternalValue(val *QuicksightThemeConfigurationDataColorPalette) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference)SetMinMaxGradient(val *[]*string) {
	if err := j.validateSetMinMaxGradientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minMaxGradient",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) ResetColors() {
	_jsii_.InvokeVoid(
		q,
		"resetColors",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) ResetEmptyFillColor() {
	_jsii_.InvokeVoid(
		q,
		"resetEmptyFillColor",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) ResetMinMaxGradient() {
	_jsii_.InvokeVoid(
		q,
		"resetMinMaxGradient",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationDataColorPaletteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

