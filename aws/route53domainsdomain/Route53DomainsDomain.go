// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53domainsdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/route53domainsdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/route53domains_domain aws_route53domains_domain}.
type Route53DomainsDomain interface {
	cdktf.TerraformResource
	AbuseContactEmail() *string
	AbuseContactPhone() *string
	AdminContact() Route53DomainsDomainAdminContactList
	AdminContactInput() interface{}
	AdminPrivacy() interface{}
	SetAdminPrivacy(val interface{})
	AdminPrivacyInput() interface{}
	AutoRenew() interface{}
	SetAutoRenew(val interface{})
	AutoRenewInput() interface{}
	BillingContact() Route53DomainsDomainBillingContactList
	BillingContactInput() interface{}
	BillingPrivacy() interface{}
	SetBillingPrivacy(val interface{})
	BillingPrivacyInput() interface{}
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
	CreationDate() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	DurationInYears() *float64
	SetDurationInYears(val *float64)
	DurationInYearsInput() *float64
	ExpirationDate() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostedZoneId() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NameServer() Route53DomainsDomainNameServerList
	NameServerInput() interface{}
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
	RegistrantContact() Route53DomainsDomainRegistrantContactList
	RegistrantContactInput() interface{}
	RegistrantPrivacy() interface{}
	SetRegistrantPrivacy(val interface{})
	RegistrantPrivacyInput() interface{}
	RegistrarName() *string
	RegistrarUrl() *string
	StatusList() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktf.StringMap
	TagsInput() *map[string]*string
	TechContact() Route53DomainsDomainTechContactList
	TechContactInput() interface{}
	TechPrivacy() interface{}
	SetTechPrivacy(val interface{})
	TechPrivacyInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() Route53DomainsDomainTimeoutsOutputReference
	TimeoutsInput() interface{}
	TransferLock() interface{}
	SetTransferLock(val interface{})
	TransferLockInput() interface{}
	UpdatedDate() *string
	WhoisServer() *string
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
	PutAdminContact(value interface{})
	PutBillingContact(value interface{})
	PutNameServer(value interface{})
	PutRegistrantContact(value interface{})
	PutTechContact(value interface{})
	PutTimeouts(value *Route53DomainsDomainTimeouts)
	ResetAdminContact()
	ResetAdminPrivacy()
	ResetAutoRenew()
	ResetBillingContact()
	ResetBillingPrivacy()
	ResetDurationInYears()
	ResetNameServer()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegistrantContact()
	ResetRegistrantPrivacy()
	ResetTags()
	ResetTechContact()
	ResetTechPrivacy()
	ResetTimeouts()
	ResetTransferLock()
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

// The jsii proxy struct for Route53DomainsDomain
type jsiiProxy_Route53DomainsDomain struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Route53DomainsDomain) AbuseContactEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abuseContactEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) AbuseContactPhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abuseContactPhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) AdminContact() Route53DomainsDomainAdminContactList {
	var returns Route53DomainsDomainAdminContactList
	_jsii_.Get(
		j,
		"adminContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) AdminContactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) AdminPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) AdminPrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminPrivacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) AutoRenew() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRenew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) AutoRenewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRenewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) BillingContact() Route53DomainsDomainBillingContactList {
	var returns Route53DomainsDomainBillingContactList
	_jsii_.Get(
		j,
		"billingContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) BillingContactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"billingContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) BillingPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"billingPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) BillingPrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"billingPrivacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) DurationInYears() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationInYears",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) DurationInYearsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationInYearsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) ExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) NameServer() Route53DomainsDomainNameServerList {
	var returns Route53DomainsDomainNameServerList
	_jsii_.Get(
		j,
		"nameServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) NameServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) RegistrantContact() Route53DomainsDomainRegistrantContactList {
	var returns Route53DomainsDomainRegistrantContactList
	_jsii_.Get(
		j,
		"registrantContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) RegistrantContactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registrantContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) RegistrantPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registrantPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) RegistrantPrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registrantPrivacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) RegistrarName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrarName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) RegistrarUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrarUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) StatusList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statusList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) TechContact() Route53DomainsDomainTechContactList {
	var returns Route53DomainsDomainTechContactList
	_jsii_.Get(
		j,
		"techContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) TechContactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"techContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) TechPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"techPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) TechPrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"techPrivacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) Timeouts() Route53DomainsDomainTimeoutsOutputReference {
	var returns Route53DomainsDomainTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) TransferLock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transferLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) TransferLockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transferLockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) UpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomain) WhoisServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoisServer",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/route53domains_domain aws_route53domains_domain} Resource.
