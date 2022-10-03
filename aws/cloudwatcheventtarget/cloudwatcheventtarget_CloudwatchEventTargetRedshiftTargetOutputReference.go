package cloudwatcheventtarget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/cloudwatcheventtarget/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchEventTargetRedshiftTargetOutputReference interface {
	cdktf.ComplexObject
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
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	DbUser() *string
	SetDbUser(val *string)
	DbUserInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CloudwatchEventTargetRedshiftTarget
	SetInternalValue(val *CloudwatchEventTargetRedshiftTarget)
	SecretsManagerArn() *string
	SetSecretsManagerArn(val *string)
	SecretsManagerArnInput() *string
	Sql() *string
	SetSql(val *string)
	SqlInput() *string
	StatementName() *string
	SetStatementName(val *string)
	StatementNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WithEvent() interface{}
	SetWithEvent(val interface{})
	WithEventInput() interface{}
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
	ResetDbUser()
	ResetSecretsManagerArn()
	ResetSql()
	ResetStatementName()
	ResetWithEvent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudwatchEventTargetRedshiftTargetOutputReference
type jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) DbUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) DbUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) InternalValue() *CloudwatchEventTargetRedshiftTarget {
	var returns *CloudwatchEventTargetRedshiftTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) SecretsManagerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) SecretsManagerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) Sql() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) SqlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) StatementName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) StatementNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) WithEvent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) WithEventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withEventInput",
		&returns,
	)
	return returns
}


func NewCloudwatchEventTargetRedshiftTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudwatchEventTargetRedshiftTargetOutputReference {
	_init_.Initialize()

	if err := validateNewCloudwatchEventTargetRedshiftTargetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchEventTarget.CloudwatchEventTargetRedshiftTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudwatchEventTargetRedshiftTargetOutputReference_Override(c CloudwatchEventTargetRedshiftTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchEventTarget.CloudwatchEventTargetRedshiftTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference)SetDbUser(val *string) {
	if err := j.validateSetDbUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbUser",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference)SetInternalValue(val *CloudwatchEventTargetRedshiftTarget) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference)SetSecretsManagerArn(val *string) {
	if err := j.validateSetSecretsManagerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerArn",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference)SetSql(val *string) {
	if err := j.validateSetSqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sql",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference)SetStatementName(val *string) {
	if err := j.validateSetStatementNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementName",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference)SetWithEvent(val interface{}) {
	if err := j.validateSetWithEventParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withEvent",
		val,
	)
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) ResetDbUser() {
	_jsii_.InvokeVoid(
		c,
		"resetDbUser",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) ResetSecretsManagerArn() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretsManagerArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) ResetSql() {
	_jsii_.InvokeVoid(
		c,
		"resetSql",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) ResetStatementName() {
	_jsii_.InvokeVoid(
		c,
		"resetStatementName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) ResetWithEvent() {
	_jsii_.InvokeVoid(
		c,
		"resetWithEvent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetRedshiftTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

