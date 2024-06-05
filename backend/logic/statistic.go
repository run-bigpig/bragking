package logic

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/bragking/backend/model/np"
	"github.com/run-bigpig/bragking/backend/types"
)

type Statistic struct {
	ctx     *fiber.Ctx
	npModel *np.NpModel
}

func NewStatistic(ctx *fiber.Ctx) *Statistic {
	return &Statistic{
		ctx:     ctx,
		npModel: np.NewNpModel("np"),
	}
}

// FindCloverBet 获取老C抽成
func (s *Statistic) FindCloverBet() (int64, error) {
	maxUser, err := s.npModel.FindTotalBet(s.ctx.Context())
	if err != nil {
		return 0, err
	}
	return int64(float64(maxUser) * 0.1), nil
}

func (s *Statistic) FindMaxBetUser(req *types.MaxCountUserReq) (*types.MaxCountUser, error) {
	//获取指定条件获取妖精最多的用户
	maxUser, err := s.npModel.FindMaxBetUser(s.ctx.Context(), req.User, map[string]interface{}{
		"status": req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &types.MaxCountUser{
		User:  maxUser.User,
		Count: maxUser.Count,
	}, nil
}

// FindCorrectAnswerCountList 获取设置答案的总概率
func (s *Statistic) FindCorrectAnswerCountList(req *types.AnswerStatisticReq) (*types.AnswerStatisticList, error) {
	//获取答案数量
	answerCount, err := s.npModel.FindCorrectAnswerCount(s.ctx.Context(), req.Answer)
	if err != nil {
		return nil, err
	}
	return &types.AnswerStatisticList{
		Count: answerCount,
	}, nil
}

// FindMyAnswerCountList 获取选择答案的总数量
func (s *Statistic) FindMyAnswerCountList(req *types.AnswerStatisticReq) (*types.AnswerStatisticList, error) {
	//获取答案
	answerCount, err := s.npModel.FindMyAnswerCount(s.ctx.Context(), req.Answer)
	if err != nil {
		return nil, err
	}
	return &types.AnswerStatisticList{
		Count: answerCount,
	}, nil
}

func (s *Statistic) FindChallengerCorrectAnswerCountList(req *types.AnswerStatisticReq) ([]*types.AnswerCountList, error) {
	var list []*types.AnswerCountList
	answerList, err := s.npModel.FindChallengerCorrectAnswerCountList(s.ctx.Context(), req.Name, req.Answer, req.Status)
	if err != nil {
		return nil, err
	}
	for _, a := range answerList {
		list = append(list, &types.AnswerCountList{
			Name:  a.Challenger,
			Count: a.Count,
			Bet:   a.Bet,
		})
	}
	return list, nil
}

// FindContestantMyAnswerCountList 获取每个回答者回答次数
func (s *Statistic) FindContestantMyAnswerCountList(req *types.AnswerStatisticReq) ([]*types.AnswerCountList, error) {
	var list []*types.AnswerCountList
	answerList, err := s.npModel.FindContestantMyAnswerCountList(s.ctx.Context(), req.Name, req.Answer, req.Status)
	if err != nil {
		return nil, err
	}
	for _, a := range answerList {
		list = append(list, &types.AnswerCountList{
			Name:  a.Contestant,
			Count: a.Count,
			Bet:   a.Bet,
		})
	}
	return list, nil
}

func (s *Statistic) FindMaxUser(req *types.MaxCountUserReq) (*types.MaxCountUser, error) {
	maxCount, err := s.npModel.FindMaxUser(s.ctx.Context(), req.User, map[string]interface{}{"status": req.Status})
	if err != nil {
		return nil, err
	}
	return &types.MaxCountUser{
		User:  maxCount.User,
		Count: maxCount.Count,
	}, nil
}
