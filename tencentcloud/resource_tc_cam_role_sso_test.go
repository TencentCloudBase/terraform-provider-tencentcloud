package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func init() {
	// go test -v ./tencentcloud -sweep=ap-guangzhou -sweep-run=tencentcloud_cam_role_sso
	resource.AddTestSweepers("tencentcloud_cam_role_sso", &resource.Sweeper{
		Name: "tencentcloud_cam_role_sso",
		F: func(r string) error {
			logId := getLogId(contextNil)
			ctx := context.WithValue(context.TODO(), logIdKey, logId)
			cli, _ := sharedClientForRegion(r)
			client := cli.(*TencentCloudClient).apiV3Conn

			request := cam.NewDescribeOIDCConfigRequest()
			request.Name = helper.String(defaultSSORoleName)
			response, err := client.UseCamClient().DescribeOIDCConfig(request)

			if err != nil {
				return err
			}

			dReq := cam.NewDeleteOIDCConfigRequest()
			dReq.Name = response.Response.Name
			dRes, err := client.UseCamClient().DeleteOIDCConfig(dReq)

			if err != nil {
				return err
			}

			log.Printf("[%s] success, request %s, response %s, logId %s", dReq.GetAction(), dReq.ToJsonString(), dRes.ToJsonString(), ctx.Value(logIdKey))

			return nil
		},
	})
}

func TestAccTencentCloudCamRoleSSO(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamRoleSSODestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamRoleSSO,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamRoleSSOExists("tencentcloud_cam_role_sso.foo"),
					resource.TestCheckResourceAttrSet("tencentcloud_cam_role_sso.foo", "name"),
					resource.TestCheckResourceAttrSet("tencentcloud_cam_role_sso.foo", "identity_url"),
					resource.TestCheckResourceAttrSet("tencentcloud_cam_role_sso.foo", "identity_key"),
					resource.TestCheckResourceAttrSet("tencentcloud_cam_role_sso.foo", "description"),
					resource.TestCheckResourceAttrSet("tencentcloud_cam_role_sso.foo", "client_ids.#"),
				),
			},
			{
				Config: testAccCamRoleSSOUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamRoleSSOExists("tencentcloud_cam_role_sso.foo"),
					resource.TestCheckResourceAttr("tencentcloud_cam_role_sso.foo", "description", "this is a description1"),
				),
			},
			{
				ResourceName:      "tencentcloud_cam_role_sso.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCamRoleSSODestroy(s *terraform.State) error {

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloud_cam_role_sso" {
			continue
		}

		request := cam.NewDescribeOIDCConfigRequest()
		request.Name = helper.String(rs.Primary.ID)
		response, err := testAccProvider.Meta().(*TencentCloudClient).apiV3Conn.UseCamClient().DescribeOIDCConfig(request)
		if err != nil {
			if e, ok := err.(*errors.TencentCloudSDKError); ok {
				if e.GetCode() == "ResourceNotFound.IdentityNotExist" {
					return nil
				}
			}
			return err
		}
		if response != nil && *response.Response.Status != 2 {
			return fmt.Errorf("[Exists] check: CAM-ROLE-SSO %s is still exist!", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckCamRoleSSOExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[Exists] check: CAM-ROLE-SSO %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[Exists] check: CAM-ROLE-SSO id is not set")
		}
		request := cam.NewDescribeOIDCConfigRequest()
		request.Name = helper.String(rs.Primary.ID)
		response, err := testAccProvider.Meta().(*TencentCloudClient).apiV3Conn.UseCamClient().DescribeOIDCConfig(request)
		if err != nil {
			return err
		}
		if response.Response == nil {
			return fmt.Errorf("[Exists] check: CAM-ROLE-SSO %s is not found", n)
		}
		return nil
	}
}

const defaultSSORoleName = "test"

