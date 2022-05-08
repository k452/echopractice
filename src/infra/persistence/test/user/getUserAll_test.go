package persistence

import (
	"reflect"
	"sake_io_api/domain/model"
	"sake_io_api/infra/persistence"
	"sake_io_api/utils"
	"sort"
	"testing"
)

func TestGetUserAll(t *testing.T) {
	t.Setenv("GO_ENV", "test")
	up := persistence.NewUserPersistence(false)
	got := up.GetUserAll()
	want1 := []model.User{
		{
			1,
			"イワーク",
			"iwaku@example.com",
			1,
			"q",
			utils.StringToTime("2020-05-02T16:01:09Z"),
			utils.StringToTime("2020-05-02T16:01:09Z"),
		},
		{
			2,
			"サイドン",
			"saidon@example.com",
			2,
			"w",
			utils.StringToTime("2020-06-02T16:01:09Z"),
			utils.StringToTime("2020-06-02T16:01:09Z"),
		},
		{
			3,
			"ゴローニャ",
			"goronya@example.com",
			7,
			"e",
			utils.StringToTime("2020-11-03T16:01:09Z"),
			utils.StringToTime("2020-11-03T16:01:09Z"),
		},
		{
			4,
			"ラムパルド",
			"rampardo@example.com",
			197,
			"r",
			utils.StringToTime("2021-01-07T16:01:09Z"),
			utils.StringToTime("2021-01-07T16:01:09Z"),
		},
		{
			5,
			"ジーランス",
			"zirans@example.com",
			30,
			"t",
			utils.StringToTime("2021-05-23T16:01:09Z"),
			utils.StringToTime("2021-05-23T16:01:09Z"),
		},
	}
	want2 := []model.User{}

	tests := []struct {
		name string
		want []model.User
	}{
		{
			name: "userテーブル全データ取得が成功",
			want: want1,
		},
		{
			name: "userテーブル全データ取得が失敗",
			want: want2,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == 0 {
				if !reflect.DeepEqual(sortUserAll(got), sortUserAll(tt.want)) {
					t.Error("error")
				}
			} else if i == 1 {
				if reflect.DeepEqual(sortUserAll(got), sortUserAll(tt.want)) {
					t.Error("error")
				}
			}
		})
	}
}

func sortUserAll(userSlice []model.User) []model.User {
	sort.Slice(userSlice, func(i, j int) bool {
		return userSlice[i].User_id < userSlice[j].User_id
	})
	return userSlice
}
