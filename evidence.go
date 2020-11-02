package swid

import "time"

// Evidence models a evidence-entry
type Evidence struct {
	EvidenceExtension

	GlobalAttributes
	ResourceCollection

	// The date and time the information was collected pertaining to the
	// evidence item.
	Date time.Time `cbor:"35,keyasint,omitempty" json:"date,omitempty" xml:"date,attr,omitempty"`

	// The endpoint's string identifier from which the evidence was
	// collected.
	DeviceID string `cbor:"36,keyasint,omitempty" json:"device-id,omitempty" xml:"deviceId,attr,omitempty"`
}

// NewEvidence instantiates a new Evidence object initialised with the given
// deviceID
func NewEvidence(deviceID string) *Evidence {
	return &Evidence{
		DeviceID: deviceID,
		Date:     time.Now(),
	}
}

// AddFile adds the supplied File to the embedded ResourceCollection of the
// Evidence receiver
func (e *Evidence) AddFile(f File) error {
	if e.Files == nil {
		e.Files = new(Files)
	}

	*e.Files = append(*e.Files, f)

	return nil
}

// AddProcess adds the supplied Process to the embedded ResourceCollection of the
// Evidence receiver
func (e *Evidence) AddProcess(p Process) error {
	if e.Processes == nil {
		e.Processes = new(Processes)
	}

	*e.Processes = append(*e.Processes, p)

	return nil
}
