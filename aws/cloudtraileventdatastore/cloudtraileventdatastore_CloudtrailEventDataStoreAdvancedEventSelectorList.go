package cloudtraileventdatastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/cloudtraileventdatastore/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudtrailEventDataStoreAdvancedEventSelectorList interface {
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
	Get(index *float64) CloudtrailEventDataStoreAdvancedEventSelectorOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudtrailEventDataStoreAdvancedEventSelectorList
type jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCloudtrailEventDataStoreAdvancedEventSelectorList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CloudtrailEventDataStoreAdvancedEventSelectorList {
	_init_.Initialize()

	if err := validateNewCloudtrailEventDataStoreAdvancedEventSelectorListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudtrailEventDataStore.CloudtrailEventDataStoreAdvancedEventSelectorList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCloudtrailEventDataStoreAdvancedEventSelectorList_Override(c CloudtrailEventDataStoreAdvancedEventSelectorList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudtrailEventDataStore.CloudtrailEventDataStoreAdvancedEventSelectorList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList) Get(index *float64) CloudtrailEventDataStoreAdvancedEventSelectorOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns CloudtrailEventDataStoreAdvancedEventSelectorOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
