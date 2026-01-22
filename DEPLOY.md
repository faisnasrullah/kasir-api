# Deployment Guide - Railway

Panduan lengkap untuk deploy Kasir API ke Railway.

## Prerequisites
1. Akun Railway (gratis sign up di https://railway.app)
2. Railway CLI (optional, tapi memudahkan)
3. Git repository dengan project ini

## Method 1: Deploy dari GitHub (Recommended)

### Step 1: Push ke GitHub
```bash
git init
git add .
git commit -m "Initial commit: Kasir API"
git remote add origin https://github.com/yourusername/kasir-api.git
git push -u origin main
```

### Step 2: Login ke Railway
Buka https://dashboard.railway.app

### Step 3: Create New Project
1. Klik "New Project"
2. Pilih "Deploy from GitHub"
3. Connect GitHub account jika belum
4. Pilih repository `kasir-api`

### Step 4: Configure
1. Railway akan auto-detect Go project
2. Klik "Deploy"
3. Tunggu hingga deployment selesai

### Step 5: Access Application
- Cek domain yang diberikan Railway
- Swagger UI: `https://<domain>/docs/`
- Health check: `https://<domain>/health`

---

## Method 2: Deploy dengan Railway CLI

### Step 1: Install Railway CLI
```bash
# macOS
brew install railway

# Linux
curl -fsSL cli.new.railway.app/install.sh | bash
```

### Step 2: Login
```bash
railway login
```

### Step 3: Initialize Railway Project
```bash
cd /path/to/kasir-api
railway init
```

### Step 4: Deploy
```bash
railway up
```

### Step 5: View Logs
```bash
railway logs
```

### Step 6: View Service URL
```bash
railway service
```

---

## Environment Variables

Jika perlu menambah env variables di Railway:

1. Buka project di Railway Dashboard
2. Pilih service > Variables
3. Tambah variable baru
4. Trigger redeploy dengan git push

Contoh variables yang mungkin dibutuhkan (untuk development):
```
PORT=3000  # Railway akan set ini otomatis
```

---

## Monitoring & Logs

### Dari Dashboard
1. Buka Railway Dashboard
2. Pilih project > service
3. Klik "Deployments" untuk melihat history
4. Klik "Logs" untuk melihat real-time logs

### Dari CLI
```bash
# Tail logs
railway logs --follow

# View specific deployment logs
railway logs --deployment <deployment-id>
```

---

## Custom Domain (Optional)

1. Buka project di Railway Dashboard
2. Pilih service
3. Klik "Settings" > "Custom Domain"
4. Enter domain Anda
5. Configure DNS records sesuai instruksi Railway

---

## Troubleshooting

### Build Failed
- Check logs di Railway Dashboard
- Pastikan `go.mod` dan `go.sum` sudah commit
- Verify Dockerfile syntax: `docker build -t kasir-api .`

### App Crashes
- Check "Logs" di Railway Dashboard
- Verify PORT environment variable dibaca dengan benar
- Check health endpoint: `https://<domain>/health`

### Database Connection Issues
- Railway memiliki PostgreSQL plugin yang bisa ditambahkan
- Untuk dev, gunakan in-memory storage (saat ini)

---

## Database Integration (Future)

Jika ingin tambah PostgreSQL:

1. Di Railway Dashboard, buka project
2. Klik "+ Add" > "Database"
3. Pilih "PostgreSQL"
4. Railway akan auto-inject DATABASE_URL env var
5. Update code untuk connect ke database

---

## Cost & Limits

- Railway memberikan $5 free credits per bulan
- Kasir API dengan in-memory storage sangat hemat
- Monitor usage di: https://railway.app/account/billing

---

## Useful Links
- Railway Docs: https://docs.railway.app
- Railway CLI: https://docs.railway.app/develop/cli
- Go Deployment: https://docs.railway.app/deploy/golang
