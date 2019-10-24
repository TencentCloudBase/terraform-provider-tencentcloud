package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func resourceTencentCloudContainerCluster() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "This resource has been deprecated in Terraform TencentCloud provider version 1.16.0. Please use 'tencentcloud_kubernetes_cluster' instead.",
		Create:             resourceTencentCloudContainerClusterCreate,
		Read:               resourceTencentCloudContainerClusterRead,
		Update:             resourceTencentCloudContainerClusterUpdate,
		Delete:             resourceTencentCloudContainerClusterDelete,

		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: "It has been deprecated from version 1.16.0. Set 'instance_type' instead.",
			},
			"mem": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: "It has been deprecated from version 1.16.0. Set 'instance_type' instead.",
			},
			"os_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bandwidth": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"bandwidth_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"require_wan_ip": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_vpc_gateway": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"storage_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"storage_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"root_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"root_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"goods_num": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_cidr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_desc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cvm_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sg_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mount_target": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"docker_graph_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_script": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unschedulable": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"instance_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"kubernetes_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nodes_num": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"nodes_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_cpu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_mem": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceTencentCloudContainerClusterUpdate(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.tencentcloud_container_cluster.update")()

	return fmt.Errorf("the container cluster resource doesn't support update")
}

func resourceTencentCloudContainerClusterCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_container_cluster.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)
	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}
	cvmService := CvmService{client: meta.(*TencentCloudClient).apiV3Conn}

	runInstancesPara := cvm.NewRunInstancesRequest()

	var place cvm.Placement
	if v, ok := d.GetOkExists("zone_id"); ok {
		var zones []*cvm.ZoneInfo
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			var e error
			zones, e = cvmService.DescribeZones(ctx)
			if e != nil {
				return retryError(e, "InternalError")
			}
			return nil
		})
		if err != nil {
			return err
		}
		zone := ""
		for _, z := range zones {
			if *z.ZoneId == v.(string) {
				zone = *z.Zone
				break
			}
		}
		place.Zone = stringToPointer(zone)
	}
	runInstancesPara.Placement = &place

	var basic ClusterBasicSetting
	var cAdvanced ClusterAdvancedSettings
	var cvms RunInstancesForNode
	var iAdvanced InstanceAdvancedSettings
	var cidrSet ClusterCidrSettings

	basic.ClusterName = d.Get("cluster_name").(string)

	if clusterDescRaw, ok := d.GetOkExists("cluster_desc"); ok {
		if len(clusterDescRaw.(string)) > 0 {
			basic.ClusterDescription = clusterDescRaw.(string)
		}
	}

	if osNameRaw, ok := d.GetOkExists("os_name"); ok {
		if len(osNameRaw.(string)) > 0 {
			basic.ClusterOs = osNameRaw.(string)
		}
	}

	if clusterVersionRaw, ok := d.GetOkExists("cluster_version"); ok {
		if len(clusterVersionRaw.(string)) > 0 {
			basic.ClusterVersion = clusterVersionRaw.(string)
		}
	}

	if v, ok := d.GetOkExists("vpc_id"); ok {
		if len(v.(string)) > 0 {
			vpcId := v.(string)
			subnetId := ""
			asVpcGateway := false
			if subnetIdRaw, ok := d.GetOkExists("subnet_id"); ok {
				subnetId = subnetIdRaw.(string)
			}
			if isVpcGatewayRaw, ok := d.GetOkExists("is_vpc_gateway"); ok {
				if isVpcGatewayRaw.(int) == 1 {
					asVpcGateway = true
				}
			}
			runInstancesPara.VirtualPrivateCloud = &cvm.VirtualPrivateCloud{
				VpcId:        common.StringPtr(vpcId),
				SubnetId:     common.StringPtr(subnetId),
				AsVpcGateway: common.BoolPtr(asVpcGateway),
			}
			basic.VpcId = vpcId
		}
	}

	if instanceTypeRaw, ok := d.GetOkExists("instance_type"); ok {
		if len(instanceTypeRaw.(string)) > 0 {
			runInstancesPara.InstanceType = common.StringPtr(instanceTypeRaw.(string))
		}
	}

	if v, ok := d.GetOkExists("require_wan_ip"); ok {
		publicIpAssigned := false
		internetMaxBandwidthOut := int64(0)
		internetChargeType := ""
		if v.(int) == 1 {
			publicIpAssigned = true
			if v, ok := d.GetOkExists("bandwidth"); ok {
				internetMaxBandwidthOut = int64(v.(int))
			}
			if v, ok := d.GetOkExists("bandwidth_type"); ok {
				bandwidthTypes := map[string]string{
					"PayByMonth":   "BANDWIDTH_PREPAID",
					"PayByTraffic": "TRAFFIC_POSTPAID_BY_HOUR",
					"PayByHour":    "TRAFFIC_POSTPAID_BY_HOUR",
				}
				if v, ok := bandwidthTypes[v.(string)]; ok {
					internetChargeType = v
				}
			}
		}
		runInstancesPara.InternetAccessible = &cvm.InternetAccessible{
			PublicIpAssigned:        common.BoolPtr(publicIpAssigned),
			InternetMaxBandwidthOut: common.Int64Ptr(internetMaxBandwidthOut),
			InternetChargeType:      common.StringPtr(internetChargeType),
		}
	}

	if v, ok := d.GetOkExists("instance_name"); ok {
		runInstancesPara.InstanceName = common.StringPtr(v.(string))
	}

	if v, ok := d.GetOkExists("goods_num"); ok {
		runInstancesPara.InstanceCount = common.Int64Ptr(int64(v.(int)))
	}

	if v, ok := d.GetOkExists("sg_id"); ok {
		runInstancesPara.SecurityGroupIds = []*string{common.StringPtr(v.(string))}
	}

	if v, ok := d.GetOkExists("password"); ok {
		runInstancesPara.LoginSettings = &cvm.LoginSettings{
			Password: common.StringPtr(v.(string)),
		}
	}

	if v, ok := d.GetOkExists("key_id"); ok {
		runInstancesPara.LoginSettings = &cvm.LoginSettings{
			KeyIds: []*string{common.StringPtr(v.(string))},
		}
	}

	runInstancesPara.SystemDisk = &cvm.SystemDisk{}
	if v, ok := d.GetOkExists("root_size"); ok {
		runInstancesPara.SystemDisk.DiskSize = common.Int64Ptr(int64(v.(int)))
	}

	if v, ok := d.GetOkExists("root_type"); ok {
		runInstancesPara.SystemDisk.DiskType = common.StringPtr(v.(string))
	}

	if v, ok := d.GetOkExists("storage_size"); ok {
		if v.(int) > 0 {
			dataDisk := &cvm.DataDisk{
				DiskSize: common.Int64Ptr(int64(v.(int))),
				DiskType: common.StringPtr("CLOUD_PREMIUM"),
			}
			if v, ok := d.GetOkExists("storage_type"); ok {
				if len(v.(string)) > 0 {
					dataDisk.DiskType = common.StringPtr(v.(string))
				}
			}
			runInstancesPara.DataDisks = []*cvm.DataDisk{dataDisk}
		}
	}

	if v, ok := d.GetOkExists("cvm_type"); ok {
		cvmTypes := map[string]string{
			"PayByHour":  "POSTPAID_BY_HOUR",
			"PayByMonth": "PREPAID",
		}
		if vv, ok := cvmTypes[v.(string)]; ok {
			runInstancesPara.InstanceChargeType = common.StringPtr(vv)
		}
	}

	if v, ok := d.GetOkExists("period"); ok {
		runInstancesPara.InstanceChargePrepaid = &cvm.InstanceChargePrepaid{
			Period: common.Int64Ptr(int64(v.(int))),
		}
	}

	if clusterCidrRaw, ok := d.GetOkExists("cluster_cidr"); ok {
		cidrSet.ClusterCidr = clusterCidrRaw.(string)
	}

	if v, ok := d.GetOkExists("mount_target"); ok {
		iAdvanced.MountTarget = v.(string)
	}

	if v, ok := d.GetOkExists("docker_graph_path"); ok {
		iAdvanced.DockerGraphPath = v.(string)
	}

	if v, ok := d.GetOkExists("user_script"); ok {
		iAdvanced.UserScript = v.(string)
	}

	if v, ok := d.GetOkExists("unschedulable"); ok {
		iAdvanced.Unschedulable = int64(v.(int))
	}

	runInstancesParas := runInstancesPara.ToJsonString()
	cvms.Work = []string{runInstancesParas}

	id, err := service.CreateCluster(ctx, basic, cAdvanced, cvms, iAdvanced, cidrSet)
	if err != nil {
		return err
	}

	d.SetId(id)
	err = resource.Retry(6*writeRetryTimeout, func() *resource.RetryError {
		_, _, err = service.DescribeClusterInstances(ctx, d.Id())
		if e, ok := err.(*errors.TencentCloudSDKError); ok {
			if e.GetCode() == "InternalError.ClusterNotFound" {
				return nil
			}
		}
		if err != nil {
			return resource.RetryableError(err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	if err = resourceTencentCloudContainerClusterRead(d, meta); err != nil {
		log.Printf("[WARN]%s resource.tencentcloud_container_cluster.read after create fail , %s", logId, err.Error())
	}

	return nil
}

func resourceTencentCloudContainerClusterRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_container_cluster.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)
	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	var info ClusterInfo
	var has bool

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		var e error
		info, has, e = service.DescribeCluster(ctx, d.Id())
		if e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		return nil
	}

	if !has {
		d.SetId("")
		return nil
	}

	d.Set("cluster_name", info.ClusterName)
	d.Set("cluster_desc", info.ClusterDescription)
	d.Set("kubernetes_version", info.ClusterVersion)
	d.Set("vpc_id", info.VpcId)
	d.Set("cluster_cidr", info.ClusterCidr)
	d.Set("nodes_num", info.ClusterNodeNum)
	d.Set("nodes_status", info.ClusterStatus)
	d.Set("total_cpu", 0)
	d.Set("total_mem", 0)

	var workers []InstanceInfo
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		var e error
		_, workers, e = service.DescribeClusterInstances(ctx, info.ClusterId)
		if e, ok := e.(*errors.TencentCloudSDKError); ok {
			if e.GetCode() == "InternalError.ClusterNotFound" {
				d.SetId("")
				return nil
			}
		}
		if e != nil {
			return resource.RetryableError(e)
		}
		return nil
	})
	if err != nil {
		return err
	}

	instanceIds := []*string{}
	for _, v := range workers {
		instanceIds = append(instanceIds, common.StringPtr(v.InstanceId))
	}

	if len(instanceIds) > 0 {
		describeInstancesreq := cvm.NewDescribeInstancesRequest()
		describeInstancesreq.InstanceIds = instanceIds
		var describeInstancesResponse *cvm.DescribeInstancesResponse
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCvmClient().DescribeInstances(describeInstancesreq)
			if e != nil {
				log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
					logId, describeInstancesreq.GetAction(), describeInstancesreq.ToJsonString(), e.Error())
				return retryError(e)
			}
			describeInstancesResponse = result
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s DescribeInstances failed, reason:%s\n ", logId, err.Error())
			return err
		}

		total_cpu := int64(0)
		total_mem := int64(0)
		for _, v := range describeInstancesResponse.Response.InstanceSet {
			total_cpu += *v.CPU
			total_mem += *v.Memory
		}

		d.Set("total_cpu", total_cpu)
		d.Set("total_mem", total_mem)
	}

	return nil
}

func resourceTencentCloudContainerClusterDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_container_cluster.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)
	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	return resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		err := service.DeleteCluster(ctx, d.Id())
		if err != nil {
			return retryError(err)
		}
		return nil
	})
}
