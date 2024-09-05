package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Field struct {
	Id int `orm:"column(id);pk"`
	// Geom      string  `orm:"column(geom);null"`
	Year      int64   `orm:"column(year);null"`
	Tip       string  `orm:"column(tip);null"`
	Areaoil   float64 `orm:"column(areaoil);null"`
	Nom       int64   `orm:"column(nom);null"`
	Oil       string  `orm:"column(oil);null"`
	Gas       string  `orm:"column(gas);null"`
	Condensat string  `orm:"column(condensat);null"`
	NameRu    string  `orm:"column(name_ru);null"`
	Oblast    string  `orm:"column(oblast);null"`
	Stadia    string  `orm:"column(stadia);null"`
	Note      string  `orm:"column(note);null"`
	Istochnik string  `orm:"column(istochnik);null"`
	Ftype     string  `orm:"column(ftype);null"`
}

func (t *Field) TableName() string {
	return "field"
}

func init() {
	orm.RegisterModel(new(Field))
}

// AddField insert a new Field into database and returns
// last inserted Id on success.
func AddField(m *Field) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFieldById retrieves Field by Id. Returns error if
// Id doesn't exist
func GetFieldById(id int) (v *Field, err error) {
	o := orm.NewOrm()
	v = &Field{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllField retrieves all Field matches certain condition. Returns empty list if
// no records exist
func GetAllField(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Field))
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

	var l []Field
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

// UpdateField updates Field by Id and returns error if
// the record to be updated doesn't exist
// func UpdateFieldById(m *Field) (err error) {
// 	o := orm.NewOrm()
// 	v := Field{Id: m.Id}
// 	// ascertain id exists in the database
// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Update(m); err == nil {
// 			fmt.Println("Number of records updated in database:", num)
// 		}
// 	}
// 	return
// }

// DeleteField deletes Field by Id and returns error if
// the record to be deleted doesn't exist
// func DeleteField(id int) (err error) {
// 	o := orm.NewOrm()
// 	v := Field{Id: id}
// 	// ascertain id exists in the database
// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Delete(&Field{Id: id}); err == nil {
// 			fmt.Println("Number of records deleted in database:", num)
// 		}
// 	}
// 	return
// }
