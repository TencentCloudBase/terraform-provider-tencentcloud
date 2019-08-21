package tencentcloud

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

func dataSourceTencentCloudGaapProxies() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudGaapProxiesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:          schema.TypeSet,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				ConflictsWith: []string{"project_id", "access_region", "realserver_region"},
			},
			"project_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"ids"},
			},
			"access_region": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"ids"},
			},
			"realserver_region": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"ids"},
			},
			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			// computed
			"proxies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"policy_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bandwidth": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"concurrent": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"access_region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"realserver_region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"create_time": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"scalable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"support_protocols": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"forward_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": {
							Type:     schema.TypeMap,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudGaapProxiesRead(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("data_source.tencentcloud_gaap_proxies.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	var (
		ids              []string
		projectId        *int
		accessRegion     *string
		realserverRegion *string
	)

	if raw, ok := d.GetOk("ids"); ok {
		set := raw.(*schema.Set).List()
		ids = make([]string, 0, len(set))
		for _, id := range set {
			ids = append(ids, id.(string))
		}
	}

	if raw, ok := d.GetOk("project_id"); ok {
		projectId = common.IntPtr(raw.(int))
	}

	if raw, ok := d.GetOk("access_region"); ok {
		accessRegion = stringToPointer(raw.(string))
	}

	if raw, ok := d.GetOk("realserver_region"); ok {
		realserverRegion = stringToPointer(raw.(string))
	}

	tags := getTags(d, "tags")

	service := GaapService{client: m.(*TencentCloudClient).apiV3Conn}

	proxies, err := service.DescribeProxies(ctx, ids, projectId, accessRegion, realserverRegion, tags)
	if err != nil {
		return err
	}

	proxyIds := make([]string, 0, len(proxies))
	respProxies := make([]map[string]interface{}, 0, len(proxies))
	for _, proxy := range proxies {
		if proxy.ProxyId == nil {
			return errors.New("proxy id is nil")
		}
		if proxy.ProxyName == nil {
			return errors.New("proxy name is nil")
		}
		if proxy.Domain == nil {
			return errors.New("proxy domain is nil")
		}
		if proxy.IP == nil {
			return errors.New("proxy ip is nil")
		}
		if proxy.Bandwidth == nil {
			return errors.New("proxy bandwidth is nil")
		}
		if proxy.Concurrent == nil {
			return errors.New("proxy concurrent is nil")
		}
		if proxy.AccessRegion == nil {
			return errors.New("proxy access region is nil")
		}
		if proxy.RealServerRegion == nil {
			return errors.New("proxy realserver region is nil")
		}
		if proxy.ProjectId == nil {
			return errors.New("proxy project id is nil")
		}
		if proxy.CreateTime == nil {
			return errors.New("proxy create time is nil")
		}
		if proxy.Status == nil {
			return errors.New("proxy status is nil")
		}
		if proxy.Scalarable == nil {
			return errors.New("proxy scalable is nil")
		}
		if proxy.SupportProtocols == nil {
			return errors.New("proxy support protocols is nil")
		}
		if proxy.ForwardIP == nil {
			return errors.New("proxy forward ip is nil")
		}
		if proxy.Version == nil {
			return errors.New("proxy version is nil")
		}

		proxyIds = append(proxyIds, *proxy.ProxyId)

		m := map[string]interface{}{
			"id":                *proxy.ProxyId,
			"name":              *proxy.ProxyName,
			"domain":            *proxy.Domain,
			"ip":                *proxy.IP,
			"bandwidth":         *proxy.Bandwidth,
			"concurrent":        *proxy.Concurrent,
			"access_region":     *proxy.AccessRegion,
			"realserver_region": *proxy.RealServerRegion,
			"project_id":        *proxy.ProjectId,
			"create_time":       *proxy.CreateTime,
			"status":            *proxy.Status,
			"scalable":          *proxy.Scalarable == 1,
			"forward_ip":        *proxy.ForwardIP,
			"version":           *proxy.Version,
		}

		if proxy.PolicyId != nil {
			m["policy_id"] = *proxy.PolicyId
		}

		supportProtocols := make([]string, 0, len(proxy.SupportProtocols))
		for _, v := range proxy.SupportProtocols {
			supportProtocols = append(supportProtocols, *v)
		}
		m["support_protocols"] = supportProtocols

		tags := make(map[string]string, len(proxy.TagSet))
		for _, t := range proxy.TagSet {
			if t.TagKey == nil {
				return errors.New("tag key is nil")
			}
			if t.TagValue == nil {
				return errors.New("tag value is nil")
			}
			tags[*t.TagKey] = *t.TagValue
		}
		m["tags"] = tags

		respProxies = append(respProxies, m)
	}

	d.Set("proxies", respProxies)

	d.SetId(dataResourceIdsHash(proxyIds))

	return nil
}
