package helper

import "gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

func QuestionAnalysis(question *model.Question) {
	question.QuestionAnalysis = &model.QuestionAnalysis{}
	question.TotalAnswers = len(question.Answers)
	for _, answer := range question.Answers {
		if answer.NA {
			question.TotalNa++
			continue
		}
		if answer.IsExist != nil {
			if *answer.IsExist {
				question.TotalIsExist++
			} else {
				question.TotalNotExist++
				continue
			}
		}
		if answer.IsGood != nil {
			if *answer.IsGood {
				question.TotalIsGood++
			} else {
				question.TotalNotGood++
			}
		}
	}
	question.Answers = nil
	question.PercentageIsExist = int(float64(question.TotalIsExist) / float64(question.TotalAnswers) * 100)
	question.PercentageNotExist = int(float64(question.TotalNotExist) / float64(question.TotalAnswers) * 100)
	question.PercentageIsGood = int(float64(question.TotalIsGood) / float64(question.TotalAnswers) * 100)
	question.PercentageNotGood = int(float64(question.TotalNotGood) / float64(question.TotalAnswers) * 100)
	question.PercentageNa = int(float64(question.TotalNa) / float64(question.TotalAnswers) * 100)
}
