# simple-chat

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
cp .env.sample .env
```
Edit the .envs file with the desired parameters

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
