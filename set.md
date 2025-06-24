````markdown
# Getting Started with the Todo App (Docker & MySQL)

This README guides you step-by-step to clone, build, run, and manage a multi-container Todo app using Docker and Docker Compose.

---

## ✅ 1. Clone the App Repository

Open your terminal and run:

```bash
git clone https://github.com/docker/getting-started-app.git
cd getting-started-app
````

---

## ✅ 2. View the Project Structure

Run:

```bash
ls -l
```

Expected output:

```
.dockerignore
Dockerfile     ← you’ll create this
package.json
README.md
spec/
src/
yarn.lock
```

---

## ✅ 3. Create a Dockerfile

Create a file named `Dockerfile` in the root folder (`getting-started-app/`) with:

```dockerfile
# syntax=docker/dockerfile:1

FROM node:lts-alpine
WORKDIR /app
COPY . .
RUN yarn install --production
CMD ["node", "src/index.js"]
EXPOSE 3000
```

---

## ✅ 4. Build the Image

Build the Docker image inside the `getting-started-app/` directory:

```bash
docker build -t getting-started .
```

**Explanation:**

* `-t getting-started`: Tags the image as `getting-started`.
* `.` : Use current directory as build context.

---

## ✅ 5. Run the App in a Container

```bash
docker run -d -p 127.0.0.1:3000:3000 getting-started
```

**Explanation:**

* `-d`: Run container in detached (background) mode.
* `-p 127.0.0.1:3000:3000`: Map port 3000 in container to localhost port 3000.

---

## ✅ 6. View the App in Browser

Open your browser to:

```
http://localhost:3000
```

You should see the todo app UI. Try adding todo items.

---

## ✅ 7. Check Running Containers

```bash
docker ps
```

Sample output:

```
CONTAINER ID   IMAGE           COMMAND               PORTS                    NAMES
a1b2c3d4e5f6   getting-started "node src/index.js"   127.0.0.1:3000->3000/tcp  hopeful_morse
```

---

## ✅ 8. Push Image to Docker Hub (Optional)

1. Sign in at [Docker Hub](https://hub.docker.com/), create a repository named `getting-started`.

2. Tag your image with your Docker Hub username:

```bash
docker tag getting-started urmibiswas/getting-started
```

3. Log in via CLI:

```bash
docker login
```

4. Push image:

```bash
docker push urmibiswas/getting-started
```

---

## ✅ 9. Create a Docker Volume for Persistence

Create volume:

```bash
docker volume create todo-db
```

List volumes:

```bash
docker volume ls
```

---

## ✅ 10. Remove Old Containers

```bash
docker ps           # find running container IDs
docker rm -f <CONTAINER_ID>
```

---

## ✅ 11. Run Container with Volume Mounted

```bash
docker run -dp 127.0.0.1:3000:3000 \
  --mount type=volume,src=todo-db,target=/etc/todos \
  getting-started
```

**Explanation:**

Mount volume `todo-db` inside container at `/etc/todos` (where DB file is saved).

---

## ✅ 12. Test Data Persistence

Stop container:

```bash
docker ps
docker rm -f <CONTAINER_ID>
```

Restart with same volume:

```bash
docker run -dp 127.0.0.1:3000:3000 \
  --mount type=volume,src=todo-db,target=/etc/todos \
  getting-started
```

Open app and verify old todo items still exist.

---

## ✅ 13. Use Bind Mount for Live Development

Run container with bind mount and nodemon:

```bash
docker run -dp 127.0.0.1:3000:3000 \
  -w /app \
  --mount type=bind,src="$(pwd)",target=/app \
  node:18-alpine \
  sh -c "yarn install && yarn run dev"
