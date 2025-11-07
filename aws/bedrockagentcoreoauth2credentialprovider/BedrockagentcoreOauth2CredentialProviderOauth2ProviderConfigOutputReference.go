// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/bedrockagentcoreoauth2credentialprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference interface {
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
	CustomOauth2ProviderConfig() BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigCustomOauth2ProviderConfigList
	CustomOauth2ProviderConfigInput() interface{}
	// Experimental.
	Fqn() *string
	GithubOauth2ProviderConfig() BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigGithubOauth2ProviderConfigList
	GithubOauth2ProviderConfigInput() interface{}
	GoogleOauth2ProviderConfig() BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigGoogleOauth2ProviderConfigList
	GoogleOauth2ProviderConfigInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MicrosoftOauth2ProviderConfig() BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigMicrosoftOauth2ProviderConfigList
	MicrosoftOauth2ProviderConfigInput() interface{}
	SalesforceOauth2ProviderConfig() BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigSalesforceOauth2ProviderConfigList
	SalesforceOauth2ProviderConfigInput() interface{}
	SlackOauth2ProviderConfig() BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigSlackOauth2ProviderConfigList
	SlackOauth2ProviderConfigInput() interface{}
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
	PutCustomOauth2ProviderConfig(value interface{})
	PutGithubOauth2ProviderConfig(value interface{})
	PutGoogleOauth2ProviderConfig(value interface{})
	PutMicrosoftOauth2ProviderConfig(value interface{})
	PutSalesforceOauth2ProviderConfig(value interface{})
	PutSlackOauth2ProviderConfig(value interface{})
	ResetCustomOauth2ProviderConfig()
	ResetGithubOauth2ProviderConfig()
	ResetGoogleOauth2ProviderConfig()
	ResetMicrosoftOauth2ProviderConfig()
	ResetSalesforceOauth2ProviderConfig()
	ResetSlackOauth2ProviderConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference
type jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) CustomOauth2ProviderConfig() BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigCustomOauth2ProviderConfigList {
	var returns BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigCustomOauth2ProviderConfigList
	_jsii_.Get(
		j,
		"customOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) CustomOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) GithubOauth2ProviderConfig() BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigGithubOauth2ProviderConfigList {
	var returns BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigGithubOauth2ProviderConfigList
	_jsii_.Get(
		j,
		"githubOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) GithubOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) GoogleOauth2ProviderConfig() BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigGoogleOauth2ProviderConfigList {
	var returns BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigGoogleOauth2ProviderConfigList
	_jsii_.Get(
		j,
		"googleOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) GoogleOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googleOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) MicrosoftOauth2ProviderConfig() BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigMicrosoftOauth2ProviderConfigList {
	var returns BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigMicrosoftOauth2ProviderConfigList
	_jsii_.Get(
		j,
		"microsoftOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) MicrosoftOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) SalesforceOauth2ProviderConfig() BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigSalesforceOauth2ProviderConfigList {
	var returns BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigSalesforceOauth2ProviderConfigList
	_jsii_.Get(
		j,
		"salesforceOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) SalesforceOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"salesforceOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) SlackOauth2ProviderConfig() BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigSlackOauth2ProviderConfigList {
	var returns BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigSlackOauth2ProviderConfigList
	_jsii_.Get(
		j,
		"slackOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) SlackOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slackOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentcoreOauth2CredentialProvider.BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference_Override(b BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentcoreOauth2CredentialProvider.BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) PutCustomOauth2ProviderConfig(value interface{}) {
	if err := b.validatePutCustomOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCustomOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) PutGithubOauth2ProviderConfig(value interface{}) {
	if err := b.validatePutGithubOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putGithubOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) PutGoogleOauth2ProviderConfig(value interface{}) {
	if err := b.validatePutGoogleOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putGoogleOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) PutMicrosoftOauth2ProviderConfig(value interface{}) {
	if err := b.validatePutMicrosoftOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMicrosoftOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) PutSalesforceOauth2ProviderConfig(value interface{}) {
	if err := b.validatePutSalesforceOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSalesforceOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) PutSlackOauth2ProviderConfig(value interface{}) {
	if err := b.validatePutSlackOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSlackOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) ResetCustomOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) ResetGithubOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetGithubOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) ResetGoogleOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetGoogleOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) ResetMicrosoftOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetMicrosoftOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) ResetSalesforceOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetSalesforceOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) ResetSlackOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetSlackOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

