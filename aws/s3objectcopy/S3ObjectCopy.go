// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3objectcopy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/s3objectcopy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_object_copy aws_s3_object_copy}.
type S3ObjectCopy interface {
	cdktf.TerraformResource
	Acl() *string
	SetAcl(val *string)
	AclInput() *string
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
	BucketKeyEnabled() interface{}
	SetBucketKeyEnabled(val interface{})
	BucketKeyEnabledInput() interface{}
	CacheControl() *string
	SetCacheControl(val *string)
	CacheControlInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChecksumAlgorithm() *string
	SetChecksumAlgorithm(val *string)
	ChecksumAlgorithmInput() *string
	ChecksumCrc32() *string
	ChecksumCrc32C() *string
	ChecksumSha1() *string
	ChecksumSha256() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentDisposition() *string
	SetContentDisposition(val *string)
	ContentDispositionInput() *string
	ContentEncoding() *string
	SetContentEncoding(val *string)
	ContentEncodingInput() *string
	ContentLanguage() *string
	SetContentLanguage(val *string)
	ContentLanguageInput() *string
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	CopyIfMatch() *string
	SetCopyIfMatch(val *string)
	CopyIfMatchInput() *string
	CopyIfModifiedSince() *string
	SetCopyIfModifiedSince(val *string)
	CopyIfModifiedSinceInput() *string
	CopyIfNoneMatch() *string
	SetCopyIfNoneMatch(val *string)
	CopyIfNoneMatchInput() *string
	CopyIfUnmodifiedSince() *string
	SetCopyIfUnmodifiedSince(val *string)
	CopyIfUnmodifiedSinceInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomerAlgorithm() *string
	SetCustomerAlgorithm(val *string)
	CustomerAlgorithmInput() *string
	CustomerKey() *string
	SetCustomerKey(val *string)
	CustomerKeyInput() *string
	CustomerKeyMd5() *string
	SetCustomerKeyMd5(val *string)
	CustomerKeyMd5Input() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Etag() *string
	ExpectedBucketOwner() *string
	SetExpectedBucketOwner(val *string)
	ExpectedBucketOwnerInput() *string
	ExpectedSourceBucketOwner() *string
	SetExpectedSourceBucketOwner(val *string)
	ExpectedSourceBucketOwnerInput() *string
	Expiration() *string
	Expires() *string
	SetExpires(val *string)
	ExpiresInput() *string
	ForceDestroy() interface{}
	SetForceDestroy(val interface{})
	ForceDestroyInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Grant() S3ObjectCopyGrantList
	GrantInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	KmsEncryptionContext() *string
	SetKmsEncryptionContext(val *string)
	KmsEncryptionContextInput() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	LastModified() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataDirective() *string
	SetMetadataDirective(val *string)
	MetadataDirectiveInput() *string
	MetadataInput() *map[string]*string
	// The tree node.
	Node() constructs.Node
	ObjectLockLegalHoldStatus() *string
	SetObjectLockLegalHoldStatus(val *string)
	ObjectLockLegalHoldStatusInput() *string
	ObjectLockMode() *string
	SetObjectLockMode(val *string)
	ObjectLockModeInput() *string
	ObjectLockRetainUntilDate() *string
	SetObjectLockRetainUntilDate(val *string)
	ObjectLockRetainUntilDateInput() *string
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
	RequestCharged() cdktf.IResolvable
	RequestPayer() *string
	SetRequestPayer(val *string)
	RequestPayerInput() *string
	ServerSideEncryption() *string
	SetServerSideEncryption(val *string)
	ServerSideEncryptionInput() *string
	Source() *string
	SetSource(val *string)
	SourceCustomerAlgorithm() *string
	SetSourceCustomerAlgorithm(val *string)
	SourceCustomerAlgorithmInput() *string
	SourceCustomerKey() *string
	SetSourceCustomerKey(val *string)
	SourceCustomerKeyInput() *string
	SourceCustomerKeyMd5() *string
	SetSourceCustomerKeyMd5(val *string)
	SourceCustomerKeyMd5Input() *string
	SourceInput() *string
	SourceVersionId() *string
	StorageClass() *string
	SetStorageClass(val *string)
	StorageClassInput() *string
	TaggingDirective() *string
	SetTaggingDirective(val *string)
	TaggingDirectiveInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VersionId() *string
	WebsiteRedirect() *string
	SetWebsiteRedirect(val *string)
	WebsiteRedirectInput() *string
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
	PutGrant(value interface{})
	ResetAcl()
	ResetBucketKeyEnabled()
	ResetCacheControl()
	ResetChecksumAlgorithm()
	ResetContentDisposition()
	ResetContentEncoding()
	ResetContentLanguage()
	ResetContentType()
	ResetCopyIfMatch()
	ResetCopyIfModifiedSince()
	ResetCopyIfNoneMatch()
	ResetCopyIfUnmodifiedSince()
	ResetCustomerAlgorithm()
	ResetCustomerKey()
	ResetCustomerKeyMd5()
	ResetExpectedBucketOwner()
	ResetExpectedSourceBucketOwner()
	ResetExpires()
	ResetForceDestroy()
	ResetGrant()
	ResetId()
	ResetKmsEncryptionContext()
	ResetKmsKeyId()
	ResetMetadata()
	ResetMetadataDirective()
	ResetObjectLockLegalHoldStatus()
	ResetObjectLockMode()
	ResetObjectLockRetainUntilDate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestPayer()
	ResetServerSideEncryption()
	ResetSourceCustomerAlgorithm()
	ResetSourceCustomerKey()
	ResetSourceCustomerKeyMd5()
	ResetStorageClass()
	ResetTaggingDirective()
	ResetTags()
	ResetTagsAll()
	ResetWebsiteRedirect()
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

// The jsii proxy struct for S3ObjectCopy
type jsiiProxy_S3ObjectCopy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_S3ObjectCopy) Acl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) AclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) BucketKeyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketKeyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) BucketKeyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketKeyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CacheControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CacheControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ChecksumAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ChecksumAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ChecksumCrc32() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumCrc32",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ChecksumCrc32C() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumCrc32C",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ChecksumSha1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumSha1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ChecksumSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ContentDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ContentDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDispositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ContentEncoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ContentEncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ContentLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ContentLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CopyIfMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CopyIfMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CopyIfModifiedSince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfModifiedSince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CopyIfModifiedSinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfModifiedSinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CopyIfNoneMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfNoneMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CopyIfNoneMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfNoneMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CopyIfUnmodifiedSince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfUnmodifiedSince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CopyIfUnmodifiedSinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfUnmodifiedSinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CustomerAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CustomerAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CustomerKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CustomerKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CustomerKeyMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerKeyMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) CustomerKeyMd5Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerKeyMd5Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ExpectedBucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ExpectedBucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ExpectedSourceBucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedSourceBucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ExpectedSourceBucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedSourceBucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Expiration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Expires() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ExpiresInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Grant() S3ObjectCopyGrantList {
	var returns S3ObjectCopyGrantList
	_jsii_.Get(
		j,
		"grant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) GrantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) KmsEncryptionContext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsEncryptionContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) KmsEncryptionContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsEncryptionContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) MetadataDirective() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataDirective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) MetadataDirectiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataDirectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ObjectLockLegalHoldStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockLegalHoldStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ObjectLockLegalHoldStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockLegalHoldStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ObjectLockMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ObjectLockModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ObjectLockRetainUntilDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockRetainUntilDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ObjectLockRetainUntilDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockRetainUntilDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) RequestCharged() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"requestCharged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) RequestPayer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) RequestPayerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPayerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ServerSideEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) ServerSideEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) SourceCustomerAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) SourceCustomerAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) SourceCustomerKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) SourceCustomerKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) SourceCustomerKeyMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerKeyMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) SourceCustomerKeyMd5Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerKeyMd5Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) SourceVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) TaggingDirective() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taggingDirective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) TaggingDirectiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taggingDirectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) WebsiteRedirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectCopy) WebsiteRedirectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteRedirectInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_object_copy aws_s3_object_copy} Resource.
