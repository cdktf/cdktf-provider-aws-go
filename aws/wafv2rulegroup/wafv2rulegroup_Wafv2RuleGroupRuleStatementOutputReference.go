package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementOutputReference interface {
	cdktf.ComplexObject
	AndStatement() Wafv2RuleGroupRuleStatementAndStatementOutputReference
	AndStatementInput() *Wafv2RuleGroupRuleStatementAndStatement
	ByteMatchStatement() Wafv2RuleGroupRuleStatementByteMatchStatementOutputReference
	ByteMatchStatementInput() *Wafv2RuleGroupRuleStatementByteMatchStatement
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
	GeoMatchStatement() Wafv2RuleGroupRuleStatementGeoMatchStatementOutputReference
	GeoMatchStatementInput() *Wafv2RuleGroupRuleStatementGeoMatchStatement
	InternalValue() *Wafv2RuleGroupRuleStatement
	SetInternalValue(val *Wafv2RuleGroupRuleStatement)
	IpSetReferenceStatement() Wafv2RuleGroupRuleStatementIpSetReferenceStatementOutputReference
	IpSetReferenceStatementInput() *Wafv2RuleGroupRuleStatementIpSetReferenceStatement
	LabelMatchStatement() Wafv2RuleGroupRuleStatementLabelMatchStatementOutputReference
	LabelMatchStatementInput() *Wafv2RuleGroupRuleStatementLabelMatchStatement
	NotStatement() Wafv2RuleGroupRuleStatementNotStatementOutputReference
	NotStatementInput() *Wafv2RuleGroupRuleStatementNotStatement
	OrStatement() Wafv2RuleGroupRuleStatementOrStatementOutputReference
	OrStatementInput() *Wafv2RuleGroupRuleStatementOrStatement
	RateBasedStatement() Wafv2RuleGroupRuleStatementRateBasedStatementOutputReference
	RateBasedStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatement
	RegexMatchStatement() Wafv2RuleGroupRuleStatementRegexMatchStatementOutputReference
	RegexMatchStatementInput() *Wafv2RuleGroupRuleStatementRegexMatchStatement
	RegexPatternSetReferenceStatement() Wafv2RuleGroupRuleStatementRegexPatternSetReferenceStatementOutputReference
	RegexPatternSetReferenceStatementInput() *Wafv2RuleGroupRuleStatementRegexPatternSetReferenceStatement
	SizeConstraintStatement() Wafv2RuleGroupRuleStatementSizeConstraintStatementOutputReference
	SizeConstraintStatementInput() *Wafv2RuleGroupRuleStatementSizeConstraintStatement
	SqliMatchStatement() Wafv2RuleGroupRuleStatementSqliMatchStatementOutputReference
	SqliMatchStatementInput() *Wafv2RuleGroupRuleStatementSqliMatchStatement
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	XssMatchStatement() Wafv2RuleGroupRuleStatementXssMatchStatementOutputReference
	XssMatchStatementInput() *Wafv2RuleGroupRuleStatementXssMatchStatement
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
	PutAndStatement(value *Wafv2RuleGroupRuleStatementAndStatement)
	PutByteMatchStatement(value *Wafv2RuleGroupRuleStatementByteMatchStatement)
	PutGeoMatchStatement(value *Wafv2RuleGroupRuleStatementGeoMatchStatement)
	PutIpSetReferenceStatement(value *Wafv2RuleGroupRuleStatementIpSetReferenceStatement)
	PutLabelMatchStatement(value *Wafv2RuleGroupRuleStatementLabelMatchStatement)
	PutNotStatement(value *Wafv2RuleGroupRuleStatementNotStatement)
	PutOrStatement(value *Wafv2RuleGroupRuleStatementOrStatement)
	PutRateBasedStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatement)
	PutRegexMatchStatement(value *Wafv2RuleGroupRuleStatementRegexMatchStatement)
	PutRegexPatternSetReferenceStatement(value *Wafv2RuleGroupRuleStatementRegexPatternSetReferenceStatement)
	PutSizeConstraintStatement(value *Wafv2RuleGroupRuleStatementSizeConstraintStatement)
	PutSqliMatchStatement(value *Wafv2RuleGroupRuleStatementSqliMatchStatement)
	PutXssMatchStatement(value *Wafv2RuleGroupRuleStatementXssMatchStatement)
	ResetAndStatement()
	ResetByteMatchStatement()
	ResetGeoMatchStatement()
	ResetIpSetReferenceStatement()
	ResetLabelMatchStatement()
	ResetNotStatement()
	ResetOrStatement()
	ResetRateBasedStatement()
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

