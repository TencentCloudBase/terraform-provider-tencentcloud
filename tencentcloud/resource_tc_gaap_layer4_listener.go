package tencentcloud

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	gaap "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gaap/v20180529"
)

func resourceTencentCloudGaapLayer4Listener() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudGaapLayer4ListenerCreate,
		Read:   resourceTencentCloudGaapLayer4ListenerRead,
		Update: resourceTencentCloudGaapLayer4ListenerUpdate,
		Delete: resourceTencentCloudGaapLayer4ListenerDelete,
		Schema: map[string]*schema.Schema{
			"protocol": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"TCP", "UDP"}),
				ForceNew:     true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringLengthInRange(1, 30),
			},
			"port": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validatePort,
				ForceNew:     true,
			},
			"scheduler": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "rr",
				ValidateFunc: validateAllowedStringValue([]string{"rr", "wrr", "lc"}),
			},
			"realserver_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"IP", "DOMAIN"}),
				ForceNew:     true,
			},
			"proxy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"health_check": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"delay_loop": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateIntegerInRange(5, 300),
			},
			"connect_timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateIntegerInRange(2, 60),
			},
			"realserver_bind_set": {
				Type:     schema.TypeSet,
				Optional: true,
				Set: func(v interface{}) int {
					m := v.(map[string]interface{})
					sb := new(strings.Builder)
					sb.WriteString(m["id"].(string))
					sb.WriteString(m["ip"].(string))
					sb.WriteString(fmt.Sprintf("%d", m["port"].(int)))
					sb.WriteString(fmt.Sprintf("%d", m["weight"].(int)))
					return hashcode.String(sb.String())
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ip": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validatePort,
						},
						"weight": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      1,
							ValidateFunc: validateIntegerMin(1),
						},
					},
				},
			},

			// computed
			"status": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceTencentCloudGaapLayer4ListenerCreate(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.tencentcloud_gaap_layer4_listener.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	protocol := d.Get("protocol").(string)
	name := d.Get("name").(string)
	port := d.Get("port").(int)
	scheduler := d.Get("scheduler").(string)
	realserverType := d.Get("realserver_type").(string)

	if protocol == "TCP" && realserverType == "DOMAIN" && scheduler == "wrr" {
		return errors.New("TCP listener DOMAIN realserver type doesn't support wrr scheduler")
	}

	proxyId := d.Get("proxy_id").(string)
	healthCheck := d.Get("health_check").(bool)

	var (
		delayLoop      *int
		connectTimeout *int
		realservers    []gaapRealserverBind
	)
	if raw, ok := d.GetOk("delay_loop"); ok {
		delayLoop = common.IntPtr(raw.(int))
	}
	if raw, ok := d.GetOk("connect_timeout"); ok {
		connectTimeout = common.IntPtr(raw.(int))
	}
	if (connectTimeout == nil && delayLoop != nil) || (connectTimeout != nil && delayLoop == nil) {
		return errors.New("connect_timeout and delay_loop must be set together")
	}

	if connectTimeout != nil {
		if *connectTimeout >= *delayLoop {
			return errors.New("connect_timeout must be less than delay_loop")
		}
	}

	if raw, ok := d.GetOk("realserver_bind_set"); ok {
		list := raw.(*schema.Set).List()
		realservers = make([]gaapRealserverBind, 0, len(list))
		for _, v := range list {
			m := v.(map[string]interface{})
			realservers = append(realservers, gaapRealserverBind{
				id:     m["id"].(string),
				ip:     m["ip"].(string),
				port:   m["port"].(int),
				weight: m["weight"].(int),
			})
		}
	}

	for _, rs := range realservers {
		if scheduler == "rr" && rs.weight != 1 {
			return errors.New("when scheduler is rr, realserver weight should be 1 or null")
		}
	}

	var (
		id  string
		err error
	)

	service := GaapService{client: m.(*TencentCloudClient).apiV3Conn}

	switch protocol {
	case "TCP":
		id, err = service.CreateTCPListener(ctx, name, scheduler, realserverType, proxyId, port, healthCheck, delayLoop, connectTimeout)
		if err != nil {
			return err
		}

	case "UDP":
		id, err = service.CreateUDPListener(ctx, name, scheduler, realserverType, proxyId, port)
		if err != nil {
			return err
		}
	}

	if err := service.BindLayer4ListenerRealservers(ctx, id, protocol, proxyId, realservers); err != nil {
		return err
	}

	d.SetId(id)
	return resourceTencentCloudGaapLayer4ListenerRead(d, m)
}

