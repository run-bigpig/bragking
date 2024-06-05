package types

type List struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
}

type AnswerReq struct {
	Id       int64 `json:"id"`
	MyAnswer int64 `json:"myanswer"`
}

type CreateReq struct {
	MyMoney  int64  `json:"mymoney"`
	Question string `json:"question"`
	Answer1  string `json:"answer1"`
	Answer2  string `json:"answer2"`
	MyAnswer int64  `json:"myanswer"`
}

type NPListReq struct {
	Type     int64 `json:"type"`
	ToUserId int64 `json:"touserid"`
	Page     int64 `json:"page"`
}

type NPListRes struct {
	Id       int64 `json:"id"`
	ToUserId int64 `json:"touserid"`
}

type NPInfoReq struct {
	Id int64 `json:"id"`
}

type NPInfoRes struct {
	Challenger    string `json:"challenger"`
	Contestant    string `json:"contestant"`
	Bet           int64  `json:"bet"`
	Question      string `json:"question"`
	Answer1       string `json:"answer1"`
	Answer2       string `json:"answer2"`
	CorrectAnswer int64  `json:"correctanswer"`
	MyAnswer      int64  `json:"myanswer"`
	Status        int64  `json:"status"`
	CreateTime    string `json:"createtime"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type AnswerStatisticReq struct {
	Name   string `json:"name,omitempty"`
	Answer int64  `json:"answer"`
	Status int64  `json:"status"`
}

type AnswerStatisticList struct {
	Count int64 `json:"count"`
}

type AnswerCountList struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
	Bet   int64  `json:"bet"`
}

type MaxCountUserReq struct {
	User   string `json:"user"`
	Status int64  `json:"status"`
}

type MaxCountUser struct {
	User  string `bson:"user,omitempty" json:"user,omitempty"`
	Count int64  `bson:"count,omitempty" json:"count,omitempty"`
}
