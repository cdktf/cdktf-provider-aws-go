package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference interface {
	cdktf.ComplexObject
	ByteMatchStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementByteMatchStatementOutputReference
	ByteMatchStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementByteMatchStatement
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
	GeoMatchStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementGeoMatchStatementOutputReference
	GeoMatchStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementGeoMatchStatement
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpSetReferenceStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementIpSetReferenceStatementOutputReference
	IpSetReferenceStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementIpSetReferenceStatement
	LabelMatchStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementLabelMatchStatementOutputReference
	LabelMatchStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementLabelMatchStatement
	RegexMatchStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexMatchStatementOutputReference
	RegexMatchStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexMatchStatement
	RegexPatternSetReferenceStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementOutputReference
	RegexPatternSetReferenceStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatement
	SizeConstraintStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementOutputReference
	SizeConstraintStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatement
	SqliMatchStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatementOutputReference
	SqliMatchStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatement
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	XssMatchStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementXssMatchStatementOutputReference
	XssMatchStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementXssMatchStatement
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
	PutByteMatchStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementByteMatchStatement)
	PutGeoMatchStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementGeoMatchStatement)
	PutIpSetReferenceStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementIpSetReferenceStatement)
	PutLabelMatchStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementLabelMatchStatement)
	PutRegexMatchStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexMatchStatement)
	PutRegexPatternSetReferenceStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatement)
	PutSizeConstraintStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatement)
	PutSqliMatchStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatement)
	PutXssMatchStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementXssMatchStatement)
	ResetByteMatchStatement()
	ResetGeoMatchStatement()
	ResetIpSetReferenceStatement()
	ResetLabelMatchStatement()
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

// The jsii proxy struct for Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ByteMatchStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementByteMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementByteMatchStatementOutputReference
	_jsii_.Get(
		j,
		"byteMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ByteMatchStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementByteMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementByteMatchStatement
	_jsii_.Get(
		j,
		"byteMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) GeoMatchStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementGeoMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementGeoMatchStatementOutputReference
	_jsii_.Get(
		j,
		"geoMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) GeoMatchStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementGeoMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementGeoMatchStatement
	_jsii_.Get(
		j,
		"geoMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) IpSetReferenceStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementIpSetReferenceStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementIpSetReferenceStatementOutputReference
	_jsii_.Get(
		j,
		"ipSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) IpSetReferenceStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementIpSetReferenceStatement {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementIpSetReferenceStatement
	_jsii_.Get(
		j,
		"ipSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) LabelMatchStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementLabelMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementLabelMatchStatementOutputReference
	_jsii_.Get(
		j,
		"labelMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) LabelMatchStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementLabelMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementLabelMatchStatement
	_jsii_.Get(
		j,
		"labelMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) RegexMatchStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexMatchStatementOutputReference
	_jsii_.Get(
		j,
		"regexMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) RegexMatchStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexMatchStatement
	_jsii_.Get(
		j,
		"regexMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) RegexPatternSetReferenceStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementOutputReference
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) RegexPatternSetReferenceStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatement {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatement
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) SizeConstraintStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementOutputReference
	_jsii_.Get(
		j,
		"sizeConstraintStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) SizeConstraintStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatement {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatement
	_jsii_.Get(
		j,
		"sizeConstraintStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) SqliMatchStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatementOutputReference
	_jsii_.Get(
		j,
		"sqliMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) SqliMatchStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatement
	_jsii_.Get(
		j,
		"sqliMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) XssMatchStatement() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementXssMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementXssMatchStatementOutputReference
	_jsii_.Get(
		j,
		"xssMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) XssMatchStatementInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementXssMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementXssMatchStatement
	_jsii_.Get(
		j,
		"xssMatchStatementInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference_Override(w Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) PutByteMatchStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementByteMatchStatement) {
	if err := w.validatePutByteMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putByteMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) PutGeoMatchStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementGeoMatchStatement) {
	if err := w.validatePutGeoMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putGeoMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) PutIpSetReferenceStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementIpSetReferenceStatement) {
	if err := w.validatePutIpSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIpSetReferenceStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) PutLabelMatchStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementLabelMatchStatement) {
	if err := w.validatePutLabelMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLabelMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) PutRegexMatchStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexMatchStatement) {
	if err := w.validatePutRegexMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRegexMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) PutRegexPatternSetReferenceStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatement) {
	if err := w.validatePutRegexPatternSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRegexPatternSetReferenceStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) PutSizeConstraintStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatement) {
	if err := w.validatePutSizeConstraintStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSizeConstraintStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) PutSqliMatchStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatement) {
	if err := w.validatePutSqliMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSqliMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) PutXssMatchStatement(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementXssMatchStatement) {
	if err := w.validatePutXssMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putXssMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ResetByteMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetByteMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ResetGeoMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetGeoMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ResetIpSetReferenceStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetIpSetReferenceStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ResetLabelMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetLabelMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ResetRegexMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetRegexMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ResetRegexPatternSetReferenceStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetRegexPatternSetReferenceStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ResetSizeConstraintStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetSizeConstraintStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ResetSqliMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetSqliMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ResetXssMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetXssMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

