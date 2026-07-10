# Platform-as-a-Service

Простое REST API приложение "Task Manager", написанное на Go и задеплоенное в локальный Kubernetes кластер (kind).

Проект демонстрирует базовые навыки DevOps:
- контейнеризацию (Docker)
- CI/CD (GitHub Actions)
- оркестрацию (Kubernetes)
- сетевое взаимодействие внутри кластера

# 🚀 Технологии

- Go (Gin framework)
- Docker
- Kubernetes (kind)
- kubectl
- GitHub Actions (CI)



---

# 📡 API

| Метод  | Endpoint       | Описание             |
|--------|---------------|----------------------|
| GET    | /health       | Проверка состояния   |
| GET    | /tasks        | Получить все задачи  |
| GET    | /tasks/{id}   | Получить задачу      |
| POST   | /tasks        | Создать задачу       |
| PATCH  | /tasks/{id}   | Обновить задачу      |
| DELETE | /tasks/{id}   | Удалить задачу       |

---

# Запуск через Docker

```bash
docker build -t task-manager .
docker run -p 8080:8080 task-manager
```

# Запуск в Kubernetes (kind)
Создание кластера
kind create cluster --name task-manager --config kind-config.yml


kubectl apply -f deployment.yml


