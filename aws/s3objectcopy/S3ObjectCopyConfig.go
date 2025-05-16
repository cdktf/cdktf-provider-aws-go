// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3objectcopy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3ObjectCopyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#bucket S3ObjectCopy#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#key S3ObjectCopy#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#source S3ObjectCopy#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#acl S3ObjectCopy#acl}.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#bucket_key_enabled S3ObjectCopy#bucket_key_enabled}.
	BucketKeyEnabled interface{} `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#cache_control S3ObjectCopy#cache_control}.
	CacheControl *string `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#checksum_algorithm S3ObjectCopy#checksum_algorithm}.
	ChecksumAlgorithm *string `field:"optional" json:"checksumAlgorithm" yaml:"checksumAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#content_disposition S3ObjectCopy#content_disposition}.
	ContentDisposition *string `field:"optional" json:"contentDisposition" yaml:"contentDisposition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#content_encoding S3ObjectCopy#content_encoding}.
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#content_language S3ObjectCopy#content_language}.
	ContentLanguage *string `field:"optional" json:"contentLanguage" yaml:"contentLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#content_type S3ObjectCopy#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#copy_if_match S3ObjectCopy#copy_if_match}.
	CopyIfMatch *string `field:"optional" json:"copyIfMatch" yaml:"copyIfMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#copy_if_modified_since S3ObjectCopy#copy_if_modified_since}.
	CopyIfModifiedSince *string `field:"optional" json:"copyIfModifiedSince" yaml:"copyIfModifiedSince"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#copy_if_none_match S3ObjectCopy#copy_if_none_match}.
	CopyIfNoneMatch *string `field:"optional" json:"copyIfNoneMatch" yaml:"copyIfNoneMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#copy_if_unmodified_since S3ObjectCopy#copy_if_unmodified_since}.
	CopyIfUnmodifiedSince *string `field:"optional" json:"copyIfUnmodifiedSince" yaml:"copyIfUnmodifiedSince"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#customer_algorithm S3ObjectCopy#customer_algorithm}.
	CustomerAlgorithm *string `field:"optional" json:"customerAlgorithm" yaml:"customerAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#customer_key S3ObjectCopy#customer_key}.
	CustomerKey *string `field:"optional" json:"customerKey" yaml:"customerKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#customer_key_md5 S3ObjectCopy#customer_key_md5}.
	CustomerKeyMd5 *string `field:"optional" json:"customerKeyMd5" yaml:"customerKeyMd5"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#expected_bucket_owner S3ObjectCopy#expected_bucket_owner}.
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#expected_source_bucket_owner S3ObjectCopy#expected_source_bucket_owner}.
	ExpectedSourceBucketOwner *string `field:"optional" json:"expectedSourceBucketOwner" yaml:"expectedSourceBucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#expires S3ObjectCopy#expires}.
	Expires *string `field:"optional" json:"expires" yaml:"expires"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#force_destroy S3ObjectCopy#force_destroy}.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// grant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#grant S3ObjectCopy#grant}
	Grant interface{} `field:"optional" json:"grant" yaml:"grant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#id S3ObjectCopy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#kms_encryption_context S3ObjectCopy#kms_encryption_context}.
	KmsEncryptionContext *string `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#kms_key_id S3ObjectCopy#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#metadata S3ObjectCopy#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#metadata_directive S3ObjectCopy#metadata_directive}.
	MetadataDirective *string `field:"optional" json:"metadataDirective" yaml:"metadataDirective"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#object_lock_legal_hold_status S3ObjectCopy#object_lock_legal_hold_status}.
	ObjectLockLegalHoldStatus *string `field:"optional" json:"objectLockLegalHoldStatus" yaml:"objectLockLegalHoldStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#object_lock_mode S3ObjectCopy#object_lock_mode}.
	ObjectLockMode *string `field:"optional" json:"objectLockMode" yaml:"objectLockMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#object_lock_retain_until_date S3ObjectCopy#object_lock_retain_until_date}.
	ObjectLockRetainUntilDate *string `field:"optional" json:"objectLockRetainUntilDate" yaml:"objectLockRetainUntilDate"`
	// override_provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#override_provider S3ObjectCopy#override_provider}
	OverrideProvider *S3ObjectCopyOverrideProvider `field:"optional" json:"overrideProvider" yaml:"overrideProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#request_payer S3ObjectCopy#request_payer}.
	RequestPayer *string `field:"optional" json:"requestPayer" yaml:"requestPayer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#server_side_encryption S3ObjectCopy#server_side_encryption}.
	ServerSideEncryption *string `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#source_customer_algorithm S3ObjectCopy#source_customer_algorithm}.
	SourceCustomerAlgorithm *string `field:"optional" json:"sourceCustomerAlgorithm" yaml:"sourceCustomerAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#source_customer_key S3ObjectCopy#source_customer_key}.
	SourceCustomerKey *string `field:"optional" json:"sourceCustomerKey" yaml:"sourceCustomerKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#source_customer_key_md5 S3ObjectCopy#source_customer_key_md5}.
	SourceCustomerKeyMd5 *string `field:"optional" json:"sourceCustomerKeyMd5" yaml:"sourceCustomerKeyMd5"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#storage_class S3ObjectCopy#storage_class}.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#tagging_directive S3ObjectCopy#tagging_directive}.
	TaggingDirective *string `field:"optional" json:"taggingDirective" yaml:"taggingDirective"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#tags S3ObjectCopy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#tags_all S3ObjectCopy#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/s3_object_copy#website_redirect S3ObjectCopy#website_redirect}.
	WebsiteRedirect *string `field:"optional" json:"websiteRedirect" yaml:"websiteRedirect"`
}

