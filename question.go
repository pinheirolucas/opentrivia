package opentrivia

type (
	// QuestionCategory is the type for category option.
	QuestionCategory uint8

	// QuestionDifficulty is the type for difficulty option.
	QuestionDifficulty string

	// QuestionType is the type for question type option.
	QuestionType string
)

const (
	// QuestionCategoryAny is the value for "any" category type.
	QuestionCategoryAny QuestionCategory = 0

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
	// QuestionDifficultyAny is the value for "any" difficulty type.
	QuestionDifficultyAny QuestionDifficulty = "any"

	// QuestionDifficultyEasy is the value for "easy" difficulty type.
	QuestionDifficultyEasy QuestionDifficulty = "easy"

	// QuestionDifficultyMedium is the value for "medium" difficulty type.
	QuestionDifficultyMedium QuestionDifficulty = "medium"

	// QuestionDifficultyHard is the value for "hard" difficulty type.
	QuestionDifficultyHard QuestionDifficulty = "hard"
)

const (
	// QuestionTypeAny is the value for "any" question type.
	QuestionTypeAny QuestionType = "any"

	// QuestionTypeMultiple is the value for "multiple" question type.
	QuestionTypeMultiple QuestionType = "multiple"

	// QuestionTypeTrueFalse is the value for "true/false" question type.
	QuestionTypeTrueFalse QuestionType = "boolean"
)

// Question is the model of the Open Trivia API Question related
// methods.
type Question struct {
	Category        string   `json:"category"`
	Type            string   `json:"type"`
	Difficulty      string   `json:"difficulty"`
	Question        string   `json:"question"`
	CorrectAnswer   string   `json:"correct_answer"`
	IncorrectAnswer []string `json:"incorrect_answer"`
}

// QuestionListOptions are the options for QuestionService List
// method.
type QuestionListOptions struct {
	// The maximum limit is 50.
	Limit      uint8
	Category   QuestionCategory
	Difficulty QuestionDifficulty
	Type       QuestionType
	Token      Token
}

// QuestionRandomOptions are the options for QuestionService Random
// method.
type QuestionRandomOptions struct {
	Category   QuestionCategory
	Difficulty QuestionDifficulty
	Type       QuestionType
	Token      Token
}

// QuestionService handles communication with the question related
// methods of the Open Trivia API.
//
// Ref.: https://opentdb.com/api_config.php
type QuestionService service

// List returns a list of random questions from Open Trivia API.
func (q *QuestionService) List(options *QuestionListOptions) ([]*Question, error) {
	return nil, nil
}

// Random returns a random question from Open Trivia API.
func (q *QuestionService) Random(options *QuestionRandomOptions) (*Question, error) {
	return nil, nil
}
