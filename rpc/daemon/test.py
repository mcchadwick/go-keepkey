from messages_pb2_twirp import DaemonClient
import messages_pb2


print("hello")
req = messages_pb2.PingRequest()
params = messages_pb2.PingParams()
params.msg = "python"
req.params = params

client = DaemonClient("http://127.0.0.1:8080")
client.ping(req)
client.ping()
