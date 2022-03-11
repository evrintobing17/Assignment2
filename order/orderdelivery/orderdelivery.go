package orderdelivery

import (
	"encoding/json"
	"net/http"

	"github.com/evrintobing17/order"
	"github.com/gin-gonic/gin"
)

type HttpDelivery struct {
	orderUC order.OrderUsecase
}

func NewHttpDelivery(r *gin.Engine, orderUC order.OrderUsecase) {
	handler := HttpDelivery{
		orderUC: orderUC,
	}

	r.POST("/orders", handler.addOrder)
	r.GET("/orders", handler.getAllOrder)
	r.DELETE("/orders/:orderId", handler.deleteOrder)
	r.PUT("/orders/:orderId", handler.updateOrder)
}

func (handler *HttpDelivery) addOrder(c *gin.Context) {
	var request InsertOrders

	errBind := c.ShouldBind(&request)

	if errBind != nil {
		c.JSON(http.StatusBadRequest, errBind)
		return
	}

	order, err := handler.orderUC.CreateOrder(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
	return
}

func (handler *HttpDelivery) getAllOrder(c *gin.Context) {

	order, err := handler.orderUC.GetAllOrder()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
	return
}

func (handler *HttpDelivery) deleteOrder(c *gin.Context) {
	orderId := c.Param("orderId")

	_, err := handler.orderUC.DeleteOrder(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	response := "Sucess Delete Data"

	c.JSON(http.StatusOK, response)
	return
}

func (handler *HttpDelivery) updateOrder(c *gin.Context) {
	var request UpdateOrder
	orderId := c.Param("orderId")

	errBind := c.ShouldBind(&request)

	if errBind != nil {
		c.JSON(http.StatusBadRequest, errBind)
		return
	}

	updateOrder, err := toMap(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	updateOrders, err := handler.orderUC.UpdateOrderByID(orderId, updateOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateOrders)
	return
}

func toMap(fromStruct interface{}) (map[string]interface{}, error) {

	var inInterface map[string]interface{}

	inrec, err := json.Marshal(&fromStruct)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(inrec, &inInterface)
	if err != nil {
		return nil, err
	}

	return inInterface, nil
}
