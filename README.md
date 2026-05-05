# grscan

Мини-утилита для поиска Git-репозиториев по remote (имени или URL) внутри директории.

## Возможности

* Рекурсивно обходит папки
* Находит все `.git` репозитории
* Проверяет `git remote -v`
* Ищет по:

  * имени remote (`person-companies`)
  * или части URL (`gitlab.khc.kz`)

---

## Установка

```bash
go mod init git-find
go build -o git-find
```

---

## Использование

### Базовый запуск

```bash
./git-find -q person-companies
```

Ищет в текущей директории (`.`)

---

### Указать путь

```bash
./git-find -path ~/projects -q person-companies
```

---

## Аргументы

| Флаг    | Описание                       | По умолчанию       |
| ------- | ------------------------------ | ------------------ |
| `-path` | путь для поиска                | `.`                |
| `-q`    | строка поиска (remote или URL) | `person-companies` |

---

## Примеры

### Найти по имени remote

```bash
./git-find -q person-companies
```

---

### Найти по URL

```bash
./git-find -q gitlab.khc.kz
```

---

## Пример вывода

```bash
FOUND: /Users/user/projects/repo1
origin  git@gitlab.khc.kz:repo.git (fetch)
origin  git@gitlab.khc.kz:repo.git (push)
-----------
```

---

## Как работает

1. Обходит директорию (`filepath.WalkDir`)
2. Ищет папки `.git`
3. Для каждой:

   * выполняет `git remote -v`
   * проверяет совпадение с `-q`
4. Выводит совпадения

---

## Требования

* Go 1.18+
* установленный `git` в системе

---

## Идеи для улучшения

* Параллельный обход (ускорение)
* JSON вывод
* фильтр по branch
* вывод только путей (`--quiet`)
* чтение `.git/config` вместо вызова git

---

## Лицензия

MIT
