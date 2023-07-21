package servicecatalogportfolioshare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/servicecatalogportfolioshare/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/servicecatalog_portfolio_share aws_servicecatalog_portfolio_share}.
type ServicecatalogPortfolioShare interface {
	cdktf.TerraformResource
	Accepted() cdktf.IResolvable
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
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
	// The tree node.
	Node() constructs.Node
	PortfolioId() *string
	SetPortfolioId(val *string)
	PortfolioIdInput() *string
	PrincipalId() *string
	SetPrincipalId(val *string)
	PrincipalIdInput() *string
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
	SharePrincipals() interface{}
	SetSharePrincipals(val interface{})
	SharePrincipalsInput() interface{}
	ShareTagOptions() interface{}
	SetShareTagOptions(val interface{})
	ShareTagOptionsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ServicecatalogPortfolioShareTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	WaitForAcceptance() interface{}
	SetWaitForAcceptance(val interface{})
	WaitForAcceptanceInput() interface{}
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
	PutTimeouts(value *ServicecatalogPortfolioShareTimeouts)
	ResetAcceptLanguage()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSharePrincipals()
	ResetShareTagOptions()
	ResetTimeouts()
	ResetWaitForAcceptance()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicecatalogPortfolioShare
type jsiiProxy_ServicecatalogPortfolioShare struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Accepted() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"accepted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) PortfolioIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) PrincipalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) SharePrincipals() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharePrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) SharePrincipalsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharePrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) ShareTagOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareTagOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) ShareTagOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareTagOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Timeouts() ServicecatalogPortfolioShareTimeoutsOutputReference {
	var returns ServicecatalogPortfolioShareTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) WaitForAcceptance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForAcceptance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) WaitForAcceptanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForAcceptanceInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/servicecatalog_portfolio_share aws_servicecatalog_portfolio_share} Resource.
func NewServicecatalogPortfolioShare(scope constructs.Construct, id *string, config *ServicecatalogPortfolioShareConfig) ServicecatalogPortfolioShare {
	_init_.Initialize()

	if err := validateNewServicecatalogPortfolioShareParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicecatalogPortfolioShare{}

	_jsii_.Create(
		"@cdktf/provider-aws.servicecatalogPortfolioShare.ServicecatalogPortfolioShare",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/servicecatalog_portfolio_share aws_servicecatalog_portfolio_share} Resource.
func NewServicecatalogPortfolioShare_Override(s ServicecatalogPortfolioShare, scope constructs.Construct, id *string, config *ServicecatalogPortfolioShareConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.servicecatalogPortfolioShare.ServicecatalogPortfolioShare",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetAcceptLanguage(val *string) {
	if err := j.validateSetAcceptLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetPortfolioId(val *string) {
	if err := j.validateSetPortfolioIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetPrincipalId(val *string) {
	if err := j.validateSetPrincipalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principalId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetSharePrincipals(val interface{}) {
	if err := j.validateSetSharePrincipalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharePrincipals",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetShareTagOptions(val interface{}) {
	if err := j.validateSetShareTagOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareTagOptions",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare)SetWaitForAcceptance(val interface{}) {
	if err := j.validateSetWaitForAcceptanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForAcceptance",
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
func ServicecatalogPortfolioShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicecatalogPortfolioShare_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.servicecatalogPortfolioShare.ServicecatalogPortfolioShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicecatalogPortfolioShare_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicecatalogPortfolioShare_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.servicecatalogPortfolioShare.ServicecatalogPortfolioShare",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicecatalogPortfolioShare_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicecatalogPortfolioShare_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.servicecatalogPortfolioShare.ServicecatalogPortfolioShare",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogPortfolioShare_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.servicecatalogPortfolioShare.ServicecatalogPortfolioShare",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServicecatalogPortfolioShare) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServicecatalogPortfolioShare) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServicecatalogPortfolioShare) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServicecatalogPortfolioShare) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServicecatalogPortfolioShare) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServicecatalogPortfolioShare) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServicecatalogPortfolioShare) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServicecatalogPortfolioShare) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServicecatalogPortfolioShare) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServicecatalogPortfolioShare) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) PutTimeouts(value *ServicecatalogPortfolioShareTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) ResetSharePrincipals() {
	_jsii_.InvokeVoid(
		s,
		"resetSharePrincipals",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) ResetShareTagOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetShareTagOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) ResetWaitForAcceptance() {
	_jsii_.InvokeVoid(
		s,
		"resetWaitForAcceptance",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

