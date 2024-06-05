package consts

const (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"
	Referer   = "https://yaohuo.me/games/gamesindex.aspx"
	Success   = "success"
	Fail      = "fail"
)

// Url地址
const (
	BaseUrl   = "https://yaohuo.me"
	ListUrl   = BaseUrl + "/games/chuiniu/index.aspx?siteid=1000&classid=0&r=%d"
	AnswerUrl = BaseUrl + "/games/chuiniu/doit.aspx?myanswer=%d&action=gomod&classid=0&siteid=1000&id=%d&bt=确定"
	CreateUrl = BaseUrl + "/games/chuiniu/add.aspx?mymoney=%d&question=%s&answer1=%s&answer2=%s&myanswer=%d&action=gomod&classid=0&siteid=1000&bt=确+定"
	NPListUrl = BaseUrl + "/games/chuiniu/book_list.aspx?type=%d&siteid=1000&classid=0&page=%d%s"
	NPInfoUrl = BaseUrl + "/games/chuiniu/book_view.aspx?id=%d"
)

type AnswerStatus int

const (
	AnswerSuccess AnswerStatus = iota + 1 //获胜
	AnswerFail                            //失败
	AnswerMoney                           //没钱
	AnswerNo                              //没回答到
)

type CreateStatus int

const (
	CreateSuccess CreateStatus = iota + 1 //成功
	CreateFail                            //失败
	CreateNoMoney                         //没钱
)
