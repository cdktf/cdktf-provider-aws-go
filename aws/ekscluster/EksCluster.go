// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ekscluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_cluster aws_eks_cluster}.
type EksCluster interface {
	cdktf.TerraformResource
	AccessConfig() EksClusterAccessConfigOutputReference
	AccessConfigInput() *EksClusterAccessConfig
	Arn() *string
	BootstrapSelfManagedAddons() interface{}
	SetBootstrapSelfManagedAddons(val interface{})
	BootstrapSelfManagedAddonsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateAuthority() EksClusterCertificateAuthorityList
	ClusterId() *string
	ComputeConfig() EksClusterComputeConfigOutputReference
	ComputeConfigInput() *EksClusterComputeConfig
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPlaneScalingConfig() EksClusterControlPlaneScalingConfigOutputReference
	ControlPlaneScalingConfigInput() *EksClusterControlPlaneScalingConfig
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnabledClusterLogTypes() *[]*string
	SetEnabledClusterLogTypes(val *[]*string)
	EnabledClusterLogTypesInput() *[]*string
	EncryptionConfig() EksClusterEncryptionConfigOutputReference
	EncryptionConfigInput() *EksClusterEncryptionConfig
	Endpoint() *string
	ForceUpdateVersion() interface{}
	SetForceUpdateVersion(val interface{})
	ForceUpdateVersionInput() interface{}
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
	Identity() EksClusterIdentityList
	IdInput() *string
	KubernetesNetworkConfig() EksClusterKubernetesNetworkConfigOutputReference
	KubernetesNetworkConfigInput() *EksClusterKubernetesNetworkConfig
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutpostConfig() EksClusterOutpostConfigOutputReference
	OutpostConfigInput() *EksClusterOutpostConfig
	PlatformVersion() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RemoteNetworkConfig() EksClusterRemoteNetworkConfigOutputReference
	RemoteNetworkConfigInput() *EksClusterRemoteNetworkConfig
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Status() *string
	StorageConfig() EksClusterStorageConfigOutputReference
	StorageConfigInput() *EksClusterStorageConfig
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
	Timeouts() EksClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpgradePolicy() EksClusterUpgradePolicyOutputReference
	UpgradePolicyInput() *EksClusterUpgradePolicy
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	VpcConfig() EksClusterVpcConfigOutputReference
	VpcConfigInput() *EksClusterVpcConfig
	ZonalShiftConfig() EksClusterZonalShiftConfigOutputReference
	ZonalShiftConfigInput() *EksClusterZonalShiftConfig
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
	PutAccessConfig(value *EksClusterAccessConfig)
	PutComputeConfig(value *EksClusterComputeConfig)
	PutControlPlaneScalingConfig(value *EksClusterControlPlaneScalingConfig)
	PutEncryptionConfig(value *EksClusterEncryptionConfig)
	PutKubernetesNetworkConfig(value *EksClusterKubernetesNetworkConfig)
	PutOutpostConfig(value *EksClusterOutpostConfig)
	PutRemoteNetworkConfig(value *EksClusterRemoteNetworkConfig)
	PutStorageConfig(value *EksClusterStorageConfig)
	PutTimeouts(value *EksClusterTimeouts)
	PutUpgradePolicy(value *EksClusterUpgradePolicy)
	PutVpcConfig(value *EksClusterVpcConfig)
	PutZonalShiftConfig(value *EksClusterZonalShiftConfig)
	ResetAccessConfig()
	ResetBootstrapSelfManagedAddons()
	ResetComputeConfig()
	ResetControlPlaneScalingConfig()
	ResetDeletionProtection()
	ResetEnabledClusterLogTypes()
	ResetEncryptionConfig()
	ResetForceUpdateVersion()
	ResetId()
	ResetKubernetesNetworkConfig()
	ResetOutpostConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetRemoteNetworkConfig()
	ResetStorageConfig()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetUpgradePolicy()
	ResetVersion()
	ResetZonalShiftConfig()
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

// The jsii proxy struct for EksCluster
type jsiiProxy_EksCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EksCluster) AccessConfig() EksClusterAccessConfigOutputReference {
	var returns EksClusterAccessConfigOutputReference
	_jsii_.Get(
		j,
		"accessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) AccessConfigInput() *EksClusterAccessConfig {
	var returns *EksClusterAccessConfig
	_jsii_.Get(
		j,
		"accessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) BootstrapSelfManagedAddons() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapSelfManagedAddons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) BootstrapSelfManagedAddonsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapSelfManagedAddonsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) CertificateAuthority() EksClusterCertificateAuthorityList {
	var returns EksClusterCertificateAuthorityList
	_jsii_.Get(
		j,
		"certificateAuthority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ComputeConfig() EksClusterComputeConfigOutputReference {
	var returns EksClusterComputeConfigOutputReference
	_jsii_.Get(
		j,
		"computeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ComputeConfigInput() *EksClusterComputeConfig {
	var returns *EksClusterComputeConfig
	_jsii_.Get(
		j,
		"computeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ControlPlaneScalingConfig() EksClusterControlPlaneScalingConfigOutputReference {
	var returns EksClusterControlPlaneScalingConfigOutputReference
	_jsii_.Get(
		j,
		"controlPlaneScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ControlPlaneScalingConfigInput() *EksClusterControlPlaneScalingConfig {
	var returns *EksClusterControlPlaneScalingConfig
	_jsii_.Get(
		j,
		"controlPlaneScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) EnabledClusterLogTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledClusterLogTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) EnabledClusterLogTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledClusterLogTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) EncryptionConfig() EksClusterEncryptionConfigOutputReference {
	var returns EksClusterEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) EncryptionConfigInput() *EksClusterEncryptionConfig {
	var returns *EksClusterEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ForceUpdateVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ForceUpdateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Identity() EksClusterIdentityList {
	var returns EksClusterIdentityList
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) KubernetesNetworkConfig() EksClusterKubernetesNetworkConfigOutputReference {
	var returns EksClusterKubernetesNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"kubernetesNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) KubernetesNetworkConfigInput() *EksClusterKubernetesNetworkConfig {
	var returns *EksClusterKubernetesNetworkConfig
	_jsii_.Get(
		j,
		"kubernetesNetworkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) OutpostConfig() EksClusterOutpostConfigOutputReference {
	var returns EksClusterOutpostConfigOutputReference
	_jsii_.Get(
		j,
		"outpostConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) OutpostConfigInput() *EksClusterOutpostConfig {
	var returns *EksClusterOutpostConfig
	_jsii_.Get(
		j,
		"outpostConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) RemoteNetworkConfig() EksClusterRemoteNetworkConfigOutputReference {
	var returns EksClusterRemoteNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"remoteNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) RemoteNetworkConfigInput() *EksClusterRemoteNetworkConfig {
	var returns *EksClusterRemoteNetworkConfig
	_jsii_.Get(
		j,
		"remoteNetworkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) StorageConfig() EksClusterStorageConfigOutputReference {
	var returns EksClusterStorageConfigOutputReference
	_jsii_.Get(
		j,
		"storageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) StorageConfigInput() *EksClusterStorageConfig {
	var returns *EksClusterStorageConfig
	_jsii_.Get(
		j,
		"storageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Timeouts() EksClusterTimeoutsOutputReference {
	var returns EksClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) UpgradePolicy() EksClusterUpgradePolicyOutputReference {
	var returns EksClusterUpgradePolicyOutputReference
	_jsii_.Get(
		j,
		"upgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) UpgradePolicyInput() *EksClusterUpgradePolicy {
	var returns *EksClusterUpgradePolicy
	_jsii_.Get(
		j,
		"upgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) VpcConfig() EksClusterVpcConfigOutputReference {
	var returns EksClusterVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) VpcConfigInput() *EksClusterVpcConfig {
	var returns *EksClusterVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ZonalShiftConfig() EksClusterZonalShiftConfigOutputReference {
	var returns EksClusterZonalShiftConfigOutputReference
	_jsii_.Get(
		j,
		"zonalShiftConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ZonalShiftConfigInput() *EksClusterZonalShiftConfig {
	var returns *EksClusterZonalShiftConfig
	_jsii_.Get(
		j,
		"zonalShiftConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_cluster aws_eks_cluster} Resource.
func NewEksCluster(scope constructs.Construct, id *string, config *EksClusterConfig) EksCluster {
	_init_.Initialize()

	if err := validateNewEksClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EksCluster{}

	_jsii_.Create(
		"@cdktf/provider-aws.eksCluster.EksCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_cluster aws_eks_cluster} Resource.
func NewEksCluster_Override(e EksCluster, scope constructs.Construct, id *string, config *EksClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.eksCluster.EksCluster",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EksCluster)SetBootstrapSelfManagedAddons(val interface{}) {
	if err := j.validateSetBootstrapSelfManagedAddonsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootstrapSelfManagedAddons",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetEnabledClusterLogTypes(val *[]*string) {
	if err := j.validateSetEnabledClusterLogTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledClusterLogTypes",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetForceUpdateVersion(val interface{}) {
	if err := j.validateSetForceUpdateVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdateVersion",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a EksCluster resource upon running "cdktf plan <stack-name>".
func EksCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEksCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.eksCluster.EksCluster",
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
func EksCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEksCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.eksCluster.EksCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EksCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEksCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.eksCluster.EksCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EksCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEksCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.eksCluster.EksCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EksCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.eksCluster.EksCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EksCluster) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EksCluster) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EksCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EksCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EksCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EksCluster) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EksCluster) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EksCluster) PutAccessConfig(value *EksClusterAccessConfig) {
	if err := e.validatePutAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAccessConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutComputeConfig(value *EksClusterComputeConfig) {
	if err := e.validatePutComputeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putComputeConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutControlPlaneScalingConfig(value *EksClusterControlPlaneScalingConfig) {
	if err := e.validatePutControlPlaneScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putControlPlaneScalingConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutEncryptionConfig(value *EksClusterEncryptionConfig) {
	if err := e.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutKubernetesNetworkConfig(value *EksClusterKubernetesNetworkConfig) {
	if err := e.validatePutKubernetesNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putKubernetesNetworkConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutOutpostConfig(value *EksClusterOutpostConfig) {
	if err := e.validatePutOutpostConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOutpostConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutRemoteNetworkConfig(value *EksClusterRemoteNetworkConfig) {
	if err := e.validatePutRemoteNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRemoteNetworkConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutStorageConfig(value *EksClusterStorageConfig) {
	if err := e.validatePutStorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStorageConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutTimeouts(value *EksClusterTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutUpgradePolicy(value *EksClusterUpgradePolicy) {
	if err := e.validatePutUpgradePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putUpgradePolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutVpcConfig(value *EksClusterVpcConfig) {
	if err := e.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutZonalShiftConfig(value *EksClusterZonalShiftConfig) {
	if err := e.validatePutZonalShiftConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putZonalShiftConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) ResetAccessConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetAccessConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetBootstrapSelfManagedAddons() {
	_jsii_.InvokeVoid(
		e,
		"resetBootstrapSelfManagedAddons",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetComputeConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetComputeConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetControlPlaneScalingConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetControlPlaneScalingConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		e,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetEnabledClusterLogTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetEnabledClusterLogTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetForceUpdateVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetForceUpdateVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetKubernetesNetworkConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetKubernetesNetworkConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetOutpostConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetOutpostConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetRemoteNetworkConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetRemoteNetworkConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetStorageConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetStorageConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetUpgradePolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetUpgradePolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetZonalShiftConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetZonalShiftConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

