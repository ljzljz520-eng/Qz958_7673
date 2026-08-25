package model

var statuses = []string{"pending", "processing", "approved", "archived", "cancelled"}

func Statuses() []string { return append([]string(nil), statuses...) }
func IsStatus(s string) bool {
	for _, v := range statuses {
		if s == v {
			return true
		}
	}
	return false
}
func CanTransition(from, to string) bool {
	if !IsStatus(from) || !IsStatus(to) || from == to {
		return false
	}
	switch from {
	case "pending":
		return to == "processing" || to == "cancelled"
	case "processing":
		return to == "approved" || to == "cancelled"
	case "approved":
		return to == "archived"
	}
	return false
}
func NormalizeStatus(s string) string {
	if !IsStatus(s) {
		return "pending"
	}
	return s
}
func NextStatus(s string) string {
	switch s {
	case "pending":
		return "processing"
	case "processing":
		return "approved"
	case "approved":
		return "archived"
	}
	return s
}
