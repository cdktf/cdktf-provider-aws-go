package glueclassifier

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/glueclassifier/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueClassifierCsvClassifierOutputReference interface {
	cdktf.ComplexObject
	AllowSingleColumn() interface{}
	SetAllowSingleColumn(val interface{})
	AllowSingleColumnInput() interface{}
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
	ContainsHeader() *string
	SetContainsHeader(val *string)
	ContainsHeaderInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomDatatypeConfigured() interface{}
	SetCustomDatatypeConfigured(val interface{})
	CustomDatatypeConfiguredInput() interface{}
	CustomDatatypes() *[]*string
	SetCustomDatatypes(val *[]*string)
	CustomDatatypesInput() *[]*string
	Delimiter() *string
	SetDelimiter(val *string)
	DelimiterInput() *string
	DisableValueTrimming() interface{}
	SetDisableValueTrimming(val interface{})
	DisableValueTrimmingInput() interface{}
	// Experimental.
	Fqn() *string
	Header() *[]*string
	SetHeader(val *[]*string)
	HeaderInput() *[]*string
	InternalValue() *GlueClassifierCsvClassifier
	SetInternalValue(val *GlueClassifierCsvClassifier)
	QuoteSymbol() *string
	SetQuoteSymbol(val *string)
	QuoteSymbolInput() *string
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
	ResetAllowSingleColumn()
	ResetContainsHeader()
	ResetCustomDatatypeConfigured()
	ResetCustomDatatypes()
	ResetDelimiter()
	ResetDisableValueTrimming()
	ResetHeader()
	ResetQuoteSymbol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueClassifierCsvClassifierOutputReference
type jsiiProxy_GlueClassifierCsvClassifierOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) AllowSingleColumn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSingleColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) AllowSingleColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSingleColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ContainsHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containsHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ContainsHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containsHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) CustomDatatypeConfigured() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customDatatypeConfigured",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) CustomDatatypeConfiguredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customDatatypeConfiguredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) CustomDatatypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDatatypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) CustomDatatypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDatatypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) Delimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) DelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) DisableValueTrimming() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableValueTrimming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) DisableValueTrimmingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableValueTrimmingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) Header() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) HeaderInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) InternalValue() *GlueClassifierCsvClassifier {
	var returns *GlueClassifierCsvClassifier
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) QuoteSymbol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteSymbol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) QuoteSymbolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteSymbolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueClassifierCsvClassifierOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GlueClassifierCsvClassifierOutputReference {
	_init_.Initialize()

	if err := validateNewGlueClassifierCsvClassifierOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueClassifierCsvClassifierOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.glueClassifier.GlueClassifierCsvClassifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueClassifierCsvClassifierOutputReference_Override(g GlueClassifierCsvClassifierOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.glueClassifier.GlueClassifierCsvClassifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference)SetAllowSingleColumn(val interface{}) {
	if err := j.validateSetAllowSingleColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowSingleColumn",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference)SetContainsHeader(val *string) {
	if err := j.validateSetContainsHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containsHeader",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference)SetCustomDatatypeConfigured(val interface{}) {
	if err := j.validateSetCustomDatatypeConfiguredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDatatypeConfigured",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference)SetCustomDatatypes(val *[]*string) {
	if err := j.validateSetCustomDatatypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDatatypes",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference)SetDelimiter(val *string) {
	if err := j.validateSetDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delimiter",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference)SetDisableValueTrimming(val interface{}) {
	if err := j.validateSetDisableValueTrimmingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableValueTrimming",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference)SetHeader(val *[]*string) {
	if err := j.validateSetHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"header",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference)SetInternalValue(val *GlueClassifierCsvClassifier) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference)SetQuoteSymbol(val *string) {
	if err := j.validateSetQuoteSymbolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quoteSymbol",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetAllowSingleColumn() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowSingleColumn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetContainsHeader() {
	_jsii_.InvokeVoid(
		g,
		"resetContainsHeader",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetCustomDatatypeConfigured() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomDatatypeConfigured",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetCustomDatatypes() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomDatatypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetDelimiter() {
	_jsii_.InvokeVoid(
		g,
		"resetDelimiter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetDisableValueTrimming() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableValueTrimming",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		g,
		"resetHeader",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetQuoteSymbol() {
	_jsii_.InvokeVoid(
		g,
		"resetQuoteSymbol",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

