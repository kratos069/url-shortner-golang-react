<!-- How to use? -->
🧪 How to Test in Postman
POST http://localhost:8080/shorten
Body (JSON):

{
  "url": "https://example.com"
}

GET the short URL returned (e.g.):

http://localhost:8080/kratos69/aB12Xz

<!-- new migration file command -->
migrate create -ext sql -dir db/migration -seq add_users

<!-- github upload steps -->
git init
git add .
git commit -m "message for commit..."
git remote add origin https://github.com/your-username/your-repo-name.git
git remote -v
git branch -M main
git push -u origin main

<!-- (1) build docker image of project -->
docker build -t url-short:latest .

<!-- sql file from dbml file -->
dbml2sql --postgres -o doc/schema.sql doc/db.dbml

<!-- before merging changes with main branch -->
git checkout -b ft/newFeature
git add .
git commit -m "new feature added"
git push origin ft/newFeature
(go to github, create pull, merge and delete feature branch)
(back in terminal)
git checkout main
git pull