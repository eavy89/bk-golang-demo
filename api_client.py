#!/usr/bin/env python3

import requests

BASE_URL = "http://127.0.0.1:8080"
HEADERS = {"Content-Type": "application/json"}

def register_user(username, password):
    res = requests.post(
        f"{BASE_URL}/register",
        json={"username": username, "password": password},
        headers=HEADERS
    )
    print("Register: Status code -", res.status_code, res.json())

def login_user(username, password):
    res = requests.post(
        f"{BASE_URL}/login",
        json={"username": username, "password": password},
        headers=HEADERS
    )
    print("Login: Status code - ", res.status_code, res.json())
    if res.status_code == 200 and "token" in res.json():
        return res.json()["token"]
    return None

def buy_product(token, item, quantity, price):
    headers = {
        **HEADERS,
        "Authorization": f"Bearer {token}"
    }
    res = requests.post(
        f"{BASE_URL}/api/purchase",
        json={"item": item, "quantity": quantity, "price": price},
        headers=headers
    )
    print("Buy: Status code -", res.status_code, res.json())

def get_history(token):
    headers = {
        **HEADERS,
        "Authorization": f"Bearer {token}"
    }
    res = requests.get(
        f"{BASE_URL}/api/purchases",
        headers=headers
    )
    print("History: Status code -", res.status_code)
    for purchase in res.json():
        print(purchase)

## ======= Client Workflow ======= ##

def client_workflow(username, password, need_register=False):
    if need_register:
        register_user(username, password)

    token = login_user(username, password)
    if token:
        buy_product(token, "item1", 1, 10.99)
        buy_product(token, "item2", 2, 20)
        get_history(token)
    else:
        print("Login failed")



client_workflow("ivan2", "1234", True)