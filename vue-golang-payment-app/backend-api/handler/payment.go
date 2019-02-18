package handler

import (
	"etc/vue-golang-payment-app/backend-api/db"
	"etc/vue-golang-payment-app/backend-api/domain"
	gpay "etc/vue-golang-payment-app/payment-server/proto"

	"google.golang.org/grpc"
)


var addr = "localhost:50051"

func Charge(c Context) {
	t := domain.Payment{}
	c.Bind(&t)
	identifer, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalSercerError, err)
	}

	res, err := db.SelectItem(int64(identifer))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	greq := &gpay.PayRequest{
		Id:		int64(identifer),
		Token:	t.Token,
		Amount:	res.Amount,
		Name:	res.Name,
		Description:	res.Description,
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusForbidden, err)
	}
	defer conn.Close()
	client := gpay.NewPayManagerClient(conn)

	gres, err := client.Charge(context.Background(), greq)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}
	c.JSON(http.StatusOK, gres)