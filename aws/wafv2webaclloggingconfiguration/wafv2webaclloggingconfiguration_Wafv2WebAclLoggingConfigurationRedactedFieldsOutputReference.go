package wafv2webaclloggingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/wafv2webaclloggingconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2WebAclLoggingConfigurationRedactedFieldsAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsAllQueryArguments
	Body() Wafv2WebAclLoggingConfigurationRedactedFieldsBodyOutputReference
	BodyInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsBody
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Method() Wafv2WebAclLoggingConfigurationRedactedFieldsMethodOutputReference
	MethodInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsMethod
	QueryString() Wafv2WebAclLoggingConfigurationRedactedFieldsQueryStringOutputReference
	QueryStringInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsQueryString
	SingleHeader() Wafv2WebAclLoggingConfigurationRedactedFieldsSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsSingleHeader
	SingleQueryArgument() Wafv2WebAclLoggingConfigurationRedactedFieldsSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2WebAclLoggingConfigurationRedactedFieldsUriPathOutputReference
	UriPathInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsUriPath
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
	PutAllQueryArguments(value *Wafv2WebAclLoggingConfigurationRedactedFieldsAllQueryArguments)
	PutBody(value *Wafv2WebAclLoggingConfigurationRedactedFieldsBody)
	PutMethod(value *Wafv2WebAclLoggingConfigurationRedactedFieldsMethod)
	PutQueryString(value *Wafv2WebAclLoggingConfigurationRedactedFieldsQueryString)
	PutSingleHeader(value *Wafv2WebAclLoggingConfigurationRedactedFieldsSingleHeader)
	PutSingleQueryArgument(value *Wafv2WebAclLoggingConfigurationRedactedFieldsSingleQueryArgument)
	PutUriPath(value *Wafv2WebAclLoggingConfigurationRedactedFieldsUriPath)
	ResetAllQueryArguments()
	ResetBody()
	ResetMethod()
	ResetQueryString()
	ResetSingleHeader()
	ResetSingleQueryArgument()
	ResetUriPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference
type jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) AllQueryArguments() Wafv2WebAclLoggingConfigurationRedactedFieldsAllQueryArgumentsOutputReference {
	var returns Wafv2WebAclLoggingConfigurationRedactedFieldsAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) AllQueryArgumentsInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsAllQueryArguments {
	var returns *Wafv2WebAclLoggingConfigurationRedactedFieldsAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) Body() Wafv2WebAclLoggingConfigurationRedactedFieldsBodyOutputReference {
	var returns Wafv2WebAclLoggingConfigurationRedactedFieldsBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) BodyInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsBody {
	var returns *Wafv2WebAclLoggingConfigurationRedactedFieldsBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) Method() Wafv2WebAclLoggingConfigurationRedactedFieldsMethodOutputReference {
	var returns Wafv2WebAclLoggingConfigurationRedactedFieldsMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) MethodInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsMethod {
	var returns *Wafv2WebAclLoggingConfigurationRedactedFieldsMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) QueryString() Wafv2WebAclLoggingConfigurationRedactedFieldsQueryStringOutputReference {
	var returns Wafv2WebAclLoggingConfigurationRedactedFieldsQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) QueryStringInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsQueryString {
	var returns *Wafv2WebAclLoggingConfigurationRedactedFieldsQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) SingleHeader() Wafv2WebAclLoggingConfigurationRedactedFieldsSingleHeaderOutputReference {
	var returns Wafv2WebAclLoggingConfigurationRedactedFieldsSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) SingleHeaderInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsSingleHeader {
	var returns *Wafv2WebAclLoggingConfigurationRedactedFieldsSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) SingleQueryArgument() Wafv2WebAclLoggingConfigurationRedactedFieldsSingleQueryArgumentOutputReference {
	var returns Wafv2WebAclLoggingConfigurationRedactedFieldsSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) SingleQueryArgumentInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsSingleQueryArgument {
	var returns *Wafv2WebAclLoggingConfigurationRedactedFieldsSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) UriPath() Wafv2WebAclLoggingConfigurationRedactedFieldsUriPathOutputReference {
	var returns Wafv2WebAclLoggingConfigurationRedactedFieldsUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) UriPathInput() *Wafv2WebAclLoggingConfigurationRedactedFieldsUriPath {
	var returns *Wafv2WebAclLoggingConfigurationRedactedFieldsUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2WebAclLoggingConfigurationRedactedFieldsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclLoggingConfigurationRedactedFieldsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAclLoggingConfiguration.Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafv2WebAclLoggingConfigurationRedactedFieldsOutputReference_Override(w Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAclLoggingConfiguration.Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) PutAllQueryArguments(value *Wafv2WebAclLoggingConfigurationRedactedFieldsAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) PutBody(value *Wafv2WebAclLoggingConfigurationRedactedFieldsBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) PutMethod(value *Wafv2WebAclLoggingConfigurationRedactedFieldsMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) PutQueryString(value *Wafv2WebAclLoggingConfigurationRedactedFieldsQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) PutSingleHeader(value *Wafv2WebAclLoggingConfigurationRedactedFieldsSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) PutSingleQueryArgument(value *Wafv2WebAclLoggingConfigurationRedactedFieldsSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) PutUriPath(value *Wafv2WebAclLoggingConfigurationRedactedFieldsUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationRedactedFieldsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

