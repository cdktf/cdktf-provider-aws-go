package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference interface {
	cdktf.ComplexObject
	All() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAllOutputReference
	AllInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAll
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
	ExcludedHeaders() *[]*string
	SetExcludedHeaders(val *[]*string)
	ExcludedHeadersInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludedHeaders() *[]*string
	SetIncludedHeaders(val *[]*string)
	IncludedHeadersInput() *[]*string
	InternalValue() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPattern
	SetInternalValue(val *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPattern)
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
	PutAll(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAll)
	ResetAll()
	ResetExcludedHeaders()
	ResetIncludedHeaders()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference
type jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) All() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAllOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAllOutputReference
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) AllInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAll {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAll
	_jsii_.Get(
		j,
		"allInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ExcludedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ExcludedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) IncludedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) IncludedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) InternalValue() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPattern {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPattern
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference_Override(w Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetExcludedHeaders(val *[]*string) {
	if err := j.validateSetExcludedHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedHeaders",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetIncludedHeaders(val *[]*string) {
	if err := j.validateSetIncludedHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedHeaders",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetInternalValue(val *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPattern) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) PutAll(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAll) {
	if err := w.validatePutAllParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAll",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ResetAll() {
	_jsii_.InvokeVoid(
		w,
		"resetAll",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ResetExcludedHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetExcludedHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ResetIncludedHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetIncludedHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

