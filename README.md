rabbitmqmgmt
============
cli tool for rabbitmq queue/exchage/bindings management.
rabbitmqctl equivalent for create/delete/bind queues, exchanges, etc.


Subcommands
===========
   queue_add      add a new queue
   queue_remove      remove an existing queue
   queue_bind     bind a queue to a exchange using a ginven topic/routing key
   queue_unbind      remove an existing binding
   exchange_add      add a new exchange
   exchange_remove   remove an existing exchange
   help, h     Shows a list of commands or help for one command

Global options
==============
   --amqp_uri, -u 'amqp://guest:guest@localhost:5672/'   broker url (including vhost)
   --version, -v              print the version
   --help, -h                 show help


Examples:
=========

```rabbitmqmgmt queue_add --durable --x-message-ttl 5000 test_queue```
```rabbitmqmgmt queue_remove test_queue```
```rabbitmqmgmt queue_bind test_queue test_exchange test_routing_key```
```rabbitmqmgmt queue_unbind test_queue test_exchange test_routing_key```
```rabbitmqmgmt exchange_add --type topic --durable test_exchange```
```rabbitmqmgmt exchange_remove test_exchange```

