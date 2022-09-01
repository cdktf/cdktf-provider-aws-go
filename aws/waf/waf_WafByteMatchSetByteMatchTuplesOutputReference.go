package waf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/waf/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafByteMatchSetByteMatchTuplesOutputReference interface {
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
	FieldToMatch() WafByteMatchSetByteMatchTuplesFieldToMatchOutputReference
	FieldToMatchInput() *WafByteMatchSetByteMatchTuplesFieldToMatch
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PositionalConstraint() *string
	SetPositionalConstraint(val *string)
	PositionalConstraintInput() *string
	TargetString() *string
	SetTargetString(val *string)
	TargetStringInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextTransformation() *string
	SetTextTransformation(val *string)
	TextTransformationInput() *string
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
	PutFieldToMatch(value *WafByteMatchSetByteMatchTuplesFieldToMatch)
	ResetTargetString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WafByteMatchSetByteMatchTuplesOutputReference
type jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) FieldToMatch() WafByteMatchSetByteMatchTuplesFieldToMatchOutputReference {
	var returns WafByteMatchSetByteMatchTuplesFieldToMatchOutputReference
	_jsii_.Get(
		j,
		"fieldToMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) FieldToMatchInput() *WafByteMatchSetByteMatchTuplesFieldToMatch {
	var returns *WafByteMatchSetByteMatchTuplesFieldToMatch
	_jsii_.Get(
		j,
		"fieldToMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) PositionalConstraint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"positionalConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) PositionalConstraintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"positionalConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) TargetString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) TargetStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) TextTransformation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textTransformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) TextTransformationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textTransformationInput",
		&returns,
	)
	return returns
}


func NewWafByteMatchSetByteMatchTuplesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WafByteMatchSetByteMatchTuplesOutputReference {
	_init_.Initialize()

	if err := validateNewWafByteMatchSetByteMatchTuplesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.waf.WafByteMatchSetByteMatchTuplesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafByteMatchSetByteMatchTuplesOutputReference_Override(w WafByteMatchSetByteMatchTuplesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.waf.WafByteMatchSetByteMatchTuplesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference)SetPositionalConstraint(val *string) {
	if err := j.validateSetPositionalConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"positionalConstraint",
		val,
	)
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference)SetTargetString(val *string) {
	if err := j.validateSetTargetStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetString",
		val,
	)
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference)SetTextTransformation(val *string) {
	if err := j.validateSetTextTransformationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textTransformation",
		val,
	)
}

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) PutFieldToMatch(value *WafByteMatchSetByteMatchTuplesFieldToMatch) {
	if err := w.validatePutFieldToMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putFieldToMatch",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) ResetTargetString() {
	_jsii_.InvokeVoid(
		w,
		"resetTargetString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WafByteMatchSetByteMatchTuplesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

