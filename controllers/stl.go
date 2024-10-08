package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"gdx2-beego/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type ResponseJson struct {
	Msg   string `json:"msg"`   // Сериализуется как "msg"
	Count int    `json:"count"` // Сериализуется как "count"
}

// StlController operations for Stl
type StlController struct {
	beego.Controller
}

// URLMapping ...
func (c *StlController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetOneROSG", c.GetOneROSG)

	// c.Mapping("Post", c.Post)
	// c.Mapping("Put", c.Put)
	// c.Mapping("Delete", c.Delete)
}

// GetOne ...
// @Title Get One
// @Description get Stl by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Stl
// @Failure 403 :id is empty
// @router /:id [get]
func (c *StlController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetStlById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One ROSG
// @Description get Stl by id
// @Param	rosg		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Stl
// @Failure 403 :rosg is empty
// @router /ROSG/:rosg [get]
func (c *StlController) GetOneROSG() {
	idStr := c.Ctx.Input.Param(":rosg")
	// id, _ := strconv.Atoi(idStr)
	v, err := models.GetStlByRosg(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetCount ...
// @Title Get Coount ROSG
// @Description get count Stl by ROSG
// @Param	rosg		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Stl
// @Failure 403 :rosg is empty
// @router /COUNT/ROSG/:rosg [get]
func (c *StlController) GetCountSTLROSG() {
	idStr := c.Ctx.Input.Param(":rosg")
	// id, _ := strconv.Atoi(idStr)
	v, err := models.GetStlCountByRosg(idStr)
	// c = strconv.Itoa(v)
	// fmt.Printf("%T, %v\n", c, c)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {

		r := &ResponseJson{Msg: "ok", Count: v}
		// data, err := json.MarshalIndent(r, "", "  ")
		data, err := json.Marshal(r)
		if err != nil {
			fmt.Println("Ошибка записи данных:", err)
		}
		fmt.Println(r)
		c.Data["json"] = string(data)
	}
	c.ServeJSON()
}

// func (c *StlController) GetCountSTLROSG() {
// 	idStr := c.Ctx.Input.Param(":rosg")
// 	// id, _ := strconv.Atoi(idStr)
// 	v, err := models.GetStlCountByRosg(idStr)
// 	if err != nil {
// 		c.Data["json"] = err.Error()
// 	} else {

// 		c.Data["json"] = v
// 	}
// 	c.ServeJSON()
// }

// GetAll ...
// @Title Get All
// @Description get Stl
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Stl
// @Failure 403
// @router / [get]
func (c *StlController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllStl(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Stl
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Stl	true		"body for Stl content"
// @Success 200 {object} models.Stl
// @Failure 403 :id is not int
// @router /:id [put]
// func (c *StlController) Put() {
// 	idStr := c.Ctx.Input.Param(":id")
// 	id, _ := strconv.Atoi(idStr)
// 	v := models.Stl{Id: id}
// 	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
// 		if err := models.UpdateStlById(&v); err == nil {
// 			c.Data["json"] = "OK"
// 		} else {
// 			c.Data["json"] = err.Error()
// 		}
// 	} else {
// 		c.Data["json"] = err.Error()
// 	}
// 	c.ServeJSON()
// }

// Delete ...
// @Title Delete
// @Description delete the Stl
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
// func (c *StlController) Delete() {
// 	idStr := c.Ctx.Input.Param(":id")
// 	id, _ := strconv.Atoi(idStr)
// 	if err := models.DeleteStl(id); err == nil {
// 		c.Data["json"] = "OK"
// 	} else {
// 		c.Data["json"] = err.Error()
// 	}
// 	c.ServeJSON()
// }

// Post ...
// @Title Post
// @Description create Stl
// @Param	body		body 	models.Stl	true		"body for Stl content"
// @Success 201 {int} models.Stl
// @Failure 403 body is empty
// @router / [post]
// func (c *StlController) Post() {
// 	var v models.Stl
// 	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
// 		if _, err := models.AddStl(&v); err == nil {
// 			c.Ctx.Output.SetStatus(201)
// 			c.Data["json"] = v
// 		} else {
// 			c.Data["json"] = err.Error()
// 		}
// 	} else {
// 		c.Data["json"] = err.Error()
// 	}
// 	c.ServeJSON()
// }
