from concurrent import futures
import time
import logging

import grpc

import streamcounters.counters.counters_pb2 as cp
import streamcounters.counters.counters_pb2_grpc as cpg

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

class StatusServicer(cpg.StatusServicer):
    """Provides methods that implement functionality of counters server"""

    def GetCounters(self, request, context):
        if request.client_id == '':
            raise ValueError('invalid client ID')

        resp = cp.CounterResp(
            ok=True
        )

        for count in range(5):
            resp.counter = count+1
            time.sleep(3)
            yield resp

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    cpg.add_StatusServicer_to_server(
        StatusServicer(), server
    )
    server.add_insecure_port('[::]:50052')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except:
        server.stop(0)


if __name__ == "__main__":
    logging.basicConfig()
    serve()