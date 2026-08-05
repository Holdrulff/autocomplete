# Project History

This document records relevant engineering decisions, experiments, abandoned alternatives, and meaningful uses of artificial intelligence. It is not intended to capture trivial autocomplete interactions or every command executed during development.

## 2026-08-05 - Repository initialization

### Context

The repository was empty and had no Git history. The only local file was `enunciado.docx`, which contains the original challenge brief. The project will be developed as a public portfolio project, but the original document should not be published without explicit permission from its owner.

Pull Requests require a shared base commit on the target branch. Because `main` did not yet exist as a real Git reference, a base commit was necessary before the normal branch-and-PR workflow could begin.

### Decision

An empty commit named `chore: initialize repository` was created directly on `main` and pushed to GitHub. This is the only planned exception to the Pull Request workflow because it contains no project files or implementation.

All subsequent changes will follow this workflow:

1. Create a branch from an up-to-date `main`.
2. Make one coherent, reviewable change.
3. Run the relevant checks locally.
4. Open a Pull Request.
5. Review the change and its reasoning before merging.

The original `enunciado.docx` file will remain local and is listed in `.gitignore`. Public documentation will summarize the requirements without reproducing the original document.

### Alternatives considered

- **Commit the original brief:** rejected because the project does not have explicit permission to redistribute the document.
- **Add project files directly to `main`:** rejected because it would bypass the review workflow required by the project.
- **Create the first project commit on a feature branch:** rejected because the remote repository first needed a valid `main` reference to act as the Pull Request base.

### AI usage

- **Asked:** Read the challenge, understand the intended portfolio project, and act as a tech lead throughout its development.
- **Accepted:** The proposed architecture direction, incremental Pull Request workflow, and the use of an empty initialization commit.
- **Changed:** The original assistant behavior attempted to implement the complete project automatically. The process was changed to one microtask at a time, with the developer writing the implementation and the assistant providing guidance and review.
- **Rejected:** Automatic end-to-end implementation, because it would undermine the primary goal of learning how every part of the system works.

## Decision record template

Copy this section for future decisions that materially affect implementation or architecture.

### YYYY-MM-DD - Decision title

#### Context

What problem, constraint, or uncertainty led to this decision?

#### Decision

What was selected, and why does it fit the current project?

#### Alternatives considered

What other options were evaluated, and why were they not selected?

#### Consequences

What benefits, limitations, risks, or follow-up work result from the decision?

#### AI usage

- **Asked:** What relevant question or prompt was given to the AI?
- **Accepted:** What was used as suggested, and why?
- **Changed:** What was adapted before use, and why?
- **Rejected:** What was intentionally not used, and why?
