# Queue Bot for MIREA

Telegram-бот для управления очередями в чатах. Позволяет:
- Добавлять участников по команде `+` или `+<число>`
- Показывать очередь (`/get`)
- Удалять очередь (`/remove`) - только для администраторов
- Управлять очердью с помощью встроенной клавиатуры

## Зависимости
- Go 1.20+
- [go-telegram/bot](https://github.com/go-telegram/bot)
- [go-sqlite3](https://github.com/mattn/go-sqlite3)

## Установка
```bash
git clone https://github.com/emmnogodetei/queueMIREA_bot.git
cd queueMIREA_bot
go mod download

Получите токен бота у @BotFather
Запустите бота:
TELEGRAM_TOKEN="ваш_токен" go run main.go