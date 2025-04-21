import base64
import json
import os
import io
import matplotlib.pyplot as plt
import pika
from math import log,exp
from pika.adapters.blocking_connection import BlockingChannel
from pika.spec import Basic

def funcs(approx , k):
    if approx == 'Linear':
        return lambda x: k[0] + x*k[1]
    if approx == 'Square':
        return lambda x: k[0] + x*k[1] + x**2*k[2]
    if approx == 'Cube':
        return lambda x: k[0] + x*k[1] + x**2*k[2] + x**3*k[3]
    if approx == 'Log':
        return lambda x: k[0] + log(x)*k[1]
    if approx == 'Exp':
        return lambda x: exp(k[1]*x) * exp(k[0])
    if approx == 'Pow':
        return lambda x:  x**k[1] * exp(k[0])


def get_interval(xs):
    sxs = sorted(xs)
    a = -abs(sxs[0]) * 0.2 + sxs[0]
    b = abs(sxs[-1]) * 0.2 + sxs[-1]
    width = (b-a)/1000
    return [a + i*width for i in range(1000)]

global from_graph
def generate_graph(ch: BlockingChannel, method: Basic.Deliver, properties: pika.BasicProperties, body : str):
    stringIObytes = io.BytesIO()
    inp = json.loads(body)
    plt.plot(inp['xs'], inp['ys'], 'o')
    interval = get_interval(inp['xs'])
    func = funcs(inp['approx'], inp['ks'])
    plt.plot(interval, list(map(func, interval)))
    #plt.show()
    plt.savefig(stringIObytes, format='jpg')
    stringIObytes.seek(0)
    base64_jpgData = base64.b64encode(stringIObytes.read()).decode()
    ch.basic_publish(exchange="",
            routing_key=from_graph,
            body=json.dumps({"graph": base64_jpgData}),
            properties=properties
    )

if __name__ == "__main__":
    hostname = os.environ.get('RABBITMQ_HOST')
    port = os.environ.get('RABBITMQ_PORT')
    user = os.environ.get('RABBITMQ_USER')
    password = os.environ.get('RABBITMQ_PASSWORD')
    to_graph = os.environ.get('RABBITMQ_QUEUE_TO_GRAPH')
    from_graph = os.environ.get('RABBITMQ_QUEUE_FROM_GRAPH')
    connection = pika.BlockingConnection(pika.ConnectionParameters(host=hostname, port=port, credentials=pika.PlainCredentials(user, password)))
    ch = connection.channel()
    ch.queue_declare(queue=to_graph)
    ch.queue_declare(queue=from_graph)
    ch.basic_consume(
        queue=to_graph,
        on_message_callback=generate_graph,
        auto_ack=True
    )
    ch.start_consuming()

