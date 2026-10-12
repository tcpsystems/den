# Web API and UI bases

## Question

Which API and UI bases should den skills cover first: REST for denapi, and later GraphQL/Apollo, SAML, shadcn, Mantine, Svelte? denapi is a small local API over the Unix socket. Do not adopt a component kit as the harness. A skill pack can name one kit when a seat is building fox or a web app.

## Sources

- Apollo, GraphQL vs. REST (Sashko Stubailo, 2017): https://www.apollographql.com/blog/graphql-vs-rest
- Svelte docs, Overview: https://svelte.dev/docs/svelte/overview
- Local: `cmd/denctl` serves `GET` handlers on a Unix socket (`0600`) and the same mux on `127.0.0.1:8793`. Paths today are `/api/state` and `/api/events` (SSE). `plan/controller.md` and `internal/console/state.go` already mark shadcn and React as later, not the controller. Fox is SvelteKit; the controller page, when it is static Svelte, is not fox.

## Finding

Cover REST for denapi first. GraphQL, Apollo, SAML, shadcn, and Mantine stay later, and only as skills a seat loads for a product. Svelte is the one UI base worth a product skill before those kits, because fox already uses it. It is not a harness dependency.

denapi and a product API are different objects.

denapi is the daemon's own HTTP. Resources are few and known: state, then an event stream. REST fits that. Apollo's note says the core of REST is a resource at a URL, fetched with `GET`, usually JSON, and that GraphQL's gain is one query that names the fields the client wants when the server would otherwise ship a rigid, over-large document. denapi does not have that problem. One JSON state document and one SSE stream are the whole contract. A schema, a query language, and an Apollo server would sit on top of two handlers.

A product skill is what a seat loads when the packet is fox or some other web app. That pack may name one kit. It does not change denapi, the socket, or the controller binary.

Order for skills:

1. REST, for denapi only: resource URLs, JSON, loopback or the socket, no public host.
2. Svelte, as a product skill when the seat is building fox or the static controller page. Official docs: a compiler for UI components, and SvelteKit only when the app is full stack.
3. Later, one pack at a time, only if a packet names it: GraphQL or Apollo for a product API with many related reads; SAML for a product that must speak that login; shadcn or Mantine as the single named kit for that seat.

## What den should not copy

Do not adopt shadcn, Mantine, or any component kit as the harness. Do not put GraphQL or Apollo on the Unix socket. Do not treat SAML as a denapi concern. Do not load a UI kit into `denctl` because a product seat uses it. One skill pack, one kit, one product. The harness stays the Go handlers it already has.

## One plan line

Skill one: REST against denapi's socket and loopback mux; every other API or UI base is a later product skill, and Svelte is the only one fox already justifies.

## Unproven

No measurement that a GraphQL product skill would be loaded often enough to write before SAML or a kit. shadcn and Mantine were not fetched this pass. SvelteKit versus plain Svelte for the controller page is already a local plan choice, not re-tested here.
