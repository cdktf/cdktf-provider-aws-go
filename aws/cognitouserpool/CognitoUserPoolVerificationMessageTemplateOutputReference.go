// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/cognitouserpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoUserPoolVerificationMessageTemplateOutputReference interface {
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
	DefaultEmailOption() *string
	SetDefaultEmailOption(val *string)
	DefaultEmailOptionInput() *string
	EmailMessage() *string
	SetEmailMessage(val *string)
	EmailMessageByLink() *string
	SetEmailMessageByLink(val *string)
	EmailMessageByLinkInput() *string
	EmailMessageInput() *string
	EmailSubject() *string
	SetEmailSubject(val *string)
	EmailSubjectByLink() *string
	SetEmailSubjectByLink(val *string)
	EmailSubjectByLinkInput() *string
	EmailSubjectInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CognitoUserPoolVerificationMessageTemplate
	SetInternalValue(val *CognitoUserPoolVerificationMessageTemplate)
	SmsMessage() *string
	SetSmsMessage(val *string)
	SmsMessageInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetDefaultEmailOption()
	ResetEmailMessage()
	ResetEmailMessageByLink()
	ResetEmailSubject()
	ResetEmailSubjectByLink()
	ResetSmsMessage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoUserPoolVerificationMessageTemplateOutputReference
type jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) DefaultEmailOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultEmailOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) DefaultEmailOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultEmailOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailMessageByLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessageByLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailMessageByLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessageByLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailSubjectByLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectByLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailSubjectByLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectByLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) InternalValue() *CognitoUserPoolVerificationMessageTemplate {
	var returns *CognitoUserPoolVerificationMessageTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) SmsMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) SmsMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCognitoUserPoolVerificationMessageTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CognitoUserPoolVerificationMessageTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoUserPoolVerificationMessageTemplateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPoolVerificationMessageTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCognitoUserPoolVerificationMessageTemplateOutputReference_Override(c CognitoUserPoolVerificationMessageTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPoolVerificationMessageTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference)SetDefaultEmailOption(val *string) {
	if err := j.validateSetDefaultEmailOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultEmailOption",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference)SetEmailMessage(val *string) {
	if err := j.validateSetEmailMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference)SetEmailMessageByLink(val *string) {
	if err := j.validateSetEmailMessageByLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailMessageByLink",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference)SetEmailSubject(val *string) {
	if err := j.validateSetEmailSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailSubject",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference)SetEmailSubjectByLink(val *string) {
	if err := j.validateSetEmailSubjectByLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailSubjectByLink",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference)SetInternalValue(val *CognitoUserPoolVerificationMessageTemplate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference)SetSmsMessage(val *string) {
	if err := j.validateSetSmsMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ResetDefaultEmailOption() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultEmailOption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ResetEmailMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailMessage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ResetEmailMessageByLink() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailMessageByLink",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ResetEmailSubject() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailSubject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ResetEmailSubjectByLink() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailSubjectByLink",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ResetSmsMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetSmsMessage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

