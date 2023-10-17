// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/quicksighttheme/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightThemeConfigurationOutputReference interface {
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
	DataColorPalette() QuicksightThemeConfigurationDataColorPaletteOutputReference
	DataColorPaletteInput() *QuicksightThemeConfigurationDataColorPalette
	// Experimental.
	Fqn() *string
	InternalValue() *QuicksightThemeConfiguration
	SetInternalValue(val *QuicksightThemeConfiguration)
	Sheet() QuicksightThemeConfigurationSheetOutputReference
	SheetInput() *QuicksightThemeConfigurationSheet
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Typography() QuicksightThemeConfigurationTypographyOutputReference
	TypographyInput() *QuicksightThemeConfigurationTypography
	UiColorPalette() QuicksightThemeConfigurationUiColorPaletteOutputReference
	UiColorPaletteInput() *QuicksightThemeConfigurationUiColorPalette
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
	PutDataColorPalette(value *QuicksightThemeConfigurationDataColorPalette)
	PutSheet(value *QuicksightThemeConfigurationSheet)
	PutTypography(value *QuicksightThemeConfigurationTypography)
	PutUiColorPalette(value *QuicksightThemeConfigurationUiColorPalette)
	ResetDataColorPalette()
	ResetSheet()
	ResetTypography()
	ResetUiColorPalette()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightThemeConfigurationOutputReference
type jsiiProxy_QuicksightThemeConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) DataColorPalette() QuicksightThemeConfigurationDataColorPaletteOutputReference {
	var returns QuicksightThemeConfigurationDataColorPaletteOutputReference
	_jsii_.Get(
		j,
		"dataColorPalette",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) DataColorPaletteInput() *QuicksightThemeConfigurationDataColorPalette {
	var returns *QuicksightThemeConfigurationDataColorPalette
	_jsii_.Get(
		j,
		"dataColorPaletteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) InternalValue() *QuicksightThemeConfiguration {
	var returns *QuicksightThemeConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) Sheet() QuicksightThemeConfigurationSheetOutputReference {
	var returns QuicksightThemeConfigurationSheetOutputReference
	_jsii_.Get(
		j,
		"sheet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) SheetInput() *QuicksightThemeConfigurationSheet {
	var returns *QuicksightThemeConfigurationSheet
	_jsii_.Get(
		j,
		"sheetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) Typography() QuicksightThemeConfigurationTypographyOutputReference {
	var returns QuicksightThemeConfigurationTypographyOutputReference
	_jsii_.Get(
		j,
		"typography",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) TypographyInput() *QuicksightThemeConfigurationTypography {
	var returns *QuicksightThemeConfigurationTypography
	_jsii_.Get(
		j,
		"typographyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) UiColorPalette() QuicksightThemeConfigurationUiColorPaletteOutputReference {
	var returns QuicksightThemeConfigurationUiColorPaletteOutputReference
	_jsii_.Get(
		j,
		"uiColorPalette",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference) UiColorPaletteInput() *QuicksightThemeConfigurationUiColorPalette {
	var returns *QuicksightThemeConfigurationUiColorPalette
	_jsii_.Get(
		j,
		"uiColorPaletteInput",
		&returns,
	)
	return returns
}


func NewQuicksightThemeConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QuicksightThemeConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightThemeConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightThemeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightTheme.QuicksightThemeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightThemeConfigurationOutputReference_Override(q QuicksightThemeConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightTheme.QuicksightThemeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference)SetInternalValue(val *QuicksightThemeConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) PutDataColorPalette(value *QuicksightThemeConfigurationDataColorPalette) {
	if err := q.validatePutDataColorPaletteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDataColorPalette",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) PutSheet(value *QuicksightThemeConfigurationSheet) {
	if err := q.validatePutSheetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSheet",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) PutTypography(value *QuicksightThemeConfigurationTypography) {
	if err := q.validatePutTypographyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTypography",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) PutUiColorPalette(value *QuicksightThemeConfigurationUiColorPalette) {
	if err := q.validatePutUiColorPaletteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putUiColorPalette",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) ResetDataColorPalette() {
	_jsii_.InvokeVoid(
		q,
		"resetDataColorPalette",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) ResetSheet() {
	_jsii_.InvokeVoid(
		q,
		"resetSheet",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) ResetTypography() {
	_jsii_.InvokeVoid(
		q,
		"resetTypography",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) ResetUiColorPalette() {
	_jsii_.InvokeVoid(
		q,
		"resetUiColorPalette",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightThemeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

