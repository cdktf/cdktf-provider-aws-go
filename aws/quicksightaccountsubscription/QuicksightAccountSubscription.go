// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightaccountsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/quicksightaccountsubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/quicksight_account_subscription aws_quicksight_account_subscription}.
type QuicksightAccountSubscription interface {
	cdktf.TerraformResource
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
	AccountSubscriptionStatus() *string
	ActiveDirectoryName() *string
	SetActiveDirectoryName(val *string)
	ActiveDirectoryNameInput() *string
	AdminGroup() *[]*string
	SetAdminGroup(val *[]*string)
	AdminGroupInput() *[]*string
	AuthenticationMethod() *string
	SetAuthenticationMethod(val *string)
	AuthenticationMethodInput() *string
	AuthorGroup() *[]*string
	SetAuthorGroup(val *[]*string)
	AuthorGroupInput() *[]*string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	AwsAccountIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContactNumber() *string
	SetContactNumber(val *string)
	ContactNumberInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DirectoryId() *string
	SetDirectoryId(val *string)
	DirectoryIdInput() *string
	Edition() *string
	SetEdition(val *string)
	EditionInput() *string
	EmailAddress() *string
	SetEmailAddress(val *string)
	EmailAddressInput() *string
	FirstName() *string
	SetFirstName(val *string)
	FirstNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamIdentityCenterInstanceArn() *string
	SetIamIdentityCenterInstanceArn(val *string)
	IamIdentityCenterInstanceArnInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LastName() *string
	SetLastName(val *string)
	LastNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NotificationEmail() *string
	SetNotificationEmail(val *string)
	NotificationEmailInput() *string
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
	ReaderGroup() *[]*string
	SetReaderGroup(val *[]*string)
	ReaderGroupInput() *[]*string
	Realm() *string
	SetRealm(val *string)
	RealmInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() QuicksightAccountSubscriptionTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *QuicksightAccountSubscriptionTimeouts)
	ResetActiveDirectoryName()
	ResetAdminGroup()
	ResetAuthorGroup()
	ResetAwsAccountId()
	ResetContactNumber()
	ResetDirectoryId()
	ResetEmailAddress()
	ResetFirstName()
	ResetIamIdentityCenterInstanceArn()
	ResetId()
	ResetLastName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReaderGroup()
	ResetRealm()
	ResetTimeouts()
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

// The jsii proxy struct for QuicksightAccountSubscription
type jsiiProxy_QuicksightAccountSubscription struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_QuicksightAccountSubscription) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) AccountSubscriptionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountSubscriptionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) ActiveDirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) ActiveDirectoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) AdminGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) AdminGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) AuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) AuthorGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) AuthorGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) AwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) ContactNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) ContactNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) EditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) EmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) EmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) IamIdentityCenterInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamIdentityCenterInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) IamIdentityCenterInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamIdentityCenterInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) NotificationEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) NotificationEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) ReaderGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readerGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) ReaderGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readerGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) RealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) Timeouts() QuicksightAccountSubscriptionTimeoutsOutputReference {
	var returns QuicksightAccountSubscriptionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAccountSubscription) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/quicksight_account_subscription aws_quicksight_account_subscription} Resource.
func NewQuicksightAccountSubscription(scope constructs.Construct, id *string, config *QuicksightAccountSubscriptionConfig) QuicksightAccountSubscription {
	_init_.Initialize()

	if err := validateNewQuicksightAccountSubscriptionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightAccountSubscription{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightAccountSubscription.QuicksightAccountSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/quicksight_account_subscription aws_quicksight_account_subscription} Resource.
func NewQuicksightAccountSubscription_Override(q QuicksightAccountSubscription, scope constructs.Construct, id *string, config *QuicksightAccountSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightAccountSubscription.QuicksightAccountSubscription",
		[]interface{}{scope, id, config},
		q,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetAccountName(val *string) {
	if err := j.validateSetAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetActiveDirectoryName(val *string) {
	if err := j.validateSetActiveDirectoryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDirectoryName",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetAdminGroup(val *[]*string) {
	if err := j.validateSetAdminGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminGroup",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetAuthenticationMethod(val *string) {
	if err := j.validateSetAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMethod",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetAuthorGroup(val *[]*string) {
	if err := j.validateSetAuthorGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorGroup",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetAwsAccountId(val *string) {
	if err := j.validateSetAwsAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetContactNumber(val *string) {
	if err := j.validateSetContactNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactNumber",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetDirectoryId(val *string) {
	if err := j.validateSetDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetEdition(val *string) {
	if err := j.validateSetEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetEmailAddress(val *string) {
	if err := j.validateSetEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddress",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetIamIdentityCenterInstanceArn(val *string) {
	if err := j.validateSetIamIdentityCenterInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamIdentityCenterInstanceArn",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetNotificationEmail(val *string) {
	if err := j.validateSetNotificationEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationEmail",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetReaderGroup(val *[]*string) {
	if err := j.validateSetReaderGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readerGroup",
		val,
	)
}

func (j *jsiiProxy_QuicksightAccountSubscription)SetRealm(val *string) {
	if err := j.validateSetRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

// Generates CDKTF code for importing a QuicksightAccountSubscription resource upon running "cdktf plan <stack-name>".
func QuicksightAccountSubscription_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateQuicksightAccountSubscription_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.quicksightAccountSubscription.QuicksightAccountSubscription",
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
func QuicksightAccountSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightAccountSubscription_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.quicksightAccountSubscription.QuicksightAccountSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QuicksightAccountSubscription_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightAccountSubscription_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.quicksightAccountSubscription.QuicksightAccountSubscription",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QuicksightAccountSubscription_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightAccountSubscription_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.quicksightAccountSubscription.QuicksightAccountSubscription",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QuicksightAccountSubscription_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.quicksightAccountSubscription.QuicksightAccountSubscription",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) AddMoveTarget(moveTarget *string) {
	if err := q.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) AddOverride(path *string, value interface{}) {
	if err := q.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := q.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) MoveFromId(id *string) {
	if err := q.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveFromId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) MoveTo(moveTarget *string, index interface{}) {
	if err := q.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) MoveToId(id *string) {
	if err := q.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveToId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) OverrideLogicalId(newLogicalId *string) {
	if err := q.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) PutTimeouts(value *QuicksightAccountSubscriptionTimeouts) {
	if err := q.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetActiveDirectoryName() {
	_jsii_.InvokeVoid(
		q,
		"resetActiveDirectoryName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetAdminGroup() {
	_jsii_.InvokeVoid(
		q,
		"resetAdminGroup",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetAuthorGroup() {
	_jsii_.InvokeVoid(
		q,
		"resetAuthorGroup",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		q,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetContactNumber() {
	_jsii_.InvokeVoid(
		q,
		"resetContactNumber",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		q,
		"resetDirectoryId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetEmailAddress() {
	_jsii_.InvokeVoid(
		q,
		"resetEmailAddress",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetFirstName() {
	_jsii_.InvokeVoid(
		q,
		"resetFirstName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetIamIdentityCenterInstanceArn() {
	_jsii_.InvokeVoid(
		q,
		"resetIamIdentityCenterInstanceArn",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetId() {
	_jsii_.InvokeVoid(
		q,
		"resetId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetLastName() {
	_jsii_.InvokeVoid(
		q,
		"resetLastName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		q,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetReaderGroup() {
	_jsii_.InvokeVoid(
		q,
		"resetReaderGroup",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetRealm() {
	_jsii_.InvokeVoid(
		q,
		"resetRealm",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) ResetTimeouts() {
	_jsii_.InvokeVoid(
		q,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAccountSubscription) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAccountSubscription) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