func NewRoute53DomainsDomain(scope constructs.Construct, id *string, config *Route53DomainsDomainConfig) Route53DomainsDomain {
	_init_.Initialize()

	if err := validateNewRoute53DomainsDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53DomainsDomain{}

	_jsii_.Create(
		"@cdktf/provider-aws.route53DomainsDomain.Route53DomainsDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/route53domains_domain aws_route53domains_domain} Resource.
func NewRoute53DomainsDomain_Override(r Route53DomainsDomain, scope constructs.Construct, id *string, config *Route53DomainsDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.route53DomainsDomain.Route53DomainsDomain",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetAdminPrivacy(val interface{}) {
	if err := j.validateSetAdminPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPrivacy",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetAutoRenew(val interface{}) {
	if err := j.validateSetAutoRenewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRenew",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetBillingPrivacy(val interface{}) {
	if err := j.validateSetBillingPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingPrivacy",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetDurationInYears(val *float64) {
	if err := j.validateSetDurationInYearsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durationInYears",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetRegistrantPrivacy(val interface{}) {
	if err := j.validateSetRegistrantPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registrantPrivacy",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetTechPrivacy(val interface{}) {
	if err := j.validateSetTechPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"techPrivacy",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomain)SetTransferLock(val interface{}) {
	if err := j.validateSetTransferLockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transferLock",
		val,
	)
}

// Generates CDKTF code for importing a Route53DomainsDomain resource upon running "cdktf plan <stack-name>".
func Route53DomainsDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRoute53DomainsDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53DomainsDomain.Route53DomainsDomain",
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
func Route53DomainsDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53DomainsDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53DomainsDomain.Route53DomainsDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53DomainsDomain_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53DomainsDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53DomainsDomain.Route53DomainsDomain",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53DomainsDomain_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53DomainsDomain_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53DomainsDomain.Route53DomainsDomain",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Route53DomainsDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.route53DomainsDomain.Route53DomainsDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_Route53DomainsDomain) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Route53DomainsDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_Route53DomainsDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53DomainsDomain) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_Route53DomainsDomain) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53DomainsDomain) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Route53DomainsDomain) PutAdminContact(value interface{}) {
	if err := r.validatePutAdminContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAdminContact",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53DomainsDomain) PutBillingContact(value interface{}) {
	if err := r.validatePutBillingContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putBillingContact",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53DomainsDomain) PutNameServer(value interface{}) {
	if err := r.validatePutNameServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putNameServer",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53DomainsDomain) PutRegistrantContact(value interface{}) {
	if err := r.validatePutRegistrantContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRegistrantContact",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53DomainsDomain) PutTechContact(value interface{}) {
	if err := r.validatePutTechContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTechContact",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53DomainsDomain) PutTimeouts(value *Route53DomainsDomainTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetAdminContact() {
	_jsii_.InvokeVoid(
		r,
		"resetAdminContact",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetAdminPrivacy() {
	_jsii_.InvokeVoid(
		r,
		"resetAdminPrivacy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetAutoRenew() {
	_jsii_.InvokeVoid(
		r,
		"resetAutoRenew",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetBillingContact() {
	_jsii_.InvokeVoid(
		r,
		"resetBillingContact",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetBillingPrivacy() {
	_jsii_.InvokeVoid(
		r,
		"resetBillingPrivacy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetDurationInYears() {
	_jsii_.InvokeVoid(
		r,
		"resetDurationInYears",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetNameServer() {
	_jsii_.InvokeVoid(
		r,
		"resetNameServer",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetRegistrantContact() {
	_jsii_.InvokeVoid(
		r,
		"resetRegistrantContact",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetRegistrantPrivacy() {
	_jsii_.InvokeVoid(
		r,
		"resetRegistrantPrivacy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetTechContact() {
	_jsii_.InvokeVoid(
		r,
		"resetTechContact",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetTechPrivacy() {
	_jsii_.InvokeVoid(
		r,
		"resetTechPrivacy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) ResetTransferLock() {
	_jsii_.InvokeVoid(
		r,
		"resetTransferLock",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

