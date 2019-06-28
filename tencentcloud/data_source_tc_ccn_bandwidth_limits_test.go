package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceTencentCloudCcnV3BandwidthLimits_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccDataSourceTencentCloudCcnBandwidthLimits,

				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloud_ccn_bandwidth_limits.limit"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_ccn_bandwidth_limits.limit", "ccn_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_ccn_bandwidth_limits.limit", "limits.#"),
				),
			},
		},
	})
}

const TestAccDataSourceTencentCloudCcnBandwidthLimits = `
resource tencentcloud_ccn main{
	name ="ci-temp-test-ccn"
	description="ci-temp-test-ccn-des"
	qos ="AG"
}

data tencentcloud_ccn_bandwidth_limits limit {
	ccn_id ="${tencentcloud_ccn.main.id}"
}

`
