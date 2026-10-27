# Agent Guidelines

These guidelines apply to all work in this repository.

## Project

- The project name is `urchin-hat.nc`.
- Implement a TCP server in Go that streams animated urchin-hat ASCII art to terminals connected with `nc`.
- The service is planned to run on Google Cloud Compute Engine (GCE).

## Communication

- Communicate with the user in Japanese by default, including questions, progress updates, explanations, and final reports.
- Follow the user's explicit language preference when they request another language.
- Keep this `AGENTS.md` in English.

## Branches and GitOps

- Treat `main` as the integration branch. Do not commit or push directly to `main`.
- Start each task on a dedicated branch based on the latest `main`.
- Use descriptive branch names such as `feat/<description>`, `fix/<description>`, `docs/<description>`, or `chore/<description>`.
- Push changes to the task branch and integrate them into `main` through a Pull Request.
- Follow GitOps principles: manage the desired state of infrastructure and deployment configuration in Git, and make changes reviewable through Pull Requests.
- Do not leave manual changes to the live environment out of sync with the configuration in Git. Record any emergency changes in Git as well.
- Never commit credentials, private keys, or other secrets.

## Commits

- Write commit messages in Japanese and clearly describe the change.
- Prefixes such as `feat:` and `docs:` are allowed, but the description must be in Japanese.
- Keep commits small, with one purpose per commit.
- Separate unrelated features, fixes, refactoring, and formatting into different commits.
- Keep each commit coherent: include tests and documentation needed for that specific change.
- Review the staged diff before committing to avoid unrelated changes and secrets.

Examples:

```text
docs: プロジェクトの概要と接続方法を追加
feat: TCP 接続の受け付け処理を追加
fix: 切断したクライアントへの配信を停止
```

## Implementation and Verification

- Format changed Go files with `gofmt`.
- Once the Go module and tests are available, run checks appropriate to the change, such as `go test ./...`.
- For documentation-only changes, review the content and run `git diff --check`.
- Do not describe planned features or unpublished connection endpoints as available.
- Report what changed, what was checked, and any remaining work.