const testAccCamRoleSSO = `
resource "tencentcloud_cam_role_sso" "foo" {
	name="` + defaultSSORoleName + `"
	identity_url="https://login.microsoftonline.com/ebc33bb0-4c95-4673-9b64-61bfa32cf297/v2.0"
	identity_key="CXsKICAgICJrZXlzIjogWwogICAgICAgIHsKICAgICAgICAgICAgImt0eSI6ICJSU0EiLAogICAgICAgICAgICAidXNlIjogInNpZyIsCiAgICAgICAgICAgICJraWQiOiAibk9vM1pEck9EWEVLMWpLV2hYc2xIUl9LWEVnIiwKICAgICAgICAgICAgIng1dCI6ICJuT28zWkRyT0RYRUsxaktXaFhzbEhSX0tYRWciLAogICAgICAgICAgICAibiI6ICJvYUxMVDloa2NTajJ0R2Zac2pidTdYejFLcnMwcUVpY1hQbUVzSktPQlFIYXVaX2tSTTFIZEVrZ09KYlV6blVzcEU2eE91T1NYamx6RXJxQnhYQXU0U0N2Y3ZWT0NZRzJ2OUczLXVJckxGNWRzdEQwc1lIQm8xVm9tdEt4ekY5MFZzbHJrbjZyTlFnVUdJV2d2dVFUeG0xdVJrbFlGUEVjVElSdzBMbllrbnpKMDZHQzlsaktSNjE3d0FCVnJaTmtCdURnUUtqMzdxY3l4b2F4SUdkeEVjbVZGWlhKeXJ4RGdkWGg5b3dSbVpuNkxJSmxHalo5bTU5ZW1mdXduQm5zSVFHN0Rpckp3ZTlTWHJMWG5leFJRV3F5ekNka1lhT3FrcEtyc2p1eFVqMi1NSFgzMUZxc2RwSkpzT0F2WVhHT1lCS0pSamhHckdkT05WclpkVWRUQlEiLAogICAgICAgICAgICAiZSI6ICJBUUFCIiwKICAgICAgICAgICAgIng1YyI6IFsKICAgICAgICAgICAgICAgICJNSUlEQlRDQ0FlMmdBd0lCQWdJUU4zM1JPYUlKNmJKQldEQ3h0bUpFYmpBTkJna3Foa2lHOXcwQkFRc0ZBREF0TVNzd0tRWURWUVFERXlKaFkyTnZkVzUwY3k1aFkyTmxjM05qYjI1MGNtOXNMbmRwYm1SdmQzTXVibVYwTUI0WERUSXdNVEl5TVRJd05UQXhOMW9YRFRJMU1USXlNREl3TlRBeE4xb3dMVEVyTUNrR0ExVUVBeE1pWVdOamIzVnVkSE11WVdOalpYTnpZMjl1ZEhKdmJDNTNhVzVrYjNkekxtNWxkRENDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFLR2l5MC9ZWkhFbzlyUm4yYkkyN3UxODlTcTdOS2hJbkZ6NWhMQ1NqZ1VCMnJtZjVFVE5SM1JKSURpVzFNNTFMS1JPc1RyamtsNDVjeEs2Z2NWd0x1RWdyM0wxVGdtQnRyL1J0L3JpS3l4ZVhiTFE5TEdCd2FOVmFKclNzY3hmZEZiSmE1SitxelVJRkJpRm9MN2tFOFp0YmtaSldCVHhIRXlFY05DNTJKSjh5ZE9oZ3ZaWXlrZXRlOEFBVmEyVFpBYmc0RUNvOSs2bk1zYUdzU0JuY1JISmxSV1Z5Y3E4UTRIVjRmYU1FWm1aK2l5Q1pSbzJmWnVmWHBuN3NKd1o3Q0VCdXc0cXljSHZVbDZ5MTUzc1VVRnFzc3duWkdHanFwS1NxN0k3c1ZJOXZqQjE5OVJhckhhU1NiRGdMMkZ4am1BU2lVWTRScXhuVGpWYTJYVkhVd1VDQXdFQUFhTWhNQjh3SFFZRFZSME9CQllFRkk1bU41ZnRIbG9FRFZOb0lhOHNRczdrSkFlVE1BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQm5hR25vanhOZ25WNCtUQ1BaOWJyNG94MW5Sbjl0elk4YjVwd0tUVzJNY0pUZTB5RXZySHlhSXRLOEtibWVLSk9CdkFTZitRd0hrcCtGMkJBWHpSaVRsNForZ05GUVVMUHpzUVdwbUtsejZmSVdoYzdrc2dwVGtNSzZBYVRid1dZVGZtcEtuUXcvS0ptLzZyYm9MRFdZeUtGcFFjU3R1NjdSWithUnZRejY4RXYyZ2E1SnNYbGNPSjNnUC9sRTVXQzFTMHJqZmFiemRNT0dQOHFaUWhYazR3Qk9ndEZCYWlzRG5ialY1cGNJcmpSUGxob0N4dktnQy8yOTBuWjkvRExCSDNUYkhrOHh3SFhlQkFuQWp5QXFPWmlqOTJ1a3NBdjdaTHE0TU9EY25Rc2hWSU5Yd3NZc2hHMXBRcU9Md01lcnROYVk1V3RydWJNUmt1NDREdzdSIgogICAgICAgICAgICBdLAogICAgICAgICAgICAiaXNzdWVyIjogImh0dHBzOi8vbG9naW4ubWljcm9zb2Z0b25saW5lLmNvbS9lYmMzM2JiMC00Yzk1LTQ2NzMtOWI2NC02MWJmYTMyY2YyOTcvdjIuMCIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgImt0eSI6ICJSU0EiLAogICAgICAgICAgICAidXNlIjogInNpZyIsCiAgICAgICAgICAgICJraWQiOiAibDNzUS01MGNDSDR4QlZaTEhUR3duU1I3NjgwIiwKICAgICAgICAgICAgIng1dCI6ICJsM3NRLTUwY0NINHhCVlpMSFRHd25TUjc2ODAiLAogICAgICAgICAgICAibiI6ICJzZnNYTVhXdU8tZG5pTGFJRUxhM1B5cXo5WV9yV2ZmX0FWckNBbkZTZFBIYThfX1Bta2J0X3lxLTZaM3UxbzRnalJwS1ducmp4SWg4ekRuMVoxUlMyNm5rS2NOZzV4Zld4UjJLOENQYlNiWThnTXJwXzRwWm43dGdyRW1vTE1rd2ZnWWFWQy00TWlGRW8xUDJnZDltQ2RnSUlDYU5lWWtHMWJJUFRuYXFxdVRNNUtmVDk3MU1wdU9WT2RNMXlzaWVqZGNORHZFYjd2Mjg0UFlaa3cyaW13cWlCWTNGUjBzVkc3amdLVW90RnZoZDdUUjVXc0EyMEdTXzZaSWtVVWxMVWJHX3JYV0dsMFlqWkxTX1VmNHE4SGJvN3UtN01hRm44QjY5RjZZYUZkRGxYbV9BMFNwZWRWRldRRkd6TXNwNDNfNnZFempmckZESlZBWWt3YjZ4VVEiLAogICAgICAgICAgICAiZSI6ICJBUUFCIiwKICAgICAgICAgICAgIng1YyI6IFsKICAgICAgICAgICAgICAgICJNSUlEQlRDQ0FlMmdBd0lCQWdJUVdQQjFvZk9wQTdGRmxPQms1aVBhTlRBTkJna3Foa2lHOXcwQkFRc0ZBREF0TVNzd0tRWURWUVFERXlKaFkyTnZkVzUwY3k1aFkyTmxjM05qYjI1MGNtOXNMbmRwYm1SdmQzTXVibVYwTUI0WERUSXhNREl3TnpFM01EQXpPVm9YRFRJMk1ESXdOakUzTURBek9Wb3dMVEVyTUNrR0ExVUVBeE1pWVdOamIzVnVkSE11WVdOalpYTnpZMjl1ZEhKdmJDNTNhVzVrYjNkekxtNWxkRENDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFMSDdGekYxcmp2blo0aTJpQkMydHo4cXMvV1A2MW4zL3dGYXdnSnhVblR4MnZQL3o1cEc3ZjhxdnVtZDd0YU9JSTBhU2xwNjQ4U0lmTXc1OVdkVVV0dXA1Q25EWU9jWDFzVWRpdkFqMjBtMlBJREs2ZitLV1orN1lLeEpxQ3pKTUg0R0dsUXZ1REloUktOVDlvSGZaZ25ZQ0NBbWpYbUpCdFd5RDA1MnFxcmt6T1NuMC9lOVRLYmpsVG5UTmNySW5vM1hEUTd4Rys3OXZPRDJHWk1Ob3BzS29nV054VWRMRlJ1NDRDbEtMUmI0WGUwMGVWckFOdEJrdittU0pGRkpTMUd4djYxMWhwZEdJMlMwdjFIK0t2QjI2Tzd2dXpHaFovQWV2UmVtR2hYUTVWNXZ3TkVxWG5WUlZrQlJzekxLZU4vK3J4TTQzNnhReVZRR0pNRytzVkVDQXdFQUFhTWhNQjh3SFFZRFZSME9CQllFRkxsUkJTeHhnbU5QT2JDRnJsK2hTc2JjdlJrY01BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQitVUUZUTnM2QlVZM0FJR2tTMlpSdVpnSnNORXIvWkVNNGFDczJkb21kMk9xajcrNWlXc25QaDVDdWdGbkk0bmQrWkxnS1ZIU0Q2YWNRMjd3ZStlTlk2Z3hmcFFDWTFmaU4vdUtPT3NBMElmOEliUGRCRWh0UGVyUmdQSkZYTEhhWVZxRDhVWURvNUtOQ2NvQjRLaDhudkNXUkdQVVVIUFJxcDdBbkFjVnJjYmlYQS9ibU1DbkZXdU5OYWhjYUFLaUpUeFlsS0RhRElpUE4zNXlFQ1liRGowUEJXSlV4b2Jydmo1STI3NWpiaWtrcDhRU0xZblNVL3Y3ZE1EVWJ4U0xmWjd6c1R1YUYyUXgrTDYyUHNZVHdMeklGWDNNOEVNU1E2aDY4VHVwRlRpNW4wTTJ5SVhRZ29Sb05FRFdOSlovYVpNWS9ncVQwMkdRR0JXcmgrL3ZKIgogICAgICAgICAgICBdLAogICAgICAgICAgICAiaXNzdWVyIjogImh0dHBzOi8vbG9naW4ubWljcm9zb2Z0b25saW5lLmNvbS9lYmMzM2JiMC00Yzk1LTQ2NzMtOWI2NC02MWJmYTMyY2YyOTcvdjIuMCIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgImt0eSI6ICJSU0EiLAogICAgICAgICAgICAidXNlIjogInNpZyIsCiAgICAgICAgICAgICJraWQiOiAiTXI1LUFVaWJmQmlpN05kMWpCZWJheGJvWFcwIiwKICAgICAgICAgICAgIng1dCI6ICJNcjUtQVVpYmZCaWk3TmQxakJlYmF4Ym9YVzAiLAogICAgICAgICAgICAibiI6ICJ5cjN2MXVFVHJGZlQxN3p2T2l5MDF3OG5PLTF0NjdjbWlaTFp4cTJJU0RkdGU5ZHctSXhDUjdsUFYyd2V6Y3pJUmdjV21ZZ0Zuc2syajZtMTBINHRLemNxWk0wSkpfTmlnWTI5cEZpbXhsTDdfcVhNQjFQb3JGSmRsQUt2cDVTZ2pTVHdMclhqa3IxQXFXd2JwekcyeVpVTk4zR0U4R3ZtVGVvNHl3ZVFiTkNkLXlPX1pwb3p4MEozNHdIQkVNdWF3LVpmQ1VrN21kS0tzZy1FY0U0WnYwWGdsOXdQMk1wS1B4MFY4Z0xhenhlNlVROVNoek51cnVTT25jcExZSk5fb1E0YUtmNXB0T3AxcnNmRFkySUs5ZnJ0bVJUS09kUS1NRW1TZGpHTF84OElRY3ZDczdqcVZ6NTNYS29YUmxYQjh0TUlHT2NnLUlDZXI2eXhlMml0SVEiLAogICAgICAgICAgICAiZSI6ICJBUUFCIiwKICAgICAgICAgICAgIng1YyI6IFsKICAgICAgICAgICAgICAgICJNSUlEQlRDQ0FlMmdBd0lCQWdJUWZmOHlyRk8zQ0lOUEhVVFQ3NnRVc1RBTkJna3Foa2lHOXcwQkFRc0ZBREF0TVNzd0tRWURWUVFERXlKaFkyTnZkVzUwY3k1aFkyTmxjM05qYjI1MGNtOXNMbmRwYm1SdmQzTXVibVYwTUI0WERUSXhNVEF5TkRFM05EVTFObG9YRFRJMk1UQXlOREUzTkRVMU5sb3dMVEVyTUNrR0ExVUVBeE1pWVdOamIzVnVkSE11WVdOalpYTnpZMjl1ZEhKdmJDNTNhVzVrYjNkekxtNWxkRENDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFNcTk3OWJoRTZ4WDA5ZTg3em9zdE5jUEp6dnRiZXUzSm9tUzJjYXRpRWczYlh2WGNQaU1Ra2U1VDFkc0hzM015RVlIRnBtSUJaN0pObytwdGRCK0xTczNLbVROQ1NmellvR052YVJZcHNaUysvNmx6QWRUNkt4U1haUUNyNmVVb0kwazhDNjE0NUs5UUtsc0c2Y3h0c21WRFRkeGhQQnI1azNxT01zSGtHelFuZnNqdjJhYU04ZENkK01Cd1JETG1zUG1Yd2xKTzVuU2lySVBoSEJPR2I5RjRKZmNEOWpLU2o4ZEZmSUMyczhYdWxFUFVvY3picTdranAzS1MyQ1RmNkVPR2luK2FiVHFkYTdIdzJOaUN2WDY3WmtVeWpuVVBqQkprbll4aS8vUENFSEx3ck80NmxjK2QxeXFGMFpWd2ZMVENCam5JUGlBbnErc3NYdG9yU0VDQXdFQUFhTWhNQjh3SFFZRFZSME9CQllFRkRpWkc2czVkOVJ2b3JwcWJWZFMyL01EOFpLaE1BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQVFBUHVxcUtqMkFnZkM5YXl4K3FVdTB2anpLWWRaNlQrM3NzSkRPR3dCMWNMTVhNVFVWZ0Z3ajhic1gxYWhEVUpkektwV3ROajdibm8rVWc4NUl5VTdrODlVMFlncjU1eldVNWg0d25uUnJDdTlRS3Z1ZFVQbmJpWG9WdUhQd2NLOHcxZmRYWlFCNVFxL2tLemhOR1k1N2NHMWJ3ajNSL2FJZENwK0JqZ0ZwcE9LakpwSzdGS1M4RzJ2NzBlSWlDTE1hcEs5bExFZVFPeEl2emN0VHNYeTlFWjd3dGFJaVlreTRaU2l0dXBoVG9KVWtha0hhUTZldmJuODJsVGc2V1p6MXRtU21ZblBxUmRBZmY3YWlRMVN3OUhwdXpsWlkvcGlUVnF2ZDZBZktacXl4dS9GaEVORTBPZHYvMGhsSHpJMTVqS1FXTDFMamMwTm0zeTFza3V0IgogICAgICAgICAgICBdLAogICAgICAgICAgICAiaXNzdWVyIjogImh0dHBzOi8vbG9naW4ubWljcm9zb2Z0b25saW5lLmNvbS9lYmMzM2JiMC00Yzk1LTQ2NzMtOWI2NC02MWJmYTMyY2YyOTcvdjIuMCIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgImt0eSI6ICJSU0EiLAogICAgICAgICAgICAidXNlIjogInNpZyIsCiAgICAgICAgICAgICJraWQiOiAialMxWG8xT1dEal81MnZid0dOZ3ZRTzJWek1jIiwKICAgICAgICAgICAgIng1dCI6ICJqUzFYbzFPV0RqXzUydmJ3R05ndlFPMlZ6TWMiLAogICAgICAgICAgICAibiI6ICJzcHZRY1hXcVlyTWN2Y3FRbWZTTVluYlVDOFUwM1ljdG5YeUxJQmUxNDhPemhCcmdkQU9tUGZNZkppX3RVVzhMOXN2VkdwazVxRzZkTjBuNjY5Y1JIS3FVNTJHbkcwdGx5WVhtekZDMWh6SFZnUXo5ZWh2ZTR0bEo3dXc5MzZYSVVPQU94eDNYMjB6ZHB4N2dtNHpIeDRqMlpCbFhza0FqNlUzYWRwSFFOdXdVRTZrbW5nSldSLWRlV2xFaWdNcFJzdlVWUTJPNWgwLVJTcThXcl94N3VkM0s2R1R0cnpBUmFtejl1azJJWGF0S1lkbmo1SnJrMmpMWTZuV3QtR3R4bEFfbDlYd0lyT2w2U3FhX3BPR0lwUzAxSktkeEt2cEJDOVZkUzhvWEItN1A1cUxrc212N3RxLVNiYmlPZWMwY3ZVN1dQN3ZVUnYxMDRWNEZpSV9xb1EiLAogICAgICAgICAgICAiZSI6ICJBUUFCIiwKICAgICAgICAgICAgIng1YyI6IFsKICAgICAgICAgICAgICAgICJNSUlEQlRDQ0FlMmdBd0lCQWdJUUhzZXRQK2k4aTZWSUFtam1mVkd2NmpBTkJna3Foa2lHOXcwQkFRc0ZBREF0TVNzd0tRWURWUVFERXlKaFkyTnZkVzUwY3k1aFkyTmxjM05qYjI1MGNtOXNMbmRwYm1SdmQzTXVibVYwTUI0WERUSXlNREV6TURJek1EWXhORm9YRFRJM01ERXpNREl6TURZeE5Gb3dMVEVyTUNrR0ExVUVBeE1pWVdOamIzVnVkSE11WVdOalpYTnpZMjl1ZEhKdmJDNTNhVzVrYjNkekxtNWxkRENDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFMS2IwSEYxcW1LekhMM0trSm4wakdKMjFBdkZOTjJITFoxOGl5QVh0ZVBEczRRYTRIUURwajN6SHlZdjdWRnZDL2JMMVJxWk9haHVuVGRKK3V2WEVSeXFsT2RocHh0TFpjbUY1c3hRdFljeDFZRU0vWG9iM3VMWlNlN3NQZCtseUZEZ0RzY2QxOXRNM2FjZTRKdU14OGVJOW1RWlY3SkFJK2xOMm5hUjBEYnNGQk9wSnA0Q1ZrZm5YbHBSSW9ES1ViTDFGVU5qdVlkUGtVcXZGcS84ZTduZHl1aGs3YTh3RVdwcy9icE5pRjJyU21IWjQrU2E1Tm95Mk9wMXJmaHJjWlFQNWZWOENLenBla3FtdjZUaGlLVXROU1NuY1NyNlFRdlZYVXZLRndmdXorYWk1TEpyKzdhdmttMjRqbm5OSEwxTzFqKzcxRWI5ZE9GZUJZaVA2cUVDQXdFQUFhTWhNQjh3SFFZRFZSME9CQllFRkd6VkZqQWJZcFUvMmVuNHJ5NExNTFVISjNHak1BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQlUwWWROVmZkQnl2cHdzUGZ3TmREOG0xUExlZUNLbUxIUW5XUkk1NjAweUVIdW9Vdm9BSmQ1ZHdlMVpVMWJISFJSS1dON0FrdFV6b2ZQM3lGNjF4dGl6aEVieVBqSEsxdG5SK2lQRXZpV3hWdkszN0h0ZkVQenVoMVZxcDA4YnFZMTVNY1lVdGY3N2wySFhUcGFrK1VXWVJZSkJpKysydW1JREtZNVVNcVUrTEVabnZhWHliTFVLTjN4RzRpeTJxMUFiOHN5R0ZhVVA3SjNuQ3RWclI3aXAzOUJudlNUVFpaTm8vT0M3ZllYSjJYNHNOMS8yWmhSNUV0bkFnd2kyUnZsWmwwYVdQcmN6QXJVQ3hEQkNic0tQTC9VcC9rSUQxaXIxVk80TFQwOXJ5ZnYybngzeTZsMFl2dUw3ZVB6NG5HWUNXSGNiTVZjVXJRVVhxdVozWHRJIgogICAgICAgICAgICBdLAogICAgICAgICAgICAiaXNzdWVyIjogImh0dHBzOi8vbG9naW4ubWljcm9zb2Z0b25saW5lLmNvbS9lYmMzM2JiMC00Yzk1LTQ2NzMtOWI2NC02MWJmYTMyY2YyOTcvdjIuMCIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgImt0eSI6ICJSU0EiLAogICAgICAgICAgICAidXNlIjogInNpZyIsCiAgICAgICAgICAgICJraWQiOiAiRHFVdThnZi1uQWdjeWpQMy1TdXBsTkFYQW5jIiwKICAgICAgICAgICAgIng1dCI6ICJEcVV1OGdmLW5BZ2N5alAzLVN1cGxOQVhBbmMiLAogICAgICAgICAgICAibiI6ICIxbjctbldTTGV1V1F6QlJsWVNiUzhSanZXdmtRZUQ3UUw5Zk9XYUdYYlc3M1ZOR0gwWWlwWmlzUENsRnY2R3p3ZldFQ1RXUXAxOVdGZV9sQVNrYTUtS0VXa1FWekNiRU1hYWFmT0lzN2hDNjFQNWNHZ3c3ZGh1VzRzN2Y2WllHWkV6UTRGNXJIRS1ZTlJidkQ1MXFpclBOektIazNuamkxd3JoMFl0YlBQSWYtLU5iSTk4YkN3TExoOWF2ZWRPbXFFU3pXT0dFQ0VNWHY4TFNNLUI5U0tnXzRRdUJ0eUJ3d0lha1R1cW84NHN3VEJNNXc4UGRocFdaWkR0UGdIODdXei1fV2pXdms5OUFqWGw3bDhwV1BRSmlLTnVqdF9jazNOREZwemFMRXBwb2RoVXNJRDBwdFJBMDA4ZUNVNmw4VC11eDE5d1ptYl95Qm5IY1YzcEZXaFEiLAogICAgICAgICAgICAiZSI6ICJBUUFCIiwKICAgICAgICAgICAgIng1YyI6IFsKICAgICAgICAgICAgICAgICJNSUlDOFRDQ0FkbWdBd0lCQWdJUVlWay90SjFlNHBoSVN2VnJBQUxOS1RBTkJna3Foa2lHOXcwQkFRc0ZBREFqTVNFd0h3WURWUVFERXhoc2IyZHBiaTV0YVdOeWIzTnZablJ2Ym14cGJtVXVkWE13SGhjTk1qQXhNakl4TURBd01EQXdXaGNOTWpVeE1qSXhNREF3TURBd1dqQWpNU0V3SHdZRFZRUURFeGhzYjJkcGJpNXRhV055YjNOdlpuUnZibXhwYm1VdWRYTXdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFEV2Z2NmRaSXQ2NVpETUZHVmhKdEx4R085YStSQjRQdEF2MTg1Wm9aZHRidmRVMFlmUmlLbG1LdzhLVVcvb2JQQjlZUUpOWkNuWDFZVjcrVUJLUnJuNG9SYVJCWE1Kc1F4cHBwODRpenVFTHJVL2x3YUREdDJHNWJpenQvcGxnWmtUTkRnWG1zY1Q1ZzFGdThQbldxS3M4M01vZVRlZU9MWEN1SFJpMXM4OGgvNzQxc2ozeHNMQXN1SDFxOTUwNmFvUkxOWTRZUUlReGUvd3RJejRIMUlxRC9oQzRHM0lIREFocVJPNnFqeml6Qk1Fem5EdzkyR2xabGtPMCtBZnp0YlA3OWFOYStUMzBDTmVYdVh5bFk5QW1JbzI2TzM5eVRjME1Xbk5vc1NtbWgyRlN3Z1BTbTFFRFRUeDRKVHFYeFA2N0hYM0JtWnYvSUdjZHhYZWtWYUZBZ01CQUFHaklUQWZNQjBHQTFVZERnUVdCQlEyci8vbGdUUGNLdWdoRGt6bUN0Umx3K1A5U3pBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQXNGZFJ5Y3pOV2gvcXBZdmNJWmJEdldZemxybUZaYzZibGNVem5zOXpmN3NVV3RRWnJaUHU1RGJldFYyR3IycjNxdE1ES1hDVWFSK3Bxb3kzSTJ6eFRYM3g4YlROaFpEOVlBZ0FGbFRMTlN5ZFRhSzVSSHlCLzVrcjZCN1pKZU5JazNQUlZoUkd0NnliQ0pTalYvVllWa0xSNWZkTFArNUdodkJFU29iQVIvZDBudHJpVHpwNy90TE1iL29YeDd3NUh1MW0zSThycE1vY29YZkgyU0gxR0xtTVhqNk14MWR0d0NEWU02YnNiM2ZoV1J6OU85T01SNlFOaVRucThxOXduMVF6QkFuUmNzd1l6VDFMS0tCUE5GU2FzQ3ZMWU9DUE9aQ0wrVzhOOGpxYTlaUllOWUtXWHptaVNwdGdCRU0yNHQzbTVGVVd6V3FvTHU5cEljbmtQUT09IgogICAgICAgICAgICBdLAogICAgICAgICAgICAiaXNzdWVyIjogImh0dHBzOi8vbG9naW4ubWljcm9zb2Z0b25saW5lLmNvbS9lYmMzM2JiMC00Yzk1LTQ2NzMtOWI2NC02MWJmYTMyY2YyOTcvdjIuMCIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgImt0eSI6ICJSU0EiLAogICAgICAgICAgICAidXNlIjogInNpZyIsCiAgICAgICAgICAgICJraWQiOiAiT3paNURibWNzbzlRenQyTW9kR21paGczMEJvIiwKICAgICAgICAgICAgIng1dCI6ICJPelo1RGJtY3NvOVF6dDJNb2RHbWloZzMwQm8iLAogICAgICAgICAgICAibiI6ICIwMXJlOWEyQlVUdE50ZEZ6TE5JLVFFSFc4WGhEaURNRGJHTWt4SFJJWVhINDF6QmNjc1h3SDl2TWkwSHV4WEhwWE96d3RVWUt3bDkzWlIzN3RwNmxwdndsVTFIZVBObVpwSjlELVhBdlU3M3gwM1lLb1pFZGFGQjM5VnNWeUxpaDNmdVB2NkRQRTJxVC1UTkUzWDVZZElXT0dGcmNNa2NYTHNqTy1CQ3E0cWNTZEJIMmxCZ0VRVXVENm5xcmVMWnNnLWdQelNEaGpWU2NJVVpHaUQ4TTJzS3hBRGlJSG81S2xhWkl5dTMydDhKa2F2UDlqTTdJdFNBanppZzFXMnl2VlF6VVFaQS14WnFKbzJqeEIzZ19meWdkUFVISzZVTi1fY3FrcmZ4bjItVldIMXdNaGxtOTBTcHhUTUQ0SG9ZT1ZpejFnZ0g4R0NYMmFCaVg1T3pRNlEiLAogICAgICAgICAgICAiZSI6ICJBUUFCIiwKICAgICAgICAgICAgIng1YyI6IFsKICAgICAgICAgICAgICAgICJNSUlDOFRDQ0FkbWdBd0lCQWdJUVFyWFBYSWxVRTRKTVRNa1ZqKzAyWVRBTkJna3Foa2lHOXcwQkFRc0ZBREFqTVNFd0h3WURWUVFERXhoc2IyZHBiaTV0YVdOeWIzTnZablJ2Ym14cGJtVXVkWE13SGhjTk1qRXdNekV3TURBd01EQXdXaGNOTWpZd016RXdNREF3TURBd1dqQWpNU0V3SHdZRFZRUURFeGhzYjJkcGJpNXRhV055YjNOdlpuUnZibXhwYm1VdWRYTXdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFEVFd0NzFyWUZSTzAyMTBYTXMwajVBUWRieGVFT0lNd05zWXlURWRFaGhjZmpYTUZ4eXhmQWYyOHlMUWU3RmNlbGM3UEMxUmdyQ1gzZGxIZnUybnFXbS9DVlRVZDQ4Mlpta24wUDVjQzlUdmZIVGRncWhrUjFvVUhmMVd4WEl1S0hkKzQrL29NOFRhcFA1TTBUZGZsaDBoWTRZV3R3eVJ4Y3V5TTc0RUtyaXB4SjBFZmFVR0FSQlM0UHFlcXQ0dG15RDZBL05JT0dOVkp3aFJrYUlQd3phd3JFQU9JZ2Vqa3FWcGtqSzdmYTN3bVJxOC8yTXpzaTFJQ1BPS0RWYmJLOVZETlJCa0Q3Rm1vbWphUEVIZUQ5L0tCMDlRY3JwUTM3OXlxU3QvR2ZiNVZZZlhBeUdXYjNSS25GTXdQZ2VoZzVXTFBXQ0Fmd1lKZlpvR0pmazdORHBBZ01CQUFHaklUQWZNQjBHQTFVZERnUVdCQlRFQ2pCUkFORFBMR3JuMXA3cXR3c3d0QlU3SnpBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQXExSWI0RVJ2WEc1a2lWbWhmTE9wdW4yRWxWT0xZK1hrdlZseVZqcTM1clptU0lHeGdmRmMwOFFPUUZWbXJXUVlybHNzMExiYm9IMGNJS2lENityVm9aVE1XeEdFaWNPY0dORnpya2NHMHVsdTBjZ2hLWUlEM0dLRFRmdFlLRVBrdnUydlFtdWVxYTR0MnRUM1BsWUY3RmkyZGJvUjVZOTZVZ3M4enFOd2RCTVJtNjc3Ti90SkJrNTNDc09mOU5uQmF4WjFFR0FybUVISEliODB2T0RTdHYzNXVlTHJmTVJ0Q0YvSGNna0d4eTJVOGthQ3pZbW16SGg0ellEa2VDd00zQ3EyYkdrRytFZmU5aEZZZkRIdzEzRHpUUitoOXBQcUZGaUF4blozb2ZUOTZOclpIZFlqd2JmbU04Y3czbGRnMHhRekdjd1pqdHlZbXdKNnNEZFJ2UT09IgogICAgICAgICAgICBdLAogICAgICAgICAgICAiaXNzdWVyIjogImh0dHBzOi8vbG9naW4ubWljcm9zb2Z0b25saW5lLmNvbS9lYmMzM2JiMC00Yzk1LTQ2NzMtOWI2NC02MWJmYTMyY2YyOTcvdjIuMCIKICAgICAgICB9CiAgICBdCn0K"
	client_ids=["c2856e29-d344-49a9-88ed-de3a01d99753"]
	description="this is a description"
}
`

