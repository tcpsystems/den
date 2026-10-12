// Package catalog is the runtime vocabulary.
//
// One word means one thing. Seats, skills, and tools use these names.
// The catalog indexes them. It does not run them.
package catalog

// Terms are the definitions the runtime uses. Do not overload them.
const Terms = `
identity: who a seat is. Stable id, display name, and avatar. Not a role.
role: what a seat is allowed to do in one wave. researcher, architect, builder, reviewer.
niche: the subject a seat is staffed for. One niche per seat. Not a skill.
profile: the bundle a session binds to one seat. Identity, role, niche, model, skill ids, memory scope.
skill: one markdown playbook loaded when its description matches the step. Not a tool.
tool: one argv the runtime runs inside the path jail. denread, denfetch, git.
mcp: one indexed server the runtime may start when a grant names it. Not loaded into every prompt.
plugin: one directory that adds skills, tools, or mcp index rows. The core stays small.
memory: facts the seat may read. Session notes on disk. OpenViking is a backend, not the definition.
context: the packet for one step. Identity, one skill body, the task. Not the whole catalog.
grant: a permission lease. Reason, parent, deadline. It expires. It is not a role.
model: a route id such as auto/coding. OmniRoute is a gateway. It is not the model.
`

// Kind is one catalog entry type.
type Kind string

const (
	KindIdentity Kind = "identity"
	KindRole     Kind = "role"
	KindNiche    Kind = "niche"
	KindProfile  Kind = "profile"
	KindSkill    Kind = "skill"
	KindTool     Kind = "tool"
	KindMCP      Kind = "mcp"
	KindPlugin   Kind = "plugin"
	KindMemory   Kind = "memory"
	KindModel    Kind = "model"
)
