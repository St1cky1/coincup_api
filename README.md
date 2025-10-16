```markdown
# 🚀 CoinCap API Client

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)
[![API Status](https://img.shields.io/badge/API-Live-brightgreen?style=for-the-badge)](https://docs.coincap.io/)

**Простой и мощный Go-клиент для работы с данными о криптовалютах через CoinCap API**

</div>

## 📖 О проекте

Этот клиент предоставляет удобный интерфейс для получения актуальной информации о криптовалютах из официального API CoinCap. Идеально подходит для создания крипто-трекеров, аналитических инструментов и торговых ботов.

### ✨ Возможности

- 📊 **Получение данных** о всех криптовалютах
- 🔍 **Поиск по конкретной валюте** (Bitcoin, Ethereum, etc.)
- ⚡ **Высокая производительность** благодаря Go
- 📝 **Автоматическое логирование** запросов
- 🛡 **Надёжная обработка ошибок**
- ⏱ **Настраиваемые таймауты** запросов

## 🚀 Быстрый старт

### Предварительные требования

- [Go](https://golang.org/dl/) версии 1.19 или выше
- [API ключ](https://coincap.io/) от CoinCap (бесплатный)

### Установка

```bash
# Клонируйте репозиторий
git clone https://github.com/Sticky1/coincup.api.git
cd coincup.api

# Установите зависимости
go mod download
```

### Настройка API ключа

1. Получите бесплатный API ключ на [coincap.io](https://coincap.io/)
2. Создайте файл `.env` в корне проекта:
```env
API_KEY=ваш_супер_секретный_ключ_здесь
```

## 💻 Использование

### Базовый пример

```go
package main

import (
    coincap "coincap_api_client/coincap_client"
    "fmt"
    "log"
    "time"
)

func main() {
    // Создаём клиент с 10-секундным таймаутом
    client, err := coincap.NewClient(10 * time.Second)
    if err != nil {
        log.Fatal("Ошибка создания клиента:", err)
    }

    // Получаем список всех криптовалют
    assets, err := client.GetAssets()
    if err != nil {
        log.Fatal("Ошибка получения данных:", err)
    }

    // Выводим топ-5 криптовалют
    fmt.Println("🏆 Топ-5 криптовалют по рыночной капитализации:")
    for i := 0; i < 5 && i < len(assets); i++ {
        asset := assets[i]
        fmt.Printf("%d. %s (%s) - $%s\n", 
            i+1, asset.Name, asset.Symbol, asset.PriceUsd)
    }
}
```

### Получение конкретной криптовалюты

```go
// Получаем данные по Bitcoin
bitcoin, err := client.GetAsset("bitcoin")
if err != nil {
    log.Fatal("Ошибка получения Bitcoin:", err)
}

