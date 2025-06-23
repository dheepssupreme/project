# Blog Application Backend

Backend aplikasi blog yang dibangun dengan Go, menggunakan Supabase untuk database dan authentication.

## 🚀 Fitur

- ✅ CRUD Artikel dengan slug URL-friendly
- ✅ Sistem tagging artikel
- ✅ View tracking artikel
- ✅ Reading time calculation
- ✅ Word count calculation
- ✅ Sistem authentication dengan JWT
- ✅ Pagination untuk daftar artikel
- ✅ Filter artikel berdasarkan status, author, dan tag
- 🔄 Sistem komentar dengan threading (dalam development)
- 🔄 Sistem like & bookmark (dalam development)
- 🔄 Sistem following user (dalam development)
- 🔄 Notifikasi real-time (dalam development)

## 🛠 Tech Stack

- **Backend**: Go (Gin Framework)
- **Database**: PostgreSQL (Supabase)
- **ORM**: GORM
- **Authentication**: JWT + Supabase Auth
- **CORS**: gin-contrib/cors

## 📦 Structure

```
project/
├── config/          # Konfigurasi aplikasi
├── database/        # Database connection
├── handlers/        # API handlers
├── middleware/      # Custom middleware
├── models/          # Database models
├── script_db/       # Database schema
├── go.mod          # Go dependencies
├── main.go         # Entry point
└── README.md       # Dokumentasi
```

## ⚙️ Setup

### 1. Environment Variables

Buat file `.env` di root project:

```env
# Database Configuration
DATABASE_URL=postgresql://username:password@hostname:port/database_name

# Supabase Configuration  
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_ANON_KEY=your-supabase-anon-key

# JWT Configuration
JWT_SECRET=your-jwt-secret-key

# Server Configuration
PORT=8080
ENVIRONMENT=development
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Setup Database

Jalankan script SQL di `script_db/maindb.sql` ke database Supabase Anda.

### 4. Run Server

```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## 📡 API Endpoints

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/api/v1/articles` | Daftar artikel (dengan pagination & filter) |
| GET | `/api/v1/articles/:slug` | Detail artikel berdasarkan slug |

### Protected Endpoints (Perlu Authentication)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/articles` | Buat artikel baru |
| PUT | `/api/v1/articles/:slug` | Update artikel |
| DELETE | `/api/v1/articles/:slug` | Hapus artikel |

### Query Parameters untuk GET /articles

- `page`: Nomor halaman (default: 1)
- `limit`: Jumlah artikel per halaman (default: 10)
- `published`: Filter artikel published (`true`/`false`)
- `featured`: Filter artikel featured (`true`/`false`)
- `author_id`: Filter berdasarkan author ID
- `tag`: Filter berdasarkan tag slug

### Authentication

Gunakan Bearer token di header:
```
Authorization: Bearer <your_jwt_token>
```

## 🏗 Database Schema

Database schema lengkap tersedia di `script_db/maindb.sql` dengan tabel:

- `profiles`: User profiles dengan social media links
- `articles`: Artikel dengan content, metadata, dan statistics  
- `tags`: Tag sistem dengan color coding
- `article_tags`: Many-to-many relationship artikel dan tag
- `follows`: User following system
- `likes`: Like artikel
- `bookmarks`: Bookmark artikel
- `comments`: Komentar dengan threading support
- `comment_likes`: Like untuk komentar
- `notifications`: Sistem notifikasi
- `article_views`: Analytics tracking views
- `tag_follows`: Following tag tertentu

## 🚀 Deployment

### Vercel (Recommended)

1. Install Vercel CLI: `npm i -g vercel`
2. Buat file `vercel.json`:

```json
{
  "version": 2,
  "builds": [
    {
      "src": "main.go",
      "use": "@vercel/go"
    }
  ],
  "routes": [
    {
      "src": "/(.*)",
      "dest": "/main.go"
    }
  ]
}
```

3. Deploy: `vercel --prod`

### Environment Variables di Vercel

Set semua environment variables yang ada di `.env` melalui Vercel dashboard.

## 🔮 Roadmap

- [ ] Implementasi handlers untuk komentar
- [ ] Implementasi like & bookmark functionality  
- [ ] User management (profile, following)
- [ ] Sistem notifikasi real-time
- [ ] Search functionality
- [ ] Rate limiting
- [ ] File upload untuk cover image
- [ ] Email notifications
- [ ] Analytics dashboard

## 🤝 Contributing

1. Fork repository
2. Buat branch feature (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push ke branch (`git push origin feature/amazing-feature`)
5. Buat Pull Request

## 📄 License

MIT License - lihat file LICENSE untuk detail lengkap. 