/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class ResetConnectorV1Request {
    /**
     * The name of the connector.
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=connector")
    public com.formance.formance_sdk.models.shared.Connector connector;

    public ResetConnectorV1Request withConnector(com.formance.formance_sdk.models.shared.Connector connector) {
        this.connector = connector;
        return this;
    }
    
    /**
     * The connector ID.
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=connectorId")
    public String connectorId;

    public ResetConnectorV1Request withConnectorId(String connectorId) {
        this.connectorId = connectorId;
        return this;
    }
    
    public ResetConnectorV1Request(@JsonProperty("connector") com.formance.formance_sdk.models.shared.Connector connector, @JsonProperty("connectorId") String connectorId) {
        this.connector = connector;
        this.connectorId = connectorId;
  }
}
