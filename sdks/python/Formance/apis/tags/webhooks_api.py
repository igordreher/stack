# coding: utf-8

"""
    Formance Stack API

    Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions />   # noqa: E501

    The version of the OpenAPI document: develop
    Contact: support@formance.com
    Generated by: https://openapi-generator.tech
"""

from Formance.paths.api_webhooks_configs_id_activate.put import ActivateConfig
from Formance.paths.api_webhooks_configs_id_secret_change.put import ChangeConfigSecret
from Formance.paths.api_webhooks_configs_id_deactivate.put import DeactivateConfig
from Formance.paths.api_webhooks_configs_id.delete import DeleteConfig
from Formance.paths.api_webhooks_configs.get import GetManyConfigs
from Formance.paths.api_webhooks_configs.post import InsertConfig
from Formance.paths.api_webhooks_configs_id_test.get import TestConfig


class WebhooksApi(
    ActivateConfig,
    ChangeConfigSecret,
    DeactivateConfig,
    DeleteConfig,
    GetManyConfigs,
    InsertConfig,
    TestConfig,
):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """
    pass
