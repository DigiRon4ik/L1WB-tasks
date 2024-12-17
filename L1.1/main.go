package main

import "fmt"

// Human parent structure with multiple fields and methods.
type Human struct {
	Name string
	Age  int
}

// A method for Human that gives you something to look at.
func (h *Human) lookAt(entity string) {
	fmt.Printf("%s is looking at %s.\n", h.Name, entity)
}

// A method for Human that gives you the opportunity to listen to something.
func (h *Human) listensTo(entity string) {
	fmt.Printf("%s is listens to %s.\n", h.Name, entity)
}

// A child structure of Action, which contains an anonymous Human field,
// which allows you to work with Action as if it had the same methods and fields as Human.
type Action struct {
	Human
	Activity string
}

// The doActivity method to be able to do something.
func (a *Action) doActivity() {
	fmt.Printf("%s is now doing: %s.\n", a.Name, a.Activity)
}

// Overriding the lookAt method, now if you don't explicitly call the method from Human
// when working with an Action, the lookAt method from the Action will be used by default.
func (a *Action) lookAt() {
	fmt.Printf("%s stares into the monitor.\n", a.Name)
}

func main() {
	action := Action{
		Human: Human{
			Name: "John Doe",
			Age:  23,
		},
		Activity: "walking",
	}

	// Explicit call of lookAt method of Human.
	action.Human.lookAt("tree")
	// listensTo method from Human. If the method is not overridden, it will be called on the parent by default.
	action.listensTo("music")
	// lookAt method from Action. Overridden method.
	action.lookAt()
	// doActivity method from Action. This method does not have Human.
	action.doActivity()
}

/*
 - Output: -
John Doe is looking at tree.
John Doe is listens to music.
John Doe stares into the monitor.
John Doe is now doing: walking.
*/
