// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawss3object

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.dataAwsS3Object.DataAwsS3Object",
		reflect.TypeOf((*DataAwsS3Object)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "body", GoGetter: "Body"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "bucketInput", GoGetter: "BucketInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketKeyEnabled", GoGetter: "BucketKeyEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "cacheControl", GoGetter: "CacheControl"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "checksumCrc32", GoGetter: "ChecksumCrc32"},
			_jsii_.MemberProperty{JsiiProperty: "checksumCrc32C", GoGetter: "ChecksumCrc32C"},
			_jsii_.MemberProperty{JsiiProperty: "checksumCrc64Nvme", GoGetter: "ChecksumCrc64Nvme"},
			_jsii_.MemberProperty{JsiiProperty: "checksumMode", GoGetter: "ChecksumMode"},
			_jsii_.MemberProperty{JsiiProperty: "checksumModeInput", GoGetter: "ChecksumModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "checksumSha1", GoGetter: "ChecksumSha1"},
			_jsii_.MemberProperty{JsiiProperty: "checksumSha256", GoGetter: "ChecksumSha256"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "contentDisposition", GoGetter: "ContentDisposition"},
			_jsii_.MemberProperty{JsiiProperty: "contentEncoding", GoGetter: "ContentEncoding"},
			_jsii_.MemberProperty{JsiiProperty: "contentLanguage", GoGetter: "ContentLanguage"},
			_jsii_.MemberProperty{JsiiProperty: "contentLength", GoGetter: "ContentLength"},
			_jsii_.MemberProperty{JsiiProperty: "contentType", GoGetter: "ContentType"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "expiration", GoGetter: "Expiration"},
			_jsii_.MemberProperty{JsiiProperty: "expires", GoGetter: "Expires"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "lastModified", GoGetter: "LastModified"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "objectLockLegalHoldStatus", GoGetter: "ObjectLockLegalHoldStatus"},
			_jsii_.MemberProperty{JsiiProperty: "objectLockMode", GoGetter: "ObjectLockMode"},
			_jsii_.MemberProperty{JsiiProperty: "objectLockRetainUntilDate", GoGetter: "ObjectLockRetainUntilDate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "range", GoGetter: "Range"},
			_jsii_.MemberProperty{JsiiProperty: "rangeInput", GoGetter: "RangeInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetChecksumMode", GoMethod: "ResetChecksumMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRange", GoMethod: "ResetRange"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetVersionId", GoMethod: "ResetVersionId"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideEncryption", GoGetter: "ServerSideEncryption"},
			_jsii_.MemberProperty{JsiiProperty: "sseKmsKeyId", GoGetter: "SseKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "storageClass", GoGetter: "StorageClass"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "versionId", GoGetter: "VersionId"},
			_jsii_.MemberProperty{JsiiProperty: "versionIdInput", GoGetter: "VersionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "websiteRedirectLocation", GoGetter: "WebsiteRedirectLocation"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsS3Object{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.dataAwsS3Object.DataAwsS3ObjectConfig",
		reflect.TypeOf((*DataAwsS3ObjectConfig)(nil)).Elem(),
	)
}
