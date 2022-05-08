package persistence

import (
	"reflect"
	"sake_io_api/domain/model"
	"sake_io_api/infra/persistence"
	"sake_io_api/utils"
	"sort"
	"testing"
)

func TestGetManagementAll(t *testing.T) {
	t.Setenv("GO_ENV", "test")
	up := persistence.NewManagementPersistence(false)
	got := up.GetManagementAll()
	want1 := []model.Management{
		{
			1,
			1,
			"グレンリベット12年",
			100,
			utils.StringToTime("2020-04-01T00:00:00Z"),
			utils.StringToTime("2020-05-02T16:01:09Z"),
			utils.StringToTime("2020-05-02T16:01:09Z"),
		}, {
			2,
			1,
			"🍺",
			1000,
			utils.StringToTime("2020-05-02T00:00:00Z"),
			utils.StringToTime("2020-06-02T16:01:09Z"),
			utils.StringToTime("2020-06-02T16:01:09Z"),
		},
		{
			3,
			2,
			"黒霧島",
			700,
			utils.StringToTime("2020-01-01T00:00:00Z"),
			utils.StringToTime("2020-11-03T16:01:09Z"),
			utils.StringToTime("2020-11-03T16:01:09Z"),
		},
		{
			4,
			3,
			"浦霞",
			197,
			utils.StringToTime("2020-06-11T00:00:00Z"),
			utils.StringToTime("2021-01-07T16:01:09Z"),
			utils.StringToTime("2021-01-07T16:01:09Z"),
		},
		{
			5,
			4,
			"バドワイザー",
			300,
			utils.StringToTime("2021-01-02T00:00:00Z"),
			utils.StringToTime("2021-05-23T16:01:09Z"),
			utils.StringToTime("2021-05-23T16:01:09Z"),
		},
	}
	want2 := []model.Management{}

	tests := []struct {
		name string
		want []model.Management
	}{
		{
			name: "managementテーブル全データ取得が成功",
			want: want1,
		},
		{
			name: "managementテーブル全データ取得が失敗",
			want: want2,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == 0 {
				if !reflect.DeepEqual(sortManagementAll(got), sortManagementAll(tt.want)) {
					t.Error("error")
				}
			} else if i == 1 {
				if reflect.DeepEqual(sortManagementAll(got), sortManagementAll(tt.want)) {
					t.Error("error")
				}
			}
		})
	}
}

func sortManagementAll(managementSlice []model.Management) []model.Management {
	sort.Slice(managementSlice, func(i, j int) bool {
		return managementSlice[i].Management_id < managementSlice[j].Management_id
	})
	return managementSlice
}
