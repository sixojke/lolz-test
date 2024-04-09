Запуск:
```
sudo docker compose up -d
```

Запросы:

Создать книгу
POST http://localhost:8009/book/create
```json
{
    "name": "book-3", 
    "author": "Лермонтов",
    "description": "тут какое-то описание",
    "genre": "триллер"
}
```

Получить информацию по книге
GET http://localhost:8009/book/e2262973-4028-4ffe-a589-6f7a67910461
```json
nil
```

Получить книги по жанру
GET http://localhost:8009/books/list?genre=триллр&limit=4
```json
nil
```

Удалить книгу
DELETE http://localhost:8009/book/delete/22ee70b9-b9fc-40f3-aa0c-96174a165c55
```json
nil
```

Найти книгу
GET http://localhost:8009/books/search?string=Лермонтов&limit=10
```json
nil
```

Получить список жанров
GET http://localhost:8009/genres
```json
nil
```
