package controllers

import (
	"gin_todo/forms"
	"gin_todo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)



func AddTodoHandler(c *gin.Context) {
	var m forms.PostTodo
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, badReq)
		return
	}

	t, err := models.AddTodo(&m, c.GetInt64("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, badReq)
		return
	}
	c.JSON(http.StatusOK, GenJson(&t))
}

func GetTodo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, badReq)
		return
	}
	m, err := models.GetTodo(id)

	if m.UserId != c.GetInt64("UserId") {
		c.JSON(http.StatusForbidden, err403)
		return
	}
	c.JSON(http.StatusOK, GenJson(m))
}

func UpdateTodo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, badReq)
		return
	}
	var m forms.PostTodo
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, badReq)
		return
	}

	t, err := models.GetTodo(id)

	uid := c.GetInt64("UserId")

	if t.UserId == uid {
		if err = models.UpdateTodoById(t, &m); err != nil {
			c.JSON(http.StatusBadRequest, badReq)
			return
		}
		c.JSON(http.StatusOK, GenJson(t))
	} else {
		c.JSON(http.StatusForbidden, err403)
	}
}

func DeleteTodo(c *gin.Context)  {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, badReq)
		return
	}


	t, err := models.GetTodo(id)

	uid := c.GetInt64("UserId")

	if t.UserId == uid {
		if err = models.DeleteTodo(id); err != nil {
			c.JSON(http.StatusBadRequest, badReq)
			return
		}
		SuccessJson(c)
	} else {
		c.JSON(http.StatusForbidden, err403)
	}
}

func AllTodo(c *gin.Context) {
	var fields []string
	var sortby []string
	var query = make(map[string]interface{})
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.Query("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)

	// offset: 0 (default is 0)
	offset, err = strconv.ParseInt(c.DefaultQuery("offset", "0"), 10, 64)

	// sortby: col1,col2
	if v := c.Query("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}

	if v := c.Query("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.JSON(http.StatusBadRequest, badReq)
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	uid := c.GetInt64("UserId")

	l, total, err := models.GetAllTodo(query, fields, sortby, offset, limit, uid)

	if err != nil {
		c.JSON(http.StatusBadRequest, badReq)
		return
	}

	c.JSON(http.StatusOK, GenJson(map[string](interface{}){"data": l, "count": total, "limit": limit, "offset": offset}))

}
