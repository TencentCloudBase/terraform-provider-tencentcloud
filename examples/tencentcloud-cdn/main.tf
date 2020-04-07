resource "tencentcloud_cdn_domain" "foo" {
  domain = "xxxx.com"
  service_type = "web"
  area = "mainland"

  origin {
    origin_type = "ip"
    origin_list = ["127.0.0.1"]
    origin_pull_protocol = "follow"
  }

  https_config {
    https_switch = "off"
    http2_switch = "off"
    ocsp_stapling_switch = "off"
    spdy_switch = "off"
    verify_client = "off"
  }

  tags = {
    hello = "world"
  }
}