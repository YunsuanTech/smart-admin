package Twitter

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	TwitterUsersApi
	ScoringRulesApi
	ScoreRecordApi
}

var (
	twitterUsersService = service.ServiceGroupApp.TwitterServiceGroup.TwitterUsersService
	scoringRulesService = service.ServiceGroupApp.TwitterServiceGroup.ScoringRulesService
	scoreRecordService  = service.ServiceGroupApp.TwitterServiceGroup.ScoreRecordService
)
