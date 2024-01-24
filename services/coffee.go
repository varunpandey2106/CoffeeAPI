package services

import (
	"context"
	"time"
)

type Coffee struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Roast string `json:"roast"`
    Image string `json:"image"`
    Region string `json:"region"`
    Price float32 `json:"price"`
    GrindUnit int16 `json:"grind_unit"` 
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (c*Coffee) GetAllCoffees()([]*Coffee, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

    query := `SELECT id, name, image, roast,region,price, grind_unit, created_at, updated_at FROM coffees`
    rows, err:=db.QueryContext(ctx, query)

    if err != nil{
        return nil, err
    }
}