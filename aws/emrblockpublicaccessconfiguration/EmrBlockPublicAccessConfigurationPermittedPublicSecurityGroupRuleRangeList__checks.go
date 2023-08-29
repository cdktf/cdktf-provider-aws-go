// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package emrblockpublicaccessconfiguration

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRange:
		val := val.(*[]*EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRange)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRange:
		val_ := val.([]*EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRange)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRange; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewEmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