func resourceTencentCloudGaapLayer4ListenerRead(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.tencentcloud_gaap_layer4_listener.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	id := d.Id()
	protocol := d.Get("protocol").(string)
	proxyId := d.Get("proxy_id").(string)
	var (
		name           string
		port           int
		scheduler      string
		realServerType string
		healthCheck    *bool
		delayLoop      *int
		connectTimeout *int
		status         *int
		createTime     *int
		realservers    []map[string]interface{}
	)

	service := GaapService{client: m.(*TencentCloudClient).apiV3Conn}

	switch protocol {
	case "TCP":
		listeners, err := service.DescribeTCPListener(ctx, proxyId, &id, nil, nil)
		if err != nil {
			return err
		}

		if len(listeners) == 0 {
			d.SetId("")
			return nil
		}

		var listener *gaap.TCPListener
		for _, l := range listeners {
			if l.ListenerId == nil {
				return errors.New("listener id is nil")
			}
			if *l.ListenerId == id {
				listener = l
				break
			}
		}
		if listener == nil {
			d.SetId("")
			return nil
		}

		if listener.ListenerName == nil {
			return fmt.Errorf("listener %s name is nil", id)
		}
		name = *listener.ListenerName

		if listener.Port == nil {
			return fmt.Errorf("listener %s port is nil", id)
		}
		port = int(*listener.Port)

		if listener.Scheduler == nil {
			return fmt.Errorf("listener %s scheduler is nil", id)
		}
		scheduler = *listener.Scheduler

		if listener.RealServerType == nil {
			return fmt.Errorf("listener %s realserver type is nil", id)
		}
		realServerType = *listener.RealServerType

		if listener.HealthCheck != nil {
			healthCheck = common.BoolPtr(*listener.HealthCheck == 1)
		}

		if listener.DelayLoop != nil {
			delayLoop = common.IntPtr(int(*listener.DelayLoop))
		}

		if listener.ConnectTimeout != nil {
			connectTimeout = common.IntPtr(int(*listener.ConnectTimeout))
		}

		if len(listener.RealServerSet) > 0 {
			realservers = make([]map[string]interface{}, 0, len(listener.RealServerSet))
			for _, rs := range listener.RealServerSet {
				if rs.RealServerId == nil {
					return errors.New("realserver id is nil")
				}
				if rs.RealServerIP == nil {
					return errors.New("realserver IP is nil")
				}
				if rs.RealServerPort == nil {
					return errors.New("realserver port is nil")
				}
				if rs.RealServerWeight == nil {
					return errors.New("realserver weight is nil")
				}

				realservers = append(realservers, map[string]interface{}{
					"id":     *rs.RealServerId,
					"ip":     *rs.RealServerIP,
					"port":   *rs.RealServerPort,
					"weight": *rs.RealServerWeight,
				})
			}
		}

		if listener.ListenerStatus != nil {
			status = common.IntPtr(int(*listener.ListenerStatus))
		}

		if listener.CreateTime != nil {
			createTime = common.IntPtr(int(*listener.CreateTime))
		}

	case "UDP":
		listeners, err := service.DescribeUDPListener(ctx, proxyId, &id, nil, nil)
		if err != nil {
			return err
		}

		if len(listeners) == 0 {
			d.SetId("")
			return nil
		}

		var listener *gaap.UDPListener
		for _, l := range listeners {
			if l.ListenerId == nil {
				return errors.New("listener id is nil")
			}
			if *l.ListenerId == id {
				listener = l
				break
			}
		}
		if listener == nil {
			d.SetId("")
			return nil
		}

		if listener.ListenerName == nil {
			return fmt.Errorf("listener %s name is nil", id)
		}
		name = *listener.ListenerName

		if listener.Port == nil {
			return fmt.Errorf("listener %s port is nil", id)
		}
		port = int(*listener.Port)

		if listener.Scheduler == nil {
			return fmt.Errorf("listener %s scheduler is nil", id)
		}
		scheduler = *listener.Scheduler

		if listener.RealServerType == nil {
			return fmt.Errorf("listener %s realserver type is nil", id)
		}
		realServerType = *listener.RealServerType

		if len(listener.RealServerSet) > 0 {
			realservers = make([]map[string]interface{}, 0, len(listener.RealServerSet))
			for _, rs := range listener.RealServerSet {
				if rs.RealServerId == nil {
					return errors.New("realserver id is nil")
				}
				if rs.RealServerIP == nil {
					return errors.New("realserver IP is nil")
				}
				if rs.RealServerPort == nil {
					return errors.New("realserver port is nil")
				}
				if rs.RealServerWeight == nil {
					return errors.New("realserver weight is nil")
				}

				realservers = append(realservers, map[string]interface{}{
					"id":     *rs.RealServerId,
					"ip":     *rs.RealServerIP,
					"port":   *rs.RealServerPort,
					"weight": *rs.RealServerWeight,
				})
			}
		}

		if listener.ListenerStatus != nil {
			status = common.IntPtr(int(*listener.ListenerStatus))
		}

		if listener.CreateTime != nil {
			createTime = common.IntPtr(int(*listener.CreateTime))
		}
	}

	d.Set("name", name)
	d.Set("port", port)
	d.Set("scheduler", scheduler)
	d.Set("realserver_type", realServerType)
	if healthCheck != nil {
		d.Set("health_check", healthCheck)
	}
	if delayLoop != nil {
		d.Set("delay_loop", delayLoop)
	}
	if connectTimeout != nil {
		d.Set("connect_timeout", connectTimeout)
	}
	if len(realservers) > 0 {
		d.Set("realserver_bind_set", realservers)
	}
	if status != nil {
		d.Set("status", status)
	}
	if createTime != nil {
		d.Set("create_time", createTime)
	}

	return nil
}

