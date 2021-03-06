//------------------------------------------------------------------------------
// From http://github.com/heitorperalles/checkCpfRestApiGoExec
//
// Distributed under The MIT License (MIT) <http://opensource.org/licenses/MIT>
//
// Copyright (c) 2020 Heitor Peralles <heitorgp@gmail.com>
//
// Permission is hereby  granted, free of charge, to any  person obtaining a copy
// of this software and associated  documentation files (the "Software"), to deal
// in the Software  without restriction, including without  limitation the rights
// to  use, copy,  modify, merge,  publish, distribute,  sublicense, and/or  sell
// copies  of  the Software,  and  to  permit persons  to  whom  the Software  is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE  IS PROVIDED "AS  IS", WITHOUT WARRANTY  OF ANY KIND,  EXPRESS OR
// IMPLIED,  INCLUDING BUT  NOT  LIMITED TO  THE  WARRANTIES OF  MERCHANTABILITY,
// FITNESS FOR  A PARTICULAR PURPOSE AND  NONINFRINGEMENT. IN NO EVENT  SHALL THE
// AUTHORS  OR COPYRIGHT  HOLDERS  BE  LIABLE FOR  ANY  CLAIM,  DAMAGES OR  OTHER
// LIABILITY, WHETHER IN AN ACTION OF  CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE  OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//------------------------------------------------------------------------------
package main

// Subject model (REQUEST specification)
//
// Represents how subject data must be submitted to the API.
//
// Request example:
// 		{
//		"name":"Heitor Peralles",
//		"cpf":"40442820135",
//		"rg":
//			{
//				"number":"209921899",
//				"issued":"2020/5/20",
//				"by":"DETRAN-RJ"
//			}
//		}
//
type Subject struct {

	// Subject's Name
	//
	// Required: False
	// Example: Heitor Peralles
	Name   string  `json:"name,omitempty"`

	// Subject's CPF (Brazilian Physical Person Register)
	//
	// Required: True
	// Example: 404.428.201-35
	// Example: 40442820135
	Cpf    string  `json:"cpf,omitempty"`

	// Subject's RG (General Registry)
	//
	// Required: False
	// Example: [RG struct]
	RG 	   *RG  	 `json:"rg,omitempty"`
}

// RG model
// Represents how General Registry data must be submitted.
type RG struct {

	// General Registry number
	//
	// Required: False
	// Example: 209921899
	Number 			 string `json:"number,omitempty"`

	// General Registry issued date
	//
	// Required: False
	// Example: 2020/05/20
	Issued   			string `json:"issued,omitempty"`

	// General Registry issuing entity
	//
	// Required: False
	// Example: DETRAN-RJ
	Entity 				string `json:"entity,omitempty"`
}

// Verdict model (RESPONSE specification)
//
// Represents how API requests will be replyed.
//
// Response example:
//		{
//		"status":"False",
//		"message":"Invalid CPF Format."
//		}
//
type Verdict struct {

		// Status (True if register OK, False otherwise)
		//
		// Required: True
		// Example: True
		// Example: False
		Status 			string `json:"status,omitempty"`

		// Error Message
		//
		// Required: False
		// Example: Invalid CPF Format
		Message 		string `json:"message,omitempty"`
}

// Messages to be attached on RESPONSE
//
// Possible messages to each expected status code
const (
		// Response message for status code 400
    INVALID_CPF_FORMAT_MESSAGE 			string = "Invalid CPF Format."

		// Response message for status code 400
    CPF_NOT_PROVIDED_MESSAGE 				string = "Required CPF field not present."

		// Response message for status code 400
    INVALID_JSON_FORMAT_MESSAGE 		string = "Invalid JSON provided."

		// Response message for status code 403
		SUBJECT_REJECTED_MESSAGE 				string = "CPF not regular or not existant."

		// Response message for status code 500
		EXTERNAL_SERVER_ERROR_MESSAGE 	string = "Problem trying to comunicate with other entities."

		// Response message for any other status code
		UNKNOWN_ERROR_MESSAGE 					string = "Unknown error."
)
