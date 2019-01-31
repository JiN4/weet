package main

import "github.com/weet/seed/createSeeds"

func main() {
	createSeeds.CreateSeedSamples()
	createSeeds.CreateSeedUserBasics()
	createSeeds.CreateSeedQuestions()
	createSeeds.CreateSeedAnswers()
	createSeeds.CreateSeedMatchingFormats()
	createSeeds.CreateSeedUserQuestionAndAnswers()
	createSeeds.CreateSeedUserIdealQuestionAndAnswers()
	createSeeds.CreateSeedUserSexes()
	createSeeds.CreateSeedMatchingAges()
	createSeeds.CreateSeedMatchingPrefectures()
	createSeeds.CreateSeedMatchingSexes()
	createSeeds.CreateSeedPrefectures()
	createSeeds.CreateSeedMatchingFormatChoices()
	createSeeds.CreateSeedFavoUsers()
	createSeeds.CreateSeedMutualFavoUsers()
}
