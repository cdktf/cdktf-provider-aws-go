package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference interface {
	cdktf.ComplexObject
	All() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAllOutputReference
	AllInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAll
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
	InternalValue() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPattern
	SetInternalValue(val *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPattern)
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
	PutAll(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAll)
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

// The jsii proxy struct for Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference
type jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) All() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAllOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAllOutputReference
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) AllInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAll {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAll
	_jsii_.Get(
		j,
		"allInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ExcludedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ExcludedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) IncludedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) IncludedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) InternalValue() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPattern {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPattern
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference_Override(w Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetExcludedHeaders(val *[]*string) {
	if err := j.validateSetExcludedHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedHeaders",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetIncludedHeaders(val *[]*string) {
	if err := j.validateSetIncludedHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedHeaders",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetInternalValue(val *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPattern) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) PutAll(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternAll) {
	if err := w.validatePutAllParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAll",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ResetAll() {
	_jsii_.InvokeVoid(
		w,
		"resetAll",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ResetExcludedHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetExcludedHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ResetIncludedHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetIncludedHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersMatchPatternOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

