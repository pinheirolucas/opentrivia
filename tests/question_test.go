package tests

import (
	"testing"

	"github.com/pinheirolucas/opentrivia"
)

func TestList(t *testing.T) {
	t.Parallel()

	t.Run("expect the default limit to be 10", func(t *testing.T) {
		t.Parallel()

		const expectedLength = 10
		options := &opentrivia.QuestionListOptions{}

		list, err := client.Question.List(options)
		if err != nil {
			t.Error(err)
		}

		resultLength := len(list)

		if resultLength != expectedLength {
			t.Errorf("Expected %d, got %d", expectedLength, resultLength)
		}
	})

	t.Run("expect the count of questions being the same as provided by options", func(t *testing.T) {
		t.Parallel()

		const expectedLength = 20
		options := &opentrivia.QuestionListOptions{
			Limit: expectedLength,
		}

		list, err := client.Question.List(options)
		if err != nil {
			t.Error(err)
		}

		resultLength := len(list)

		if resultLength != expectedLength {
			t.Errorf("Expected %d, got %d", expectedLength, resultLength)
		}
	})

	t.Run("expect all the questions difficulties being the same as provided by options", func(t *testing.T) {
		t.Parallel()

		const expectedDifficulty = "easy"
		options := &opentrivia.QuestionListOptions{
			Difficulty: opentrivia.QuestionDifficultyEasy,
			Limit:      20,
		}

		list, err := client.Question.List(options)
		if err != nil {
			t.Error(err)
		}

		for _, v := range list {
			if v.Difficulty != expectedDifficulty {
				t.Errorf(
					"All the results should be of difficulty %s, got a result with difficulty %s",
					expectedDifficulty, v.Difficulty,
				)
				return
			}
		}
	})

	t.Run("expect all the questions categories being the same as provided by options", func(t *testing.T) {
		t.Parallel()

		const expectedCategory = "Entertainment: Video Games"
		options := &opentrivia.QuestionListOptions{
			Category: opentrivia.QuestionCategoryVideoGame,
			Limit:    20,
		}

		list, err := client.Question.List(options)
		if err != nil {
			t.Error(err)
		}

		for _, v := range list {
			if v.Category != expectedCategory {
				t.Errorf(
					"All the results should be of category %s, got a result with category %s",
					expectedCategory, v.Category,
				)
				return
			}
		}
	})

	t.Run("expect all the questions types being the same as provided by options", func(t *testing.T) {
		t.Parallel()

		const expectedType = "multiple"
		options := &opentrivia.QuestionListOptions{
			Type:  opentrivia.QuestionTypeMultiple,
			Limit: 20,
		}

		list, err := client.Question.List(options)
		if err != nil {
			t.Error(err)
		}

		for _, v := range list {
			if v.Type != expectedType {
				t.Errorf(
					"All the results should be of type %s, got a result with type %s",
					expectedType, v.Type,
				)
				return
			}
		}
	})

	t.Run("expect to use the default options in case of no options", func(t *testing.T) {
		t.Parallel()

		const expectedLength = 10

		list, err := client.Question.List(nil)
		if err != nil {
			t.Error(err)
		}

		resultLength := len(list)

		if resultLength != expectedLength {
			t.Errorf("Expected length of %d, got %d", expectedLength, resultLength)
		}
	})

	t.Run("expect to compose the options", func(t *testing.T) {
		t.Parallel()

		const expectedLength = 20
		const expectedDifficulty = "easy"
		const expectedCategory = "Entertainment: Video Games"
		const expectedType = "multiple"
		options := &opentrivia.QuestionListOptions{
			Category:   opentrivia.QuestionCategoryVideoGame,
			Difficulty: opentrivia.QuestionDifficultyEasy,
			Limit:      expectedLength,
			Type:       opentrivia.QuestionTypeMultiple,
		}

		list, err := client.Question.List(options)
		if err != nil {
			t.Error(err)
		}
		resultLength := len(list)

		if resultLength != expectedLength {
			t.Errorf("Expected %d, got %d", expectedLength, resultLength)
		}
		for _, v := range list {
			if v.Category != expectedCategory {
				t.Errorf(
					"All the results should be of category %s, got a result with category %s",
					expectedCategory, v.Category,
				)
				return
			} else if v.Difficulty != expectedDifficulty {
				t.Errorf(
					"All the results should be of difficulty %s, got a result with difficulty %s",
					expectedDifficulty, v.Difficulty,
				)
				return
			} else if v.Type != expectedType {
				t.Errorf(
					"All the results should be of type %s, got a result with type %s",
					expectedType, v.Type,
				)
				return
			}
		}
	})

	t.Run("expect to return opentrivia.ErrNoResults", func(t *testing.T) {
		t.Parallel()

		options := &opentrivia.QuestionListOptions{
			Category: 1,
		}

		_, err := client.Question.List(options)
		if err != opentrivia.ErrNoResults {
			t.Error(err)
		}
	})

	t.Run("expect to return opentrivia.ErrInvalidParameter", func(t *testing.T) {
		t.Parallel()

		options := &opentrivia.QuestionListOptions{
			Difficulty: "jasldjalkdkalsd",
		}

		_, err := client.Question.List(options)
		if err != opentrivia.ErrInvalidParameter {
			t.Error(err)
		}
	})

	t.Run("expect to return opentrivia.ErrTokenNotFound", func(t *testing.T) {
		t.Parallel()

		const fakeToken = "not_a_token"
		options := &opentrivia.QuestionListOptions{
			Token: fakeToken,
		}

		_, err := client.Question.List(options)
		if err != opentrivia.ErrTokenNotFound {
			t.Error(err)
		}
	})
}

