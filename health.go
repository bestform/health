package health

const (
	// Up Convenient constant value representing up state.
	Up = "up"

	// Down Convenient constant value representing down state.
	Down = "down"

	// OutOfService Convenient constant value representing out-of-service state.
	OutOfService = "outOfService"

	// Unknown Convenient constant value representing unknown state.
	Unknown = "unknown"
)

// Health is a health status struct
type Health struct {
	Status string      `json:"status"`
	Info   interface{} `json:"info,omitempty"`
}

// NewHealth return a new Health with status Down
func NewHealth() Health {
	h := Health{}
	h.Down()

	return h
}

// IsUnknown returns true if Status is Unknown
func (h Health) IsUnknown() bool {
	return h.Status == Unknown
}

// IsUp returns true if Status is Up
func (h Health) IsUp() bool {
	return h.Status == Up
}

// IsDown returns true if Status is Down
func (h Health) IsDown() bool {
	return h.Status == Down
}

// IsOutOfService returns true if Status is IsOutOfService
func (h Health) IsOutOfService() bool {
	return h.Status == OutOfService
}

// Down set the status to Down
func (h *Health) Down() {
	h.Status = Down
}

// OutOfService set the status to OutOfService
func (h *Health) OutOfService() {
	h.Status = OutOfService
}

// Unknown set the status to Unknown
func (h *Health) Unknown() {
	h.Status = Unknown
}

// Up set the status to Up
func (h *Health) Up() {
	h.Status = Up
}