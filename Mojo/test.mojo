# import time
from python import Python


# def count_to_one_billion():

#     for i in range(1, 100000000): 
#         pass 
#    #try:
#         # let time = Python.import_module("time")
#         # start_time = time.time()

#         # end_time = time.time()
#         # duration = end_time - start_time
#     # except:
#     #     print("What")
#     #     pass
#     # finally: 
#     #     print("What")

fn main():
    try:
        let time = Python.import_module("time")
        let start_time = time.time()
        for i in range(1, 1000000000):  # Count from 1 to 1,000,000,000
            pass  # Just a placeholder to indicate a loop iteration, does nothing
        let end_time = time.time()
        let duration = end_time - start_time
        print("Counting to 1 trillion took", duration, "seconds.")
    except: 
        print("Error")

