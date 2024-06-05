package logic

import (
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/run-bigpig/bragking/backend/config"
	"github.com/run-bigpig/bragking/backend/consts"
	"github.com/run-bigpig/bragking/backend/types"
	"github.com/run-bigpig/bragking/backend/utils"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type YaoHuo struct {
	ctx    context.Context
	userId int64
	header map[string]string
}

func NewYaoHuo(ctx context.Context) *YaoHuo {
	return &YaoHuo{
		ctx:    ctx,
		userId: utils.ExtractUid(config.Get().Cookie),
		header: map[string]string{"User-Agent": consts.UserAgent, "Referer": consts.Referer, "Cookie": config.Get().Cookie},
	}
}

func (a *YaoHuo) List() ([]*types.List, error) {
	url := fmt.Sprintf(consts.ListUrl, rand.Intn(99999))
	list := make([]*types.List, 0)
	c := colly.NewCollector(colly.UserAgent(consts.UserAgent))
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", config.Get().Cookie)
		r.Headers.Set("Referer", consts.Referer)
	})
	c.OnResponse(func(r *colly.Response) {
		c.OnHTML("body>div[class^=line]", func(e *colly.HTMLElement) {
			e.ForEach("a", func(i int, e *colly.HTMLElement) {
				if strings.Contains(e.Attr("href"), "doit.aspx") {
					idstr := utils.GetUrlParam(e.Attr("href"), "id")
					id, _ := strconv.Atoi(idstr)
					list = append(list, &types.List{
						Id:      int64(id),
						Content: e.Text,
					})
				}
			})
		})
	})
	err := c.Visit(url)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (a *YaoHuo) Answer(answer *types.AnswerReq) consts.AnswerStatus {
	url := fmt.Sprintf(consts.AnswerUrl, answer.MyAnswer, answer.Id)
	data, err := utils.SendRequest(http.MethodPost, url, a.header, nil)
	if err != nil {
		return consts.AnswerFail
	}
	switch {
	case strings.Contains(string(data), "妖晶不够了"):
		return consts.AnswerMoney
	case strings.Contains(string(data), "获胜"):
		return consts.AnswerSuccess
	case strings.Contains(string(data), "失败"):
		return consts.AnswerFail
	default:
		return consts.AnswerNo
	}
}

func (a *YaoHuo) Create(create *types.CreateReq) consts.CreateStatus {
	url := fmt.Sprintf(consts.CreateUrl, create.MyMoney, create.Question, create.Answer1, create.Answer2, create.MyAnswer)
	data, err := utils.SendRequest(http.MethodPost, url, a.header, nil)
	if err != nil {
		return consts.CreateFail
	}
	switch {
	case strings.Contains(string(data), "妖精不够了"):
		return consts.CreateNoMoney
	case strings.Contains(string(data), "成功创建"):
		return consts.CreateSuccess
	default:
		return consts.CreateFail
	}
}

// NPList 破产列表
func (a *YaoHuo) NPList(nplist *types.NPListReq) (int64, []*types.NPListRes, error) {
	var total int
	var touserid string
	if nplist.ToUserId != 0 {
		touserid = fmt.Sprintf("&touserid=%d", nplist.ToUserId)
	}
	url := fmt.Sprintf(consts.NPListUrl, nplist.Type, nplist.Page, touserid)
	list := make([]*types.NPListRes, 0)
	c := colly.NewCollector(colly.UserAgent(consts.UserAgent))
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", config.Get().Cookie)
		r.Headers.Set("Referer", consts.Referer)
	})
	c.OnResponse(func(r *colly.Response) {
		c.OnHTML("body>div[class^=line]", func(e *colly.HTMLElement) {
			item := &types.NPListRes{}
			e.ForEach("a", func(i int, e *colly.HTMLElement) {
				switch {
				case strings.Contains(e.Attr("href"), "userinfo.aspx"):
					idstr := utils.GetUrlParam(e.Attr("href"), "touserid")
					id, _ := strconv.Atoi(idstr)
					item.ToUserId = int64(id)
				case strings.Contains(e.Attr("href"), "book_view.aspx"):
					idstr := utils.GetUrlParam(e.Attr("href"), "id")
					id, _ := strconv.Atoi(idstr)
					item.Id = int64(id)
				}
			})
			if item.Id != 0 {
				list = append(list, item)
			}
		})
		c.OnHTML("body>div.showpage>form>input[name=getTotal]", func(e *colly.HTMLElement) {
			total, _ = strconv.Atoi(e.Attr("value"))
		})
	})
	err := c.Visit(url)
	if err != nil {
		return 0, nil, err
	}
	return int64(total), list, nil
}

// NPInfo 破产信息
func (a *YaoHuo) NPInfo(npinfo *types.NPInfoReq) (*types.NPInfoRes, error) {
	url := fmt.Sprintf(consts.NPInfoUrl, npinfo.Id)
	c := colly.NewCollector(colly.UserAgent(consts.UserAgent))
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", config.Get().Cookie)
		r.Headers.Set("Referer", consts.Referer)
	})
	info := &types.NPInfoRes{}
	c.OnResponse(func(r *colly.Response) {
		c.OnHTML("body>div[class=content]", func(e *colly.HTMLElement) {
			htmlContent, err := e.DOM.Html()
			if err != nil {
				return
			}
			lists := strings.Split(htmlContent, "<br/>")
			for _, list := range lists {
				list = strings.TrimSpace(list)
				listSlice := strings.Split(list, ":")
				option := listSlice[0] + ":"
				switch {
				case strings.Contains(option, "挑战者:"):
					info.Challenger = listSlice[1]
				case strings.Contains(option, "赌注是:"):
					info.Bet = utils.StringToInt64(strings.ReplaceAll(listSlice[1], "妖晶", ""))
				case strings.Contains(option, "问题是:"):
					info.Question = listSlice[1]
				case strings.Contains(option, "答案1:"):
					info.Answer1 = listSlice[1]
				case strings.Contains(option, "答案2:"):
					info.Answer2 = listSlice[1]
				case strings.Contains(option, "发起时间:"):
					info.CreateTime = listSlice[1]
					if len(listSlice) == 4 {
						info.CreateTime = fmt.Sprintf("%s:%s:%s", listSlice[1], listSlice[2], listSlice[3])
					}
				case strings.Contains(option, "应战者:"):
					info.Contestant = listSlice[1]
				case strings.Contains(option, "状态:"):
					status := listSlice[1]
					switch {
					case strings.Contains(status, "获胜"):
						info.Status = 1
					case strings.Contains(status, "失败"):
						info.Status = 2
					case strings.Contains(status, "进行中"):
						info.Status = 3
					default:
						info.Status = 4
					}
				case strings.Contains(option, "挑战方出的是"):
					info.CorrectAnswer = a.extractAnswer(listSlice[0])
				case strings.Contains(option, "应战方出的是"):
					info.MyAnswer = a.extractAnswer(listSlice[0])
				}
			}
		})
	})
	err := c.Visit(url)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// 提取答案
func (a *YaoHuo) extractAnswer(s string) int64 {
	s = strings.TrimRight(s, "]")
	tempSlice := strings.Split(s, "[")
	if len(tempSlice) != 2 {
		return 0
	}
	switch strings.TrimSpace(tempSlice[1]) {
	case "答案1":
		return 1
	case "答案2":
		return 2
	default:
		return 0
	}
}
