// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ami

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/ami/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/ami aws_ami}.
type Ami interface {
	cdktf.TerraformResource
	Architecture() *string
	SetArchitecture(val *string)
	ArchitectureInput() *string
	Arn() *string
	BootMode() *string
	SetBootMode(val *string)
	BootModeInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeprecationTime() *string
	SetDeprecationTime(val *string)
	DeprecationTimeInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EbsBlockDevice() AmiEbsBlockDeviceList
	EbsBlockDeviceInput() interface{}
	EnaSupport() interface{}
	SetEnaSupport(val interface{})
	EnaSupportInput() interface{}
	EphemeralBlockDevice() AmiEphemeralBlockDeviceList
	EphemeralBlockDeviceInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hypervisor() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageLocation() *string
	SetImageLocation(val *string)
	ImageLocationInput() *string
	ImageOwnerAlias() *string
	ImageType() *string
	ImdsSupport() *string
	SetImdsSupport(val *string)
	ImdsSupportInput() *string
	KernelId() *string
	SetKernelId(val *string)
	KernelIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManageEbsSnapshots() cdktf.IResolvable
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OwnerId() *string
	Platform() *string
	PlatformDetails() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Public() cdktf.IResolvable
	RamdiskId() *string
	SetRamdiskId(val *string)
	RamdiskIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	RootDeviceName() *string
	SetRootDeviceName(val *string)
	RootDeviceNameInput() *string
	RootSnapshotId() *string
	SriovNetSupport() *string
	SetSriovNetSupport(val *string)
	SriovNetSupportInput() *string
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
	Timeouts() AmiTimeoutsOutputReference
	TimeoutsInput() interface{}
	TpmSupport() *string
	SetTpmSupport(val *string)
	TpmSupportInput() *string
	UsageOperation() *string
	VirtualizationType() *string
	SetVirtualizationType(val *string)
	VirtualizationTypeInput() *string
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
	PutEbsBlockDevice(value interface{})
	PutEphemeralBlockDevice(value interface{})
	PutTimeouts(value *AmiTimeouts)
	ResetArchitecture()
	ResetBootMode()
	ResetDeprecationTime()
	ResetDescription()
	ResetEbsBlockDevice()
	ResetEnaSupport()
	ResetEphemeralBlockDevice()
	ResetId()
	ResetImageLocation()
	ResetImdsSupport()
	ResetKernelId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRamdiskId()
	ResetRootDeviceName()
	ResetSriovNetSupport()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetTpmSupport()
	ResetVirtualizationType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Ami
type jsiiProxy_Ami struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Ami) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) BootMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) BootModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) DeprecationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprecationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) DeprecationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprecationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) EbsBlockDevice() AmiEbsBlockDeviceList {
	var returns AmiEbsBlockDeviceList
	_jsii_.Get(
		j,
		"ebsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) EbsBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) EnaSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enaSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) EnaSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enaSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) EphemeralBlockDevice() AmiEphemeralBlockDeviceList {
	var returns AmiEphemeralBlockDeviceList
	_jsii_.Get(
		j,
		"ephemeralBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) EphemeralBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Hypervisor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hypervisor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) ImageLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) ImageLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) ImageOwnerAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageOwnerAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) ImdsSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imdsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) ImdsSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imdsSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) KernelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) ManageEbsSnapshots() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"manageEbsSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) PlatformDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Public() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"public",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) RamdiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramdiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) RamdiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramdiskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) RootDeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDeviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) RootDeviceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDeviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) RootSnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootSnapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) SriovNetSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sriovNetSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) SriovNetSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sriovNetSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) Timeouts() AmiTimeoutsOutputReference {
	var returns AmiTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) TpmSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpmSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) TpmSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpmSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) UsageOperation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) VirtualizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ami) VirtualizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualizationTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/ami aws_ami} Resource.
