package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Stl struct {
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

func (t *Stl) TableName() string {
	return "stl"
}

func init() {
	orm.RegisterModel(new(Stl))
}

// AddStl insert a new Stl into database and returns
// last inserted Id on success.
// func AddStl(m *Stl) (id int64, err error) {
// 	o := orm.NewOrm()
// 	id, err = o.Insert(m)
// 	return
// }

// GetStlById retrieves Stl by Id. Returns error if
// Id doesn't exist
func GetStlById(id int) (v *Stl, err error) {
	o := orm.NewOrm()
	v = &Stl{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetStlByRosg(rosg string) (dataList []Stl, err error) {
	// list encapsulated data
	var list []Stl
	o := orm.NewOrm()
	nums, err := o.Raw("SELECT * FROM stl WHERE in_n_rosg = ?", rosg).QueryRows(&list)
	if nums == 0 {

	}
	return list, err
}

// sql := fmt.Sprintf("SELECT %sid%s,%sname%s FROM %suser%s WHERE id = ?",Q,Q,Q,Q,Q,Q)
// rs := Ormer.Raw(sql, 1)

// func GetStlByRosg(rosg string) (v *Stl, err error) {
// 	o := orm.NewOrm()
// 	v = &Stl{InNRosg: rosg}
// 	qb, _ := orm.NewQueryBuilder("postgres")
// 	qb.Select("in_n_rosg").From("stl").Where("in_n_rosg == rosg")

// 	// query = fmt.Sprintf("SELECT 'id','in_n_rosg' FROM %sstl%s", Q, Q)
// 	// num, err = dORM.Raw(query).QueryRows(&ids,&names) // ids=>{1,2},names=>{"nobody","slene"}

// 	// row, err := o.Raw(qb.String(), 1).QueryRows(v)
// 	// qb, _ := orm.NewQueryBuilder("postgres")
// 	// qb.Select("id", "in_n_rosg").From("stl").Where("in_n_rosg == rosg").OrderBy("id").Desc().Limit(10)
// 	// orm.NewOrm().Raw(qb.String(), 1).QueryRows(v)

// 	fmt.Println(rosg)
// 	fmt.Println(v)
// 	if err = o.Raw(qb.String(), 1).QueryRows(v); err == nil {
// 		return v, nil
// 	}

// 	// if err == nil {
// 	// 	return fmt.Print("Error: " + err.Error())
// 	// }
// 	// if err = o.Read(v); err == nil {
// 	// 	return v, nil
// 	// }
// 	// return nil, err
// 	return nil, err
// }

// GetAllStl retrieves all Stl matches certain condition. Returns empty list if
// no records exist
func GetAllStl(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Stl))
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

	var l []Stl
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

// UpdateStl updates Stl by Id and returns error if
// the record to be updated doesn't exist
// func UpdateStlById(m *Stl) (err error) {
// 	o := orm.NewOrm()
// 	v := Stl{Id: m.Id}
// 	// ascertain id exists in the database
// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Update(m); err == nil {
// 			fmt.Println("Number of records updated in database:", num)
// 		}
// 	}
// 	return
// }

// DeleteStl deletes Stl by Id and returns error if
// the record to be deleted doesn't exist
// func DeleteStl(id int) (err error) {
// 	o := orm.NewOrm()
// 	v := Stl{Id: id}
// 	// ascertain id exists in the database
// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Delete(&Stl{Id: id}); err == nil {
// 			fmt.Println("Number of records deleted in database:", num)
// 		}
// 	}
// 	return
// }
