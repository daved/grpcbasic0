#!/usr/bin/env python2.7

from __future__ import print_function

import grpc
import grpcbasic0_pb2
import grpcbasic0_pb2_grpc

if __name__ == '__main__':
    tmpl = "Iteration {} - ID: {}, Name: {}, Fortune: {}"
    chan = grpc.insecure_channel('localhost:3323')
    stub = grpcbasic0_pb2_grpc.UserServiceStub(chan)
    resp = stub.GetUsers(grpcbasic0_pb2.GetUsersReq(start=2, count=2))
    ct = 0
    for u in resp.users:
        ct += 1
        print(tmpl.format(ct, u.id, u.name, u.fortune))
