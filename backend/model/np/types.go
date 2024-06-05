package np

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Np struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	ChallengerId  int64         `bson:"challenger_id,omitempty" json:"challenger_id,omitempty"`
	ContestantId  int64         `bson:"contestant_id,omitempty" json:"contestant_id,omitempty"`
	Challenger    string        `bson:"challenger,omitempty" json:"challenger,omitempty"`
	Contestant    string        `bson:"contestant,omitempty" json:"contestant,omitempty"`
	QuestionId    int64         `bson:"question_id,omitempty" json:"question_id,omitempty"`
	Question      string        `bson:"question,omitempty" json:"question,omitempty"`
	Answer1       string        `bson:"answer1,omitempty" json:"answer1,omitempty"`
	Answer2       string        `bson:"answer2,omitempty" json:"answer2,omitempty"`
	CorrectAnswer int64         `bson:"correct_answer,omitempty" json:"correct_answer,omitempty"`
	MyAnswer      int64         `bson:"my_answer,omitempty" json:"my_answer,omitempty"`
	Status        int64         `bson:"status,omitempty" json:"status,omitempty"`
	Bet           int64         `bson:"bet,omitempty" json:"bet,omitempty"`
	QuestionTime  time.Time     `bson:"question_time,omitempty" json:"question_time,omitempty"`
	CreateTime    time.Time     `bson:"create_time,omitempty" json:"create_time,omitempty"`
	UpdateTime    time.Time     `bson:"update_time,omitempty" json:"update_time,omitempty"`
}

type ChallengerCorrectAnswerCountList struct {
	Challenger string `bson:"challenger,omitempty" json:"challenger,omitempty"`
	Count      int64  `bson:"count,omitempty" json:"count,omitempty"`
	Bet        int64  `bson:"bet,omitempty" json:"bet,omitempty"`
}

type ContestantMyAnswerCountList struct {
	Contestant string `bson:"contestant,omitempty" json:"contestant,omitempty"`
	Count      int64  `bson:"count,omitempty" json:"count,omitempty"`
	Bet        int64  `bson:"bet,omitempty" json:"bet,omitempty"`
}

type MaxCountUser struct {
	User  string `bson:"user,omitempty" json:"user,omitempty"`
	Count int64  `bson:"count,omitempty" json:"count,omitempty"`
}

type Bet struct {
	Total int64 `bson:"total,omitempty" json:"total,omitempty"`
}
