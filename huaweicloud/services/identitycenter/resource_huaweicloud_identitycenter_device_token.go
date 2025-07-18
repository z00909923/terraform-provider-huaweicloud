// ---------------------------------------------------------------
// *** AUTO GENERATED CODE ***
// @Product IdentityCenter
// ---------------------------------------------------------------

package identitycenter

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API IdentityStore POST /v1/clients
func ResourceIdentityCenterDeviceToken() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIdentityCenterDeviceTokenCreate,
		ReadContext:   resourceIdentityCenterDeviceTokenRead,
		DeleteContext: resourceIdentityCenterDeviceTokenDelete,
		Description:   "schema: Internal",
		Schema: map[string]*schema.Schema{
			"client_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"client_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"token_endpoint_auth_method": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scopes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"grant_types": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"response_types": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceIdentityCenterDeviceTokenCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// createIdentityCenterClient: create IdentityCenter client
	var (
		createIdentityCenterClientHttpUrl = "v1/clients"
		createIdentityCenterClientProduct = "identityoidc"
	)
	createIdentityCenterClientClient, err := cfg.NewServiceClient(createIdentityCenterClientProduct, region)
	if err != nil {
		return diag.Errorf("error creating Identity Center Client: %s", err)
	}

	createIdentityCenterClientPath := createIdentityCenterClientClient.Endpoint + createIdentityCenterClientHttpUrl
	createIdentityCenterClientOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}
	log.Println(createIdentityCenterClientOpt)
	createIdentityCenterClientOpt.JSONBody = utils.RemoveNil(buildCreateIdentityCenterDeviceTokenBodyParams(d))
	createIdentityCenterClientResp, err := createIdentityCenterClientClient.Request("POST",
		createIdentityCenterClientPath, &createIdentityCenterClientOpt)
	if err != nil {
		return diag.Errorf("error creating Identity Center Client: %s", err)
	}

	createIdentityCenterClientRespBody, err := utils.FlattenResponse(createIdentityCenterClientResp)
	if err != nil {
		return diag.FromErr(err)
	}

	deviceCode := utils.PathSearch("device_code", createIdentityCenterClientRespBody, "").(string)
	if deviceCode == "" {
		return diag.Errorf("unable to find the Identity Center device_code from the API response")
	}
	d.SetId(deviceCode)

	return nil
}

func buildCreateIdentityCenterDeviceTokenBodyParams(d *schema.ResourceData) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"client_id":     utils.ValueIgnoreEmpty(d.Get("client_id")),
		"client_secret": utils.ValueIgnoreEmpty(d.Get("client_secret")),
		"start_url":     utils.ValueIgnoreEmpty(d.Get("start_url")),
	}
	return bodyParams
}

func resourceIdentityCenterDeviceTokenRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceIdentityCenterDeviceTokenDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	errorMsg := "Deleting client is not supported."
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  errorMsg,
		},
	}
}
