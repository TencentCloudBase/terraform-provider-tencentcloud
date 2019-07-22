package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccTencentCloudSecurityGroup_basic(t *testing.T) {
	var sgId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityGroupDestroy(&sgId),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupExists("tencentcloud_security_group.foo", &sgId),
					resource.TestCheckResourceAttr("tencentcloud_security_group.foo", "name", "ci-temp-test-sg"),
					resource.TestCheckResourceAttr("tencentcloud_security_group.foo", "description", ""),
				),
			},
			{
				ResourceName:      "tencentcloud_security_group.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudSecurityGroup_update(t *testing.T) {
	var sgId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityGroupDestroy(&sgId),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupExists("tencentcloud_security_group.foo", &sgId),
					resource.TestCheckResourceAttr("tencentcloud_security_group.foo", "name", "ci-temp-test-sg"),
					resource.TestCheckResourceAttr("tencentcloud_security_group.foo", "description", "ci-temp-test-sg-desc"),
				),
			},
			{
				Config: testAccSecurityGroupConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupExists("tencentcloud_security_group.foo", &sgId),
					resource.TestCheckResourceAttr("tencentcloud_security_group.foo", "name", "ci-temp-test-sg-updated"),
					resource.TestCheckResourceAttr("tencentcloud_security_group.foo", "description", "ci-temp-test-sg-desc-updated"),
				),
			},
		},
	})
}

func testAccCheckSecurityGroupDestroy(id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*TencentCloudClient).apiV3Conn
		service := VpcService{client: client}

		_, has, err := service.DescribeSecurityGroup(context.TODO(), *id)
		if err != nil {
			return err
		}

		if has != 0 {
			return fmt.Errorf("security group still exists")
		}

		return nil
	}
}

func testAccCheckSecurityGroupExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no security group ID is set")
		}

		service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

		_, has, err := service.DescribeSecurityGroup(context.TODO(), rs.Primary.ID)
		if err != nil {
			return err
		}

		if has == 0 {
			return fmt.Errorf("security group not found: %s", rs.Primary.ID)
		}

		*id = rs.Primary.ID

		return nil
	}
}

const testAccSecurityGroupConfigBasic = `
resource "tencentcloud_security_group" "foo" {
  name = "ci-temp-test-sg"
}
`

const testAccSecurityGroupConfig = `
resource "tencentcloud_security_group" "foo" {
  name = "ci-temp-test-sg"
  description = "ci-temp-test-sg-desc"
}
`
const testAccSecurityGroupConfigUpdate = `
resource "tencentcloud_security_group" "foo" {
  name = "ci-temp-test-sg-updated"
  description = "ci-temp-test-sg-desc-updated"
}
`
