package quicksightdataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/quicksightdataset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightDataSetPhysicalTableMapOutputReference interface {
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
	CustomSql() QuicksightDataSetPhysicalTableMapCustomSqlOutputReference
	CustomSqlInput() *QuicksightDataSetPhysicalTableMapCustomSql
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PhysicalTableMapId() *string
	SetPhysicalTableMapId(val *string)
	PhysicalTableMapIdInput() *string
	RelationalTable() QuicksightDataSetPhysicalTableMapRelationalTableOutputReference
	RelationalTableInput() *QuicksightDataSetPhysicalTableMapRelationalTable
	S3Source() QuicksightDataSetPhysicalTableMapS3SourceOutputReference
	S3SourceInput() *QuicksightDataSetPhysicalTableMapS3Source
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
	PutCustomSql(value *QuicksightDataSetPhysicalTableMapCustomSql)
	PutRelationalTable(value *QuicksightDataSetPhysicalTableMapRelationalTable)
	PutS3Source(value *QuicksightDataSetPhysicalTableMapS3Source)
	ResetCustomSql()
	ResetRelationalTable()
	ResetS3Source()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDataSetPhysicalTableMapOutputReference
type jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) CustomSql() QuicksightDataSetPhysicalTableMapCustomSqlOutputReference {
	var returns QuicksightDataSetPhysicalTableMapCustomSqlOutputReference
	_jsii_.Get(
		j,
		"customSql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) CustomSqlInput() *QuicksightDataSetPhysicalTableMapCustomSql {
	var returns *QuicksightDataSetPhysicalTableMapCustomSql
	_jsii_.Get(
		j,
		"customSqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) PhysicalTableMapId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalTableMapId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) PhysicalTableMapIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalTableMapIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) RelationalTable() QuicksightDataSetPhysicalTableMapRelationalTableOutputReference {
	var returns QuicksightDataSetPhysicalTableMapRelationalTableOutputReference
	_jsii_.Get(
		j,
		"relationalTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) RelationalTableInput() *QuicksightDataSetPhysicalTableMapRelationalTable {
	var returns *QuicksightDataSetPhysicalTableMapRelationalTable
	_jsii_.Get(
		j,
		"relationalTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) S3Source() QuicksightDataSetPhysicalTableMapS3SourceOutputReference {
	var returns QuicksightDataSetPhysicalTableMapS3SourceOutputReference
	_jsii_.Get(
		j,
		"s3Source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) S3SourceInput() *QuicksightDataSetPhysicalTableMapS3Source {
	var returns *QuicksightDataSetPhysicalTableMapS3Source
	_jsii_.Get(
		j,
		"s3SourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightDataSetPhysicalTableMapOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QuicksightDataSetPhysicalTableMapOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSetPhysicalTableMapOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDataSet.QuicksightDataSetPhysicalTableMapOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQuicksightDataSetPhysicalTableMapOutputReference_Override(q QuicksightDataSetPhysicalTableMapOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDataSet.QuicksightDataSetPhysicalTableMapOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference)SetPhysicalTableMapId(val *string) {
	if err := j.validateSetPhysicalTableMapIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"physicalTableMapId",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) PutCustomSql(value *QuicksightDataSetPhysicalTableMapCustomSql) {
	if err := q.validatePutCustomSqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putCustomSql",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) PutRelationalTable(value *QuicksightDataSetPhysicalTableMapRelationalTable) {
	if err := q.validatePutRelationalTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRelationalTable",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) PutS3Source(value *QuicksightDataSetPhysicalTableMapS3Source) {
	if err := q.validatePutS3SourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putS3Source",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) ResetCustomSql() {
	_jsii_.InvokeVoid(
		q,
		"resetCustomSql",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) ResetRelationalTable() {
	_jsii_.InvokeVoid(
		q,
		"resetRelationalTable",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) ResetS3Source() {
	_jsii_.InvokeVoid(
		q,
		"resetS3Source",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetPhysicalTableMapOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

