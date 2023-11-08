---
created: 2023-11-07T11:11:03.102Z
updated: 2023-11-07T15:40:45.604Z
assigned: ""
progress: 0
tags:
  - 'Main Feature'
---

# User / Admin Sign In

MUST USE COOKIES FOR SESSIONS

- Create a login session to access the forum and be able to add posts and comments
- Use cookies to allow each user to have only one opened session
- Each of this sessions must contain an expiration date. It is up to you to decide how long the cookie stays "alive"
- Must be able to check if the email provided is present in the database and if all credentials are correct
- Must check if the password is the same with the one provided and, if the password is not the same, it will return an error response

## Sub-tasks

- [ ] Backend
- [ ] Frontend

## Comments

- author: aabdeen
  date: 2023-11-07T15:40:35.404Z
  - Use struct if there is no access to table
  - Build the struct to match proposed table attributes
