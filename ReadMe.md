# Small API Rest with GO

This project has a 4 endpoints, two of them are connected with a Database on Supabase.

---


| Method | Endpoint | Description |
|------|---------|------------|
| GET | /albums | Listar álbumes |
| POST | /albums | Crear álbum |
| DELETE | /albums/{id} | Eliminar álbum |
| GET | /schools | Listar escuelas |
| POST | /schools | Crear escuelas (permite un `Array`) |


---

## 📌 GET /schools

Make a `GET` to `/schools`.

**Response**

```json
{
  "data": [
    {
      "uuid": "hh2b0c47-56c5-4c1b-ca39-e8e31be1d71a",
      "id": 15,
      "name": "School 1",
      "address": "Address school 1"
    },
    {
      "uuid": "aah2b0c47-56c5-4dsb-ca39-e8e31be1d81x",
      "id": 16,
      "name": "School 2",
      "address": "Address school 2"
    }
  ],

  "status": 200
}
```

---

## 📌 POST /schools

**Body**

```json
{
  "name": "Name school",
  "address": "Address school"
}
```

**Response**

The server response when status code is `201`

```json
{
  "data": [
    {
      "uuid": "hh2b0c47-56c5-4c1b-ca39-e8e31be1d71a",
      "id": 15,
      "name": "school",
      "address": "Address school"
    },
  ],

  "status": 201
}
```

## 📌 DELETE /school/{uuid}

---

### Path Parameters
| Param | Type | Required | Description |
|------|------|----------|-------------|
| uuid | string | ✅ Sí | UUID de la escuela a eliminar |


---

### Request Example
```http
DELETE /schools/9a3f2c1e-7b21-4c8a-9e4a-123456789abc
```

### Response Example when all was success
```json
{
  "message": "School deleted successfully.",
  "status": 204
}
```

