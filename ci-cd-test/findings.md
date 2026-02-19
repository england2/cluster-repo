# Findings: CI/CD, GHCR, Helm/Kustomize, and GitOps Repo Structure

## Context

You are building microservices and want a setup that can:

- Build and push images (eventually to GHCR).
- Run basic tests in CI.
- Deploy to a Kubernetes cluster from Git.
- Potentially use Helm or Kustomize.
- Potentially use Flux/Argo-style GitOps with a separate deployment repo.

This note summarizes the information requests and conclusions from our discussion.

## Direct Answers to Your Questions

### 1) Do Kustomize or Helm build images themselves?

Finding:

- By default, no. Neither Helm nor Kustomize is an image build system.
- They are deployment configuration tools for Kubernetes manifests.

What they do:

- Set/patch `image` references (repo/tag/digest).
- Configure ports, env vars, resources, replicas, probes, etc.
- Render/apply manifests (directly or via GitOps controller).

What builds images:

- CI pipeline steps (GitHub Actions, etc.) using Docker Buildx or similar.

### 2) Is build+push in GitHub Actions after tests the right CI/CD plan?

Finding:

- Yes, that is the standard and recommended approach.

Typical sequence:

1. Run unit tests and basic checks.
2. Build image(s).
3. Push to GHCR with immutable references (commit SHA and/or digest).
4. Update deployment config to use the new image.
5. Let cluster reconciliation deploy it (Flux/Argo) or run an explicit deploy step.

Recommendation:

- Prefer immutable deployment refs (`image@sha256:...` or SHA tags) over mutable tags like `latest`.

### 3) Is keeping a Dockerfile in each service directory sane?

Finding:

- Yes, this is a common and sane monorepo pattern.

Pros:

- Service ownership is clear.
- Service-specific build context/dependencies stay local.
- Easier multi-service CI matrix builds.

Tradeoff:

- Some duplication may appear across Dockerfiles (can be handled later with shared base images or templates).

### 4) If deploying with Helm/Kustomize, should YAML here configure server port, etc.?

Finding:

- Yes. Kubernetes/deployment concerns should be expressed in deployment YAML/chart values.

Examples:

- `containerPort`, Service ports.
- replicas, resource limits/requests.
- env vars and secrets references.
- probes and rollout strategy.

Code should generally:

- Read runtime config via env vars where practical.
- Avoid hardcoding environment-specific deployment details.

### 5) Should this repo contain only microservice code, or also Kubernetes YAML?

Finding:

- Both models are valid. Choose based on team size and operational maturity.

Option A: Single repo (app + deploy together)

- App code and deployment manifests/charts live in one place.
- Faster to start, simpler workflow for small teams.

Option B: Split repos (app repo + infra/GitOps repo)

- App repo builds/pushes images.
- Separate Flux/Argo repo contains Kubernetes desired state.
- Stronger separation of concerns, better audit trail and access boundaries.

Practical recommendation from discussion:

- Start simple (single repo with `deploy/` folder) unless you already need strict separation.
- Move to split repos when governance/team boundaries demand it.

### 6) In split-repo GitOps, does app CI update image hash in Flux repo?

Finding:

- Yes, that is a common pattern.

Typical flow:

1. App CI builds/tests and pushes image to GHCR.
2. App CI creates a commit/PR in Flux repo updating image tag/digest in manifests/values.
3. Flux reconciles that repo change into the cluster.

Alternative:

- Use Flux Image Automation to watch image tags and auto-update manifests in the GitOps repo.

Recommendation from discussion:

- Start with CI creating PRs to the Flux repo (clear review/audit).
- Adopt Flux image automation later if you want less CI scripting.

## Operational Guidance (Consolidated)

- Keep service code and Dockerfiles close together (`containerX/serviceX/...`).
- Introduce a deployment layer (`deploy/helm` or `deploy/kustomize`) when ready.
- Treat image build/push and deploy config update as separate pipeline responsibilities.
- Prefer immutable image references for reliable rollouts/rollbacks.
- If using GitOps, cluster changes should come from GitOps repo reconciliation, not direct `kubectl apply` in app CI.

## Suggested Near-Term Repo Shape

If you stay single-repo for now:

- `container1/service1/` (code + Dockerfile)
- `container2/service2/` (code + Dockerfile)
- `deploy/helm/...` or `deploy/kustomize/...` (cluster config)
- `.github/workflows/` (test/build/push and deploy-config update logic)

If you split later:

- App repo: code, Dockerfiles, CI for test/build/push.
- GitOps repo: manifests/charts only, reconciled by Flux/Argo.

