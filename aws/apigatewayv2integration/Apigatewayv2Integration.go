// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2integration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/apigatewayv2integration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/apigatewayv2_integration aws_apigatewayv2_integration}.
type Apigatewayv2Integration interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionId() *string
	SetConnectionId(val *string)
	ConnectionIdInput() *string
	ConnectionType() *string
	SetConnectionType(val *string)
	ConnectionTypeInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentHandlingStrategy() *string
	SetContentHandlingStrategy(val *string)
	ContentHandlingStrategyInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CredentialsArn() *string
	SetCredentialsArn(val *string)
	CredentialsArnInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IntegrationMethod() *string
	SetIntegrationMethod(val *string)
	IntegrationMethodInput() *string
	IntegrationResponseSelectionExpression() *string
	IntegrationSubtype() *string
	SetIntegrationSubtype(val *string)
	IntegrationSubtypeInput() *string
	IntegrationType() *string
	SetIntegrationType(val *string)
	IntegrationTypeInput() *string
	IntegrationUri() *string
	SetIntegrationUri(val *string)
	IntegrationUriInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PassthroughBehavior() *string
	SetPassthroughBehavior(val *string)
	PassthroughBehaviorInput() *string
	PayloadFormatVersion() *string
	SetPayloadFormatVersion(val *string)
	PayloadFormatVersionInput() *string
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
	RequestParameters() *map[string]*string
	SetRequestParameters(val *map[string]*string)
	RequestParametersInput() *map[string]*string
	RequestTemplates() *map[string]*string
	SetRequestTemplates(val *map[string]*string)
	RequestTemplatesInput() *map[string]*string
	ResponseParameters() Apigatewayv2IntegrationResponseParametersList
	ResponseParametersInput() interface{}
	TemplateSelectionExpression() *string
	SetTemplateSelectionExpression(val *string)
	TemplateSelectionExpressionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeoutMilliseconds() *float64
	SetTimeoutMilliseconds(val *float64)
	TimeoutMillisecondsInput() *float64
	TlsConfig() Apigatewayv2IntegrationTlsConfigOutputReference
	TlsConfigInput() *Apigatewayv2IntegrationTlsConfig
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutResponseParameters(value interface{})
	PutTlsConfig(value *Apigatewayv2IntegrationTlsConfig)
	ResetConnectionId()
	ResetConnectionType()
	ResetContentHandlingStrategy()
	ResetCredentialsArn()
	ResetDescription()
	ResetId()
	ResetIntegrationMethod()
	ResetIntegrationSubtype()
	ResetIntegrationUri()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassthroughBehavior()
	ResetPayloadFormatVersion()
	ResetRequestParameters()
	ResetRequestTemplates()
	ResetResponseParameters()
	ResetTemplateSelectionExpression()
	ResetTimeoutMilliseconds()
	ResetTlsConfig()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Apigatewayv2Integration
type jsiiProxy_Apigatewayv2Integration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2Integration) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ConnectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ContentHandlingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandlingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ContentHandlingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandlingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) CredentialsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) CredentialsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationResponseSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationResponseSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationSubtype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationSubtype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationSubtypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationSubtypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) PassthroughBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) PassthroughBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) PayloadFormatVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadFormatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) PayloadFormatVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadFormatVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) RequestParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) RequestParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) RequestTemplates() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) RequestTemplatesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestTemplatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ResponseParameters() Apigatewayv2IntegrationResponseParametersList {
	var returns Apigatewayv2IntegrationResponseParametersList
	_jsii_.Get(
		j,
		"responseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ResponseParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TemplateSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TemplateSelectionExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateSelectionExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TimeoutMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TimeoutMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TlsConfig() Apigatewayv2IntegrationTlsConfigOutputReference {
	var returns Apigatewayv2IntegrationTlsConfigOutputReference
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TlsConfigInput() *Apigatewayv2IntegrationTlsConfig {
	var returns *Apigatewayv2IntegrationTlsConfig
	_jsii_.Get(
		j,
		"tlsConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/apigatewayv2_integration aws_apigatewayv2_integration} Resource.
func NewApigatewayv2Integration(scope constructs.Construct, id *string, config *Apigatewayv2IntegrationConfig) Apigatewayv2Integration {
	_init_.Initialize()

	if err := validateNewApigatewayv2IntegrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Apigatewayv2Integration{}

	_jsii_.Create(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2Integration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/apigatewayv2_integration aws_apigatewayv2_integration} Resource.
func NewApigatewayv2Integration_Override(a Apigatewayv2Integration, scope constructs.Construct, id *string, config *Apigatewayv2IntegrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2Integration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetApiId(val *string) {
	if err := j.validateSetApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetConnectionId(val *string) {
	if err := j.validateSetConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetConnectionType(val *string) {
	if err := j.validateSetConnectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetContentHandlingStrategy(val *string) {
	if err := j.validateSetContentHandlingStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentHandlingStrategy",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetCredentialsArn(val *string) {
	if err := j.validateSetCredentialsArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialsArn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetIntegrationMethod(val *string) {
	if err := j.validateSetIntegrationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationMethod",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetIntegrationSubtype(val *string) {
	if err := j.validateSetIntegrationSubtypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationSubtype",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetIntegrationType(val *string) {
	if err := j.validateSetIntegrationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetIntegrationUri(val *string) {
	if err := j.validateSetIntegrationUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationUri",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetPassthroughBehavior(val *string) {
	if err := j.validateSetPassthroughBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passthroughBehavior",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetPayloadFormatVersion(val *string) {
	if err := j.validateSetPayloadFormatVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetRequestParameters(val *map[string]*string) {
	if err := j.validateSetRequestParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestParameters",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetRequestTemplates(val *map[string]*string) {
	if err := j.validateSetRequestTemplatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTemplates",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetTemplateSelectionExpression(val *string) {
	if err := j.validateSetTemplateSelectionExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateSelectionExpression",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration)SetTimeoutMilliseconds(val *float64) {
	if err := j.validateSetTimeoutMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutMilliseconds",
		val,
	)
}

// Generates CDKTF code for importing a Apigatewayv2Integration resource upon running "cdktf plan <stack-name>".
func Apigatewayv2Integration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApigatewayv2Integration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2Integration",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func Apigatewayv2Integration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayv2Integration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2Integration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Apigatewayv2Integration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayv2Integration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2Integration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Apigatewayv2Integration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayv2Integration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2Integration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2Integration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2Integration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_Apigatewayv2Integration) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_Apigatewayv2Integration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_Apigatewayv2Integration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_Apigatewayv2Integration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_Apigatewayv2Integration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_Apigatewayv2Integration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_Apigatewayv2Integration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_Apigatewayv2Integration) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_Apigatewayv2Integration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_Apigatewayv2Integration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Integration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_Apigatewayv2Integration) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) PutResponseParameters(value interface{}) {
	if err := a.validatePutResponseParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResponseParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) PutTlsConfig(value *Apigatewayv2IntegrationTlsConfig) {
	if err := a.validatePutTlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTlsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetConnectionId() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetConnectionType() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetContentHandlingStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetContentHandlingStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetCredentialsArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCredentialsArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetIntegrationMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetIntegrationSubtype() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationSubtype",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetIntegrationUri() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetPassthroughBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetPassthroughBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetPayloadFormatVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetPayloadFormatVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetRequestParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetRequestTemplates() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestTemplates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetResponseParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetTemplateSelectionExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetTemplateSelectionExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetTimeoutMilliseconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutMilliseconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetTlsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Integration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Integration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Integration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Integration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Integration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