func resourceTencentCloudGaapLayer4ListenerUpdate(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.tencentcloud_gaap_layer4_listener.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	id := d.Id()
	protocol := d.Get("protocol").(string)
	proxyId := d.Get("proxy_id").(string)
	var (
		name           *string
		scheduler      *string
		healthCheck    *bool
		delayLoop      *int
		connectTimeout *int
		realservers    []gaapRealserverBind
		hasChange      []string
	)

	service := GaapService{client: m.(*TencentCloudClient).apiV3Conn}

	d.Partial(true)

	if d.HasChange("name") {
		hasChange = append(hasChange, "name")
		name = stringToPointer(d.Get("name").(string))
	}
	if d.HasChange("scheduler") {
		hasChange = append(hasChange, "scheduler")
		scheduler = stringToPointer(d.Get("scheduler").(string))
	}
	if d.HasChange("health_check") {
		hasChange = append(hasChange, "health_check")
		healthCheck = common.BoolPtr(d.Get("health_check").(bool))
	}
	if d.HasChange("delay_loop") {
		hasChange = append(hasChange, "delay_loop")
		delayLoop = common.IntPtr(d.Get("delay_loop").(int))
	}
	if d.HasChange("connect_timeout") {
		hasChange = append(hasChange, "connect_timeout")
		connectTimeout = common.IntPtr(d.Get("connect_timeout").(int))
	}

	if d.HasChange("realserver_bind_set") {
		list := d.Get("realserver_bind_set").(*schema.Set).List()
		realservers = make([]gaapRealserverBind, 0, len(list))
		for _, v := range list {
			m := v.(map[string]interface{})
			realservers = append(realservers, gaapRealserverBind{
				id:     m["id"].(string),
				ip:     m["ip"].(string),
				port:   m["port"].(int),
				weight: m["weight"].(int),
			})
		}
	}

	if len(hasChange) > 0 {
		switch protocol {
		case "TCP":
			if err := service.ModifyTCPListenerAttribute(ctx, proxyId, id, name, scheduler, healthCheck, delayLoop, connectTimeout); err != nil {
				return err
			}

		case "UDP":
			if err := service.ModifyUDPListenerAttribute(ctx, proxyId, id, name, scheduler); err != nil {
				return err
			}
		}

		for _, attr := range hasChange {
			d.SetPartial(attr)
		}
	}

	if len(realservers) > 0 {
		if err := service.BindLayer4ListenerRealservers(ctx, id, protocol, proxyId, realservers); err != nil {
			return err
		}
		d.SetPartial("realserver_bind_set")
	}

	d.Partial(false)

	return resourceTencentCloudGaapLayer4ListenerRead(d, m)
}

func resourceTencentCloudGaapLayer4ListenerDelete(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.tencentcloud_gaap_layer4_listener.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	id := d.Id()
	protocol := d.Get("protocol").(string)
	proxyId := d.Get("proxy_id").(string)

	service := GaapService{client: m.(*TencentCloudClient).apiV3Conn}

	return service.DeleteLayer4Listener(ctx, id, proxyId, protocol)
}
