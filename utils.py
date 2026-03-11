
from dataclasses import dataclass
@dataclass
class UserSession:
    session_id: str
    user_id: int
    expires_at: int

