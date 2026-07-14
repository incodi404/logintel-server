import logging
import datetime
from database import collection


class MongoDBHandler(logging.Handler):
    def emit(self, record):
        try:
            log_document = {
                "timestamp": datetime.datetime.utcnow(),
                "logger_name": record.name,
                "level": record.levelname,
                "message": self.format(record),
                "module": record.module,
                "line_number": record.lineno
            }

            collection.insert_one(log_document)

        except Exception:
            self.handleError(record)
            
        logger = logging.getLogger("api_logger")
        logger.setLevel(logging.INFO)

        formatter = logging.Formatter("%(asctime)s - %(levelname)s - %(message)s")

        mongo_handler = MongoDBHandler()
        mongo_handler.setFormatter(formatter)
        logger.addHandler(mongo_handler)

        file_handler = logging.FileHandler('api_logs.log')
        file_handler.setFormatter(formatter)
        logger.addHandler(file_handler)



        