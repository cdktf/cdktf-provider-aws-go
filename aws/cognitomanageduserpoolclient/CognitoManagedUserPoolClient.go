package cognitomanageduserpoolclient

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/cognitomanageduserpoolclient/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cognito_managed_user_pool_client aws_cognito_managed_user_pool_client}.
type CognitoManagedUserPoolClient interface {
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
	AnalyticsConfiguration() CognitoManagedUserPoolClientAnalyticsConfigurationList
	AnalyticsConfigurationInput() interface{}
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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
	Id() *string
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
	NamePattern() *string
	SetNamePattern(val *string)
	NamePatternInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
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
	TokenValidityUnits() CognitoManagedUserPoolClientTokenValidityUnitsList
	TokenValidityUnitsInput() interface{}
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
	PutAnalyticsConfiguration(value interface{})
	PutTokenValidityUnits(value interface{})
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
	ResetIdTokenValidity()
	ResetLogoutUrls()
	ResetNamePattern()
	ResetNamePrefix()
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

// The jsii proxy struct for CognitoManagedUserPoolClient
type jsiiProxy_CognitoManagedUserPoolClient struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) AccessTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) AccessTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) AllowedOauthFlows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) AllowedOauthFlowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthFlowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) AllowedOauthFlowsUserPoolClient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedOauthFlowsUserPoolClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) AllowedOauthFlowsUserPoolClientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedOauthFlowsUserPoolClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) AllowedOauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) AllowedOauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) AnalyticsConfiguration() CognitoManagedUserPoolClientAnalyticsConfigurationList {
	var returns CognitoManagedUserPoolClientAnalyticsConfigurationList
	_jsii_.Get(
		j,
		"analyticsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) AnalyticsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analyticsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) AuthSessionValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authSessionValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) AuthSessionValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authSessionValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) CallbackUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callbackUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) CallbackUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callbackUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) DefaultRedirectUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRedirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) DefaultRedirectUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRedirectUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) EnablePropagateAdditionalUserContextData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePropagateAdditionalUserContextData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) EnablePropagateAdditionalUserContextDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePropagateAdditionalUserContextDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) EnableTokenRevocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTokenRevocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) EnableTokenRevocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTokenRevocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) ExplicitAuthFlows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"explicitAuthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) ExplicitAuthFlowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"explicitAuthFlowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) IdTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) IdTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) LogoutUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logoutUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) LogoutUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logoutUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) NamePattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) NamePatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) PreventUserExistenceErrors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preventUserExistenceErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) PreventUserExistenceErrorsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preventUserExistenceErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) ReadAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) ReadAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) RefreshTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) RefreshTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) SupportedIdentityProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedIdentityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) SupportedIdentityProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedIdentityProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) TokenValidityUnits() CognitoManagedUserPoolClientTokenValidityUnitsList {
	var returns CognitoManagedUserPoolClientTokenValidityUnitsList
	_jsii_.Get(
		j,
		"tokenValidityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) TokenValidityUnitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenValidityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) WriteAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"writeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoManagedUserPoolClient) WriteAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"writeAttributesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cognito_managed_user_pool_client aws_cognito_managed_user_pool_client} Resource.
func NewCognitoManagedUserPoolClient(scope constructs.Construct, id *string, config *CognitoManagedUserPoolClientConfig) CognitoManagedUserPoolClient {
	_init_.Initialize()

	if err := validateNewCognitoManagedUserPoolClientParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoManagedUserPoolClient{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoManagedUserPoolClient.CognitoManagedUserPoolClient",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cognito_managed_user_pool_client aws_cognito_managed_user_pool_client} Resource.
func NewCognitoManagedUserPoolClient_Override(c CognitoManagedUserPoolClient, scope constructs.Construct, id *string, config *CognitoManagedUserPoolClientConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoManagedUserPoolClient.CognitoManagedUserPoolClient",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetAccessTokenValidity(val *float64) {
	if err := j.validateSetAccessTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetAllowedOauthFlows(val *[]*string) {
	if err := j.validateSetAllowedOauthFlowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOauthFlows",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetAllowedOauthFlowsUserPoolClient(val interface{}) {
	if err := j.validateSetAllowedOauthFlowsUserPoolClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOauthFlowsUserPoolClient",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetAllowedOauthScopes(val *[]*string) {
	if err := j.validateSetAllowedOauthScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOauthScopes",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetAuthSessionValidity(val *float64) {
	if err := j.validateSetAuthSessionValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authSessionValidity",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetCallbackUrls(val *[]*string) {
	if err := j.validateSetCallbackUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callbackUrls",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetDefaultRedirectUri(val *string) {
	if err := j.validateSetDefaultRedirectUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRedirectUri",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetEnablePropagateAdditionalUserContextData(val interface{}) {
	if err := j.validateSetEnablePropagateAdditionalUserContextDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePropagateAdditionalUserContextData",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetEnableTokenRevocation(val interface{}) {
	if err := j.validateSetEnableTokenRevocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTokenRevocation",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetExplicitAuthFlows(val *[]*string) {
	if err := j.validateSetExplicitAuthFlowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"explicitAuthFlows",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetIdTokenValidity(val *float64) {
	if err := j.validateSetIdTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetLogoutUrls(val *[]*string) {
	if err := j.validateSetLogoutUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoutUrls",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetNamePattern(val *string) {
	if err := j.validateSetNamePatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePattern",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetPreventUserExistenceErrors(val *string) {
	if err := j.validateSetPreventUserExistenceErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventUserExistenceErrors",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetReadAttributes(val *[]*string) {
	if err := j.validateSetReadAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readAttributes",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetRefreshTokenValidity(val *float64) {
	if err := j.validateSetRefreshTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetSupportedIdentityProviders(val *[]*string) {
	if err := j.validateSetSupportedIdentityProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedIdentityProviders",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetUserPoolId(val *string) {
	if err := j.validateSetUserPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

func (j *jsiiProxy_CognitoManagedUserPoolClient)SetWriteAttributes(val *[]*string) {
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
func CognitoManagedUserPoolClient_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoManagedUserPoolClient_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cognitoManagedUserPoolClient.CognitoManagedUserPoolClient",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitoManagedUserPoolClient_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoManagedUserPoolClient_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cognitoManagedUserPoolClient.CognitoManagedUserPoolClient",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitoManagedUserPoolClient_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoManagedUserPoolClient_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cognitoManagedUserPoolClient.CognitoManagedUserPoolClient",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoManagedUserPoolClient_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.cognitoManagedUserPoolClient.CognitoManagedUserPoolClient",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoManagedUserPoolClient) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoManagedUserPoolClient) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoManagedUserPoolClient) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoManagedUserPoolClient) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoManagedUserPoolClient) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoManagedUserPoolClient) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoManagedUserPoolClient) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoManagedUserPoolClient) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoManagedUserPoolClient) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoManagedUserPoolClient) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) PutAnalyticsConfiguration(value interface{}) {
	if err := c.validatePutAnalyticsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAnalyticsConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) PutTokenValidityUnits(value interface{}) {
	if err := c.validatePutTokenValidityUnitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTokenValidityUnits",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetAccessTokenValidity() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessTokenValidity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetAllowedOauthFlows() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedOauthFlows",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetAllowedOauthFlowsUserPoolClient() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedOauthFlowsUserPoolClient",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetAllowedOauthScopes() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedOauthScopes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetAnalyticsConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetAnalyticsConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetAuthSessionValidity() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthSessionValidity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetCallbackUrls() {
	_jsii_.InvokeVoid(
		c,
		"resetCallbackUrls",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetDefaultRedirectUri() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultRedirectUri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetEnablePropagateAdditionalUserContextData() {
	_jsii_.InvokeVoid(
		c,
		"resetEnablePropagateAdditionalUserContextData",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetEnableTokenRevocation() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableTokenRevocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetExplicitAuthFlows() {
	_jsii_.InvokeVoid(
		c,
		"resetExplicitAuthFlows",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetIdTokenValidity() {
	_jsii_.InvokeVoid(
		c,
		"resetIdTokenValidity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetLogoutUrls() {
	_jsii_.InvokeVoid(
		c,
		"resetLogoutUrls",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetNamePattern() {
	_jsii_.InvokeVoid(
		c,
		"resetNamePattern",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetPreventUserExistenceErrors() {
	_jsii_.InvokeVoid(
		c,
		"resetPreventUserExistenceErrors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetReadAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetReadAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetRefreshTokenValidity() {
	_jsii_.InvokeVoid(
		c,
		"resetRefreshTokenValidity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetSupportedIdentityProviders() {
	_jsii_.InvokeVoid(
		c,
		"resetSupportedIdentityProviders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetTokenValidityUnits() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenValidityUnits",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ResetWriteAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetWriteAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoManagedUserPoolClient) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

