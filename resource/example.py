# -*- coding: utf-8 -*-

import os, sys
import json, time
from confluent_kafka.cimpl import Consumer, KafkaError
import functools
from content_common.content_grpc_stubs.termite import Termite

TermiteHost = ""
m = Termite(TermiteHost, 3)

def termite(func):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        r, r_info = m.start_work(work_id=kwargs["work_id"])
        if r != 1:
            logging.getLogger("task").info("Termite Rpc Error: %s [work_id: %s]", r_info, kwargs["work_id"])
            return
        try:
            func(*args, **kwargs)
        except Exception as e:
            if settings.DEBUG:
                traceback.print_exc()
            m.raise_work(kwargs["work_id"], traceback.format_exc())
    return wrapper


@termite
def t_video_tag_detect(**args, **kwargs):
    """ """
    vid = long(kwargs["cid"])
    # code
    m.finish_work(kwargs["work_id"])


def deamon():
    """ Termite Client """
    group = ""
    KAFKA_HOST = ""
    KAFKA_TOPIC=""
    print KAFKA_HOST
    c = Consumer({
        "bootstrap.servers": KAFKA_HOST,
        'group.id': group,
    })
    c.subscribe([KAFKA_TOPIC])

    running = True
    while running:
        msg = c.poll(1)
        if msg is None:
            continue
        if not msg.error():
            data = json.loads(msg.value())
            print("receive msg:", data)
            kwargs = {
                "work_id": data["Work_id"],
                "flow_id": data["Flow_id"],
                "cid": data["Cid"]
            }
            # 视频美女标签识别
            if data.get("Work", "") == "video_tag_detect":
                t_video_tag_detect(**kwargs)
            else:
                pass
        else:
            if msg.error().code() == KafkaError._PARTITION_EOF:
                print "Skip-Error Message-Topic: {} Partition: {} Offset: {}Error: {}".format(msg.topic(),
                                                                                              msg.partition(),
                                                                                              msg.offset(),
                                                                                              msg.error())
            else:
                print "Error Message: {}".format(msg.error())
            time.sleep(0.01)
    c.close()

if __name__ == "__main__":
    deamon()
