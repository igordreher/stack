"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
from ..shared import workflowinstancehistory as shared_workflowinstancehistory
from dataclasses_json import Undefined, dataclass_json
from sdk import utils


@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class GetWorkflowInstanceHistoryResponse:
    r"""The workflow instance history"""
    data: list[shared_workflowinstancehistory.WorkflowInstanceHistory] = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('data') }})
    

