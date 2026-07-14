import requests


def main():
    with open("app.log", "r") as file:
        logs = file.read()

        prompt = f"""Analyze the following logs.
        Identify:
        1. Errors
        2. Root causes
        3. Severity levels
        4. Recommendations for resolution
        Logs:
        {logs}"""

    response = requests.post(
        "http://localhost:8000/analyze-logs",
        json={"logs": logs},
    )

    print(response.json())


