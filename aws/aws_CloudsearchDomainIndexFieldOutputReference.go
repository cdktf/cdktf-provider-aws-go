// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudsearchDomainIndexFieldOutputReference interface {
	cdktf.ComplexObject
	AnalysisScheme() *string
	SetAnalysisScheme(val *string)
	AnalysisSchemeInput() *string
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
	DefaultValue() *string
	SetDefaultValue(val *string)
	DefaultValueInput() *string
	Facet() interface{}
	SetFacet(val interface{})
	FacetInput() interface{}
	// Experimental.
	Fqn() *string
	Highlight() interface{}
	SetHighlight(val interface{})
	HighlightInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Return() interface{}
	SetReturn(val interface{})
	ReturnInput() interface{}
	Search() interface{}
	SetSearch(val interface{})
	SearchInput() interface{}
	Sort() interface{}
	SetSort(val interface{})
	SortInput() interface{}
	SourceFields() *string
	SetSourceFields(val *string)
	SourceFieldsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetAnalysisScheme()
	ResetDefaultValue()
	ResetFacet()
	ResetHighlight()
	ResetReturn()
	ResetSearch()
	ResetSort()
	ResetSourceFields()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudsearchDomainIndexFieldOutputReference
type jsiiProxy_CloudsearchDomainIndexFieldOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) AnalysisScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) AnalysisSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) DefaultValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) DefaultValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) Facet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) FacetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) Highlight() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"highlight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) HighlightInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"highlightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) Return() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"return",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) ReturnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) Search() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) SearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) Sort() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) SortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) SourceFields() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) SourceFieldsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCloudsearchDomainIndexFieldOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudsearchDomainIndexFieldOutputReference {
	_init_.Initialize()

	if err := validateNewCloudsearchDomainIndexFieldOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudsearchDomainIndexFieldOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.CloudsearchDomainIndexFieldOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudsearchDomainIndexFieldOutputReference_Override(c CloudsearchDomainIndexFieldOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.CloudsearchDomainIndexFieldOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetAnalysisScheme(val *string) {
	if err := j.validateSetAnalysisSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analysisScheme",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetDefaultValue(val *string) {
	if err := j.validateSetDefaultValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultValue",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetFacet(val interface{}) {
	if err := j.validateSetFacetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facet",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetHighlight(val interface{}) {
	if err := j.validateSetHighlightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"highlight",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetReturn(val interface{}) {
	if err := j.validateSetReturnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"return",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetSearch(val interface{}) {
	if err := j.validateSetSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"search",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetSort(val interface{}) {
	if err := j.validateSetSortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sort",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetSourceFields(val *string) {
	if err := j.validateSetSourceFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFields",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainIndexFieldOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) ResetAnalysisScheme() {
	_jsii_.InvokeVoid(
		c,
		"resetAnalysisScheme",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) ResetDefaultValue() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) ResetFacet() {
	_jsii_.InvokeVoid(
		c,
		"resetFacet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) ResetHighlight() {
	_jsii_.InvokeVoid(
		c,
		"resetHighlight",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) ResetReturn() {
	_jsii_.InvokeVoid(
		c,
		"resetReturn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) ResetSearch() {
	_jsii_.InvokeVoid(
		c,
		"resetSearch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		c,
		"resetSort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) ResetSourceFields() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceFields",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainIndexFieldOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

