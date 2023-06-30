package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleStatementOutputReference interface {
	cdktf.ComplexObject
	AndStatement() Wafv2WebAclRuleStatementAndStatementOutputReference
	AndStatementInput() *Wafv2WebAclRuleStatementAndStatement
	ByteMatchStatement() Wafv2WebAclRuleStatementByteMatchStatementOutputReference
	ByteMatchStatementInput() *Wafv2WebAclRuleStatementByteMatchStatement
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
	GeoMatchStatement() Wafv2WebAclRuleStatementGeoMatchStatementOutputReference
	GeoMatchStatementInput() *Wafv2WebAclRuleStatementGeoMatchStatement
	InternalValue() *Wafv2WebAclRuleStatement
	SetInternalValue(val *Wafv2WebAclRuleStatement)
	IpSetReferenceStatement() Wafv2WebAclRuleStatementIpSetReferenceStatementOutputReference
	IpSetReferenceStatementInput() *Wafv2WebAclRuleStatementIpSetReferenceStatement
	LabelMatchStatement() Wafv2WebAclRuleStatementLabelMatchStatementOutputReference
	LabelMatchStatementInput() *Wafv2WebAclRuleStatementLabelMatchStatement
	ManagedRuleGroupStatement() Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference
	ManagedRuleGroupStatementInput() *Wafv2WebAclRuleStatementManagedRuleGroupStatement
	NotStatement() Wafv2WebAclRuleStatementNotStatementOutputReference
	NotStatementInput() *Wafv2WebAclRuleStatementNotStatement
	OrStatement() Wafv2WebAclRuleStatementOrStatementOutputReference
	OrStatementInput() *Wafv2WebAclRuleStatementOrStatement
	RateBasedStatement() Wafv2WebAclRuleStatementRateBasedStatementOutputReference
	RateBasedStatementInput() *Wafv2WebAclRuleStatementRateBasedStatement
	RegexMatchStatement() Wafv2WebAclRuleStatementRegexMatchStatementOutputReference
	RegexMatchStatementInput() *Wafv2WebAclRuleStatementRegexMatchStatement
	RegexPatternSetReferenceStatement() Wafv2WebAclRuleStatementRegexPatternSetReferenceStatementOutputReference
	RegexPatternSetReferenceStatementInput() *Wafv2WebAclRuleStatementRegexPatternSetReferenceStatement
	RuleGroupReferenceStatement() Wafv2WebAclRuleStatementRuleGroupReferenceStatementOutputReference
	RuleGroupReferenceStatementInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatement
	SizeConstraintStatement() Wafv2WebAclRuleStatementSizeConstraintStatementOutputReference
	SizeConstraintStatementInput() *Wafv2WebAclRuleStatementSizeConstraintStatement
	SqliMatchStatement() Wafv2WebAclRuleStatementSqliMatchStatementOutputReference
	SqliMatchStatementInput() *Wafv2WebAclRuleStatementSqliMatchStatement
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	XssMatchStatement() Wafv2WebAclRuleStatementXssMatchStatementOutputReference
	XssMatchStatementInput() *Wafv2WebAclRuleStatementXssMatchStatement
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
	PutAndStatement(value *Wafv2WebAclRuleStatementAndStatement)
	PutByteMatchStatement(value *Wafv2WebAclRuleStatementByteMatchStatement)
	PutGeoMatchStatement(value *Wafv2WebAclRuleStatementGeoMatchStatement)
	PutIpSetReferenceStatement(value *Wafv2WebAclRuleStatementIpSetReferenceStatement)
	PutLabelMatchStatement(value *Wafv2WebAclRuleStatementLabelMatchStatement)
	PutManagedRuleGroupStatement(value *Wafv2WebAclRuleStatementManagedRuleGroupStatement)
	PutNotStatement(value *Wafv2WebAclRuleStatementNotStatement)
	PutOrStatement(value *Wafv2WebAclRuleStatementOrStatement)
	PutRateBasedStatement(value *Wafv2WebAclRuleStatementRateBasedStatement)
	PutRegexMatchStatement(value *Wafv2WebAclRuleStatementRegexMatchStatement)
	PutRegexPatternSetReferenceStatement(value *Wafv2WebAclRuleStatementRegexPatternSetReferenceStatement)
	PutRuleGroupReferenceStatement(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatement)
	PutSizeConstraintStatement(value *Wafv2WebAclRuleStatementSizeConstraintStatement)
	PutSqliMatchStatement(value *Wafv2WebAclRuleStatementSqliMatchStatement)
	PutXssMatchStatement(value *Wafv2WebAclRuleStatementXssMatchStatement)
	ResetAndStatement()
	ResetByteMatchStatement()
	ResetGeoMatchStatement()
	ResetIpSetReferenceStatement()
	ResetLabelMatchStatement()
	ResetManagedRuleGroupStatement()
	ResetNotStatement()
	ResetOrStatement()
	ResetRateBasedStatement()
	ResetRegexMatchStatement()
	ResetRegexPatternSetReferenceStatement()
	ResetRuleGroupReferenceStatement()
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

// The jsii proxy struct for Wafv2WebAclRuleStatementOutputReference
type jsiiProxy_Wafv2WebAclRuleStatementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) AndStatement() Wafv2WebAclRuleStatementAndStatementOutputReference {
	var returns Wafv2WebAclRuleStatementAndStatementOutputReference
	_jsii_.Get(
		j,
		"andStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) AndStatementInput() *Wafv2WebAclRuleStatementAndStatement {
	var returns *Wafv2WebAclRuleStatementAndStatement
	_jsii_.Get(
		j,
		"andStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ByteMatchStatement() Wafv2WebAclRuleStatementByteMatchStatementOutputReference {
	var returns Wafv2WebAclRuleStatementByteMatchStatementOutputReference
	_jsii_.Get(
		j,
		"byteMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ByteMatchStatementInput() *Wafv2WebAclRuleStatementByteMatchStatement {
	var returns *Wafv2WebAclRuleStatementByteMatchStatement
	_jsii_.Get(
		j,
		"byteMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) GeoMatchStatement() Wafv2WebAclRuleStatementGeoMatchStatementOutputReference {
	var returns Wafv2WebAclRuleStatementGeoMatchStatementOutputReference
	_jsii_.Get(
		j,
		"geoMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) GeoMatchStatementInput() *Wafv2WebAclRuleStatementGeoMatchStatement {
	var returns *Wafv2WebAclRuleStatementGeoMatchStatement
	_jsii_.Get(
		j,
		"geoMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) InternalValue() *Wafv2WebAclRuleStatement {
	var returns *Wafv2WebAclRuleStatement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) IpSetReferenceStatement() Wafv2WebAclRuleStatementIpSetReferenceStatementOutputReference {
	var returns Wafv2WebAclRuleStatementIpSetReferenceStatementOutputReference
	_jsii_.Get(
		j,
		"ipSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) IpSetReferenceStatementInput() *Wafv2WebAclRuleStatementIpSetReferenceStatement {
	var returns *Wafv2WebAclRuleStatementIpSetReferenceStatement
	_jsii_.Get(
		j,
		"ipSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) LabelMatchStatement() Wafv2WebAclRuleStatementLabelMatchStatementOutputReference {
	var returns Wafv2WebAclRuleStatementLabelMatchStatementOutputReference
	_jsii_.Get(
		j,
		"labelMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) LabelMatchStatementInput() *Wafv2WebAclRuleStatementLabelMatchStatement {
	var returns *Wafv2WebAclRuleStatementLabelMatchStatement
	_jsii_.Get(
		j,
		"labelMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ManagedRuleGroupStatement() Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference {
	var returns Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference
	_jsii_.Get(
		j,
		"managedRuleGroupStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ManagedRuleGroupStatementInput() *Wafv2WebAclRuleStatementManagedRuleGroupStatement {
	var returns *Wafv2WebAclRuleStatementManagedRuleGroupStatement
	_jsii_.Get(
		j,
		"managedRuleGroupStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) NotStatement() Wafv2WebAclRuleStatementNotStatementOutputReference {
	var returns Wafv2WebAclRuleStatementNotStatementOutputReference
	_jsii_.Get(
		j,
		"notStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) NotStatementInput() *Wafv2WebAclRuleStatementNotStatement {
	var returns *Wafv2WebAclRuleStatementNotStatement
	_jsii_.Get(
		j,
		"notStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) OrStatement() Wafv2WebAclRuleStatementOrStatementOutputReference {
	var returns Wafv2WebAclRuleStatementOrStatementOutputReference
	_jsii_.Get(
		j,
		"orStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) OrStatementInput() *Wafv2WebAclRuleStatementOrStatement {
	var returns *Wafv2WebAclRuleStatementOrStatement
	_jsii_.Get(
		j,
		"orStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) RateBasedStatement() Wafv2WebAclRuleStatementRateBasedStatementOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementOutputReference
	_jsii_.Get(
		j,
		"rateBasedStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) RateBasedStatementInput() *Wafv2WebAclRuleStatementRateBasedStatement {
	var returns *Wafv2WebAclRuleStatementRateBasedStatement
	_jsii_.Get(
		j,
		"rateBasedStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) RegexMatchStatement() Wafv2WebAclRuleStatementRegexMatchStatementOutputReference {
	var returns Wafv2WebAclRuleStatementRegexMatchStatementOutputReference
	_jsii_.Get(
		j,
		"regexMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) RegexMatchStatementInput() *Wafv2WebAclRuleStatementRegexMatchStatement {
	var returns *Wafv2WebAclRuleStatementRegexMatchStatement
	_jsii_.Get(
		j,
		"regexMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) RegexPatternSetReferenceStatement() Wafv2WebAclRuleStatementRegexPatternSetReferenceStatementOutputReference {
	var returns Wafv2WebAclRuleStatementRegexPatternSetReferenceStatementOutputReference
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) RegexPatternSetReferenceStatementInput() *Wafv2WebAclRuleStatementRegexPatternSetReferenceStatement {
	var returns *Wafv2WebAclRuleStatementRegexPatternSetReferenceStatement
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) RuleGroupReferenceStatement() Wafv2WebAclRuleStatementRuleGroupReferenceStatementOutputReference {
	var returns Wafv2WebAclRuleStatementRuleGroupReferenceStatementOutputReference
	_jsii_.Get(
		j,
		"ruleGroupReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) RuleGroupReferenceStatementInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatement {
	var returns *Wafv2WebAclRuleStatementRuleGroupReferenceStatement
	_jsii_.Get(
		j,
		"ruleGroupReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) SizeConstraintStatement() Wafv2WebAclRuleStatementSizeConstraintStatementOutputReference {
	var returns Wafv2WebAclRuleStatementSizeConstraintStatementOutputReference
	_jsii_.Get(
		j,
		"sizeConstraintStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) SizeConstraintStatementInput() *Wafv2WebAclRuleStatementSizeConstraintStatement {
	var returns *Wafv2WebAclRuleStatementSizeConstraintStatement
	_jsii_.Get(
		j,
		"sizeConstraintStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) SqliMatchStatement() Wafv2WebAclRuleStatementSqliMatchStatementOutputReference {
	var returns Wafv2WebAclRuleStatementSqliMatchStatementOutputReference
	_jsii_.Get(
		j,
		"sqliMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) SqliMatchStatementInput() *Wafv2WebAclRuleStatementSqliMatchStatement {
	var returns *Wafv2WebAclRuleStatementSqliMatchStatement
	_jsii_.Get(
		j,
		"sqliMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) XssMatchStatement() Wafv2WebAclRuleStatementXssMatchStatementOutputReference {
	var returns Wafv2WebAclRuleStatementXssMatchStatementOutputReference
	_jsii_.Get(
		j,
		"xssMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) XssMatchStatementInput() *Wafv2WebAclRuleStatementXssMatchStatement {
	var returns *Wafv2WebAclRuleStatementXssMatchStatement
	_jsii_.Get(
		j,
		"xssMatchStatementInput",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleStatementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2WebAclRuleStatementOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleStatementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleStatementOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleStatementOutputReference_Override(w Wafv2WebAclRuleStatementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference)SetInternalValue(val *Wafv2WebAclRuleStatement) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutAndStatement(value *Wafv2WebAclRuleStatementAndStatement) {
	if err := w.validatePutAndStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAndStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutByteMatchStatement(value *Wafv2WebAclRuleStatementByteMatchStatement) {
	if err := w.validatePutByteMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putByteMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutGeoMatchStatement(value *Wafv2WebAclRuleStatementGeoMatchStatement) {
	if err := w.validatePutGeoMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putGeoMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutIpSetReferenceStatement(value *Wafv2WebAclRuleStatementIpSetReferenceStatement) {
	if err := w.validatePutIpSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIpSetReferenceStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutLabelMatchStatement(value *Wafv2WebAclRuleStatementLabelMatchStatement) {
	if err := w.validatePutLabelMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLabelMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutManagedRuleGroupStatement(value *Wafv2WebAclRuleStatementManagedRuleGroupStatement) {
	if err := w.validatePutManagedRuleGroupStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putManagedRuleGroupStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutNotStatement(value *Wafv2WebAclRuleStatementNotStatement) {
	if err := w.validatePutNotStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putNotStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutOrStatement(value *Wafv2WebAclRuleStatementOrStatement) {
	if err := w.validatePutOrStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putOrStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutRateBasedStatement(value *Wafv2WebAclRuleStatementRateBasedStatement) {
	if err := w.validatePutRateBasedStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRateBasedStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutRegexMatchStatement(value *Wafv2WebAclRuleStatementRegexMatchStatement) {
	if err := w.validatePutRegexMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRegexMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutRegexPatternSetReferenceStatement(value *Wafv2WebAclRuleStatementRegexPatternSetReferenceStatement) {
	if err := w.validatePutRegexPatternSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRegexPatternSetReferenceStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutRuleGroupReferenceStatement(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatement) {
	if err := w.validatePutRuleGroupReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRuleGroupReferenceStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutSizeConstraintStatement(value *Wafv2WebAclRuleStatementSizeConstraintStatement) {
	if err := w.validatePutSizeConstraintStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSizeConstraintStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutSqliMatchStatement(value *Wafv2WebAclRuleStatementSqliMatchStatement) {
	if err := w.validatePutSqliMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSqliMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) PutXssMatchStatement(value *Wafv2WebAclRuleStatementXssMatchStatement) {
	if err := w.validatePutXssMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putXssMatchStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetAndStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetAndStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetByteMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetByteMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetGeoMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetGeoMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetIpSetReferenceStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetIpSetReferenceStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetLabelMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetLabelMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetManagedRuleGroupStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetManagedRuleGroupStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetNotStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetNotStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetOrStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetOrStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetRateBasedStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetRateBasedStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetRegexMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetRegexMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetRegexPatternSetReferenceStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetRegexPatternSetReferenceStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetRuleGroupReferenceStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetRuleGroupReferenceStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetSizeConstraintStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetSizeConstraintStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetSqliMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetSqliMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ResetXssMatchStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetXssMatchStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

