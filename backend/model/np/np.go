package np

import (
	"context"
	"errors"
	"github.com/run-bigpig/bragking/backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type NpModel struct {
	conn *mongo.Collection
}

func NewNpModel(collection string) *NpModel {
	return &NpModel{
		conn: model.GetDataBase().Collection(collection),
	}
}

// Count 获取列表总数
func (m *NpModel) Count() (int64, error) {
	total, err := m.conn.CountDocuments(context.Background(), bson.M{})
	return total, err
}

// FindMaxUser 查询指定条件下次数最多的用户
func (m *NpModel) FindMaxUser(ctx context.Context, userFiled string, where map[string]interface{}) (*MaxCountUser, error) {
	var data []*MaxCountUser
	option := &options.AggregateOptions{}
	res, err := m.conn.Aggregate(ctx, []bson.M{
		{"$match": where},
		{"$group": bson.M{
			"_id":   "$" + userFiled,
			"count": bson.M{"$sum": 1},
		}},
		{"$sort": bson.M{"count": -1}},
		{"$project": bson.M{
			"_id":   0,
			"user":  "$_id",
			"count": 1,
		}},
		{"$limit": 1},
	}, option)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	err = res.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return data[0], err
}

// FindNotExistFiled 查找不包含某个字段的文档列表
func (m *NpModel) FindNotExistFiled(ctx context.Context, filed string) ([]*Np, error) {
	var np []*Np
	option := &options.FindOptions{}
	res, err := m.conn.Find(ctx, bson.M{filed: bson.M{"$exists": false}}, option)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	err = res.All(ctx, &np)
	return np, err
}

// FindOneAndUpdate 查找并更新
func (m *NpModel) FindOneAndUpdate(ctx context.Context, whereMap map[string]interface{}, updateMap map[string]interface{}) error {
	option := &options.FindOneAndUpdateOptions{}
	option.SetUpsert(true)
	res := m.conn.FindOneAndUpdate(ctx,
		whereMap,
		bson.M{
			"$set": updateMap,
			"$currentDate": bson.M{
				"update_time": true,
			},
			"$setOnInsert": bson.M{
				"create_time": time.Now(),
			},
		}, option)
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil
	}
	return res.Err()
}

// FindCorrectAnswerCount 获取设置答案的次数
func (m *NpModel) FindCorrectAnswerCount(ctx context.Context, answer int64) (int64, error) {
	total, err := m.conn.CountDocuments(ctx, bson.M{"correct_answer": answer})
	return total, err
}

// FindMyAnswerCount 获取选择答案的次数
func (m *NpModel) FindMyAnswerCount(ctx context.Context, answer int64) (int64, error) {
	total, err := m.conn.CountDocuments(ctx, bson.M{"my_answer": answer})
	return total, err
}

// FindTotalBet 获取总妖晶数
func (m *NpModel) FindTotalBet(ctx context.Context) (int64, error) {
	var data []*Bet
	option := &options.AggregateOptions{}
	res, err := m.conn.Aggregate(ctx, []bson.M{
		{"$group": bson.M{
			"_id": nil,
			"bet": bson.M{"$sum": "$bet"},
		}},
		{"$project": bson.M{
			"_id":   0,
			"total": "$bet",
		}},
	}, option)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return 0, nil
	}
	err = res.All(ctx, &data)
	if err != nil {
		return 0, err
	}
	if len(data) == 0 {
		return 0, nil
	}
	return data[0].Total, err
}

// FindMaxBetUser 获取获得妖晶最多的
func (m *NpModel) FindMaxBetUser(ctx context.Context, userFiled string, where map[string]interface{}) (*MaxCountUser, error) {
	var data []*MaxCountUser
	option := &options.AggregateOptions{}
	res, err := m.conn.Aggregate(ctx, []bson.M{
		{"$match": where},
		{"$group": bson.M{
			"_id":   "$" + userFiled,
			"count": bson.M{"$sum": "$bet"},
		},
		},
		{"$sort": bson.M{"count": -1}},
		{"$project": bson.M{
			"_id":   0,
			"user":  "$_id",
			"count": 1,
		}},
		{"$limit": 1},
	}, option)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	err = res.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return data[0], err
}

// FindChallengerCorrectAnswerCountList 获取挑战者设置指定答案的次数列表
func (m *NpModel) FindChallengerCorrectAnswerCountList(ctx context.Context, Challenger string, answer int64, status int64) ([]*ChallengerCorrectAnswerCountList, error) {
	var list []*ChallengerCorrectAnswerCountList
	option := &options.AggregateOptions{}
	match := bson.M{
		"correct_answer": answer,
	}
	if Challenger != "" {
		match["challenger"] = Challenger
	}
	if status != 0 {
		match["status"] = status
	}
	pipeline := bson.A{
		bson.M{
			"$match": match,
		},
		bson.M{
			"$group": bson.M{
				"_id": "$challenger",
				"count": bson.M{
					"$sum": 1,
				},
				"bet": bson.M{
					"$sum": "$bet",
				},
			},
		},
		bson.M{
			"$sort": bson.M{
				"bet": -1,
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":        0,
				"challenger": "$_id",
				"count":      1,
				"bet":        1,
			},
		},
	}
	res, err := m.conn.Aggregate(ctx, pipeline, option)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	err = res.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return list, err
}

func (m *NpModel) FindContestantMyAnswerCountList(ctx context.Context, Contestant string, answer int64, status int64) ([]*ContestantMyAnswerCountList, error) {
	var list []*ContestantMyAnswerCountList
	option := &options.AggregateOptions{}
	match := bson.M{
		"my_answer": answer,
	}
	if Contestant != "" {
		match["contestant"] = Contestant
	}
	if status != 0 {
		match["status"] = status
	}
	pipeline := bson.A{
		bson.M{
			"$match": match,
		},
		bson.M{
			"$group": bson.M{
				"_id": "$contestant",
				"count": bson.M{
					"$sum": 1,
				},
				"bet": bson.M{
					"$sum": "$bet",
				},
			},
		},
		bson.M{
			"$sort": bson.M{
				"bet": -1,
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":        0,
				"contestant": "$_id",
				"count":      1,
				"bet":        1,
			},
		},
	}
	res, err := m.conn.Aggregate(ctx, pipeline, option)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	err = res.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return list, err
}

// FindDateCountList 获取按日期分组的指定条件数量的列表
func (m *NpModel) FindDateCountList(ctx context.Context, where map[string]interface{}) ([]*DateCountList, error) {
	var list []*DateCountList
	option := &options.AggregateOptions{}
	pipeline := bson.A{
		bson.M{"$match": where},
		bson.M{"$addFields": bson.M{"date": bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": "$question_time", "timezone": "Asia/Shanghai"}}}},
		bson.M{"$group": bson.M{"_id": "$date", "count": bson.M{"$sum": 1}}},
		bson.M{"$sort": bson.M{"_id": 1}},
		bson.M{"$project": bson.M{"_id": 0, "date": "$_id", "count": 1}},
	}
	res, err := m.conn.Aggregate(ctx, pipeline, option)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	err = res.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return list, err
}
