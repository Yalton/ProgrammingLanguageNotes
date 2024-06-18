import time

def count_to_one_billion():
    start_time = time.time()
    number = 0
    for i in range(1, 1000000000):  # Count from 1 to 1000000000
        number = i
        pass  # Just a placeholder to indicate a loop iteration, does nothing
    end_time = time.time()
    duration = end_time - start_time
    print(f"Counting to 1 10000000000 took {duration} seconds.")

if __name__ == "__main__":
    count_to_one_billion()
