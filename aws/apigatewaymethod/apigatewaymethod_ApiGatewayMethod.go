package apigatewaymethod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/apigatewaymethod/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method aws_api_gateway_method}.
type ApiGatewayMethod interface {
	cdktf.TerraformResource
	ApiKeyRequired() interface{}
	SetApiKeyRequired(val interface{})
	ApiKeyRequiredInput() interface{}
	Authorization() *string
	SetAuthorization(val *string)
	AuthorizationInput() *string
	AuthorizationScopes() *[]*string
	SetAuthorizationScopes(val *[]*string)
	AuthorizationScopesInput() *[]*string
	AuthorizerId() *string
	SetAuthorizerId(val *string)
	AuthorizerIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OperationName() *string
	SetOperationName(val *string)
	OperationNameInput() *string
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
	RequestModels() *map[string]*string
	SetRequestModels(val *map[string]*string)
	RequestModelsInput() *map[string]*string
	RequestParameters() *map[string]interface{}
	SetRequestParameters(val *map[string]interface{})
	RequestParametersInput() *map[string]interface{}
	RequestValidatorId() *string
	SetRequestValidatorId(val *string)
	RequestValidatorIdInput() *string
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetApiKeyRequired()
	ResetAuthorizationScopes()
	ResetAuthorizerId()
	ResetId()
	ResetOperationName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestModels()
	ResetRequestParameters()
	ResetRequestValidatorId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayMethod
type jsiiProxy_ApiGatewayMethod struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayMethod) ApiKeyRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) ApiKeyRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Authorization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) AuthorizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) AuthorizationScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizationScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) AuthorizationScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizationScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) AuthorizerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) OperationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RequestModels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RequestModelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestModelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RequestParameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"requestParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RequestParametersInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"requestParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RequestValidatorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestValidatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RequestValidatorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestValidatorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method aws_api_gateway_method} Resource.
func NewApiGatewayMethod(scope constructs.Construct, id *string, config *ApiGatewayMethodConfig) ApiGatewayMethod {
	_init_.Initialize()

	if err := validateNewApiGatewayMethodParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayMethod{}

	_jsii_.Create(
		"@cdktf/provider-aws.apiGatewayMethod.ApiGatewayMethod",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method aws_api_gateway_method} Resource.
func NewApiGatewayMethod_Override(a ApiGatewayMethod, scope constructs.Construct, id *string, config *ApiGatewayMethodConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.apiGatewayMethod.ApiGatewayMethod",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetApiKeyRequired(val interface{}) {
	if err := j.validateSetApiKeyRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKeyRequired",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetAuthorization(val *string) {
	if err := j.validateSetAuthorizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorization",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetAuthorizationScopes(val *[]*string) {
	if err := j.validateSetAuthorizationScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationScopes",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetAuthorizerId(val *string) {
	if err := j.validateSetAuthorizerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetOperationName(val *string) {
	if err := j.validateSetOperationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetRequestModels(val *map[string]*string) {
	if err := j.validateSetRequestModelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestModels",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetRequestParameters(val *map[string]interface{}) {
	if err := j.validateSetRequestParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestParameters",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetRequestValidatorId(val *string) {
	if err := j.validateSetRequestValidatorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestValidatorId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod)SetRestApiId(val *string) {
	if err := j.validateSetRestApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restApiId",
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
func ApiGatewayMethod_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiGatewayMethod_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.apiGatewayMethod.ApiGatewayMethod",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiGatewayMethod_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiGatewayMethod_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.apiGatewayMethod.ApiGatewayMethod",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiGatewayMethod_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiGatewayMethod_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.apiGatewayMethod.ApiGatewayMethod",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayMethod_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.apiGatewayMethod.ApiGatewayMethod",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiGatewayMethod) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApiGatewayMethod) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiGatewayMethod) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiGatewayMethod) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiGatewayMethod) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiGatewayMethod) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiGatewayMethod) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiGatewayMethod) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiGatewayMethod) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiGatewayMethod) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiGatewayMethod) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethod) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetApiKeyRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKeyRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetAuthorizationScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetAuthorizerId() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetOperationName() {
	_jsii_.InvokeVoid(
		a,
		"resetOperationName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetRequestModels() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestModels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetRequestParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetRequestValidatorId() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestValidatorId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethod) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethod) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethod) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

