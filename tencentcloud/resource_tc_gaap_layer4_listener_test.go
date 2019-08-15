package tencentcloud

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccTencentCloudGaapLayer4Listener_basic(t *testing.T) {
	id := new(string)
	proxyId := new(string)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGaapLayer4ListenerDestroy(id, proxyId, "TCP"),
		Steps: []resource.TestStep{
			{
				Config: testAccGaapLayer4ListenerBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGaapLayer4ListenerExists("tencentcloud_gaap_layer4_listener.foo", id, proxyId, "TCP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "protocol", "TCP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "name", "ci-test-gaap-4-listener"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "port", "80"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_type", "IP"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "proxy_id"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "health_check", "true"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "delay_loop", "5"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "connect_timeout", "3"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_bind_set.#", "2"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "status"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "create_time"),
				),
			},
		},
	})
}

func TestAccTencentCloudGaapLayer4Listener_TcpDomain(t *testing.T) {
	id := new(string)
	proxyId := new(string)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGaapLayer4ListenerDestroy(id, proxyId, "TCP"),
		Steps: []resource.TestStep{
			{
				Config: testAccGaapLayer4ListenerTcpDomain,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGaapLayer4ListenerExists("tencentcloud_gaap_layer4_listener.foo", id, proxyId, "TCP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "protocol", "TCP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "name", "ci-test-gaap-4-listener"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "port", "80"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_type", "DOMAIN"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "proxy_id"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "health_check", "true"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "delay_loop", "5"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "connect_timeout", "3"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_bind_set.#", "2"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "status"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "create_time"),
				),
			},
		},
	})
}

func TestAccTencentCloudGaapLayer4Listener_update(t *testing.T) {
	id := new(string)
	proxyId := new(string)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGaapLayer4ListenerDestroy(id, proxyId, "TCP"),
		Steps: []resource.TestStep{
			{
				Config: testAccGaapLayer4ListenerBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGaapLayer4ListenerExists("tencentcloud_gaap_layer4_listener.foo", id, proxyId, "TCP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "protocol", "TCP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "name", "ci-test-gaap-4-listener"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "port", "80"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "scheduler", "rr"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_type", "IP"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "proxy_id"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "health_check", "true"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "delay_loop", "5"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "connect_timeout", "3"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_bind_set.#", "2"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "status"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "create_time"),
				),
			},
			{
				Config: testAccGaapLayer4ListenerUpdateNameAndHealthConfigAndScheduler,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGaapLayer4ListenerExists("tencentcloud_gaap_layer4_listener.foo", id, proxyId, "TCP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "name", "ci-test-gaap-4-listener-new"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "scheduler", "wrr"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "delay_loop", "11"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "connect_timeout", "10"),
				),
			},
			{
				Config: testAccGaapLayer4ListenerUpdateNoHealthCheck,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGaapLayer4ListenerExists("tencentcloud_gaap_layer4_listener.foo", id, proxyId, "TCP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "health_check", "false"),
					resource.TestCheckNoResourceAttr("tencentcloud_gaap_layer4_listener.foo", "delay_loop"),
					resource.TestCheckNoResourceAttr("tencentcloud_gaap_layer4_listener.foo", "connect_timeout"),
				),
			},
		},
	})
}

func TestAccTencentCloudGaapLayer4Listener_updateRealserverSet(t *testing.T) {
	id := new(string)
	proxyId := new(string)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGaapLayer4ListenerDestroy(id, proxyId, "TCP"),
		Steps: []resource.TestStep{
			{
				Config: testAccGaapLayer4ListenerBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGaapLayer4ListenerExists("tencentcloud_gaap_layer4_listener.foo", id, proxyId, "TCP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "protocol", "TCP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "name", "ci-test-gaap-4-listener"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "port", "80"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "scheduler", "rr"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_type", "IP"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "proxy_id"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "health_check", "true"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "delay_loop", "5"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "connect_timeout", "3"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_bind_set.#", "2"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "status"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "create_time"),
				),
			},
			{
				Config: testAccGaapLayer4ListenerTcpUpdateRealserverSet,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGaapLayer4ListenerExists("tencentcloud_gaap_layer4_listener.foo", id, proxyId, "TCP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_bind_set.#", "1"),
				),
			},
		},
	})
}

