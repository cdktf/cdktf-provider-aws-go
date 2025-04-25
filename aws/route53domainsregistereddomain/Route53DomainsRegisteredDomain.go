// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53domainsregistereddomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/route53domainsregistereddomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/route53domains_registered_domain aws_route53domains_registered_domain}.
type Route53DomainsRegisteredDomain interface {
	cdktf.TerraformResource
	AbuseContactEmail() *string
	AbuseContactPhone() *string
	AdminContact() Route53DomainsRegisteredDomainAdminContactOutputReference
	AdminContactInput() *Route53DomainsRegisteredDomainAdminContact
	AdminPrivacy() interface{}
	SetAdminPrivacy(val interface{})
	AdminPrivacyInput() interface{}
	AutoRenew() interface{}
	SetAutoRenew(val interface{})
	AutoRenewInput() interface{}
	BillingContact() Route53DomainsRegisteredDomainBillingContactOutputReference
	BillingContactInput() *Route53DomainsRegisteredDomainBillingContact
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
	ExpirationDate() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NameServer() Route53DomainsRegisteredDomainNameServerList
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
	RegistrantContact() Route53DomainsRegisteredDomainRegistrantContactOutputReference
	RegistrantContactInput() *Route53DomainsRegisteredDomainRegistrantContact
	RegistrantPrivacy() interface{}
	SetRegistrantPrivacy(val interface{})
	RegistrantPrivacyInput() interface{}
	RegistrarName() *string
	RegistrarUrl() *string
	Reseller() *string
	StatusList() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TechContact() Route53DomainsRegisteredDomainTechContactOutputReference
	TechContactInput() *Route53DomainsRegisteredDomainTechContact
	TechPrivacy() interface{}
	SetTechPrivacy(val interface{})
	TechPrivacyInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() Route53DomainsRegisteredDomainTimeoutsOutputReference
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
	PutAdminContact(value *Route53DomainsRegisteredDomainAdminContact)
	PutBillingContact(value *Route53DomainsRegisteredDomainBillingContact)
	PutNameServer(value interface{})
	PutRegistrantContact(value *Route53DomainsRegisteredDomainRegistrantContact)
	PutTechContact(value *Route53DomainsRegisteredDomainTechContact)
	PutTimeouts(value *Route53DomainsRegisteredDomainTimeouts)
	ResetAdminContact()
	ResetAdminPrivacy()
	ResetAutoRenew()
	ResetBillingContact()
	ResetBillingPrivacy()
	ResetId()
	ResetNameServer()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegistrantContact()
	ResetRegistrantPrivacy()
	ResetTags()
	ResetTagsAll()
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

// The jsii proxy struct for Route53DomainsRegisteredDomain
type jsiiProxy_Route53DomainsRegisteredDomain struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) AbuseContactEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abuseContactEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) AbuseContactPhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abuseContactPhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) AdminContact() Route53DomainsRegisteredDomainAdminContactOutputReference {
	var returns Route53DomainsRegisteredDomainAdminContactOutputReference
	_jsii_.Get(
		j,
		"adminContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) AdminContactInput() *Route53DomainsRegisteredDomainAdminContact {
	var returns *Route53DomainsRegisteredDomainAdminContact
	_jsii_.Get(
		j,
		"adminContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) AdminPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) AdminPrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminPrivacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) AutoRenew() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRenew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) AutoRenewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRenewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) BillingContact() Route53DomainsRegisteredDomainBillingContactOutputReference {
	var returns Route53DomainsRegisteredDomainBillingContactOutputReference
	_jsii_.Get(
		j,
		"billingContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) BillingContactInput() *Route53DomainsRegisteredDomainBillingContact {
	var returns *Route53DomainsRegisteredDomainBillingContact
	_jsii_.Get(
		j,
		"billingContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) BillingPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"billingPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) BillingPrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"billingPrivacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) ExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) NameServer() Route53DomainsRegisteredDomainNameServerList {
	var returns Route53DomainsRegisteredDomainNameServerList
	_jsii_.Get(
		j,
		"nameServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) NameServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) RegistrantContact() Route53DomainsRegisteredDomainRegistrantContactOutputReference {
	var returns Route53DomainsRegisteredDomainRegistrantContactOutputReference
	_jsii_.Get(
		j,
		"registrantContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) RegistrantContactInput() *Route53DomainsRegisteredDomainRegistrantContact {
	var returns *Route53DomainsRegisteredDomainRegistrantContact
	_jsii_.Get(
		j,
		"registrantContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) RegistrantPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registrantPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) RegistrantPrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registrantPrivacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) RegistrarName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrarName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) RegistrarUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrarUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) Reseller() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reseller",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) StatusList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statusList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) TechContact() Route53DomainsRegisteredDomainTechContactOutputReference {
	var returns Route53DomainsRegisteredDomainTechContactOutputReference
	_jsii_.Get(
		j,
		"techContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) TechContactInput() *Route53DomainsRegisteredDomainTechContact {
	var returns *Route53DomainsRegisteredDomainTechContact
	_jsii_.Get(
		j,
		"techContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) TechPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"techPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) TechPrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"techPrivacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) Timeouts() Route53DomainsRegisteredDomainTimeoutsOutputReference {
	var returns Route53DomainsRegisteredDomainTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) TransferLock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transferLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) TransferLockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transferLockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) UpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain) WhoisServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoisServer",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/route53domains_registered_domain aws_route53domains_registered_domain} Resource.
func NewRoute53DomainsRegisteredDomain(scope constructs.Construct, id *string, config *Route53DomainsRegisteredDomainConfig) Route53DomainsRegisteredDomain {
	_init_.Initialize()

	if err := validateNewRoute53DomainsRegisteredDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53DomainsRegisteredDomain{}

	_jsii_.Create(
		"@cdktf/provider-aws.route53DomainsRegisteredDomain.Route53DomainsRegisteredDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/route53domains_registered_domain aws_route53domains_registered_domain} Resource.
func NewRoute53DomainsRegisteredDomain_Override(r Route53DomainsRegisteredDomain, scope constructs.Construct, id *string, config *Route53DomainsRegisteredDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.route53DomainsRegisteredDomain.Route53DomainsRegisteredDomain",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetAdminPrivacy(val interface{}) {
	if err := j.validateSetAdminPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPrivacy",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetAutoRenew(val interface{}) {
	if err := j.validateSetAutoRenewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRenew",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetBillingPrivacy(val interface{}) {
	if err := j.validateSetBillingPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingPrivacy",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetRegistrantPrivacy(val interface{}) {
	if err := j.validateSetRegistrantPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registrantPrivacy",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetTechPrivacy(val interface{}) {
	if err := j.validateSetTechPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"techPrivacy",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomain)SetTransferLock(val interface{}) {
	if err := j.validateSetTransferLockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transferLock",
		val,
	)
}

// Generates CDKTF code for importing a Route53DomainsRegisteredDomain resource upon running "cdktf plan <stack-name>".
func Route53DomainsRegisteredDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRoute53DomainsRegisteredDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53DomainsRegisteredDomain.Route53DomainsRegisteredDomain",
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
func Route53DomainsRegisteredDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53DomainsRegisteredDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53DomainsRegisteredDomain.Route53DomainsRegisteredDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53DomainsRegisteredDomain_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53DomainsRegisteredDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53DomainsRegisteredDomain.Route53DomainsRegisteredDomain",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53DomainsRegisteredDomain_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53DomainsRegisteredDomain_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53DomainsRegisteredDomain.Route53DomainsRegisteredDomain",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Route53DomainsRegisteredDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.route53DomainsRegisteredDomain.Route53DomainsRegisteredDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomain) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomain) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomain) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomain) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) PutAdminContact(value *Route53DomainsRegisteredDomainAdminContact) {
	if err := r.validatePutAdminContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAdminContact",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) PutBillingContact(value *Route53DomainsRegisteredDomainBillingContact) {
	if err := r.validatePutBillingContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putBillingContact",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) PutNameServer(value interface{}) {
	if err := r.validatePutNameServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putNameServer",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) PutRegistrantContact(value *Route53DomainsRegisteredDomainRegistrantContact) {
	if err := r.validatePutRegistrantContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRegistrantContact",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) PutTechContact(value *Route53DomainsRegisteredDomainTechContact) {
	if err := r.validatePutTechContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTechContact",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) PutTimeouts(value *Route53DomainsRegisteredDomainTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetAdminContact() {
	_jsii_.InvokeVoid(
		r,
		"resetAdminContact",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetAdminPrivacy() {
	_jsii_.InvokeVoid(
		r,
		"resetAdminPrivacy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetAutoRenew() {
	_jsii_.InvokeVoid(
		r,
		"resetAutoRenew",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetBillingContact() {
	_jsii_.InvokeVoid(
		r,
		"resetBillingContact",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetBillingPrivacy() {
	_jsii_.InvokeVoid(
		r,
		"resetBillingPrivacy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetNameServer() {
	_jsii_.InvokeVoid(
		r,
		"resetNameServer",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetRegistrantContact() {
	_jsii_.InvokeVoid(
		r,
		"resetRegistrantContact",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetRegistrantPrivacy() {
	_jsii_.InvokeVoid(
		r,
		"resetRegistrantPrivacy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetTagsAll() {
	_jsii_.InvokeVoid(
		r,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetTechContact() {
	_jsii_.InvokeVoid(
		r,
		"resetTechContact",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetTechPrivacy() {
	_jsii_.InvokeVoid(
		r,
		"resetTechPrivacy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ResetTransferLock() {
	_jsii_.InvokeVoid(
		r,
		"resetTransferLock",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsRegisteredDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

