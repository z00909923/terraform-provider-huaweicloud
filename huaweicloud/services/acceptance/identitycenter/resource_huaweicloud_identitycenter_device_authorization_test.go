package identitycenter

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccIdentityCenterDeviceAuthorization_basic(t *testing.T) {
	rName := "huaweicloud_identitycenter_device_authorization.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckMultiAccount(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testIdentityCenterDeviceAuthorization_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rName, "device_code"),
					resource.TestCheckResourceAttrSet(rName, "verification_uri_complete"),
				),
			},
		},
	})
}

func testIdentityCenterDeviceAuthorization_basic() string {
	return fmt.Sprintf(`
%s
resource "huaweicloud_identitycenter_device_authorization" "test"{
  client_id    = huaweicloud_identitycenter_client.test.client_id
  client_secret   = huaweicloud_identitycenter_client.test.client_secret
  start_url  = "https://idcenter.huaweicloud.com/d-47f958c067/portal"
}
`, testIdentityCenterClient_basic())
}
