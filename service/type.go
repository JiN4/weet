package service

// sampleデータ
type Sample struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// マイページ情報
type Mypage struct {
	Basic  Basic    `json:"basic"`
	Format []Format `json:"format"`
}

//基本的な情報
type Basic struct {
	FormatName string `json:"format_name"`
	UserName   string `json:"user_name"`
	Image1     string `json:"image1"`
	Image2     string `json:"image2"`
	Image3     string `json:"image3"`
	Age        string `json:"age"`
	Hitokoto   string `json:"hitokoto"`
	Comment    string `json:"comment"`
}

//マッチング形式ごとのタイトルと質疑応答
type Format struct {
	FormatName        string              `json:"format_name"`
	QuestionAndAnswer []QuestionAndAnswer `json:"question_and_answer"`
}

//１セットの質疑応答
type QuestionAndAnswer struct {
	QuestionName string `json:"question_name"`
	AnswerName   string `json:"answer_name"`
}

/*
{
	{
		"qId":"1"
		"q":"タバコは吸いますか？"
	}

	{
		{
			"aId":"1"
			"qId":"1"
			"a":"はい"
		}
		{
			"aId":"2"
			"qId":"1"
			"a":"いいえ"
			}
	}






	 {
		"q":"職業"
		"a":"プログラマ"
	 }

	 {
		"q":"お酒は飲みますか？"
		"a":"いいえ"
	 }
}
*/
