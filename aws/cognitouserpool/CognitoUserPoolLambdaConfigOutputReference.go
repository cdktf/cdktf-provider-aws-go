// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/cognitouserpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoUserPoolLambdaConfigOutputReference interface {
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
	CreateAuthChallenge() *string
	SetCreateAuthChallenge(val *string)
	CreateAuthChallengeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomEmailSender() CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference
	CustomEmailSenderInput() *CognitoUserPoolLambdaConfigCustomEmailSender
	CustomMessage() *string
	SetCustomMessage(val *string)
	CustomMessageInput() *string
	CustomSmsSender() CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference
	CustomSmsSenderInput() *CognitoUserPoolLambdaConfigCustomSmsSender
	DefineAuthChallenge() *string
	SetDefineAuthChallenge(val *string)
	DefineAuthChallengeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CognitoUserPoolLambdaConfig
	SetInternalValue(val *CognitoUserPoolLambdaConfig)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	PostAuthentication() *string
	SetPostAuthentication(val *string)
	PostAuthenticationInput() *string
	PostConfirmation() *string
	SetPostConfirmation(val *string)
	PostConfirmationInput() *string
	PreAuthentication() *string
	SetPreAuthentication(val *string)
	PreAuthenticationInput() *string
	PreSignUp() *string
	SetPreSignUp(val *string)
	PreSignUpInput() *string
	PreTokenGeneration() *string
	SetPreTokenGeneration(val *string)
	PreTokenGenerationInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserMigration() *string
	SetUserMigration(val *string)
	UserMigrationInput() *string
	VerifyAuthChallengeResponse() *string
	SetVerifyAuthChallengeResponse(val *string)
	VerifyAuthChallengeResponseInput() *string
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
	PutCustomEmailSender(value *CognitoUserPoolLambdaConfigCustomEmailSender)
	PutCustomSmsSender(value *CognitoUserPoolLambdaConfigCustomSmsSender)
	ResetCreateAuthChallenge()
	ResetCustomEmailSender()
	ResetCustomMessage()
	ResetCustomSmsSender()
	ResetDefineAuthChallenge()
	ResetKmsKeyId()
	ResetPostAuthentication()
	ResetPostConfirmation()
	ResetPreAuthentication()
	ResetPreSignUp()
	ResetPreTokenGeneration()
	ResetUserMigration()
	ResetVerifyAuthChallengeResponse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoUserPoolLambdaConfigOutputReference
type jsiiProxy_CognitoUserPoolLambdaConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CreateAuthChallenge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuthChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CreateAuthChallengeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuthChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CustomEmailSender() CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference {
	var returns CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference
	_jsii_.Get(
		j,
		"customEmailSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CustomEmailSenderInput() *CognitoUserPoolLambdaConfigCustomEmailSender {
	var returns *CognitoUserPoolLambdaConfigCustomEmailSender
	_jsii_.Get(
		j,
		"customEmailSenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CustomMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CustomMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CustomSmsSender() CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference {
	var returns CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference
	_jsii_.Get(
		j,
		"customSmsSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CustomSmsSenderInput() *CognitoUserPoolLambdaConfigCustomSmsSender {
	var returns *CognitoUserPoolLambdaConfigCustomSmsSender
	_jsii_.Get(
		j,
		"customSmsSenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) DefineAuthChallenge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defineAuthChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) DefineAuthChallengeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defineAuthChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) InternalValue() *CognitoUserPoolLambdaConfig {
	var returns *CognitoUserPoolLambdaConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PostAuthentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PostAuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PostConfirmation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postConfirmation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PostConfirmationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postConfirmationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PreAuthentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PreAuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PreSignUp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSignUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PreSignUpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSignUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PreTokenGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preTokenGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PreTokenGenerationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preTokenGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) UserMigration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userMigration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) UserMigrationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userMigrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) VerifyAuthChallengeResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyAuthChallengeResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) VerifyAuthChallengeResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyAuthChallengeResponseInput",
		&returns,
	)
	return returns
}


func NewCognitoUserPoolLambdaConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CognitoUserPoolLambdaConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoUserPoolLambdaConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUserPoolLambdaConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPoolLambdaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCognitoUserPoolLambdaConfigOutputReference_Override(c CognitoUserPoolLambdaConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPoolLambdaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetCreateAuthChallenge(val *string) {
	if err := j.validateSetCreateAuthChallengeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAuthChallenge",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetCustomMessage(val *string) {
	if err := j.validateSetCustomMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetDefineAuthChallenge(val *string) {
	if err := j.validateSetDefineAuthChallengeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defineAuthChallenge",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetInternalValue(val *CognitoUserPoolLambdaConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetPostAuthentication(val *string) {
	if err := j.validateSetPostAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postAuthentication",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetPostConfirmation(val *string) {
	if err := j.validateSetPostConfirmationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postConfirmation",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetPreAuthentication(val *string) {
	if err := j.validateSetPreAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preAuthentication",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetPreSignUp(val *string) {
	if err := j.validateSetPreSignUpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preSignUp",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetPreTokenGeneration(val *string) {
	if err := j.validateSetPreTokenGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preTokenGeneration",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetUserMigration(val *string) {
	if err := j.validateSetUserMigrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userMigration",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference)SetVerifyAuthChallengeResponse(val *string) {
	if err := j.validateSetVerifyAuthChallengeResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyAuthChallengeResponse",
		val,
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PutCustomEmailSender(value *CognitoUserPoolLambdaConfigCustomEmailSender) {
	if err := c.validatePutCustomEmailSenderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomEmailSender",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PutCustomSmsSender(value *CognitoUserPoolLambdaConfigCustomSmsSender) {
	if err := c.validatePutCustomSmsSenderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomSmsSender",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetCreateAuthChallenge() {
	_jsii_.InvokeVoid(
		c,
		"resetCreateAuthChallenge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetCustomEmailSender() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomEmailSender",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetCustomMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomMessage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetCustomSmsSender() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomSmsSender",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetDefineAuthChallenge() {
	_jsii_.InvokeVoid(
		c,
		"resetDefineAuthChallenge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetPostAuthentication() {
	_jsii_.InvokeVoid(
		c,
		"resetPostAuthentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetPostConfirmation() {
	_jsii_.InvokeVoid(
		c,
		"resetPostConfirmation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetPreAuthentication() {
	_jsii_.InvokeVoid(
		c,
		"resetPreAuthentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetPreSignUp() {
	_jsii_.InvokeVoid(
		c,
		"resetPreSignUp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetPreTokenGeneration() {
	_jsii_.InvokeVoid(
		c,
		"resetPreTokenGeneration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetUserMigration() {
	_jsii_.InvokeVoid(
		c,
		"resetUserMigration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetVerifyAuthChallengeResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetVerifyAuthChallengeResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

