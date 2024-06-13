// Code generated by iacg; DO NOT EDIT.
package tke

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tke "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20180525"

	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func ResourceTencentCloudKubernetesScaleWorker() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudKubernetesScaleWorkerCreate,
		Read:   resourceTencentCloudKubernetesScaleWorkerRead,
		Delete: resourceTencentCloudKubernetesScaleWorkerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: customScaleWorkerResourceImporter,
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the cluster.",
			},

			"data_disk": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				MaxItems:    11,
				Description: "Configurations of data disk.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_format_and_mount": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Default:     false,
							Description: "Indicate whether to auto format and mount or not. Default is `false`.",
						},
						"disk_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Default:     0,
							Description: "Volume of disk in GB. Default is `0`.",
						},
						"disk_type": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Default:     "CLOUD_PREMIUM",
							Description: "Types of disk, available values: `CLOUD_PREMIUM` and `CLOUD_SSD` and `CLOUD_HSSD` and `CLOUD_TSSD`.",
						},
						"file_system": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Default:     "",
							Description: "File system, e.g. `ext3/ext4/xfs`.",
						},
						"mount_target": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Default:     "",
							Description: "Mount target.",
						},
					},
				},
			},

			"desired_pod_num": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.",
			},

			"docker_graph_path": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Docker graph path. Default is `/var/lib/docker`.",
			},

			"extra_args": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Custom parameter information related to the node.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"gpu_args": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "GPU driver parameters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cuda": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "CUDA  version. Format like: `{ version: String, name: String }`. `version`: Version of GPU driver or CUDA; `name`: Name of GPU driver or CUDA.",
						},
						"cudnn": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "cuDNN version. Format like: `{ version: String, name: String, doc_name: String, dev_name: String }`. `version`: cuDNN version; `name`: cuDNN name; `doc_name`: Doc name of cuDNN; `dev_name`: Dev name of cuDNN.",
						},
						"custom_driver": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Custom GPU driver. Format like: `{address: String}`. `address`: URL of custom GPU driver address.",
						},
						"driver": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "GPU driver version. Format like: `{ version: String, name: String }`. `version`: Version of GPU driver or CUDA; `name`: Name of GPU driver or CUDA.",
						},
						"mig_enable": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether to enable MIG.",
						},
					},
				},
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Labels of kubernetes scale worker created nodes.",
			},

			"mount_target": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Mount target. Default is not mounting.",
			},

			"unschedulable": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Default:     0,
				Description: "Set whether the added node participates in scheduling. The default value is 0, which means participating in scheduling; non-0 means not participating in scheduling. After the node initialization is completed, you can execute kubectl uncordon nodename to join the node in scheduling.",
			},

			"worker_config": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				MinItems:    1,
				Description: "Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"availability_zone": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Indicates which availability zone will be used.",
						},
						"bandwidth_package_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.",
						},
						"cam_role_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "CAM role name authorized to access.",
						},
						"count": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Default:     1,
							Description: "Number of cvm.",
						},
						"data_disk": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							MaxItems:    11,
							Description: "Configurations of data disk.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_format_and_mount": {
										Type:        schema.TypeBool,
										Optional:    true,
										ForceNew:    true,
										Default:     false,
										Description: "Indicate whether to auto format and mount or not. Default is `false`.",
									},
									"disk_partition": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The name of the device or partition to mount.",
									},
									"disk_size": {
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Default:     0,
										Description: "Volume of disk in GB. Default is `0`.",
									},
									"disk_type": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Default:     "CLOUD_PREMIUM",
										Description: "Types of disk, available values: `CLOUD_PREMIUM` and `CLOUD_SSD` and `CLOUD_HSSD` and `CLOUD_TSSD`.",
									},
									"encrypt": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether to encrypt data disk, default `false`.",
									},
									"file_system": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "File system, e.g. `ext3/ext4/xfs`.",
									},
									"kms_key_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "ID of the custom CMK in the format of UUID or `kms-abcd1234`. This parameter is used to encrypt cloud disks.",
									},
									"mount_target": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Mount target.",
									},
									"snapshot_id": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Data disk snapshot ID.",
									},
								},
							},
						},
						"desired_pod_num": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Default:     0,
							Description: "Indicate to set desired pod number in node. valid when enable_customized_pod_cidr=true, and it override `[globe_]desired_pod_num` for current node. Either all the fields `desired_pod_num` or none.",
						},
						"disaster_recover_group_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Description: "Disaster recover groups to which a CVM instance belongs. Only support maximum 1.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"enhanced_monitor_service": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Default:     true,
							Description: "To specify whether to enable cloud monitor service. Default is TRUE.",
						},
						"enhanced_security_service": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Default:     true,
							Description: "To specify whether to enable cloud security service. Default is TRUE.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).",
						},
						"hpc_cluster_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Id of cvm hpc cluster.",
						},
						"img_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The valid image id, format of img-xxx.",
						},
						"instance_charge_type": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Default:     "POSTPAID_BY_HOUR",
							Description: "The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. The default is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`, `PREPAID` instance will not terminated after cluster deleted, and may not allow to delete before expired.",
						},
						"instance_charge_type_prepaid_period": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Default:     1,
							Description: "The tenancy (time unit is month) of the prepaid instance. NOTE: it only works when instance_charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.",
						},
						"instance_charge_type_prepaid_renew_flag": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to `PREPAID`.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Default:     "sub machine of tke",
							Description: "Name of the CVMs.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Specified types of CVM instance.",
						},
						"internet_charge_type": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Default:     "TRAFFIC_POSTPAID_BY_HOUR",
							Description: "Charge types for network traffic. Available values include `TRAFFIC_POSTPAID_BY_HOUR`.",
						},
						"internet_max_bandwidth_out": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "Max bandwidth of Internet access in Mbps. Default is 0.",
						},
						"key_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Description: "ID list of keys, should be set if `password` not set.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Sensitive:   true,
							Description: "Password to access, should be set if `key_ids` not set.",
						},
						"public_ip_assigned": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: "Specify whether to assign an Internet IP address.",
						},
						"security_group_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: "Security groups to which a CVM instance belongs.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Private network ID.",
						},
						"system_disk_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Default:     50,
							Description: "Volume of system disk in GB. Default is `50`.",
						},
						"system_disk_type": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Default:     "CLOUD_PREMIUM",
							Description: "System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_SSD`: SSD, `CLOUD_PREMIUM`: Premium Cloud Storage. NOTE: `CLOUD_BASIC`, `LOCAL_BASIC` and `LOCAL_SSD` are deprecated.",
						},
						"user_data": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "User data provided to instances, needs to be encoded in base64, and the maximum supported data size is 16KB.",
						},
					},
				},
			},

			"pre_start_user_script": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Base64-encoded user script, executed before initializing the node, currently only effective for adding existing nodes.",
			},

			"user_script": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Base64 encoded user script, this script will be executed after the k8s component is run. The user needs to ensure that the script is reentrant and retry logic. The script and its generated log files can be viewed in the /data/ccs_userscript/ path of the node, if required. The node needs to be initialized before it can be added to the schedule. It can be used with the unschedulable parameter. After the final initialization of userScript is completed, add the kubectl uncordon nodename --kubeconfig=/root/.kube/config command to add the node to the schedule.",
			},

			"worker_instances_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"failed_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Information of the cvm when it is failed.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the cvm.",
						},
						"instance_role": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Role of the cvm.",
						},
						"instance_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "State of the cvm.",
						},
						"lan_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "LAN IP of the cvm.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudKubernetesScaleWorkerCreate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_kubernetes_scale_worker.create")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	if err := resourceTencentCloudKubernetesScaleWorkerCreateOnStart(ctx); err != nil {
		return err
	}

	_ = ctx
	return resourceTencentCloudKubernetesScaleWorkerRead(d, meta)
}

