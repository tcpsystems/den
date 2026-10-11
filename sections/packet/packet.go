// Package packet is the section record handed to a seat.
// The record is Section, Function, Read, Do, Write, Next, and Rail.
// Empty Function, Read, Do, Write, Next, and Rail pass these checks.
// That pass does not open a tool and does not store.
package packet

import (
	"bytes"
	"encoding/json"
	"strings"
	"unicode"
)

const (
	CodeExtraField   = "extra-field"
	CodeMissingField = "missing-field"
	CodeShape        = "shape"
	CodeSeatRail     = "seat-rail"

	ActorSeat  Actor = "seat"
	ActorHuman Actor = "human"

	// NextStop is the only Next value that is not a section id.
	NextStop = "stop"
)

// SectionID is one stable job id: a non-empty token with no whitespace.
type SectionID string

// RailSwitch is empty, or one switch token the human sets. It is not an object.
type RailSwitch string

// SectionPacket is the seven fields and no others.
type SectionPacket struct {
	Section  SectionID  `json:"Section"`
	Function string     `json:"Function"`
	Read     []string   `json:"Read"`
	Do       []string   `json:"Do"`
	Write    string     `json:"Write"`
	Next     string     `json:"Next"`
	Rail     RailSwitch `json:"Rail"`
}

// Actor is who proposed a packet. A seat cannot change Rail.
type Actor string

// ToolAttach is a tool result on a live turn. It is not a section field.
type ToolAttach struct {
	Name   string
	Result string
}

// Turn is a packet plus an optional tool result.
// Marshaling a turn writes the packet only.
type Turn struct {
	Packet SectionPacket
	Tool   *ToolAttach
}

// Refusal is one shape or authority failure. A nil pointer means the check passed.
type Refusal struct {
	Code  string
	Field string
}

// Task is the mesh task from task.claim.
// HandOffBody reads Body only. The other fields stay on the task.
type Task struct {
	ID        string
	Title     string
	Body      string
	Dept      string
	Role      string
	State     string
	Owns      []string
	DependsOn []string
	Lease     string
	Evidence  string
}

var fieldOrder = []string{"Section", "Function", "Read", "Do", "Write", "Next", "Rail"}

func (r *Refusal) Error() string {
	if r == nil {
		return ""
	}
	if r.Field == "" {
		return r.Code
	}
	return r.Code + ": " + r.Field
}

func refuse(code, field string) *Refusal {
	return &Refusal{Code: code, Field: field}
}

// Parse decodes a task body. The object must hold exactly the seven keys.
// An unknown key is extra-field, including when a required key is also absent.
// An absent key is missing-field. A present key of the wrong JSON shape is shape.
// Parse does not query the mesh and does not open Read.
func Parse(body string) (SectionPacket, *Refusal) {
	return parseBytes([]byte(body))
}

func parseBytes(body []byte) (SectionPacket, *Refusal) {
	raw, ref := decodeObject(body)
	if ref != nil {
		return SectionPacket{}, ref
	}
	if ref := exactKeys(raw); ref != nil {
		return SectionPacket{}, ref
	}
	pkt, ref := decodeFields(raw)
	if ref != nil {
		return SectionPacket{}, ref
	}
	if ref := Validate(pkt); ref != nil {
		return SectionPacket{}, ref
	}
	return pkt, nil
}

// Validate checks shape only.
// Section is one non-empty token with no whitespace.
// Function may be empty and must not contain a newline.
// Read and Do are string arrays and may be empty; a nil slice is shape.
// Write is one string or empty, never a comma-joined list.
// Next is empty, stop, or one section id.
// Rail is empty or one token with no whitespace.
func Validate(p SectionPacket) *Refusal {
	if !oneToken(string(p.Section)) {
		return refuse(CodeShape, "Section")
	}
	if strings.ContainsAny(p.Function, "\r\n") {
		return refuse(CodeShape, "Function")
	}
	if p.Read == nil {
		return refuse(CodeShape, "Read")
	}
	if p.Do == nil {
		return refuse(CodeShape, "Do")
	}
	if strings.Contains(p.Write, ",") {
		return refuse(CodeShape, "Write")
	}
	if p.Next != "" && p.Next != NextStop && !oneToken(p.Next) {
		return refuse(CodeShape, "Next")
	}
	if p.Rail != "" && !oneToken(string(p.Rail)) {
		return refuse(CodeShape, "Rail")
	}
	return nil
}

// ValidateSeat checks proposed, then Rail authority.
// A seat that sets or clears Rail is seat-rail.
// A human may change Rail only when the new packet still passes Validate.
func ValidateSeat(actor Actor, handed, proposed SectionPacket) *Refusal {
	if ref := Validate(proposed); ref != nil {
		return ref
	}
	switch actor {
	case ActorHuman:
		return nil
	case ActorSeat:
		if proposed.Rail != handed.Rail {
			return refuse(CodeSeatRail, "Rail")
		}
		return nil
	default:
		return refuse(CodeShape, "Actor")
	}
}

// Attach stores the tool result on the turn.
// It does not change Next, Write, or Rail.
func Attach(p SectionPacket, tool *ToolAttach) Turn {
	return Turn{Packet: p, Tool: tool}
}

