package services

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/seccret404/BE-payment-gateway-Go-Fiber/models"
	"gorm.io/gorm"
)

func CreatePaymentService(db *gorm.DB, req models.PaymentRequest) (*models.Payment, error) {
    amount := int64(req.Price * float64(req.Quantity))


    // Generate order ID
    orderID := fmt.Sprintf("ORDER-%s", uuid.New().String())

    //snap 
    var snapClient snap.Client
    snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), getMidtransEnv())

    //request snap
    snapReq := &snap.Request{
        TransactionDetails: midtrans.TransactionDetails{
            OrderID:  orderID,
            GrossAmt: amount,
        },
        CustomerDetail: &midtrans.CustomerDetails{
            FName: req.CustomerName,
            Email: req.CustomerEmail,
        },
    }

    //request ke midtrans
    snapResp, err := snapClient.CreateTransaction(snapReq)
    if err != nil {
        return nil, fmt.Errorf("midtrans error: %v", err)
    }

    //simpan data ke DB
    payment := &models.Payment{
        OrderID:       orderID,
        ProductID:     req.ProductID,
        Quantity:      req.Quantity,
        CustomerName:  req.CustomerName,
        CustomerEmail: req.CustomerEmail,
        Amount:        amount,
        Status:        "pending",
        SnapURL:       snapResp.RedirectURL,
        CreatedAt:     time.Now(),
        UpdatedAt:     time.Now(),
    }

    if err := db.Create(payment).Error; err != nil {
        return nil, fmt.Errorf("db error: %v", err)
    }

    return payment, nil
}

func getMidtransEnv() midtrans.EnvironmentType {
	if os.Getenv("MIDTRANS_IS_PRODUCTION") == "true" {
		return midtrans.Production
	}
	return midtrans.Sandbox
}
