rabbitmqmgmt
============
cli tool for rabbitmq queue/exchage/bindings management.
rabbitmqctl equivalent for create/delete/bind queues, exchanges, etc.

Subcommands
===========
 * **queue_add** add a new queue
 * **queue_remove** remove an existing queue
 * **queue_bind** bind a queue to a exchange using a ginven topic/routing key
 * **queue_unbind** remove an existing binding
 * **exchange_add** add a new exchange
 * **exchange_remove** remove an existing exchange
 
Global options
==============
   --amqp_uri, -u 'amqp://guest:guest@localhost:5672/'   broker url (including vhost)


Examples:
=========

Create a queue test_queue with a message ttl of 5000 ms and durable
```
rabbitmqmgmt queue_add --durable --x-message-ttl 5000 test_queue
```

Create durable topic exchange named test_exchange
```
rabbitmqmgmt exchange_add --type topic --durable test_exchange
```


Bind the test_queue to the test_exchange for all the topics starting with 'error'
```
rabbitmqmgmt queue_bind test_queue test_exchange 'error.#'
```

Unbind the previus binding
```
rabbitmqmgmt queue_unbind test_queue test_exchange 'error.#'
```

Remove the test_queue
```
rabbitmqmgmt queue_remove test_queue
```

Remove the test_exchange
```
rabbitmqmgmt exchange_remove test_exchange
```

