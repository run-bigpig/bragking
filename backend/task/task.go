package task

import (
	"context"
	yaohuo "github.com/run-bigpig/bragking/backend/logic"
	"github.com/run-bigpig/bragking/backend/model/np"
	"github.com/run-bigpig/bragking/backend/types"
	"github.com/run-bigpig/bragking/backend/utils"
	"log"
	"math"
	"slices"
	"strings"
	"time"
)

type Task struct {
	ctx     context.Context
	yh      *yaohuo.YaoHuo
	npModel *np.NpModel
}

func NewTask(ctx context.Context) *Task {
	return &Task{
		yh:      yaohuo.NewYaoHuo(ctx),
		ctx:     ctx,
		npModel: np.NewNpModel("np"),
	}
}

func (t *Task) StartUp() {
	t.getCreateNpList()
	t.getQuestionInfo()
}

func (t *Task) getCreateNpList() {
	//获取数据库总数
	total, err := t.npModel.Count()
	if err != nil {
		return
	}
	//获取列表总数
	listTotal := t.getListTotal()
	diff := listTotal - total
	if diff > 0 {
		totalPage := math.Ceil(float64(diff / 15))
		for i := 1; i <= int(totalPage); i++ {
			time.Sleep(1500 * time.Millisecond)
			log.Println("获取抢话第", i, "页")
			req := &types.NPListReq{
				Type:     1,
				ToUserId: 0,
				Page:     int64(i),
			}
			//<-t.createChan
			_, list, err := t.yh.NPList(req)
			if err != nil {
				log.Println(err)
				continue
			}
			for _, v := range list {
				whereMap := map[string]interface{}{
					"question_id": v.Id,
				}
				updateMap := map[string]interface{}{
					"challenger_id": v.ToUserId,
				}
				err = t.npModel.FindOneAndUpdate(t.ctx, whereMap, updateMap)
				if err != nil {
					log.Println(err)
					continue
				}
			}
		}
	}
}

func (t *Task) getAnswerNpList() {
	//获取数据库总数
	total, err := t.npModel.Count()
	if err != nil {
		return
	}
	//获取列表总数
	listTotal := t.getListTotal()
	diff := listTotal - total
	if diff > 0 {
		totalPage := math.Ceil(float64(diff / 15))
		for i := 1; i <= int(totalPage); i++ {
			time.Sleep(1 * time.Second)
			log.Println("获取大话第", i, "页")
			req := &types.NPListReq{
				Type:     0,
				ToUserId: 0,
				Page:     int64(i),
			}
			_, list, err := t.yh.NPList(req)
			if err != nil {
				log.Println(err)
				continue
			}
			for _, v := range list {
				if v.ToUserId == 0 {
					continue
				}
				whereMap := map[string]interface{}{
					"question_id": v.Id,
				}
				updateMap := map[string]interface{}{
					"contestant_id": v.ToUserId,
				}
				err = t.npModel.FindOneAndUpdate(t.ctx, whereMap, updateMap)
				if err != nil {
					log.Println(err)
					continue
				}
			}
		}
	}
}

func (t *Task) getQuestionInfo() {
	list, err := t.npModel.FindNotExistFiled(t.ctx, "challenger")
	if err != nil {
		return
	}
	for _, v := range list {
		time.Sleep(2000 * time.Millisecond)
		log.Println("处理问题ID:", v.QuestionId)
		req := &types.NPInfoReq{
			Id: v.QuestionId,
		}
		info, err := t.yh.NPInfo(req)
		if err != nil {
			log.Println(err)
			continue
		}
		statusSlice := []int64{1, 2}
		_, insert := slices.BinarySearch(statusSlice, info.Status)
		if insert {
			whereMap := map[string]interface{}{
				"question_id": v.QuestionId,
			}
			updateMap := map[string]interface{}{
				"challenger":     strings.TrimSpace(info.Challenger),
				"contestant":     strings.TrimSpace(info.Contestant),
				"status":         info.Status,
				"bet":            info.Bet,
				"question":       info.Question,
				"answer1":        info.Answer1,
				"answer2":        info.Answer2,
				"correct_answer": info.CorrectAnswer,
				"my_answer":      info.MyAnswer,
				"question_time":  utils.ParseTime(info.CreateTime, "2006/1/02 15:04:05"),
			}
			err = t.npModel.FindOneAndUpdate(t.ctx, whereMap, updateMap)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}

// 获取列表总数
func (t *Task) getListTotal() int64 {
	req := &types.NPListReq{
		Type:     1,
		ToUserId: 0,
		Page:     1,
	}
	total, _, err := t.yh.NPList(req)
	if err != nil {
		return 0
	}
	return total
}
