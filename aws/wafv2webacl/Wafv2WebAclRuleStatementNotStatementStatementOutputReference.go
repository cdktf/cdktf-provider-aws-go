// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleStatementNotStatementStatementOutputReference interface {
	cdktf.ComplexObject
	AndStatement() Wafv2WebAclRuleStatementNotStatementStatementAndStatementOutputReference
	AndStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementAndStatement
	ByteMatchStatement() Wafv2WebAclRuleStatementNotStatementStatementByteMatchStatementOutputReference
	ByteMatchStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementByteMatchStatement
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
	GeoMatchStatement() Wafv2WebAclRuleStatementNotStatementStatementGeoMatchStatementOutputReference
	GeoMatchStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementGeoMatchStatement
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpSetReferenceStatement() Wafv2WebAclRuleStatementNotStatementStatementIpSetReferenceStatementOutputReference
	IpSetReferenceStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementIpSetReferenceStatement
	LabelMatchStatement() Wafv2WebAclRuleStatementNotStatementStatementLabelMatchStatementOutputReference
	LabelMatchStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementLabelMatchStatement
	NotStatement() Wafv2WebAclRuleStatementNotStatementOutputReference
	NotStatementInput() *Wafv2WebAclRuleStatementNotStatement
	OrStatement() Wafv2WebAclRuleStatementNotStatementStatementOrStatementOutputReference
	OrStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementOrStatement
	RegexMatchStatement() Wafv2WebAclRuleStatementNotStatementStatementRegexMatchStatementOutputReference
	RegexMatchStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementRegexMatchStatement
	RegexPatternSetReferenceStatement() Wafv2WebAclRuleStatementNotStatementStatementRegexPatternSetReferenceStatementOutputReference
	RegexPatternSetReferenceStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementRegexPatternSetReferenceStatement
	SizeConstraintStatement() Wafv2WebAclRuleStatementNotStatementStatementSizeConstraintStatementOutputReference
	SizeConstraintStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementSizeConstraintStatement
	SqliMatchStatement() Wafv2WebAclRuleStatementNotStatementStatementSqliMatchStatementOutputReference
	SqliMatchStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementSqliMatchStatement
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	XssMatchStatement() Wafv2WebAclRuleStatementNotStatementStatementXssMatchStatementOutputReference
	XssMatchStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementXssMatchStatement
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
	PutAndStatement(value *Wafv2WebAclRuleStatementNotStatementStatementAndStatement)
	PutByteMatchStatement(value *Wafv2WebAclRuleStatementNotStatementStatementByteMatchStatement)
	PutGeoMatchStatement(value *Wafv2WebAclRuleStatementNotStatementStatementGeoMatchStatement)
	PutIpSetReferenceStatement(value *Wafv2WebAclRuleStatementNotStatementStatementIpSetReferenceStatement)
	PutLabelMatchStatement(value *Wafv2WebAclRuleStatementNotStatementStatementLabelMatchStatement)
	PutNotStatement(value *Wafv2WebAclRuleStatementNotStatement)
	PutOrStatement(value *Wafv2WebAclRuleStatementNotStatementStatementOrStatement)
	PutRegexMatchStatement(value *Wafv2WebAclRuleStatementNotStatementStatementRegexMatchStatement)
	PutRegexPatternSetReferenceStatement(value *Wafv2WebAclRuleStatementNotStatementStatementRegexPatternSetReferenceStatement)
	PutSizeConstraintStatement(value *Wafv2WebAclRuleStatementNotStatementStatementSizeConstraintStatement)
	PutSqliMatchStatement(value *Wafv2WebAclRuleStatementNotStatementStatementSqliMatchStatement)
	PutXssMatchStatement(value *Wafv2WebAclRuleStatementNotStatementStatementXssMatchStatement)
	ResetAndStatement()
	ResetByteMatchStatement()
	ResetGeoMatchStatement()
	ResetIpSetReferenceStatement()
	ResetLabelMatchStatement()
	ResetNotStatement()
	ResetOrStatement()
	ResetRegexMatchStatement()
	ResetRegexPatternSetReferenceStatement()
	ResetSizeConstraintStatement()
	ResetSqliMatchStatement()
	ResetXssMatchStatement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2WebAclRuleStatementNotStatementStatementOutputReference
type jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) AndStatement() Wafv2WebAclRuleStatementNotStatementStatementAndStatementOutputReference {
	var returns Wafv2WebAclRuleStatementNotStatementStatementAndStatementOutputReference
	_jsii_.Get(
		j,
		"andStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) AndStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementAndStatement {
	var returns *Wafv2WebAclRuleStatementNotStatementStatementAndStatement
	_jsii_.Get(
		j,
		"andStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ByteMatchStatement() Wafv2WebAclRuleStatementNotStatementStatementByteMatchStatementOutputReference {
	var returns Wafv2WebAclRuleStatementNotStatementStatementByteMatchStatementOutputReference
	_jsii_.Get(
		j,
		"byteMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ByteMatchStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementByteMatchStatement {
	var returns *Wafv2WebAclRuleStatementNotStatementStatementByteMatchStatement
	_jsii_.Get(
		j,
		"byteMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) GeoMatchStatement() Wafv2WebAclRuleStatementNotStatementStatementGeoMatchStatementOutputReference {
	var returns Wafv2WebAclRuleStatementNotStatementStatementGeoMatchStatementOutputReference
	_jsii_.Get(
		j,
		"geoMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) GeoMatchStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementGeoMatchStatement {
	var returns *Wafv2WebAclRuleStatementNotStatementStatementGeoMatchStatement
	_jsii_.Get(
		j,
		"geoMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) IpSetReferenceStatement() Wafv2WebAclRuleStatementNotStatementStatementIpSetReferenceStatementOutputReference {
	var returns Wafv2WebAclRuleStatementNotStatementStatementIpSetReferenceStatementOutputReference
	_jsii_.Get(
		j,
		"ipSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) IpSetReferenceStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementIpSetReferenceStatement {
	var returns *Wafv2WebAclRuleStatementNotStatementStatementIpSetReferenceStatement
	_jsii_.Get(
		j,
		"ipSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) LabelMatchStatement() Wafv2WebAclRuleStatementNotStatementStatementLabelMatchStatementOutputReference {
	var returns Wafv2WebAclRuleStatementNotStatementStatementLabelMatchStatementOutputReference
	_jsii_.Get(
		j,
		"labelMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) LabelMatchStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementLabelMatchStatement {
	var returns *Wafv2WebAclRuleStatementNotStatementStatementLabelMatchStatement
	_jsii_.Get(
		j,
		"labelMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) NotStatement() Wafv2WebAclRuleStatementNotStatementOutputReference {
	var returns Wafv2WebAclRuleStatementNotStatementOutputReference
	_jsii_.Get(
		j,
		"notStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) NotStatementInput() *Wafv2WebAclRuleStatementNotStatement {
	var returns *Wafv2WebAclRuleStatementNotStatement
	_jsii_.Get(
		j,
		"notStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) OrStatement() Wafv2WebAclRuleStatementNotStatementStatementOrStatementOutputReference {
	var returns Wafv2WebAclRuleStatementNotStatementStatementOrStatementOutputReference
	_jsii_.Get(
		j,
		"orStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) OrStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementOrStatement {
	var returns *Wafv2WebAclRuleStatementNotStatementStatementOrStatement
	_jsii_.Get(
		j,
		"orStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) RegexMatchStatement() Wafv2WebAclRuleStatementNotStatementStatementRegexMatchStatementOutputReference {
	var returns Wafv2WebAclRuleStatementNotStatementStatementRegexMatchStatementOutputReference
	_jsii_.Get(
		j,
		"regexMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) RegexMatchStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementRegexMatchStatement {
	var returns *Wafv2WebAclRuleStatementNotStatementStatementRegexMatchStatement
	_jsii_.Get(
		j,
		"regexMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) RegexPatternSetReferenceStatement() Wafv2WebAclRuleStatementNotStatementStatementRegexPatternSetReferenceStatementOutputReference {
	var returns Wafv2WebAclRuleStatementNotStatementStatementRegexPatternSetReferenceStatementOutputReference
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) RegexPatternSetReferenceStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementRegexPatternSetReferenceStatement {
	var returns *Wafv2WebAclRuleStatementNotStatementStatementRegexPatternSetReferenceStatement
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) SizeConstraintStatement() Wafv2WebAclRuleStatementNotStatementStatementSizeConstraintStatementOutputReference {
	var returns Wafv2WebAclRuleStatementNotStatementStatementSizeConstraintStatementOutputReference
	_jsii_.Get(
		j,
		"sizeConstraintStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) SizeConstraintStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementSizeConstraintStatement {
	var returns *Wafv2WebAclRuleStatementNotStatementStatementSizeConstraintStatement
	_jsii_.Get(
		j,
		"sizeConstraintStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) SqliMatchStatement() Wafv2WebAclRuleStatementNotStatementStatementSqliMatchStatementOutputReference {
	var returns Wafv2WebAclRuleStatementNotStatementStatementSqliMatchStatementOutputReference
	_jsii_.Get(
		j,
		"sqliMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) SqliMatchStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementSqliMatchStatement {
	var returns *Wafv2WebAclRuleStatementNotStatementStatementSqliMatchStatement
	_jsii_.Get(
		j,
		"sqliMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) XssMatchStatement() Wafv2WebAclRuleStatementNotStatementStatementXssMatchStatementOutputReference {
	var returns Wafv2WebAclRuleStatementNotStatementStatementXssMatchStatementOutputReference
	_jsii_.Get(
		j,
		"xssMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) XssMatchStatementInput() *Wafv2WebAclRuleStatementNotStatementStatementXssMatchStatement {
	var returns *Wafv2WebAclRuleStatementNotStatementStatementXssMatchStatement
	_jsii_.Get(
		j,
		"xssMatchStatementInput",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleStatementNotStatementStatementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Wafv2WebAclRuleStatementNotStatementStatementOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleStatementNotStatementStatementOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementNotStatementStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleStatementNotStatementStatementOutputReference_Override(w Wafv2WebAclRuleStatementNotStatementStatementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementNotStatementStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) PutAndStatement(value *Wafv2WebAclRuleStatementNotStatementStatementAndStatement) {
	if err := w.validatePutAndStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAndStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) PutByteMatchStatement(value *Wafv2WebAclRuleStatementNotStatementStatementByteMatchStatement) {
	if err := w.validatePutByteMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putByteMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) PutGeoMatchStatement(value *Wafv2WebAclRuleStatementNotStatementStatementGeoMatchStatement) {
	if err := w.validatePutGeoMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putGeoMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) PutIpSetReferenceStatement(value *Wafv2WebAclRuleStatementNotStatementStatementIpSetReferenceStatement) {
	if err := w.validatePutIpSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIpSetReferenceStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) PutLabelMatchStatement(value *Wafv2WebAclRuleStatementNotStatementStatementLabelMatchStatement) {
	if err := w.validatePutLabelMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLabelMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) PutNotStatement(value *Wafv2WebAclRuleStatementNotStatement) {
	if err := w.validatePutNotStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putNotStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) PutOrStatement(value *Wafv2WebAclRuleStatementNotStatementStatementOrStatement) {
	if err := w.validatePutOrStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putOrStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) PutRegexMatchStatement(value *Wafv2WebAclRuleStatementNotStatementStatementRegexMatchStatement) {
	if err := w.validatePutRegexMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRegexMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) PutRegexPatternSetReferenceStatement(value *Wafv2WebAclRuleStatementNotStatementStatementRegexPatternSetReferenceStatement) {
	if err := w.validatePutRegexPatternSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRegexPatternSetReferenceStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) PutSizeConstraintStatement(value *Wafv2WebAclRuleStatementNotStatementStatementSizeConstraintStatement) {
	if err := w.validatePutSizeConstraintStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSizeConstraintStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) PutSqliMatchStatement(value *Wafv2WebAclRuleStatementNotStatementStatementSqliMatchStatement) {
	if err := w.validatePutSqliMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSqliMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) PutXssMatchStatement(value *Wafv2WebAclRuleStatementNotStatementStatementXssMatchStatement) {
	if err := w.validatePutXssMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putXssMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ResetAndStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetAndStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ResetByteMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetByteMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ResetGeoMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetGeoMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ResetIpSetReferenceStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetIpSetReferenceStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ResetLabelMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetLabelMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ResetNotStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetNotStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ResetOrStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetOrStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ResetRegexMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetRegexMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ResetRegexPatternSetReferenceStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetRegexPatternSetReferenceStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ResetSizeConstraintStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetSizeConstraintStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ResetSqliMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetSqliMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ResetXssMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetXssMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementNotStatementStatementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