func resourceTencentCloudKubernetesScaleWorkerRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_kubernetes_scale_worker.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	service := TkeService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	idSplit := strings.Split(d.Id(), tccommon.FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	clusterId := idSplit[0]
	instanceIdSet := idSplit[1]

	_ = d.Set("cluster_id", clusterId)

	respData, err := service.DescribeKubernetesScaleWorkerById(ctx, clusterId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `kubernetes_scale_worker` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	respData1, err := service.DescribeKubernetesScaleWorkerById1(ctx, clusterId)
	if err != nil {
		return err
	}

	if respData1 == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `kubernetes_scale_worker` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	respData2, err := service.DescribeKubernetesScaleWorkerById2(ctx)
	if err != nil {
		return err
	}

	if respData2 == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `kubernetes_scale_worker` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	_ = instanceIdSet
	return nil
}

func resourceTencentCloudKubernetesScaleWorkerDelete(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_kubernetes_scale_worker.delete")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)
	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	idSplit := strings.Split(d.Id(), tccommon.FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	clusterId := idSplit[0]
	instanceIdSet := idSplit[1]

	var (
		request  = tke.NewDescribeClustersRequest()
		response = tke.NewDescribeClustersResponse()
	)

	request.ClusterIds = []*string{helper.String(clusterId)}

	err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
		result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseTkeClient().DescribeClustersWithContext(ctx, request)
		if e != nil {
			return tccommon.RetryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		if err := resourceTencentCloudKubernetesScaleWorkerDeletePostRequest0(ctx, request, result); err != nil {
			return err
		}

		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete kubernetes scale worker failed, reason:%+v", logId, err)
		return err
	}

	_ = response
	if err := resourceTencentCloudKubernetesScaleWorkerDeleteOnExit(ctx); err != nil {
		return err
	}

	_ = instanceIdSet
	return nil
}
