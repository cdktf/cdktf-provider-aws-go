// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalogprovisionedproduct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/servicecatalogprovisionedproduct/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/servicecatalog_provisioned_product aws_servicecatalog_provisioned_product}.
type ServicecatalogProvisionedProduct interface {
	cdktf.TerraformResource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudwatchDashboardNames() *[]*string
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
	CreatedTime() *string
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
	IgnoreErrors() interface{}
	SetIgnoreErrors(val interface{})
	IgnoreErrorsInput() interface{}
	LastProvisioningRecordId() *string
	LastRecordId() *string
	LastSuccessfulProvisioningRecordId() *string
	LaunchRoleArn() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotificationArns() *[]*string
	SetNotificationArns(val *[]*string)
	NotificationArnsInput() *[]*string
	Outputs() ServicecatalogProvisionedProductOutputsList
	PathId() *string
	SetPathId(val *string)
	PathIdInput() *string
	PathName() *string
	SetPathName(val *string)
	PathNameInput() *string
	ProductId() *string
	SetProductId(val *string)
	ProductIdInput() *string
	ProductName() *string
	SetProductName(val *string)
	ProductNameInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProvisioningArtifactId() *string
	SetProvisioningArtifactId(val *string)
	ProvisioningArtifactIdInput() *string
	ProvisioningArtifactName() *string
	SetProvisioningArtifactName(val *string)
	ProvisioningArtifactNameInput() *string
	ProvisioningParameters() ServicecatalogProvisionedProductProvisioningParametersList
	ProvisioningParametersInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RetainPhysicalResources() interface{}
	SetRetainPhysicalResources(val interface{})
	RetainPhysicalResourcesInput() interface{}
	StackSetProvisioningPreferences() ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference
	StackSetProvisioningPreferencesInput() *ServicecatalogProvisionedProductStackSetProvisioningPreferences
	Status() *string
	StatusMessage() *string
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
	Timeouts() ServicecatalogProvisionedProductTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
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
	PutProvisioningParameters(value interface{})
	PutStackSetProvisioningPreferences(value *ServicecatalogProvisionedProductStackSetProvisioningPreferences)
	PutTimeouts(value *ServicecatalogProvisionedProductTimeouts)
	ResetAcceptLanguage()
	ResetId()
	ResetIgnoreErrors()
	ResetNotificationArns()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPathId()
	ResetPathName()
	ResetProductId()
	ResetProductName()
	ResetProvisioningArtifactId()
	ResetProvisioningArtifactName()
	ResetProvisioningParameters()
	ResetRetainPhysicalResources()
	ResetStackSetProvisioningPreferences()
	ResetTags()
	ResetTagsAll()
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

// The jsii proxy struct for ServicecatalogProvisionedProduct
type jsiiProxy_ServicecatalogProvisionedProduct struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) CloudwatchDashboardNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloudwatchDashboardNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) IgnoreErrors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) IgnoreErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) LastProvisioningRecordId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastProvisioningRecordId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) LastRecordId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastRecordId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) LastSuccessfulProvisioningRecordId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSuccessfulProvisioningRecordId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) LaunchRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) NotificationArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Outputs() ServicecatalogProvisionedProductOutputsList {
	var returns ServicecatalogProvisionedProductOutputsList
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) PathId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) PathIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) PathName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) PathNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProductName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProductNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProvisioningArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProvisioningArtifactIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProvisioningArtifactName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProvisioningArtifactNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProvisioningParameters() ServicecatalogProvisionedProductProvisioningParametersList {
	var returns ServicecatalogProvisionedProductProvisioningParametersList
	_jsii_.Get(
		j,
		"provisioningParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProvisioningParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisioningParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) RetainPhysicalResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainPhysicalResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) RetainPhysicalResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainPhysicalResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) StackSetProvisioningPreferences() ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference {
	var returns ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference
	_jsii_.Get(
		j,
		"stackSetProvisioningPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) StackSetProvisioningPreferencesInput() *ServicecatalogProvisionedProductStackSetProvisioningPreferences {
	var returns *ServicecatalogProvisionedProductStackSetProvisioningPreferences
	_jsii_.Get(
		j,
		"stackSetProvisioningPreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) StatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Timeouts() ServicecatalogProvisionedProductTimeoutsOutputReference {
	var returns ServicecatalogProvisionedProductTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/servicecatalog_provisioned_product aws_servicecatalog_provisioned_product} Resource.
func NewServicecatalogProvisionedProduct(scope constructs.Construct, id *string, config *ServicecatalogProvisionedProductConfig) ServicecatalogProvisionedProduct {
	_init_.Initialize()

	if err := validateNewServicecatalogProvisionedProductParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicecatalogProvisionedProduct{}

	_jsii_.Create(
		"@cdktf/provider-aws.servicecatalogProvisionedProduct.ServicecatalogProvisionedProduct",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/servicecatalog_provisioned_product aws_servicecatalog_provisioned_product} Resource.
func NewServicecatalogProvisionedProduct_Override(s ServicecatalogProvisionedProduct, scope constructs.Construct, id *string, config *ServicecatalogProvisionedProductConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.servicecatalogProvisionedProduct.ServicecatalogProvisionedProduct",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetAcceptLanguage(val *string) {
	if err := j.validateSetAcceptLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetIgnoreErrors(val interface{}) {
	if err := j.validateSetIgnoreErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreErrors",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetNotificationArns(val *[]*string) {
	if err := j.validateSetNotificationArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationArns",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetPathId(val *string) {
	if err := j.validateSetPathIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetPathName(val *string) {
	if err := j.validateSetPathNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathName",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetProductId(val *string) {
	if err := j.validateSetProductIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetProductName(val *string) {
	if err := j.validateSetProductNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productName",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetProvisioningArtifactId(val *string) {
	if err := j.validateSetProvisioningArtifactIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningArtifactId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetProvisioningArtifactName(val *string) {
	if err := j.validateSetProvisioningArtifactNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningArtifactName",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetRetainPhysicalResources(val interface{}) {
	if err := j.validateSetRetainPhysicalResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainPhysicalResources",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a ServicecatalogProvisionedProduct resource upon running "cdktf plan <stack-name>".
func ServicecatalogProvisionedProduct_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServicecatalogProvisionedProduct_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.servicecatalogProvisionedProduct.ServicecatalogProvisionedProduct",
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
func ServicecatalogProvisionedProduct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicecatalogProvisionedProduct_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.servicecatalogProvisionedProduct.ServicecatalogProvisionedProduct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicecatalogProvisionedProduct_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicecatalogProvisionedProduct_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.servicecatalogProvisionedProduct.ServicecatalogProvisionedProduct",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicecatalogProvisionedProduct_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicecatalogProvisionedProduct_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.servicecatalogProvisionedProduct.ServicecatalogProvisionedProduct",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogProvisionedProduct_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.servicecatalogProvisionedProduct.ServicecatalogProvisionedProduct",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServicecatalogProvisionedProduct) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServicecatalogProvisionedProduct) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) PutProvisioningParameters(value interface{}) {
	if err := s.validatePutProvisioningParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProvisioningParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) PutStackSetProvisioningPreferences(value *ServicecatalogProvisionedProductStackSetProvisioningPreferences) {
	if err := s.validatePutStackSetProvisioningPreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStackSetProvisioningPreferences",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) PutTimeouts(value *ServicecatalogProvisionedProductTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetIgnoreErrors() {
	_jsii_.InvokeVoid(
		s,
		"resetIgnoreErrors",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetNotificationArns() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationArns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetPathId() {
	_jsii_.InvokeVoid(
		s,
		"resetPathId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetPathName() {
	_jsii_.InvokeVoid(
		s,
		"resetPathName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetProductId() {
	_jsii_.InvokeVoid(
		s,
		"resetProductId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetProductName() {
	_jsii_.InvokeVoid(
		s,
		"resetProductName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetProvisioningArtifactId() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisioningArtifactId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetProvisioningArtifactName() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisioningArtifactName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetProvisioningParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisioningParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetRetainPhysicalResources() {
	_jsii_.InvokeVoid(
		s,
		"resetRetainPhysicalResources",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetStackSetProvisioningPreferences() {
	_jsii_.InvokeVoid(
		s,
		"resetStackSetProvisioningPreferences",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

