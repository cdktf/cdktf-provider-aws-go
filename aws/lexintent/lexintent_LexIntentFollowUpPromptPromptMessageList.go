package lexintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/lexintent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LexIntentFollowUpPromptPromptMessageList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) LexIntentFollowUpPromptPromptMessageOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LexIntentFollowUpPromptPromptMessageList
type jsiiProxy_LexIntentFollowUpPromptPromptMessageList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptMessageList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptMessageList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptMessageList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptMessageList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptMessageList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptMessageList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewLexIntentFollowUpPromptPromptMessageList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) LexIntentFollowUpPromptPromptMessageList {
	_init_.Initialize()

	if err := validateNewLexIntentFollowUpPromptPromptMessageListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LexIntentFollowUpPromptPromptMessageList{}

	_jsii_.Create(
		"@cdktf/provider-aws.lexIntent.LexIntentFollowUpPromptPromptMessageList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewLexIntentFollowUpPromptPromptMessageList_Override(l LexIntentFollowUpPromptPromptMessageList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lexIntent.LexIntentFollowUpPromptPromptMessageList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		l,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptMessageList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptMessageList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptMessageList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptMessageList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (l *jsiiProxy_LexIntentFollowUpPromptPromptMessageList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexIntentFollowUpPromptPromptMessageList) Get(index *float64) LexIntentFollowUpPromptPromptMessageOutputReference {
	if err := l.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns LexIntentFollowUpPromptPromptMessageOutputReference

	_jsii_.Invoke(
		l,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexIntentFollowUpPromptPromptMessageList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexIntentFollowUpPromptPromptMessageList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

