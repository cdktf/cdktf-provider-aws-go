package storagegatewaysmbfileshare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v15/storagegatewaysmbfileshare/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/storagegateway_smb_file_share aws_storagegateway_smb_file_share}.
type StoragegatewaySmbFileShare interface {
	cdktf.TerraformResource
	AccessBasedEnumeration() interface{}
	SetAccessBasedEnumeration(val interface{})
	AccessBasedEnumerationInput() interface{}
	AdminUserList() *[]*string
	SetAdminUserList(val *[]*string)
	AdminUserListInput() *[]*string
	Arn() *string
	AuditDestinationArn() *string
	SetAuditDestinationArn(val *string)
	AuditDestinationArnInput() *string
	Authentication() *string
	SetAuthentication(val *string)
	AuthenticationInput() *string
	BucketRegion() *string
	SetBucketRegion(val *string)
	BucketRegionInput() *string
	CacheAttributes() StoragegatewaySmbFileShareCacheAttributesOutputReference
	CacheAttributesInput() *StoragegatewaySmbFileShareCacheAttributes
	CaseSensitivity() *string
	SetCaseSensitivity(val *string)
	CaseSensitivityInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	InvalidUserList() *[]*string
	SetInvalidUserList(val *[]*string)
	InvalidUserListInput() *[]*string
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
	// The tree node.
	Node() constructs.Node
	NotificationPolicy() *string
	SetNotificationPolicy(val *string)
	NotificationPolicyInput() *string
	ObjectAcl() *string
	SetObjectAcl(val *string)
	ObjectAclInput() *string
	OplocksEnabled() interface{}
	SetOplocksEnabled(val interface{})
	OplocksEnabledInput() interface{}
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
	SmbAclEnabled() interface{}
	SetSmbAclEnabled(val interface{})
	SmbAclEnabledInput() interface{}
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
	Timeouts() StoragegatewaySmbFileShareTimeoutsOutputReference
	TimeoutsInput() interface{}
	ValidUserList() *[]*string
	SetValidUserList(val *[]*string)
	ValidUserListInput() *[]*string
	VpcEndpointDnsName() *string
	SetVpcEndpointDnsName(val *string)
	VpcEndpointDnsNameInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCacheAttributes(value *StoragegatewaySmbFileShareCacheAttributes)
	PutTimeouts(value *StoragegatewaySmbFileShareTimeouts)
	ResetAccessBasedEnumeration()
	ResetAdminUserList()
	ResetAuditDestinationArn()
	ResetAuthentication()
	ResetBucketRegion()
	ResetCacheAttributes()
	ResetCaseSensitivity()
	ResetDefaultStorageClass()
	ResetFileShareName()
	ResetGuessMimeTypeEnabled()
	ResetId()
	ResetInvalidUserList()
	ResetKmsEncrypted()
	ResetKmsKeyArn()
	ResetNotificationPolicy()
	ResetObjectAcl()
	ResetOplocksEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReadOnly()
	ResetRequesterPays()
	ResetSmbAclEnabled()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetValidUserList()
	ResetVpcEndpointDnsName()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for StoragegatewaySmbFileShare
type jsiiProxy_StoragegatewaySmbFileShare struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AccessBasedEnumeration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessBasedEnumeration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AccessBasedEnumerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessBasedEnumerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AdminUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AdminUserListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminUserListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AuditDestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AuditDestinationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Authentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) BucketRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) BucketRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) CacheAttributes() StoragegatewaySmbFileShareCacheAttributesOutputReference {
	var returns StoragegatewaySmbFileShareCacheAttributesOutputReference
	_jsii_.Get(
		j,
		"cacheAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) CacheAttributesInput() *StoragegatewaySmbFileShareCacheAttributes {
	var returns *StoragegatewaySmbFileShareCacheAttributes
	_jsii_.Get(
		j,
		"cacheAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) CaseSensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caseSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) CaseSensitivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caseSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) DefaultStorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) DefaultStorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) FileshareId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileshareId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) FileShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) FileShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) GuessMimeTypeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guessMimeTypeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) GuessMimeTypeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guessMimeTypeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) InvalidUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invalidUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) InvalidUserListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invalidUserListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) KmsEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) KmsEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) LocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) LocationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) NotificationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) NotificationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ObjectAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ObjectAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) OplocksEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oplocksEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) OplocksEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oplocksEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) RequesterPays() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) RequesterPaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SmbAclEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbAclEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SmbAclEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbAclEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Timeouts() StoragegatewaySmbFileShareTimeoutsOutputReference {
	var returns StoragegatewaySmbFileShareTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ValidUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ValidUserListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validUserListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) VpcEndpointDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) VpcEndpointDnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointDnsNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/storagegateway_smb_file_share aws_storagegateway_smb_file_share} Resource.
