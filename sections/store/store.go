// Package store is the append-only section store.
// Put, Query, and AppendRefusal are the mesh path.
// ReadFast and ReadRecent read this machine.
// The OmniRoute copy is backup, not a read path.
package store

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	sectionsFile = "sections.jsonl"
	refusalsFile = "refusals.jsonl"
	progressFile = "progress.jsonl"
	currentFile  = "current"

	// BackupRoot is the OmniRoute copy. Seats do not read it to remember.
	BackupRoot = "tcp:/opt/harness-memory"
)

var (
	ErrEmptyWrite    = errors.New("store: Write is empty")
	ErrEmptyBody     = errors.New("store: section body is empty")
	ErrMissingFields = errors.New("store: Section, Function, Read, or Do is empty")
	ErrShrink        = errors.New("store: refused to shrink stored section")
	ErrSecondStore   = errors.New("store: second store of current section")
	ErrRewriteCorpus = errors.New("store: in-place rewrite of stored section")
	ErrDelete        = errors.New("store: delete is refused")
	ErrCutNotStop    = errors.New("store: lead cut is not a clean stop")
	ErrChildWrite    = errors.New("store: child cannot write parent Write, Next, or Rail")
	ErrNotFound      = errors.New("store: section not found")
	ErrBackupRead    = errors.New("store: backup is not the read path")
	ErrEmptyReason   = errors.New("store: refusal reason is empty")
)

// SectionID is one job. It is the key identity for a stored section.
type SectionID string

// RailSwitch is empty, or a named switch only the human flips.
type RailSwitch string

// TaskID is the mesh task on a progress or refusal row.
type TaskID string

// NextRef.Kind is "id", "stop", or "". Empty is unset. "stop" is a clean stop.
type NextRef struct {
	Kind string
	ID   SectionID
}

// SectionRecord is the section. Seven fields, no eighth.
type SectionRecord struct {
	Section  SectionID
	Function string
	Read     []string
	Do       []string
	Write    string
	Next     NextRef
	Rail     RailSwitch
}

// ProgressRow is one append-only ledger line.
type ProgressRow struct {
	At       time.Time
	Task     TaskID
	State    string // stored | cut | commit-refuse | failed | stop
	Artifact string
	Next     SectionID
}

// RefusalRow is a lead cut. Not a new section. Not NEXT stop.
type RefusalRow struct {
	At      time.Time
	Task    TaskID
	State   string // cut | commit-refuse
	Refused NextRef
	Reason  string
}

// MemoryLayer is one local fast or recent note.
type MemoryLayer struct {
	Tier      string // fast | recent
	Scope     string
	Key       string
	Value     string
	Source    string
	Timestamp time.Time
}

// Store is one writer over append-only JSONL under root.
type Store struct {
	root string
	home string
	mu   sync.Mutex
}

type sectionWire struct {
	Section  string   `json:"Section"`
	Function string   `json:"Function"`
	Read     []string `json:"Read"`
	Do       []string `json:"Do"`
	Write    string   `json:"Write"`
	Next     string   `json:"Next"`
	Rail     string   `json:"Rail"`
}

type refusalWire struct {
	At      time.Time `json:"at"`
	Task    string    `json:"task"`
	State   string    `json:"state"`
	Refused string    `json:"refused"`
	Reason  string    `json:"reason"`
}

type progressWire struct {
	At       time.Time `json:"at"`
	Task     string    `json:"task"`
	State    string    `json:"state"`
	Artifact string    `json:"artifact"`
	Next     string    `json:"next"`
}

// Open creates root if needed. A backup path is refused.
func Open(root string) (*Store, error) {
	return open(root, "")
}

func open(root, home string) (*Store, error) {
	if isBackupPath(root) {
		return nil, ErrBackupRead
	}
	if err := os.MkdirAll(root, 0o755); err != nil {
		return nil, err
	}
	if home == "" {
		h, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		home = h
	}
	if isBackupPath(home) {
		return nil, ErrBackupRead
	}
	return &Store{root: root, home: home}, nil
}

