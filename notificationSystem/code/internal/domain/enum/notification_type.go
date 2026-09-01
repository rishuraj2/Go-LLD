package enum

type NotificationType int

const (
	EMAIL NotificationType = iota
	SMS
	PUSH
)

func (this NotificationType) String() string {
	val := []string{"email", "sms", "push"}

	if int(this) < len(val) {
		return val[this]
	}

	return "unknown"
}
