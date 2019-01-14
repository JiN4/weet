package service

// sampleデータ
type Sample struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// ユーザ情報-----！
type User struct {
	UserBasics   UserBasics   `json:"User_basics"`
	UserSpecials UserSpecials `json:"User_pecials"`
}

//基本的な情報
type UserBasics struct {
	MatchingFormatName string `json:"matching_format_name"`
	UserName           string `json:"user_name"`
	Image1             string `json:"image1"`
	Image2             string `json:"image2"`
	Image3             string `json:"image3"`
	Age                string `json:"age"`
	Hitokoto           string `json:"hitokoto"`
	Comment            string `json:"comment"`
}

//マッチング形式ごとのタイトルと質疑応答
type UserSpecial struct {
	MatchingFormatName      string                  `json:"matching_format_name"`
	UserQuestionsAndAnswers UserQuestionsAndAnswers `json:"user_questions_and_answers"`
}

type UserSpecials []UserSpecial

//１セットの質疑応答
//select user_basics.user_name, questions.name, answers.name from user_question_and_answers join user_basics on (user_question_and_answers.user_id = user_basics.id) left join questions on (user_question_and_answers.question_id = questions.id) left join answers on (user_question_and_answers.answer_id = answers.id);
type UserQuestionAndAnswer struct {
	QuestionID   uint   `json:"question_id"`
	QuestionName string `json:"question_name"`
	AnswerName   string `json:"answer_name"`
}

type UserQuestionsAndAnswers []UserQuestionAndAnswer

//一つの質問に対する答えの候補-----！
type Answers struct {
	QuestionID uint         `json:"question_id"`
	Answers    AnswersCards `json:"answers"`
}

type AnswersCard struct {
	AmswerID   uint   `json:"answer_id"`
	AnswerName string `json:"answer_name"`
}

type AnswersCards []AnswersCard
