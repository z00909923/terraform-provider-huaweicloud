package identitycenter

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccIdentityCenterDeviceToken_basic(t *testing.T) {
	rName := "huaweicloud_identitycenter_device_token.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckMultiAccount(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testIdentityCenterDeviceToken_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rName, "access_token"),
					resource.TestCheckResourceAttrSet(rName, "expires_in"),
				),
			},
		},
	})
}

func testIdentityCenterDeviceToken_basic() string {
	return fmt.Sprintf(`
%s
resource "huaweicloud_identitycenter_device_token" "test"{
	client_id      = huaweicloud_identitycenter_client.test.id
	client_secret   = huaweicloud_identitycenter_client.test.client_secret
    device_code	= "v6rk0nzerzggvc43uyejdo0r2w1ld1rgz3rg8hetl5fzg4q5elldl7xtlb9o1g1v"
	grant_type		= "urn:ietf:params:oauth:grant-type:device_code"
}
`, testIdentityCenterClient_basic())
}
