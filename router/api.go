package router

import (
	"github.com/gin-gonic/gin"
	"github.com/weet/controller"
)

func apiRouter(api *gin.RouterGroup) {

	// sampleAPI
	api.GET("/sample/:sample_id", controller.GetSampleById)

	//---マイページ画面---
	//ID別にユーザページ情報を取得
	api.GET("/user/:user_id", controller.GetUserById)

	//基本情報を更新
	api.PUT("/user/:user_id/update/basics", controller.UpdateUserBasics)

	//全ての回答を取得
	api.GET("/answers", controller.GetAnswers)

	//マッチング相手
	//性別、居住地（複数）、マッチ項目数
	api.GET("/matching/player/:user_id/matching-format/:matching_format_id", controller.GetMatchingUser)

	//プレイヤーをお気に入りにしているユーザを全て取得
	api.GET("/favo/user/:user_id", controller.GetFavoUsersById)

	//プレイヤーがお気に入りのユーザを登録する
	api.POST("/favo/user", controller.PostFavoUsers)
	//curl http://localhost:8080/api/v1/favo/user -X POST -H "Content-Type: application/json" -d '{"MatchingFormatID": 1,"PlayerUserID": 1,"FavoUserID": 2}'

	//プレイヤーと相互にお気に入りにしているユーザを全て取得
	api.GET("/mutual-favo/user/:user_id", controller.GetMutualFavoUsersById)

	//プレイヤーがお気に入りのユーザを登録する
	api.POST("/mutual-favo/user", controller.PostMutualFavoUsers)
	//curl http://localhost:8080/api/v1/mutual-favo/user -X POST -H "Content-Type: application/json" -d '{"MatchingFormatID": 1,"UserID1": 1,"UserID2": 2}'

	//求める居住地を指定する
	api.PUT("/matching-prefectures/:user_id", controller.PutMatchingPrefectures)
	//curl http://localhost:8080/api/v1/matching-prefectures/1 -X PUT -H "Content-Type: application/json" -d '{"PrefecturesID": "1"}'

	//求める性別を指定する
	api.PUT("/matching-sexes/:user_id", controller.PutMatchingSexes)
	//curl http://localhost:8080/api/v1/matching-sexes/1 -X PUT -H "Content-Type: application/json" -d '{"SexID": 3}'

	//求める相手の年齢を指定する
	api.PUT("/matching-ages/:user_id", controller.PutMatcingAges)

	//マッチングフォーマット判別
	api.PUT("/matching-format-choices/:user_id", controller.PutMatchingFormatChoices)
	/*
		//基本情報の取得
		api.GET("/mypage/basic/:user_id", controller.GetBasicMypageFilteredById)

		//フレンドマッチング情報の取得
		api.GET("/mypage/friend/:user_id", controller.GetFriendMypageFilteredById)

		//恋愛マッチング情報の取得
		api.GET("/mypage/love/:user_id", controller.GetLoveMypageFilteredById)

		//婚活マッチング情報の取得
		api.GET("/mypage/marriage/:user_id", controller.GetMarriageMypageFilteredById)

		//ルームメイトマッチング情報の取得
		api.GET("/mypage/roommate/:user_id", controller.GetRoommateMypageFilteredById)

		//---マッチング画面---
		//選択したマッチング形式でのマッチング相手を取得
		api.GET("/matching/:matching_id/:user_id", controller.GetMatchingFormatFilteredById)
	*/
}
