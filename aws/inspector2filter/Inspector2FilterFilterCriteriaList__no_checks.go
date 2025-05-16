// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package inspector2filter

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Inspector2FilterFilterCriteriaList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewInspector2FilterFilterCriteriaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

