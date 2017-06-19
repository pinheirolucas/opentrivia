package opentrivia

import (
	"github.com/google/go-querystring/query"
	shuffle "github.com/shogo82148/go-shuffle"
)

type (
	// QuestionCategory is the type for category option.
	QuestionCategory uint8

	// QuestionDifficulty is the type for difficulty option.
	QuestionDifficulty string

	// QuestionType is the type for question type option.
	QuestionType string
)

const (
	// QuestionCategoryGeneralKnowledge is the value for
	// "general knowledge" category type.
	QuestionCategoryGeneralKnowledge QuestionCategory = 9

	// QuestionCategoryBook is the value for "book" category type.
	QuestionCategoryBook QuestionCategory = 10

	// QuestionCategoryFilm is the value for "film" category type.
	QuestionCategoryFilm QuestionCategory = 11

	// QuestionCategoryMusic is the value for "music" category type.
	QuestionCategoryMusic QuestionCategory = 12

	// QuestionCategoryMusical is the value for "musical" category type.
	QuestionCategoryMusical QuestionCategory = 13

	// QuestionCategoryTelevision is the value for "television" category type.
	QuestionCategoryTelevision QuestionCategory = 14

	// QuestionCategoryVideoGame is the value for "video game" category type.
	QuestionCategoryVideoGame QuestionCategory = 15

	// QuestionCategoryBoardGame is the value for "board game" category type.
	QuestionCategoryBoardGame QuestionCategory = 16

	// QuestionCategoryNature is the value for "nature" category type.
	QuestionCategoryNature QuestionCategory = 17

	// QuestionCategoryComputer is the value for "computer" category type.
	QuestionCategoryComputer QuestionCategory = 18

	// QuestionCategoryMath is the value for "math" category type.
	QuestionCategoryMath QuestionCategory = 19

	// QuestionCategoryMythology is the value for "mythology" category type.
	QuestionCategoryMythology QuestionCategory = 20

	// QuestionCategorySport is the value for "sport" category type.
	QuestionCategorySport QuestionCategory = 21

	// QuestionCategoryGeography is the value for "geography" category type.
	QuestionCategoryGeography QuestionCategory = 22

	// QuestionCategoryHistory is the value for "history" category type.
	QuestionCategoryHistory QuestionCategory = 23

	// QuestionCategoryPolitics is the value for "politics" category type.
	QuestionCategoryPolitics QuestionCategory = 24

	// QuestionCategoryArt is the value for "art" category type.
	QuestionCategoryArt QuestionCategory = 25

	// QuestionCategoryCelebrity is the value for "celebrity" category type.
	QuestionCategoryCelebrity QuestionCategory = 26

	// QuestionCategoryAnimal is the value for "animal" category type.
	QuestionCategoryAnimal QuestionCategory = 27

	// QuestionCategoryVehicles is the value for "vehicle" category type.
	QuestionCategoryVehicles QuestionCategory = 28

	// QuestionCategoryComic is the value for "comic" category type.
	QuestionCategoryComic QuestionCategory = 29

	// QuestionCategoryGadget is the value for "gadget" category type.
	QuestionCategoryGadget QuestionCategory = 30

	// QuestionCategoryAnime is the value for "anime" category type.
	QuestionCategoryAnime QuestionCategory = 31

	// QuestionCategoryCartoon is the value for "cartoon" category type.
	QuestionCategoryCartoon QuestionCategory = 32
)

const (
	// QuestionDifficultyEasy is the value for "easy" difficulty type.
	QuestionDifficultyEasy QuestionDifficulty = "easy"

	// QuestionDifficultyMedium is the value for "medium" difficulty type.
	QuestionDifficultyMedium QuestionDifficulty = "medium"

	// QuestionDifficultyHard is the value for "hard" difficulty type.
	QuestionDifficultyHard QuestionDifficulty = "hard"
)

const (
	// QuestionTypeMultiple is the value for "multiple" question type.
	QuestionTypeMultiple QuestionType = "multiple"

	// QuestionTypeTrueFalse is the value for "true/false" question type.
	QuestionTypeTrueFalse QuestionType = "boolean"
)

var (
	// DefaultQuestionListOptions is the default options of Question List
	// method.
	DefaultQuestionListOptions = &QuestionListOptions{
		AutoRefresh: false,
		Limit:       10,
	}

	// DefaultQuestionRandomOptions is the default options of Question Random
	// method.
	DefaultQuestionRandomOptions = &QuestionRandomOptions{
		AutoRefresh: false,
	}
)