func NewAmi(scope constructs.Construct, id *string, config *AmiConfig) Ami {
	_init_.Initialize()

	if err := validateNewAmiParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ami{}

	_jsii_.Create(
		"@cdktf/provider-aws.ami.Ami",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/ami aws_ami} Resource.
func NewAmi_Override(a Ami, scope constructs.Construct, id *string, config *AmiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ami.Ami",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Ami)SetArchitecture(val *string) {
	if err := j.validateSetArchitectureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_Ami)SetBootMode(val *string) {
	if err := j.validateSetBootModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootMode",
		val,
	)
}

func (j *jsiiProxy_Ami)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ami)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ami)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ami)SetDeprecationTime(val *string) {
	if err := j.validateSetDeprecationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deprecationTime",
		val,
	)
}

func (j *jsiiProxy_Ami)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Ami)SetEnaSupport(val interface{}) {
	if err := j.validateSetEnaSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enaSupport",
		val,
	)
}

func (j *jsiiProxy_Ami)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ami)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Ami)SetImageLocation(val *string) {
	if err := j.validateSetImageLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageLocation",
		val,
	)
}

func (j *jsiiProxy_Ami)SetImdsSupport(val *string) {
	if err := j.validateSetImdsSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imdsSupport",
		val,
	)
}

func (j *jsiiProxy_Ami)SetKernelId(val *string) {
	if err := j.validateSetKernelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kernelId",
		val,
	)
}

func (j *jsiiProxy_Ami)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ami)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Ami)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ami)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Ami)SetRamdiskId(val *string) {
	if err := j.validateSetRamdiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ramdiskId",
		val,
	)
}

func (j *jsiiProxy_Ami)SetRootDeviceName(val *string) {
	if err := j.validateSetRootDeviceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootDeviceName",
		val,
	)
}

func (j *jsiiProxy_Ami)SetSriovNetSupport(val *string) {
	if err := j.validateSetSriovNetSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sriovNetSupport",
		val,
	)
}

func (j *jsiiProxy_Ami)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Ami)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_Ami)SetTpmSupport(val *string) {
	if err := j.validateSetTpmSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpmSupport",
		val,
	)
}

func (j *jsiiProxy_Ami)SetVirtualizationType(val *string) {
	if err := j.validateSetVirtualizationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualizationType",
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
func Ami_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ami.Ami",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ami_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmi_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ami.Ami",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ami_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmi_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ami.Ami",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Ami_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.ami.Ami",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_Ami) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_Ami) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Ami) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Ami) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Ami) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Ami) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Ami) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Ami) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Ami) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Ami) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Ami) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Ami) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Ami) PutEbsBlockDevice(value interface{}) {
	if err := a.validatePutEbsBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEbsBlockDevice",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Ami) PutEphemeralBlockDevice(value interface{}) {
	if err := a.validatePutEphemeralBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEphemeralBlockDevice",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Ami) PutTimeouts(value *AmiTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Ami) ResetArchitecture() {
	_jsii_.InvokeVoid(
		a,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetBootMode() {
	_jsii_.InvokeVoid(
		a,
		"resetBootMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetDeprecationTime() {
	_jsii_.InvokeVoid(
		a,
		"resetDeprecationTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetEbsBlockDevice() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsBlockDevice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetEnaSupport() {
	_jsii_.InvokeVoid(
		a,
		"resetEnaSupport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetEphemeralBlockDevice() {
	_jsii_.InvokeVoid(
		a,
		"resetEphemeralBlockDevice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetImageLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetImageLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetImdsSupport() {
	_jsii_.InvokeVoid(
		a,
		"resetImdsSupport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetKernelId() {
	_jsii_.InvokeVoid(
		a,
		"resetKernelId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetRamdiskId() {
	_jsii_.InvokeVoid(
		a,
		"resetRamdiskId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetRootDeviceName() {
	_jsii_.InvokeVoid(
		a,
		"resetRootDeviceName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetSriovNetSupport() {
	_jsii_.InvokeVoid(
		a,
		"resetSriovNetSupport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetTpmSupport() {
	_jsii_.InvokeVoid(
		a,
		"resetTpmSupport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) ResetVirtualizationType() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualizationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Ami) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Ami) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Ami) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Ami) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