// HandOffBody parses task.Body. Mesh identity and state are not copied onto the packet.
func HandOffBody(task Task) (SectionPacket, *Refusal) {
	return Parse(task.Body)
}

// MarshalJSON writes the seven fields. A turn's tool result is omitted.
func (t Turn) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Packet)
}

// MarshalJSON writes the seven fields after Validate.
func (p SectionPacket) MarshalJSON() ([]byte, error) {
	if ref := Validate(p); ref != nil {
		return nil, ref
	}
	type plain SectionPacket
	return json.Marshal(plain(p))
}

// UnmarshalJSON is Parse. Missing and unknown keys fail closed.
func (p *SectionPacket) UnmarshalJSON(data []byte) error {
	parsed, ref := parseBytes(data)
	if ref != nil {
		return ref
	}
	*p = parsed
	return nil
}

func oneToken(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func isField(name string) bool {
	switch name {
	case "Section", "Function", "Read", "Do", "Write", "Next", "Rail":
		return true
	default:
		return false
	}
}

func decodeObject(body []byte) (map[string]json.RawMessage, *Refusal) {
	dec := json.NewDecoder(bytes.NewReader(body))
	tok, err := dec.Token()
	if err != nil {
		return nil, refuse(CodeShape, "")
	}
	delim, ok := tok.(json.Delim)
	if !ok || delim != '{' {
		return nil, refuse(CodeShape, "")
	}
	raw := make(map[string]json.RawMessage)
	for dec.More() {
		keyTok, err := dec.Token()
		if err != nil {
			return nil, refuse(CodeShape, "")
		}
		key, ok := keyTok.(string)
		if !ok {
			return nil, refuse(CodeShape, "")
		}
		var msg json.RawMessage
		if err := dec.Decode(&msg); err != nil {
			return nil, refuse(CodeShape, "")
		}
		if _, exists := raw[key]; exists {
			return nil, refuse(CodeShape, key)
		}
		raw[key] = msg
	}
	tok, err = dec.Token()
	if err != nil {
		return nil, refuse(CodeShape, "")
	}
	end, ok := tok.(json.Delim)
	if !ok || end != '}' {
		return nil, refuse(CodeShape, "")
	}
	if dec.More() {
		return nil, refuse(CodeShape, "")
	}
	return raw, nil
}

func exactKeys(raw map[string]json.RawMessage) *Refusal {
	var extra []string
	for key := range raw {
		if !isField(key) {
			extra = append(extra, key)
		}
	}
	if len(extra) > 0 {
		sortStrings(extra)
		return refuse(CodeExtraField, extra[0])
	}
	for _, field := range fieldOrder {
		if _, ok := raw[field]; !ok {
			return refuse(CodeMissingField, field)
		}
	}
	return nil
}

func sortStrings(vals []string) {
	for i := 1; i < len(vals); i++ {
		j := i
		for j > 0 && vals[j] < vals[j-1] {
			vals[j], vals[j-1] = vals[j-1], vals[j]
			j--
		}
	}
}

func decodeFields(raw map[string]json.RawMessage) (SectionPacket, *Refusal) {
	var pkt SectionPacket
	section, ref := decodeString(raw["Section"], "Section")
	if ref != nil {
		return SectionPacket{}, ref
	}
	pkt.Section = SectionID(section)
	pkt.Function, ref = decodeString(raw["Function"], "Function")
	if ref != nil {
		return SectionPacket{}, ref
	}
	pkt.Read, ref = decodeList(raw["Read"], "Read")
	if ref != nil {
		return SectionPacket{}, ref
	}
	pkt.Do, ref = decodeList(raw["Do"], "Do")
	if ref != nil {
		return SectionPacket{}, ref
	}
	pkt.Write, ref = decodeString(raw["Write"], "Write")
	if ref != nil {
		return SectionPacket{}, ref
	}
	pkt.Next, ref = decodeString(raw["Next"], "Next")
	if ref != nil {
		return SectionPacket{}, ref
	}
	rail, ref := decodeString(raw["Rail"], "Rail")
	if ref != nil {
		return SectionPacket{}, ref
	}
	pkt.Rail = RailSwitch(rail)
	return pkt, nil
}

func decodeString(raw json.RawMessage, field string) (string, *Refusal) {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 || raw[0] != '"' {
		return "", refuse(CodeShape, field)
	}
	var s string
	if err := json.Unmarshal(raw, &s); err != nil {
		return "", refuse(CodeShape, field)
	}
	return s, nil
}

func decodeList(raw json.RawMessage, field string) ([]string, *Refusal) {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 || raw[0] != '[' {
		return nil, refuse(CodeShape, field)
	}
	var elems []json.RawMessage
	if err := json.Unmarshal(raw, &elems); err != nil {
		return nil, refuse(CodeShape, field)
	}
	out := make([]string, 0, len(elems))
	for _, el := range elems {
		el = bytes.TrimSpace(el)
		if len(el) == 0 || el[0] != '"' {
			return nil, refuse(CodeShape, field)
		}
		var s string
		if err := json.Unmarshal(el, &s); err != nil {
			return nil, refuse(CodeShape, field)
		}
		out = append(out, s)
	}
	return out, nil
}