func TestAccTencentCloudGaapLayer4Listener_udp(t *testing.T) {
	id := new(string)
	proxyId := new(string)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGaapLayer4ListenerDestroy(id, proxyId, "UDP"),
		Steps: []resource.TestStep{
			{
				Config: testAccGaapLayer4ListenerUdp,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGaapLayer4ListenerExists("tencentcloud_gaap_layer4_listener.foo", id, proxyId, "UDP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "protocol", "UDP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "name", "ci-test-gaap-4-udp-listener"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "port", "80"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "scheduler", "rr"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_type", "IP"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "proxy_id"),
					resource.TestCheckNoResourceAttr("tencentcloud_gaap_layer4_listener.foo", "health_check"),
					resource.TestCheckNoResourceAttr("tencentcloud_gaap_layer4_listener.foo", "delay_loop"),
					resource.TestCheckNoResourceAttr("tencentcloud_gaap_layer4_listener.foo", "connect_timeout"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_bind_set.#", "2"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "status"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "create_time"),
				),
			},
		},
	})
}

func TestAccTencentCloudGaapLayer4Listener_udpDomain(t *testing.T) {
	id := new(string)
	proxyId := new(string)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGaapLayer4ListenerDestroy(id, proxyId, "UDP"),
		Steps: []resource.TestStep{
			{
				Config: testAccGaapLayer4ListenerUdpDomain,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGaapLayer4ListenerExists("tencentcloud_gaap_layer4_listener.foo", id, proxyId, "UDP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "protocol", "UDP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "name", "ci-test-gaap-4-udp-listener"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "port", "80"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "scheduler", "rr"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_type", "DOMAIN"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "proxy_id"),
					resource.TestCheckNoResourceAttr("tencentcloud_gaap_layer4_listener.foo", "health_check"),
					resource.TestCheckNoResourceAttr("tencentcloud_gaap_layer4_listener.foo", "delay_loop"),
					resource.TestCheckNoResourceAttr("tencentcloud_gaap_layer4_listener.foo", "connect_timeout"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_bind_set.#", "2"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "status"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "create_time"),
				),
			},
		},
	})
}

func TestAccTencentCloudGaapLayer4Listener_udpUpdate(t *testing.T) {
	id := new(string)
	proxyId := new(string)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGaapLayer4ListenerDestroy(id, proxyId, "UDP"),
		Steps: []resource.TestStep{
			{
				Config: testAccGaapLayer4ListenerUdp,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGaapLayer4ListenerExists("tencentcloud_gaap_layer4_listener.foo", id, proxyId, "UDP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "protocol", "UDP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "name", "ci-test-gaap-4-udp-listener"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "port", "80"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "scheduler", "rr"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_type", "IP"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "proxy_id"),
					resource.TestCheckNoResourceAttr("tencentcloud_gaap_layer4_listener.foo", "health_check"),
					resource.TestCheckNoResourceAttr("tencentcloud_gaap_layer4_listener.foo", "delay_loop"),
					resource.TestCheckNoResourceAttr("tencentcloud_gaap_layer4_listener.foo", "connect_timeout"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "realserver_bind_set.#", "2"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "status"),
					resource.TestCheckResourceAttrSet("tencentcloud_gaap_layer4_listener.foo", "create_time"),
				),
			},
			{
				Config: testAccGaapLayer4ListenerUdpUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGaapLayer4ListenerExists("tencentcloud_gaap_layer4_listener.foo", id, proxyId, "UDP"),
					resource.TestCheckResourceAttr("tencentcloud_gaap_layer4_listener.foo", "name", "ci-test-gaap-4-udpListener-new"),
				),
			},
		},
	})
}

