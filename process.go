package swid

// Process models a process-entry
type Process struct {
	ProcessExtension

	GlobalAttributes

	// The software component's process name as it will appear in an endpoint's
	// process list.
	ProcessName string `cbor:"27,keyasint" json:"process-name" xml:"name,attr"`

	//  The process ID identified for a running instance of the software
	//  component in the endpoint's process list. This is used as part of the
	//  evidence item.
	Pid *int `cbor:"28,keyasint,omitempty" json:"pid,omitempty" xml:"pid,attr,omitempty"`
}
