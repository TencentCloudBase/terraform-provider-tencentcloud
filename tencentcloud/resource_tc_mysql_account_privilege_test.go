package tencentcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	sdkError "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

func TestAccTencentCloudMysqlAccountPrivilege(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccMysqlAccountPrivilegeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMysqlAccountPrivilege(mysqlInstanceCommonTestCase),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccMysqlAccountPrivilegeExists("tencentcloud_mysql_account_privilege.mysql_account_privilege"),
					resource.TestCheckResourceAttrSet("tencentcloud_mysql_account_privilege.mysql_account_privilege", "mysql_id"),
					resource.TestCheckResourceAttrSet("tencentcloud_mysql_account_privilege.mysql_account_privilege", "account_name"),
					resource.TestCheckResourceAttr("tencentcloud_mysql_account_privilege.mysql_account_privilege", "database_names.#", "1"),

					resource.TestCheckResourceAttr("tencentcloud_mysql_account_privilege.mysql_account_privilege", "privileges.#", "4"),
					resource.TestCheckResourceAttr("tencentcloud_mysql_account_privilege.mysql_account_privilege", "privileges.1274211008", "SELECT"),
					resource.TestCheckResourceAttr("tencentcloud_mysql_account_privilege.mysql_account_privilege", "privileges.2552575352", "UPDATE"),
					resource.TestCheckResourceAttr("tencentcloud_mysql_account_privilege.mysql_account_privilege", "privileges.3318521589", "INSERT"),
					resource.TestCheckResourceAttr("tencentcloud_mysql_account_privilege.mysql_account_privilege", "privileges.974290055", "DELETE"),
				),
			},
			{
				Config: testAccMysqlAccountPrivilegeUpdate(mysqlInstanceCommonTestCase),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccMysqlAccountPrivilegeExists("tencentcloud_mysql_account_privilege.mysql_account_privilege"),
					resource.TestCheckResourceAttrSet("tencentcloud_mysql_account_privilege.mysql_account_privilege", "mysql_id"),
					resource.TestCheckResourceAttrSet("tencentcloud_mysql_account_privilege.mysql_account_privilege", "account_name"),
					resource.TestCheckResourceAttr("tencentcloud_mysql_account_privilege.mysql_account_privilege", "database_names.#", "1"),

					resource.TestCheckResourceAttr("tencentcloud_mysql_account_privilege.mysql_account_privilege", "privileges.#", "1"),
					resource.TestCheckResourceAttr("tencentcloud_mysql_account_privilege.mysql_account_privilege", "privileges.443223901", "TRIGGER"),
				),
			},
		},
	})
}

func testAccMysqlAccountPrivilegeExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), "logId", logId)
		mysqlService := MysqlService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}
		var privilegeId resourceTencentCloudMysqlAccountPrivilegeId

		if err := json.Unmarshal([]byte(rs.Primary.ID), &privilegeId); err != nil {
			return fmt.Errorf("Local data[terraform.tfstate] corruption,can not got old account privilege id")
		}

		_, err := mysqlService.DescribeAccountPrivileges(ctx, privilegeId.MysqlId, privilegeId.AccountName, privilegeId.AccountHost, []string{"test"})
		if err != nil {
			if sdkErr, ok := err.(*sdkError.TencentCloudSDKError); ok {
				if sdkErr.Code == MysqlInstanceIdNotFound {
					return fmt.Errorf("privilege not exists in mysql")
				}
				if sdkErr.Code == "InvalidParameter" && strings.Contains(sdkErr.GetMessage(), "instance not found") {
					return fmt.Errorf("privilege not exists in mysql")
				}
				if sdkErr.Code == "InternalError.TaskError" && strings.Contains(sdkErr.Message, "User does not exist") {
					return fmt.Errorf("privilege not exists in mysql")
				}
			}
		}

		if err != nil {
			return err
		}

		var accountInfo *cdb.AccountInfo = nil

		accountInfos, err := mysqlService.DescribeAccounts(ctx, privilegeId.MysqlId)
		if err != nil {
			return err
		}

		for _, account := range accountInfos {
			if *account.User == privilegeId.AccountName && *account.Host == privilegeId.AccountHost {
				accountInfo = account
				break
			}
		}

		if accountInfo == nil {
			return fmt.Errorf("mysql account %s is not found", privilegeId.MysqlId)
		}
		return nil
	}

}

func testAccMysqlAccountPrivilegeDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	mysqlService := MysqlService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloud_mysql_account_privilege" {
			continue
		}
		var privilegeId resourceTencentCloudMysqlAccountPrivilegeId

		if err := json.Unmarshal([]byte(rs.Primary.ID), &privilegeId); err != nil {
			return fmt.Errorf("Local data[terraform.tfstate] corruption,can not got old account privilege id")
		}

		privileges, err := mysqlService.DescribeAccountPrivileges(ctx, privilegeId.MysqlId, privilegeId.AccountName, privilegeId.AccountHost, []string{"test"})

		if err != nil {
			if sdkErr, ok := err.(*sdkError.TencentCloudSDKError); ok {
				if sdkErr.Code == MysqlInstanceIdNotFound {
					return nil
				}
				if sdkErr.Code == "InvalidParameter" && strings.Contains(sdkErr.GetMessage(), "instance not found") {
					return nil
				}
				if sdkErr.Code == "InternalError.TaskError" && strings.Contains(sdkErr.Message, "User does not exist") {
					return nil
				}

			}
		}

		if err != nil {
			return err
		}
		if len(privileges) != 1 || privileges[0] != MYSQL_DATABASE_MUST_PRIVILEGE {
			return fmt.Errorf("mysql  aacount privilege not clean ok")
		}
	}

	return nil
}

func testAccMysqlAccountPrivilege(commonTestCase string) string {
	return fmt.Sprintf(`
%s
resource "tencentcloud_mysql_account" "mysql_account" {
  mysql_id    = tencentcloud_mysql_instance.default.id
  name        = "test"
  host        = "119.168.110.%%"
  password    = "test1234"
  description = "test from terraform"
}
resource "tencentcloud_mysql_account_privilege" "mysql_account_privilege" {
  mysql_id       = tencentcloud_mysql_instance.default.id
  account_name   = tencentcloud_mysql_account.mysql_account.name
  account_host   = tencentcloud_mysql_account.mysql_account.host
  privileges     = ["SELECT", "INSERT", "UPDATE", "DELETE"]
  database_names = ["test"]
}`, commonTestCase)
}

func testAccMysqlAccountPrivilegeUpdate(commonTestCase string) string {
	return fmt.Sprintf(`
%s
resource "tencentcloud_mysql_account" "mysql_account" {
  mysql_id    = tencentcloud_mysql_instance.default.id
  name        = "test"
  host        = "119.168.110.%%"
  password    = "test1234"
  description = "test from terraform"
}
resource "tencentcloud_mysql_account_privilege" "mysql_account_privilege" {
  mysql_id       = tencentcloud_mysql_instance.default.id
  account_name   = tencentcloud_mysql_account.mysql_account.name
  account_host   = tencentcloud_mysql_account.mysql_account.host
  privileges     = ["TRIGGER"]
  database_names = ["test"]
}`, commonTestCase)

}
