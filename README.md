#rest-api-tutorial

#user-serveice

#REST API

GET /users -- list of users -- 200, 404, 500
GET /users/:id -- user by id -- 200, 404, 500
POST /users/:id -- create user -- 204, 4xx, Header Location: url
PUT /users/:id -- fully update user -- 200/204
PATCH /users/:id -- partially udpate user -- 200/204, 400, 404, 500
DELETE /users/:id -- delete user by id-- 204, 400, 404 