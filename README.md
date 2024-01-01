# Checkout Journey

This project aims to simulate a checkout system. It uses event-driven architecture to simulate the checkout process. Event sourcing is used to store the state of the checkout system. 
It consist of these technologies:

- Kafka
- Go + Fiber
- Docker
- Next.js(Client)

It aims to be a simple example of how things can be done with kafka and event sourcing. Not a ready to use project for production.

## Diagram

Whenever a checkout is created, a checkout event is published to the `checkout` topic. There are consumers that listen to this topic and do their work. 

![diagram](./diagram.png)

## How to run
```
docker-compose up -d
```

## How to test
```
cd client
pnpm i
pnpm dev
```

Go to `localhost:3000` and you will see the client app. Click the Checkout button, check the inventory, shipment and notification services. You will see the related events in the console.
