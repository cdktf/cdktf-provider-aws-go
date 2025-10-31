// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pinpointsmsvoicev2phonenumber

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/pinpointsmsvoicev2phonenumber/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/pinpointsmsvoicev2_phone_number aws_pinpointsmsvoicev2_phone_number}.
type Pinpointsmsvoicev2PhoneNumber interface {
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
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DeletionProtectionEnabled() interface{}
	SetDeletionProtectionEnabled(val interface{})
	DeletionProtectionEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IsoCountryCode() *string
	SetIsoCountryCode(val *string)
	IsoCountryCodeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MessageType() *string
	SetMessageType(val *string)
	MessageTypeInput() *string
	MonthlyLeasingPrice() *string
	// The tree node.
	Node() constructs.Node
	NumberCapabilities() *[]*string
	SetNumberCapabilities(val *[]*string)
	NumberCapabilitiesInput() *[]*string
	NumberType() *string
	SetNumberType(val *string)
	NumberTypeInput() *string
	OptOutListName() *string
	SetOptOutListName(val *string)
	OptOutListNameInput() *string
	PhoneNumber() *string
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
	RegistrationId() *string
	SetRegistrationId(val *string)
	RegistrationIdInput() *string
	SelfManagedOptOutsEnabled() interface{}
	SetSelfManagedOptOutsEnabled(val interface{})
	SelfManagedOptOutsEnabledInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktf.StringMap
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() Pinpointsmsvoicev2PhoneNumberTimeoutsOutputReference
	TimeoutsInput() interface{}
	TwoWayChannelArn() *string
	SetTwoWayChannelArn(val *string)
	TwoWayChannelArnInput() *string
	TwoWayChannelEnabled() interface{}
	SetTwoWayChannelEnabled(val interface{})
	TwoWayChannelEnabledInput() interface{}
	TwoWayChannelRole() *string
	SetTwoWayChannelRole(val *string)
	TwoWayChannelRoleInput() *string
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
	PutTimeouts(value *Pinpointsmsvoicev2PhoneNumberTimeouts)
	ResetDeletionProtectionEnabled()
	ResetOptOutListName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetRegistrationId()
	ResetSelfManagedOptOutsEnabled()
	ResetTags()
	ResetTimeouts()
	ResetTwoWayChannelArn()
	ResetTwoWayChannelEnabled()
	ResetTwoWayChannelRole()
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

// The jsii proxy struct for Pinpointsmsvoicev2PhoneNumber
type jsiiProxy_Pinpointsmsvoicev2PhoneNumber struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) IsoCountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isoCountryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) IsoCountryCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isoCountryCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) MessageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) MessageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) MonthlyLeasingPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyLeasingPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) NumberCapabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"numberCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) NumberCapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"numberCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) NumberType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"numberType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) NumberTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"numberTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) OptOutListName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optOutListName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) OptOutListNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optOutListNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) RegistrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) RegistrationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) SelfManagedOptOutsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedOptOutsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) SelfManagedOptOutsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedOptOutsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) Timeouts() Pinpointsmsvoicev2PhoneNumberTimeoutsOutputReference {
	var returns Pinpointsmsvoicev2PhoneNumberTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) TwoWayChannelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) TwoWayChannelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) TwoWayChannelEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"twoWayChannelEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) TwoWayChannelEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"twoWayChannelEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) TwoWayChannelRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) TwoWayChannelRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelRoleInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/pinpointsmsvoicev2_phone_number aws_pinpointsmsvoicev2_phone_number} Resource.
func NewPinpointsmsvoicev2PhoneNumber(scope constructs.Construct, id *string, config *Pinpointsmsvoicev2PhoneNumberConfig) Pinpointsmsvoicev2PhoneNumber {
	_init_.Initialize()

	if err := validateNewPinpointsmsvoicev2PhoneNumberParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Pinpointsmsvoicev2PhoneNumber{}

	_jsii_.Create(
		"@cdktf/provider-aws.pinpointsmsvoicev2PhoneNumber.Pinpointsmsvoicev2PhoneNumber",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/pinpointsmsvoicev2_phone_number aws_pinpointsmsvoicev2_phone_number} Resource.
func NewPinpointsmsvoicev2PhoneNumber_Override(p Pinpointsmsvoicev2PhoneNumber, scope constructs.Construct, id *string, config *Pinpointsmsvoicev2PhoneNumberConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pinpointsmsvoicev2PhoneNumber.Pinpointsmsvoicev2PhoneNumber",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetIsoCountryCode(val *string) {
	if err := j.validateSetIsoCountryCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isoCountryCode",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetMessageType(val *string) {
	if err := j.validateSetMessageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageType",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetNumberCapabilities(val *[]*string) {
	if err := j.validateSetNumberCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberCapabilities",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetNumberType(val *string) {
	if err := j.validateSetNumberTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberType",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetOptOutListName(val *string) {
	if err := j.validateSetOptOutListNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optOutListName",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetRegistrationId(val *string) {
	if err := j.validateSetRegistrationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registrationId",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetSelfManagedOptOutsEnabled(val interface{}) {
	if err := j.validateSetSelfManagedOptOutsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfManagedOptOutsEnabled",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetTwoWayChannelArn(val *string) {
	if err := j.validateSetTwoWayChannelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayChannelArn",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetTwoWayChannelEnabled(val interface{}) {
	if err := j.validateSetTwoWayChannelEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayChannelEnabled",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2PhoneNumber)SetTwoWayChannelRole(val *string) {
	if err := j.validateSetTwoWayChannelRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayChannelRole",
		val,
	)
}

// Generates CDKTF code for importing a Pinpointsmsvoicev2PhoneNumber resource upon running "cdktf plan <stack-name>".
func Pinpointsmsvoicev2PhoneNumber_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePinpointsmsvoicev2PhoneNumber_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.pinpointsmsvoicev2PhoneNumber.Pinpointsmsvoicev2PhoneNumber",
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
func Pinpointsmsvoicev2PhoneNumber_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePinpointsmsvoicev2PhoneNumber_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.pinpointsmsvoicev2PhoneNumber.Pinpointsmsvoicev2PhoneNumber",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Pinpointsmsvoicev2PhoneNumber_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePinpointsmsvoicev2PhoneNumber_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.pinpointsmsvoicev2PhoneNumber.Pinpointsmsvoicev2PhoneNumber",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Pinpointsmsvoicev2PhoneNumber_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePinpointsmsvoicev2PhoneNumber_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.pinpointsmsvoicev2PhoneNumber.Pinpointsmsvoicev2PhoneNumber",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Pinpointsmsvoicev2PhoneNumber_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.pinpointsmsvoicev2PhoneNumber.Pinpointsmsvoicev2PhoneNumber",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) PutTimeouts(value *Pinpointsmsvoicev2PhoneNumberTimeouts) {
	if err := p.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ResetOptOutListName() {
	_jsii_.InvokeVoid(
		p,
		"resetOptOutListName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ResetRegion() {
	_jsii_.InvokeVoid(
		p,
		"resetRegion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ResetRegistrationId() {
	_jsii_.InvokeVoid(
		p,
		"resetRegistrationId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ResetSelfManagedOptOutsEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetSelfManagedOptOutsEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ResetTwoWayChannelArn() {
	_jsii_.InvokeVoid(
		p,
		"resetTwoWayChannelArn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ResetTwoWayChannelEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetTwoWayChannelEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ResetTwoWayChannelRole() {
	_jsii_.InvokeVoid(
		p,
		"resetTwoWayChannelRole",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2PhoneNumber) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

