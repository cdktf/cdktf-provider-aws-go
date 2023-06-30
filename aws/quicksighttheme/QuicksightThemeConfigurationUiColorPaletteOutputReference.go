package quicksighttheme

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/quicksighttheme/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightThemeConfigurationUiColorPaletteOutputReference interface {
	cdktf.ComplexObject
	Accent() *string
	SetAccent(val *string)
	AccentForeground() *string
	SetAccentForeground(val *string)
	AccentForegroundInput() *string
	AccentInput() *string
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
	Danger() *string
	SetDanger(val *string)
	DangerForeground() *string
	SetDangerForeground(val *string)
	DangerForegroundInput() *string
	DangerInput() *string
	Dimension() *string
	SetDimension(val *string)
	DimensionForeground() *string
	SetDimensionForeground(val *string)
	DimensionForegroundInput() *string
	DimensionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *QuicksightThemeConfigurationUiColorPalette
	SetInternalValue(val *QuicksightThemeConfigurationUiColorPalette)
	Measure() *string
	SetMeasure(val *string)
	MeasureForeground() *string
	SetMeasureForeground(val *string)
	MeasureForegroundInput() *string
	MeasureInput() *string
	PrimaryBackground() *string
	SetPrimaryBackground(val *string)
	PrimaryBackgroundInput() *string
	PrimaryForeground() *string
	SetPrimaryForeground(val *string)
	PrimaryForegroundInput() *string
	SecondaryBackground() *string
	SetSecondaryBackground(val *string)
	SecondaryBackgroundInput() *string
	SecondaryForeground() *string
	SetSecondaryForeground(val *string)
	SecondaryForegroundInput() *string
	Success() *string
	SetSuccess(val *string)
	SuccessForeground() *string
	SetSuccessForeground(val *string)
	SuccessForegroundInput() *string
	SuccessInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Warning() *string
	SetWarning(val *string)
	WarningForeground() *string
	SetWarningForeground(val *string)
	WarningForegroundInput() *string
	WarningInput() *string
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
	ResetAccent()
	ResetAccentForeground()
	ResetDanger()
	ResetDangerForeground()
	ResetDimension()
	ResetDimensionForeground()
	ResetMeasure()
	ResetMeasureForeground()
	ResetPrimaryBackground()
	ResetPrimaryForeground()
	ResetSecondaryBackground()
	ResetSecondaryForeground()
	ResetSuccess()
	ResetSuccessForeground()
	ResetWarning()
	ResetWarningForeground()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightThemeConfigurationUiColorPaletteOutputReference
type jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) Accent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) AccentForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accentForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) AccentForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accentForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) AccentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) Danger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"danger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) DangerForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dangerForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) DangerForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dangerForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) DangerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dangerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) Dimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) DimensionForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) DimensionForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) DimensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) InternalValue() *QuicksightThemeConfigurationUiColorPalette {
	var returns *QuicksightThemeConfigurationUiColorPalette
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) Measure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) MeasureForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) MeasureForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) MeasureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) PrimaryBackground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) PrimaryBackgroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBackgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) PrimaryForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) PrimaryForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) SecondaryBackground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) SecondaryBackgroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBackgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) SecondaryForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) SecondaryForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) Success() *string {
	var returns *string
	_jsii_.Get(
		j,
		"success",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) SuccessForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) SuccessForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) SuccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) Warning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) WarningForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) WarningForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) WarningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningInput",
		&returns,
	)
	return returns
}


func NewQuicksightThemeConfigurationUiColorPaletteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QuicksightThemeConfigurationUiColorPaletteOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightThemeConfigurationUiColorPaletteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightTheme.QuicksightThemeConfigurationUiColorPaletteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightThemeConfigurationUiColorPaletteOutputReference_Override(q QuicksightThemeConfigurationUiColorPaletteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightTheme.QuicksightThemeConfigurationUiColorPaletteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetAccent(val *string) {
	if err := j.validateSetAccentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accent",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetAccentForeground(val *string) {
	if err := j.validateSetAccentForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accentForeground",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetDanger(val *string) {
	if err := j.validateSetDangerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"danger",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetDangerForeground(val *string) {
	if err := j.validateSetDangerForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dangerForeground",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetDimension(val *string) {
	if err := j.validateSetDimensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimension",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetDimensionForeground(val *string) {
	if err := j.validateSetDimensionForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensionForeground",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetInternalValue(val *QuicksightThemeConfigurationUiColorPalette) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetMeasure(val *string) {
	if err := j.validateSetMeasureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measure",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetMeasureForeground(val *string) {
	if err := j.validateSetMeasureForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureForeground",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetPrimaryBackground(val *string) {
	if err := j.validateSetPrimaryBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryBackground",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetPrimaryForeground(val *string) {
	if err := j.validateSetPrimaryForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryForeground",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetSecondaryBackground(val *string) {
	if err := j.validateSetSecondaryBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryBackground",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetSecondaryForeground(val *string) {
	if err := j.validateSetSecondaryForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryForeground",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetSuccess(val *string) {
	if err := j.validateSetSuccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"success",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetSuccessForeground(val *string) {
	if err := j.validateSetSuccessForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successForeground",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetWarning(val *string) {
	if err := j.validateSetWarningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warning",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference)SetWarningForeground(val *string) {
	if err := j.validateSetWarningForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warningForeground",
		val,
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetAccent() {
	_jsii_.InvokeVoid(
		q,
		"resetAccent",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetAccentForeground() {
	_jsii_.InvokeVoid(
		q,
		"resetAccentForeground",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetDanger() {
	_jsii_.InvokeVoid(
		q,
		"resetDanger",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetDangerForeground() {
	_jsii_.InvokeVoid(
		q,
		"resetDangerForeground",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		q,
		"resetDimension",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetDimensionForeground() {
	_jsii_.InvokeVoid(
		q,
		"resetDimensionForeground",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetMeasure() {
	_jsii_.InvokeVoid(
		q,
		"resetMeasure",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetMeasureForeground() {
	_jsii_.InvokeVoid(
		q,
		"resetMeasureForeground",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetPrimaryBackground() {
	_jsii_.InvokeVoid(
		q,
		"resetPrimaryBackground",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetPrimaryForeground() {
	_jsii_.InvokeVoid(
		q,
		"resetPrimaryForeground",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetSecondaryBackground() {
	_jsii_.InvokeVoid(
		q,
		"resetSecondaryBackground",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetSecondaryForeground() {
	_jsii_.InvokeVoid(
		q,
		"resetSecondaryForeground",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetSuccess() {
	_jsii_.InvokeVoid(
		q,
		"resetSuccess",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetSuccessForeground() {
	_jsii_.InvokeVoid(
		q,
		"resetSuccessForeground",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetWarning() {
	_jsii_.InvokeVoid(
		q,
		"resetWarning",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ResetWarningForeground() {
	_jsii_.InvokeVoid(
		q,
		"resetWarningForeground",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightThemeConfigurationUiColorPaletteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

