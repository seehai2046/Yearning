package commom

import (
	"Yearning-go/src/lib"
	"Yearning-go/src/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
)

type _dbInfo struct {
	Results   []string
	Query     []map[string]interface{}
	BaseList  []map[string]interface{}
	Highlight []map[string]string
}

func ScanDataRows(s model.CoreDataSource, database, sql, meta string, isQuery bool) (res _dbInfo, err error) {

	ps := lib.Decrypt(s.Password)

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", s.Username, ps, s.IP, strconv.Itoa(int(s.Port)), database))

	defer func() {
		_ = db.Close()
	}()

	var _tmp string

	if err != nil {
		return _dbInfo{}, err
	}

	rows, err := db.Raw(sql).Rows()

	if err != nil {
		return _dbInfo{}, err
	}

	for rows.Next() {
		_ = rows.Scan(&_tmp)
		if isQuery {
			res.Query = append(res.Query, map[string]interface{}{"title": _tmp})
			res.BaseList = append(res.BaseList, map[string]interface{}{"title": _tmp, "children": []map[string]string{{}}})
		} else {
			res.Results = append(res.Results, _tmp)
		}
		res.Highlight = append(res.Highlight, map[string]string{"vl": _tmp, "meta": meta})
	}
	return res, nil
}