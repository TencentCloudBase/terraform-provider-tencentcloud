/*
Provide a resource to create a CLB listener rule.

Example Usage

```hcl
resource "tencentcloud_clb_listener_forward_rule" "rule" {
  listener_id                = "lbl-hh141sn9"
  clb_id                     = "lb-k2zjp9lv"
  domain                     = "foo.net"
  url                        = "/bar"
  health_check_switch        = "0"
  health_check_interval_time = "5"
  health_check_health_num    = "3"
  health_check_unhealth_num  = "3"
  health_check_http_code     = "http_1xx"
  health_check_http_path     = "Default Path"
  health_check_http_domain   = "Default Domain"
  health_check_http_method   = "GET"
  certificate_server_id      = "my server certificate ID "
  certificate_ca_id          = "my client certificate ID"
  session_expire_time        = "0"
  schedule                   = "WRR"
}
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	clb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
)

func resourceTencentCloudClbListenerRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudClbListenerRuleCreate,
		Read:   resourceTencentCloudClbListenerRuleRead,
		Update: resourceTencentCloudClbListenerRuleUpdate,
		Delete: resourceTencentCloudClbListenerRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"listener_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of CLB listener.",
			},
			"clb_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of CLB instance. ",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Domain name of the forwarding rules",
			},
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Url of the forwarding rules",
			},
			"health_check_switch": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateIntegerInRange(0, 1),
				Computed:     true,
				Description:  "Indicates whether health check is enabled.",
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
				Description:  "Health threshold of health check, and the default is 3. If a success result is returned for the health check for 3 consecutive times, the backend CVM is identified as healthy. The value range is 2-10",
			},
			"health_check_unhealth_num": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateIntegerInRange(2, 10),
				Description:  "Health threshold of health check, and the default is 3. If a success result is returned for the health check for 3 consecutive times, the backend CVM is identified as healthy. The value range is 2-10",
			},
			"health_check_http_code": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateIntegerInRange(1, 31),
				Description:  "Path of health check (applicable only to HTTP/HTTPS check methods).",
			},
			"health_check_http_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Path of health check (applicable only to HTTP/HTTPS check methods). ",
			},
			"health_check_http_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Domain name of health check (applicable only to HTTP/HTTPS check methods)",
			},
			"health_check_http_method": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue(CLB_HTTP_METHOD),
				Description:  "Methods of health check (applicable only to HTTP/HTTPS check methods). Available values include 'HEAD' and 'GET'.",
			},
			"certificate_ssl_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue(CERT_SSL_MODE),
				Description:  "Type of SSL Mode. Available values are 'UNIDRECTIONAL', 'MUTUAL' ",
			},
			"certificate_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Id of the server certificate.If not set, the content, key, name of server certificate must be set, only supported by listners of protocol 'HTTPS'. ",
			},
			"certificate_ca_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Id of the client certificate.If not set, the content, key, name of client certificate must be set when SSLMode is 'mutual', only supported by listners of protocol 'HTTPS'. ",
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
		},
	}
}

func resourceTencentCloudClbListenerRuleCreate(d *schema.ResourceData, meta interface{}) error {
	logId := GetLogId(nil)
	defer LogElapsed(logId + "resource.tencentcloud_clb_listener_rule.create")()

	ctx := context.WithValue(context.TODO(), "logId", logId)

	items := strings.Split(d.Get("listener_id").(string), "#")
	if len(items) != 2 {
		return fmt.Errorf("id of resource.tencentcloud_clb_listener is wrong")
	}

	listenerId := items[0]
	clbId := items[1]
	//get listener protocol
	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	instance, err := clbService.DescribeListenerById(ctx, listenerId+"#"+clbId)
	if err != nil {
		return err
	}

	protocol := *(instance.Protocol)

	request := clb.NewCreateRuleRequest()
	request.LoadBalancerId = stringToPointer(clbId)
	request.ListenerId = stringToPointer(listenerId)

	//rule set
	var rule clb.RuleInput

	domain := d.Get("domain").(string)
	rule.Domain = stringToPointer(domain)
	url := d.Get("url").(string)
	rule.Url = stringToPointer(url)

	if v, ok := d.GetOk("session_expire_time"); ok {
		sessionExpireTime := int64(v.(int))
		rule.SessionExpireTime = &sessionExpireTime
	}
	if v, ok := d.GetOk("scheduler"); ok {
		rule.Scheduler = stringToPointer(v.(string))
	}

	healthSetFlag, healthCheck, healthErr := checkHealthCheckPara(ctx, d, protocol, HEALTH_APPLY_TYPE_RULE)
	if healthErr != nil {
		return healthErr
	}
	if healthSetFlag == true {
		rule.HealthCheck = healthCheck
	}

	certificateSetFlag, certificateInput, certErr := checkCertificateInputPara(ctx, d)
	if certErr != nil {
		return certErr
	}
	if certificateSetFlag == true {
		rule.Certificate = certificateInput
	}

	request.Rules = []*clb.RuleInput{&rule}
	requestId := ""
	response, err := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().CreateRule(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	} else {
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		requestId = *response.Response.RequestId

		retryErr := retrySet(requestId, meta.(*TencentCloudClient).apiV3Conn.UseClbClient())
		if retryErr != nil {
			return retryErr
		}
	}

	ruleInstance, ruleErr := clbService.DescribeRuleByPara(ctx, clbId, listenerId, domain, url)
	if ruleErr != nil {
		return ruleErr
	}
	d.SetId(*ruleInstance.LocationId + "#" + listenerId + "#" + clbId)
	return resourceTencentCloudClbListenerRuleRead(d, meta)
}

func resourceTencentCloudClbListenerRuleRead(d *schema.ResourceData, meta interface{}) error {
	logId := GetLogId(nil)
	defer LogElapsed(logId + "resource.tencentcloud_clb_listener_rule.read")()
	ctx := context.WithValue(context.TODO(), "logId", logId)

	ruleId := d.Id()
	items := strings.Split(ruleId, "#")
	if len(items) != 3 {
		return fmt.Errorf("id of resource.tencentcloud_clb_listener is wrong")
	}
	locationId := items[0]
	listenerId := items[1]
	clbId := items[2]
	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	//this function is not supported by api, need to be travelled
	filter := map[string]string{"location_id": locationId, "listener_id": listenerId, "clb_id": clbId}
	instances, err := clbService.DescribeRulesByFilter(ctx, filter)
	if err != nil {
		return err
	}
	if len(instances) == 0 {
		fmt.Errorf("rule not found!")
	}
	instance := instances[0]
	d.Set("clb_id", clbId)
	d.Set("listener_id", listenerId+"#"+clbId)
	d.Set("domain", instance.Domain)
	d.Set("location_id", instance.LocationId)
	d.Set("url", instance.Url)
	d.Set("scheduler", instance.Scheduler)
	d.Set("session_expire_time", instance.SessionExpireTime)

	//health check
	if instance.HealthCheck != nil {
		d.Set("health_check_switch", instance.HealthCheck.HealthSwitch)
		d.Set("health_check_interval_time", instance.HealthCheck.IntervalTime)
		//d.Set("health_check_time_out", instance.HealthCheck.TimeOut)
		d.Set("health_check_interval_time", instance.HealthCheck.IntervalTime)
		d.Set("health_check_health_num", instance.HealthCheck.HealthNum)
		d.Set("health_check_unhealth_num", instance.HealthCheck.UnHealthNum)
		d.Set("health_check_http_method", stringToPointer(strings.ToUpper(*instance.HealthCheck.HttpCheckMethod)))
		d.Set("health_check_http_domain", instance.HealthCheck.HttpCheckDomain)
		d.Set("health_check_http_path", instance.HealthCheck.HttpCheckPath)
		d.Set("health_check_http_code", instance.HealthCheck.HttpCode)
	}

	if instance.Certificate != nil {
		d.Set("certificate_ssl_mode", instance.Certificate.SSLMode)
		d.Set("certificate_id", instance.Certificate.CertId)
		d.Set("certificate_ca_id", instance.Certificate.CertCaId)

	}

	return nil
}

func resourceTencentCloudClbListenerRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	logId := GetLogId(nil)
	defer LogElapsed(logId + "resource.tencentcloud_clb_listener_rule.update")()
	ctx := context.WithValue(context.TODO(), "logId", logId)

	items := strings.Split(d.Get("listener_id").(string), "#")
	if len(items) != 2 {
		return fmt.Errorf("id of resource.tencentcloud_clb_listener is wrong")
	}

	listenerId := items[0]
	clbId := items[1]
	//get listener protocol
	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	instance, err := clbService.DescribeListenerById(ctx, listenerId+"#"+clbId)
	if err != nil {
		return err
	}

	protocol := *(instance.Protocol)

	locationId := d.Get("location_id").(string)
	changed := false
	url := ""
	scheduler := ""
	sessionExpireTime := 0

	if d.HasChange("url") {
		changed = true
		url = d.Get("url").(string)
	}

	if d.HasChange("scheduler") {
		changed = true
		scheduler = d.Get("scheduler").(string)
	}
	if d.HasChange("session_expire_time") {
		changed = true
		sessionExpireTime = d.Get("session_expire_time").(int)
	}

	healthSetFlag, healthCheck, healthErr := checkHealthCheckPara(ctx, d, protocol, HEALTH_APPLY_TYPE_RULE)
	if healthErr != nil {
		return healthErr
	}

	if healthSetFlag == true {
		changed = true
	}

	if changed {
		request := clb.NewModifyRuleRequest()
		request.ListenerId = stringToPointer(listenerId)
		request.LoadBalancerId = stringToPointer(clbId)
		request.LocationId = stringToPointer(locationId)
		if d.HasChange("url") {
			request.Url = stringToPointer(url)
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

		response, err := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().ModifyRule(request)

		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			return err
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
			requestId := *response.Response.RequestId
			retryErr := retrySet(requestId, meta.(*TencentCloudClient).apiV3Conn.UseClbClient())
			if retryErr != nil {
				return retryErr
			}
		}

	}

	return nil
}
func resourceTencentCloudClbListenerRuleDelete(d *schema.ResourceData, meta interface{}) error {
	logId := GetLogId(nil)
	defer LogElapsed(logId + "resource.tencentcloud_clb_listener_rule.delete")()
	ctx := context.WithValue(context.TODO(), "logId", logId)
	ruleId := d.Id()
	items := strings.Split(ruleId, "#")
	if len(items) != 3 {
		return fmt.Errorf("id of resource.tencentcloud_clb_listener is wrong")
	}
	locationId := items[0]
	listenerId := items[1]
	clbId := items[2]
	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := clbService.DeleteRuleById(ctx, clbId, listenerId, locationId)
	if err != nil {
		log.Printf("[CRITAL]%s reason[%s]\n", logId, err.Error())
		return err
	}
	return nil
}
