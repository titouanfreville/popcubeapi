package api

// Success -------------------------------------------

// generalOk default object style to return if correct
//
// swagger:response defaultOk
type generalOk struct {
    Status int `json:"-"`
    // What you want to say
    Message string `json:"message,omitempty"`
} 

// ---------------------------------------------------
// Errors --------------------------------------------

// wrongEntityError is an error object to inform that the provided object was not correctly formated
//
// swagger:response wrongEntity
type wrongEntityError struct {
	// Status code
	Status int `json:"status"`
	// Location of the error
	Where string `json:"where,omitempty"`
	// Message for what get wrong
	Message string `json:"message,omitempty"`
	// Culprit
	Object interface{} `json:"object"`
}

// databaseError is an error object to tell what is happening when we encounter issue with database
//
// swagger:response databaseError
type databaseError struct {
	// Status code
	Status int `json:"status"`
	// Location of the error
	Where string `json:"where,omitempty"`
	// Message for what get wrong
	Message string `json:"message,omitempty"`
}

// ---------------------------------------------------
// Unknow --------------------------------------------

// deleteMessage return object to confirm correct deletion of an item.
//
// swagger:response deleteMessage
type deleteMessage struct {
	// Status
	Status int `json:"status"`
	// Correctly removed ?
	Success bool `json:"success"`
	// More information about why is it or isn't it removed
	Message string `json:"message,omitempty"`
	// The object we where trying to remove
	Object interface{} `json:"removed_object, omitempty"`
}
// ---------------------------------------------------
// Generators ----------------------------------------

func newGeneralOk(message string) generalOk {
    return generalOk {
        Status: 200,
        Message: message
    }
}

func newDeleteMessage(succes bool, message string) {
    return deleteMessage {
        Status: 200,
        Message: message,
        Success: succes
    }
}

func newEntityError(code int, where string, message string) wrongEntityError {
	return wrongEntityError{
		Status:  code,
		Where:   where,
		Message: message,
	}
}

func newDatabaseError(code int, where string, message string) databaseError {
	return databaseError{
		Status:  code,
		Where:   where,
		Message: message,
	}
}

// ---------------------------------------------------