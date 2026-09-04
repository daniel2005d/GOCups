# GOCups

Herramienta de línea de comandos en Go para la enumeración y auditoría de impresoras y servicios expuestos en servidores **CUPS** (*Common Unix Printing System*).

---

## Características

* **Consulta HTTP nativa:** Interacción directa con los endpoints de gestión de CUPS.
* **Extracción de datos
  * Nombre de la cola de impresión
  * Descripción
  * Ubicación física
  * Fabricante y modelo
  * Estado actual (*Idle*, *Processing*, *Stopped*)
  * JOBS

* **Soporte multiplataforma:** Compilación estática sin dependencias de runtime para Linux, Windows y macOS.

---

## Requisitos previos

* [Go](https://go.dev/dl/) 1.21 o superior instalado en el sistema.

---

# Binarios

Descargar los binarios de la carpeta [bin](https://github.com/daniel2005d/GOCups/tree/main/bin)

## Instalación y compilación

### Clonar el repositorio

```bash
git clone https://github.com/daniel2005d/GOCups.git
cd gocups
```

### Instalar dependencias

```bash
go mod tidy
```

### Compilación rápida

```bash
# Compilar binario nativo para el sistema actual
go build -ldflags="-s -w" -trimpath -o bin/gocups .
```

### Compilación cruzada

Si cuentas con `make` (Linux/macOS):

```bash
make build-all     # Compila binarios para Linux, Windows y macOS
make build-linux   # Solo binario Linux amd64
make build-windows # Solo binario Windows amd64
```

Si ejecutas desde Windows (PowerShell):

```powershell
.\build.ps1 -Target all
```

---

## Uso

```bash
./bin/gocups -host 10.10.10.15 [-port]
```

### Opciones y flags

| Parámetro | Tipo | Valor por defecto | Descripción |
| :--- | :--- | :--- | :--- |
| `-host` | `string` | `127.0.0.1` | Dirección IP o nombre de host del servidor CUPS |
| `-port` | `int` | `631` | Puerto donde escucha el servicio CUPS |
| `-timeout`| `duration` | `10s` | Tiempo máximo de espera para la solicitud |

---

## Ejemplo de salida

```text
───────────────────────────────────────────────────────────────────────────
  Name:         PDF
  Description:  None
  Location:     Local
  Make & Model: Lexmark C9200 Foomatic/hpijs-pcl5e
  Status:       Processing - "Unable to locate printer \"hostname\"."
───────────────────────────────────────────────────────────────────────────
  ACTIVE PRINT JOBS:
┌───────┬─────────┬──────────┬──────┬───────┬────────────────────────────────────────────────────────────┐
│ ID    │ NAME    │ USER     │ SIZE │ PAGES │ STATE                                                      │
├───────┼─────────┼──────────┼──────┼───────┼────────────────────────────────────────────────────────────┤
│ PDF-1 │ Unknown │ Withheld │ 1k   │ 1     │ processing since Fri Sep 4 20:46:41 2026 "Filter failed"   │
└───────┴─────────┴──────────┴──────┴───────┴────────────────────────────────────────────────────────────┘
───────────────────────────────────────────────────────────────────────────
```
