package validator

import "regexp"

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)


type Validator struct {
	Errors map[string]string
}

func New() *Validator{
	return &Validator{Errors:make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(k, msg string) {
	if _, exists := v.Errors[k]; !exists{
		v.Errors[k] = msg
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func In(key string, lists ...string) bool {
	for _, v := range lists {
		if key == v {
			return true
		}
	}
	return false
}


func Matches(value string, re * regexp.Regexp) bool {
	return re.MatchString(value)
}

func Unique(values []string) bool {
	valueMap := map[string]bool{}

	for _, value := range values {
		valueMap[value] = true
	}

	return len(values) == len(valueMap)
}