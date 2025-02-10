package messaging

import (
	"errors"
	"fmt"
	"net/http"
)

func ExampleLogError_Notify() {
	fmt.Printf("test: LogErrorNotifier() -> [status:%v]\n", LogErrorNotifier.Notify(aspect.StatusNotFound()))

	//Output:
	//test: LogErrorNotifier() -> [status:Not Found]

}

func ExampleOutputError_Notify() {
	status := aspect.NewStatusError(http.StatusTeapot, errors.New("kettle on the boil"))
	fmt.Printf("test: OutputErrorNotifier() -> [status:%v]\n", LogErrorNotifier.Notify(status))

	//Output:
	//test: OutputErrorNotifier() -> [status:I'm A Teapot [kettle on the boil]]

}
