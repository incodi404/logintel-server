from unittest import result

from fastapi import FastAPI, Request
from fastapi import UploadFile, File
from pydantic import BaseModel
from typing import Any, List
import json
import requests
import logging
from database import collection
import time
import sys
from database import save_log
import csv
import io

logging.basicConfig(
                    level=logging.INFO, 
                    format="%(asctime)s - %(levelname)s - %(message)s",

                    handlers=[logging.StreamHandler(sys.stdout)]
                    )

logging.info("Application started")
app = FastAPI()

OLLAMA_URL = "http://ollama:11434/api/generate"

class AnalysisResponse(BaseModel):
    title: str
    summary: List[str]
    suspicious_activity: List[str]
    errors_or_anomalies: List[str]
    severity: List[str]
    recommended_actions: List[str]
    raw_response: Any

@app.get("/")
def home():
    logging.info("Root endpoint accessed")
    return {"message": "Ollama FastAPI Server Running"}


@app.middleware("http")
async def log_requests(request:Request, call_next):

    start_time = time.time()

    response = await call_next(request) 
    
    duration = time.time() - start_time
    
    logging.info(
        f"{request.method} {request.url.path}"
        f"Status:{response.status_code}"
          f"Time:{duration:.3f}s"
    )

    return response


@app.get("/history")
async def get_log_history():
    logs = list(collection.find({}, {"_id": 0}))
                    
    return {"logs": logs}


@app.post("/analyze-csv-file", response_model=AnalysisResponse)
async def analyze_csv_file(file: UploadFile = File(...)):

    contents = await file.read()

    csv_text = contents.decode("utf-8")

    reader = csv.DictReader(io.StringIO(csv_text))

    logs = list(reader)[:20]

    print("Number of rows:", len(logs))

    payload = {
        "model": "llama3.2:latest",
        "prompt": f"""
         You are a detection engineer and SOC analyst.
            
        Analyze the provided CSV log data and identify security incidents, suspicious activity, 
        system errors, performance issues, and operational anomalies.

        Return ONLY valid JSON.

        Requirements:
        1. Return exactly one JSON object.
        2. Do not include markdown, explanations, or extra text.
        3. Use this exact schema and do not rename or omit any key.
        4. `title` must be a string summarizing the overall analysis.
        5. All other fields must be arrays containing only plain text strings.
        6. Do not return CSV rows, raw log entries, timestamps, objects, dictionaries, or nested arrays.
        7. Use `[]` when no findings exist.
        8. `summary` should contain key observations.
        9. `suspicious_activity` should contain only suspicious or malicious behavior.
        10. `errors_or_anomalies` should contain errors, failures, or unusual events.
        11. `severity` must contain exactly one value: `"Low"`, `"Medium"`, `"High"`, or `"Critical"`.
        12. `recommended_actions` should contain clear recommendations based only on the findings.
        13. Do not invent information that is not supported by the CSV logs.
        14. Limit each array to a maximum of 5 items.
        
        Return exactly this JSON schema:

    {{
        "title": "",
        "summary": [],
        "suspicious_activity": [],
        "errors_or_anomalies": [],
        "severity": [],
        "recommended_actions": []
    }}
        
        Logs:
        {json.dumps(logs, indent=2)}""",
        "format": "json",
        "stream": False
    }

    response = requests.post(OLLAMA_URL, json=payload, timeout=300)

    response.raise_for_status()

    result = response.json()

    print("=" * 80)
    print("OLLAMA RESULT:")
    print(json.dumps(result, indent=2))
    print("=" * 80)

    raw_llm_response = result.get("response", "")

    required_fields = [
        "title",
        "summary",
        "suspicious_activity",
        "errors_or_anomalies",
        "severity",
        "recommended_actions",
    ]

    try:
        analysis = json.loads(raw_llm_response)

        if not isinstance(analysis, dict):
            raise ValueError("LLM response is not a JSON object.")

        missing = [f for f in required_fields if f not in analysis]

        if missing:
            print("Missing fields:", missing)

            return AnalysisResponse(
                title="Invalid LLM Response",
                summary=["The LLM did not return the expected schema."],
                suspicious_activity=[],
                errors_or_anomalies=[],
                severity=["Unknown"],
                recommended_actions=["Review the raw LLM response."],
                raw_response={
                    "llm_response": raw_llm_response,
                    "parsed_response": analysis,
                    "missing_fields": missing,
                    "input_logs": logs,
                },
            )

        analysis["raw_response"] = logs

        save_log(
            title=analysis.get("title", ""),
            log_contents=logs,
            analysis=analysis,
        )

        return AnalysisResponse(**analysis)

    except Exception as e:
        print("LLM parsing error:", e)

        return AnalysisResponse(
            title="LLM Parsing Failed",
            summary=["Unable to parse the LLM response."],
            suspicious_activity=[],
            errors_or_anomalies=[],
            severity=["Unknown"],
            recommended_actions=["Check the raw LLM response."],
            raw_response={
                "llm_response": raw_llm_response,
                "error": str(e),
                "input_logs": logs,
            },
        )

