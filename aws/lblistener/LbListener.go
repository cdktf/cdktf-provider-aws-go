// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistener

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lblistener/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/lb_listener aws_lb_listener}.
type LbListener interface {
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
	DefaultAction() LbListenerDefaultActionList
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
	MutualAuthentication() LbListenerMutualAuthenticationOutputReference
	MutualAuthenticationInput() *LbListenerMutualAuthentication
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
	Timeouts() LbListenerTimeoutsOutputReference
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
	PutMutualAuthentication(value *LbListenerMutualAuthentication)
	PutTimeouts(value *LbListenerTimeouts)
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

// The jsii proxy struct for LbListener
type jsiiProxy_LbListener struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LbListener) AlpnPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alpnPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) AlpnPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alpnPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) DefaultAction() LbListenerDefaultActionList {
	var returns LbListenerDefaultActionList
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) DefaultActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) LoadBalancerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) MutualAuthentication() LbListenerMutualAuthenticationOutputReference {
	var returns LbListenerMutualAuthenticationOutputReference
	_jsii_.Get(
		j,
		"mutualAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) MutualAuthenticationInput() *LbListenerMutualAuthentication {
	var returns *LbListenerMutualAuthentication
	_jsii_.Get(
		j,
		"mutualAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznMtlsClientcertHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznMtlsClientcertHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertIssuerHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertIssuerHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznMtlsClientcertLeafHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertLeafHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznMtlsClientcertLeafHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertLeafHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSubjectHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSubjectHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznMtlsClientcertValidityHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertValidityHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznMtlsClientcertValidityHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertValidityHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznTlsCipherSuiteHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsCipherSuiteHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznTlsCipherSuiteHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsCipherSuiteHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznTlsVersionHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsVersionHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpRequestXAmznTlsVersionHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsVersionHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseAccessControlAllowCredentialsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowCredentialsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseAccessControlAllowCredentialsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowCredentialsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseAccessControlAllowHeadersHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowHeadersHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseAccessControlAllowHeadersHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowHeadersHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseAccessControlAllowMethodsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowMethodsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseAccessControlAllowMethodsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowMethodsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseAccessControlAllowOriginHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowOriginHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseAccessControlAllowOriginHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowOriginHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseAccessControlExposeHeadersHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlExposeHeadersHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseAccessControlExposeHeadersHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlExposeHeadersHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseAccessControlMaxAgeHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlMaxAgeHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseAccessControlMaxAgeHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlMaxAgeHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseContentSecurityPolicyHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseContentSecurityPolicyHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseContentSecurityPolicyHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseContentSecurityPolicyHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseServerEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingHttpResponseServerEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseServerEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingHttpResponseServerEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseStrictTransportSecurityHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseStrictTransportSecurityHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseStrictTransportSecurityHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseStrictTransportSecurityHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseXContentTypeOptionsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXContentTypeOptionsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseXContentTypeOptionsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXContentTypeOptionsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseXFrameOptionsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXFrameOptionsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) RoutingHttpResponseXFrameOptionsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXFrameOptionsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) SslPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) SslPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) TcpIdleTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpIdleTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) TcpIdleTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpIdleTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) Timeouts() LbListenerTimeoutsOutputReference {
	var returns LbListenerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListener) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/lb_listener aws_lb_listener} Resource.
func NewLbListener(scope constructs.Construct, id *string, config *LbListenerConfig) LbListener {
	_init_.Initialize()

	if err := validateNewLbListenerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbListener{}

	_jsii_.Create(
		"@cdktf/provider-aws.lbListener.LbListener",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/lb_listener aws_lb_listener} Resource.
func NewLbListener_Override(l LbListener, scope constructs.Construct, id *string, config *LbListenerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lbListener.LbListener",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LbListener)SetAlpnPolicy(val *string) {
	if err := j.validateSetAlpnPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alpnPolicy",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetLoadBalancerArn(val *string) {
	if err := j.validateSetLoadBalancerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerArn",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpRequestXAmznMtlsClientcertHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertHeaderName",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertIssuerHeaderName",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertLeafHeaderName",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertSubjectHeaderName",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertValidityHeaderName",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznTlsCipherSuiteHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznTlsCipherSuiteHeaderName",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpRequestXAmznTlsVersionHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznTlsVersionHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznTlsVersionHeaderName",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowCredentialsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowCredentialsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpResponseAccessControlAllowHeadersHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowHeadersHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowHeadersHeaderValue",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpResponseAccessControlAllowMethodsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowMethodsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowMethodsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpResponseAccessControlAllowOriginHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowOriginHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowOriginHeaderValue",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpResponseAccessControlExposeHeadersHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlExposeHeadersHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlExposeHeadersHeaderValue",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpResponseAccessControlMaxAgeHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlMaxAgeHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlMaxAgeHeaderValue",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpResponseContentSecurityPolicyHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseContentSecurityPolicyHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseContentSecurityPolicyHeaderValue",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpResponseServerEnabled(val interface{}) {
	if err := j.validateSetRoutingHttpResponseServerEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseServerEnabled",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpResponseStrictTransportSecurityHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseStrictTransportSecurityHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseStrictTransportSecurityHeaderValue",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpResponseXContentTypeOptionsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseXContentTypeOptionsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseXContentTypeOptionsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetRoutingHttpResponseXFrameOptionsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseXFrameOptionsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseXFrameOptionsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetSslPolicy(val *string) {
	if err := j.validateSetSslPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslPolicy",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_LbListener)SetTcpIdleTimeoutSeconds(val *float64) {
	if err := j.validateSetTcpIdleTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpIdleTimeoutSeconds",
		val,
	)
}

// Generates CDKTF code for importing a LbListener resource upon running "cdktf plan <stack-name>".
func LbListener_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLbListener_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lbListener.LbListener",
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
func LbListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbListener_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lbListener.LbListener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbListener_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbListener_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lbListener.LbListener",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbListener_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbListener_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lbListener.LbListener",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LbListener_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.lbListener.LbListener",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LbListener) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LbListener) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LbListener) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LbListener) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LbListener) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LbListener) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LbListener) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LbListener) PutDefaultAction(value interface{}) {
	if err := l.validatePutDefaultActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDefaultAction",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbListener) PutMutualAuthentication(value *LbListenerMutualAuthentication) {
	if err := l.validatePutMutualAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMutualAuthentication",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbListener) PutTimeouts(value *LbListenerTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbListener) ResetAlpnPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetAlpnPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		l,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetMutualAuthentication() {
	_jsii_.InvokeVoid(
		l,
		"resetMutualAuthentication",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetPort() {
	_jsii_.InvokeVoid(
		l,
		"resetPort",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetProtocol() {
	_jsii_.InvokeVoid(
		l,
		"resetProtocol",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpRequestXAmznMtlsClientcertHeaderName() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpRequestXAmznMtlsClientcertHeaderName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpRequestXAmznTlsVersionHeaderName() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpRequestXAmznTlsVersionHeaderName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpResponseAccessControlAllowHeadersHeaderValue() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpResponseAccessControlAllowHeadersHeaderValue",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpResponseAccessControlAllowMethodsHeaderValue() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpResponseAccessControlAllowMethodsHeaderValue",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpResponseAccessControlAllowOriginHeaderValue() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpResponseAccessControlAllowOriginHeaderValue",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpResponseAccessControlExposeHeadersHeaderValue() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpResponseAccessControlExposeHeadersHeaderValue",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpResponseAccessControlMaxAgeHeaderValue() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpResponseAccessControlMaxAgeHeaderValue",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpResponseContentSecurityPolicyHeaderValue() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpResponseContentSecurityPolicyHeaderValue",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpResponseServerEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpResponseServerEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpResponseStrictTransportSecurityHeaderValue() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpResponseStrictTransportSecurityHeaderValue",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpResponseXContentTypeOptionsHeaderValue() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpResponseXContentTypeOptionsHeaderValue",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetRoutingHttpResponseXFrameOptionsHeaderValue() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingHttpResponseXFrameOptionsHeaderValue",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetSslPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetSslPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetTagsAll() {
	_jsii_.InvokeVoid(
		l,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetTcpIdleTimeoutSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetTcpIdleTimeoutSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListener) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListener) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

