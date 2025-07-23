# Stockwise

Sistema para consultar, almacenar y recomendar acciones bursátiles usando Golang, Vue 3, CockroachDB y Tailwind CSS.

## Estructura del proyecto

```
stockwise/
│
├── backend/                # Backend en Golang
│   ├── cmd/                # Entrypoints
│   ├── internal/           # Lógica interna
│   │   ├── api/            # Handlers HTTP/API
│   │   ├── db/             # Acceso a CockroachDB
│   │   ├── models/         # Modelos
│   │   ├── services/       # Lógica de negocio
│   │   └── utils/          # Utilidades
│   └── pkg/                # Código exportable
│
├── frontend/               # Frontend en Vue 3
│   ├── public/             # Archivos estáticos
│   └── src/
│       ├── assets/         # Imágenes, fuentes
│       ├── components/     # Componentes Vue
│       ├── views/          # Vistas principales
│       ├── store/          # Pinia stores
│       ├── router/         # Rutas
│       ├── api/            # Llamadas a la API
│       └── types/          # Tipos TypeScript
│
├── migrations/             # Migraciones de base de datos
├── scripts/                # Scripts útiles
├── tests/                  # Pruebas
├── README.md
└── docker-compose.yml      # Orquestación de servicios
```

## Tecnologías
- Backend: Golang
- Frontend: Vue 3, TypeScript, Pinia, Tailwind CSS
- Base de datos: CockroachDB 