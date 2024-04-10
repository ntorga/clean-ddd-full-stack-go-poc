
# Full Stack Go with Clean Architecture and DDD: A Proof of Concept

## Introduction
> The center of your application is not the database. Nor is it one or more of the frameworks you may be using. **The center of your application is the use cases of your application**  -  _Unclebob_ ([source](https://blog.8thlight.com/uncle-bob/2012/05/15/NODB.html "NODB"))

This project builds upon the principles demonstrated in the [PHP version](https://github.com/ntorga/clean-ddd-php-poc-contacts) of our Clean Architecture and Domain-Driven Design (DDD) Proof of Concept (PoC). Before diving into this Go implementation, we recommend familiarizing yourself with the PHP version to grasp the foundational concepts that drive this project.

![Architecture](./architecture.jpg)

## Objective
The primary aim with this Go-based PoC is to demonstrate how Clean Architecture and DDD principles can be effectively implemented in a full stack Go project. However, I wanted to do things a bit differently this time around.

The PHP version of the PoC was a simple REST JSON API application using files. In this version, we also have:

- **Database**: Data needs to be in an actual database. SQLite was chosen for its simplicity.
- **CLI**: A command-line interface (CLI) to interact with the application.
- **Server Side Rendering (SSR)**: Real DOM from the server side with reactivity, but without a framework such as Next.js (React) or Nuxt.js (Vue) AND without having to create endpoints for HTML fragments.

## Technologies
- Go
- Echo
- Cobra
- GORM
- SQLite
- Templ
- Unpoly

## Q&A

**Why not use HTMX + Alpine.js?**

Although HTMX and Alpine.js are a great combo, I chose Unpoly cause I didn't want to return HTML fragments from the server, I wanted to morph the DOM on the client side using the return of the updated page.

All the routes that return HTML are GET only, so no need to duplicate controllers, I just need a presenter on the UI layer. The write operations, such as POST, PUT and DELETE consumes the JSON REST API and returns JSON, no HTML fragments. 🤨

**Why not Next.js (React) or Nuxt.js (Vue)?**

I wanted to keep things simple and avoid the complexity of a full-fledged framework. I wanted to see if I could achieve the same result with a simpler approach, without returning JSON that would be transformed into HTML on the server side and then sent to the client.

**Why not PostgreSQL, MySQL or MongoDB?**

No need for a heavy database for this PoC. SQLite is more than enough for this project. The focus is on the architecture, not the database. However, since we are using GORM, it would be easy to switch to another database if needed.

**Why not Python, Ruby, PHP or JavaScript/TypeScript?**

I think Go has a great developer experience, it's fast, has a great standard library, and it's easy to deploy. With Go, the entire project is a single binary. No need for a runtime or a virtual environment. REST API, CLI and HTML with a single command.