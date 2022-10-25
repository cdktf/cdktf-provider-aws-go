//go:build !no_runtime_type_checking

package dataawscloudfrontcachepolicy

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

