package logic

import (
	"errors"
	"strconv"
	"strings"
	"wyatt/api/model"
	"wyatt/util"
)

func SplitFlowNumber(flow string) (tableName string, TableID int64, timestamp int64, err error) {
	splits := strings.Split(flow, ".")
	if 3 != len(splits) {
		err = errors.New("Incorrect flow number")
		return
	}

	tn := splits[0]
	var ok bool
	tableName, ok = model.TabelMap[tn]
	if !ok {
		err = errors.New("Incorrect flow number: table not exist")
		return
	}

	TID := splits[1]
	TableID, err = strconv.ParseInt(TID, 10, 64)
	if err != nil {
		util.LoggerError(err)
		return
	}

	t := splits[2]
	timestamp, err = strconv.ParseInt(t, 10, 64)
	if err != nil {
		util.LoggerError(err)
		return
	}

	return
}
