package formaterror

import (
	//"errors"
	"strings"
)

var errorMessages = make(map[string]string)
var err error

// FormatError : handle errors when they occur
func FormatError(errString string) map[string]string {
	if strings.Contains(errString, "email") {
		errorMessages["Taken_email"] = "Un compte existe déjà avec cet email"
	}

	if strings.Contains(errString, "title") {
		errorMessages["Taken_title"] = "Title Already Taken"
	}

	if strings.Contains(errString, "hashedPassword") {
		errorMessages["Incorrect_password"] = "Mot de passe incorrect"
	}

	if strings.Contains(errString, "record not found") {
		errorMessages["No_record"] = "No Record Found"
	}

	if strings.Contains(errString, "You already liked this trip") {
		errorMessages["Double_like"] = "Vous aimez déjà ce voyage"
	}

	if strings.Contains(errString, "You already participated to this trip") {
		errorMessages["Double_participation"] = "Vous participez déjà au voyage"
	}

	if len(errorMessages) > 0 {
		return errorMessages
	}

	if len(errorMessages) == 0 {
		errorMessages["Incorrect_details"] = "Incorrect Details"
		return errorMessages
	}

	return nil
}
