// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alblistener

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/alblistener/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/alb_listener aws_alb_listener}.
type AlbListener interface {
	cdktf.TerraformResource
	AlpnPolicy() *string
	SetAlpnPolicy(val *string)
	AlpnPolicyInput() *string
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
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
	DefaultAction() AlbListenerDefaultActionList
	DefaultActionInput() interface{}
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancerArn() *string
	SetLoadBalancerArn(val *string)
	LoadBalancerArnInput() *string
	MutualAuthentication() AlbListenerMutualAuthenticationOutputReference
	MutualAuthenticationInput() *AlbListenerMutualAuthentication
	// The tree node.
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	RoutingHttpRequestXAmznMtlsClientcertHeaderName() *string
	SetRoutingHttpRequestXAmznMtlsClientcertHeaderName(val *string)
	RoutingHttpRequestXAmznMtlsClientcertHeaderNameInput() *string
	RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName() *string
	SetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName(val *string)
	RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderNameInput() *string
	RoutingHttpRequestXAmznMtlsClientcertLeafHeaderName() *string
	SetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName(val *string)
	RoutingHttpRequestXAmznMtlsClientcertLeafHeaderNameInput() *string
	RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName() *string
	SetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName(val *string)
	RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameInput() *string
	RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName() *string
	SetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName(val *string)
	RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderNameInput() *string
	RoutingHttpRequestXAmznMtlsClientcertValidityHeaderName() *string
	SetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName(val *string)
	RoutingHttpRequestXAmznMtlsClientcertValidityHeaderNameInput() *string
	RoutingHttpRequestXAmznTlsCipherSuiteHeaderName() *string
	SetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName(val *string)
	RoutingHttpRequestXAmznTlsCipherSuiteHeaderNameInput() *string
	RoutingHttpRequestXAmznTlsVersionHeaderName() *string
	SetRoutingHttpRequestXAmznTlsVersionHeaderName(val *string)
	RoutingHttpRequestXAmznTlsVersionHeaderNameInput() *string
	RoutingHttpResponseAccessControlAllowCredentialsHeaderValue() *string
	SetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue(val *string)
	RoutingHttpResponseAccessControlAllowCredentialsHeaderValueInput() *string
	RoutingHttpResponseAccessControlAllowHeadersHeaderValue() *string
	SetRoutingHttpResponseAccessControlAllowHeadersHeaderValue(val *string)
	RoutingHttpResponseAccessControlAllowHeadersHeaderValueInput() *string
	RoutingHttpResponseAccessControlAllowMethodsHeaderValue() *string
	SetRoutingHttpResponseAccessControlAllowMethodsHeaderValue(val *string)
	RoutingHttpResponseAccessControlAllowMethodsHeaderValueInput() *string
	RoutingHttpResponseAccessControlAllowOriginHeaderValue() *string
	SetRoutingHttpResponseAccessControlAllowOriginHeaderValue(val *string)
	RoutingHttpResponseAccessControlAllowOriginHeaderValueInput() *string
	RoutingHttpResponseAccessControlExposeHeadersHeaderValue() *string
	SetRoutingHttpResponseAccessControlExposeHeadersHeaderValue(val *string)
	RoutingHttpResponseAccessControlExposeHeadersHeaderValueInput() *string
	RoutingHttpResponseAccessControlMaxAgeHeaderValue() *string
	SetRoutingHttpResponseAccessControlMaxAgeHeaderValue(val *string)
	RoutingHttpResponseAccessControlMaxAgeHeaderValueInput() *string
	RoutingHttpResponseContentSecurityPolicyHeaderValue() *string
	SetRoutingHttpResponseContentSecurityPolicyHeaderValue(val *string)
	RoutingHttpResponseContentSecurityPolicyHeaderValueInput() *string
	RoutingHttpResponseServerEnabled() interface{}
	SetRoutingHttpResponseServerEnabled(val interface{})
	RoutingHttpResponseServerEnabledInput() interface{}
	RoutingHttpResponseStrictTransportSecurityHeaderValue() *string
	SetRoutingHttpResponseStrictTransportSecurityHeaderValue(val *string)
	RoutingHttpResponseStrictTransportSecurityHeaderValueInput() *string
	RoutingHttpResponseXContentTypeOptionsHeaderValue() *string
	SetRoutingHttpResponseXContentTypeOptionsHeaderValue(val *string)
	RoutingHttpResponseXContentTypeOptionsHeaderValueInput() *string
	RoutingHttpResponseXFrameOptionsHeaderValue() *string
	SetRoutingHttpResponseXFrameOptionsHeaderValue(val *string)
	RoutingHttpResponseXFrameOptionsHeaderValueInput() *string
	SslPolicy() *string
	SetSslPolicy(val *string)
	SslPolicyInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TcpIdleTimeoutSeconds() *float64
	SetTcpIdleTimeoutSeconds(val *float64)
	TcpIdleTimeoutSecondsInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AlbListenerTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutDefaultAction(value interface{})
	PutMutualAuthentication(value *AlbListenerMutualAuthentication)
	PutTimeouts(value *AlbListenerTimeouts)
	ResetAlpnPolicy()
	ResetCertificateArn()
	ResetId()
	ResetMutualAuthentication()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetProtocol()
	ResetRoutingHttpRequestXAmznMtlsClientcertHeaderName()
	ResetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName()
	ResetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName()
	ResetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName()
	ResetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName()
	ResetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName()
	ResetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName()
	ResetRoutingHttpRequestXAmznTlsVersionHeaderName()
	ResetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue()
	ResetRoutingHttpResponseAccessControlAllowHeadersHeaderValue()
	ResetRoutingHttpResponseAccessControlAllowMethodsHeaderValue()
	ResetRoutingHttpResponseAccessControlAllowOriginHeaderValue()
	ResetRoutingHttpResponseAccessControlExposeHeadersHeaderValue()
	ResetRoutingHttpResponseAccessControlMaxAgeHeaderValue()
	ResetRoutingHttpResponseContentSecurityPolicyHeaderValue()
	ResetRoutingHttpResponseServerEnabled()
	ResetRoutingHttpResponseStrictTransportSecurityHeaderValue()
	ResetRoutingHttpResponseXContentTypeOptionsHeaderValue()
	ResetRoutingHttpResponseXFrameOptionsHeaderValue()
	ResetSslPolicy()
	ResetTags()
	ResetTagsAll()
	ResetTcpIdleTimeoutSeconds()
	ResetTimeouts()
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

// The jsii proxy struct for AlbListener
type jsiiProxy_AlbListener struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AlbListener) AlpnPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alpnPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) AlpnPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alpnPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) DefaultAction() AlbListenerDefaultActionList {
	var returns AlbListenerDefaultActionList
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) DefaultActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) LoadBalancerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) MutualAuthentication() AlbListenerMutualAuthenticationOutputReference {
	var returns AlbListenerMutualAuthenticationOutputReference
	_jsii_.Get(
		j,
		"mutualAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) MutualAuthenticationInput() *AlbListenerMutualAuthentication {
	var returns *AlbListenerMutualAuthentication
	_jsii_.Get(
		j,
		"mutualAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznMtlsClientcertHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznMtlsClientcertHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertIssuerHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertIssuerHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznMtlsClientcertLeafHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertLeafHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznMtlsClientcertLeafHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertLeafHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSubjectHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSubjectHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznMtlsClientcertValidityHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertValidityHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznMtlsClientcertValidityHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertValidityHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznTlsCipherSuiteHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsCipherSuiteHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznTlsCipherSuiteHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsCipherSuiteHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznTlsVersionHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsVersionHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpRequestXAmznTlsVersionHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsVersionHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseAccessControlAllowCredentialsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowCredentialsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseAccessControlAllowCredentialsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowCredentialsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseAccessControlAllowHeadersHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowHeadersHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseAccessControlAllowHeadersHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowHeadersHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseAccessControlAllowMethodsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowMethodsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseAccessControlAllowMethodsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowMethodsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseAccessControlAllowOriginHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowOriginHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseAccessControlAllowOriginHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowOriginHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseAccessControlExposeHeadersHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlExposeHeadersHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseAccessControlExposeHeadersHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlExposeHeadersHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseAccessControlMaxAgeHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlMaxAgeHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseAccessControlMaxAgeHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlMaxAgeHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseContentSecurityPolicyHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseContentSecurityPolicyHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseContentSecurityPolicyHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseContentSecurityPolicyHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseServerEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingHttpResponseServerEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseServerEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingHttpResponseServerEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseStrictTransportSecurityHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseStrictTransportSecurityHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseStrictTransportSecurityHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseStrictTransportSecurityHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseXContentTypeOptionsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXContentTypeOptionsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseXContentTypeOptionsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXContentTypeOptionsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseXFrameOptionsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXFrameOptionsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) RoutingHttpResponseXFrameOptionsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXFrameOptionsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) SslPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) SslPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) TcpIdleTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpIdleTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) TcpIdleTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpIdleTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) Timeouts() AlbListenerTimeoutsOutputReference {
	var returns AlbListenerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListener) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/alb_listener aws_alb_listener} Resource.
func NewAlbListener(scope constructs.Construct, id *string, config *AlbListenerConfig) AlbListener {
	_init_.Initialize()

	if err := validateNewAlbListenerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlbListener{}

	_jsii_.Create(
		"@cdktf/provider-aws.albListener.AlbListener",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/alb_listener aws_alb_listener} Resource.
func NewAlbListener_Override(a AlbListener, scope constructs.Construct, id *string, config *AlbListenerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.albListener.AlbListener",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AlbListener)SetAlpnPolicy(val *string) {
	if err := j.validateSetAlpnPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alpnPolicy",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetLoadBalancerArn(val *string) {
	if err := j.validateSetLoadBalancerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerArn",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpRequestXAmznMtlsClientcertHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertHeaderName",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertIssuerHeaderName",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertLeafHeaderName",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertSubjectHeaderName",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertValidityHeaderName",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznTlsCipherSuiteHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznTlsCipherSuiteHeaderName",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpRequestXAmznTlsVersionHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznTlsVersionHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznTlsVersionHeaderName",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowCredentialsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowCredentialsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpResponseAccessControlAllowHeadersHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowHeadersHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowHeadersHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpResponseAccessControlAllowMethodsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowMethodsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowMethodsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpResponseAccessControlAllowOriginHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowOriginHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowOriginHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpResponseAccessControlExposeHeadersHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlExposeHeadersHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlExposeHeadersHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpResponseAccessControlMaxAgeHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlMaxAgeHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlMaxAgeHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpResponseContentSecurityPolicyHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseContentSecurityPolicyHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseContentSecurityPolicyHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpResponseServerEnabled(val interface{}) {
	if err := j.validateSetRoutingHttpResponseServerEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseServerEnabled",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpResponseStrictTransportSecurityHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseStrictTransportSecurityHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseStrictTransportSecurityHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpResponseXContentTypeOptionsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseXContentTypeOptionsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseXContentTypeOptionsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetRoutingHttpResponseXFrameOptionsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseXFrameOptionsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseXFrameOptionsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetSslPolicy(val *string) {
	if err := j.validateSetSslPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslPolicy",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AlbListener)SetTcpIdleTimeoutSeconds(val *float64) {
	if err := j.validateSetTcpIdleTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpIdleTimeoutSeconds",
		val,
	)
}

// Generates CDKTF code for importing a AlbListener resource upon running "cdktf plan <stack-name>".
func AlbListener_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAlbListener_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.albListener.AlbListener",
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
func AlbListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlbListener_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.albListener.AlbListener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AlbListener_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlbListener_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.albListener.AlbListener",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AlbListener_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlbListener_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.albListener.AlbListener",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AlbListener_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.albListener.AlbListener",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AlbListener) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AlbListener) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AlbListener) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlbListener) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlbListener) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlbListener) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlbListener) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlbListener) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlbListener) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlbListener) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlbListener) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlbListener) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListener) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AlbListener) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlbListener) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AlbListener) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AlbListener) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AlbListener) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AlbListener) PutDefaultAction(value interface{}) {
	if err := a.validatePutDefaultActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListener) PutMutualAuthentication(value *AlbListenerMutualAuthentication) {
	if err := a.validatePutMutualAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMutualAuthentication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListener) PutTimeouts(value *AlbListenerTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListener) ResetAlpnPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetAlpnPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetMutualAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetMutualAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpRequestXAmznMtlsClientcertHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznMtlsClientcertHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpRequestXAmznTlsVersionHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznTlsVersionHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpResponseAccessControlAllowHeadersHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseAccessControlAllowHeadersHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpResponseAccessControlAllowMethodsHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseAccessControlAllowMethodsHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpResponseAccessControlAllowOriginHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseAccessControlAllowOriginHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpResponseAccessControlExposeHeadersHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseAccessControlExposeHeadersHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpResponseAccessControlMaxAgeHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseAccessControlMaxAgeHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpResponseContentSecurityPolicyHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseContentSecurityPolicyHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpResponseServerEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseServerEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpResponseStrictTransportSecurityHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseStrictTransportSecurityHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpResponseXContentTypeOptionsHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseXContentTypeOptionsHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetRoutingHttpResponseXFrameOptionsHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseXFrameOptionsHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetSslPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSslPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetTcpIdleTimeoutSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTcpIdleTimeoutSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListener) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListener) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListener) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListener) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListener) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

