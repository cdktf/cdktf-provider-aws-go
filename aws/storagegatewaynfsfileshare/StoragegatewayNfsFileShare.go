// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagegatewaynfsfileshare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/storagegatewaynfsfileshare/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/storagegateway_nfs_file_share aws_storagegateway_nfs_file_share}.
type StoragegatewayNfsFileShare interface {
	cdktf.TerraformResource
	Arn() *string
	AuditDestinationArn() *string
	SetAuditDestinationArn(val *string)
	AuditDestinationArnInput() *string
	BucketRegion() *string
	SetBucketRegion(val *string)
	BucketRegionInput() *string
	CacheAttributes() StoragegatewayNfsFileShareCacheAttributesOutputReference
	CacheAttributesInput() *StoragegatewayNfsFileShareCacheAttributes
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientList() *[]*string
	SetClientList(val *[]*string)
	ClientListInput() *[]*string
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
	DefaultStorageClass() *string
	SetDefaultStorageClass(val *string)
	DefaultStorageClassInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FileshareId() *string
	FileShareName() *string
	SetFileShareName(val *string)
	FileShareNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayArn() *string
	SetGatewayArn(val *string)
	GatewayArnInput() *string
	GuessMimeTypeEnabled() interface{}
	SetGuessMimeTypeEnabled(val interface{})
	GuessMimeTypeEnabledInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	KmsEncrypted() interface{}
	SetKmsEncrypted(val interface{})
	KmsEncryptedInput() interface{}
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocationArn() *string
	SetLocationArn(val *string)
	LocationArnInput() *string
	NfsFileShareDefaults() StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference
	NfsFileShareDefaultsInput() *StoragegatewayNfsFileShareNfsFileShareDefaults
	// The tree node.
	Node() constructs.Node
	NotificationPolicy() *string
	SetNotificationPolicy(val *string)
	NotificationPolicyInput() *string
	ObjectAcl() *string
	SetObjectAcl(val *string)
	ObjectAclInput() *string
	Path() *string
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
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	RequesterPays() interface{}
	SetRequesterPays(val interface{})
	RequesterPaysInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Squash() *string
	SetSquash(val *string)
	SquashInput() *string
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
	Timeouts() StoragegatewayNfsFileShareTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcEndpointDnsName() *string
	SetVpcEndpointDnsName(val *string)
	VpcEndpointDnsNameInput() *string
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
	PutCacheAttributes(value *StoragegatewayNfsFileShareCacheAttributes)
	PutNfsFileShareDefaults(value *StoragegatewayNfsFileShareNfsFileShareDefaults)
	PutTimeouts(value *StoragegatewayNfsFileShareTimeouts)
	ResetAuditDestinationArn()
	ResetBucketRegion()
	ResetCacheAttributes()
	ResetDefaultStorageClass()
	ResetFileShareName()
	ResetGuessMimeTypeEnabled()
	ResetId()
	ResetKmsEncrypted()
	ResetKmsKeyArn()
	ResetNfsFileShareDefaults()
	ResetNotificationPolicy()
	ResetObjectAcl()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReadOnly()
	ResetRequesterPays()
	ResetSquash()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetVpcEndpointDnsName()
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

// The jsii proxy struct for StoragegatewayNfsFileShare
type jsiiProxy_StoragegatewayNfsFileShare struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) AuditDestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) AuditDestinationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) BucketRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) BucketRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) CacheAttributes() StoragegatewayNfsFileShareCacheAttributesOutputReference {
	var returns StoragegatewayNfsFileShareCacheAttributesOutputReference
	_jsii_.Get(
		j,
		"cacheAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) CacheAttributesInput() *StoragegatewayNfsFileShareCacheAttributes {
	var returns *StoragegatewayNfsFileShareCacheAttributes
	_jsii_.Get(
		j,
		"cacheAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ClientList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ClientListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) DefaultStorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) DefaultStorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) FileshareId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileshareId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) FileShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) FileShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) GuessMimeTypeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guessMimeTypeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) GuessMimeTypeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guessMimeTypeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) KmsEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) KmsEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) LocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) LocationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) NfsFileShareDefaults() StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference {
	var returns StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference
	_jsii_.Get(
		j,
		"nfsFileShareDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) NfsFileShareDefaultsInput() *StoragegatewayNfsFileShareNfsFileShareDefaults {
	var returns *StoragegatewayNfsFileShareNfsFileShareDefaults
	_jsii_.Get(
		j,
		"nfsFileShareDefaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) NotificationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) NotificationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ObjectAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ObjectAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) RequesterPays() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) RequesterPaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Squash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SquashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Timeouts() StoragegatewayNfsFileShareTimeoutsOutputReference {
	var returns StoragegatewayNfsFileShareTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) VpcEndpointDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) VpcEndpointDnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointDnsNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/storagegateway_nfs_file_share aws_storagegateway_nfs_file_share} Resource.
func NewStoragegatewayNfsFileShare(scope constructs.Construct, id *string, config *StoragegatewayNfsFileShareConfig) StoragegatewayNfsFileShare {
	_init_.Initialize()

	if err := validateNewStoragegatewayNfsFileShareParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StoragegatewayNfsFileShare{}

	_jsii_.Create(
		"@cdktf/provider-aws.storagegatewayNfsFileShare.StoragegatewayNfsFileShare",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/storagegateway_nfs_file_share aws_storagegateway_nfs_file_share} Resource.
func NewStoragegatewayNfsFileShare_Override(s StoragegatewayNfsFileShare, scope constructs.Construct, id *string, config *StoragegatewayNfsFileShareConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.storagegatewayNfsFileShare.StoragegatewayNfsFileShare",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetAuditDestinationArn(val *string) {
	if err := j.validateSetAuditDestinationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditDestinationArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetBucketRegion(val *string) {
	if err := j.validateSetBucketRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketRegion",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetClientList(val *[]*string) {
	if err := j.validateSetClientListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientList",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetDefaultStorageClass(val *string) {
	if err := j.validateSetDefaultStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultStorageClass",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetFileShareName(val *string) {
	if err := j.validateSetFileShareNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileShareName",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetGatewayArn(val *string) {
	if err := j.validateSetGatewayArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetGuessMimeTypeEnabled(val interface{}) {
	if err := j.validateSetGuessMimeTypeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guessMimeTypeEnabled",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetKmsEncrypted(val interface{}) {
	if err := j.validateSetKmsEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsEncrypted",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetLocationArn(val *string) {
	if err := j.validateSetLocationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetNotificationPolicy(val *string) {
	if err := j.validateSetNotificationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationPolicy",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetObjectAcl(val *string) {
	if err := j.validateSetObjectAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectAcl",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetRequesterPays(val interface{}) {
	if err := j.validateSetRequesterPaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requesterPays",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetSquash(val *string) {
	if err := j.validateSetSquashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"squash",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare)SetVpcEndpointDnsName(val *string) {
	if err := j.validateSetVpcEndpointDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcEndpointDnsName",
		val,
	)
}

// Generates CDKTF code for importing a StoragegatewayNfsFileShare resource upon running "cdktf plan <stack-name>".
func StoragegatewayNfsFileShare_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStoragegatewayNfsFileShare_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewayNfsFileShare.StoragegatewayNfsFileShare",
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
func StoragegatewayNfsFileShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStoragegatewayNfsFileShare_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewayNfsFileShare.StoragegatewayNfsFileShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StoragegatewayNfsFileShare_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStoragegatewayNfsFileShare_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewayNfsFileShare.StoragegatewayNfsFileShare",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StoragegatewayNfsFileShare_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStoragegatewayNfsFileShare_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewayNfsFileShare.StoragegatewayNfsFileShare",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewayNfsFileShare_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.storagegatewayNfsFileShare.StoragegatewayNfsFileShare",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StoragegatewayNfsFileShare) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StoragegatewayNfsFileShare) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StoragegatewayNfsFileShare) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StoragegatewayNfsFileShare) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StoragegatewayNfsFileShare) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StoragegatewayNfsFileShare) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StoragegatewayNfsFileShare) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StoragegatewayNfsFileShare) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StoragegatewayNfsFileShare) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StoragegatewayNfsFileShare) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) PutCacheAttributes(value *StoragegatewayNfsFileShareCacheAttributes) {
	if err := s.validatePutCacheAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCacheAttributes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) PutNfsFileShareDefaults(value *StoragegatewayNfsFileShareNfsFileShareDefaults) {
	if err := s.validatePutNfsFileShareDefaultsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNfsFileShareDefaults",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) PutTimeouts(value *StoragegatewayNfsFileShareTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetAuditDestinationArn() {
	_jsii_.InvokeVoid(
		s,
		"resetAuditDestinationArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetBucketRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetBucketRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetCacheAttributes() {
	_jsii_.InvokeVoid(
		s,
		"resetCacheAttributes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetDefaultStorageClass() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultStorageClass",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetFileShareName() {
	_jsii_.InvokeVoid(
		s,
		"resetFileShareName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetGuessMimeTypeEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetGuessMimeTypeEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetKmsEncrypted() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsEncrypted",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetNfsFileShareDefaults() {
	_jsii_.InvokeVoid(
		s,
		"resetNfsFileShareDefaults",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetNotificationPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetObjectAcl() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectAcl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetReadOnly() {
	_jsii_.InvokeVoid(
		s,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetRequesterPays() {
	_jsii_.InvokeVoid(
		s,
		"resetRequesterPays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetSquash() {
	_jsii_.InvokeVoid(
		s,
		"resetSquash",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetVpcEndpointDnsName() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcEndpointDnsName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