```

Check logs:

```bash
docker logs -f <container-id>
```

Expected:

```
[nodemon] starting `node src/index.js`
Using sqlite database at /etc/todos/todo.db
Listening on port 3000
```

---

## ✅ 14. Make a Live Code Change

Edit file: `src/static/js/app.js`

Change line near 109 from:

```js
{submitting ? 'Adding...' : 'Add Item'}
```

To:

```js
{submitting ? 'Adding...' : 'Add'}
```

Save, nodemon will detect the change, app reloads automatically.

---

## ✅ 15. Create Custom Docker Network

```bash
docker network create todo-app
```

---

## ✅ 16. Run MySQL Container on Network

```bash
docker run -d \
  --network todo-app --network-alias mysql \
  -v todo-mysql-data:/var/lib/mysql \
  -e MYSQL_ROOT_PASSWORD=secret \
  -e MYSQL_DATABASE=todos \
  mysql:8.0
```

---

## ✅ 17. Verify MySQL is Running

Find container ID:

```bash
docker ps
```

Connect to MySQL:

```bash
docker exec -it <mysql-container-id> mysql -u root -p
```

Enter password: `secret`

List databases:

```sql
SHOW DATABASES;
```

Expected output:

```
information_schema
mysql
performance_schema
sys
todos
```

Exit MySQL shell:

```sql
exit
```

---

## ✅ 18. Run Node.js App Connected to MySQL

```bash
docker run -dp 127.0.0.1:3000:3000 \
  -w /app -v "$(pwd):/app" \
  --network todo-app \
  -e MYSQL_HOST=mysql \
  -e MYSQL_USER=root \
  -e MYSQL_PASSWORD=secret \
  -e MYSQL_DB=todos \
  node:18-alpine \
  sh -c "yarn install && yarn run dev"
```

---

## ✅ 19. Check App Logs

```bash
docker logs -f <app-container-id>
```

Expected:

```
Connected to mysql db at host mysql
Listening on port 3000
```

Open [http://localhost:3000](http://localhost:3000) and test.

---

## ✅ 20. Use Docker Compose for Multi-Container Setup

Create `compose.yaml` with:

```yaml
services:
  app:
    image: node:18-alpine
    command: sh -c "yarn install && yarn run dev"
    ports:
      - 127.0.0.1:3000:3000
    working_dir: /app
    volumes:
      - ./:/app
    environment:
      MYSQL_HOST: mysql
      MYSQL_USER: root
      MYSQL_PASSWORD: secret
      MYSQL_DB: todos

  mysql:
    image: mysql:8.0
    volumes:
      - todo-mysql-data:/var/lib/mysql
    environment:
      MYSQL_ROOT_PASSWORD: secret
      MYSQL_DATABASE: todos

volumes:
  todo-mysql-data:
```

---

## ✅ 21. Run App Stack with Compose

Stop old containers:

```bash
docker ps -a
docker rm -f <container-ids>
```

Start app stack:

```bash
docker compose up -d
```

Expected output:

```
Creating network "getting-started-app_default" ...
Creating volume "getting-started-app_todo-mysql-data" ...
Creating getting-started-app_mysql_1 ...
Creating getting-started-app_app_1 ...
```

---

## ✅ 22. View Logs

For all services:

```bash
docker compose logs -f
```

Expected:

```
mysql_1  | ... ready for connections ...
app_1    | Connected to mysql db at host mysql
app_1    | Listening on port 3000
```

---

## ✅ 23. Open the App

Visit:

```
http://localhost:3000
```

Add todos and verify persistence.

---

## ✅ 24. Inspect Volumes or MySQL DB (Optional)

List volumes:

```bash
docker volume ls
docker volume inspect getting-started-app_todo-mysql-data
```

Access MySQL:

```bash
docker exec -it getting-started-app_mysql_1 mysql -uroot -p
```

Enter password `secret`.

Use DB and query:

```sql
USE todos;
SELECT * FROM todo_items;
```

---

## ✅ 25. Stop or Remove Everything

Stop app stack:

```bash
docker compose down
```

Stop and remove volumes as well:

```bash
docker compose down --volumes
```

---

# Summary

| Task                  | Command                         |
| --------------------- | ------------------------------- |
| Start app             | `docker compose up -d`          |
| View logs             | `docker compose logs -f`        |
| Stop app              | `docker compose down`           |
| Remove volumes & data | `docker compose down --volumes` |

---
