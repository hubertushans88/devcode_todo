package activity

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hubertushans88/devcode_todo/controllers"
	"github.com/hubertushans88/devcode_todo/models"
	"gorm.io/gorm"
	"time"
)

var activityCache []models.Activity
var cnt =1

var ReadAll = func(c *fiber.Ctx) error{
	if activityCache==nil{
		var activities []models.Activity
		models.GetDB().Find(&activities)
		activityCache=activities
	}

	return controllers.SendResponse(c, 200, "Success","Success", activityCache)
}

var ReadOne = func(c *fiber.Ctx) error{
	id := c.Params("id")
	var activity []models.Activity
	models.GetDB().First(&activity, id)
	if len(activity) == 0{
		return controllers.SendResponse(c, 404, "Not Found", "Activity with ID "+id+" Not Found", map[string]string{})
	}
	return controllers.SendResponse(c, 200, "Success","Success", activity[0])
}

type createRequest struct {
	Title *string `json:"title"`
	Email *string `json:"email"`
}

var Create = func(c *fiber.Ctx) error{
	var req createRequest

	if err := c.BodyParser(&req); err != nil {
		return controllers.SendInternalError(c)
	}
	if req.Title == nil {
		return controllers.SendResponse(c, 400, "Bad Request", "title cannot be null", map[string]string{})
	}

	activity := &models.Activity{
		ID: 	   uint(cnt),
		Title:     *req.Title,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	cnt++
	go 	models.GetDB().Create(&activity)

	activityCache=nil
	return controllers.SendResponse(c, 201, "Success", "Success", activity)
}

var Update = func(c *fiber.Ctx)error {
	id := c.Params("id")
	var req createRequest
	if err := c.BodyParser(&req); err != nil {
		return controllers.SendInternalError(c)
	}
	if req.Title == nil {
		return controllers.SendResponse(c, 400, "Bad Request", "title cannot be null", map[string]string{})
	}

	var activity models.Activity
	q:= models.GetDB().First(&activity, id)
	if q.Error== gorm.ErrRecordNotFound{
		return controllers.SendResponse(c, 404, "Not Found", "Activity with ID "+id+" Not Found", map[string]string{})
	}
	activity.Title = *req.Title
	go models.GetDB().Save(activity)
	activityCache=nil
	return controllers.SendResponse(c,200,"Success","Success", &activity)
}

var Delete = func(c *fiber.Ctx)error {
	id := c.Params( "id")
	var activity models.Activity

	q:= models.GetDB().First(&activity, id)
	if q.Error == gorm.ErrRecordNotFound{
		return controllers.SendResponse(c, 404, "Not Found", "Activity with ID "+id+" Not Found", map[string]string{})
	}

	activityCache=nil
	go models.GetDB().Delete(&activity, id)
	return controllers.SendResponse(c, 200, "Success", "Success", map[string]string{})
}