package main

// colletion of key and value pair.

// An anonymous struct is just like a normal struct, but it is defined without a name and therefore cannot be referenced elsewhere in the code.

// myCar1 := struct {
//    Make string
//    Model string
// }{
//    Make: "telsa",
//    Model: "model 3"
// }

// ---------------------

// type messageTosend struct {
// 	message   string
// 	sender    user
// 	recipient user
// }

// type user struct {
// 	name   string
// 	number int
// }

// func canSendMessage(mToSend messageTosend) bool {

// 	if mToSend.sender.name == "" {
// 		return false
// 	}
// 	if mToSend.recipient.name == "" {
// 		return false
// 	}
// 	if mToSend.sender.number == 0 {
// 		return false
// 	}
// 	if mToSend.recipient.number == 0 {
// 		return false
// 	}

// 	return true
// }

// func test(m messageTosend) {
// 	fmt.Printf("sending message : '%s' at number '%s' is '%d' \n", m.message, m.recipient.name, m.recipient.number)
// }

// func main() {
// 	sender := user{name: "jenil", number: 123}
// 	recipient := user{name: "tanish", number: 123}

// 	message := messageTosend{
// 		message:   "Thanks for siging up!",
// 		sender:    sender,
// 		recipient: recipient,
// 	}

// 	if canSendMessage(message) {
// 		test(message)
// 	} else {
// 		fmt.Println("This is Send the message")
// 	}

// }

type rect struct{
  width int
  height int
}
func (r rect) area() int {
   return r.width * r.height
}

r := rect{
   width: 5,
   height: 10,
}

fmt.Println(r.area())