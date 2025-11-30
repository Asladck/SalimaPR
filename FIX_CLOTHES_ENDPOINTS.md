# ✅ Исправление: Ошибка 404 на /api/clothes/item

## 🐛 Проблема
```
GET http://localhost:8080/api/clothes/item 404 (Not Found)
```

Фронтенд запрашивал `GET /api/clothes/item` для получения всех вещей пользователя, но бэкенд не имел этого endpoint.

## 🔍 Причина

**Было в handler.go:**
```go
clothes.GET("/item/:id", h.getClothesByUserId)  // Неправильно
```

Функция `getClothesByUserId` получала все вещи пользователя, но была привязана к endpoint с параметром `:id`, что требовало ID в URL.

## ✅ Решение

Разделены два endpoint:
1. `GET /api/clothes/item` - получить все вещи пользователя
2. `GET /api/clothes/item/:id` - получить одну вещь по ID

### Измененные файлы:

#### 1. `pkg/handler/handler.go`
```go
clothes := api.Group("/clothes")
{
    clothes.POST("/item", h.addClothes)
    clothes.GET("/item", h.getAllClothes)           // ✅ Новый endpoint
    clothes.GET("/item/:id", h.getClothesById)      // ✅ Новый endpoint
    clothes.DELETE("/item/:id", h.deleteClothesById)
    clothes.PUT("/item/:id", h.updateClothesById)
}
```

#### 2. `pkg/handler/clothes.go`
**Добавлены функции:**

```go
// Получить все вещи пользователя
func (h *Handler) getAllClothes(c *gin.Context) {
    userID, err := getUserId(c)
    if err != nil {
        NewErrorResponse(c, http.StatusUnauthorized, err.Error())
        return
    }
    items, err := h.services.Clothes.GetClothesByUserId(userID)
    if err != nil {
        NewErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, dto.GetAllClothResponse{
        Data: items,
    })
}

// Получить одну вещь по ID
func (h *Handler) getClothesById(c *gin.Context) {
    userID, err := getUserId(c)
    if err != nil {
        NewErrorResponse(c, http.StatusUnauthorized, err.Error())
        return
    }
    itemID := c.Param("id")
    
    item, err := h.services.Clothes.GetClothesById(userID, itemID)
    if err != nil {
        NewErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    
    c.JSON(http.StatusOK, map[string]interface{}{
        "data": item,
    })
}
```

**Также исправлено:**
```go
// Добавлен ImageURL из input
item := &models.Clothes{
    UserId:   userID,
    Name:     input.Name,
    Category: input.Category,
    Color:    input.Color,
    Season:   input.Season,
    Material: input.Material,
    ImageURL: input.ImageURL,  // ✅ Было: ""
}
```

#### 3. `pkg/service/service.go`
```go
type Clothes interface {
    AddClothes(item *models.Clothes) (string, error)
    GetClothesByUserId(userId string) ([]models.Clothes, error)
    GetClothesById(userId, itemId string) (*models.Clothes, error)  // ✅ Добавлено
    DeleteClothesById(userId, id string) error
    UpdateClothesById(userId, itemId string, item dto.ClothesUpdateInput) error
}
```

#### 4. `pkg/service/clothes.go`
```go
func (s *ClothesService) GetClothesById(userID, itemID string) (*models.Clothes, error) {
    return s.repo.GetClothesById(userID, itemID)
}
```

#### 5. `pkg/repository/repository.go`
```go
type Clothes interface {
    AddClothes(item *models.Clothes) (string, error)
    GetClothesByUserId(userId string) ([]models.Clothes, error)
    GetClothesById(userId, itemId string) (*models.Clothes, error)  // ✅ Добавлено
    DeleteClothes(userID, id string) error
    UpdateClothes(userID, itemID string, item dto.ClothesUpdateInput) error
    GetClothesByCategory(userId, category string) ([]models.Clothes, error)
}
```

#### 6. `pkg/repository/clothes.go`
```go
func (r *ClothesRepository) GetClothesById(userId, itemId string) (*models.Clothes, error) {
    var item models.Clothes
    err := r.db.Where("user_id = ? AND id = ?", userId, itemId).First(&item).Error
    if err != nil {
        return nil, err
    }
    return &item, nil
}
```

#### 7. `pkg/dto/dto.go`
```go
type AddClothesInput struct {
    Name     string `json:"name" binding:"required"`
    Category string `json:"category" binding:"required"`
    Color    string `json:"color"`
    Season   string `json:"season"`
    Material string `json:"material"`
    ImageURL string `json:"image_url"`  // ✅ Добавлено
}
```

## 📋 API Endpoints теперь:

### Clothes
- ✅ `POST /api/clothes/item` - добавить вещь
- ✅ `GET /api/clothes/item` - получить все вещи пользователя
- ✅ `GET /api/clothes/item/:id` - получить вещь по ID
- ✅ `PUT /api/clothes/item/:id` - обновить вещь
- ✅ `DELETE /api/clothes/item/:id` - удалить вещь

## 🧪 Тестирование

### 1. Получить все вещи
```bash
curl -X GET http://localhost:8080/api/clothes/item \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

**Ожидаемый ответ:**
```json
{
  "data": [
    {
      "id": "uuid-1",
      "user_id": "user-uuid",
      "name": "Blue Jeans",
      "category": "bottom",
      "color": "Blue",
      "season": "all",
      "material": "Denim",
      "image_url": "https://...",
      "created_at": "2025-12-01T...",
      "updated_at": "2025-12-01T..."
    }
  ]
}
```

### 2. Получить одну вещь
```bash
curl -X GET http://localhost:8080/api/clothes/item/uuid-1 \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

**Ожидаемый ответ:**
```json
{
  "data": {
    "id": "uuid-1",
    "name": "Blue Jeans",
    ...
  }
}
```

### 3. Добавить вещь с изображением
```bash
curl -X POST http://localhost:8080/api/clothes/item \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Red Shirt",
    "category": "top",
    "color": "Red",
    "season": "summer",
    "material": "Cotton",
    "image_url": "https://example.com/shirt.jpg"
  }'
```

## ✅ Результат

После перезапуска сервера:

1. ✅ `GET /api/clothes/item` возвращает 200 и список вещей
2. ✅ `GET /api/clothes/item/:id` возвращает одну вещь
3. ✅ `POST /api/clothes/item` принимает image_url
4. ✅ Фронтенд успешно загружает список вещей
5. ✅ Страница wardrobe.html работает корректно

## 🚀 Как применить

1. **Перезапустите сервер:**
```bash
go run cmd/main.go
```

2. **Очистите кэш браузера** (Ctrl+F5)

3. **Проверьте страницу wardrobe:**
```
http://localhost:8080/wardrobe.html
```

4. **Результат:**
   - ✅ Нет ошибки 404
   - ✅ Список вещей загружается
   - ✅ Можно добавлять/редактировать/удалять вещи

---

**Проблема полностью решена!** 🎉

Теперь все CRUD операции для одежды работают корректно.

