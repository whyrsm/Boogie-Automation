# 🔄 NocoDB Data Sync Script

Script ringan berbasis **Node.js** untuk menyinkronkan data antar tabel di **NocoDB**. Sudah preconfigured — tinggal sesuaikan `.env`, lalu jalankan:

```bash
docker-compose up
```

---

## ⚙️ Cara Pakai

1. Clone repo ini
2. Sesuaikan value di file `.env` di root:

```env
NOCO_URL=http://noco.example.com
NOCO_TOKEN=your-noco-token
CUSTOMER_TABLE_ID=tbl_customer
PO_CUSTOMER_TABLE_ID=tbl_po_customer
SPH_CUSTOMER_TABLE_ID=tbl_sph_customer
ARTICLE_TABLE_ID=tbl_article
```

3. Jalankan:

```bash
docker-compose up
```

---

## 🔌 Endpoint Sync

```http
POST http://{base_url}:3001/api/sync?type={sync-type}
```

Ganti `{base_url}` dengan host lokal kamu (misal: `localhost`).

---

## 🔀 Query Param `type`

| type          | Deskripsi                        |
|---------------|----------------------------------|
| customer-po   | Link Customer ke PO              |
| customer-sph  | Link Customer ke SPH             |
| article-sph   | Link Article ke SPH              |
| po-sph        | Link PO ke SPH                   |

### Contoh:

```http
POST http://localhost:3001/api/sync?type=customer-po
```

---

## 📝 Catatan

- Semua error akan tampil di log container
- Endpoint hanya aktif setelah `docker-compose up` berhasil jalan

---

## 🛠 Melihat Log

```bash
docker-compose logs -f
```
