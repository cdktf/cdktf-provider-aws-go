package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.provider.AwsProvider",
		reflect.TypeOf((*AwsProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessKey", GoGetter: "AccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "accessKeyInput", GoGetter: "AccessKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedAccountIds", GoGetter: "AllowedAccountIds"},
			_jsii_.MemberProperty{JsiiProperty: "allowedAccountIdsInput", GoGetter: "AllowedAccountIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRole", GoGetter: "AssumeRole"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleInput", GoGetter: "AssumeRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleWithWebIdentity", GoGetter: "AssumeRoleWithWebIdentity"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleWithWebIdentityInput", GoGetter: "AssumeRoleWithWebIdentityInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "customCaBundle", GoGetter: "CustomCaBundle"},
			_jsii_.MemberProperty{JsiiProperty: "customCaBundleInput", GoGetter: "CustomCaBundleInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTags", GoGetter: "DefaultTags"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTagsInput", GoGetter: "DefaultTagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "ec2MetadataServiceEndpoint", GoGetter: "Ec2MetadataServiceEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "ec2MetadataServiceEndpointInput", GoGetter: "Ec2MetadataServiceEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "ec2MetadataServiceEndpointMode", GoGetter: "Ec2MetadataServiceEndpointMode"},
			_jsii_.MemberProperty{JsiiProperty: "ec2MetadataServiceEndpointModeInput", GoGetter: "Ec2MetadataServiceEndpointModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberProperty{JsiiProperty: "endpointsInput", GoGetter: "EndpointsInput"},
			_jsii_.MemberProperty{JsiiProperty: "forbiddenAccountIds", GoGetter: "ForbiddenAccountIds"},
			_jsii_.MemberProperty{JsiiProperty: "forbiddenAccountIdsInput", GoGetter: "ForbiddenAccountIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "httpProxy", GoGetter: "HttpProxy"},
			_jsii_.MemberProperty{JsiiProperty: "httpProxyInput", GoGetter: "HttpProxyInput"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreTags", GoGetter: "IgnoreTags"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreTagsInput", GoGetter: "IgnoreTagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "insecure", GoGetter: "Insecure"},
			_jsii_.MemberProperty{JsiiProperty: "insecureInput", GoGetter: "InsecureInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetries", GoGetter: "MaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesInput", GoGetter: "MaxRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "profile", GoGetter: "Profile"},
			_jsii_.MemberProperty{JsiiProperty: "profileInput", GoGetter: "ProfileInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessKey", GoMethod: "ResetAccessKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedAccountIds", GoMethod: "ResetAllowedAccountIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssumeRole", GoMethod: "ResetAssumeRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssumeRoleWithWebIdentity", GoMethod: "ResetAssumeRoleWithWebIdentity"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomCaBundle", GoMethod: "ResetCustomCaBundle"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTags", GoMethod: "ResetDefaultTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetEc2MetadataServiceEndpoint", GoMethod: "ResetEc2MetadataServiceEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetEc2MetadataServiceEndpointMode", GoMethod: "ResetEc2MetadataServiceEndpointMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndpoints", GoMethod: "ResetEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resetForbiddenAccountIds", GoMethod: "ResetForbiddenAccountIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpProxy", GoMethod: "ResetHttpProxy"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreTags", GoMethod: "ResetIgnoreTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecure", GoMethod: "ResetInsecure"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRetries", GoMethod: "ResetMaxRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProfile", GoMethod: "ResetProfile"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3ForcePathStyle", GoMethod: "ResetS3ForcePathStyle"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3UsePathStyle", GoMethod: "ResetS3UsePathStyle"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretKey", GoMethod: "ResetSecretKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharedConfigFiles", GoMethod: "ResetSharedConfigFiles"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharedCredentialsFile", GoMethod: "ResetSharedCredentialsFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharedCredentialsFiles", GoMethod: "ResetSharedCredentialsFiles"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipCredentialsValidation", GoMethod: "ResetSkipCredentialsValidation"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipGetEc2Platforms", GoMethod: "ResetSkipGetEc2Platforms"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipMetadataApiCheck", GoMethod: "ResetSkipMetadataApiCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipRegionValidation", GoMethod: "ResetSkipRegionValidation"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipRequestingAccountId", GoMethod: "ResetSkipRequestingAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStsRegion", GoMethod: "ResetStsRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseDualstackEndpoint", GoMethod: "ResetUseDualstackEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseFipsEndpoint", GoMethod: "ResetUseFipsEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "s3ForcePathStyle", GoGetter: "S3ForcePathStyle"},
			_jsii_.MemberProperty{JsiiProperty: "s3ForcePathStyleInput", GoGetter: "S3ForcePathStyleInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3UsePathStyle", GoGetter: "S3UsePathStyle"},
			_jsii_.MemberProperty{JsiiProperty: "s3UsePathStyleInput", GoGetter: "S3UsePathStyleInput"},
			_jsii_.MemberProperty{JsiiProperty: "secretKey", GoGetter: "SecretKey"},
			_jsii_.MemberProperty{JsiiProperty: "secretKeyInput", GoGetter: "SecretKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "sharedConfigFiles", GoGetter: "SharedConfigFiles"},
			_jsii_.MemberProperty{JsiiProperty: "sharedConfigFilesInput", GoGetter: "SharedConfigFilesInput"},
			_jsii_.MemberProperty{JsiiProperty: "sharedCredentialsFile", GoGetter: "SharedCredentialsFile"},
			_jsii_.MemberProperty{JsiiProperty: "sharedCredentialsFileInput", GoGetter: "SharedCredentialsFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "sharedCredentialsFiles", GoGetter: "SharedCredentialsFiles"},
			_jsii_.MemberProperty{JsiiProperty: "sharedCredentialsFilesInput", GoGetter: "SharedCredentialsFilesInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipCredentialsValidation", GoGetter: "SkipCredentialsValidation"},
			_jsii_.MemberProperty{JsiiProperty: "skipCredentialsValidationInput", GoGetter: "SkipCredentialsValidationInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipGetEc2Platforms", GoGetter: "SkipGetEc2Platforms"},
			_jsii_.MemberProperty{JsiiProperty: "skipGetEc2PlatformsInput", GoGetter: "SkipGetEc2PlatformsInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipMetadataApiCheck", GoGetter: "SkipMetadataApiCheck"},
			_jsii_.MemberProperty{JsiiProperty: "skipMetadataApiCheckInput", GoGetter: "SkipMetadataApiCheckInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipRegionValidation", GoGetter: "SkipRegionValidation"},
			_jsii_.MemberProperty{JsiiProperty: "skipRegionValidationInput", GoGetter: "SkipRegionValidationInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipRequestingAccountId", GoGetter: "SkipRequestingAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "skipRequestingAccountIdInput", GoGetter: "SkipRequestingAccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "stsRegion", GoGetter: "StsRegion"},
			_jsii_.MemberProperty{JsiiProperty: "stsRegionInput", GoGetter: "StsRegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "useDualstackEndpoint", GoGetter: "UseDualstackEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "useDualstackEndpointInput", GoGetter: "UseDualstackEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "useFipsEndpoint", GoGetter: "UseFipsEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "useFipsEndpointInput", GoGetter: "UseFipsEndpointInput"},
		},
		func() interface{} {
			j := jsiiProxy_AwsProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.provider.AwsProviderAssumeRole",
		reflect.TypeOf((*AwsProviderAssumeRole)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.provider.AwsProviderAssumeRoleWithWebIdentity",
		reflect.TypeOf((*AwsProviderAssumeRoleWithWebIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.provider.AwsProviderConfig",
		reflect.TypeOf((*AwsProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.provider.AwsProviderDefaultTags",
		reflect.TypeOf((*AwsProviderDefaultTags)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.provider.AwsProviderEndpoints",
		reflect.TypeOf((*AwsProviderEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.provider.AwsProviderIgnoreTags",
		reflect.TypeOf((*AwsProviderIgnoreTags)(nil)).Elem(),
	)
}