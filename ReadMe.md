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

**Body**

```json
{
  "name": "Name school",
  "address": "Address school"
}
```

**Body**
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
