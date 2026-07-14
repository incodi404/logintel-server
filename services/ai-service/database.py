from pymongo import MongoClient
from datetime import datetime
import logging

client = MongoClient("mongodb://admin:password123@mongodb:27017")
db = client["log_analysis_API"]
collection = db["logs"]

def save_log(title, log_contents, analysis):
    log_document = {
        "title": title,
        "uploaded_at": datetime.utcnow(),
        "log_contents": log_contents,
        "analysis": analysis
    }
    result=collection.insert_one(log_document)
    logging.info(f"Log successfully written to MongoDB for event: {title}")
    print(f"Log saved to MongoDB: {result.inserted_id}")