// Put appends one section. Same Section id a second time is refused.
// A later Section id supersedes current. Previous rows stay.
func (s *Store) Put(rec SectionRecord) error {
	if rec.Write == "" {
		return ErrEmptyWrite
	}
	if rec.Section == "" || rec.Function == "" || len(rec.Read) == 0 || len(rec.Do) == 0 {
		return ErrMissingFields
	}
	body, err := EncodeBody(rec)
	if err != nil {
		return err
	}
	if len(bytes.TrimSpace(body)) == 0 {
		return ErrEmptyBody
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if isBackupPath(s.root) {
		return ErrBackupRead
	}

	rows, err := s.readSections()
	if err != nil {
		return err
	}
	for _, prev := range rows {
		if prev.Section == rec.Section {
			return ErrSecondStore
		}
		if prev.Write == rec.Write {
			old, err := EncodeBody(prev)
			if err != nil {
				return err
			}
			if len(body) < len(old) {
				return ErrShrink
			}
			return ErrRewriteCorpus
		}
	}
	if err := appendJSONL(filepath.Join(s.root, sectionsFile), body); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(s.root, currentFile), []byte(rec.Section+"\n"), 0o644)
}

// Query returns the stored record for id. Empty id is current.
// It does not read BackupRoot. It does not return a steer or ledger body.
func (s *Store) Query(id SectionID) (SectionRecord, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if isBackupPath(s.root) {
		return SectionRecord{}, ErrBackupRead
	}
	if id == "" {
		cur, err := s.readCurrent()
		if err != nil {
			return SectionRecord{}, err
		}
		id = cur
	}
	rows, err := s.readSections()
	if err != nil {
		return SectionRecord{}, err
	}
	var found *SectionRecord
	for i := range rows {
		if rows[i].Section == id {
			found = &rows[i]
		}
	}
	if found == nil {
		return SectionRecord{}, ErrNotFound
	}
	return *found, nil
}

