package quicksightdashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/quicksightdashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightDashboardDashboardPublishOptionsOutputReference interface {
	cdktf.ComplexObject
	AdHocFilteringOption() QuicksightDashboardDashboardPublishOptionsAdHocFilteringOptionOutputReference
	AdHocFilteringOptionInput() *QuicksightDashboardDashboardPublishOptionsAdHocFilteringOption
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
	DataPointDrillUpDownOption() QuicksightDashboardDashboardPublishOptionsDataPointDrillUpDownOptionOutputReference
	DataPointDrillUpDownOptionInput() *QuicksightDashboardDashboardPublishOptionsDataPointDrillUpDownOption
	DataPointMenuLabelOption() QuicksightDashboardDashboardPublishOptionsDataPointMenuLabelOptionOutputReference
	DataPointMenuLabelOptionInput() *QuicksightDashboardDashboardPublishOptionsDataPointMenuLabelOption
	DataPointTooltipOption() QuicksightDashboardDashboardPublishOptionsDataPointTooltipOptionOutputReference
	DataPointTooltipOptionInput() *QuicksightDashboardDashboardPublishOptionsDataPointTooltipOption
	ExportToCsvOption() QuicksightDashboardDashboardPublishOptionsExportToCsvOptionOutputReference
	ExportToCsvOptionInput() *QuicksightDashboardDashboardPublishOptionsExportToCsvOption
	ExportWithHiddenFieldsOption() QuicksightDashboardDashboardPublishOptionsExportWithHiddenFieldsOptionOutputReference
	ExportWithHiddenFieldsOptionInput() *QuicksightDashboardDashboardPublishOptionsExportWithHiddenFieldsOption
	// Experimental.
	Fqn() *string
	InternalValue() *QuicksightDashboardDashboardPublishOptions
	SetInternalValue(val *QuicksightDashboardDashboardPublishOptions)
	SheetControlsOption() QuicksightDashboardDashboardPublishOptionsSheetControlsOptionOutputReference
	SheetControlsOptionInput() *QuicksightDashboardDashboardPublishOptionsSheetControlsOption
	SheetLayoutElementMaximizationOption() QuicksightDashboardDashboardPublishOptionsSheetLayoutElementMaximizationOptionOutputReference
	SheetLayoutElementMaximizationOptionInput() *QuicksightDashboardDashboardPublishOptionsSheetLayoutElementMaximizationOption
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VisualAxisSortOption() QuicksightDashboardDashboardPublishOptionsVisualAxisSortOptionOutputReference
	VisualAxisSortOptionInput() *QuicksightDashboardDashboardPublishOptionsVisualAxisSortOption
	VisualMenuOption() QuicksightDashboardDashboardPublishOptionsVisualMenuOptionOutputReference
	VisualMenuOptionInput() *QuicksightDashboardDashboardPublishOptionsVisualMenuOption
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
	PutAdHocFilteringOption(value *QuicksightDashboardDashboardPublishOptionsAdHocFilteringOption)
	PutDataPointDrillUpDownOption(value *QuicksightDashboardDashboardPublishOptionsDataPointDrillUpDownOption)
	PutDataPointMenuLabelOption(value *QuicksightDashboardDashboardPublishOptionsDataPointMenuLabelOption)
	PutDataPointTooltipOption(value *QuicksightDashboardDashboardPublishOptionsDataPointTooltipOption)
	PutExportToCsvOption(value *QuicksightDashboardDashboardPublishOptionsExportToCsvOption)
	PutExportWithHiddenFieldsOption(value *QuicksightDashboardDashboardPublishOptionsExportWithHiddenFieldsOption)
	PutSheetControlsOption(value *QuicksightDashboardDashboardPublishOptionsSheetControlsOption)
	PutSheetLayoutElementMaximizationOption(value *QuicksightDashboardDashboardPublishOptionsSheetLayoutElementMaximizationOption)
	PutVisualAxisSortOption(value *QuicksightDashboardDashboardPublishOptionsVisualAxisSortOption)
	PutVisualMenuOption(value *QuicksightDashboardDashboardPublishOptionsVisualMenuOption)
	ResetAdHocFilteringOption()
	ResetDataPointDrillUpDownOption()
	ResetDataPointMenuLabelOption()
	ResetDataPointTooltipOption()
	ResetExportToCsvOption()
	ResetExportWithHiddenFieldsOption()
	ResetSheetControlsOption()
	ResetSheetLayoutElementMaximizationOption()
	ResetVisualAxisSortOption()
	ResetVisualMenuOption()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDashboardDashboardPublishOptionsOutputReference
type jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) AdHocFilteringOption() QuicksightDashboardDashboardPublishOptionsAdHocFilteringOptionOutputReference {
	var returns QuicksightDashboardDashboardPublishOptionsAdHocFilteringOptionOutputReference
	_jsii_.Get(
		j,
		"adHocFilteringOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) AdHocFilteringOptionInput() *QuicksightDashboardDashboardPublishOptionsAdHocFilteringOption {
	var returns *QuicksightDashboardDashboardPublishOptionsAdHocFilteringOption
	_jsii_.Get(
		j,
		"adHocFilteringOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) DataPointDrillUpDownOption() QuicksightDashboardDashboardPublishOptionsDataPointDrillUpDownOptionOutputReference {
	var returns QuicksightDashboardDashboardPublishOptionsDataPointDrillUpDownOptionOutputReference
	_jsii_.Get(
		j,
		"dataPointDrillUpDownOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) DataPointDrillUpDownOptionInput() *QuicksightDashboardDashboardPublishOptionsDataPointDrillUpDownOption {
	var returns *QuicksightDashboardDashboardPublishOptionsDataPointDrillUpDownOption
	_jsii_.Get(
		j,
		"dataPointDrillUpDownOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) DataPointMenuLabelOption() QuicksightDashboardDashboardPublishOptionsDataPointMenuLabelOptionOutputReference {
	var returns QuicksightDashboardDashboardPublishOptionsDataPointMenuLabelOptionOutputReference
	_jsii_.Get(
		j,
		"dataPointMenuLabelOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) DataPointMenuLabelOptionInput() *QuicksightDashboardDashboardPublishOptionsDataPointMenuLabelOption {
	var returns *QuicksightDashboardDashboardPublishOptionsDataPointMenuLabelOption
	_jsii_.Get(
		j,
		"dataPointMenuLabelOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) DataPointTooltipOption() QuicksightDashboardDashboardPublishOptionsDataPointTooltipOptionOutputReference {
	var returns QuicksightDashboardDashboardPublishOptionsDataPointTooltipOptionOutputReference
	_jsii_.Get(
		j,
		"dataPointTooltipOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) DataPointTooltipOptionInput() *QuicksightDashboardDashboardPublishOptionsDataPointTooltipOption {
	var returns *QuicksightDashboardDashboardPublishOptionsDataPointTooltipOption
	_jsii_.Get(
		j,
		"dataPointTooltipOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ExportToCsvOption() QuicksightDashboardDashboardPublishOptionsExportToCsvOptionOutputReference {
	var returns QuicksightDashboardDashboardPublishOptionsExportToCsvOptionOutputReference
	_jsii_.Get(
		j,
		"exportToCsvOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ExportToCsvOptionInput() *QuicksightDashboardDashboardPublishOptionsExportToCsvOption {
	var returns *QuicksightDashboardDashboardPublishOptionsExportToCsvOption
	_jsii_.Get(
		j,
		"exportToCsvOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ExportWithHiddenFieldsOption() QuicksightDashboardDashboardPublishOptionsExportWithHiddenFieldsOptionOutputReference {
	var returns QuicksightDashboardDashboardPublishOptionsExportWithHiddenFieldsOptionOutputReference
	_jsii_.Get(
		j,
		"exportWithHiddenFieldsOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ExportWithHiddenFieldsOptionInput() *QuicksightDashboardDashboardPublishOptionsExportWithHiddenFieldsOption {
	var returns *QuicksightDashboardDashboardPublishOptionsExportWithHiddenFieldsOption
	_jsii_.Get(
		j,
		"exportWithHiddenFieldsOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) InternalValue() *QuicksightDashboardDashboardPublishOptions {
	var returns *QuicksightDashboardDashboardPublishOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) SheetControlsOption() QuicksightDashboardDashboardPublishOptionsSheetControlsOptionOutputReference {
	var returns QuicksightDashboardDashboardPublishOptionsSheetControlsOptionOutputReference
	_jsii_.Get(
		j,
		"sheetControlsOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) SheetControlsOptionInput() *QuicksightDashboardDashboardPublishOptionsSheetControlsOption {
	var returns *QuicksightDashboardDashboardPublishOptionsSheetControlsOption
	_jsii_.Get(
		j,
		"sheetControlsOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) SheetLayoutElementMaximizationOption() QuicksightDashboardDashboardPublishOptionsSheetLayoutElementMaximizationOptionOutputReference {
	var returns QuicksightDashboardDashboardPublishOptionsSheetLayoutElementMaximizationOptionOutputReference
	_jsii_.Get(
		j,
		"sheetLayoutElementMaximizationOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) SheetLayoutElementMaximizationOptionInput() *QuicksightDashboardDashboardPublishOptionsSheetLayoutElementMaximizationOption {
	var returns *QuicksightDashboardDashboardPublishOptionsSheetLayoutElementMaximizationOption
	_jsii_.Get(
		j,
		"sheetLayoutElementMaximizationOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) VisualAxisSortOption() QuicksightDashboardDashboardPublishOptionsVisualAxisSortOptionOutputReference {
	var returns QuicksightDashboardDashboardPublishOptionsVisualAxisSortOptionOutputReference
	_jsii_.Get(
		j,
		"visualAxisSortOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) VisualAxisSortOptionInput() *QuicksightDashboardDashboardPublishOptionsVisualAxisSortOption {
	var returns *QuicksightDashboardDashboardPublishOptionsVisualAxisSortOption
	_jsii_.Get(
		j,
		"visualAxisSortOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) VisualMenuOption() QuicksightDashboardDashboardPublishOptionsVisualMenuOptionOutputReference {
	var returns QuicksightDashboardDashboardPublishOptionsVisualMenuOptionOutputReference
	_jsii_.Get(
		j,
		"visualMenuOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) VisualMenuOptionInput() *QuicksightDashboardDashboardPublishOptionsVisualMenuOption {
	var returns *QuicksightDashboardDashboardPublishOptionsVisualMenuOption
	_jsii_.Get(
		j,
		"visualMenuOptionInput",
		&returns,
	)
	return returns
}


func NewQuicksightDashboardDashboardPublishOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QuicksightDashboardDashboardPublishOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDashboardDashboardPublishOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDashboard.QuicksightDashboardDashboardPublishOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDashboardDashboardPublishOptionsOutputReference_Override(q QuicksightDashboardDashboardPublishOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDashboard.QuicksightDashboardDashboardPublishOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference)SetInternalValue(val *QuicksightDashboardDashboardPublishOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) PutAdHocFilteringOption(value *QuicksightDashboardDashboardPublishOptionsAdHocFilteringOption) {
	if err := q.validatePutAdHocFilteringOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAdHocFilteringOption",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) PutDataPointDrillUpDownOption(value *QuicksightDashboardDashboardPublishOptionsDataPointDrillUpDownOption) {
	if err := q.validatePutDataPointDrillUpDownOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDataPointDrillUpDownOption",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) PutDataPointMenuLabelOption(value *QuicksightDashboardDashboardPublishOptionsDataPointMenuLabelOption) {
	if err := q.validatePutDataPointMenuLabelOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDataPointMenuLabelOption",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) PutDataPointTooltipOption(value *QuicksightDashboardDashboardPublishOptionsDataPointTooltipOption) {
	if err := q.validatePutDataPointTooltipOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDataPointTooltipOption",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) PutExportToCsvOption(value *QuicksightDashboardDashboardPublishOptionsExportToCsvOption) {
	if err := q.validatePutExportToCsvOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putExportToCsvOption",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) PutExportWithHiddenFieldsOption(value *QuicksightDashboardDashboardPublishOptionsExportWithHiddenFieldsOption) {
	if err := q.validatePutExportWithHiddenFieldsOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putExportWithHiddenFieldsOption",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) PutSheetControlsOption(value *QuicksightDashboardDashboardPublishOptionsSheetControlsOption) {
	if err := q.validatePutSheetControlsOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSheetControlsOption",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) PutSheetLayoutElementMaximizationOption(value *QuicksightDashboardDashboardPublishOptionsSheetLayoutElementMaximizationOption) {
	if err := q.validatePutSheetLayoutElementMaximizationOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSheetLayoutElementMaximizationOption",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) PutVisualAxisSortOption(value *QuicksightDashboardDashboardPublishOptionsVisualAxisSortOption) {
	if err := q.validatePutVisualAxisSortOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putVisualAxisSortOption",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) PutVisualMenuOption(value *QuicksightDashboardDashboardPublishOptionsVisualMenuOption) {
	if err := q.validatePutVisualMenuOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putVisualMenuOption",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ResetAdHocFilteringOption() {
	_jsii_.InvokeVoid(
		q,
		"resetAdHocFilteringOption",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ResetDataPointDrillUpDownOption() {
	_jsii_.InvokeVoid(
		q,
		"resetDataPointDrillUpDownOption",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ResetDataPointMenuLabelOption() {
	_jsii_.InvokeVoid(
		q,
		"resetDataPointMenuLabelOption",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ResetDataPointTooltipOption() {
	_jsii_.InvokeVoid(
		q,
		"resetDataPointTooltipOption",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ResetExportToCsvOption() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToCsvOption",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ResetExportWithHiddenFieldsOption() {
	_jsii_.InvokeVoid(
		q,
		"resetExportWithHiddenFieldsOption",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ResetSheetControlsOption() {
	_jsii_.InvokeVoid(
		q,
		"resetSheetControlsOption",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ResetSheetLayoutElementMaximizationOption() {
	_jsii_.InvokeVoid(
		q,
		"resetSheetLayoutElementMaximizationOption",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ResetVisualAxisSortOption() {
	_jsii_.InvokeVoid(
		q,
		"resetVisualAxisSortOption",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ResetVisualMenuOption() {
	_jsii_.InvokeVoid(
		q,
		"resetVisualMenuOption",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

