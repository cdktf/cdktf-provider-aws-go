package cognitouserpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/cognitouserpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoUserPoolEmailConfigurationOutputReference interface {
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
	ConfigurationSet() *string
	SetConfigurationSet(val *string)
	ConfigurationSetInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EmailSendingAccount() *string
	SetEmailSendingAccount(val *string)
	EmailSendingAccountInput() *string
	// Experimental.
	Fqn() *string
	FromEmailAddress() *string
	SetFromEmailAddress(val *string)
	FromEmailAddressInput() *string
	InternalValue() *CognitoUserPoolEmailConfiguration
	SetInternalValue(val *CognitoUserPoolEmailConfiguration)
	ReplyToEmailAddress() *string
	SetReplyToEmailAddress(val *string)
	ReplyToEmailAddressInput() *string
	SourceArn() *string
	SetSourceArn(val *string)
	SourceArnInput() *string
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
	ResetConfigurationSet()
	ResetEmailSendingAccount()
	ResetFromEmailAddress()
	ResetReplyToEmailAddress()
	ResetSourceArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoUserPoolEmailConfigurationOutputReference
type jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ConfigurationSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ConfigurationSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) EmailSendingAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSendingAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) EmailSendingAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSendingAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) FromEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) FromEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) InternalValue() *CognitoUserPoolEmailConfiguration {
	var returns *CognitoUserPoolEmailConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ReplyToEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ReplyToEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) SourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCognitoUserPoolEmailConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CognitoUserPoolEmailConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoUserPoolEmailConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPoolEmailConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCognitoUserPoolEmailConfigurationOutputReference_Override(c CognitoUserPoolEmailConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPoolEmailConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference)SetConfigurationSet(val *string) {
	if err := j.validateSetConfigurationSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationSet",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference)SetEmailSendingAccount(val *string) {
	if err := j.validateSetEmailSendingAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailSendingAccount",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference)SetFromEmailAddress(val *string) {
	if err := j.validateSetFromEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromEmailAddress",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference)SetInternalValue(val *CognitoUserPoolEmailConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference)SetReplyToEmailAddress(val *string) {
	if err := j.validateSetReplyToEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replyToEmailAddress",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference)SetSourceArn(val *string) {
	if err := j.validateSetSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceArn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ResetConfigurationSet() {
	_jsii_.InvokeVoid(
		c,
		"resetConfigurationSet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ResetEmailSendingAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailSendingAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ResetFromEmailAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetFromEmailAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ResetReplyToEmailAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetReplyToEmailAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ResetSourceArn() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

