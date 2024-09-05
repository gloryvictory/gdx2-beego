package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Lu struct {
	Id         int       `orm:"column(id);pk"`
	Geom       string    `orm:"column(geom);null"`
	Areaoil    float64   `orm:"column(areaoil);null"`
	AreaLic    string    `orm:"column(area_lic);null"`
	Year       int64     `orm:"column(year);null"`
	NomZsngp   int64     `orm:"column(nom_zsngp);null"`
	NomList    string    `orm:"column(nom_list);null"`
	Nom        int64     `orm:"column(nom);null"`
	DataStart  time.Time `orm:"column(data_start);type(date);null"`
	DataEnd    time.Time `orm:"column(data_end);type(date);null"`
	Vid        string    `orm:"column(vid);null"`
	Ftype      string    `orm:"column(ftype);null"`
	NameRus    string    `orm:"column(name_rus);null"`
	Anumber    string    `orm:"column(anumber);null"`
	Sostiyanie string    `orm:"column(sostiyanie);null"`
	Priznak    string    `orm:"column(priznak);null"`
	NomLic     string    `orm:"column(nom_lic);null"`
	HeadNedro  string    `orm:"column(head_nedro);null"`
	Oblast     string    `orm:"column(oblast);null"`
	Zngp       string    `orm:"column(zngp);null"`
	Nedropolz  string    `orm:"column(nedropolz);null"`
	Nedropol   string    `orm:"column(nedropol);null"`
	NomUrfo    int64     `orm:"column(nom_urfo);null"`
	Authority  string    `orm:"column(authority);null"`
}

func (t *Lu) TableName() string {
	return "lu"
}

func init() {
	orm.RegisterModel(new(Lu))
}

// AddLu insert a new Lu into database and returns
// last inserted Id on success.
func AddLu(m *Lu) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLuById retrieves Lu by Id. Returns error if
// Id doesn't exist
func GetLuById(id int) (v *Lu, err error) {
	o := orm.NewOrm()
	v = &Lu{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLu retrieves all Lu matches certain condition. Returns empty list if
// no records exist
func GetAllLu(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Lu))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Lu
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateLu updates Lu by Id and returns error if
// the record to be updated doesn't exist
func UpdateLuById(m *Lu) (err error) {
	o := orm.NewOrm()
	v := Lu{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLu deletes Lu by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLu(id int) (err error) {
	o := orm.NewOrm()
	v := Lu{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Lu{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
