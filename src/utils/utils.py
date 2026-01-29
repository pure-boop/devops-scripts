
def process_data(data: dict) -> dict:
    '''Process incoming data payload'''
    cleaned = {k: v.strip() if isinstance(v, str) else v for k, v in data.items()}
    return cleaned

