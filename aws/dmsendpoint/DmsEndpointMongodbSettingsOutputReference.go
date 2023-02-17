package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/dmsendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsEndpointMongodbSettingsOutputReference interface {
	cdktf.ComplexObject
	AuthMechanism() *string
	SetAuthMechanism(val *string)
	AuthMechanismInput() *string
	AuthSource() *string
	SetAuthSource(val *string)
	AuthSourceInput() *string
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
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
	DocsToInvestigate() *string
	SetDocsToInvestigate(val *string)
	DocsToInvestigateInput() *string
	ExtractDocId() *string
	SetExtractDocId(val *string)
	ExtractDocIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DmsEndpointMongodbSettings
	SetInternalValue(val *DmsEndpointMongodbSettings)
	NestingLevel() *string
	SetNestingLevel(val *string)
	NestingLevelInput() *string
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
	ResetAuthMechanism()
	ResetAuthSource()
	ResetAuthType()
	ResetDocsToInvestigate()
	ResetExtractDocId()
	ResetNestingLevel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsEndpointMongodbSettingsOutputReference
type jsiiProxy_DmsEndpointMongodbSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) AuthMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) AuthMechanismInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) AuthSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) AuthSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) DocsToInvestigate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsToInvestigate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) DocsToInvestigateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsToInvestigateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ExtractDocId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extractDocId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ExtractDocIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extractDocIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) InternalValue() *DmsEndpointMongodbSettings {
	var returns *DmsEndpointMongodbSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) NestingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nestingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) NestingLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nestingLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsEndpointMongodbSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DmsEndpointMongodbSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsEndpointMongodbSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpointMongodbSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpointMongodbSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsEndpointMongodbSettingsOutputReference_Override(d DmsEndpointMongodbSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpointMongodbSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference)SetAuthMechanism(val *string) {
	if err := j.validateSetAuthMechanismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authMechanism",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference)SetAuthSource(val *string) {
	if err := j.validateSetAuthSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authSource",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference)SetDocsToInvestigate(val *string) {
	if err := j.validateSetDocsToInvestigateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"docsToInvestigate",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference)SetExtractDocId(val *string) {
	if err := j.validateSetExtractDocIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extractDocId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference)SetInternalValue(val *DmsEndpointMongodbSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference)SetNestingLevel(val *string) {
	if err := j.validateSetNestingLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nestingLevel",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ResetAuthMechanism() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthMechanism",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ResetAuthSource() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ResetAuthType() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ResetDocsToInvestigate() {
	_jsii_.InvokeVoid(
		d,
		"resetDocsToInvestigate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ResetExtractDocId() {
	_jsii_.InvokeVoid(
		d,
		"resetExtractDocId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ResetNestingLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetNestingLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

