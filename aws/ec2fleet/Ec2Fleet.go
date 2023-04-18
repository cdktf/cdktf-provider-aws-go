package ec2fleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/ec2fleet/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/ec2_fleet aws_ec2_fleet}.
type Ec2Fleet interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Context() *string
	SetContext(val *string)
	ContextInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExcessCapacityTerminationPolicy() *string
	SetExcessCapacityTerminationPolicy(val *string)
	ExcessCapacityTerminationPolicyInput() *string
	FleetInstanceSet() Ec2FleetFleetInstanceSetList
	FleetInstanceSetInput() interface{}
	FleetState() *string
	SetFleetState(val *string)
	FleetStateInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FulfilledCapacity() *float64
	SetFulfilledCapacity(val *float64)
	FulfilledCapacityInput() *float64
	FulfilledOnDemandCapacity() *float64
	SetFulfilledOnDemandCapacity(val *float64)
	FulfilledOnDemandCapacityInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	LaunchTemplateConfig() Ec2FleetLaunchTemplateConfigList
	LaunchTemplateConfigInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OnDemandOptions() Ec2FleetOnDemandOptionsOutputReference
	OnDemandOptionsInput() *Ec2FleetOnDemandOptions
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
	ReplaceUnhealthyInstances() interface{}
	SetReplaceUnhealthyInstances(val interface{})
	ReplaceUnhealthyInstancesInput() interface{}
	SpotOptions() Ec2FleetSpotOptionsOutputReference
	SpotOptionsInput() *Ec2FleetSpotOptions
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TargetCapacitySpecification() Ec2FleetTargetCapacitySpecificationOutputReference
	TargetCapacitySpecificationInput() *Ec2FleetTargetCapacitySpecification
	TerminateInstances() interface{}
	SetTerminateInstances(val interface{})
	TerminateInstancesInput() interface{}
	TerminateInstancesWithExpiration() interface{}
	SetTerminateInstancesWithExpiration(val interface{})
	TerminateInstancesWithExpirationInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() Ec2FleetTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	ValidFrom() *string
	SetValidFrom(val *string)
	ValidFromInput() *string
	ValidUntil() *string
	SetValidUntil(val *string)
	ValidUntilInput() *string
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
	PutFleetInstanceSet(value interface{})
	PutLaunchTemplateConfig(value interface{})
	PutOnDemandOptions(value *Ec2FleetOnDemandOptions)
	PutSpotOptions(value *Ec2FleetSpotOptions)
	PutTargetCapacitySpecification(value *Ec2FleetTargetCapacitySpecification)
	PutTimeouts(value *Ec2FleetTimeouts)
	ResetContext()
	ResetExcessCapacityTerminationPolicy()
	ResetFleetInstanceSet()
	ResetFleetState()
	ResetFulfilledCapacity()
	ResetFulfilledOnDemandCapacity()
	ResetId()
	ResetOnDemandOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReplaceUnhealthyInstances()
	ResetSpotOptions()
	ResetTags()
	ResetTagsAll()
	ResetTerminateInstances()
	ResetTerminateInstancesWithExpiration()
	ResetTimeouts()
	ResetType()
	ResetValidFrom()
	ResetValidUntil()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Ec2Fleet
