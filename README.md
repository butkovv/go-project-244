### Hexlet tests and linter status:
[![Actions Status](https://github.com/butkovv/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/butkovv/go-project-244/actions)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=butkovv_go-project-244&metric=coverage)](https://sonarcloud.io/summary/new_code?id=butkovv_go-project-244)

## gendiff

CLI‑утилита для сравнения двух конфигурационных файлов и вывода различий.

### Установка

```sh
go install ./cmd/gendiff
```

### Использование

```sh
gendiff [--format FORMAT] <path1> <path2>
```

### Примеры

```sh
gendiff testdata/fixture/file1.json testdata/fixture/file2.yaml
```

```sh
gendiff --format plain testdata/fixture/file1.json testdata/fixture/file2.yaml
```

```sh
gendiff --format json testdata/fixture/file1.json testdata/fixture/file2.yaml
```

### Форматы вывода

- `stylish` (по умолчанию)
- `plain`
- `json`

### Поддерживаемые файлы

- JSON (`.json`)
- YAML (`.yaml`, `.yml`)

### Коды завершения

- `0` при успехе
- не‑ноль при ошибке
