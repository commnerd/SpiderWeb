// The messages package provides message constructs to communicate between nodes
package messages

// The groundwork for all node communications
interface Message{
	getPayload() interface{}
}
