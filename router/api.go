package router

import (
	"github.com/gin-gonic/gin"
	"github.com/weet/controller"
)

func apiRouter(api *gin.RouterGroup) {

	// sampleAPI
	api.GET("/sample/:sample_id", controller.GetSampleById)

	//---マイページ画面---
	//ユーザページ情報を全て取得
	api.GET("/user/:user_id", controller.GetUserFilteredById)

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
