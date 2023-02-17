package emrcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/emrcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster aws_emr_cluster}.
type EmrCluster interface {
	cdktf.TerraformResource
	AdditionalInfo() *string
	SetAdditionalInfo(val *string)
	AdditionalInfoInput() *string
	Applications() *[]*string
	SetApplications(val *[]*string)
	ApplicationsInput() *[]*string
	Arn() *string
	AutoscalingRole() *string
	SetAutoscalingRole(val *string)
	AutoscalingRoleInput() *string
	AutoTerminationPolicy() EmrClusterAutoTerminationPolicyOutputReference
	AutoTerminationPolicyInput() *EmrClusterAutoTerminationPolicy
	BootstrapAction() EmrClusterBootstrapActionList
	BootstrapActionInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterState() *string
	Configurations() *string
	SetConfigurations(val *string)
	ConfigurationsInput() *string
	ConfigurationsJson() *string
	SetConfigurationsJson(val *string)
	ConfigurationsJsonInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CoreInstanceFleet() EmrClusterCoreInstanceFleetOutputReference
	CoreInstanceFleetInput() *EmrClusterCoreInstanceFleet
	CoreInstanceGroup() EmrClusterCoreInstanceGroupOutputReference
	CoreInstanceGroupInput() *EmrClusterCoreInstanceGroup
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CustomAmiId() *string
	SetCustomAmiId(val *string)
	CustomAmiIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EbsRootVolumeSize() *float64
	SetEbsRootVolumeSize(val *float64)
	EbsRootVolumeSizeInput() *float64
	Ec2Attributes() EmrClusterEc2AttributesOutputReference
	Ec2AttributesInput() *EmrClusterEc2Attributes
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
	KeepJobFlowAliveWhenNoSteps() interface{}
	SetKeepJobFlowAliveWhenNoSteps(val interface{})
	KeepJobFlowAliveWhenNoStepsInput() interface{}
	KerberosAttributes() EmrClusterKerberosAttributesOutputReference
	KerberosAttributesInput() *EmrClusterKerberosAttributes
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListStepsStates() *[]*string
	SetListStepsStates(val *[]*string)
	ListStepsStatesInput() *[]*string
	LogEncryptionKmsKeyId() *string
	SetLogEncryptionKmsKeyId(val *string)
	LogEncryptionKmsKeyIdInput() *string
	LogUri() *string
	SetLogUri(val *string)
	LogUriInput() *string
	MasterInstanceFleet() EmrClusterMasterInstanceFleetOutputReference
	MasterInstanceFleetInput() *EmrClusterMasterInstanceFleet
	MasterInstanceGroup() EmrClusterMasterInstanceGroupOutputReference
	MasterInstanceGroupInput() *EmrClusterMasterInstanceGroup
	MasterPublicDns() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	ReleaseLabel() *string
	SetReleaseLabel(val *string)
	ReleaseLabelInput() *string
	ScaleDownBehavior() *string
	SetScaleDownBehavior(val *string)
	ScaleDownBehaviorInput() *string
	SecurityConfiguration() *string
	SetSecurityConfiguration(val *string)
	SecurityConfigurationInput() *string
	ServiceRole() *string
	SetServiceRole(val *string)
	ServiceRoleInput() *string
	Step() EmrClusterStepList
	StepConcurrencyLevel() *float64
	SetStepConcurrencyLevel(val *float64)
	StepConcurrencyLevelInput() *float64
	StepInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerminationProtection() interface{}
	SetTerminationProtection(val interface{})
	TerminationProtectionInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VisibleToAllUsers() interface{}
	SetVisibleToAllUsers(val interface{})
	VisibleToAllUsersInput() interface{}
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
	PutAutoTerminationPolicy(value *EmrClusterAutoTerminationPolicy)
	PutBootstrapAction(value interface{})
	PutCoreInstanceFleet(value *EmrClusterCoreInstanceFleet)
	PutCoreInstanceGroup(value *EmrClusterCoreInstanceGroup)
	PutEc2Attributes(value *EmrClusterEc2Attributes)
	PutKerberosAttributes(value *EmrClusterKerberosAttributes)
	PutMasterInstanceFleet(value *EmrClusterMasterInstanceFleet)
	PutMasterInstanceGroup(value *EmrClusterMasterInstanceGroup)
	PutStep(value interface{})
	ResetAdditionalInfo()
	ResetApplications()
	ResetAutoscalingRole()
	ResetAutoTerminationPolicy()
	ResetBootstrapAction()
	ResetConfigurations()
	ResetConfigurationsJson()
	ResetCoreInstanceFleet()
	ResetCoreInstanceGroup()
	ResetCustomAmiId()
	ResetEbsRootVolumeSize()
	ResetEc2Attributes()
	ResetId()
	ResetKeepJobFlowAliveWhenNoSteps()
	ResetKerberosAttributes()
	ResetListStepsStates()
	ResetLogEncryptionKmsKeyId()
	ResetLogUri()
	ResetMasterInstanceFleet()
	ResetMasterInstanceGroup()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScaleDownBehavior()
	ResetSecurityConfiguration()
	ResetStep()
	ResetStepConcurrencyLevel()
	ResetTags()
	ResetTagsAll()
	ResetTerminationProtection()
	ResetVisibleToAllUsers()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for EmrCluster
type jsiiProxy_EmrCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EmrCluster) AdditionalInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) AdditionalInfoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Applications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ApplicationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) AutoscalingRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) AutoscalingRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) AutoTerminationPolicy() EmrClusterAutoTerminationPolicyOutputReference {
	var returns EmrClusterAutoTerminationPolicyOutputReference
	_jsii_.Get(
		j,
		"autoTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) AutoTerminationPolicyInput() *EmrClusterAutoTerminationPolicy {
	var returns *EmrClusterAutoTerminationPolicy
	_jsii_.Get(
		j,
		"autoTerminationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) BootstrapAction() EmrClusterBootstrapActionList {
	var returns EmrClusterBootstrapActionList
	_jsii_.Get(
		j,
		"bootstrapAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) BootstrapActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ClusterState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Configurations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ConfigurationsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ConfigurationsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ConfigurationsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CoreInstanceFleet() EmrClusterCoreInstanceFleetOutputReference {
	var returns EmrClusterCoreInstanceFleetOutputReference
	_jsii_.Get(
		j,
		"coreInstanceFleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CoreInstanceFleetInput() *EmrClusterCoreInstanceFleet {
	var returns *EmrClusterCoreInstanceFleet
	_jsii_.Get(
		j,
		"coreInstanceFleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CoreInstanceGroup() EmrClusterCoreInstanceGroupOutputReference {
	var returns EmrClusterCoreInstanceGroupOutputReference
	_jsii_.Get(
		j,
		"coreInstanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CoreInstanceGroupInput() *EmrClusterCoreInstanceGroup {
	var returns *EmrClusterCoreInstanceGroup
	_jsii_.Get(
		j,
		"coreInstanceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CustomAmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CustomAmiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) EbsRootVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) EbsRootVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Ec2Attributes() EmrClusterEc2AttributesOutputReference {
	var returns EmrClusterEc2AttributesOutputReference
	_jsii_.Get(
		j,
		"ec2Attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Ec2AttributesInput() *EmrClusterEc2Attributes {
	var returns *EmrClusterEc2Attributes
	_jsii_.Get(
		j,
		"ec2AttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) KeepJobFlowAliveWhenNoSteps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepJobFlowAliveWhenNoSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) KeepJobFlowAliveWhenNoStepsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepJobFlowAliveWhenNoStepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) KerberosAttributes() EmrClusterKerberosAttributesOutputReference {
	var returns EmrClusterKerberosAttributesOutputReference
	_jsii_.Get(
		j,
		"kerberosAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) KerberosAttributesInput() *EmrClusterKerberosAttributes {
	var returns *EmrClusterKerberosAttributes
	_jsii_.Get(
		j,
		"kerberosAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ListStepsStates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listStepsStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ListStepsStatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listStepsStatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) LogEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) LogEncryptionKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEncryptionKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) LogUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) LogUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) MasterInstanceFleet() EmrClusterMasterInstanceFleetOutputReference {
	var returns EmrClusterMasterInstanceFleetOutputReference
	_jsii_.Get(
		j,
		"masterInstanceFleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) MasterInstanceFleetInput() *EmrClusterMasterInstanceFleet {
	var returns *EmrClusterMasterInstanceFleet
	_jsii_.Get(
		j,
		"masterInstanceFleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) MasterInstanceGroup() EmrClusterMasterInstanceGroupOutputReference {
	var returns EmrClusterMasterInstanceGroupOutputReference
	_jsii_.Get(
		j,
		"masterInstanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) MasterInstanceGroupInput() *EmrClusterMasterInstanceGroup {
	var returns *EmrClusterMasterInstanceGroup
	_jsii_.Get(
		j,
		"masterInstanceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) MasterPublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPublicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ScaleDownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ScaleDownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Step() EmrClusterStepList {
	var returns EmrClusterStepList
	_jsii_.Get(
		j,
		"step",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) StepConcurrencyLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stepConcurrencyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) StepConcurrencyLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stepConcurrencyLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) StepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TerminationProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TerminationProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) VisibleToAllUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleToAllUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) VisibleToAllUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleToAllUsersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster aws_emr_cluster} Resource.
func NewEmrCluster(scope constructs.Construct, id *string, config *EmrClusterConfig) EmrCluster {
	_init_.Initialize()

	if err := validateNewEmrClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrCluster{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrCluster.EmrCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster aws_emr_cluster} Resource.
func NewEmrCluster_Override(e EmrCluster, scope constructs.Construct, id *string, config *EmrClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrCluster.EmrCluster",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EmrCluster)SetAdditionalInfo(val *string) {
	if err := j.validateSetAdditionalInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalInfo",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetApplications(val *[]*string) {
	if err := j.validateSetApplicationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applications",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetAutoscalingRole(val *string) {
	if err := j.validateSetAutoscalingRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingRole",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetConfigurations(val *string) {
	if err := j.validateSetConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurations",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetConfigurationsJson(val *string) {
	if err := j.validateSetConfigurationsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationsJson",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetCustomAmiId(val *string) {
	if err := j.validateSetCustomAmiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAmiId",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetEbsRootVolumeSize(val *float64) {
	if err := j.validateSetEbsRootVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsRootVolumeSize",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetKeepJobFlowAliveWhenNoSteps(val interface{}) {
	if err := j.validateSetKeepJobFlowAliveWhenNoStepsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepJobFlowAliveWhenNoSteps",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetListStepsStates(val *[]*string) {
	if err := j.validateSetListStepsStatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listStepsStates",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetLogEncryptionKmsKeyId(val *string) {
	if err := j.validateSetLogEncryptionKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetLogUri(val *string) {
	if err := j.validateSetLogUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logUri",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetReleaseLabel(val *string) {
	if err := j.validateSetReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetScaleDownBehavior(val *string) {
	if err := j.validateSetScaleDownBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownBehavior",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetSecurityConfiguration(val *string) {
	if err := j.validateSetSecurityConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetServiceRole(val *string) {
	if err := j.validateSetServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetStepConcurrencyLevel(val *float64) {
	if err := j.validateSetStepConcurrencyLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stepConcurrencyLevel",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetTerminationProtection(val interface{}) {
	if err := j.validateSetTerminationProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationProtection",
		val,
	)
}

func (j *jsiiProxy_EmrCluster)SetVisibleToAllUsers(val interface{}) {
	if err := j.validateSetVisibleToAllUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibleToAllUsers",
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
func EmrCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.emrCluster.EmrCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EmrCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.emrCluster.EmrCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EmrCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.emrCluster.EmrCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EmrCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.emrCluster.EmrCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EmrCluster) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EmrCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EmrCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EmrCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EmrCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EmrCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EmrCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EmrCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EmrCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EmrCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrCluster) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EmrCluster) PutAutoTerminationPolicy(value *EmrClusterAutoTerminationPolicy) {
	if err := e.validatePutAutoTerminationPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutoTerminationPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutBootstrapAction(value interface{}) {
	if err := e.validatePutBootstrapActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBootstrapAction",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutCoreInstanceFleet(value *EmrClusterCoreInstanceFleet) {
	if err := e.validatePutCoreInstanceFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCoreInstanceFleet",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutCoreInstanceGroup(value *EmrClusterCoreInstanceGroup) {
	if err := e.validatePutCoreInstanceGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCoreInstanceGroup",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutEc2Attributes(value *EmrClusterEc2Attributes) {
	if err := e.validatePutEc2AttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEc2Attributes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutKerberosAttributes(value *EmrClusterKerberosAttributes) {
	if err := e.validatePutKerberosAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putKerberosAttributes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutMasterInstanceFleet(value *EmrClusterMasterInstanceFleet) {
	if err := e.validatePutMasterInstanceFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMasterInstanceFleet",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutMasterInstanceGroup(value *EmrClusterMasterInstanceGroup) {
	if err := e.validatePutMasterInstanceGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMasterInstanceGroup",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutStep(value interface{}) {
	if err := e.validatePutStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStep",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) ResetAdditionalInfo() {
	_jsii_.InvokeVoid(
		e,
		"resetAdditionalInfo",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetApplications() {
	_jsii_.InvokeVoid(
		e,
		"resetApplications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetAutoscalingRole() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscalingRole",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetAutoTerminationPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoTerminationPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetBootstrapAction() {
	_jsii_.InvokeVoid(
		e,
		"resetBootstrapAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetConfigurations() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigurations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetConfigurationsJson() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigurationsJson",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetCoreInstanceFleet() {
	_jsii_.InvokeVoid(
		e,
		"resetCoreInstanceFleet",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetCoreInstanceGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetCoreInstanceGroup",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetCustomAmiId() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomAmiId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetEbsRootVolumeSize() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsRootVolumeSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetEc2Attributes() {
	_jsii_.InvokeVoid(
		e,
		"resetEc2Attributes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetKeepJobFlowAliveWhenNoSteps() {
	_jsii_.InvokeVoid(
		e,
		"resetKeepJobFlowAliveWhenNoSteps",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetKerberosAttributes() {
	_jsii_.InvokeVoid(
		e,
		"resetKerberosAttributes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetListStepsStates() {
	_jsii_.InvokeVoid(
		e,
		"resetListStepsStates",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetLogEncryptionKmsKeyId() {
	_jsii_.InvokeVoid(
		e,
		"resetLogEncryptionKmsKeyId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetLogUri() {
	_jsii_.InvokeVoid(
		e,
		"resetLogUri",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetMasterInstanceFleet() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterInstanceFleet",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetMasterInstanceGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterInstanceGroup",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetScaleDownBehavior() {
	_jsii_.InvokeVoid(
		e,
		"resetScaleDownBehavior",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetStep() {
	_jsii_.InvokeVoid(
		e,
		"resetStep",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetStepConcurrencyLevel() {
	_jsii_.InvokeVoid(
		e,
		"resetStepConcurrencyLevel",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetTerminationProtection() {
	_jsii_.InvokeVoid(
		e,
		"resetTerminationProtection",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetVisibleToAllUsers() {
	_jsii_.InvokeVoid(
		e,
		"resetVisibleToAllUsers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

