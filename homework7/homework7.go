package homework7

import "fmt"

type UserService struct {
	// not need to implement
	NotEmptyStruct bool
}
type MessageService struct {
	// not need to implement
	NotEmptyStruct bool
}

type Container struct {
	deps map[string]interface{}
}

func NewContainer() *Container {

	return &Container{make(map[string]interface{})}
}

func (c *Container) RegisterType(name string, constructor interface{}) (isAdded bool, err error) {
	_, isCallable := constructor.(func() interface{})
	if isCallable {
		c.deps[name] = constructor
		return true, nil
	} else {
		return false, fmt.Errorf("constructor isn't callable")
	}
}

func (c *Container) Resolve(name string) (interface{}, error) {
	r, v := c.deps[name]
	if v {
		return r.(func() interface{})(), nil
	}
	return nil, fmt.Errorf("dependency not found: %s", name)
}
