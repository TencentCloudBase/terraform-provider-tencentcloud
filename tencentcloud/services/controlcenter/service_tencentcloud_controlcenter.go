// Code generated by iacg; DO NOT EDIT.
package controlcenter

import "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/connectivity"

func NewControlcenterService(client *connectivity.TencentCloudClient) ControlcenterService {
	return ControlcenterService{client: client}
}

type ControlcenterService struct {
	client *connectivity.TencentCloudClient
}