const testAccCamRoleSSOUpdate = `
resource "tencentcloud_cam_role_sso" "foo" {
	name="` + defaultSSORoleName + `"
	identity_url="https://login.microsoftonline.com/ebc33bb0-4c95-4673-9b64-61bfa32cf297/v2.0"
	identity_key="CXsKICAgICJrZXlzIjogWwogICAgICAgIHsKICAgICAgICAgICAgImt0eSI6ICJSU0EiLAogICAgICAgICAgICAidXNlIjogInNpZyIsCiAgICAgICAgICAgICJraWQiOiAibk9vM1pEck9EWEVLMWpLV2hYc2xIUl9LWEVnIiwKICAgICAgICAgICAgIng1dCI6ICJuT28zWkRyT0RYRUsxaktXaFhzbEhSX0tYRWciLAogICAgICAgICAgICAibiI6ICJvYUxMVDloa2NTajJ0R2Zac2pidTdYejFLcnMwcUVpY1hQbUVzSktPQlFIYXVaX2tSTTFIZEVrZ09KYlV6blVzcEU2eE91T1NYamx6RXJxQnhYQXU0U0N2Y3ZWT0NZRzJ2OUczLXVJckxGNWRzdEQwc1lIQm8xVm9tdEt4ekY5MFZzbHJrbjZyTlFnVUdJV2d2dVFUeG0xdVJrbFlGUEVjVElSdzBMbllrbnpKMDZHQzlsaktSNjE3d0FCVnJaTmtCdURnUUtqMzdxY3l4b2F4SUdkeEVjbVZGWlhKeXJ4RGdkWGg5b3dSbVpuNkxJSmxHalo5bTU5ZW1mdXduQm5zSVFHN0Rpckp3ZTlTWHJMWG5leFJRV3F5ekNka1lhT3FrcEtyc2p1eFVqMi1NSFgzMUZxc2RwSkpzT0F2WVhHT1lCS0pSamhHckdkT05WclpkVWRUQlEiLAogICAgICAgICAgICAiZSI6ICJBUUFCIiwKICAgICAgICAgICAgIng1YyI6IFsKICAgICAgICAgICAgICAgICJNSUlEQlRDQ0FlMmdBd0lCQWdJUU4zM1JPYUlKNmJKQldEQ3h0bUpFYmpBTkJna3Foa2lHOXcwQkFRc0ZBREF0TVNzd0tRWURWUVFERXlKaFkyTnZkVzUwY3k1aFkyTmxjM05qYjI1MGNtOXNMbmRwYm1SdmQzTXVibVYwTUI0WERUSXdNVEl5TVRJd05UQXhOMW9YRFRJMU1USXlNREl3TlRBeE4xb3dMVEVyTUNrR0ExVUVBeE1pWVdOamIzVnVkSE11WVdOalpYTnpZMjl1ZEhKdmJDNTNhVzVrYjNkekxtNWxkRENDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFLR2l5MC9ZWkhFbzlyUm4yYkkyN3UxODlTcTdOS2hJbkZ6NWhMQ1NqZ1VCMnJtZjVFVE5SM1JKSURpVzFNNTFMS1JPc1RyamtsNDVjeEs2Z2NWd0x1RWdyM0wxVGdtQnRyL1J0L3JpS3l4ZVhiTFE5TEdCd2FOVmFKclNzY3hmZEZiSmE1SitxelVJRkJpRm9MN2tFOFp0YmtaSldCVHhIRXlFY05DNTJKSjh5ZE9oZ3ZaWXlrZXRlOEFBVmEyVFpBYmc0RUNvOSs2bk1zYUdzU0JuY1JISmxSV1Z5Y3E4UTRIVjRmYU1FWm1aK2l5Q1pSbzJmWnVmWHBuN3NKd1o3Q0VCdXc0cXljSHZVbDZ5MTUzc1VVRnFzc3duWkdHanFwS1NxN0k3c1ZJOXZqQjE5OVJhckhhU1NiRGdMMkZ4am1BU2lVWTRScXhuVGpWYTJYVkhVd1VDQXdFQUFhTWhNQjh3SFFZRFZSME9CQllFRkk1bU41ZnRIbG9FRFZOb0lhOHNRczdrSkFlVE1BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQm5hR25vanhOZ25WNCtUQ1BaOWJyNG94MW5Sbjl0elk4YjVwd0tUVzJNY0pUZTB5RXZySHlhSXRLOEtibWVLSk9CdkFTZitRd0hrcCtGMkJBWHpSaVRsNForZ05GUVVMUHpzUVdwbUtsejZmSVdoYzdrc2dwVGtNSzZBYVRid1dZVGZtcEtuUXcvS0ptLzZyYm9MRFdZeUtGcFFjU3R1NjdSWithUnZRejY4RXYyZ2E1SnNYbGNPSjNnUC9sRTVXQzFTMHJqZmFiemRNT0dQOHFaUWhYazR3Qk9ndEZCYWlzRG5ialY1cGNJcmpSUGxob0N4dktnQy8yOTBuWjkvRExCSDNUYkhrOHh3SFhlQkFuQWp5QXFPWmlqOTJ1a3NBdjdaTHE0TU9EY25Rc2hWSU5Yd3NZc2hHMXBRcU9Md01lcnROYVk1V3RydWJNUmt1NDREdzdSIgogICAgICAgICAgICBdLAogICAgICAgICAgICAiaXNzdWVyIjogImh0dHBzOi8vbG9naW4ubWljcm9zb2Z0b25saW5lLmNvbS9lYmMzM2JiMC00Yzk1LTQ2NzMtOWI2NC02MWJmYTMyY2YyOTcvdjIuMCIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgImt0eSI6ICJSU0EiLAogICAgICAgICAgICAidXNlIjogInNpZyIsCiAgICAgICAgICAgICJraWQiOiAibDNzUS01MGNDSDR4QlZaTEhUR3duU1I3NjgwIiwKICAgICAgICAgICAgIng1dCI6ICJsM3NRLTUwY0NINHhCVlpMSFRHd25TUjc2ODAiLAogICAgICAgICAgICAibiI6ICJzZnNYTVhXdU8tZG5pTGFJRUxhM1B5cXo5WV9yV2ZmX0FWckNBbkZTZFBIYThfX1Bta2J0X3lxLTZaM3UxbzRnalJwS1ducmp4SWg4ekRuMVoxUlMyNm5rS2NOZzV4Zld4UjJLOENQYlNiWThnTXJwXzRwWm43dGdyRW1vTE1rd2ZnWWFWQy00TWlGRW8xUDJnZDltQ2RnSUlDYU5lWWtHMWJJUFRuYXFxdVRNNUtmVDk3MU1wdU9WT2RNMXlzaWVqZGNORHZFYjd2Mjg0UFlaa3cyaW13cWlCWTNGUjBzVkc3amdLVW90RnZoZDdUUjVXc0EyMEdTXzZaSWtVVWxMVWJHX3JYV0dsMFlqWkxTX1VmNHE4SGJvN3UtN01hRm44QjY5RjZZYUZkRGxYbV9BMFNwZWRWRldRRkd6TXNwNDNfNnZFempmckZESlZBWWt3YjZ4VVEiLAogICAgICAgICAgICAiZSI6ICJBUUFCIiwKICAgICAgICAgICAgIng1YyI6IFsKICAgICAgICAgICAgICAgICJNSUlEQlRDQ0FlMmdBd0lCQWdJUVdQQjFvZk9wQTdGRmxPQms1aVBhTlRBTkJna3Foa2lHOXcwQkFRc0ZBREF0TVNzd0tRWURWUVFERXlKaFkyTnZkVzUwY3k1aFkyTmxjM05qYjI1MGNtOXNMbmRwYm1SdmQzTXVibVYwTUI0WERUSXhNREl3TnpFM01EQXpPVm9YRFRJMk1ESXdOakUzTURBek9Wb3dMVEVyTUNrR0ExVUVBeE1pWVdOamIzVnVkSE11WVdOalpYTnpZMjl1ZEhKdmJDNTNhVzVrYjNkekxtNWxkRENDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFMSDdGekYxcmp2blo0aTJpQkMydHo4cXMvV1A2MW4zL3dGYXdnSnhVblR4MnZQL3o1cEc3ZjhxdnVtZDd0YU9JSTBhU2xwNjQ4U0lmTXc1OVdkVVV0dXA1Q25EWU9jWDFzVWRpdkFqMjBtMlBJREs2ZitLV1orN1lLeEpxQ3pKTUg0R0dsUXZ1REloUktOVDlvSGZaZ25ZQ0NBbWpYbUpCdFd5RDA1MnFxcmt6T1NuMC9lOVRLYmpsVG5UTmNySW5vM1hEUTd4Rys3OXZPRDJHWk1Ob3BzS29nV054VWRMRlJ1NDRDbEtMUmI0WGUwMGVWckFOdEJrdittU0pGRkpTMUd4djYxMWhwZEdJMlMwdjFIK0t2QjI2Tzd2dXpHaFovQWV2UmVtR2hYUTVWNXZ3TkVxWG5WUlZrQlJzekxLZU4vK3J4TTQzNnhReVZRR0pNRytzVkVDQXdFQUFhTWhNQjh3SFFZRFZSME9CQllFRkxsUkJTeHhnbU5QT2JDRnJsK2hTc2JjdlJrY01BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQitVUUZUTnM2QlVZM0FJR2tTMlpSdVpnSnNORXIvWkVNNGFDczJkb21kMk9xajcrNWlXc25QaDVDdWdGbkk0bmQrWkxnS1ZIU0Q2YWNRMjd3ZStlTlk2Z3hmcFFDWTFmaU4vdUtPT3NBMElmOEliUGRCRWh0UGVyUmdQSkZYTEhhWVZxRDhVWURvNUtOQ2NvQjRLaDhudkNXUkdQVVVIUFJxcDdBbkFjVnJjYmlYQS9ibU1DbkZXdU5OYWhjYUFLaUpUeFlsS0RhRElpUE4zNXlFQ1liRGowUEJXSlV4b2Jydmo1STI3NWpiaWtrcDhRU0xZblNVL3Y3ZE1EVWJ4U0xmWjd6c1R1YUYyUXgrTDYyUHNZVHdMeklGWDNNOEVNU1E2aDY4VHVwRlRpNW4wTTJ5SVhRZ29Sb05FRFdOSlovYVpNWS9ncVQwMkdRR0JXcmgrL3ZKIgogICAgICAgICAgICBdLAogICAgICAgICAgICAiaXNzdWVyIjogImh0dHBzOi8vbG9naW4ubWljcm9zb2Z0b25saW5lLmNvbS9lYmMzM2JiMC00Yzk1LTQ2NzMtOWI2NC02MWJmYTMyY2YyOTcvdjIuMCIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgImt0eSI6ICJSU0EiLAogICAgICAgICAgICAidXNlIjogInNpZyIsCiAgICAgICAgICAgICJraWQiOiAiTXI1LUFVaWJmQmlpN05kMWpCZWJheGJvWFcwIiwKICAgICAgICAgICAgIng1dCI6ICJNcjUtQVVpYmZCaWk3TmQxakJlYmF4Ym9YVzAiLAogICAgICAgICAgICAibiI6ICJ5cjN2MXVFVHJGZlQxN3p2T2l5MDF3OG5PLTF0NjdjbWlaTFp4cTJJU0RkdGU5ZHctSXhDUjdsUFYyd2V6Y3pJUmdjV21ZZ0Zuc2syajZtMTBINHRLemNxWk0wSkpfTmlnWTI5cEZpbXhsTDdfcVhNQjFQb3JGSmRsQUt2cDVTZ2pTVHdMclhqa3IxQXFXd2JwekcyeVpVTk4zR0U4R3ZtVGVvNHl3ZVFiTkNkLXlPX1pwb3p4MEozNHdIQkVNdWF3LVpmQ1VrN21kS0tzZy1FY0U0WnYwWGdsOXdQMk1wS1B4MFY4Z0xhenhlNlVROVNoek51cnVTT25jcExZSk5fb1E0YUtmNXB0T3AxcnNmRFkySUs5ZnJ0bVJUS09kUS1NRW1TZGpHTF84OElRY3ZDczdqcVZ6NTNYS29YUmxYQjh0TUlHT2NnLUlDZXI2eXhlMml0SVEiLAogICAgICAgICAgICAiZSI6ICJBUUFCIiwKICAgICAgICAgICAgIng1YyI6IFsKICAgICAgICAgICAgICAgICJNSUlEQlRDQ0FlMmdBd0lCQWdJUWZmOHlyRk8zQ0lOUEhVVFQ3NnRVc1RBTkJna3Foa2lHOXcwQkFRc0ZBREF0TVNzd0tRWURWUVFERXlKaFkyTnZkVzUwY3k1aFkyTmxjM05qYjI1MGNtOXNMbmRwYm1SdmQzTXVibVYwTUI0WERUSXhNVEF5TkRFM05EVTFObG9YRFRJMk1UQXlOREUzTkRVMU5sb3dMVEVyTUNrR0ExVUVBeE1pWVdOamIzVnVkSE11WVdOalpYTnpZMjl1ZEhKdmJDNTNhVzVrYjNkekxtNWxkRENDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFNcTk3OWJoRTZ4WDA5ZTg3em9zdE5jUEp6dnRiZXUzSm9tUzJjYXRpRWczYlh2WGNQaU1Ra2U1VDFkc0hzM015RVlIRnBtSUJaN0pObytwdGRCK0xTczNLbVROQ1NmellvR052YVJZcHNaUysvNmx6QWRUNkt4U1haUUNyNmVVb0kwazhDNjE0NUs5UUtsc0c2Y3h0c21WRFRkeGhQQnI1azNxT01zSGtHelFuZnNqdjJhYU04ZENkK01Cd1JETG1zUG1Yd2xKTzVuU2lySVBoSEJPR2I5RjRKZmNEOWpLU2o4ZEZmSUMyczhYdWxFUFVvY3picTdranAzS1MyQ1RmNkVPR2luK2FiVHFkYTdIdzJOaUN2WDY3WmtVeWpuVVBqQkprbll4aS8vUENFSEx3ck80NmxjK2QxeXFGMFpWd2ZMVENCam5JUGlBbnErc3NYdG9yU0VDQXdFQUFhTWhNQjh3SFFZRFZSME9CQllFRkRpWkc2czVkOVJ2b3JwcWJWZFMyL01EOFpLaE1BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQVFBUHVxcUtqMkFnZkM5YXl4K3FVdTB2anpLWWRaNlQrM3NzSkRPR3dCMWNMTVhNVFVWZ0Z3ajhic1gxYWhEVUpkektwV3ROajdibm8rVWc4NUl5VTdrODlVMFlncjU1eldVNWg0d25uUnJDdTlRS3Z1ZFVQbmJpWG9WdUhQd2NLOHcxZmRYWlFCNVFxL2tLemhOR1k1N2NHMWJ3ajNSL2FJZENwK0JqZ0ZwcE9LakpwSzdGS1M4RzJ2NzBlSWlDTE1hcEs5bExFZVFPeEl2emN0VHNYeTlFWjd3dGFJaVlreTRaU2l0dXBoVG9KVWtha0hhUTZldmJuODJsVGc2V1p6MXRtU21ZblBxUmRBZmY3YWlRMVN3OUhwdXpsWlkvcGlUVnF2ZDZBZktacXl4dS9GaEVORTBPZHYvMGhsSHpJMTVqS1FXTDFMamMwTm0zeTFza3V0IgogICAgICAgICAgICBdLAogICAgICAgICAgICAiaXNzdWVyIjogImh0dHBzOi8vbG9naW4ubWljcm9zb2Z0b25saW5lLmNvbS9lYmMzM2JiMC00Yzk1LTQ2NzMtOWI2NC02MWJmYTMyY2YyOTcvdjIuMCIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgImt0eSI6ICJSU0EiLAogICAgICAgICAgICAidXNlIjogInNpZyIsCiAgICAgICAgICAgICJraWQiOiAialMxWG8xT1dEal81MnZid0dOZ3ZRTzJWek1jIiwKICAgICAgICAgICAgIng1dCI6ICJqUzFYbzFPV0RqXzUydmJ3R05ndlFPMlZ6TWMiLAogICAgICAgICAgICAibiI6ICJzcHZRY1hXcVlyTWN2Y3FRbWZTTVluYlVDOFUwM1ljdG5YeUxJQmUxNDhPemhCcmdkQU9tUGZNZkppX3RVVzhMOXN2VkdwazVxRzZkTjBuNjY5Y1JIS3FVNTJHbkcwdGx5WVhtekZDMWh6SFZnUXo5ZWh2ZTR0bEo3dXc5MzZYSVVPQU94eDNYMjB6ZHB4N2dtNHpIeDRqMlpCbFhza0FqNlUzYWRwSFFOdXdVRTZrbW5nSldSLWRlV2xFaWdNcFJzdlVWUTJPNWgwLVJTcThXcl94N3VkM0s2R1R0cnpBUmFtejl1azJJWGF0S1lkbmo1SnJrMmpMWTZuV3QtR3R4bEFfbDlYd0lyT2w2U3FhX3BPR0lwUzAxSktkeEt2cEJDOVZkUzhvWEItN1A1cUxrc212N3RxLVNiYmlPZWMwY3ZVN1dQN3ZVUnYxMDRWNEZpSV9xb1EiLAogICAgICAgICAgICAiZSI6ICJBUUFCIiwKICAgICAgICAgICAgIng1YyI6IFsKICAgICAgICAgICAgICAgICJNSUlEQlRDQ0FlMmdBd0lCQWdJUUhzZXRQK2k4aTZWSUFtam1mVkd2NmpBTkJna3Foa2lHOXcwQkFRc0ZBREF0TVNzd0tRWURWUVFERXlKaFkyTnZkVzUwY3k1aFkyTmxjM05qYjI1MGNtOXNMbmRwYm1SdmQzTXVibVYwTUI0WERUSXlNREV6TURJek1EWXhORm9YRFRJM01ERXpNREl6TURZeE5Gb3dMVEVyTUNrR0ExVUVBeE1pWVdOamIzVnVkSE11WVdOalpYTnpZMjl1ZEhKdmJDNTNhVzVrYjNkekxtNWxkRENDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFMS2IwSEYxcW1LekhMM0trSm4wakdKMjFBdkZOTjJITFoxOGl5QVh0ZVBEczRRYTRIUURwajN6SHlZdjdWRnZDL2JMMVJxWk9haHVuVGRKK3V2WEVSeXFsT2RocHh0TFpjbUY1c3hRdFljeDFZRU0vWG9iM3VMWlNlN3NQZCtseUZEZ0RzY2QxOXRNM2FjZTRKdU14OGVJOW1RWlY3SkFJK2xOMm5hUjBEYnNGQk9wSnA0Q1ZrZm5YbHBSSW9ES1ViTDFGVU5qdVlkUGtVcXZGcS84ZTduZHl1aGs3YTh3RVdwcy9icE5pRjJyU21IWjQrU2E1Tm95Mk9wMXJmaHJjWlFQNWZWOENLenBla3FtdjZUaGlLVXROU1NuY1NyNlFRdlZYVXZLRndmdXorYWk1TEpyKzdhdmttMjRqbm5OSEwxTzFqKzcxRWI5ZE9GZUJZaVA2cUVDQXdFQUFhTWhNQjh3SFFZRFZSME9CQllFRkd6VkZqQWJZcFUvMmVuNHJ5NExNTFVISjNHak1BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQlUwWWROVmZkQnl2cHdzUGZ3TmREOG0xUExlZUNLbUxIUW5XUkk1NjAweUVIdW9Vdm9BSmQ1ZHdlMVpVMWJISFJSS1dON0FrdFV6b2ZQM3lGNjF4dGl6aEVieVBqSEsxdG5SK2lQRXZpV3hWdkszN0h0ZkVQenVoMVZxcDA4YnFZMTVNY1lVdGY3N2wySFhUcGFrK1VXWVJZSkJpKysydW1JREtZNVVNcVUrTEVabnZhWHliTFVLTjN4RzRpeTJxMUFiOHN5R0ZhVVA3SjNuQ3RWclI3aXAzOUJudlNUVFpaTm8vT0M3ZllYSjJYNHNOMS8yWmhSNUV0bkFnd2kyUnZsWmwwYVdQcmN6QXJVQ3hEQkNic0tQTC9VcC9rSUQxaXIxVk80TFQwOXJ5ZnYybngzeTZsMFl2dUw3ZVB6NG5HWUNXSGNiTVZjVXJRVVhxdVozWHRJIgogICAgICAgICAgICBdLAogICAgICAgICAgICAiaXNzdWVyIjogImh0dHBzOi8vbG9naW4ubWljcm9zb2Z0b25saW5lLmNvbS9lYmMzM2JiMC00Yzk1LTQ2NzMtOWI2NC02MWJmYTMyY2YyOTcvdjIuMCIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgImt0eSI6ICJSU0EiLAogICAgICAgICAgICAidXNlIjogInNpZyIsCiAgICAgICAgICAgICJraWQiOiAiRHFVdThnZi1uQWdjeWpQMy1TdXBsTkFYQW5jIiwKICAgICAgICAgICAgIng1dCI6ICJEcVV1OGdmLW5BZ2N5alAzLVN1cGxOQVhBbmMiLAogICAgICAgICAgICAibiI6ICIxbjctbldTTGV1V1F6QlJsWVNiUzhSanZXdmtRZUQ3UUw5Zk9XYUdYYlc3M1ZOR0gwWWlwWmlzUENsRnY2R3p3ZldFQ1RXUXAxOVdGZV9sQVNrYTUtS0VXa1FWekNiRU1hYWFmT0lzN2hDNjFQNWNHZ3c3ZGh1VzRzN2Y2WllHWkV6UTRGNXJIRS1ZTlJidkQ1MXFpclBOektIazNuamkxd3JoMFl0YlBQSWYtLU5iSTk4YkN3TExoOWF2ZWRPbXFFU3pXT0dFQ0VNWHY4TFNNLUI5U0tnXzRRdUJ0eUJ3d0lha1R1cW84NHN3VEJNNXc4UGRocFdaWkR0UGdIODdXei1fV2pXdms5OUFqWGw3bDhwV1BRSmlLTnVqdF9jazNOREZwemFMRXBwb2RoVXNJRDBwdFJBMDA4ZUNVNmw4VC11eDE5d1ptYl95Qm5IY1YzcEZXaFEiLAogICAgICAgICAgICAiZSI6ICJBUUFCIiwKICAgICAgICAgICAgIng1YyI6IFsKICAgICAgICAgICAgICAgICJNSUlDOFRDQ0FkbWdBd0lCQWdJUVlWay90SjFlNHBoSVN2VnJBQUxOS1RBTkJna3Foa2lHOXcwQkFRc0ZBREFqTVNFd0h3WURWUVFERXhoc2IyZHBiaTV0YVdOeWIzTnZablJ2Ym14cGJtVXVkWE13SGhjTk1qQXhNakl4TURBd01EQXdXaGNOTWpVeE1qSXhNREF3TURBd1dqQWpNU0V3SHdZRFZRUURFeGhzYjJkcGJpNXRhV055YjNOdlpuUnZibXhwYm1VdWRYTXdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFEV2Z2NmRaSXQ2NVpETUZHVmhKdEx4R085YStSQjRQdEF2MTg1Wm9aZHRidmRVMFlmUmlLbG1LdzhLVVcvb2JQQjlZUUpOWkNuWDFZVjcrVUJLUnJuNG9SYVJCWE1Kc1F4cHBwODRpenVFTHJVL2x3YUREdDJHNWJpenQvcGxnWmtUTkRnWG1zY1Q1ZzFGdThQbldxS3M4M01vZVRlZU9MWEN1SFJpMXM4OGgvNzQxc2ozeHNMQXN1SDFxOTUwNmFvUkxOWTRZUUlReGUvd3RJejRIMUlxRC9oQzRHM0lIREFocVJPNnFqeml6Qk1Fem5EdzkyR2xabGtPMCtBZnp0YlA3OWFOYStUMzBDTmVYdVh5bFk5QW1JbzI2TzM5eVRjME1Xbk5vc1NtbWgyRlN3Z1BTbTFFRFRUeDRKVHFYeFA2N0hYM0JtWnYvSUdjZHhYZWtWYUZBZ01CQUFHaklUQWZNQjBHQTFVZERnUVdCQlEyci8vbGdUUGNLdWdoRGt6bUN0Umx3K1A5U3pBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQXNGZFJ5Y3pOV2gvcXBZdmNJWmJEdldZemxybUZaYzZibGNVem5zOXpmN3NVV3RRWnJaUHU1RGJldFYyR3IycjNxdE1ES1hDVWFSK3Bxb3kzSTJ6eFRYM3g4YlROaFpEOVlBZ0FGbFRMTlN5ZFRhSzVSSHlCLzVrcjZCN1pKZU5JazNQUlZoUkd0NnliQ0pTalYvVllWa0xSNWZkTFArNUdodkJFU29iQVIvZDBudHJpVHpwNy90TE1iL29YeDd3NUh1MW0zSThycE1vY29YZkgyU0gxR0xtTVhqNk14MWR0d0NEWU02YnNiM2ZoV1J6OU85T01SNlFOaVRucThxOXduMVF6QkFuUmNzd1l6VDFMS0tCUE5GU2FzQ3ZMWU9DUE9aQ0wrVzhOOGpxYTlaUllOWUtXWHptaVNwdGdCRU0yNHQzbTVGVVd6V3FvTHU5cEljbmtQUT09IgogICAgICAgICAgICBdLAogICAgICAgICAgICAiaXNzdWVyIjogImh0dHBzOi8vbG9naW4ubWljcm9zb2Z0b25saW5lLmNvbS9lYmMzM2JiMC00Yzk1LTQ2NzMtOWI2NC02MWJmYTMyY2YyOTcvdjIuMCIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgImt0eSI6ICJSU0EiLAogICAgICAgICAgICAidXNlIjogInNpZyIsCiAgICAgICAgICAgICJraWQiOiAiT3paNURibWNzbzlRenQyTW9kR21paGczMEJvIiwKICAgICAgICAgICAgIng1dCI6ICJPelo1RGJtY3NvOVF6dDJNb2RHbWloZzMwQm8iLAogICAgICAgICAgICAibiI6ICIwMXJlOWEyQlVUdE50ZEZ6TE5JLVFFSFc4WGhEaURNRGJHTWt4SFJJWVhINDF6QmNjc1h3SDl2TWkwSHV4WEhwWE96d3RVWUt3bDkzWlIzN3RwNmxwdndsVTFIZVBObVpwSjlELVhBdlU3M3gwM1lLb1pFZGFGQjM5VnNWeUxpaDNmdVB2NkRQRTJxVC1UTkUzWDVZZElXT0dGcmNNa2NYTHNqTy1CQ3E0cWNTZEJIMmxCZ0VRVXVENm5xcmVMWnNnLWdQelNEaGpWU2NJVVpHaUQ4TTJzS3hBRGlJSG81S2xhWkl5dTMydDhKa2F2UDlqTTdJdFNBanppZzFXMnl2VlF6VVFaQS14WnFKbzJqeEIzZ19meWdkUFVISzZVTi1fY3FrcmZ4bjItVldIMXdNaGxtOTBTcHhUTUQ0SG9ZT1ZpejFnZ0g4R0NYMmFCaVg1T3pRNlEiLAogICAgICAgICAgICAiZSI6ICJBUUFCIiwKICAgICAgICAgICAgIng1YyI6IFsKICAgICAgICAgICAgICAgICJNSUlDOFRDQ0FkbWdBd0lCQWdJUVFyWFBYSWxVRTRKTVRNa1ZqKzAyWVRBTkJna3Foa2lHOXcwQkFRc0ZBREFqTVNFd0h3WURWUVFERXhoc2IyZHBiaTV0YVdOeWIzTnZablJ2Ym14cGJtVXVkWE13SGhjTk1qRXdNekV3TURBd01EQXdXaGNOTWpZd016RXdNREF3TURBd1dqQWpNU0V3SHdZRFZRUURFeGhzYjJkcGJpNXRhV055YjNOdlpuUnZibXhwYm1VdWRYTXdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFEVFd0NzFyWUZSTzAyMTBYTXMwajVBUWRieGVFT0lNd05zWXlURWRFaGhjZmpYTUZ4eXhmQWYyOHlMUWU3RmNlbGM3UEMxUmdyQ1gzZGxIZnUybnFXbS9DVlRVZDQ4Mlpta24wUDVjQzlUdmZIVGRncWhrUjFvVUhmMVd4WEl1S0hkKzQrL29NOFRhcFA1TTBUZGZsaDBoWTRZV3R3eVJ4Y3V5TTc0RUtyaXB4SjBFZmFVR0FSQlM0UHFlcXQ0dG15RDZBL05JT0dOVkp3aFJrYUlQd3phd3JFQU9JZ2Vqa3FWcGtqSzdmYTN3bVJxOC8yTXpzaTFJQ1BPS0RWYmJLOVZETlJCa0Q3Rm1vbWphUEVIZUQ5L0tCMDlRY3JwUTM3OXlxU3QvR2ZiNVZZZlhBeUdXYjNSS25GTXdQZ2VoZzVXTFBXQ0Fmd1lKZlpvR0pmazdORHBBZ01CQUFHaklUQWZNQjBHQTFVZERnUVdCQlRFQ2pCUkFORFBMR3JuMXA3cXR3c3d0QlU3SnpBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQXExSWI0RVJ2WEc1a2lWbWhmTE9wdW4yRWxWT0xZK1hrdlZseVZqcTM1clptU0lHeGdmRmMwOFFPUUZWbXJXUVlybHNzMExiYm9IMGNJS2lENityVm9aVE1XeEdFaWNPY0dORnpya2NHMHVsdTBjZ2hLWUlEM0dLRFRmdFlLRVBrdnUydlFtdWVxYTR0MnRUM1BsWUY3RmkyZGJvUjVZOTZVZ3M4enFOd2RCTVJtNjc3Ti90SkJrNTNDc09mOU5uQmF4WjFFR0FybUVISEliODB2T0RTdHYzNXVlTHJmTVJ0Q0YvSGNna0d4eTJVOGthQ3pZbW16SGg0ellEa2VDd00zQ3EyYkdrRytFZmU5aEZZZkRIdzEzRHpUUitoOXBQcUZGaUF4blozb2ZUOTZOclpIZFlqd2JmbU04Y3czbGRnMHhRekdjd1pqdHlZbXdKNnNEZFJ2UT09IgogICAgICAgICAgICBdLAogICAgICAgICAgICAiaXNzdWVyIjogImh0dHBzOi8vbG9naW4ubWljcm9zb2Z0b25saW5lLmNvbS9lYmMzM2JiMC00Yzk1LTQ2NzMtOWI2NC02MWJmYTMyY2YyOTcvdjIuMCIKICAgICAgICB9CiAgICBdCn0K"
	client_ids=["c2856e29-d344-49a9-88ed-de3a01d99753"]
	description="this is a description1"
}
`
