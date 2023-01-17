package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference interface {
	cdktf.ComplexObject
	AndStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementOutputReference
	AndStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatement
	ByteMatchStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementByteMatchStatementOutputReference
	ByteMatchStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementByteMatchStatement
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
	GeoMatchStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatementOutputReference
	GeoMatchStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatement
	InternalValue() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatement
	SetInternalValue(val *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatement)
	IpSetReferenceStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatementOutputReference
	IpSetReferenceStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatement
	LabelMatchStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementLabelMatchStatementOutputReference
	LabelMatchStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementLabelMatchStatement
	NotStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementOutputReference
	NotStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatement
	OrStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementOutputReference
	OrStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatement
	RegexMatchStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementOutputReference
	RegexMatchStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatement
	RegexPatternSetReferenceStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementOutputReference
	RegexPatternSetReferenceStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatement
	SizeConstraintStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatementOutputReference
	SizeConstraintStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatement
	SqliMatchStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementOutputReference
	SqliMatchStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatement
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	XssMatchStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementOutputReference
	XssMatchStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementXssMatchStatement
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
	PutAndStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatement)
	PutByteMatchStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementByteMatchStatement)
	PutGeoMatchStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatement)
	PutIpSetReferenceStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatement)
	PutLabelMatchStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementLabelMatchStatement)
	PutNotStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatement)
	PutOrStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatement)
	PutRegexMatchStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatement)
	PutRegexPatternSetReferenceStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatement)
	PutSizeConstraintStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatement)
	PutSqliMatchStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatement)
	PutXssMatchStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementXssMatchStatement)
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

// The jsii proxy struct for Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) AndStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementOutputReference
	_jsii_.Get(
		j,
		"andStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) AndStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatement
	_jsii_.Get(
		j,
		"andStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ByteMatchStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementByteMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementByteMatchStatementOutputReference
	_jsii_.Get(
		j,
		"byteMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ByteMatchStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementByteMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementByteMatchStatement
	_jsii_.Get(
		j,
		"byteMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) GeoMatchStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatementOutputReference
	_jsii_.Get(
		j,
		"geoMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) GeoMatchStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatement
	_jsii_.Get(
		j,
		"geoMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) InternalValue() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) IpSetReferenceStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatementOutputReference
	_jsii_.Get(
		j,
		"ipSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) IpSetReferenceStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatement
	_jsii_.Get(
		j,
		"ipSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) LabelMatchStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementLabelMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementLabelMatchStatementOutputReference
	_jsii_.Get(
		j,
		"labelMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) LabelMatchStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementLabelMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementLabelMatchStatement
	_jsii_.Get(
		j,
		"labelMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) NotStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementOutputReference
	_jsii_.Get(
		j,
		"notStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) NotStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatement
	_jsii_.Get(
		j,
		"notStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) OrStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementOutputReference
	_jsii_.Get(
		j,
		"orStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) OrStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatement
	_jsii_.Get(
		j,
		"orStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) RegexMatchStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementOutputReference
	_jsii_.Get(
		j,
		"regexMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) RegexMatchStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatement
	_jsii_.Get(
		j,
		"regexMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) RegexPatternSetReferenceStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementOutputReference
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) RegexPatternSetReferenceStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatement
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) SizeConstraintStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatementOutputReference
	_jsii_.Get(
		j,
		"sizeConstraintStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) SizeConstraintStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatement
	_jsii_.Get(
		j,
		"sizeConstraintStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) SqliMatchStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementOutputReference
	_jsii_.Get(
		j,
		"sqliMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) SqliMatchStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatement
	_jsii_.Get(
		j,
		"sqliMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) XssMatchStatement() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementOutputReference
	_jsii_.Get(
		j,
		"xssMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) XssMatchStatementInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementXssMatchStatement {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementXssMatchStatement
	_jsii_.Get(
		j,
		"xssMatchStatementInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference_Override(w Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatement) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) PutAndStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatement) {
	if err := w.validatePutAndStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAndStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) PutByteMatchStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementByteMatchStatement) {
	if err := w.validatePutByteMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putByteMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) PutGeoMatchStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatement) {
	if err := w.validatePutGeoMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putGeoMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) PutIpSetReferenceStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatement) {
	if err := w.validatePutIpSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIpSetReferenceStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) PutLabelMatchStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementLabelMatchStatement) {
	if err := w.validatePutLabelMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLabelMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) PutNotStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatement) {
	if err := w.validatePutNotStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putNotStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) PutOrStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatement) {
	if err := w.validatePutOrStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putOrStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) PutRegexMatchStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatement) {
	if err := w.validatePutRegexMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRegexMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) PutRegexPatternSetReferenceStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatement) {
	if err := w.validatePutRegexPatternSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRegexPatternSetReferenceStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) PutSizeConstraintStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatement) {
	if err := w.validatePutSizeConstraintStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSizeConstraintStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) PutSqliMatchStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatement) {
	if err := w.validatePutSqliMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSqliMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) PutXssMatchStatement(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementXssMatchStatement) {
	if err := w.validatePutXssMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putXssMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ResetAndStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetAndStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ResetByteMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetByteMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ResetGeoMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetGeoMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ResetIpSetReferenceStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetIpSetReferenceStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ResetLabelMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetLabelMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ResetNotStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetNotStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ResetOrStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetOrStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ResetRegexMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetRegexMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ResetRegexPatternSetReferenceStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetRegexPatternSetReferenceStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ResetSizeConstraintStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetSizeConstraintStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ResetSqliMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetSqliMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ResetXssMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetXssMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

