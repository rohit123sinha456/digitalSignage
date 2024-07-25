import pika , sys, os

# Set the connection parameters to connect to rabbit-server1 on port 5672
# on the / virtual host using the username "guest" and password "guest"
credentials = pika.PlainCredentials('DSU1187e3cc', 'password')
parameters = pika.ConnectionParameters('91.108.110.157',
                                       5672,
                                       'DSUVHOST1187e3cc',
                                       credentials)


connection = pika.BlockingConnection(parameters)
channel = connection.channel()

channel.exchange_declare(exchange='PLExchange', exchange_type='direct',durable=True)

result = channel.queue_declare(queue='', exclusive=True)
queue_name = result.method.queue

channel.queue_bind(exchange='PLExchange', queue=queue_name,routing_key="")

print(' [*] Waiting for logs. To exit press CTRL+C')

def callback(ch, method, properties, body):
    print(f" [x] {body}")

channel.basic_consume(
    queue=queue_name, on_message_callback=callback, auto_ack=True)

channel.start_consuming()