// QuestionListOptions are the options for QuestionService List
// method.
type QuestionListOptions struct {
	// If true, the request will refresh the provided token when needed.
	AutoRefresh bool `url:"-"`

	// The maximum limit is 50.
	Category   QuestionCategory   `url:"category,omitempty"`
	Difficulty QuestionDifficulty `url:"difficulty,omitempty"`
	Limit      uint8              `url:"amount,omitempty"`
	Token      Token              `url:"token,omitempty"`
	Type       QuestionType       `url:"type,omitempty"`
}

// QuestionRandomOptions are the options for QuestionService Random
// method.
type QuestionRandomOptions struct {
	// If true, the request will refresh the provided token when needed.
	AutoRefresh bool `url:"-"`

	Category   QuestionCategory   `url:"category,omitempty"`
	Difficulty QuestionDifficulty `url:"difficulty,omitempty"`
	Token      Token              `url:"token,omitempty"`
	Type       QuestionType       `url:"type,omitempty"`
}

type questionResponse struct {
	ResponseCode responseCode `json:"response_code"`
	Results      []Question   `json:"results"`
}

// Question is the model of the Open Trivia API Question related
// methods.
type Question struct {
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

// IsAnswerCorrect helps to find out if the provided answer is correct
func (q *Question) IsAnswerCorrect(answer string) bool {
	return answer == q.CorrectAnswer
}

// ShuffleAnswers merging the correct answer with the incorrect answers
func (q *Question) ShuffleAnswers() []string {
	answers := append(q.IncorrectAnswers, q.CorrectAnswer)

	shuffle.Strings(answers)

	return answers
}

// QuestionService handles communication with the question related
// methods of the Open Trivia API.
//
// Ref.: https://opentdb.com/api_config.php
type QuestionService service

// List returns a list of random questions from Open Trivia API.
//
// If options is nil, List will use opentrivia.DefaultQuestionListOptions.
func (q *QuestionService) List(options *QuestionListOptions) ([]Question, error) {
	if options == nil {
		options = DefaultQuestionListOptions
	} else if options.Limit <= 0 {
		options.Limit = DefaultQuestionListOptions.Limit
	} else if options.Limit > 50 {
		options.Limit = 50
	}

	v, err := query.Values(options)
	if err != nil {
		return []Question{}, err
	}

	req, err := q.client.NewRequest(defaultAPIRoute, v)
	if err != nil {
		return []Question{}, err
	}

	var resp questionResponse
	if _, err := q.client.Do(req, &resp); err != nil {
		return []Question{}, err
	}

	switch resp.ResponseCode {
	case responseCodeInvalidParameter:
		return []Question{}, ErrInvalidParameter
	case responseCodeNoResults:
		return []Question{}, ErrNoResults
	case responseCodeTokenEmpty:
		if options.AutoRefresh {
			t, err := q.client.Token.Refresh(options.Token)
			if err != nil {
				return []Question{}, err
			}

			options.Token = t
			return q.client.Question.List(options)
		}

		return []Question{}, ErrTokenEmpty
	case responseCodeTokenNotFound:
		return []Question{}, ErrTokenNotFound
	}

	return resp.Results, nil
}

// Random returns a random question from Open Trivia API.
//
// If options is nil, Random will use opentrivia.DefaultQuestionRandomOptions.
func (q *QuestionService) Random(options *QuestionRandomOptions) (Question, error) {
	if options == nil {
		options = DefaultQuestionRandomOptions
	}

	v, err := query.Values(options)
	if err != nil {
		return Question{}, err
	}
	// This method requires a single result
	v.Set("amount", "1")

	req, err := q.client.NewRequest(defaultAPIRoute, v)
	if err != nil {
		return Question{}, err
	}

	var resp questionResponse
	if _, err := q.client.Do(req, &resp); err != nil {
		return Question{}, err
	}

	switch resp.ResponseCode {
	case responseCodeInvalidParameter:
		return Question{}, ErrInvalidParameter
	case responseCodeNoResults:
		return Question{}, ErrNoResults
	case responseCodeTokenEmpty:
		if options.AutoRefresh {
			t, err := q.client.Token.Refresh(options.Token)
			if err != nil {
				return Question{}, err
			}

			options.Token = t
			return q.client.Question.Random(options)
		}

		return Question{}, ErrTokenEmpty
	case responseCodeTokenNotFound:
		return Question{}, ErrTokenNotFound
	}

	return resp.Results[0], nil
}
