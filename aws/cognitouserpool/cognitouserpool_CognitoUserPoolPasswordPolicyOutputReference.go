package cognitouserpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/cognitouserpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoUserPoolPasswordPolicyOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *CognitoUserPoolPasswordPolicy
	SetInternalValue(val *CognitoUserPoolPasswordPolicy)
	MinimumLength() *float64
	SetMinimumLength(val *float64)
	MinimumLengthInput() *float64
	RequireLowercase() interface{}
	SetRequireLowercase(val interface{})
	RequireLowercaseInput() interface{}
	RequireNumbers() interface{}
	SetRequireNumbers(val interface{})
	RequireNumbersInput() interface{}
	RequireSymbols() interface{}
	SetRequireSymbols(val interface{})
	RequireSymbolsInput() interface{}
	RequireUppercase() interface{}
	SetRequireUppercase(val interface{})
	RequireUppercaseInput() interface{}
	TemporaryPasswordValidityDays() *float64
	SetTemporaryPasswordValidityDays(val *float64)
	TemporaryPasswordValidityDaysInput() *float64
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
	ResetMinimumLength()
	ResetRequireLowercase()
	ResetRequireNumbers()
	ResetRequireSymbols()
	ResetRequireUppercase()
	ResetTemporaryPasswordValidityDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoUserPoolPasswordPolicyOutputReference
type jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) InternalValue() *CognitoUserPoolPasswordPolicy {
	var returns *CognitoUserPoolPasswordPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) MinimumLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) MinimumLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireLowercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLowercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireLowercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLowercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireNumbers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireNumbersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireSymbols() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSymbols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireSymbolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSymbolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireUppercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireUppercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireUppercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireUppercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) TemporaryPasswordValidityDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temporaryPasswordValidityDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) TemporaryPasswordValidityDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temporaryPasswordValidityDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCognitoUserPoolPasswordPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CognitoUserPoolPasswordPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoUserPoolPasswordPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPoolPasswordPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCognitoUserPoolPasswordPolicyOutputReference_Override(c CognitoUserPoolPasswordPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPoolPasswordPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference)SetInternalValue(val *CognitoUserPoolPasswordPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference)SetMinimumLength(val *float64) {
	if err := j.validateSetMinimumLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumLength",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference)SetRequireLowercase(val interface{}) {
	if err := j.validateSetRequireLowercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireLowercase",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference)SetRequireNumbers(val interface{}) {
	if err := j.validateSetRequireNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireNumbers",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference)SetRequireSymbols(val interface{}) {
	if err := j.validateSetRequireSymbolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireSymbols",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference)SetRequireUppercase(val interface{}) {
	if err := j.validateSetRequireUppercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireUppercase",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference)SetTemporaryPasswordValidityDays(val *float64) {
	if err := j.validateSetTemporaryPasswordValidityDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temporaryPasswordValidityDays",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ResetMinimumLength() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimumLength",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ResetRequireLowercase() {
	_jsii_.InvokeVoid(
		c,
		"resetRequireLowercase",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ResetRequireNumbers() {
	_jsii_.InvokeVoid(
		c,
		"resetRequireNumbers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ResetRequireSymbols() {
	_jsii_.InvokeVoid(
		c,
		"resetRequireSymbols",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ResetRequireUppercase() {
	_jsii_.InvokeVoid(
		c,
		"resetRequireUppercase",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ResetTemporaryPasswordValidityDays() {
	_jsii_.InvokeVoid(
		c,
		"resetTemporaryPasswordValidityDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

