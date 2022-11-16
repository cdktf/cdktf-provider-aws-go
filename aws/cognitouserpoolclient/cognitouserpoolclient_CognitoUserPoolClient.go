package cognitouserpoolclient

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/cognitouserpoolclient/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client aws_cognito_user_pool_client}.
type CognitoUserPoolClient interface {
	cdktf.TerraformResource
	AccessTokenValidity() *float64
	SetAccessTokenValidity(val *float64)
	AccessTokenValidityInput() *float64
	AllowedOauthFlows() *[]*string
	SetAllowedOauthFlows(val *[]*string)
	AllowedOauthFlowsInput() *[]*string
	AllowedOauthFlowsUserPoolClient() interface{}
	SetAllowedOauthFlowsUserPoolClient(val interface{})
	AllowedOauthFlowsUserPoolClientInput() interface{}
	AllowedOauthScopes() *[]*string
	SetAllowedOauthScopes(val *[]*string)
	AllowedOauthScopesInput() *[]*string
	AnalyticsConfiguration() CognitoUserPoolClientAnalyticsConfigurationOutputReference
	AnalyticsConfigurationInput() *CognitoUserPoolClientAnalyticsConfiguration
	AuthSessionValidity() *float64
	SetAuthSessionValidity(val *float64)
	AuthSessionValidityInput() *float64
	CallbackUrls() *[]*string
	SetCallbackUrls(val *[]*string)
	CallbackUrlsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientSecret() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DefaultRedirectUri() *string
	SetDefaultRedirectUri(val *string)
	DefaultRedirectUriInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnablePropagateAdditionalUserContextData() interface{}
	SetEnablePropagateAdditionalUserContextData(val interface{})
	EnablePropagateAdditionalUserContextDataInput() interface{}
	EnableTokenRevocation() interface{}
	SetEnableTokenRevocation(val interface{})
	EnableTokenRevocationInput() interface{}
	ExplicitAuthFlows() *[]*string
	SetExplicitAuthFlows(val *[]*string)
	ExplicitAuthFlowsInput() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GenerateSecret() interface{}
	SetGenerateSecret(val interface{})
	GenerateSecretInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	IdTokenValidity() *float64
	SetIdTokenValidity(val *float64)
	IdTokenValidityInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogoutUrls() *[]*string
	SetLogoutUrls(val *[]*string)
	LogoutUrlsInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PreventUserExistenceErrors() *string
	SetPreventUserExistenceErrors(val *string)
	PreventUserExistenceErrorsInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReadAttributes() *[]*string
	SetReadAttributes(val *[]*string)
	ReadAttributesInput() *[]*string
	RefreshTokenValidity() *float64
	SetRefreshTokenValidity(val *float64)
	RefreshTokenValidityInput() *float64
	SupportedIdentityProviders() *[]*string
	SetSupportedIdentityProviders(val *[]*string)
	SupportedIdentityProvidersInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenValidityUnits() CognitoUserPoolClientTokenValidityUnitsOutputReference
	TokenValidityUnitsInput() *CognitoUserPoolClientTokenValidityUnits
	UserPoolId() *string
	SetUserPoolId(val *string)
	UserPoolIdInput() *string
	WriteAttributes() *[]*string
	SetWriteAttributes(val *[]*string)
	WriteAttributesInput() *[]*string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAnalyticsConfiguration(value *CognitoUserPoolClientAnalyticsConfiguration)
	PutTokenValidityUnits(value *CognitoUserPoolClientTokenValidityUnits)
	ResetAccessTokenValidity()
	ResetAllowedOauthFlows()
	ResetAllowedOauthFlowsUserPoolClient()
	ResetAllowedOauthScopes()
	ResetAnalyticsConfiguration()
	ResetAuthSessionValidity()
	ResetCallbackUrls()
	ResetDefaultRedirectUri()
	ResetEnablePropagateAdditionalUserContextData()
	ResetEnableTokenRevocation()
	ResetExplicitAuthFlows()
	ResetGenerateSecret()
	ResetId()
	ResetIdTokenValidity()
	ResetLogoutUrls()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreventUserExistenceErrors()
	ResetReadAttributes()
	ResetRefreshTokenValidity()
	ResetSupportedIdentityProviders()
	ResetTokenValidityUnits()
	ResetWriteAttributes()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CognitoUserPoolClient
type jsiiProxy_CognitoUserPoolClient struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoUserPoolClient) AccessTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AccessTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AllowedOauthFlows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AllowedOauthFlowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthFlowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AllowedOauthFlowsUserPoolClient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedOauthFlowsUserPoolClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AllowedOauthFlowsUserPoolClientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedOauthFlowsUserPoolClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AllowedOauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AllowedOauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AnalyticsConfiguration() CognitoUserPoolClientAnalyticsConfigurationOutputReference {
	var returns CognitoUserPoolClientAnalyticsConfigurationOutputReference
	_jsii_.Get(
		j,
		"analyticsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AnalyticsConfigurationInput() *CognitoUserPoolClientAnalyticsConfiguration {
	var returns *CognitoUserPoolClientAnalyticsConfiguration
	_jsii_.Get(
		j,
		"analyticsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AuthSessionValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authSessionValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AuthSessionValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authSessionValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) CallbackUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callbackUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) CallbackUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callbackUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) DefaultRedirectUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRedirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) DefaultRedirectUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRedirectUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) EnablePropagateAdditionalUserContextData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePropagateAdditionalUserContextData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) EnablePropagateAdditionalUserContextDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePropagateAdditionalUserContextDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) EnableTokenRevocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTokenRevocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) EnableTokenRevocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTokenRevocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) ExplicitAuthFlows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"explicitAuthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) ExplicitAuthFlowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"explicitAuthFlowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) GenerateSecret() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) GenerateSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) IdTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) IdTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) LogoutUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logoutUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) LogoutUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logoutUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) PreventUserExistenceErrors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preventUserExistenceErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) PreventUserExistenceErrorsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preventUserExistenceErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) ReadAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) ReadAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) RefreshTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) RefreshTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) SupportedIdentityProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedIdentityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) SupportedIdentityProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedIdentityProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) TokenValidityUnits() CognitoUserPoolClientTokenValidityUnitsOutputReference {
	var returns CognitoUserPoolClientTokenValidityUnitsOutputReference
	_jsii_.Get(
		j,
		"tokenValidityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) TokenValidityUnitsInput() *CognitoUserPoolClientTokenValidityUnits {
	var returns *CognitoUserPoolClientTokenValidityUnits
	_jsii_.Get(
		j,
		"tokenValidityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) WriteAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"writeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) WriteAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"writeAttributesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client aws_cognito_user_pool_client} Resource.
func NewCognitoUserPoolClient(scope constructs.Construct, id *string, config *CognitoUserPoolClientConfig) CognitoUserPoolClient {
	_init_.Initialize()

	if err := validateNewCognitoUserPoolClientParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUserPoolClient{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClient",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client aws_cognito_user_pool_client} Resource.
func NewCognitoUserPoolClient_Override(c CognitoUserPoolClient, scope constructs.Construct, id *string, config *CognitoUserPoolClientConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClient",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetAccessTokenValidity(val *float64) {
	if err := j.validateSetAccessTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetAllowedOauthFlows(val *[]*string) {
	if err := j.validateSetAllowedOauthFlowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOauthFlows",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetAllowedOauthFlowsUserPoolClient(val interface{}) {
	if err := j.validateSetAllowedOauthFlowsUserPoolClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOauthFlowsUserPoolClient",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetAllowedOauthScopes(val *[]*string) {
	if err := j.validateSetAllowedOauthScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOauthScopes",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetAuthSessionValidity(val *float64) {
	if err := j.validateSetAuthSessionValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authSessionValidity",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetCallbackUrls(val *[]*string) {
	if err := j.validateSetCallbackUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callbackUrls",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetDefaultRedirectUri(val *string) {
	if err := j.validateSetDefaultRedirectUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRedirectUri",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetEnablePropagateAdditionalUserContextData(val interface{}) {
	if err := j.validateSetEnablePropagateAdditionalUserContextDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePropagateAdditionalUserContextData",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetEnableTokenRevocation(val interface{}) {
	if err := j.validateSetEnableTokenRevocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTokenRevocation",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetExplicitAuthFlows(val *[]*string) {
	if err := j.validateSetExplicitAuthFlowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"explicitAuthFlows",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetGenerateSecret(val interface{}) {
	if err := j.validateSetGenerateSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generateSecret",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetIdTokenValidity(val *float64) {
	if err := j.validateSetIdTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetLogoutUrls(val *[]*string) {
	if err := j.validateSetLogoutUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoutUrls",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetPreventUserExistenceErrors(val *string) {
	if err := j.validateSetPreventUserExistenceErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventUserExistenceErrors",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetReadAttributes(val *[]*string) {
	if err := j.validateSetReadAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readAttributes",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetRefreshTokenValidity(val *float64) {
	if err := j.validateSetRefreshTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetSupportedIdentityProviders(val *[]*string) {
	if err := j.validateSetSupportedIdentityProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedIdentityProviders",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetUserPoolId(val *string) {
	if err := j.validateSetUserPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient)SetWriteAttributes(val *[]*string) {
	if err := j.validateSetWriteAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeAttributes",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CognitoUserPoolClient_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoUserPoolClient_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClient",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitoUserPoolClient_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoUserPoolClient_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClient",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitoUserPoolClient_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoUserPoolClient_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClient",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoUserPoolClient_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClient",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CognitoUserPoolClient) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoUserPoolClient) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolClient) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoUserPoolClient) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoUserPoolClient) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoUserPoolClient) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolClient) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolClient) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoUserPoolClient) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoUserPoolClient) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolClient) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) PutAnalyticsConfiguration(value *CognitoUserPoolClientAnalyticsConfiguration) {
	if err := c.validatePutAnalyticsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAnalyticsConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) PutTokenValidityUnits(value *CognitoUserPoolClientTokenValidityUnits) {
	if err := c.validatePutTokenValidityUnitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTokenValidityUnits",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetAccessTokenValidity() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessTokenValidity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetAllowedOauthFlows() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedOauthFlows",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetAllowedOauthFlowsUserPoolClient() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedOauthFlowsUserPoolClient",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetAllowedOauthScopes() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedOauthScopes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetAnalyticsConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetAnalyticsConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetAuthSessionValidity() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthSessionValidity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetCallbackUrls() {
	_jsii_.InvokeVoid(
		c,
		"resetCallbackUrls",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetDefaultRedirectUri() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultRedirectUri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetEnablePropagateAdditionalUserContextData() {
	_jsii_.InvokeVoid(
		c,
		"resetEnablePropagateAdditionalUserContextData",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetEnableTokenRevocation() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableTokenRevocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetExplicitAuthFlows() {
	_jsii_.InvokeVoid(
		c,
		"resetExplicitAuthFlows",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetGenerateSecret() {
	_jsii_.InvokeVoid(
		c,
		"resetGenerateSecret",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetIdTokenValidity() {
	_jsii_.InvokeVoid(
		c,
		"resetIdTokenValidity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetLogoutUrls() {
	_jsii_.InvokeVoid(
		c,
		"resetLogoutUrls",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetPreventUserExistenceErrors() {
	_jsii_.InvokeVoid(
		c,
		"resetPreventUserExistenceErrors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetReadAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetReadAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetRefreshTokenValidity() {
	_jsii_.InvokeVoid(
		c,
		"resetRefreshTokenValidity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetSupportedIdentityProviders() {
	_jsii_.InvokeVoid(
		c,
		"resetSupportedIdentityProviders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetTokenValidityUnits() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenValidityUnits",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetWriteAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetWriteAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolClient) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolClient) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolClient) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

