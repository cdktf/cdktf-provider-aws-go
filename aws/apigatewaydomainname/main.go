// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewaydomainname

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.apiGatewayDomainName.ApiGatewayDomainName",
		reflect.TypeOf((*ApiGatewayDomainName)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "certificateArn", GoGetter: "CertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "certificateArnInput", GoGetter: "CertificateArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "certificateBody", GoGetter: "CertificateBody"},
			_jsii_.MemberProperty{JsiiProperty: "certificateBodyInput", GoGetter: "CertificateBodyInput"},
			_jsii_.MemberProperty{JsiiProperty: "certificateChain", GoGetter: "CertificateChain"},
			_jsii_.MemberProperty{JsiiProperty: "certificateChainInput", GoGetter: "CertificateChainInput"},
			_jsii_.MemberProperty{JsiiProperty: "certificateName", GoGetter: "CertificateName"},
			_jsii_.MemberProperty{JsiiProperty: "certificateNameInput", GoGetter: "CertificateNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "certificatePrivateKey", GoGetter: "CertificatePrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "certificatePrivateKeyInput", GoGetter: "CertificatePrivateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "certificateUploadDate", GoGetter: "CertificateUploadDate"},
			_jsii_.MemberProperty{JsiiProperty: "cloudfrontDomainName", GoGetter: "CloudfrontDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "cloudfrontZoneId", GoGetter: "CloudfrontZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "domainNameId", GoGetter: "DomainNameId"},
			_jsii_.MemberProperty{JsiiProperty: "domainNameInput", GoGetter: "DomainNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "endpointConfiguration", GoGetter: "EndpointConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "endpointConfigurationInput", GoGetter: "EndpointConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "mutualTlsAuthentication", GoGetter: "MutualTlsAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "mutualTlsAuthenticationInput", GoGetter: "MutualTlsAuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ownershipVerificationCertificateArn", GoGetter: "OwnershipVerificationCertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "ownershipVerificationCertificateArnInput", GoGetter: "OwnershipVerificationCertificateArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "policyInput", GoGetter: "PolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putEndpointConfiguration", GoMethod: "PutEndpointConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putMutualTlsAuthentication", GoMethod: "PutMutualTlsAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "regionalCertificateArn", GoGetter: "RegionalCertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "regionalCertificateArnInput", GoGetter: "RegionalCertificateArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "regionalCertificateName", GoGetter: "RegionalCertificateName"},
			_jsii_.MemberProperty{JsiiProperty: "regionalCertificateNameInput", GoGetter: "RegionalCertificateNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "regionalDomainName", GoGetter: "RegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "regionalZoneId", GoGetter: "RegionalZoneId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificateArn", GoMethod: "ResetCertificateArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificateBody", GoMethod: "ResetCertificateBody"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificateChain", GoMethod: "ResetCertificateChain"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificateName", GoMethod: "ResetCertificateName"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificatePrivateKey", GoMethod: "ResetCertificatePrivateKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndpointConfiguration", GoMethod: "ResetEndpointConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMutualTlsAuthentication", GoMethod: "ResetMutualTlsAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOwnershipVerificationCertificateArn", GoMethod: "ResetOwnershipVerificationCertificateArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicy", GoMethod: "ResetPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegionalCertificateArn", GoMethod: "ResetRegionalCertificateArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegionalCertificateName", GoMethod: "ResetRegionalCertificateName"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityPolicy", GoMethod: "ResetSecurityPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "securityPolicy", GoGetter: "SecurityPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "securityPolicyInput", GoGetter: "SecurityPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayDomainName{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.apiGatewayDomainName.ApiGatewayDomainNameConfig",
		reflect.TypeOf((*ApiGatewayDomainNameConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.apiGatewayDomainName.ApiGatewayDomainNameEndpointConfiguration",
		reflect.TypeOf((*ApiGatewayDomainNameEndpointConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.apiGatewayDomainName.ApiGatewayDomainNameEndpointConfigurationOutputReference",
		reflect.TypeOf((*ApiGatewayDomainNameEndpointConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "types", GoGetter: "Types"},
			_jsii_.MemberProperty{JsiiProperty: "typesInput", GoGetter: "TypesInput"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.apiGatewayDomainName.ApiGatewayDomainNameMutualTlsAuthentication",
		reflect.TypeOf((*ApiGatewayDomainNameMutualTlsAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.apiGatewayDomainName.ApiGatewayDomainNameMutualTlsAuthenticationOutputReference",
		reflect.TypeOf((*ApiGatewayDomainNameMutualTlsAuthenticationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetTruststoreVersion", GoMethod: "ResetTruststoreVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "truststoreUri", GoGetter: "TruststoreUri"},
			_jsii_.MemberProperty{JsiiProperty: "truststoreUriInput", GoGetter: "TruststoreUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "truststoreVersion", GoGetter: "TruststoreVersion"},
			_jsii_.MemberProperty{JsiiProperty: "truststoreVersionInput", GoGetter: "TruststoreVersionInput"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
