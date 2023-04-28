package dataawss3bucketobjects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/dataawss3bucketobjects/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/s3_bucket_objects aws_s3_bucket_objects}.
type DataAwsS3BucketObjects interface {
	cdktf.TerraformDataSource
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CommonPrefixes() *[]*string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Delimiter() *string
	SetDelimiter(val *string)
	DelimiterInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EncodingType() *string
	SetEncodingType(val *string)
	EncodingTypeInput() *string
	FetchOwner() interface{}
	SetFetchOwner(val interface{})
	FetchOwnerInput() interface{}
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
	Keys() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxKeys() *float64
	SetMaxKeys(val *float64)
	MaxKeysInput() *float64
	// The tree node.
	Node() constructs.Node
	Owners() *[]*string
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	StartAfter() *string
	SetStartAfter(val *string)
	StartAfterInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetDelimiter()
	ResetEncodingType()
	ResetFetchOwner()
	ResetId()
	ResetMaxKeys()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrefix()
	ResetStartAfter()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsS3BucketObjects
type jsiiProxy_DataAwsS3BucketObjects struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsS3BucketObjects) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) CommonPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commonPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) Delimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) DelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) EncodingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) EncodingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) FetchOwner() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) FetchOwnerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) Keys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) MaxKeys() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) MaxKeysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) Owners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"owners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) StartAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) StartAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketObjects) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/s3_bucket_objects aws_s3_bucket_objects} Data Source.
func NewDataAwsS3BucketObjects(scope constructs.Construct, id *string, config *DataAwsS3BucketObjectsConfig) DataAwsS3BucketObjects {
	_init_.Initialize()

	if err := validateNewDataAwsS3BucketObjectsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsS3BucketObjects{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsS3BucketObjects.DataAwsS3BucketObjects",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/s3_bucket_objects aws_s3_bucket_objects} Data Source.
func NewDataAwsS3BucketObjects_Override(d DataAwsS3BucketObjects, scope constructs.Construct, id *string, config *DataAwsS3BucketObjectsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsS3BucketObjects.DataAwsS3BucketObjects",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsS3BucketObjects)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketObjects)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketObjects)SetDelimiter(val *string) {
	if err := j.validateSetDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delimiter",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketObjects)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketObjects)SetEncodingType(val *string) {
	if err := j.validateSetEncodingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encodingType",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketObjects)SetFetchOwner(val interface{}) {
	if err := j.validateSetFetchOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fetchOwner",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketObjects)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketObjects)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketObjects)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketObjects)SetMaxKeys(val *float64) {
	if err := j.validateSetMaxKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxKeys",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketObjects)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketObjects)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketObjects)SetStartAfter(val *string) {
	if err := j.validateSetStartAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startAfter",
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
func DataAwsS3BucketObjects_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsS3BucketObjects_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsS3BucketObjects.DataAwsS3BucketObjects",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsS3BucketObjects_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsS3BucketObjects_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsS3BucketObjects.DataAwsS3BucketObjects",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsS3BucketObjects_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsS3BucketObjects_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsS3BucketObjects.DataAwsS3BucketObjects",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsS3BucketObjects_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsS3BucketObjects.DataAwsS3BucketObjects",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsS3BucketObjects) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsS3BucketObjects) ResetDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsS3BucketObjects) ResetEncodingType() {
	_jsii_.InvokeVoid(
		d,
		"resetEncodingType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsS3BucketObjects) ResetFetchOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetFetchOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsS3BucketObjects) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsS3BucketObjects) ResetMaxKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsS3BucketObjects) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsS3BucketObjects) ResetPrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetPrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsS3BucketObjects) ResetStartAfter() {
	_jsii_.InvokeVoid(
		d,
		"resetStartAfter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsS3BucketObjects) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketObjects) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

