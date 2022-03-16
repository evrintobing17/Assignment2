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

// @title           Orders API
// @version         1
// @description     This is API collection to hit Orders API
// @termsOfService  http://swagger.io/terms/
// @contact.name    Evrin Lumbantobing
// @host            localhost:8081
// @BasePath        /docs
func NewHttpDelivery(r *gin.Engine, orderUC order.OrderUsecase) {
	handler := HttpDelivery{
		orderUC: orderUC,
	}

	r.POST("/orders", handler.addOrder)
	r.GET("/orders", handler.getAllOrder)
	r.DELETE("/orders/:orderId", handler.deleteOrder)
	r.PUT("/orders/:orderId", handler.updateOrder)
}

// Add Order godoc
// @Summary      Add new order
// @Description  Add new order
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param  		 order body orderdelivery.Orders true "Create Order"
// @Success      200  {string}  orderdelivery.InsertOrder
// @Router       /orders [post]
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

// GetAllOrder godoc
// @Summary      Show all order
// @Description  Show all order
// @Tags         orders
// @Accept       json
// @Produce      json
// @Success      200  {object}  orderdelivery.GetAllOrderResponse
// @Router       /orders [get]
func (handler *HttpDelivery) getAllOrder(c *gin.Context) {

	order, err := handler.orderUC.GetAllOrder()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
	return
}

// DeleteOrder godoc
// @Summary      delete order
// @Description  delete order by id
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param id path integer true "orderId"
// @Success      200  {string}  success delete
// @Router       /orders/{orderId} [delete]
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

// UpdateOrder godoc
// @Summary      Update Order
// @Description  Update Order Data
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param id path integer true "orderId"
// @Param  		 order body orderdelivery.UpdateOrder true "Update Order"
// @Success      200  {string}  orderdelivery.UpdateOrder
// @Router       /orders/{orderId} [put]
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
