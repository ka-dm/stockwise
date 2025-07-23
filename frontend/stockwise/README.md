# stockwise

This template should help get you started developing with Vue 3 in Vite.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) to make the TypeScript language service aware of `.vue` types.

## Customize configuration

See [Vite Configuration Reference](https://vite.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```

---

## 1. `frontend/Dockerfile` para desarrollo

```dockerfile
FROM node:20-alpine3.19

WORKDIR /app

# Instala dependencias primero (mejor cache)
COPY package*.json ./
RUN npm install

# Copia el resto del código (en desarrollo, esto se sobreescribe con el volumen)
COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host"]
```

---

## 2. Sección frontend en `docker-compose.yml`

```yaml
frontend:
  build: ./frontend
  ports:
    - "5173:5173"
  volumes:
    - ./frontend:/app
    - /app/node_modules # Evita que node_modules de tu host sobrescriba los del contenedor
  depends_on:
    - backend
  restart: unless-stopped
```

---

### ¿Por qué así?

- El comando `npm run dev -- --host` permite que Vite acepte conexiones externas (necesario en Docker).
- El volumen `./frontend:/app` hace que los cambios en tu código se reflejen instantáneamente en el contenedor.
- El volumen `/app/node_modules` evita conflictos entre dependencias del host y del contenedor.

---

## ¿Y para producción?

Para producción, usa el Dockerfile que construye y sirve con Nginx (como el que ya tienes).

---

¿Te gustaría que te deje ambos Dockerfile (dev y prod) o solo el de desarrollo?  
¿Quieres que te ayude a alternar entre ambos fácilmente?
