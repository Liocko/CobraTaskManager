
# 🐍 Cobra Task Manager

Простой менеджер задач с CLI-интерфейсом, написанный на Go с использованием **Cobra** и **SQLite**.

Теперь данные хранятся в SQLite-базе, а не в JSON.

![](https://media.tenor.com/TCMWkxIkF9IAAAAj/dancing-gopher.gif)

## 📦 Установка

1. **Клонируйте репозиторий**:
```
git clone https://github.com/Liocko/CobraTaskManager.git
cd CobraTaskManager
```
Соберите проект:

```
go build -o cobraTaskManager main.go
```
Добавьте алиас для удобства (опционально):
Добавьте в ваш .bashrc/.zshrc:


```
alias taskmgr="/путь/к/проекту/cobraTaskManager"
```
Или установите бинарник в системный путь:
```
sudo cp cobraTaskManager /usr/local/bin  # Теперь можно вызывать из любой папки!
```
## 🚀 Использование

**Основные команды:**

###### (все комманды прописаны уже с алиас)

### 1. Добавить задачу
```
taskmgr add "Название | Описание"
```
### 2. Показать список задач
```
taskmgr list
```

### 3. Отметить задачу выполненной
````
taskmgr done 2
````

### 4. Удалить задачу
```
taskmgr remove 3
```
Пример вывода:
```
┌────┬──────────────────────┬─────────────────────────────┬────────┐
│ ID │      Task Name       │       Task Description      │ Status │
├────┼──────────────────────┼─────────────────────────────┼────────┤
│ 1  │ Покупки              │ Купить молоко и хлеб        │   ✓    │
│ 2  │ Изучение Go          │ Разобраться с горутинами    │   ✗    │
└────┴──────────────────────┴─────────────────────────────┴────────┘
```
🔧 Технологии
Go 1.21+

Cobra (CLI-фреймворк)

SQLite (хранение данных)

go-pretty (табличный вывод)

📁 Особенности
SQLite вместо JSON
Теперь все задачи хранятся в локальной SQLite-базе (data.db)

Автоматическая инициализация
База данных создается автоматически при первом запуске.

Алиасы
Можно настроить короткие команды (например, taskmgr вместо полного пути).

🛠 В планах:

* Коррекция нумерации


* Пользователи и авторизация


* Улучшение взаимодействия
