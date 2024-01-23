package services

import(
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

