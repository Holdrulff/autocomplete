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

## 2026-08-05 - Split local product and cloud deployment milestones

### Context

The original roadmap treated the local application, AWS deployment, production operations, and a single technical article as one delivery. The developer is learning Go, web development, React, and Docker from a beginner level. Adding AWS at the same time would mix product behavior, packaging, distributed infrastructure, security, and cost management into one learning cycle.

The project can already provide portfolio value as a complete, tested, and reproducible Docker Compose application. Cloud deployment is valuable, but it answers a different question: how an existing local architecture should evolve for a public production environment.

### Decision

The roadmap is divided into two explicit milestones:

1. **Milestone 1 - Local Product:** deliver the full autocomplete system with Go, GraphQL, React, Docker Compose, CI, documentation, release `v1.0.0`, and an English article focused on the product and its local architecture.
2. **Milestone 2 - Cloud Deployment:** adapt the stable application to AWS, add continuous deployment and operations, publish release `v2.0.0`, and write a second English article focused on the cloud evolution.

Milestone 1 is considered complete without a public URL. The GitHub repository will provide a one-command Docker Compose workflow, screenshots or a GIF, tests, and troubleshooting instructions. Milestone 2 starts only after the first release and article are complete.

The application core will still be designed for future cloud deployment. Business logic must not depend on Docker, HTTP frameworks, or AWS SDKs. Services will remain stateless, configuration will come from the environment, and transport-specific behavior will be isolated behind interfaces.

### Alternatives considered

- **Keep one milestone and one article:** rejected because the scope would delay feedback, combine too many new concepts, and make the article less focused.
- **Deploy to AWS before completing Docker:** rejected because infrastructure would make application defects harder to distinguish from deployment defects.
- **Remove cloud deployment permanently:** rejected because AWS remains valuable portfolio material once the local system is stable and understood.
- **Use a managed hosting platform for the first release:** deferred because the selected learning goal is to understand packaging with Docker first and AWS architecture later.

### Consequences

- The first release becomes achievable sooner and can be reproduced without an AWS account or ongoing cost.
- The first article can focus deeply on algorithms, APIs, frontend behavior, testing, and Docker.
- The second article gains a concrete before-and-after architecture comparison instead of presenting AWS as isolated configuration.
- A public live demo will not exist until Milestone 2; the first release will rely on repository documentation and visual evidence.
- Interfaces and configuration boundaries must be maintained from the beginning so the later AWS adapters do not require rewriting the core.

### AI usage

- **Asked:** Evaluate whether the project should first ship end-to-end with Docker and defer AWS to a second article and milestone.
- **Accepted:** The recommendation to separate local product development from cloud deployment and keep CI in the first milestone.
- **Changed:** The original plan used one release and one article; it now defines local `v1.0.0` and cloud `v2.0.0` with separate articles.
- **Rejected:** Treating AWS as a requirement for the first portfolio release.

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
