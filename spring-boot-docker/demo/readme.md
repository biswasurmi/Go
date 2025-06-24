✅ Why Multi-Stage Builds?
Traditional Docker Builds:
Everything (dependencies, compiler, tools, runtime) is bundled together.

Results in large image sizes.

Higher security risks and slower deployments.

Multi-Stage Builds:
Split the build and runtime environments.

Final image contains only what is needed to run.

Smaller, faster, and more secure images.



*************** volume *******************urmi@urmi-HP-Laptop:~/Desktop/Go/spring-boot-docker/demo$ docker run --name=db \
  -e POSTGRES_PASSWORD=secret \
  -d \
  -v postgres_data:/var/lib/postgresql/data \
  postgres
Unable to find image 'postgres:latest' locally
latest: Pulling from library/postgres
dad67da3f26b: Already exists 
eb3a531023c8: Pull complete 
05b641b3bdab: Pull complete 
64e8f1b2b243: Pull complete 
603ef9fcdd8e: Pull complete 
8a1f652e0c97: Pull complete 
c6def2c6e21d: Pull complete 
b47a445a47f0: Pull complete 
c95f49cc11b3: Pull complete 
3664068a9b37: Pull complete 
abfd68ef219e: Pull complete 
928d00623a6e: Pull complete 
db3ab53631e4: Pull complete 
f4ce9941f6e3: Pull complete 
Digest: sha256:6cf6142afacfa89fb28b894d6391c7dcbf6523c33178bdc33e782b3b533a9342
Status: Downloaded newer image for postgres:latest
32e9d63f4553406cd11774269ca3df659bf94b466a35c0ce0f7c2c891b14a7d2
urmi@urmi-HP-Laptop:~/Desktop/Go/spring-boot-docker/demo$ docker exec -ti db psql -U postgres
psql (17.5 (Debian 17.5-1.pgdg120+1))
Type "help" for help.

postgres=# 
postgres=# CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    description VARCHAR(100)
);

INSERT INTO tasks (description) VALUES ('Finish work'), ('Have fun');
CREATE TABLE
INSERT 0 2
postgres=# SELECT * FROM tasks;
 id | description 
----+-------------
  1 | Finish work
  2 | Have fun
(2 rows)

postgres=# \q
urmi@urmi-HP-Laptop:~/Desktop/Go/spring-boot-docker/demo$ docker stop db
db
urmi@urmi-HP-Laptop:~/Desktop/Go/spring-boot-docker/demo$ docker rm db
db
urmi@urmi-HP-Laptop:~/Desktop/Go/spring-boot-docker/demo$ docker run --name=new-db \
  -d \
  -v postgres_data:/var/lib/postgresql/data \
  postgres
4fd3a4e19bc9fb5e1e9b182a22ec9ebff3cfe3108143511b75c36a107c2d0ee0
urmi@urmi-HP-Laptop:~/Desktop/Go/spring-boot-docker/demo$ docker exec -ti new-db \
  psql -U postgres -c "SELECT * FROM tasks"
 id | description 
----+-------------
  1 | Finish work
  2 | Have fun
(2 rows)

urmi@urmi-HP-Laptop:~/Desktop/Go/spring-boot-docker/demo$ docker volume inspect postgres_data
[
    {
        "CreatedAt": "2025-06-24T15:28:02+06:00",
        "Driver": "local",
        "Labels": null,
        "Mountpoint": "/var/lib/docker/volumes/postgres_data/_data",
        "Name": "postgres_data",
        "Options": null,
        "Scope": "local"
    }
]
urmi@urmi-HP-Laptop:~/Desktop/Go/spring-boot-docker/demo$ 


******************* volume vs bind mount ***********************

✅ PART 1: THEORY — Docker Bind Mounts vs Volumes
🧠 The Problem
Containers run in isolation and cannot access your host files directly. But often:

You want to share files between the container and your host

Or you want containers to persist files across restarts or rebuilds

Docker gives you two solutions:

🔷 1. Volumes
Managed by Docker (docker volume create)

Data is stored outside the container, in Docker-managed space

Used for: persistent storage, like a database

Data survives container removal

Docker manages location and permissions

Good for: Production environments

🔶 2. Bind Mounts
Mounts a host file/folder directly into a container

The host path is explicitly specified (e.g., /home/urmi/public_html:/usr/local/apache2/htdocs/)

Ideal for development, so changes on the host reflect instantly inside the container

More flexible and manual

Can use :ro (read-only) or :rw (read-write) at the end

🆚 Comparison Table
Feature	Volume	Bind Mount
Data managed by	Docker	You (host path explicitly specified)
Use case	Persistent app data (DBs, logs)	Live development files, config files
Host control	Minimal (Docker-managed)	Full (you control file location)
Portability	More portable	Less portable (host-specific paths)
Ease of use	Simpler for long-term storage	Flexible for development