func TestRandom(t *testing.T) {
	t.Parallel()

	t.Run("expect the question difficulty to be the same as provided by options", func(t *testing.T) {
		t.Parallel()

		const expectedDifficulty = "easy"
		options := &opentrivia.QuestionRandomOptions{
			Difficulty: opentrivia.QuestionDifficultyEasy,
		}

		q, err := client.Question.Random(options)
		if err != nil {
			t.Error(err)
		}

		if q.Difficulty != expectedDifficulty {
			t.Errorf("Expect the difficulty to be %s, got %s", expectedDifficulty, q.Difficulty)
		}
	})

	t.Run("expect the question category to be the same as provided by options", func(t *testing.T) {
		t.Parallel()

		const expectedCategory = "Entertainment: Video Games"
		options := &opentrivia.QuestionRandomOptions{
			Category: opentrivia.QuestionCategoryVideoGame,
		}

		q, err := client.Question.Random(options)
		if err != nil {
			t.Error(err)
		}

		if q.Category != expectedCategory {
			t.Errorf("Expect the category to be %s, got %s", expectedCategory, q.Category)
		}
	})

	t.Run("expect the question type to be the same as provided by options", func(t *testing.T) {
		t.Parallel()

		const expectedType = "multiple"
		options := &opentrivia.QuestionRandomOptions{
			Type: opentrivia.QuestionTypeMultiple,
		}

		q, err := client.Question.Random(options)
		if err != nil {
			t.Error(err)
		}

		if q.Type != expectedType {
			t.Errorf("Expect the type to be %s, got %s", expectedType, q.Type)
		}
	})

	t.Run("expect to compose the options", func(t *testing.T) {
		t.Parallel()

		const expectedDifficulty = "easy"
		const expectedCategory = "Entertainment: Video Games"
		const expectedType = "multiple"
		options := &opentrivia.QuestionRandomOptions{
			Category:   opentrivia.QuestionCategoryVideoGame,
			Difficulty: opentrivia.QuestionDifficultyEasy,
			Type:       opentrivia.QuestionTypeMultiple,
		}

		q, err := client.Question.Random(options)
		if err != nil {
			t.Error(err)
		}

		if q.Category != expectedCategory || q.Difficulty != expectedDifficulty || q.Type != expectedType {
			t.Error("The composition failed")
		}
	})

	t.Run("expect to return opentrivia.ErrNoResults", func(t *testing.T) {
		t.Parallel()

		options := &opentrivia.QuestionRandomOptions{
			Category: 1,
		}

		_, err := client.Question.Random(options)
		if err != opentrivia.ErrNoResults {
			t.Error(err)
		}
	})

	t.Run("expect to return opentrivia.ErrInvalidParameter", func(t *testing.T) {
		t.Parallel()

		options := &opentrivia.QuestionRandomOptions{
			Difficulty: "jasldjalkdkalsd",
		}

		_, err := client.Question.Random(options)
		if err != opentrivia.ErrInvalidParameter {
			t.Error(err)
		}
	})

	t.Run("expect to return opentrivia.ErrTokenNotFound", func(t *testing.T) {
		t.Parallel()

		const fakeToken = "not_a_token"
		options := &opentrivia.QuestionRandomOptions{
			Token: fakeToken,
		}

		_, err := client.Question.Random(options)
		if err != opentrivia.ErrTokenNotFound {
			t.Error(err)
		}
	})
}
