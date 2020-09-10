package preseed

import (
	"time"
)

// Data available when rendering the preseed file
type Data struct {
	Hostname   string
	RenderedAt time.Time
}