@app.post("/analyze-json", response_model=AnalysisResponse)
async def analyze_json(request: Request):

    logs = await request.json()

    payload = {
        "model": "llama3.2:latest",
        "prompt": f"""
     
        You are a detection engineer and master in SOC analysis. 

        Your task is to analyze the provided log data and identify security incidents, suspicious behavior, 
        system errors, performance issues, and operational anomalies.

        Return ONLY valid JSON.

        Requirements:

        1. Return only a single valid JSON object.
        2. Do not include markdown, explanations, comments, or extra text.
        3. Include all schema keys exactly as provided; do not rename or omit any key.
        4. `title` must be a string.
        5. All other fields must be arrays of plain text strings only.
        6. Do not return nested arrays, JSON objects, log entries, timestamps, or dictionaries inside any array.
        7. Use an empty array (`[]`) when no findings exist.
        8. `title` should briefly summarize the overall analysis.
        9. `summary` should contain concise observations.
        10. `suspicious_activity` should list only suspicious or malicious behavior.
        11. `errors_or_anomalies` should list errors, failures, or unusual events.
        12. `severity` must contain exactly one of: `"Low"`, `"Medium"`, `"High"`, or `"Critical"`.
        13. `recommended_actions` should contain clear, actionable recommendations.
        14. Do not repeat the input logs in the output.
        15. Include no more than 5 items in each array.


        
        Return the JSON using exactly this schema:

        {{
    
            "title": "",
            "summary": [],
            "suspicious_activity": [],
            "errors_or_anomalies": [],
            "severity": [],
            "recommended_actions": []

        }}


        Logs:
        {json.dumps(logs, indent=2)}""",
        "format": "json",
        "stream": False
    }

    response = requests.post(OLLAMA_URL, json=payload)

    response.raise_for_status()

    result = response.json()

    print("LLM Response:")
    print(result["response"])

    try:
        analysis = json.loads(result["response"])
        analysis["raw_response"] = logs

        save_log(

            title=analysis.get("title", ""),
            log_contents=logs,
            analysis=analysis
        )


    except json.JSONDecodeError:
         return {
            "title": "",
            "summary": ["LLM returned invalid JSON."],
            "suspicious_activity": [""],
            "errors_or_anomalies": [""],
            "severity": [""],
            "recommended_actions": ["Retry analysis."],
            "raw_response": {}
        }

    return analysis


@app.post("/analyze-log-file", response_model=AnalysisResponse)
async def analyse_logs_file(file: UploadFile = File(...)):
    
    contents = await file.read()
    logs = json.loads(contents.decode("utf-8"))


    payload = {
        "model": "llama3.2:latest",
        "prompt": f"""
        You are a detection engineer and master in SOC analysis.

        Analyze the given log and based on the log, give me informations.

        Return ONLY valid JSON in the following format:

        {{
            "title": "",
            "summary": ["......"],
            "suspicious_activity": ["....."],
            "errors_or_anomalies": ["....."],
            "severity": ["...."],
            "recommended_actions": ["...."],
        }}


        Logs:
        {json.dumps(logs, indent=2)}""",
        "format": "json",
        "stream": False
    }
    
    response = requests.post(OLLAMA_URL, json=payload)

    response.raise_for_status()
         
    result = response.json()

    print(result["response"])

    try:
      analysis = json.loads(result["response"])
      analysis["raw_response"] = logs

      save_log(
          
        title=analysis.get("title", ""),
        log_contents=logs,
        analysis=analysis
    )
    except json.JSONDecodeError:
     return {
         "title": "",
        "summary": ["LLM returned invalid JSON."],
        "suspicious_activity": [""],
        "errors_or_anomalies": [""],
        "severity": ["Unknown"],
        "recommended_actions": ["Retry analysis."],
        "raw_response": {}
    }

    return AnalysisResponse(**analysis)
    
   
   