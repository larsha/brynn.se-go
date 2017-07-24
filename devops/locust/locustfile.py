from locust import HttpLocust, TaskSet, task

import random

urls = [l.strip() for l in open('accesslog.txt').readlines()]

class WebsiteTasks(TaskSet):

    @task
    def request(self):
        url = random.choice(urls)
        self.client.get(url, verify=False)

class WebsiteUser(HttpLocust):
    min_wait = 2000
    max_wait = 5000
    task_set = WebsiteTasks
