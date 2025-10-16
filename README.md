# CoinCap API Client

Go-клиент для работы с API сервиса [CoinCap](https://coincap.io/) для получения данных о криптовалютах.

## 🚀 Возможности

- Получение списка всех криптовалют
- Получение информации о конкретной криптовалюте по названию
- Логирование HTTP-запросов
- Поддержка таймаутов запросов

## 📦 Установка

### Требования
- Go 1.19 или выше
- API ключ от CoinCap

### Клонирование репозитория
```bash
git clone https://github.com/Sticky1/coincup.api.git
cd coincup.api
```

### Настройка окружения
1. Получите API ключ на [CoinCap.io](https://coincap.io/)
2. Создайте файл `.env` в корне проекта:
```env
API_KEY=ваш_api_ключ_здесь
```

## 🔧 Использование

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
    // Создание клиента с таймаутом 10 секунд
    client, err := coincap.NewClient(10 * time.Second)
    if err != nil {
        log.Fatal(err)
    }

    // Получение списка всех криптовалют
    assets, err := client.GetAssets()
    if err != nil {
        log.Fatal(err)
    }

    // Вывод информации о криптовалютах
    for _, asset := range assets {
        fmt.Println(asset.Info())
    }

    // Получение информации о конкретной криптовалюте
    bitcoin, err := client.GetAsset("bitcoin")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Цена Bitcoin: $%s\n", bitcoin.PriceUsd)
}
```

## 📁 Структура проекта

```
coincup.api/
├── coincap_client/
│   ├── client.go          # Основной клиент и методы API
│   ├── responses.go       # Структуры ответов API
│   └── roundtrip.go       # Логирование HTTP-запросов
├── main.go                # Пример использования
├── go.mod                 # Зависимости Go
├── go.sum                 # Контрольные суммы зависимостей
├── .env                   # Файл с API ключом (не в репозитории)
└── .gitignore            # Игнорируемые файлы
```

## 🔑 Модели данных

### Структура Asset
```go
type Asset struct {
    ID            string `json:"id"`
    Rank          string `json:"rank"`
    Symbol        string `json:"symbol"`
    Name          string `json:"name"`
    Supply        string `json:"supply"`
    MaxSupply     string `json:"maxSupply"`
    MarketCapUsd  string `json:"marketCapUsd"`
    VolumeUsd24Hr string `json:"volumeUsd24Hr"`
    PriceUsd      string `json:"priceUsd"`
}
```

## 🛠 Методы API

### `GetAssets()`
Возвращает массив всех криптовалют.

**Пример ответа:**
```
Id ~ bitcoin | Rank ~ 1 | Symbol ~ BTC | Name ~ Bitcoin | Supply ~ 19624143 | MaxSupply ~ 21000000 | MaketCapUsd ~ 1133726090024 | VolumeUsd24Hr ~ 16563596754 | PriceUsd ~ 57747.47
```

### `GetAsset(name string)`
Возвращает информацию о конкретной криптовалюте по её ID.

**Пример использования:**
```go
bitcoin, err := client.GetAsset("bitcoin")
ethereum, err := client.GetAsset("ethereum")
```

## ⚙️ Настройка

### Таймауты
При создании клиента можно указать таймаут:
```go
// Клиент с таймаутом 5 секунд
client, err := coincap.NewClient(5 * time.Second)

// Клиент с таймаутом 30 секунд  
client, err := coincap.NewClient(30 * time.Second)
```

### Логирование
Клиент автоматически логирует все HTTP-запросы в формате:
```
[Thu Oct 16 11:03:04 2025] GET https://rest.coincap.io/v3/assets/
```

## 🐛 Решение проблем

### Ошибка: "API key not found"
Убедитесь, что файл `.env` существует и содержит корректный API ключ.

### Ошибка: "timeout can't be zero"
Укажите положительное значение таймаута при создании клиента.

### Ошибка: "file .env not find"
Создайте файл `.env` в корне проекта с вашим API ключом.


## 🔗 Полезные ссылки

- [CoinCap API Documentation](https://docs.coincap.io/)