func NewS3ObjectCopy(scope constructs.Construct, id *string, config *S3ObjectCopyConfig) S3ObjectCopy {
	_init_.Initialize()

	if err := validateNewS3ObjectCopyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ObjectCopy{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3ObjectCopy.S3ObjectCopy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_object_copy aws_s3_object_copy} Resource.
func NewS3ObjectCopy_Override(s S3ObjectCopy, scope constructs.Construct, id *string, config *S3ObjectCopyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3ObjectCopy.S3ObjectCopy",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetAcl(val *string) {
	if err := j.validateSetAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acl",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetBucketKeyEnabled(val interface{}) {
	if err := j.validateSetBucketKeyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketKeyEnabled",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetCacheControl(val *string) {
	if err := j.validateSetCacheControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheControl",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetChecksumAlgorithm(val *string) {
	if err := j.validateSetChecksumAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checksumAlgorithm",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetContentDisposition(val *string) {
	if err := j.validateSetContentDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentDisposition",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetContentEncoding(val *string) {
	if err := j.validateSetContentEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentEncoding",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetContentLanguage(val *string) {
	if err := j.validateSetContentLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentLanguage",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetCopyIfMatch(val *string) {
	if err := j.validateSetCopyIfMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyIfMatch",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetCopyIfModifiedSince(val *string) {
	if err := j.validateSetCopyIfModifiedSinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyIfModifiedSince",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetCopyIfNoneMatch(val *string) {
	if err := j.validateSetCopyIfNoneMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyIfNoneMatch",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetCopyIfUnmodifiedSince(val *string) {
	if err := j.validateSetCopyIfUnmodifiedSinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyIfUnmodifiedSince",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetCustomerAlgorithm(val *string) {
	if err := j.validateSetCustomerAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerAlgorithm",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetCustomerKey(val *string) {
	if err := j.validateSetCustomerKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerKey",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetCustomerKeyMd5(val *string) {
	if err := j.validateSetCustomerKeyMd5Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerKeyMd5",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetExpectedBucketOwner(val *string) {
	if err := j.validateSetExpectedBucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedBucketOwner",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetExpectedSourceBucketOwner(val *string) {
	if err := j.validateSetExpectedSourceBucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedSourceBucketOwner",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetExpires(val *string) {
	if err := j.validateSetExpiresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expires",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetKmsEncryptionContext(val *string) {
	if err := j.validateSetKmsEncryptionContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsEncryptionContext",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetMetadataDirective(val *string) {
	if err := j.validateSetMetadataDirectiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataDirective",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetObjectLockLegalHoldStatus(val *string) {
	if err := j.validateSetObjectLockLegalHoldStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockLegalHoldStatus",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetObjectLockMode(val *string) {
	if err := j.validateSetObjectLockModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockMode",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetObjectLockRetainUntilDate(val *string) {
	if err := j.validateSetObjectLockRetainUntilDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockRetainUntilDate",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetRequestPayer(val *string) {
	if err := j.validateSetRequestPayerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestPayer",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetServerSideEncryption(val *string) {
	if err := j.validateSetServerSideEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideEncryption",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetSourceCustomerAlgorithm(val *string) {
	if err := j.validateSetSourceCustomerAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCustomerAlgorithm",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetSourceCustomerKey(val *string) {
	if err := j.validateSetSourceCustomerKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCustomerKey",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetSourceCustomerKeyMd5(val *string) {
	if err := j.validateSetSourceCustomerKeyMd5Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCustomerKeyMd5",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetTaggingDirective(val *string) {
	if err := j.validateSetTaggingDirectiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taggingDirective",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_S3ObjectCopy)SetWebsiteRedirect(val *string) {
	if err := j.validateSetWebsiteRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"websiteRedirect",
		val,
	)
}