func testAccCheckGaapLayer4ListenerExists(n string, id, proxyId *string, protocol string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("no listener ID is set")
		}

		attrProxyId := rs.Primary.Attributes["proxy_id"]
		if attrProxyId == "" {
			return errors.New("no proxy ID is set")
		}

		service := GaapService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

		switch protocol {
		case "TCP":
			listeners, err := service.DescribeTCPListener(context.TODO(), attrProxyId, &rs.Primary.ID, nil, nil)
			if err != nil {
				return err
			}

			for _, l := range listeners {
				if l.ListenerId == nil {
					return errors.New("listener id is nil")
				}
				if rs.Primary.ID == *l.ListenerId {
					*id = *l.ListenerId
					*proxyId = attrProxyId
					break
				}
			}

		case "UDP":
			listeners, err := service.DescribeUDPListener(context.TODO(), attrProxyId, &rs.Primary.ID, nil, nil)
			if err != nil {
				return err
			}

			for _, l := range listeners {
				if l.ListenerId == nil {
					return errors.New("listener id is nil")
				}
				if rs.Primary.ID == *l.ListenerId {
					*id = *l.ListenerId
					*proxyId = attrProxyId
					break
				}
			}
		}

		if id == nil {
			return fmt.Errorf("listener not found: %s", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckGaapLayer4ListenerDestroy(id, proxyId *string, protocol string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*TencentCloudClient).apiV3Conn
		service := GaapService{client: client}

		switch protocol {
		case "TCP":
			listeners, err := service.DescribeTCPListener(context.TODO(), *proxyId, id, nil, nil)
			if err != nil {
				return err
			}
			if len(listeners) > 0 {
				return errors.New("listener still exists")
			}

		case "UDP":
			listeners, err := service.DescribeUDPListener(context.TODO(), *proxyId, id, nil, nil)
			if err != nil {
				return err
			}
			if len(listeners) > 0 {
				return errors.New("listener still exists")
			}
		}

		return nil
	}
}

const testAccGaapLayer4ListenerBasic = `
resource tencentcloud_gaap_proxy "foo" {
  name              = "ci-test-gaap-proxy"
  bandwidth         = 10
  concurrent        = 2
  access_region     = "SouthChina"
  realserver_region = "NorthChina"
}

resource tencentcloud_gaap_realserver "foo" {
  ip   = "1.1.1.1"
  name = "ci-test-gaap-realserver"
}

resource tencentcloud_gaap_realserver "bar" {
  ip   = "119.29.29.29"
  name = "ci-test-gaap-realserver2"
}

resource tencentcloud_gaap_layer4_listener "foo" {
  protocol            = "TCP"
  name                = "ci-test-gaap-4-listener"
  port                = 80
  realserver_type     = "IP"
  proxy_id            = "${tencentcloud_gaap_proxy.foo.id}"
  health_check        = true
  delay_loop          = 5
  connect_timeout     = 3

  realserver_bind_set {
    id   = "${tencentcloud_gaap_realserver.foo.id}"
    ip   = "${tencentcloud_gaap_realserver.foo.ip}"
    port = 80
  }
  realserver_bind_set {
    id     = "${tencentcloud_gaap_realserver.bar.id}"
    ip     = "${tencentcloud_gaap_realserver.bar.ip}"
    port   = 80
  }
}
`

const testAccGaapLayer4ListenerTcpDomain = `
resource tencentcloud_gaap_proxy "foo" {
  name              = "ci-test-gaap-proxy"
  bandwidth         = 10
  concurrent        = 2
  access_region     = "SouthChina"
  realserver_region = "NorthChina"
}

resource tencentcloud_gaap_realserver "foo" {
  domain = "qq.com"
  name   = "ci-test-gaap-realserver"
}

resource tencentcloud_gaap_realserver "bar" {
  domain = "www.qq.com"
  name   = "ci-test-gaap-realserver2"
}

resource tencentcloud_gaap_layer4_listener "foo" {
  protocol            = "TCP"
  name                = "ci-test-gaap-4-listener"
  port                = 80
  realserver_type     = "DOMAIN"
  proxy_id            = "${tencentcloud_gaap_proxy.foo.id}"
  health_check        = true
  delay_loop          = 5
  connect_timeout     = 3

  realserver_bind_set {
    id   = "${tencentcloud_gaap_realserver.foo.id}"
    ip   = "${tencentcloud_gaap_realserver.foo.domain}"
    port = 80
  }
  realserver_bind_set {
    id     = "${tencentcloud_gaap_realserver.bar.id}"
    ip     = "${tencentcloud_gaap_realserver.bar.domain}"
    port   = 80
  }
}
`

