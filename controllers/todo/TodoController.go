package todo

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/hubertushans88/devcode_todo/controllers"
	"github.com/hubertushans88/devcode_todo/models"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var s2b = map[string]bool{"true": true, "1": true, "false": false, "0": false}
var todoCache = map[string][]models.Todo{}
var cnt = 1

var ReadAll = func(c *fiber.Ctx) error {
	activityId := c.Query("activity_group_id")

	qry := "all"
	if activityId != "" {
		qry = activityId
	}

	if todoCache[qry] == nil {
		var todos []models.Todo
		if activityId == "" {
			models.GetDB().Find(&todos, "is_active=true")
		} else {
			models.GetDB().Find(&todos, "is_active=true AND activity_group_id=?", activityId)
		}

		if len(todos) == 0 {
			return controllers.SendResponse(c, 200, "Success", "Success", []string{})
		}
		todoCache[qry] = todos
	}

	return controllers.SendResponse(c, 200, "Success", "Success", todoCache[qry])
}

var ReadOne = func(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	q := models.GetDB().First(&todo, id)
	if q.Error == gorm.ErrRecordNotFound {
		return controllers.SendResponse(c, 404, "Not Found", "Todo with ID "+id+" Not Found", map[string]string{})
	}
	return controllers.SendResponse(c, 200, "Success", "Success", todo)
}

type createTodoRequest struct {
	Title           *string `json:"title"`
	ActivityGroupID *uint   `json:"activity_group_id"`
}

var Create = func(c *fiber.Ctx) error {
	var req createTodoRequest

	if err := c.BodyParser(&req); err != nil {
		return controllers.SendInternalError(c)
	}
	if req.Title == nil {
		return controllers.SendResponse(c, 400, "Bad Request", "title cannot be null", map[string]string{})
	}
	if req.ActivityGroupID == nil {
		return controllers.SendResponse(c, 400, "Bad Request", "activity_group_id cannot be null", map[string]string{})
	}
	id := fmt.Sprint(req.ActivityGroupID)

	now := time.Now()
	todo := &models.Todo{
		//ID:              uint(cnt),
		Title:           *req.Title,
		ActivityGroupID: *req.ActivityGroupID,
	}

	//req["id"] = todo.ID
	//req["priority"] = "very-high"
	//req["activity_group_id"] = activityID
	//req["is_active"] = true
	//req["created_at"] = todo.CreatedAt
	//req["updated_at"] = todo.UpdatedAt

	if cnt == 1 || cnt >= 800 { //600 working
		//if cnt%4 != 0 {
		models.GetDB().Create(&todo)
	} else {
		x := todo
		go func(data *models.Todo) {
			models.GetDB().Create(&data)
		}(x)
		todo.CreatedAt = now
		todo.UpdatedAt = now
		todo.ID = uint(cnt)
	}
	//models.GetDB().Create(&todo)
	cnt++
	todoCache["all"] = nil
	todoCache[id] = nil

	return controllers.SendResponse(c, 201, "Success", "Success", todo)
}

var Update = func(c *fiber.Ctx) error {
	id := c.Params("id")
	nid := id
	var req map[string]interface{}
	if err := c.BodyParser(&req); err != nil {
		return controllers.SendInternalError(c)
	}
	if req["title"] == nil && req["activity_group_id"] == nil && req["is_active"] == nil {
		return controllers.SendResponse(c, 400, "Bad Request", "title cannot be null", map[string]string{})
	}

	var todo models.Todo
	q := models.GetDB().First(&todo, id)
	if q.Error == gorm.ErrRecordNotFound {
		return controllers.SendResponse(c, 404, "Not Found", "Todo with ID "+id+" Not Found", map[string]string{})
	}

	todoCache[fmt.Sprint(todo.ActivityGroupID)] = nil

	if req["title"] != nil {
		todo.Title = fmt.Sprint(req["title"])
	}
	if req["is_active"] != nil {

		todo.IsActive = s2b[fmt.Sprint(req["is_active"])]
	}
	if req["activity_group_id"] != nil {
		nid = fmt.Sprint(req["activity_group_id"])
		nID, _ := strconv.Atoi(id)
		todo.ActivityGroupID = uint(nID)
	}
	go models.GetDB().Save(todo)
	todoCache["all"] = nil
	todoCache[nid] = nil
	return controllers.SendResponse(c, 200, "Success", "Success", &todo)
}

var Delete = func(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo

	q := models.GetDB().First(&todo, id)
	if q.Error == gorm.ErrRecordNotFound {
		return controllers.SendResponse(c, 404, "Not Found", "Todo with ID "+id+" Not Found", map[string]string{})
	}

	todoCache["all"] = nil
	todoCache[fmt.Sprint(todo.ActivityGroupID)] = nil
	go models.GetDB().Delete(&todo, id)
	return controllers.SendResponse(c, 200, "Success", "Success", map[string]string{})
}
