package router

import (
	"github.com/gin-gonic/gin"
	"github.com/weet/controller"
)

func apiRouter(api1 *gin.RouterGroup, api2 *gin.RouterGroup) {

	//ID別にユーザページ情報を取得
	api1.GET("/user/:user_id", controller.GetUserById)

	//ID別にユーザページ情報を取得2
	api2.GET("/user/:user_id", controller.GetUserById2)

	//全ての回答を取得
	api1.GET("/answers", controller.GetAnswers)

	//マッチング相手を性別・居住地・年齢で絞り込みランダムで取得
	api1.GET("/matching/player/:user_id/matching-format/:matching_format_id", controller.GetMatchingUser)

	//プレイヤーをお気に入りにしているユーザを全て取得
	api1.GET("/favo/user/:user_id", controller.GetFavoUsersById)

	//プレイヤーと相互にお気に入りにしているユーザを全て取得
	api1.GET("/mutual-favo/user/:user_id", controller.GetMutualFavoUsersById)

	//プレイヤーがお気に入りのユーザを登録する
	api1.POST("/favo/user", controller.PostFavoUsers)
	//curl http://localhost:8080/api/v1/favo/user -X POST -H "Content-Type: application/json" -d '{"MatchingFormatID": 1,"PlayerUserID": 1,"FavoUserID": 2}'

	//プレイヤーの相互いいね！ユーザを登録する
	api1.POST("/mutual-favo/user", controller.PostMutualFavoUsers)
	//curl http://localhost:8080/api/v1/mutual-favo/user -X POST -H "Content-Type: application/json" -d '{"MatchingFormatID": 1,"UserID1": 1,"UserID2": 2}'

	//ユーザの基本情報を更新する
	api1.PUT("/user/:user_id/update/basics", controller.UpdateUserBasics)

	//ユーザの質疑応答の自分情報を更新する
	api1.PUT("/user/:user_id/update/specials", controller.UpdateUserSpecials)
	//curl -w '\n' 'http://localhost:8080/api/v1/user/1/update/specials' --data 'question_id=1&answer_id=27' -XPUT

	//ユーザの質疑応答の理想像を更新する
	api1.PUT("/user/:user_id/update/ideal-specials", controller.UpdateUserIdealSpecials)
	//curl -w '\n' 'http://localhost:8080/api/v1/user/1/update/ideal-specials' --data 'matching_format_id=1&question_id=1&answer_ids=3,4,47' -XPUT

	//求める居住地を更新する
	api1.PUT("/matching-prefectures/:user_id", controller.PutMatchingPrefectures)
	//curl http://localhost:8080/api/v1/matching-prefectures/1 -X PUT -H "Content-Type: application/json" -d '{"PrefecturesID": "1"}'

	//求める性別を更新する
	api1.PUT("/matching-sexes/:user_id", controller.PutMatchingSexes)
	//curl http://localhost:8080/api/v1/matching-sexes/1 -X PUT -H "Content-Type: application/json" -d '{"SexID": 3}'

	//求める相手の年齢を更新する
	api1.PUT("/matching-ages/:user_id", controller.PutMatcingAges)

	//マッチングされる形式を更新する
	api1.PUT("/matching-format-choices/:user_id", controller.PutMatchingFormatChoices)

	//プレイヤーをお気に入りにしているユーザを削除
	api1.DELETE("/favo/player/:player_user_id/favo-user/:favo_user_id/matching-format/:matching_format_id", controller.DeleteFavoUsers)
	//curl http://http://54.238.92.95:8080/api/v1/favo/player/1/favo-user/7/matching-format/2 -X DELETE -H "Content-Type: application/json"
}
