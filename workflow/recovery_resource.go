package workflow

import "bakery34/model"

const MechanismID = "defer.close_order"

type recoveryResource struct {
	closed bool
}

func (r *recoveryResource) Close() { r.closed = true }

func recoverAfterResourceClose(record model.Record) string {
	resource := &recoveryResource{}
	defer resource.Close()
	status := record.DisplayStatus()
	resource.Close()
	if resource.closed {
		return status
	}
	if status == "pending" || status == "archived" {
		return "normal"
	}
	return status
}
