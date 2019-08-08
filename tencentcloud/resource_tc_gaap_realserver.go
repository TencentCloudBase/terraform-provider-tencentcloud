package tencentcloud

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTencentCloudGaapRealserver() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:          schema.TypeString,
				Optional:      true,
				ValidateFunc:  validateIp,
				ConflictsWith: []string{"domain"},
				ForceNew:      true,
			},
			"domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"ip"},
				ForceNew:      true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringLengthInRange(1, 30),
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				Default:  0,
				ForceNew: true,
			},
			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}
