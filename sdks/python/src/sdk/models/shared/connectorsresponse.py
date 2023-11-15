"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
from ..shared import connector as shared_connector
from dataclasses_json import Undefined, dataclass_json
from sdk import utils


@dataclass_json(undefined=Undefined.EXCLUDE)
@dataclasses.dataclass
class ConnectorsResponseData:
    
    connector_id: str = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('connectorID') }})
    name: str = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('name') }})
    provider: shared_connector.Connector = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('provider') }})
    

@dataclass_json(undefined=Undefined.EXCLUDE)
@dataclasses.dataclass
class ConnectorsResponse:
    r"""OK"""
    
    data: list[ConnectorsResponseData] = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('data') }})
    