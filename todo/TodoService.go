package todo

import "errors"

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Status == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
