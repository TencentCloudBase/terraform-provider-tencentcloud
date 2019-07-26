/*
Provide a resource to create a CLB listener.

Example Usage

```hcl
resource "tencentcloud_clb_listener" "clb_listener" {
  clb_id                     = "lb-k2zjp9lv"
  listener_name              = "mylistener"
  port                       = "80"
  protocol                   = "HTTP"
  health_check_switch        = "0"
  health_check_time_out      = "2"
  health_check_interval_time = "5"
  health_check_health_num    = "3"
  health_check_unhealth_num  = "3"
  certificate_ssl_mode       = "MUTUAL"
  certificate_id             = "mycert server ID "
  certificate_ca_id          = "mycert ca ID"
  session_expire_time        = "0"
  scheduler                  = "WRR"
}
```

Import

CLB listener can be imported using the id, e.g.

```
$ terraform import tencentcloud_clb.listener lbl-qckdffns#lb-p7nlgs4t

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	clb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
	"log"
	"strings"
	"time"
)

func resourceTencentCloudClbListener() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudClbListenerCreate,
		Read:   resourceTencentCloudClbListenerRead,
		Update: resourceTencentCloudClbListenerUpdate,
		Delete: resourceTencentCloudClbListenerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"clb_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateStringLengthInRange(1, 60),
				Description:  "ID of the CLB to be queried.",
			},
			"listener_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringLengthInRange(1, 60),
				Description:  "Name of the CLB listener to be queried, and available values can only be Chinese characters, English letters, numbers, underscore and hyphen '-'",
			},
			"port": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateIntegerInRange(1, 65535),
				Description:  "Port of the CLB listener.",
			},
			"protocol": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue(CLB_LISTENER_PROTOCOL),
				Description:  "Type of protocol within the listener, and available values include 'TCP', 'UDP', 'HTTP', 'HTTPS' and 'TCP_SSL' ('TCP_SSL' is in the internal test, please apply if you need to use).",
			},
			"health_check_switch": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateIntegerInRange(0, 1),
				Computed:     true,
				Description:  "Indicates whether health check is enabled.",
			},
			"health_check_time_out": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateIntegerInRange(2, 60),
				Description:  "Response timeout of health check. The value range is 2-60 sec, and the default is 2 sec. Response timeout needs to be less than check interval.",
			},
			"health_check_interval_time": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateIntegerInRange(5, 300),
				Description:  "Interval time of health check. The value range is 5-300 sec, and the default is 5 sec.",
			},
			"health_check_health_num": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateIntegerInRange(2, 10),
				Description:  "Health threshold of health check, and the default is 3. If a success result is returned for the health check for 3 consecutive times, the backend CVM is identified as healthy. The value range is 2-10.",
			},
			"health_check_unhealth_num": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateIntegerInRange(2, 10),
				Description:  "Unhealth threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, the CVM is identified as unhealthy. The value range is 2-10.",
			},
			"certificate_ssl_mode": {
				Type:     schema.TypeString,
				Optional: true,

				ValidateFunc: validateAllowedStringValue(CERT_SSL_MODE),
				Description:  "Type of SSL Mode, and available values inclue 'UNIDRECTIONAL', 'MUTUAL'.",
			},
			"certificate_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the server certificate. If not specified, the content, key, and name of the server certificate must be set. NOTES: only supported by listners of protocol 'HTTPS'.",
			},
			"certificate_ca_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the client certificate. If not specified, the content, key, name of client certificate must be set when SSLMode is 'mutual'. NOTES: only supported by listners of protocol 'HTTPS'.",
			},
			"session_expire_time": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateIntegerInRange(30, 300),
				Description:  "Time of session persistence within the CLB listener.",
			},
			"scheduler": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue(CLB_LISTENER_SCHEDULER),
				Description:  "Scheduling method of the CLB listener, and available values include 'WRR' and 'LEAST_CONN'. The defaule is 'WRR'.",
			},
			"sni_switch": {
				Type:         schema.TypeInt,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validateIntegerInRange(0, 1),
				Description:  "Indicates whether SNI is enabled, and only supported with protocol 'HTTPS'.",
			},
		},
	}
}

func resourceTencentCloudClbListenerCreate(d *schema.ResourceData, meta interface{}) error {
	logId := GetLogId(nil)
	defer LogElapsed(logId + "resource.tencentcloud_clb_listener.create")()
	request := clb.NewCreateListenerRequest()
	clbId := d.Get("clb_id").(string)
	request.LoadBalancerId = stringToPointer(clbId)
	listenerName := d.Get("listener_name").(string)
	request.ListenerNames = []*string{&listenerName}

	port := int64(d.Get("port").(int))
	ports := []*int64{&port}
	request.Ports = ports
	protocol := d.Get("protocol").(string)
	if protocol == "TCP_SSL" {
		return fmt.Errorf("TCP_SSL protocol type needs manual application")
	} else {
		request.Protocol = stringToPointer(protocol)
	}

	ctx := context.WithValue(context.TODO(), "logId", logId)

	healthSetFlag, healthCheck, healthErr := checkHealthCheckPara(ctx, d, protocol)
	if healthErr != nil {
		return healthErr
	}
	if healthSetFlag == true {
		request.HealthCheck = healthCheck
	}

	certificateSetFlag, certificateInput, certErr := checkCertificateInputPara(ctx, d)
	if certErr != nil {
		return certErr
	}
	if certificateSetFlag == true {
		request.Certificate = certificateInput
	}

	if v, ok := d.GetOk("session_expire_time"); ok {
		if !(protocol == CLB_LISTENER_PROTOCOL_TCP || protocol == CLB_LISTENER_PROTOCOL_UDP) {
			return fmt.Errorf("session_expire_time can only be set with protocol TCP/UDP")
		}
		vv := int64(v.(int))
		request.SessionExpireTime = &vv
	}

	if v, ok := d.GetOk("scheduler"); ok {
		if !(protocol == CLB_LISTENER_PROTOCOL_TCP || protocol == CLB_LISTENER_PROTOCOL_UDP) {
			return fmt.Errorf("Scheduler can only be set with protocol TCP/UDP")
		}

		request.Scheduler = stringToPointer(v.(string))
	}
	if v, ok := d.GetOk("sni_switch"); ok {
		if protocol != CLB_LISTENER_PROTOCOL_HTTPS {
			return fmt.Errorf("sni_switch can only be set with protocol HTTPS ")
		} else {
			vv := int64(v.(int))
			request.SniSwitch = &vv
		}
	}
	response, err := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().CreateListener(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	} else {
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		request_id := *response.Response.RequestId

		retryErr := retrySet(request_id, meta.(*TencentCloudClient).apiV3Conn.UseClbClient())
		if retryErr != nil {
			return retryErr
		}
	}

	if len(response.Response.ListenerIds) < 1 {
		return fmt.Errorf("load balancer id is nil")
	}
	listenerId := *response.Response.ListenerIds[0]
	d.SetId(listenerId + "#" + clbId)
	time.Sleep(6 * time.Second)
	return resourceTencentCloudClbListenerRead(d, meta)
}

func resourceTencentCloudClbListenerRead(d *schema.ResourceData, meta interface{}) error {
	logId := GetLogId(nil)
	defer LogElapsed(logId + "resource.tencentcloud_clb_listener.read")()
	ctx := context.WithValue(context.TODO(), "logId", logId)

	items := strings.Split(d.Id(), "#")
	if len(items) != 2 {
		return fmt.Errorf("id of resource.tencentcloud_clb_listener is wrong")
	}

	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	instance, err := clbService.DescribeListenerById(ctx, d.Id())
	if err != nil {
		return err
	}

	d.Set("clb_id", items[1])
	d.Set("listener_name", instance.ListenerName)
	d.Set("port", instance.Port)
	d.Set("protocol", instance.Protocol)
	d.Set("session_expire_time", instance.SessionExpireTime)
	d.Set("scheduler", instance.Scheduler)
	d.Set("sni_switch", instance.SniSwitch)
	//health check
	d.Set("health_check_switch", instance.HealthCheck.HealthSwitch)
	d.Set("health_check_interval_time", instance.HealthCheck.IntervalTime)
	d.Set("health_check_time_out", instance.HealthCheck.TimeOut)
	d.Set("health_check_interval_time", instance.HealthCheck.IntervalTime)
	d.Set("health_check_health_num ", instance.HealthCheck.HealthNum)
	d.Set("health_check_unhealth_num", instance.HealthCheck.UnHealthNum)

	if instance.Certificate != nil {
		d.Set("certificate_ssl_mode", instance.Certificate.SSLMode)
		d.Set("certificate_id", instance.Certificate.CertId)
		d.Set("certificate_ca_id", instance.Certificate.CertCaId)

	}

	return nil
}

func resourceTencentCloudClbListenerUpdate(d *schema.ResourceData, meta interface{}) error {
	logId := GetLogId(nil)
	defer LogElapsed(logId + "resource.tencentcloud_clb_listener.update")()

	items := strings.Split(d.Id(), "#")
	if len(items) != 2 {
		return fmt.Errorf("id of resource.tencentcloud_clb_listener is wrong")
	}

	listenerId := items[0]
	clbId := items[1]
	changed := false
	scheduler := ""
	listenerName := ""
	sessionExpireTime := 0

	protocol := d.Get("protocol").(string)

	if d.HasChange("listener_name") {
		listenerName = d.Get("listener_name").(string)
	}

	if d.HasChange("scheduler") {
		changed = true
		scheduler = d.Get("scheduler").(string)
	}
	if d.HasChange("session_expire_time") {
		changed = true
		sessionExpireTime = d.Get("session_expire_time").(int)
	}

	ctx := context.WithValue(context.TODO(), "logId", logId)
	healthSetFlag, healthCheck, healthErr := checkHealthCheckPara(ctx, d, protocol)
	if healthErr != nil {
		return healthErr
	}
	if healthSetFlag == true {
		changed = true
	}

	certificateSetFlag, certificateInput, certErr := checkCertificateInputPara(ctx, d)
	if certErr != nil {
		return certErr
	}
	if certificateSetFlag == true {
		changed = true
	}

	if changed {
		request := clb.NewModifyListenerRequest()
		request.ListenerId = stringToPointer(listenerId)
		request.LoadBalancerId = stringToPointer(clbId)

		if d.HasChange("listener_name") {
			request.ListenerName = stringToPointer(listenerName)
		}

		if d.HasChange("scheduler") {
			request.Scheduler = stringToPointer(scheduler)
		}
		if d.HasChange("session_expire_time") {
			sessionExpireTime64 := int64(sessionExpireTime)
			request.SessionExpireTime = &sessionExpireTime64
		}
		if healthSetFlag == true {
			request.HealthCheck = healthCheck
		}
		if certificateSetFlag == true {
			request.Certificate = certificateInput
		}

		response, err := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().ModifyListener(request)

		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			return err
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
			request_id := *response.Response.RequestId
			retryErr := retrySet(request_id, meta.(*TencentCloudClient).apiV3Conn.UseClbClient())
			if retryErr != nil {
				return retryErr
			}
		}

	}

	return nil
}

func resourceTencentCloudClbListenerDelete(d *schema.ResourceData, meta interface{}) error {
	logId := GetLogId(nil)
	defer LogElapsed(logId + "resource.tencentcloud_clb_listener.delete")()
	ctx := context.WithValue(context.TODO(), "logId", logId)

	items := strings.Split(d.Id(), "#")
	if len(items) != 2 {
		return fmt.Errorf("id of resource.tencentcloud_clb_listener is wrong")
	}
	listenerId := items[0]
	clbId := items[1]
	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := clbService.DeleteListenerById(ctx, clbId, listenerId)
	if err != nil {
		log.Printf("[CRITAL]%s reason[%s]\n", logId, err.Error())
		return err
	}

	return nil
}