type jsiiProxy_Ec2Fleet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Ec2Fleet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) ExcessCapacityTerminationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) ExcessCapacityTerminationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) FleetInstanceSet() Ec2FleetFleetInstanceSetList {
	var returns Ec2FleetFleetInstanceSetList
	_jsii_.Get(
		j,
		"fleetInstanceSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) FleetInstanceSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fleetInstanceSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) FleetState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) FleetStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) FulfilledCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fulfilledCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) FulfilledCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fulfilledCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) FulfilledOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fulfilledOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) FulfilledOnDemandCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fulfilledOnDemandCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) LaunchTemplateConfig() Ec2FleetLaunchTemplateConfigList {
	var returns Ec2FleetLaunchTemplateConfigList
	_jsii_.Get(
		j,
		"launchTemplateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) LaunchTemplateConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) OnDemandOptions() Ec2FleetOnDemandOptionsOutputReference {
	var returns Ec2FleetOnDemandOptionsOutputReference
	_jsii_.Get(
		j,
		"onDemandOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) OnDemandOptionsInput() *Ec2FleetOnDemandOptions {
	var returns *Ec2FleetOnDemandOptions
	_jsii_.Get(
		j,
		"onDemandOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) ReplaceUnhealthyInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceUnhealthyInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) ReplaceUnhealthyInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceUnhealthyInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) SpotOptions() Ec2FleetSpotOptionsOutputReference {
	var returns Ec2FleetSpotOptionsOutputReference
	_jsii_.Get(
		j,
		"spotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) SpotOptionsInput() *Ec2FleetSpotOptions {
	var returns *Ec2FleetSpotOptions
	_jsii_.Get(
		j,
		"spotOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TargetCapacitySpecification() Ec2FleetTargetCapacitySpecificationOutputReference {
	var returns Ec2FleetTargetCapacitySpecificationOutputReference
	_jsii_.Get(
		j,
		"targetCapacitySpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TargetCapacitySpecificationInput() *Ec2FleetTargetCapacitySpecification {
	var returns *Ec2FleetTargetCapacitySpecification
	_jsii_.Get(
		j,
		"targetCapacitySpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TerminateInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TerminateInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TerminateInstancesWithExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesWithExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TerminateInstancesWithExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesWithExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) Timeouts() Ec2FleetTimeoutsOutputReference {
	var returns Ec2FleetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) ValidFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) ValidFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) ValidUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Fleet) ValidUntilInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntilInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/ec2_fleet aws_ec2_fleet} Resource.
func NewEc2Fleet(scope constructs.Construct, id *string, config *Ec2FleetConfig) Ec2Fleet {
	_init_.Initialize()

	if err := validateNewEc2FleetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Fleet{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2Fleet.Ec2Fleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/ec2_fleet aws_ec2_fleet} Resource.
func NewEc2Fleet_Override(e Ec2Fleet, scope constructs.Construct, id *string, config *Ec2FleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2Fleet.Ec2Fleet",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetExcessCapacityTerminationPolicy(val *string) {
	if err := j.validateSetExcessCapacityTerminationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excessCapacityTerminationPolicy",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetFleetState(val *string) {
	if err := j.validateSetFleetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fleetState",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetFulfilledCapacity(val *float64) {
	if err := j.validateSetFulfilledCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fulfilledCapacity",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetFulfilledOnDemandCapacity(val *float64) {
	if err := j.validateSetFulfilledOnDemandCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fulfilledOnDemandCapacity",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetReplaceUnhealthyInstances(val interface{}) {
	if err := j.validateSetReplaceUnhealthyInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceUnhealthyInstances",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetTerminateInstances(val interface{}) {
	if err := j.validateSetTerminateInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateInstances",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetTerminateInstancesWithExpiration(val interface{}) {
	if err := j.validateSetTerminateInstancesWithExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateInstancesWithExpiration",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetValidFrom(val *string) {
	if err := j.validateSetValidFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validFrom",
		val,
	)
}

func (j *jsiiProxy_Ec2Fleet)SetValidUntil(val *string) {
	if err := j.validateSetValidUntilParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validUntil",
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
func Ec2Fleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Fleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2Fleet.Ec2Fleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2Fleet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Fleet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2Fleet.Ec2Fleet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2Fleet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Fleet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2Fleet.Ec2Fleet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Ec2Fleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.ec2Fleet.Ec2Fleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2Fleet) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Ec2Fleet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2Fleet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2Fleet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2Fleet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2Fleet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2Fleet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2Fleet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2Fleet) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2Fleet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2Fleet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2Fleet) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Ec2Fleet) PutFleetInstanceSet(value interface{}) {
	if err := e.validatePutFleetInstanceSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putFleetInstanceSet",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Fleet) PutLaunchTemplateConfig(value interface{}) {
	if err := e.validatePutLaunchTemplateConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLaunchTemplateConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Fleet) PutOnDemandOptions(value *Ec2FleetOnDemandOptions) {
	if err := e.validatePutOnDemandOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOnDemandOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Fleet) PutSpotOptions(value *Ec2FleetSpotOptions) {
	if err := e.validatePutSpotOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSpotOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Fleet) PutTargetCapacitySpecification(value *Ec2FleetTargetCapacitySpecification) {
	if err := e.validatePutTargetCapacitySpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTargetCapacitySpecification",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Fleet) PutTimeouts(value *Ec2FleetTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetContext() {
	_jsii_.InvokeVoid(
		e,
		"resetContext",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetExcessCapacityTerminationPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetExcessCapacityTerminationPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetFleetInstanceSet() {
	_jsii_.InvokeVoid(
		e,
		"resetFleetInstanceSet",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetFleetState() {
	_jsii_.InvokeVoid(
		e,
		"resetFleetState",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetFulfilledCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetFulfilledCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetFulfilledOnDemandCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetFulfilledOnDemandCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetOnDemandOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetOnDemandOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetReplaceUnhealthyInstances() {
	_jsii_.InvokeVoid(
		e,
		"resetReplaceUnhealthyInstances",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetSpotOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetSpotOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetTerminateInstances() {
	_jsii_.InvokeVoid(
		e,
		"resetTerminateInstances",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetTerminateInstancesWithExpiration() {
	_jsii_.InvokeVoid(
		e,
		"resetTerminateInstancesWithExpiration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetType() {
	_jsii_.InvokeVoid(
		e,
		"resetType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetValidFrom() {
	_jsii_.InvokeVoid(
		e,
		"resetValidFrom",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) ResetValidUntil() {
	_jsii_.InvokeVoid(
		e,
		"resetValidUntil",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Fleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Fleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Fleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Fleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

