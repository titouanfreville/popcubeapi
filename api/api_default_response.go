package api

import "github.com/titouanfreville/popcubeapi/utils"

// Success -------------------------------------------

// generalOk default object style to return if correct
//
// swagger:response defaultOk
type generalOk struct {
	// What you want to say
	// in:body
	Message string `json:"message,omitempty"`
	Status  int    `json:"-"`
}

// ---------------------------------------------------
// Errors --------------------------------------------

// genericError general error when unexpected errors occured
//
// swagger:response genericError
type genericError struct {
	// in:body
	Error utils.AppError
}

// wrongEntityError is an error object to inform that the provided object was not correctly formated
//
// swagger:response wrongEntity
type wrongEntityError struct {
	// in:body
	Error utils.AppError
}

// databaseError is an error object to tell what is happening when we encounter issue with database
//
// swagger:response databaseError
type databaseError struct {
	// in:body
	Error utils.AppError
}

// ---------------------------------------------------
// Unknow --------------------------------------------

// deleteMessageModel is an object to confirm correct deletion of an item.
//
// swagger:model deleteMessageModel
type deleteMessageModel struct {
	// Status
	Status int `json:"status"`
	// Correctly removed ?
	Success bool `json:"success"`
	// More information about why is it or isn't it removed
	Message string `json:"message,omitempty"`
	// The object we where trying to remove
	Object interface{} `json:"removed_object, omitempty"`
}

// deleteMessage return object to confirm correct deletion of an item.
//
// swagger:model deleteMessage
type deleteMessage struct {
	// in:body
	Message deleteMessageModel
}

// ---------------------------------------------------
// Generators ----------------------------------------

func newGeneralOk(message string) generalOk {
	return generalOk{
		Status:  200,
		Message: message,
	}
}

func newDeleteMessage(succes bool, message string) deleteMessageModel {
	return deleteMessageModel{
		Status:  200,
		Message: message,
		Success: succes,
	}
}

// ---------------------------------------------------
