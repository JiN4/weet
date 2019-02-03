package service

// ユーザ情報-----！
type User struct {
	UserBasics        UserBasics        `json:"user_basics"`
	UserSpecials      UserSpecials      `json:"user_specials"`
	UserIdealSpecials UserIdealSpecials `json:"user_ideal_specials"`
}

type MatchingUser struct {
	UserBasics             UserBasics        `json:"user_basics"`
	UserSpecials           UserSpecials      `json:"user_specials"`
	UserIdealSpecials      UserIdealSpecials `json:"user_ideal_specials"`
	MatchingQuestions      []uint            `json:"matching_questions"`
	MatchingIdealQuestions []uint            `json:"matching_ideal_questions"`
}

//基本的な情報
type UserBasics struct {
	MatchingFormatName string `json:"matching_format_name"`
	UserID             uint   `json:"user_id"`
	UserName           string `json:"user_name"`
	Image1             string `json:"image1"`
	Image2             string `json:"image2"`
	Image3             string `json:"image3"`
	Sex                string `json:"sex"`
	Age                uint   `json:"age"`
	Residence          string `json:"residence"`
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

//マッチング形式ごとのタイトルと質疑応答
type UserIdealSpecial struct {
	MatchingFormatName           string                       `json:"matching_format_name"`
	UserIdealQuestionsAndAnswers UserIdealQuestionsAndAnswers `json:"user_ideal_questions_and_answers"`
}

type UserIdealSpecials []UserIdealSpecial

type UserIdealQuestionAndAnswer struct {
	QuestionID   uint   `json:"ideal_question_id"`
	QuestionName string `json:"ideal_question_name"`
	AnswerName   string `json:"ideal_answer_name"`
}

type UserIdealQuestionsAndAnswers []UserIdealQuestionAndAnswer

type AnswersByQuestion struct {
	QuestionID       uint              `json:"question_id"`
	CandidateAnswers []CandidateAnswer `json:"candidate_answer"`
}

//一つの質問に対する答えの候補-----！
type Answers []AnswersByQuestion

type CandidateAnswer struct {
	AnswerID   uint   `json:"answer_id"`
	AnswerName string `json:"answer_name"`
}

type FavoUser struct {
	UserID    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	Image1    string `json:"image1"`
	Sex       string `json:"sex"`
	Age       uint   `json:"age"`
	Residence string `json:"residence"`
	Hitokoto  string `json:"hitokoto"`
}

type FriendFavoUsers []FavoUser

type LoveFavoUsers []FavoUser

type MarriageFavoUsers []FavoUser

type RoommateFavoUsers []FavoUser

type AllFavoUsers struct {
	FriendFavoUsers   FriendFavoUsers   `json:"friend_favo_users"`
	LoveFavoUsers     LoveFavoUsers     `json:"love_favo_users"`
	MarriageFavoUsers MarriageFavoUsers `json:"marriage_favo_users"`
	RoommateFavoUsers RoommateFavoUsers `json:"roommate_favo_users"`
}

//-------
// ユーザ情報-----！
type User2 struct {
	UserBasics         UserBasics         `json:"user_basics"`
	UserSpecials       UserSpecials       `json:"user_specials"`
	UserIdealSpecials2 UserIdealSpecials2 `json:"user_ideal_specials"`
}

//マッチング形式ごとのタイトルと質疑応答
type UserIdealSpecial2 struct {
	MatchingFormatName            string                        `json:"matching_format_name"`
	UserIdealQuestionsAndAnswers2 UserIdealQuestionsAndAnswers2 `json:"user_ideal_questions_and_answers"`
}

type UserIdealSpecials2 []UserIdealSpecial2

type UserIdealQuestionsAndAnswer2 struct {
	QuestionID   uint     `json:"ideal_question_id"`
	QuestionName string   `json:"ideal_question_name"`
	AnswerName   []string `json:"ideal_answer_name"`
}

type UserIdealQuestionsAndAnswers2 []UserIdealQuestionsAndAnswer2
