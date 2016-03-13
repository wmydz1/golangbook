package validator
import (

    "net/url"
)
type ProfilePage struct {
    Form *url.Values

}
type Errors struct {
    Errors []error
}
//Goes through the form object and validates each elememts
//Attach