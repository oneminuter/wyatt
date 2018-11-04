package model

import (
	"time"
	"wyatt/db"

	"wyatt/util"

	"wyatt/api/constant"

	"errors"

	_ "github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
)

type TableModel struct {
	ID        int64 `json:"-" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	FlowId    int64      `json:"flowId" sql:"index"` //子流水号，创建时的时间戳， 10为数字
}

func init() {
	mdb := db.GetMysqlDB()
	mdb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&Comment{}, &Community{}, &JoinedCommunity{}, &Message{}, &Topic{}, &Zan{}, &User{},
		&CommunityManager{}, &TopicCollect{})
}

//验证流水号是否合法
func ValidateFlowId(tableName string, tableID int64, timestamp int64) bool {
	mdb := db.GetMysqlDB()
	var c int
	err := mdb.Table(tableName).Where("id = ? AND flow_id = ?", tableID, timestamp).Count(&c).Error
	if err != nil {
		util.LoggerError(err)
		return false
	}
	if 1 > c {
		return false
	}
	return true
}

//根据流水号信息查询记录
func QueryByFlowIdInfo(tableName string, tableID int64, timestamp int64) (map[string]interface{}, error) {
	/*
		mdb := db.GetMysqlDB()
		rows, err := mdb.Table(tableName).Where("id = ? AND flow_id = ?", tableID, timestamp).Rows()
		if err != nil {
			util.LoggerError(err)
			return nil, err
		}
		cols, _ := rows.Columns()

		m := make(map[string]interface{})

		for rows.Next() {
			// Create a slice of interface{}'s to represent each column,
			// and a second slice to contain pointers to each item in the columns slice.
			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i, _ := range columns {
				columnPointers[i] = &columns[i]
			}

			// Scan the result into the column pointers...
			if err := rows.Scan(columnPointers...); err != nil {
				return nil, err
			}

			// Create our map, and retrieve the value for each column from the pointers slice,
			// storing it in the map with the name of the column as the key.

			for i, colName := range cols {
				val := columnPointers[i].(*interface{})
				//log.Println(colName, ":", reflect.TypeOf(*val) == reflect.Slice)

				m[colName] = *val
			}

			// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
			fmt.Println(m)
		}

		return m, nil
	*/
	var (
		result = make(map[string]interface{})
		err    error
	)
	switch tableName {
	case constant.TabelMap[constant.CM]:
		caller := new(Comment)
		err = caller.QueryOne("*", "id = ? AND flow_id = ?", tableID, timestamp)
		if err != nil {
			util.LoggerError(err)
			return nil, err
		}
		result = recordToMap(caller)
	case constant.TabelMap[constant.CMT]:
		caller := new(Comment)
		err = caller.QueryOne("*", "id = ? AND flow_id = ?", tableID, timestamp)
		if err != nil {
			util.LoggerError(err)
			return nil, err
		}
		result = recordToMap(caller)
	case constant.TabelMap[constant.TP]:
		caller := new(Topic)
		err = caller.QueryOne("*", "id = ? AND flow_id = ?", tableID, timestamp)
		if err != nil {
			util.LoggerError(err)
			return nil, err
		}
		result = recordToMap(caller)
	default:
		util.LoggerError(errors.New("tableName 不在匹配选项中"))
		return nil, errors.New("tableName 不在匹配选项中")
	}
	return result, nil
}

func recordToMap(record interface{}) (result map[string]interface{}) {
	bytes, err := jsoniter.Marshal(record)
	if err != nil {
		util.LoggerError(err)
		return
	}
	err = jsoniter.Unmarshal(bytes, &result)
	if err != nil {
		util.LoggerError(err)
		return
	}
	return result
}