const testAccGaapLayer4ListenerUpdateNameAndHealthConfigAndScheduler = `
resource tencentcloud_gaap_proxy "foo" {
  name              = "ci-test-gaap-proxy"
  bandwidth         = 10
  concurrent        = 2
  access_region     = "SouthChina"
  realserver_region = "NorthChina"
}

resource tencentcloud_gaap_realserver "foo" {
  ip   = "1.1.1.1"
  name = "ci-test-gaap-realserver"
}

resource tencentcloud_gaap_realserver "bar" {
  ip   = "119.29.29.29"
  name = "ci-test-gaap-realserver2"
}

resource tencentcloud_gaap_layer4_listener "foo" {
  protocol            = "TCP"
  name                = "ci-test-gaap-4-listener-new"
  port                = 80
  scheduler           = "wrr"
  realserver_type     = "IP"
  proxy_id            = "${tencentcloud_gaap_proxy.foo.id}"
  health_check        = true
  delay_loop          = 11
  connect_timeout     = 10

  realserver_bind_set {
    id   = "${tencentcloud_gaap_realserver.foo.id}"
    ip   = "${tencentcloud_gaap_realserver.foo.ip}"
    port = 80
  }
  realserver_bind_set {
    id     = "${tencentcloud_gaap_realserver.bar.id}"
    ip     = "${tencentcloud_gaap_realserver.bar.ip}"
    port   = 80
  }
}
`

const testAccGaapLayer4ListenerUpdateNoHealthCheck = `
resource tencentcloud_gaap_proxy "foo" {
  name              = "ci-test-gaap-proxy"
  bandwidth         = 10
  concurrent        = 2
  access_region     = "SouthChina"
  realserver_region = "NorthChina"
}

resource tencentcloud_gaap_realserver "foo" {
  ip   = "1.1.1.1"
  name = "ci-test-gaap-realserver"
}

resource tencentcloud_gaap_realserver "bar" {
  ip   = "119.29.29.29"
  name = "ci-test-gaap-realserver2"
}

resource tencentcloud_gaap_layer4_listener "foo" {
  protocol            = "TCP"
  name                = "ci-test-gaap-4-listener-new"
  port                = 80
  realserver_type     = "IP"
  proxy_id            = "${tencentcloud_gaap_proxy.foo.id}"
  health_check        = false

  realserver_bind_set {
    id   = "${tencentcloud_gaap_realserver.foo.id}"
    ip   = "${tencentcloud_gaap_realserver.foo.ip}"
    port = 80
  }
  realserver_bind_set {
    id     = "${tencentcloud_gaap_realserver.bar.id}"
    ip     = "${tencentcloud_gaap_realserver.bar.ip}"
    port   = 80
  }
}
`

const testAccGaapLayer4ListenerTcpUpdateRealserverSet = `
resource tencentcloud_gaap_proxy "foo" {
  name              = "ci-test-gaap-proxy"
  bandwidth         = 10
  concurrent        = 2
  access_region     = "SouthChina"
  realserver_region = "NorthChina"
}

resource tencentcloud_gaap_realserver "foo" {
  ip   = "1.1.1.1"
  name = "ci-test-gaap-realserver"
}

resource tencentcloud_gaap_realserver "bar" {
  ip   = "119.29.29.29"
  name = "ci-test-gaap-realserver2"
}

resource tencentcloud_gaap_layer4_listener "foo" {
  protocol            = "TCP"
  name                = "ci-test-gaap-4-listener"
  port                = 80
  realserver_type     = "IP"
  proxy_id            = "${tencentcloud_gaap_proxy.foo.id}"
  health_check        = true
  delay_loop          = 5
  connect_timeout     = 3

  realserver_bind_set {
    id   = "${tencentcloud_gaap_realserver.foo.id}"
    ip   = "${tencentcloud_gaap_realserver.foo.ip}"
    port = 80
  }
}
`

