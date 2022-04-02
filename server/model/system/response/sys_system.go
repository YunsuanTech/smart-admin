package response

import "smart-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
