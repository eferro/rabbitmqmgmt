rabbitmqmgmt
============

rabbitmqctl equivalent for create/delete/bind queues, exchanges, etc.

USAGE:
   rabbitmqmgmt [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   queue_add		add a new queue
   queue_remove		remove an existing queue
   queue_bind		bind a queue to a exchange using a ginven topic/routing key
   queue_unbind		remove an existing binding
   exchange_add		add a new exchange
   exchange_remove	remove an existing exchange
   help, h		Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --amqp_uri, -u 'amqp://guest:guest@localhost:5672/'	broker url (including vhost)
   --version, -v					print the version
   --help, -h						show help
   