const testAccGaapLayer4ListenerUdp = `
resource tencentcloud_gaap_proxy "foo" {
  name              = "ci-test-gaap-proxy"
  bandwidth         = 10
  concurrent        = 2
  access_region     = "SouthChina"
  realserver_region = "NorthChina"
}

resource tencentcloud_gaap_realserver "foo" {
  ip   = "1.1.1.1"
  name = "ci-test-gaap-realserver"
}

resource tencentcloud_gaap_realserver "bar" {
  ip   = "119.29.29.29"
  name = "ci-test-gaap-realserver2"
}

resource tencentcloud_gaap_layer4_listener "foo" {
  protocol            = "UDP"
  name                = "ci-test-gaap-4-udp-listener"
  port                = 80
  realserver_type     = "IP"
  proxy_id            = "${tencentcloud_gaap_proxy.foo.id}"

  realserver_bind_set {
    id   = "${tencentcloud_gaap_realserver.foo.id}"
    ip   = "${tencentcloud_gaap_realserver.foo.ip}"
    port = 80
  }
  realserver_bind_set {
    id     = "${tencentcloud_gaap_realserver.bar.id}"
    ip     = "${tencentcloud_gaap_realserver.bar.ip}"
    port   = 80
  }
}
`

const testAccGaapLayer4ListenerUdpDomain = `
resource tencentcloud_gaap_proxy "foo" {
  name              = "ci-test-gaap-proxy"
  bandwidth         = 10
  concurrent        = 2
  access_region     = "SouthChina"
  realserver_region = "NorthChina"
}

resource tencentcloud_gaap_realserver "foo" {
  domain = "www.qq.com"
  name   = "ci-test-gaap-realserver"
}

resource tencentcloud_gaap_realserver "bar" {
  domain = "qq.com"
  name   = "ci-test-gaap-realserver2"
}

resource tencentcloud_gaap_layer4_listener "foo" {
  protocol            = "UDP"
  name                = "ci-test-gaap-4-udp-listener"
  port                = 80
  realserver_type     = "DOMAIN"
  proxy_id            = "${tencentcloud_gaap_proxy.foo.id}"

  realserver_bind_set {
    id   = "${tencentcloud_gaap_realserver.foo.id}"
    ip   = "${tencentcloud_gaap_realserver.foo.domain}"
    port = 80
  }
  realserver_bind_set {
    id     = "${tencentcloud_gaap_realserver.bar.id}"
    ip     = "${tencentcloud_gaap_realserver.bar.domain}"
    port   = 80
  }
}
`

const testAccGaapLayer4ListenerUdpUpdate = `
resource tencentcloud_gaap_proxy "foo" {
  name              = "ci-test-gaap-proxy"
  bandwidth         = 10
  concurrent        = 2
  access_region     = "SouthChina"
  realserver_region = "NorthChina"
}

resource tencentcloud_gaap_realserver "foo" {
  ip   = "1.1.1.1"
  name = "ci-test-gaap-realserver"
}

resource tencentcloud_gaap_realserver "bar" {
  ip   = "119.29.29.29"
  name = "ci-test-gaap-realserver2"
}

resource tencentcloud_gaap_layer4_listener "foo" {
  protocol            = "UDP"
  name                = "ci-test-gaap-4-udpListener-new"
  port                = 80
  realserver_type     = "IP"
  proxy_id            = "${tencentcloud_gaap_proxy.foo.id}"

  realserver_bind_set {
    id   = "${tencentcloud_gaap_realserver.foo.id}"
    ip   = "${tencentcloud_gaap_realserver.foo.ip}"
    port = 80
  }
  realserver_bind_set {
    id     = "${tencentcloud_gaap_realserver.bar.id}"
    ip     = "${tencentcloud_gaap_realserver.bar.ip}"
    port   = 80
  }
}
`
