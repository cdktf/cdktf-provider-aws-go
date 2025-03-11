// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistener

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbListenerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// default_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#default_action LbListener#default_action}
	DefaultAction interface{} `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#load_balancer_arn LbListener#load_balancer_arn}.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#alpn_policy LbListener#alpn_policy}.
	AlpnPolicy *string `field:"optional" json:"alpnPolicy" yaml:"alpnPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#certificate_arn LbListener#certificate_arn}.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#id LbListener#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// mutual_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#mutual_authentication LbListener#mutual_authentication}
	MutualAuthentication *LbListenerMutualAuthentication `field:"optional" json:"mutualAuthentication" yaml:"mutualAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#port LbListener#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#protocol LbListener#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_request_x_amzn_mtls_clientcert_header_name LbListener#routing_http_request_x_amzn_mtls_clientcert_header_name}.
	RoutingHttpRequestXAmznMtlsClientcertHeaderName *string `field:"optional" json:"routingHttpRequestXAmznMtlsClientcertHeaderName" yaml:"routingHttpRequestXAmznMtlsClientcertHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_request_x_amzn_mtls_clientcert_issuer_header_name LbListener#routing_http_request_x_amzn_mtls_clientcert_issuer_header_name}.
	RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName *string `field:"optional" json:"routingHttpRequestXAmznMtlsClientcertIssuerHeaderName" yaml:"routingHttpRequestXAmznMtlsClientcertIssuerHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_request_x_amzn_mtls_clientcert_leaf_header_name LbListener#routing_http_request_x_amzn_mtls_clientcert_leaf_header_name}.
	RoutingHttpRequestXAmznMtlsClientcertLeafHeaderName *string `field:"optional" json:"routingHttpRequestXAmznMtlsClientcertLeafHeaderName" yaml:"routingHttpRequestXAmznMtlsClientcertLeafHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_request_x_amzn_mtls_clientcert_serial_number_header_name LbListener#routing_http_request_x_amzn_mtls_clientcert_serial_number_header_name}.
	RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName *string `field:"optional" json:"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName" yaml:"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_request_x_amzn_mtls_clientcert_subject_header_name LbListener#routing_http_request_x_amzn_mtls_clientcert_subject_header_name}.
	RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName *string `field:"optional" json:"routingHttpRequestXAmznMtlsClientcertSubjectHeaderName" yaml:"routingHttpRequestXAmznMtlsClientcertSubjectHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_request_x_amzn_mtls_clientcert_validity_header_name LbListener#routing_http_request_x_amzn_mtls_clientcert_validity_header_name}.
	RoutingHttpRequestXAmznMtlsClientcertValidityHeaderName *string `field:"optional" json:"routingHttpRequestXAmznMtlsClientcertValidityHeaderName" yaml:"routingHttpRequestXAmznMtlsClientcertValidityHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_request_x_amzn_tls_cipher_suite_header_name LbListener#routing_http_request_x_amzn_tls_cipher_suite_header_name}.
	RoutingHttpRequestXAmznTlsCipherSuiteHeaderName *string `field:"optional" json:"routingHttpRequestXAmznTlsCipherSuiteHeaderName" yaml:"routingHttpRequestXAmznTlsCipherSuiteHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_request_x_amzn_tls_version_header_name LbListener#routing_http_request_x_amzn_tls_version_header_name}.
	RoutingHttpRequestXAmznTlsVersionHeaderName *string `field:"optional" json:"routingHttpRequestXAmznTlsVersionHeaderName" yaml:"routingHttpRequestXAmznTlsVersionHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_response_access_control_allow_credentials_header_value LbListener#routing_http_response_access_control_allow_credentials_header_value}.
	RoutingHttpResponseAccessControlAllowCredentialsHeaderValue *string `field:"optional" json:"routingHttpResponseAccessControlAllowCredentialsHeaderValue" yaml:"routingHttpResponseAccessControlAllowCredentialsHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_response_access_control_allow_headers_header_value LbListener#routing_http_response_access_control_allow_headers_header_value}.
	RoutingHttpResponseAccessControlAllowHeadersHeaderValue *string `field:"optional" json:"routingHttpResponseAccessControlAllowHeadersHeaderValue" yaml:"routingHttpResponseAccessControlAllowHeadersHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_response_access_control_allow_methods_header_value LbListener#routing_http_response_access_control_allow_methods_header_value}.
	RoutingHttpResponseAccessControlAllowMethodsHeaderValue *string `field:"optional" json:"routingHttpResponseAccessControlAllowMethodsHeaderValue" yaml:"routingHttpResponseAccessControlAllowMethodsHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_response_access_control_allow_origin_header_value LbListener#routing_http_response_access_control_allow_origin_header_value}.
	RoutingHttpResponseAccessControlAllowOriginHeaderValue *string `field:"optional" json:"routingHttpResponseAccessControlAllowOriginHeaderValue" yaml:"routingHttpResponseAccessControlAllowOriginHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_response_access_control_expose_headers_header_value LbListener#routing_http_response_access_control_expose_headers_header_value}.
	RoutingHttpResponseAccessControlExposeHeadersHeaderValue *string `field:"optional" json:"routingHttpResponseAccessControlExposeHeadersHeaderValue" yaml:"routingHttpResponseAccessControlExposeHeadersHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_response_access_control_max_age_header_value LbListener#routing_http_response_access_control_max_age_header_value}.
	RoutingHttpResponseAccessControlMaxAgeHeaderValue *string `field:"optional" json:"routingHttpResponseAccessControlMaxAgeHeaderValue" yaml:"routingHttpResponseAccessControlMaxAgeHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_response_content_security_policy_header_value LbListener#routing_http_response_content_security_policy_header_value}.
	RoutingHttpResponseContentSecurityPolicyHeaderValue *string `field:"optional" json:"routingHttpResponseContentSecurityPolicyHeaderValue" yaml:"routingHttpResponseContentSecurityPolicyHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_response_server_enabled LbListener#routing_http_response_server_enabled}.
	RoutingHttpResponseServerEnabled interface{} `field:"optional" json:"routingHttpResponseServerEnabled" yaml:"routingHttpResponseServerEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_response_strict_transport_security_header_value LbListener#routing_http_response_strict_transport_security_header_value}.
	RoutingHttpResponseStrictTransportSecurityHeaderValue *string `field:"optional" json:"routingHttpResponseStrictTransportSecurityHeaderValue" yaml:"routingHttpResponseStrictTransportSecurityHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_response_x_content_type_options_header_value LbListener#routing_http_response_x_content_type_options_header_value}.
	RoutingHttpResponseXContentTypeOptionsHeaderValue *string `field:"optional" json:"routingHttpResponseXContentTypeOptionsHeaderValue" yaml:"routingHttpResponseXContentTypeOptionsHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#routing_http_response_x_frame_options_header_value LbListener#routing_http_response_x_frame_options_header_value}.
	RoutingHttpResponseXFrameOptionsHeaderValue *string `field:"optional" json:"routingHttpResponseXFrameOptionsHeaderValue" yaml:"routingHttpResponseXFrameOptionsHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#ssl_policy LbListener#ssl_policy}.
	SslPolicy *string `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#tags LbListener#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#tags_all LbListener#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#tcp_idle_timeout_seconds LbListener#tcp_idle_timeout_seconds}.
	TcpIdleTimeoutSeconds *float64 `field:"optional" json:"tcpIdleTimeoutSeconds" yaml:"tcpIdleTimeoutSeconds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lb_listener#timeouts LbListener#timeouts}
	Timeouts *LbListenerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

