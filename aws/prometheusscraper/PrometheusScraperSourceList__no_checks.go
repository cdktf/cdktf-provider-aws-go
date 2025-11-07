// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package prometheusscraper

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrometheusScraperSourceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PrometheusScraperSourceList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrometheusScraperSourceList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperSourceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperSourceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperSourceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperSourceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrometheusScraperSourceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

