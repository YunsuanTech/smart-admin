package Twitter

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	TwitterUsersRouter
	ScoringRulesRouter
	ScoreRecordRouter
}

var (
	twitterUsersApi = api.ApiGroupApp.TwitterApiGroup.TwitterUsersApi
	scoringRulesApi = api.ApiGroupApp.TwitterApiGroup.ScoringRulesApi
	scoreRecordApi  = api.ApiGroupApp.TwitterApiGroup.ScoreRecordApi
)