// The jsii proxy struct for Wafv2RuleGroupRuleStatementOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) AndStatement() Wafv2RuleGroupRuleStatementAndStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementOutputReference
	_jsii_.Get(
		j,
		"andStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) AndStatementInput() *Wafv2RuleGroupRuleStatementAndStatement {
	var returns *Wafv2RuleGroupRuleStatementAndStatement
	_jsii_.Get(
		j,
		"andStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ByteMatchStatement() Wafv2RuleGroupRuleStatementByteMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementByteMatchStatementOutputReference
	_jsii_.Get(
		j,
		"byteMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ByteMatchStatementInput() *Wafv2RuleGroupRuleStatementByteMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementByteMatchStatement
	_jsii_.Get(
		j,
		"byteMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) GeoMatchStatement() Wafv2RuleGroupRuleStatementGeoMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementGeoMatchStatementOutputReference
	_jsii_.Get(
		j,
		"geoMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) GeoMatchStatementInput() *Wafv2RuleGroupRuleStatementGeoMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementGeoMatchStatement
	_jsii_.Get(
		j,
		"geoMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) InternalValue() *Wafv2RuleGroupRuleStatement {
	var returns *Wafv2RuleGroupRuleStatement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) IpSetReferenceStatement() Wafv2RuleGroupRuleStatementIpSetReferenceStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementIpSetReferenceStatementOutputReference
	_jsii_.Get(
		j,
		"ipSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) IpSetReferenceStatementInput() *Wafv2RuleGroupRuleStatementIpSetReferenceStatement {
	var returns *Wafv2RuleGroupRuleStatementIpSetReferenceStatement
	_jsii_.Get(
		j,
		"ipSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) LabelMatchStatement() Wafv2RuleGroupRuleStatementLabelMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementLabelMatchStatementOutputReference
	_jsii_.Get(
		j,
		"labelMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) LabelMatchStatementInput() *Wafv2RuleGroupRuleStatementLabelMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementLabelMatchStatement
	_jsii_.Get(
		j,
		"labelMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) NotStatement() Wafv2RuleGroupRuleStatementNotStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementOutputReference
	_jsii_.Get(
		j,
		"notStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) NotStatementInput() *Wafv2RuleGroupRuleStatementNotStatement {
	var returns *Wafv2RuleGroupRuleStatementNotStatement
	_jsii_.Get(
		j,
		"notStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) OrStatement() Wafv2RuleGroupRuleStatementOrStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementOutputReference
	_jsii_.Get(
		j,
		"orStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) OrStatementInput() *Wafv2RuleGroupRuleStatementOrStatement {
	var returns *Wafv2RuleGroupRuleStatementOrStatement
	_jsii_.Get(
		j,
		"orStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) RateBasedStatement() Wafv2RuleGroupRuleStatementRateBasedStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementOutputReference
	_jsii_.Get(
		j,
		"rateBasedStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) RateBasedStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatement
	_jsii_.Get(
		j,
		"rateBasedStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) RegexMatchStatement() Wafv2RuleGroupRuleStatementRegexMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRegexMatchStatementOutputReference
	_jsii_.Get(
		j,
		"regexMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) RegexMatchStatementInput() *Wafv2RuleGroupRuleStatementRegexMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementRegexMatchStatement
	_jsii_.Get(
		j,
		"regexMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) RegexPatternSetReferenceStatement() Wafv2RuleGroupRuleStatementRegexPatternSetReferenceStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRegexPatternSetReferenceStatementOutputReference
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) RegexPatternSetReferenceStatementInput() *Wafv2RuleGroupRuleStatementRegexPatternSetReferenceStatement {
	var returns *Wafv2RuleGroupRuleStatementRegexPatternSetReferenceStatement
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) SizeConstraintStatement() Wafv2RuleGroupRuleStatementSizeConstraintStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementSizeConstraintStatementOutputReference
	_jsii_.Get(
		j,
		"sizeConstraintStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) SizeConstraintStatementInput() *Wafv2RuleGroupRuleStatementSizeConstraintStatement {
	var returns *Wafv2RuleGroupRuleStatementSizeConstraintStatement
	_jsii_.Get(
		j,
		"sizeConstraintStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) SqliMatchStatement() Wafv2RuleGroupRuleStatementSqliMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementSqliMatchStatementOutputReference
	_jsii_.Get(
		j,
		"sqliMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) SqliMatchStatementInput() *Wafv2RuleGroupRuleStatementSqliMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementSqliMatchStatement
	_jsii_.Get(
		j,
		"sqliMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) XssMatchStatement() Wafv2RuleGroupRuleStatementXssMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementXssMatchStatementOutputReference
	_jsii_.Get(
		j,
		"xssMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) XssMatchStatementInput() *Wafv2RuleGroupRuleStatementXssMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementXssMatchStatement
	_jsii_.Get(
		j,
		"xssMatchStatementInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementOutputReference_Override(w Wafv2RuleGroupRuleStatementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatement) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) PutAndStatement(value *Wafv2RuleGroupRuleStatementAndStatement) {
	if err := w.validatePutAndStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAndStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) PutByteMatchStatement(value *Wafv2RuleGroupRuleStatementByteMatchStatement) {
	if err := w.validatePutByteMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putByteMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) PutGeoMatchStatement(value *Wafv2RuleGroupRuleStatementGeoMatchStatement) {
	if err := w.validatePutGeoMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putGeoMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) PutIpSetReferenceStatement(value *Wafv2RuleGroupRuleStatementIpSetReferenceStatement) {
	if err := w.validatePutIpSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIpSetReferenceStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) PutLabelMatchStatement(value *Wafv2RuleGroupRuleStatementLabelMatchStatement) {
	if err := w.validatePutLabelMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLabelMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) PutNotStatement(value *Wafv2RuleGroupRuleStatementNotStatement) {
	if err := w.validatePutNotStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putNotStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) PutOrStatement(value *Wafv2RuleGroupRuleStatementOrStatement) {
	if err := w.validatePutOrStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putOrStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) PutRateBasedStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatement) {
	if err := w.validatePutRateBasedStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRateBasedStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) PutRegexMatchStatement(value *Wafv2RuleGroupRuleStatementRegexMatchStatement) {
	if err := w.validatePutRegexMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRegexMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) PutRegexPatternSetReferenceStatement(value *Wafv2RuleGroupRuleStatementRegexPatternSetReferenceStatement) {
	if err := w.validatePutRegexPatternSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRegexPatternSetReferenceStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) PutSizeConstraintStatement(value *Wafv2RuleGroupRuleStatementSizeConstraintStatement) {
	if err := w.validatePutSizeConstraintStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSizeConstraintStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) PutSqliMatchStatement(value *Wafv2RuleGroupRuleStatementSqliMatchStatement) {
	if err := w.validatePutSqliMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSqliMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) PutXssMatchStatement(value *Wafv2RuleGroupRuleStatementXssMatchStatement) {
	if err := w.validatePutXssMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putXssMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ResetAndStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetAndStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ResetByteMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetByteMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ResetGeoMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetGeoMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ResetIpSetReferenceStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetIpSetReferenceStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ResetLabelMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetLabelMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ResetNotStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetNotStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ResetOrStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetOrStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ResetRateBasedStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetRateBasedStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ResetRegexMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetRegexMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ResetRegexPatternSetReferenceStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetRegexPatternSetReferenceStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ResetSizeConstraintStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetSizeConstraintStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ResetSqliMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetSqliMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ResetXssMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetXssMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

