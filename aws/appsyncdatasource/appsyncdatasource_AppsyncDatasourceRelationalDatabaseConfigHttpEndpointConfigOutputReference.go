package appsyncdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/appsyncdatasource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference interface {
	cdktf.ComplexObject
	AwsSecretStoreArn() *string
	SetAwsSecretStoreArn(val *string)
	AwsSecretStoreArnInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DbClusterIdentifier() *string
	SetDbClusterIdentifier(val *string)
	DbClusterIdentifierInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig
	SetInternalValue(val *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig)
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDatabaseName()
	ResetRegion()
	ResetSchema()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference
type jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) AwsSecretStoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretStoreArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) AwsSecretStoreArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretStoreArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) DbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) DbClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) InternalValue() *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig {
	var returns *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncDatasource.AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference_Override(a AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncDatasource.AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference)SetAwsSecretStoreArn(val *string) {
	if err := j.validateSetAwsSecretStoreArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsSecretStoreArn",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference)SetDbClusterIdentifier(val *string) {
	if err := j.validateSetDbClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference)SetInternalValue(val *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) ResetSchema() {
	_jsii_.InvokeVoid(
		a,
		"resetSchema",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

