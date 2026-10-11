package packet

import (
	"bytes"
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
)

func baseMap() map[string]json.RawMessage {
	return map[string]json.RawMessage{
		"Section":  json.RawMessage(`"packet"`),
		"Function": json.RawMessage(`"Check the seven fields and refuse an eighth."`),
		"Read":     json.RawMessage(`["plan.md"]`),
		"Do":       json.RawMessage(`["Parse the body","Validate the shape"]`),
		"Write":    json.RawMessage(`"packet/lead"`),
		"Next":     json.RawMessage(`"store"`),
		"Rail":     json.RawMessage(`""`),
	}
}

func mustObject(t *testing.T, raw map[string]json.RawMessage) string {
	t.Helper()
	b, err := json.Marshal(raw)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func mustParse(t *testing.T, body string) SectionPacket {
	t.Helper()
	pkt, ref := Parse(body)
	if ref != nil {
		t.Fatalf("Parse refusal %s field %q", ref.Code, ref.Field)
	}
	return pkt
}

func wantRefusal(t *testing.T, ref *Refusal, code, field string) {
	t.Helper()
	if ref == nil {
		t.Fatalf("expected %s field %q", code, field)
	}
	if ref.Code != code || ref.Field != field {
		t.Fatalf("got %s field %q, want %s field %q", ref.Code, ref.Field, code, field)
	}
}

func TestParseSevenFields(t *testing.T) {
	body := mustObject(t, baseMap())
	pkt := mustParse(t, body)
	if pkt.Section != "packet" {
		t.Fatalf("Section %q", pkt.Section)
	}
	if pkt.Function != "Check the seven fields and refuse an eighth." {
		t.Fatalf("Function %q", pkt.Function)
	}
	if !reflect.DeepEqual(pkt.Read, []string{"plan.md"}) {
		t.Fatalf("Read %#v", pkt.Read)
	}
	if !reflect.DeepEqual(pkt.Do, []string{"Parse the body", "Validate the shape"}) {
		t.Fatalf("Do %#v", pkt.Do)
	}
	if pkt.Write != "packet/lead" || pkt.Next != "store" || pkt.Rail != "" {
		t.Fatalf("Write %q Next %q Rail %q", pkt.Write, pkt.Next, pkt.Rail)
	}
}

func TestParseRoundTrip(t *testing.T) {
	pkt := mustParse(t, mustObject(t, baseMap()))
	b, err := json.Marshal(pkt)
	if err != nil {
		t.Fatal(err)
	}
	var keys map[string]json.RawMessage
	if err := json.Unmarshal(b, &keys); err != nil {
		t.Fatal(err)
	}
	if len(keys) != 7 {
		t.Fatalf("keys %d", len(keys))
	}
	for _, field := range fieldOrder {
		if _, ok := keys[field]; !ok {
			t.Fatalf("missing %s in %s", field, b)
		}
	}
	again := mustParse(t, string(b))
	if !reflect.DeepEqual(pkt, again) {
		t.Fatalf("round trip %#v vs %#v", pkt, again)
	}
}

func TestParseRejectsExtraBeforeMissing(t *testing.T) {
	raw := baseMap()
	delete(raw, "Rail")
	raw["thread_id"] = json.RawMessage(`"thr"`)
	raw["checkpoint"] = json.RawMessage(`"cp"`)
	_, ref := Parse(mustObject(t, raw))
	wantRefusal(t, ref, CodeExtraField, "checkpoint")
}

func TestParseRejectsMissingInOrder(t *testing.T) {
	for i, field := range fieldOrder {
		raw := baseMap()
		delete(raw, field)
		_, ref := Parse(mustObject(t, raw))
		wantRefusal(t, ref, CodeMissingField, field)
		if i == 0 && ref.Field != "Section" {
			t.Fatal(ref)
		}
	}
}

func TestParseRejectsEighthKeys(t *testing.T) {
	keys := []string{
		"Skill", "Catalog", "Model", "Steer", "Ledger", "Approval",
		"checkpoint", "checkpoint_id", "parent_checkpoint_id", "tasks",
		"metadata", "values", "input", "interrupt_before", "interrupt_after",
		"conversation_references", "thread_id", "objective", "DependsOn",
		"TaskState", "State", "Tokens", "room.say", "ID", "Evidence",
	}
	for _, key := range keys {
		raw := baseMap()
		raw[key] = json.RawMessage(`"no"`)
		_, ref := Parse(mustObject(t, raw))
		wantRefusal(t, ref, CodeExtraField, key)
	}
}

func TestParseRejectsShapes(t *testing.T) {
	cases := []struct {
		field string
		raw   string
	}{
		{"Section", `""`},
		{"Section", `"two words"`},
		{"Section", `1`},
		{"Section", `null`},
		{"Function", `"one\ntwo"`},
		{"Function", `["sentence"]`},
		{"Function", `null`},
		{"Read", `null`},
		{"Read", `"plan.md"`},
		{"Read", `[1]`},
		{"Read", `[null]`},
		{"Read", `[{"path":"plan.md"}]`},
		{"Do", `null`},
		{"Do", `{}`},
		{"Write", `["packet/lead","other"]`},
		{"Write", `"packet/lead,other"`},
		{"Write", `null`},
		{"Next", `["store"]`},
		{"Next", `"two words"`},
		{"Next", `{"id":"store"}`},
		{"Next", `null`},
		{"Rail", `{"switch":true}`},
		{"Rail", `["approve"]`},
		{"Rail", `"human flip"`},
		{"Rail", `null`},
		{"Rail", `true`},
	}
	for _, tc := range cases {
		raw := baseMap()
		raw[tc.field] = json.RawMessage(tc.raw)
		_, ref := Parse(mustObject(t, raw))
		wantRefusal(t, ref, CodeShape, tc.field)
	}
}

func TestParseAllowsEmptyOptionalFields(t *testing.T) {
	raw := baseMap()
	raw["Function"] = json.RawMessage(`""`)
	raw["Read"] = json.RawMessage(`[]`)
	raw["Do"] = json.RawMessage(`[]`)
	raw["Write"] = json.RawMessage(`""`)
	raw["Next"] = json.RawMessage(`""`)
	raw["Rail"] = json.RawMessage(`""`)
	pkt := mustParse(t, mustObject(t, raw))
	if pkt.Read == nil || pkt.Do == nil {
		t.Fatal("empty arrays became nil")
	}
	if len(pkt.Read) != 0 || len(pkt.Do) != 0 {
		t.Fatalf("Read %#v Do %#v", pkt.Read, pkt.Do)
	}
	if ref := Validate(pkt); ref != nil {
		t.Fatal(ref)
	}
}

func TestParseNextStopAndSectionID(t *testing.T) {
	raw := baseMap()
	raw["Next"] = json.RawMessage(`"stop"`)
	pkt := mustParse(t, mustObject(t, raw))
	if pkt.Next != NextStop {
		t.Fatal(pkt.Next)
	}
	raw["Next"] = json.RawMessage(`"sections/packet"`)
	pkt = mustParse(t, mustObject(t, raw))
	if pkt.Next != "sections/packet" {
		t.Fatal(pkt.Next)
	}
}

func TestParseDoesNotOpenRead(t *testing.T) {
	missing := "/tmp/den-packet-read-is-not-opened"
	if _, err := os.Stat(missing); !os.IsNotExist(err) {
		t.Fatal(err)
	}
	raw := baseMap()
	raw["Read"] = json.RawMessage(`["/tmp/den-packet-read-is-not-opened"]`)
	pkt := mustParse(t, mustObject(t, raw))
	if pkt.Read[0] != missing {
		t.Fatal(pkt.Read)
	}
	if _, err := os.Stat(missing); !os.IsNotExist(err) {
		t.Fatal("Parse touched Read")
	}
}

func TestParseRejectsNonObject(t *testing.T) {
	for _, body := range []string{``, `[]`, `"packet"`, `null`, `1`, `{"Section":"packet"}{"Next":"stop"}`} {
		_, ref := Parse(body)
		wantRefusal(t, ref, CodeShape, "")
	}
}

func TestParseRejectsDuplicateKey(t *testing.T) {
	body := `{"Section":"packet","Section":"other","Function":"","Read":[],"Do":[],"Write":"","Next":"stop","Rail":""}`
	_, ref := Parse(body)
	wantRefusal(t, ref, CodeShape, "Section")
}

func TestValidateConstructedValues(t *testing.T) {
	ok := SectionPacket{
		Section:  "packet",
		Function: "",
		Read:     []string{},
		Do:       []string{},
		Write:    "",
		Next:     "",
		Rail:     "",
	}
	if ref := Validate(ok); ref != nil {
		t.Fatal(ref)
	}
	badRead := ok
	badRead.Read = nil
	wantRefusal(t, Validate(badRead), CodeShape, "Read")
	badWrite := ok
	badWrite.Write = "a,b"
	wantRefusal(t, Validate(badWrite), CodeShape, "Write")
	badFn := ok
	badFn.Function = "one\ntwo"
	wantRefusal(t, Validate(badFn), CodeShape, "Function")
	if _, err := json.Marshal(badWrite); err == nil {
		t.Fatal("marshal stored a comma-joined Write")
	}
}

func TestValidateSeat(t *testing.T) {
	handed := SectionPacket{
		Section: "packet",
		Read:    []string{},
		Do:      []string{},
		Rail:    "approve",
		Next:    "store",
		Write:   "packet/lead",
	}
	same := handed
	same.Function = "Keep the handed rail."
	same.Next = NextStop
	same.Write = ""
	if ref := ValidateSeat(ActorSeat, handed, same); ref != nil {
		t.Fatal(ref)
	}
	cleared := handed
	cleared.Rail = ""
	wantRefusal(t, ValidateSeat(ActorSeat, handed, cleared), CodeSeatRail, "Rail")
	flipped := handed
	flipped.Rail = "hold"
	wantRefusal(t, ValidateSeat(ActorSeat, handed, flipped), CodeSeatRail, "Rail")
	if ref := ValidateSeat(ActorHuman, handed, cleared); ref != nil {
		t.Fatal(ref)
	}
	if ref := ValidateSeat(ActorHuman, handed, flipped); ref != nil {
		t.Fatal(ref)
	}
	spaced := handed
	spaced.Rail = "human flip"
	wantRefusal(t, ValidateSeat(ActorHuman, handed, spaced), CodeShape, "Rail")
	wantRefusal(t, ValidateSeat(ActorSeat, handed, spaced), CodeShape, "Rail")
	wantRefusal(t, ValidateSeat(Actor("model"), handed, handed), CodeShape, "Actor")
}

func TestAttachLeavesPacketFields(t *testing.T) {
	pkt := mustParse(t, mustObject(t, baseMap()))
	tool := &ToolAttach{Name: "kb", Result: "stop"}
	turn := Attach(pkt, tool)
	if turn.Tool != tool {
		t.Fatal("tool not attached")
	}
	if turn.Packet.Next != pkt.Next || turn.Packet.Write != pkt.Write || turn.Packet.Rail != pkt.Rail {
		t.Fatal("Attach wrote Next, Write, or Rail")
	}
	if !reflect.DeepEqual(turn.Packet, pkt) {
		t.Fatal("Attach changed the packet")
	}
	bare := Attach(pkt, nil)
	if bare.Tool != nil {
		t.Fatal("nil tool was filled in")
	}
	b, err := json.Marshal(turn)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(b, []byte(`"stop"`)) || bytes.Contains(b, []byte(`"kb"`)) {
		t.Fatalf("tool result stored in %s", b)
	}
	var keys map[string]json.RawMessage
	if err := json.Unmarshal(b, &keys); err != nil {
		t.Fatal(err)
	}
	if len(keys) != len(fieldOrder) {
		t.Fatalf("stored keys %d", len(keys))
	}
}

func TestHandOffBodyIgnoresMeshFields(t *testing.T) {
	body := mustObject(t, baseMap())
	task := Task{
		ID:        "task-9",
		Title:     "not the function",
		Body:      body,
		Dept:      "packet",
		Role:      "lead",
		State:     "completed",
		Owns:      []string{"sections/packet/*"},
		DependsOn: []string{"other-task"},
		Lease:     "lease-1",
		Evidence:  "hash",
	}
	pkt, ref := HandOffBody(task)
	if ref != nil {
		t.Fatal(ref)
	}
	direct := mustParse(t, body)
	if !reflect.DeepEqual(pkt, direct) {
		t.Fatalf("handoff %#v vs %#v", pkt, direct)
	}
	stored, err := json.Marshal(pkt)
	if err != nil {
		t.Fatal(err)
	}
	for _, leak := range []string{"task-9", "not the function", "completed", "other-task", "lease-1", "hash", "sections/packet/*"} {
		if bytes.Contains(stored, []byte(leak)) {
			t.Fatalf("mesh field %q leaked into %s", leak, stored)
		}
	}
}

func TestUnmarshalJSONFailsClosed(t *testing.T) {
	var pkt SectionPacket
	err := json.Unmarshal([]byte(`{"Section":"packet"}`), &pkt)
	if err == nil {
		t.Fatal("partial object unmarshaled")
	}
	ref, ok := err.(*Refusal)
	if !ok {
		t.Fatal(err)
	}
	wantRefusal(t, ref, CodeMissingField, "Function")
}

func TestSourceDoesNotCallMeshOrOpenFiles(t *testing.T) {
	src, err := os.ReadFile("packet.go")
	if err != nil {
		t.Fatal(err)
	}
	text := string(src)
	banned := []string{
		"CutNext", "PutKb", "CreateTask", "kb.query", "task.get",
		"os.Open", "os.ReadFile", "deerflow", "deer-flow", "http.",
	}
	for _, word := range banned {
		if strings.Contains(text, word) {
			t.Fatalf("packet.go contains %q", word)
		}
	}
}
