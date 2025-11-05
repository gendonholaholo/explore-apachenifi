# NiFi Parser - Dokumentasi Teknis

## Tujuan Pembangunan

NiFi Parser dikembangkan untuk mengatasi kompleksitas dalam menganalisis dan memvalidasi konfigurasi Apache NiFi flow tanpa memerlukan akses ke NiFi UI atau server yang sedang berjalan. Tool ini menyediakan solusi command-line yang cepat dan efisien untuk inspeksi konfigurasi NiFi dalam format yang human-readable.

## Deskripsi Fungsional

NiFi Parser adalah utilitas CLI berbasis Go yang melakukan parsing terhadap file konfigurasi Apache NiFi dalam format JSON dan menghasilkan laporan terstruktur yang mencakup:

- **Flow Summary**: Ringkasan metadata flow (ID, nama, versi)
- **Processor Analysis**: Detail setiap processor termasuk tipe, status, dan properti konfigurasi
- **Connection Mapping**: Visualisasi hubungan antar processor dalam flow
- **Runtime Statistics**: Metrik agregat tentang status operational processor

Tool ini mengimplementasikan flag-based filtering untuk memungkinkan pengguna fokus pada aspek spesifik dari konfigurasi yang relevan dengan kebutuhan analisis mereka.

## Manfaat Operasional

### 1. Efisiensi Workflow
Mengeliminasi kebutuhan untuk membuka NiFi UI atau parsing manual file JSON yang kompleks, mengurangi waktu inspeksi konfigurasi dari menit menjadi detik.

### 2. Validasi Pre-Deployment
Memungkinkan validasi konfigurasi flow sebelum deployment ke environment production, mengurangi risiko error runtime akibat misconfiguration.

### 3. Dokumentasi & Audit
Menghasilkan snapshot tekstual dari konfigurasi flow yang dapat disimpan untuk keperluan dokumentasi, version control, atau compliance audit.

### 4. Debugging & Troubleshooting
Menyediakan quick overview status processor untuk identifikasi cepat komponen yang stopped atau misconfigured tanpa akses ke live system.

### 5. CI/CD Integration
Dapat diintegrasikan ke dalam pipeline CI/CD untuk automated configuration validation sebagai bagian dari quality assurance process.

## Target Pengguna

### Primary Users

**1. Data Engineers**
- Mengembangkan dan maintain NiFi data pipelines
- Memerlukan validasi konfigurasi sebelum deployment
- Membutuhkan quick reference untuk flow architecture

**2. DevOps Engineers**
- Mengelola NiFi infrastructure dan deployment
- Melakukan troubleshooting operational issues
- Mengimplementasikan automated configuration checks

**3. Platform Engineers**
- Membangun tooling dan automation untuk data platform
- Mengintegrasikan NiFi configuration validation ke dalam CI/CD
- Memonitor compliance terhadap configuration standards

### Secondary Users

**4. Solution Architects**
- Melakukan review terhadap flow design
- Memvalidasi adherence terhadap architectural guidelines
- Mendokumentasikan data flow patterns

**5. QA Engineers**
- Melakukan testing terhadap NiFi configurations
- Memverifikasi deployment artifacts sebelum release
- Mengotomasi configuration validation tests

## Spesifikasi Teknis

### Requirements
- Go 1.21 atau lebih tinggi
- Input: JSON-formatted NiFi flow configuration
- Platform: Cross-platform (Linux, macOS, Windows)

### Key Features
- Zero external dependencies (pure Go standard library)
- Stateless operation (no database atau persistent storage)
- Read-only analysis (tidak memodifikasi file input)
- Flexible output filtering via command-line flags

## Use Case Examples

### Use Case 1: Pre-Deployment Validation
```bash
nifiparser -file production_flow.json -stats
# Output: Quick check untuk memastikan semua processor running
```

### Use Case 2: Flow Documentation
```bash
nifiparser -file dataflow.json > flow_documentation.txt
# Output: Tekstual snapshot untuk version control
```

### Use Case 3: Troubleshooting
```bash
nifiparser -file current_flow.json -processors
# Output: Identifikasi processor yang stopped atau error
```

### Use Case 4: CI/CD Pipeline
```bash
# Dalam GitLab CI/Jenkins
nifiparser -file configs/*.json || exit 1
# Output: Fail build jika ada configuration issue
```

## Roadmap & Future Enhancements

- **XML Support**: Parsing untuk XML-formatted NiFi configurations
- **Schema Validation**: Validasi terhadap NiFi schema definitions
- **Diff Tool**: Membandingkan dua versi flow configuration
- **Export Formats**: JSON, YAML, atau Markdown output
- **REST API Integration**: Direct fetch dari NiFi REST API

## Kontribusi

Tool ini dikembangkan sebagai bagian dari pembelajaran Go programming dengan fokus pada:
- Idiomatic Go project structure
- Comprehensive testing dan code quality
- Production-ready tooling practices
- Developer experience optimization

---

**Version**: 1.0.0  
**Language**: Go 1.21+  
**License**: MIT  
**Maintainer**: Ghaws Shafadonia
