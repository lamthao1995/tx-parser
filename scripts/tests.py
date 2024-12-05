import requests
import json

# Base URL of the API
BASE_URL = "http://localhost:8080"


# Helper function to test an API
def test_api(endpoint, method="GET", params=None, json_data=None, expected_status=200):
    print(f"Testing {method} {BASE_URL + endpoint}")

    try:
        if method == "GET":
            response = requests.get(BASE_URL + endpoint, params=params)
        elif method == "POST":
            response = requests.post(BASE_URL + endpoint, json=json_data)
        else:
            print(f"Unsupported method: {method}")
            return False

        # Validate response
        if response.status_code != expected_status:
            print(f"❌ Test failed: {response.status_code} != {expected_status}")
            print(f"Response: {response.text}")
            return False

        print(f"✅ Test passed: {response.status_code} == {expected_status}")
        print(f"Response: {response.json()}")
        return True

    except requests.RequestException as e:
        print(f"❌ Test failed with exception: {e}")
        return False


# Test Cases
def run_tests():
    # 1. Test /currentBlock API
    test_api("/currentBlock")

    # 2. Test /subscribe API
    test_api("/subscribe", params={"address": "0x789ghi"})

    # 3. Test /transactions API
    test_api("/transactions", params={"address": "0x789ghi"})


if __name__ == "__main__":
    run_tests()
