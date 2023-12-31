# simple-chat
A simple chat application to play with websockets. Made with svelte frontend and rust/go switchable backends.

## Dependencies
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) (Because it uses the new `docker compose up` command).

## Init
### 1. Clone the repo
```bash
git clone https://github.com/MarkelCA/simple-chat
cd simple-chat
```
### 2. Create the envs file
*(From the root folder)*
```bash
cp .env.example .env
```
Edit the .envs file with the desired parameters. You'll have to change the variables inside the <> chars. **Don't change the already defined ones, like the `VITE_SERVER_PORT`**

## Frontend
(From the root folder)

```bash
npm run dev
```

## Backend
There's a script to run the backend server.
```bash
./start-server <go|rust>
```
