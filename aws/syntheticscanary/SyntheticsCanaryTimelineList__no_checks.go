// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package syntheticscanary

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsCanaryTimelineList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SyntheticsCanaryTimelineList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsCanaryTimelineList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SyntheticsCanaryTimelineList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SyntheticsCanaryTimelineList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SyntheticsCanaryTimelineList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSyntheticsCanaryTimelineListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

