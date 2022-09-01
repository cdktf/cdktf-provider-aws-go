package appsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/appsync/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference interface {
	cdktf.ComplexObject
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LambdaAuthorizerConfig() AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference
	LambdaAuthorizerConfigInput() *AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfig
	OpenidConnectConfig() AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference
	OpenidConnectConfigInput() *AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserPoolConfig() AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference
	UserPoolConfigInput() *AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig
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
	PutLambdaAuthorizerConfig(value *AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfig)
	PutOpenidConnectConfig(value *AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfig)
	PutUserPoolConfig(value *AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig)
	ResetLambdaAuthorizerConfig()
	ResetOpenidConnectConfig()
	ResetUserPoolConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference
type jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) LambdaAuthorizerConfig() AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference {
	var returns AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) LambdaAuthorizerConfigInput() *AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfig {
	var returns *AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfig
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) OpenidConnectConfig() AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference {
	var returns AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference
	_jsii_.Get(
		j,
		"openidConnectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) OpenidConnectConfigInput() *AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfig {
	var returns *AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfig
	_jsii_.Get(
		j,
		"openidConnectConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) UserPoolConfig() AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference {
	var returns AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference
	_jsii_.Get(
		j,
		"userPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) UserPoolConfigInput() *AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig {
	var returns *AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig
	_jsii_.Get(
		j,
		"userPoolConfigInput",
		&returns,
	)
	return returns
}


func NewAppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference {
	_init_.Initialize()

	if err := validateNewAppsyncGraphqlApiAdditionalAuthenticationProviderOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appsync.AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference_Override(a AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appsync.AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) PutLambdaAuthorizerConfig(value *AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfig) {
	if err := a.validatePutLambdaAuthorizerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaAuthorizerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) PutOpenidConnectConfig(value *AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfig) {
	if err := a.validatePutOpenidConnectConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpenidConnectConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) PutUserPoolConfig(value *AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig) {
	if err := a.validatePutUserPoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserPoolConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) ResetLambdaAuthorizerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaAuthorizerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) ResetOpenidConnectConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOpenidConnectConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) ResetUserPoolConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetUserPoolConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

