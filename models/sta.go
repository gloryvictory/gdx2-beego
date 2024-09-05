package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Sta struct {
	Id int `orm:"column(id);pk"`
	// Geom      string `orm:"column(geom);null"`
	WebUkId   string `orm:"column(web_uk_id);null"`
	VidIz     string `orm:"column(vid_iz);null"`
	Tgf       string `orm:"column(tgf);null"`
	NUkTgf    string `orm:"column(n_uk_tgf);null"`
	NUkRosg   string `orm:"column(n_uk_rosg);null"`
	NameOtch  string `orm:"column(name_otch);null"`
	NameOtch1 string `orm:"column(name_otch1);null"`
	Avts      string `orm:"column(avts);null"`
	GodNach   string `orm:"column(god_nach);null"`
	GodEnd    string `orm:"column(god_end);null"`
	OrgIsp    string `orm:"column(org_isp);null"`
	InNTgf    string `orm:"column(in_n_tgf);null"`
	InNRosg   string `orm:"column(in_n_rosg);null"`
	Nom1000   string `orm:"column(nom_1000);null"`
	Method    string `orm:"column(method);null"`
	Scale     string `orm:"column(scale);null"`
}

func (t *Sta) TableName() string {
	return "sta"
}

func init() {
	orm.RegisterModel(new(Sta))
}

// AddSta insert a new Sta into database and returns
// last inserted Id on success.
func AddSta(m *Sta) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetStaById retrieves Sta by Id. Returns error if
// Id doesn't exist
func GetStaById(id int) (v *Sta, err error) {
	o := orm.NewOrm()
	v = &Sta{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSta retrieves all Sta matches certain condition. Returns empty list if
// no records exist
func GetAllSta(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Sta))
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

	var l []Sta
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

// UpdateSta updates Sta by Id and returns error if
// the record to be updated doesn't exist
// func UpdateStaById(m *Sta) (err error) {
// 	o := orm.NewOrm()
// 	v := Sta{Id: m.Id}
// 	// ascertain id exists in the database
// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Update(m); err == nil {
// 			fmt.Println("Number of records updated in database:", num)
// 		}
// 	}
// 	return
// }

// DeleteSta deletes Sta by Id and returns error if
// the record to be deleted doesn't exist
// func DeleteSta(id int) (err error) {
// 	o := orm.NewOrm()
// 	v := Sta{Id: id}
// 	// ascertain id exists in the database
// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Delete(&Sta{Id: id}); err == nil {
// 			fmt.Println("Number of records deleted in database:", num)
// 		}
// 	}
// 	return
// }
