/*
Use this data source to query dayu L4 rules
Example Usage
```hcl
data "tencentcloud_dayu_l4_rules" "name_test" {
  resource_type = tencentcloud_dayu_l4_rule.test_rule.resource_type
  resource_id      = tencentcloud_dayu_l4_rule.test_rule.resource_id
  name		= tencentcloud_dayu_l4_rule.test_rule.name
}
data "tencentcloud_dayu_l4_rules" "id_test" {
  resource_type = tencentcloud_dayu_l4_rule.test_rule.resource_type
  resource_id      = tencentcloud_dayu_l4_rule.test_rule.resource_id
  rule_id		= tencentcloud_dayu_l4_rule.test_rule.rule_id
}
```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func dataSourceTencentCloudDayuL4Rules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudDayuL4RulesRead,
		Schema: map[string]*schema.Schema{
			"resource_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue(DAYU_RESOURCE_TYPE),
				Description:  "Type of the resource that the L4 rule works for, valid values are `bgpip`, `bgp`, `bgp-multip`, `net`.",
			},
			"resource_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of the resource that the L4 rule works for.",
			},
			"rule_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Id of the L4 rule to be queried.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the L4 rule to be queried.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of L4 rules. Each element contains the following attributes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Source type, 1 for source of host, 2 for source of ip.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the rule.",
						},
						"source_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The source port of the L4 rule.",
						},
						"dest_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The destination port of the L4 rule.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protocol of the rule.",
						},
						"source_list": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Source ip or domain.",
									},
									"weight": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Weight of the source.",
									},
								},
							},
							Description: "Source list of the rule.",
						},
						"health_check_switch": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether health check is enabled.",
						},
						"health_check_interval": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Interval time of health check.",
						},
						"health_check_health_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Health threshold of health check.",
						},
						"health_check_unhealth_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Unhealthy threshold of health check.",
						},
						"health_check_timeout": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "HTTP Status Code. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.",
						},
						"session_switch": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicate that the session will keep or not.",
						},
						"session_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Session keep time, only valid when `session_switch` is true, the available value ranges from 1 to 300 and unit is second.",
						},
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Id of the 4 layer rule.",
						},
						"lb_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "LB type of the rule, 1 for weight cycling and 2 for IP hash.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudDayuL4RulesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloud_dayu_l7_rules.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	service := DayuService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	resourceType := d.Get("resource_type").(string)
	resourceId := d.Get("resource_id").(string)
	ruleId := ""
	if v, ok := d.GetOk("rule_id"); ok {
		ruleId = v.(string)
	}
	name := ""
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	rules, healths, _, err := service.DescribeL4Rules(ctx, resourceType, resourceId, name, ruleId)
	if err != nil {
		rules, healths, _, err = service.DescribeL4Rules(ctx, resourceType, resourceId, name, ruleId)
	}

	if err != nil {
		return err
	}

	list := make([]map[string]interface{}, 0, len(rules))
	ids := make([]string, 0, len(rules))

	listItem := make(map[string]interface{})
	for k, rule := range rules {
		health := healths[k]
		listItem["name"] = *rule.RuleName
		listItem["protocol"] = *rule.Protocol
		listItem["source_port"] = int(*rule.SourcePort)
		listItem["dest_port"] = int(*rule.VirtualPort)
		listItem["rule_id"] = *rule.RuleId
		listItem["lb_type"] = int(*rule.LbType)
		listItem["source_type"] = int(*rule.SourceType)
		listItem["session_time"] = int(*rule.KeepTime)
		listItem["session_switch"] = intToBool(int(*rule.KeepEnable))
		listItem["source_list"] = flattenSourceList(rule.SourceList)
		if health.Enable != nil {
			listItem["health_check_switch"] = intToBool(int(*health.Enable))
		}
		if health.TimeOut != nil {
			listItem["health_check_timeout"] = int(*health.TimeOut)
		}
		if health.Interval != nil {
			listItem["health_check_interval"] = int(*health.Interval)
		}
		if health.KickNum != nil {
			listItem["health_check_unhealth_num"] = int(*health.KickNum)
		}
		if health.AliveNum != nil {
			listItem["health_check_health_num"] = int(*health.AliveNum)
		}
		list = append(list, listItem)
		ids = append(ids, listItem["rule_id"].(string))
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	if e := d.Set("list", list); e != nil {
		log.Printf("[CRITAL]%s provider set list fail, reason:%s\n", logId, e.Error())
		return e
	}
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		return writeToFile(output.(string), list)
	}
	return nil

}
