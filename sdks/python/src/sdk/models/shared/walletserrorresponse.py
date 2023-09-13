"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
from dataclasses_json import Undefined, dataclass_json
from enum import Enum
from sdk import utils

class WalletsErrorResponseErrorCode(str, Enum):
    VALIDATION = 'VALIDATION'
    INTERNAL_ERROR = 'INTERNAL_ERROR'
    INSUFFICIENT_FUND = 'INSUFFICIENT_FUND'
    HOLD_CLOSED = 'HOLD_CLOSED'


@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class WalletsErrorResponse:
    r"""Error"""
    error_code: WalletsErrorResponseErrorCode = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('errorCode') }})
    error_message: str = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('errorMessage') }})
    