// AppendRefusal stores a lead cut. Current section stays.
// State "stop" is ErrCutNotStop. This is not Put of a successor.
func (s *Store) AppendRefusal(row RefusalRow) error {
	if row.State == "stop" {
		return ErrCutNotStop
	}
	if row.State != "cut" && row.State != "commit-refuse" {
		return ErrCutNotStop
	}
	if strings.TrimSpace(row.Reason) == "" {
		return ErrEmptyReason
	}
	if row.At.IsZero() {
		row.At = time.Now().UTC()
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if isBackupPath(s.root) {
		return ErrBackupRead
	}
	wire := refusalWire{
		At:      row.At.UTC(),
		Task:    string(row.Task),
		State:   row.State,
		Refused: formatNext(row.Refused),
		Reason:  row.Reason,
	}
	body, err := json.Marshal(wire)
	if err != nil {
		return err
	}
	return appendJSONL(filepath.Join(s.root, refusalsFile), body)
}

// AppendProgress appends timestamp, task id, state, artifact, next id.
func (s *Store) AppendProgress(row ProgressRow) error {
	if row.State == "" {
		return ErrMissingFields
	}
	if row.At.IsZero() {
		row.At = time.Now().UTC()
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if isBackupPath(s.root) {
		return ErrBackupRead
	}
	wire := progressWire{
		At:       row.At.UTC(),
		Task:     string(row.Task),
		State:    row.State,
		Artifact: row.Artifact,
		Next:     string(row.Next),
	}
	body, err := json.Marshal(wire)
	if err != nil {
		return err
	}
	return appendJSONL(filepath.Join(s.root, progressFile), body)
}

// ReadFast reads ~/.local/share/harness/memory/fast on this machine.
func (s *Store) ReadFast() ([]MemoryLayer, error) {
	return s.readMemory("fast", s.fastDir())
}

// ReadRecent reads ~/.local/share/harness/memory/recent on this machine.
func (s *Store) ReadRecent() ([]MemoryLayer, error) {
	return s.readMemory("recent", s.recentDir())
}

// ReadBackup is not a read path.
func ReadBackup() ([]MemoryLayer, error) {
	return nil, ErrBackupRead
}

// ChildWrite refuses a child changing the parent's Write, Next, or Rail.
func ChildWrite(parent, proposed SectionRecord) error {
	if proposed.Write != parent.Write || formatNext(proposed.Next) != formatNext(parent.Next) || proposed.Rail != parent.Rail {
		return ErrChildWrite
	}
	return nil
}

// EncodeBody encodes exactly the seven fields.
func EncodeBody(rec SectionRecord) ([]byte, error) {
	read := rec.Read
	if read == nil {
		read = []string{}
	}
	do := rec.Do
	if do == nil {
		do = []string{}
	}
	return json.Marshal(sectionWire{
		Section:  string(rec.Section),
		Function: rec.Function,
		Read:     read,
		Do:       do,
		Write:    rec.Write,
		Next:     formatNext(rec.Next),
		Rail:     string(rec.Rail),
	})
}

// DecodeBody decodes seven fields and refuses an eighth.
func DecodeBody(body []byte) (SectionRecord, error) {
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()
	var wire sectionWire
	if err := dec.Decode(&wire); err != nil {
		return SectionRecord{}, err
	}
	if wire.Write == "" && wire.Section == "" && wire.Function == "" {
		return SectionRecord{}, ErrEmptyBody
	}
	return SectionRecord{
		Section:  SectionID(wire.Section),
		Function: wire.Function,
		Read:     wire.Read,
		Do:       wire.Do,
		Write:    wire.Write,
		Next:     parseNext(wire.Next),
		Rail:     RailSwitch(wire.Rail),
	}, nil
}

func (s *Store) fastDir() string {
	return filepath.Join(s.home, ".local", "share", "harness", "memory", "fast")
}

func (s *Store) recentDir() string {
	return filepath.Join(s.home, ".local", "share", "harness", "memory", "recent")
}

func (s *Store) readMemory(tier, dir string) ([]MemoryLayer, error) {
	if isBackupPath(dir) || isBackupPath(s.home) {
		return nil, ErrBackupRead
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	out := make([]MemoryLayer, 0, len(entries))
	for _, ent := range entries {
		if ent.IsDir() || strings.HasPrefix(ent.Name(), ".") {
			continue
		}
		path := filepath.Join(dir, ent.Name())
		if isBackupPath(path) {
			return nil, ErrBackupRead
		}
		body, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		info, err := ent.Info()
		if err != nil {
			return nil, err
		}
		out = append(out, MemoryLayer{
			Tier:      tier,
			Scope:     "local",
			Key:       ent.Name(),
			Value:     string(body),
			Source:    path,
			Timestamp: info.ModTime().UTC(),
		})
	}
	return out, nil
}

func (s *Store) readSections() ([]SectionRecord, error) {
	data, err := os.ReadFile(filepath.Join(s.root, sectionsFile))
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	var rows []SectionRecord
	for _, line := range bytes.Split(data, []byte("\n")) {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		rec, err := DecodeBody(line)
		if err != nil {
			return nil, err
		}
		rows = append(rows, rec)
	}
	return rows, nil
}

func (s *Store) readCurrent() (SectionID, error) {
	b, err := os.ReadFile(filepath.Join(s.root, currentFile))
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return "", ErrNotFound
		}
		return "", err
	}
	id := SectionID(strings.TrimSpace(string(b)))
	if id == "" {
		return "", ErrNotFound
	}
	return id, nil
}

func appendJSONL(path string, body []byte) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(append(body, '\n')); err != nil {
		return err
	}
	return f.Sync()
}

func formatNext(n NextRef) string {
	switch n.Kind {
	case "stop":
		return "stop"
	case "id":
		return string(n.ID)
	default:
		return ""
	}
}

func parseNext(s string) NextRef {
	switch s {
	case "":
		return NextRef{}
	case "stop":
		return NextRef{Kind: "stop"}
	default:
		return NextRef{Kind: "id", ID: SectionID(s)}
	}
}

func isBackupPath(p string) bool {
	if p == "" {
		return false
	}
	n := filepath.ToSlash(p)
	if strings.HasPrefix(n, "tcp:") {
		return true
	}
	return strings.Contains(n, "/opt/harness-memory")
}