fmt.Printf("💰 Bitcoin (BTC)\n")
fmt.Printf("   Цена: $%s\n", bitcoin.PriceUsd)
fmt.Printf("   Капитализация: $%s\n", bitcoin.MarketCapUsd)
fmt.Printf("   Объем (24ч): $%s\n", bitcoin.VolumeUsd24Hr)
```

### Расширенное использование

```go
// Мониторинг нескольких криптовалют
func monitorCrypto(client *coincap.Client) {
    currencies := []string{"bitcoin", "ethereum", "cardano", "solana"}
    
    for _, currency := range currencies {
        asset, err := client.GetAsset(currency)
        if err != nil {
            fmt.Printf("Ошибка получения %s: %v\n", currency, err)
            continue
        }
        
        fmt.Printf("📈 %s: $%s (изменение за 24ч: %s)\n",
            asset.Name, asset.PriceUsd, asset.ChangePercent24Hr)
    }
}
```

## 📊 Структура данных

### Модель Asset

```go
type Asset struct {
    ID            string `json:"id"`           // Уникальный идентификатор
    Rank          string `json:"rank"`         // Ранг по капитализации
    Symbol        string `json:"symbol"`       // Символ (BTC, ETH)
    Name          string `json:"name"`         // Название
    Supply        string `json:"supply"`       // Текущее предложение
    MaxSupply     string `json:"maxSupply"`    // Максимальное предложение
    MarketCapUsd  string `json:"marketCapUsd"` // Рыночная капитализация
    VolumeUsd24Hr string `json:"volumeUsd24Hr"`// Объем торгов (24ч)
    PriceUsd      string `json:"priceUsd"`     // Цена в USD
}
```

## 🛠 API методы

### `GetAssets() ([]Asset, error)`
Возвращает массив всех доступных криптовалют, отсортированных по рыночной капитализации.

**Пример вывода:**
```
1. Bitcoin (BTC) - $57,747.47
2. Ethereum (ETH) - $3,501.29
3. Binance Coin (BNB) - $623.84
...
```

### `GetAsset(name string) (Asset, error)`
Возвращает детальную информацию о конкретной криптовалюте по её ID.

**Поддерживаемые ID:** `bitcoin`, `ethereum`, `cardano`, `solana`, и [другие](https://docs.coincap.io/#89deffa0-ab03-4e0a-8d92-637a857d4c91)

## ⚙️ Конфигурация

### Настройка таймаутов

```go
// Для быстрых запросов
fastClient, _ := coincap.NewClient(5 * time.Second)

// Для нестабильного соединения
stableClient, _ := coincap.NewClient(30 * time.Second)
```

### Логирование

Клиент автоматически логирует все запросы:
```
[Thu Oct 16 11:03:04 2025] GET https://rest.coincap.io/v3/assets/
[Thu Oct 16 11:03:05 2025] GET https://rest.coincap.io/v3/assets/bitcoin
```

## 🐛 Решение проблем

### Частые ошибки и их решение

| Ошибка | Причина | Решение |
|--------|---------|---------|
| `API key not found` | Отсутствует .env файл | Создайте .env с API_KEY |
| `timeout can't be zero` | Нулевой таймаут | Укажите время > 0 |
| `Auth service unavailable` | Проблемы с сетью | Проверьте подключение |
| `Invalid JSON` | Изменения в API | Обновите клиент |

### Тестирование подключения

```bash
# Проверка API ключа
curl -H "Authorization: Bearer YOUR_API_KEY" \
  https://rest.coincap.io/v3/assets/bitcoin
```

## 🤝 Участие в разработке

Мы приветствуем вклад в проект! Вот как вы можете помочь:

1. **Сообщите о баге** через [Issues](https://github.com/Sticky1/coincup.api/issues)
2. **Предложите новую функцию** 
3. **Сделайте pull request**

### Процесс внесения изменений

```bash
# 1. Сделайте форк репозитория
# 2. Создайте ветку для функции
git checkout -b feature/amazing-feature

# 3. Закоммитьте изменения
git commit -m 'Add amazing feature'

# 4. Запушите ветку
git push origin feature/amazing-feature

# 5. Откройте Pull Request
```

## 📈 Примеры реального использования

### Крипто-портфель

```go
func trackPortfolio(client *coincap.Client, holdings map[string]float64) {
    total := 0.0
    
    for asset, amount := range holdings {
        data, err := client.GetAsset(asset)
        if err != nil {
            continue
        }
        
        price, _ := strconv.ParseFloat(data.PriceUsd, 64)
        value := price * amount
        total += value
        
        fmt.Printf("%s: %.2f × $%.2f = $%.2f\n", 
            data.Name, amount, price, value)
    }
    
    fmt.Printf("\n💰 Общая стоимость портфеля: $%.2f\n", total)
}
```

## 🛡 Безопасность

- 🔐 **API ключи** хранятся в .env файле
- 📁 **.env добавлен в .gitignore**
- 🌐 **HTTPS** для всех запросов
- ⏱ **Защита от долгих запросов** через таймауты

## 🔗 Полезные ссылки

- [📚 Документация CoinCap API](https://docs.coincap.io/)