func NewStoragegatewaySmbFileShare(scope constructs.Construct, id *string, config *StoragegatewaySmbFileShareConfig) StoragegatewaySmbFileShare {
	_init_.Initialize()

	if err := validateNewStoragegatewaySmbFileShareParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StoragegatewaySmbFileShare{}

	_jsii_.Create(
		"@cdktf/provider-aws.storagegatewaySmbFileShare.StoragegatewaySmbFileShare",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/storagegateway_smb_file_share aws_storagegateway_smb_file_share} Resource.
func NewStoragegatewaySmbFileShare_Override(s StoragegatewaySmbFileShare, scope constructs.Construct, id *string, config *StoragegatewaySmbFileShareConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.storagegatewaySmbFileShare.StoragegatewaySmbFileShare",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetAccessBasedEnumeration(val interface{}) {
	if err := j.validateSetAccessBasedEnumerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessBasedEnumeration",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetAdminUserList(val *[]*string) {
	if err := j.validateSetAdminUserListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUserList",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetAuditDestinationArn(val *string) {
	if err := j.validateSetAuditDestinationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditDestinationArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetAuthentication(val *string) {
	if err := j.validateSetAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authentication",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetBucketRegion(val *string) {
	if err := j.validateSetBucketRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketRegion",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetCaseSensitivity(val *string) {
	if err := j.validateSetCaseSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caseSensitivity",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetDefaultStorageClass(val *string) {
	if err := j.validateSetDefaultStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultStorageClass",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetFileShareName(val *string) {
	if err := j.validateSetFileShareNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileShareName",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetGatewayArn(val *string) {
	if err := j.validateSetGatewayArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetGuessMimeTypeEnabled(val interface{}) {
	if err := j.validateSetGuessMimeTypeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guessMimeTypeEnabled",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetInvalidUserList(val *[]*string) {
	if err := j.validateSetInvalidUserListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invalidUserList",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetKmsEncrypted(val interface{}) {
	if err := j.validateSetKmsEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsEncrypted",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetLocationArn(val *string) {
	if err := j.validateSetLocationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetNotificationPolicy(val *string) {
	if err := j.validateSetNotificationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationPolicy",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetObjectAcl(val *string) {
	if err := j.validateSetObjectAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectAcl",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetOplocksEnabled(val interface{}) {
	if err := j.validateSetOplocksEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oplocksEnabled",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetRequesterPays(val interface{}) {
	if err := j.validateSetRequesterPaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requesterPays",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetSmbAclEnabled(val interface{}) {
	if err := j.validateSetSmbAclEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbAclEnabled",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetValidUserList(val *[]*string) {
	if err := j.validateSetValidUserListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validUserList",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare)SetVpcEndpointDnsName(val *string) {
	if err := j.validateSetVpcEndpointDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcEndpointDnsName",
		val,
	)
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
func StoragegatewaySmbFileShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStoragegatewaySmbFileShare_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewaySmbFileShare.StoragegatewaySmbFileShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StoragegatewaySmbFileShare_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStoragegatewaySmbFileShare_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewaySmbFileShare.StoragegatewaySmbFileShare",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StoragegatewaySmbFileShare_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStoragegatewaySmbFileShare_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewaySmbFileShare.StoragegatewaySmbFileShare",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewaySmbFileShare_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.storagegatewaySmbFileShare.StoragegatewaySmbFileShare",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StoragegatewaySmbFileShare) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StoragegatewaySmbFileShare) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StoragegatewaySmbFileShare) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StoragegatewaySmbFileShare) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StoragegatewaySmbFileShare) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StoragegatewaySmbFileShare) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StoragegatewaySmbFileShare) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StoragegatewaySmbFileShare) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StoragegatewaySmbFileShare) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StoragegatewaySmbFileShare) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) PutCacheAttributes(value *StoragegatewaySmbFileShareCacheAttributes) {
	if err := s.validatePutCacheAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCacheAttributes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) PutTimeouts(value *StoragegatewaySmbFileShareTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetAccessBasedEnumeration() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessBasedEnumeration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetAdminUserList() {
	_jsii_.InvokeVoid(
		s,
		"resetAdminUserList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetAuditDestinationArn() {
	_jsii_.InvokeVoid(
		s,
		"resetAuditDestinationArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetAuthentication() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetBucketRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetBucketRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetCacheAttributes() {
	_jsii_.InvokeVoid(
		s,
		"resetCacheAttributes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetCaseSensitivity() {
	_jsii_.InvokeVoid(
		s,
		"resetCaseSensitivity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetDefaultStorageClass() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultStorageClass",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetFileShareName() {
	_jsii_.InvokeVoid(
		s,
		"resetFileShareName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetGuessMimeTypeEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetGuessMimeTypeEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetInvalidUserList() {
	_jsii_.InvokeVoid(
		s,
		"resetInvalidUserList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetKmsEncrypted() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsEncrypted",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetNotificationPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetObjectAcl() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectAcl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetOplocksEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetOplocksEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetReadOnly() {
	_jsii_.InvokeVoid(
		s,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetRequesterPays() {
	_jsii_.InvokeVoid(
		s,
		"resetRequesterPays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetSmbAclEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSmbAclEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetValidUserList() {
	_jsii_.InvokeVoid(
		s,
		"resetValidUserList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetVpcEndpointDnsName() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcEndpointDnsName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

