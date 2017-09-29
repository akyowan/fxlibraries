package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"fxlibraries/mysql"
	"fxservice/service/chatcenter/domain"
	"reflect"
	"regexp"
	"time"
	"unicode"

	"github.com/jinzhu/gorm"
)

func init() {

}

var (
	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func BuildSql(values ...interface{}) string {
	var (
		sql             string
		formattedValues []string
	)

	for _, value := range values[1].([]interface{}) {
		indirectValue := reflect.Indirect(reflect.ValueOf(value))
		if indirectValue.IsValid() {
			value = indirectValue.Interface()
			if t, ok := value.(time.Time); ok {
				formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
			} else if b, ok := value.([]byte); ok {
				if str := string(b); isPrintable(str) {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
				} else {
					formattedValues = append(formattedValues, "'<binary>'")
				}
			} else if r, ok := value.(driver.Valuer); ok {
				if value, err := r.Value(); err == nil && value != nil {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			} else {
				formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
			}
		} else {
			formattedValues = append(formattedValues, "NULL")
		}
	}

	// differentiate between $n placeholders or else treat like ?
	if numericPlaceHolderRegexp.MatchString(values[0].(string)) {
		sql = values[0].(string)
		for index, value := range formattedValues {
			placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
			sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
		}
	} else {
		formattedValuesLength := len(formattedValues)
		for index, value := range sqlRegexp.Split(values[0].(string), -1) {
			sql += value
			if index < formattedValuesLength {
				sql += formattedValues[index]
			}
		}
	}
	//
	//		messages = append(messages, sql)
	//		messages = append(messages, fmt.Sprintf(" \n\033[36;31m[%v]\033[0m ", strconv.FormatInt(values[5].(int64), 10)+" rows affected or returned "))
	//	} else {
	//		messages = append(messages, "\033[31;1m")
	//		messages = append(messages, values[2:]...)
	//		messages = append(messages, "\033[0m")
	//	}
	//}
	return sql

}

func AfterUpdate(scope *gorm.Scope) {
	sql := BuildSql(scope.SQL, scope.SQLVars)
	scope.SQL = sql
	fmt.Println("after update")
}

func main() {
	pool := mysql.NewDBPool(mysql.DBPoolConfig{
		Host:     "10.0.0.200",
		Port:     3306,
		User:     "wans",
		Password: "koudai123456",
		DBName:   "chat",
		Debug:    true,
	})
	db := pool.NewConn()
	db.Callback().Update().After("gorm:update").Register("gorm:momo_acocunt_update", AfterUpdate)

	account := domain.MomoAccount{
		ID:       100,
		NickName: "Test",
	}
	if err := db.Model(&account).Update(&account).Error; err != nil {
		fmt.Println(err)
	}
	account.City = "上海"
	db2 := pool.NewConn()
	if err := db2.Model(&account).Update(&account).Error; err != nil {
		fmt.Println(err)
	}
}
