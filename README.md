<!--# opentrivia
API client for Open Trivia DB: https://opentdb.com/

Click [here](https://godoc.org/github.com/pinheirolucas/opentrivia) to see the docs.-->

# opentrivia

This repository contains an API client for [Open Trivia DB](https://opentdb.com/).

opentrivia was developed under golang version 1.8 and not tested on previous versions of the
language.

## Usage

### opentrivia ([godoc](https://godoc.org/github.com/pinheirolucas/opentrivia))

	go get github.com/pinheirolucas/opentrivia

Construct a new Open Trivia API client, then use the services on the client to access
different parts of the API. For example:

```go
client := opentrivia.DefaultClient

// list 10 random questions
questions, err := client.Question.List(nil)
```

Some API methods have optional parameters that can be passed. For example:

```go
client := opentrivia.DefaultClient

// list 20 questions with difficulty easy
options := &opentrivia.QuestionListOptions{
	Difficulty: opentrivia.QuestionDifficultyEasy,
	Limit: 20,
}
questions, err := client.Question.List(options)
```

## License

This library is distributed under the MIT license found in the
[LICENSE](https://github.com/pinheirolucas/opentrivia/blob/master/LICENSE) file.
