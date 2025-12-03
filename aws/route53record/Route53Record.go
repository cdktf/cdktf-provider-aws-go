// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53record

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/route53record/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/route53_record aws_route53_record}.
type Route53Record interface {
	cdktf.TerraformResource
	Alias() Route53RecordAliasOutputReference
	AliasInput() *Route53RecordAlias
	AllowOverwrite() interface{}
	SetAllowOverwrite(val interface{})
	AllowOverwriteInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CidrRoutingPolicy() Route53RecordCidrRoutingPolicyOutputReference
	CidrRoutingPolicyInput() *Route53RecordCidrRoutingPolicy
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
	FailoverRoutingPolicy() Route53RecordFailoverRoutingPolicyOutputReference
	FailoverRoutingPolicyInput() *Route53RecordFailoverRoutingPolicy
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdn() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeolocationRoutingPolicy() Route53RecordGeolocationRoutingPolicyOutputReference
	GeolocationRoutingPolicyInput() *Route53RecordGeolocationRoutingPolicy
	GeoproximityRoutingPolicy() Route53RecordGeoproximityRoutingPolicyOutputReference
	GeoproximityRoutingPolicyInput() *Route53RecordGeoproximityRoutingPolicy
	HealthCheckId() *string
	SetHealthCheckId(val *string)
	HealthCheckIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LatencyRoutingPolicy() Route53RecordLatencyRoutingPolicyOutputReference
	LatencyRoutingPolicyInput() *Route53RecordLatencyRoutingPolicy
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MultivalueAnswerRoutingPolicy() interface{}
	SetMultivalueAnswerRoutingPolicy(val interface{})
	MultivalueAnswerRoutingPolicyInput() interface{}
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
	Records() *[]*string
	SetRecords(val *[]*string)
	RecordsInput() *[]*string
	SetIdentifier() *string
	SetSetIdentifier(val *string)
	SetIdentifierInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() Route53RecordTimeoutsOutputReference
	TimeoutsInput() interface{}
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
	Type() *string
	SetType(val *string)
	TypeInput() *string
	WeightedRoutingPolicy() Route53RecordWeightedRoutingPolicyOutputReference
	WeightedRoutingPolicyInput() *Route53RecordWeightedRoutingPolicy
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	PutAlias(value *Route53RecordAlias)
	PutCidrRoutingPolicy(value *Route53RecordCidrRoutingPolicy)
	PutFailoverRoutingPolicy(value *Route53RecordFailoverRoutingPolicy)
	PutGeolocationRoutingPolicy(value *Route53RecordGeolocationRoutingPolicy)
	PutGeoproximityRoutingPolicy(value *Route53RecordGeoproximityRoutingPolicy)
	PutLatencyRoutingPolicy(value *Route53RecordLatencyRoutingPolicy)
	PutTimeouts(value *Route53RecordTimeouts)
	PutWeightedRoutingPolicy(value *Route53RecordWeightedRoutingPolicy)
	ResetAlias()
	ResetAllowOverwrite()
	ResetCidrRoutingPolicy()
	ResetFailoverRoutingPolicy()
	ResetGeolocationRoutingPolicy()
	ResetGeoproximityRoutingPolicy()
	ResetHealthCheckId()
	ResetId()
	ResetLatencyRoutingPolicy()
	ResetMultivalueAnswerRoutingPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRecords()
	ResetSetIdentifier()
	ResetTimeouts()
	ResetTtl()
	ResetWeightedRoutingPolicy()
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

// The jsii proxy struct for Route53Record
type jsiiProxy_Route53Record struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Route53Record) Alias() Route53RecordAliasOutputReference {
	var returns Route53RecordAliasOutputReference
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) AliasInput() *Route53RecordAlias {
	var returns *Route53RecordAlias
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) AllowOverwrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOverwrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) AllowOverwriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOverwriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) CidrRoutingPolicy() Route53RecordCidrRoutingPolicyOutputReference {
	var returns Route53RecordCidrRoutingPolicyOutputReference
	_jsii_.Get(
		j,
		"cidrRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) CidrRoutingPolicyInput() *Route53RecordCidrRoutingPolicy {
	var returns *Route53RecordCidrRoutingPolicy
	_jsii_.Get(
		j,
		"cidrRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) FailoverRoutingPolicy() Route53RecordFailoverRoutingPolicyOutputReference {
	var returns Route53RecordFailoverRoutingPolicyOutputReference
	_jsii_.Get(
		j,
		"failoverRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) FailoverRoutingPolicyInput() *Route53RecordFailoverRoutingPolicy {
	var returns *Route53RecordFailoverRoutingPolicy
	_jsii_.Get(
		j,
		"failoverRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) GeolocationRoutingPolicy() Route53RecordGeolocationRoutingPolicyOutputReference {
	var returns Route53RecordGeolocationRoutingPolicyOutputReference
	_jsii_.Get(
		j,
		"geolocationRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) GeolocationRoutingPolicyInput() *Route53RecordGeolocationRoutingPolicy {
	var returns *Route53RecordGeolocationRoutingPolicy
	_jsii_.Get(
		j,
		"geolocationRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) GeoproximityRoutingPolicy() Route53RecordGeoproximityRoutingPolicyOutputReference {
	var returns Route53RecordGeoproximityRoutingPolicyOutputReference
	_jsii_.Get(
		j,
		"geoproximityRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) GeoproximityRoutingPolicyInput() *Route53RecordGeoproximityRoutingPolicy {
	var returns *Route53RecordGeoproximityRoutingPolicy
	_jsii_.Get(
		j,
		"geoproximityRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) HealthCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) HealthCheckIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) LatencyRoutingPolicy() Route53RecordLatencyRoutingPolicyOutputReference {
	var returns Route53RecordLatencyRoutingPolicyOutputReference
	_jsii_.Get(
		j,
		"latencyRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) LatencyRoutingPolicyInput() *Route53RecordLatencyRoutingPolicy {
	var returns *Route53RecordLatencyRoutingPolicy
	_jsii_.Get(
		j,
		"latencyRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) MultivalueAnswerRoutingPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multivalueAnswerRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) MultivalueAnswerRoutingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multivalueAnswerRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Records() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"records",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) RecordsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) SetIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) SetIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Timeouts() Route53RecordTimeoutsOutputReference {
	var returns Route53RecordTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) WeightedRoutingPolicy() Route53RecordWeightedRoutingPolicyOutputReference {
	var returns Route53RecordWeightedRoutingPolicyOutputReference
	_jsii_.Get(
		j,
		"weightedRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) WeightedRoutingPolicyInput() *Route53RecordWeightedRoutingPolicy {
	var returns *Route53RecordWeightedRoutingPolicy
	_jsii_.Get(
		j,
		"weightedRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53Record) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/route53_record aws_route53_record} Resource.
func NewRoute53Record(scope constructs.Construct, id *string, config *Route53RecordConfig) Route53Record {
	_init_.Initialize()

	if err := validateNewRoute53RecordParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53Record{}

	_jsii_.Create(
		"@cdktf/provider-aws.route53Record.Route53Record",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/route53_record aws_route53_record} Resource.
func NewRoute53Record_Override(r Route53Record, scope constructs.Construct, id *string, config *Route53RecordConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.route53Record.Route53Record",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Route53Record)SetAllowOverwrite(val interface{}) {
	if err := j.validateSetAllowOverwriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowOverwrite",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetHealthCheckId(val *string) {
	if err := j.validateSetHealthCheckIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckId",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetMultivalueAnswerRoutingPolicy(val interface{}) {
	if err := j.validateSetMultivalueAnswerRoutingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multivalueAnswerRoutingPolicy",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetRecords(val *[]*string) {
	if err := j.validateSetRecordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"records",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetSetIdentifier(val *string) {
	if err := j.validateSetSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setIdentifier",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_Route53Record)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

// Generates CDKTF code for importing a Route53Record resource upon running "cdktf plan <stack-name>".
func Route53Record_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRoute53Record_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53Record.Route53Record",
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
func Route53Record_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53Record_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53Record.Route53Record",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53Record_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53Record_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53Record.Route53Record",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53Record_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53Record_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53Record.Route53Record",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Route53Record_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.route53Record.Route53Record",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Route53Record) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_Route53Record) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Route53Record) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_Route53Record) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53Record) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_Route53Record) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_Route53Record) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_Route53Record) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_Route53Record) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_Route53Record) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_Route53Record) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_Route53Record) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53Record) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_Route53Record) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53Record) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53Record) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_Route53Record) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53Record) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Route53Record) PutAlias(value *Route53RecordAlias) {
	if err := r.validatePutAliasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAlias",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53Record) PutCidrRoutingPolicy(value *Route53RecordCidrRoutingPolicy) {
	if err := r.validatePutCidrRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCidrRoutingPolicy",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53Record) PutFailoverRoutingPolicy(value *Route53RecordFailoverRoutingPolicy) {
	if err := r.validatePutFailoverRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putFailoverRoutingPolicy",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53Record) PutGeolocationRoutingPolicy(value *Route53RecordGeolocationRoutingPolicy) {
	if err := r.validatePutGeolocationRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGeolocationRoutingPolicy",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53Record) PutGeoproximityRoutingPolicy(value *Route53RecordGeoproximityRoutingPolicy) {
	if err := r.validatePutGeoproximityRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGeoproximityRoutingPolicy",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53Record) PutLatencyRoutingPolicy(value *Route53RecordLatencyRoutingPolicy) {
	if err := r.validatePutLatencyRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putLatencyRoutingPolicy",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53Record) PutTimeouts(value *Route53RecordTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53Record) PutWeightedRoutingPolicy(value *Route53RecordWeightedRoutingPolicy) {
	if err := r.validatePutWeightedRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putWeightedRoutingPolicy",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53Record) ResetAlias() {
	_jsii_.InvokeVoid(
		r,
		"resetAlias",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetAllowOverwrite() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowOverwrite",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetCidrRoutingPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetCidrRoutingPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetFailoverRoutingPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetFailoverRoutingPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetGeolocationRoutingPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetGeolocationRoutingPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetGeoproximityRoutingPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetGeoproximityRoutingPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetHealthCheckId() {
	_jsii_.InvokeVoid(
		r,
		"resetHealthCheckId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetLatencyRoutingPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetLatencyRoutingPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetMultivalueAnswerRoutingPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetMultivalueAnswerRoutingPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetRecords() {
	_jsii_.InvokeVoid(
		r,
		"resetRecords",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetSetIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSetIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetTtl() {
	_jsii_.InvokeVoid(
		r,
		"resetTtl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) ResetWeightedRoutingPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetWeightedRoutingPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53Record) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53Record) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53Record) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53Record) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53Record) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53Record) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

