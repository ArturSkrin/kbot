# kbot

# kbot

Telegram-бот на Go. Cobra + telebot.v4.

Бот: t.me/prometheus_kbot_bot

## Что умеет

- `/start` — здоровается
- `/help` — показывает список команд
- на любой текст отвечает эхом

## Как запустить

```bash
git clone https://github.com/ArturSkrin/kbot.git
cd kbot
go mod tidy
```

Создай `.env` и впиши туда токен от @BotFather:

## CI/CD Pipeline

```mermaid
flowchart LR
    A[Push to develop] --> B[GitHub Actions]
    B --> C[go build kbot]
    C --> D[Build image linux/amd64]
    D --> E[Push to ghcr.io]
    E --> F[yq bumps tag in helm/values.yaml]
    F --> G[Commit back to develop]
    G --> H[ArgoCD polls repo]
    H --> I[Sync Helm chart]
    I --> J[kbot Pod in Kubernetes]
```