// Generates CDKTF code for importing a S3ObjectCopy resource upon running "cdktf plan <stack-name>".
func S3ObjectCopy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateS3ObjectCopy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3ObjectCopy.S3ObjectCopy",
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
func S3ObjectCopy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3ObjectCopy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3ObjectCopy.S3ObjectCopy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func S3ObjectCopy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3ObjectCopy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3ObjectCopy.S3ObjectCopy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func S3ObjectCopy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3ObjectCopy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3ObjectCopy.S3ObjectCopy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func S3ObjectCopy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.s3ObjectCopy.S3ObjectCopy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_S3ObjectCopy) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_S3ObjectCopy) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_S3ObjectCopy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_S3ObjectCopy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_S3ObjectCopy) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_S3ObjectCopy) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_S3ObjectCopy) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_S3ObjectCopy) PutGrant(value interface{}) {
	if err := s.validatePutGrantParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGrant",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetAcl() {
	_jsii_.InvokeVoid(
		s,
		"resetAcl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetBucketKeyEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetBucketKeyEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetCacheControl() {
	_jsii_.InvokeVoid(
		s,
		"resetCacheControl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetChecksumAlgorithm() {
	_jsii_.InvokeVoid(
		s,
		"resetChecksumAlgorithm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetContentDisposition() {
	_jsii_.InvokeVoid(
		s,
		"resetContentDisposition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetContentEncoding() {
	_jsii_.InvokeVoid(
		s,
		"resetContentEncoding",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetContentLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetContentLanguage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetContentType() {
	_jsii_.InvokeVoid(
		s,
		"resetContentType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetCopyIfMatch() {
	_jsii_.InvokeVoid(
		s,
		"resetCopyIfMatch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetCopyIfModifiedSince() {
	_jsii_.InvokeVoid(
		s,
		"resetCopyIfModifiedSince",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetCopyIfNoneMatch() {
	_jsii_.InvokeVoid(
		s,
		"resetCopyIfNoneMatch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetCopyIfUnmodifiedSince() {
	_jsii_.InvokeVoid(
		s,
		"resetCopyIfUnmodifiedSince",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetCustomerAlgorithm() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomerAlgorithm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetCustomerKey() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomerKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetCustomerKeyMd5() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomerKeyMd5",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetExpectedBucketOwner() {
	_jsii_.InvokeVoid(
		s,
		"resetExpectedBucketOwner",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetExpectedSourceBucketOwner() {
	_jsii_.InvokeVoid(
		s,
		"resetExpectedSourceBucketOwner",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetExpires() {
	_jsii_.InvokeVoid(
		s,
		"resetExpires",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		s,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetGrant() {
	_jsii_.InvokeVoid(
		s,
		"resetGrant",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetKmsEncryptionContext() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsEncryptionContext",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetMetadata() {
	_jsii_.InvokeVoid(
		s,
		"resetMetadata",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetMetadataDirective() {
	_jsii_.InvokeVoid(
		s,
		"resetMetadataDirective",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetObjectLockLegalHoldStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectLockLegalHoldStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetObjectLockMode() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectLockMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetObjectLockRetainUntilDate() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectLockRetainUntilDate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetRequestPayer() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestPayer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetServerSideEncryption() {
	_jsii_.InvokeVoid(
		s,
		"resetServerSideEncryption",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetSourceCustomerAlgorithm() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceCustomerAlgorithm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetSourceCustomerKey() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceCustomerKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetSourceCustomerKeyMd5() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceCustomerKeyMd5",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetStorageClass() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetTaggingDirective() {
	_jsii_.InvokeVoid(
		s,
		"resetTaggingDirective",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) ResetWebsiteRedirect() {
	_jsii_.InvokeVoid(
		s,
		"resetWebsiteRedirect",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectCopy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectCopy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

