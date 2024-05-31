// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketobject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/s3bucketobject/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/s3_bucket_object aws_s3_bucket_object}.
type S3BucketObject interface {
	cdktf.TerraformResource
	Acl() *string
	SetAcl(val *string)
	AclInput() *string
	Arn() *string
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
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Content() *string
	SetContent(val *string)
	ContentBase64() *string
	SetContentBase64(val *string)
	ContentBase64Input() *string
	ContentDisposition() *string
	SetContentDisposition(val *string)
	ContentDispositionInput() *string
	ContentEncoding() *string
	SetContentEncoding(val *string)
	ContentEncodingInput() *string
	ContentInput() *string
	ContentLanguage() *string
	SetContentLanguage(val *string)
	ContentLanguageInput() *string
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
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
	ServerSideEncryption() *string
	SetServerSideEncryption(val *string)
	ServerSideEncryptionInput() *string
	Source() *string
	SetSource(val *string)
	SourceHash() *string
	SetSourceHash(val *string)
	SourceHashInput() *string
	SourceInput() *string
	StorageClass() *string
	SetStorageClass(val *string)
	StorageClassInput() *string
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
	ResetAcl()
	ResetBucketKeyEnabled()
	ResetCacheControl()
	ResetContent()
	ResetContentBase64()
	ResetContentDisposition()
	ResetContentEncoding()
	ResetContentLanguage()
	ResetContentType()
	ResetEtag()
	ResetForceDestroy()
	ResetId()
	ResetKmsKeyId()
	ResetMetadata()
	ResetObjectLockLegalHoldStatus()
	ResetObjectLockMode()
	ResetObjectLockRetainUntilDate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServerSideEncryption()
	ResetSource()
	ResetSourceHash()
	ResetStorageClass()
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

// The jsii proxy struct for S3BucketObject
type jsiiProxy_S3BucketObject struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_S3BucketObject) Acl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) AclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) BucketKeyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketKeyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) BucketKeyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketKeyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) CacheControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) CacheControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ContentBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ContentBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ContentDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ContentDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDispositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ContentEncoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ContentEncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ContentLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ContentLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ObjectLockLegalHoldStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockLegalHoldStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ObjectLockLegalHoldStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockLegalHoldStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ObjectLockMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ObjectLockModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ObjectLockRetainUntilDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockRetainUntilDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ObjectLockRetainUntilDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockRetainUntilDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ServerSideEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) ServerSideEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) SourceHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) SourceHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) WebsiteRedirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketObject) WebsiteRedirectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteRedirectInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/s3_bucket_object aws_s3_bucket_object} Resource.
func NewS3BucketObject(scope constructs.Construct, id *string, config *S3BucketObjectConfig) S3BucketObject {
	_init_.Initialize()

	if err := validateNewS3BucketObjectParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketObject{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3BucketObject.S3BucketObject",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/s3_bucket_object aws_s3_bucket_object} Resource.
func NewS3BucketObject_Override(s S3BucketObject, scope constructs.Construct, id *string, config *S3BucketObjectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3BucketObject.S3BucketObject",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_S3BucketObject)SetAcl(val *string) {
	if err := j.validateSetAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acl",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetBucketKeyEnabled(val interface{}) {
	if err := j.validateSetBucketKeyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketKeyEnabled",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetCacheControl(val *string) {
	if err := j.validateSetCacheControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheControl",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetContentBase64(val *string) {
	if err := j.validateSetContentBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentBase64",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetContentDisposition(val *string) {
	if err := j.validateSetContentDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentDisposition",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetContentEncoding(val *string) {
	if err := j.validateSetContentEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentEncoding",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetContentLanguage(val *string) {
	if err := j.validateSetContentLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentLanguage",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetEtag(val *string) {
	if err := j.validateSetEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetObjectLockLegalHoldStatus(val *string) {
	if err := j.validateSetObjectLockLegalHoldStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockLegalHoldStatus",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetObjectLockMode(val *string) {
	if err := j.validateSetObjectLockModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockMode",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetObjectLockRetainUntilDate(val *string) {
	if err := j.validateSetObjectLockRetainUntilDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockRetainUntilDate",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetServerSideEncryption(val *string) {
	if err := j.validateSetServerSideEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideEncryption",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetSourceHash(val *string) {
	if err := j.validateSetSourceHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceHash",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_S3BucketObject)SetWebsiteRedirect(val *string) {
	if err := j.validateSetWebsiteRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"websiteRedirect",
		val,
	)
}

// Generates CDKTF code for importing a S3BucketObject resource upon running "cdktf plan <stack-name>".
func S3BucketObject_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateS3BucketObject_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3BucketObject.S3BucketObject",
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
func S3BucketObject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3BucketObject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3BucketObject.S3BucketObject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func S3BucketObject_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3BucketObject_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3BucketObject.S3BucketObject",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func S3BucketObject_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3BucketObject_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3BucketObject.S3BucketObject",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func S3BucketObject_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.s3BucketObject.S3BucketObject",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_S3BucketObject) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_S3BucketObject) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_S3BucketObject) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3BucketObject) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3BucketObject) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3BucketObject) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3BucketObject) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3BucketObject) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3BucketObject) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3BucketObject) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3BucketObject) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3BucketObject) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketObject) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_S3BucketObject) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3BucketObject) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_S3BucketObject) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_S3BucketObject) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_S3BucketObject) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_S3BucketObject) ResetAcl() {
	_jsii_.InvokeVoid(
		s,
		"resetAcl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetBucketKeyEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetBucketKeyEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetCacheControl() {
	_jsii_.InvokeVoid(
		s,
		"resetCacheControl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetContent() {
	_jsii_.InvokeVoid(
		s,
		"resetContent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetContentBase64() {
	_jsii_.InvokeVoid(
		s,
		"resetContentBase64",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetContentDisposition() {
	_jsii_.InvokeVoid(
		s,
		"resetContentDisposition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetContentEncoding() {
	_jsii_.InvokeVoid(
		s,
		"resetContentEncoding",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetContentLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetContentLanguage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetContentType() {
	_jsii_.InvokeVoid(
		s,
		"resetContentType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetEtag() {
	_jsii_.InvokeVoid(
		s,
		"resetEtag",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		s,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetMetadata() {
	_jsii_.InvokeVoid(
		s,
		"resetMetadata",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetObjectLockLegalHoldStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectLockLegalHoldStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetObjectLockMode() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectLockMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetObjectLockRetainUntilDate() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectLockRetainUntilDate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetServerSideEncryption() {
	_jsii_.InvokeVoid(
		s,
		"resetServerSideEncryption",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetSource() {
	_jsii_.InvokeVoid(
		s,
		"resetSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetSourceHash() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceHash",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetStorageClass() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) ResetWebsiteRedirect() {
	_jsii_.InvokeVoid(
		s,
		"resetWebsiteRedirect",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketObject) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketObject) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketObject) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketObject) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketObject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketObject) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

