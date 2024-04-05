## API-документация

Здесь вы найдете описание доступных API-эндпоинтов, их методов и параметров запросов.

## Эндпоинты для работы с данными об автомобилях

#### Эндпоинт 1: Добавление нового автомобиля

**Путь**: /car/new

**Метод**: POST

**Описание**: Этот эндпоинт предназначен для добавления нового автомобиля.

**Тело запроса:**

```json
{
    "regNums": ["X123XX150"] // массив гос. номеров
}
```

**Пример запроса:**

```text
POST /car/new
Content-Type: application/json

{
    "regNums": ["X123XX150"]
}
```

**Примеры ответов**
- Статус 201 Created
- Статус 400 BadRequest
- Статус 500 InternalServerError

#### Эндпоинт 2: Получение списка автомобилей

**Путь**: /car/all

**Метод**: GET

**Описание**: Этот эндпоинт предназначен для получения списка автомобилей. Есть пагинация, реализованная через limit и offset. 

**Параметры запроса:**

- limit - максимальное количество автомобилей в списке
- offset - смещение от начала списка для получения нужных данных

**Пример запроса:**

```text
GET /car/all?limit=10&offset=0
```

**Примеры ответов:**
- Статус 200 OK
  ```json
  [
    {
      "id": "5d6406f9-4f0e-4a9a-86f6-cf283c0d670d",
      "regNum": "X125XX150",
      "mark": "Nissan",
      "model": "Skyline",
      "year": 2022,
      "owner": {
        "name": "Ivan",
        "surname": "Ivanov",
        "patronymic": "Ivanovich"
      }
    },
  ]
  ```
- Статус 400 BadRequest
- Статус 500 InternalServerError

#### Эндпоинт 3: Обновление информации об автомобиле

**Путь**: /car/update/{id}

**Метод**: PUT

**Описание**: Этот эндпоинт предназначен для обновления информации об автомобиле.

**Параметры запроса:**

- id - идентификатор автомобиля в формате UUID4

**Тело запроса:**

```json
{
  "regNum": "X125XX150", // новый гос. номер
  "mark": "Nissan", // новая марка
  "model": "Skyline", // новая модель
  "year": 2022, // новый год выпуска
  "owner": {
    "name": "Ivan", // новое имя владельца
    "surname": "Ivanov", // новая фамилия владельца
    "patronymic": "Ivanovich" // новое отчество владельца
  }
}
```

Тело запроса также может содержать только отдельные поля для обновления.

```json
{
  "regNum": "Y125YY150"
}
```

**Пример запроса:**

```text
PUT /car/update/5d6406f9-4f0e-4a9a-86f6-cf283c0d670d
Content-Type: application/json

{
  "regNum": "Y125YY150"
}
```

**Примеры ответов:**
- Статус 200 OK
- Статус 400 BadRequest
- Статус 500 InternalServerError

#### Эндпоинт 4: Удаление информации об автомобиле

**Путь**: /car/delete/{id}

**Метод**: DELETE

**Описание**: Этот эндпоинт предназначен для удаления информации об автомобиле.

**Параметры запроса:**

- id - идентификатор автомобиля в формате UUID4

**Пример запроса:**

```text
DELETE /car/update/5d6406f9-4f0e-4a9a-86f6-cf283c0d670d
```

**Примеры ответов:**
- Статус 204 NoContent
- Статус 400 BadRequest
- Статус 500 InternalServerError