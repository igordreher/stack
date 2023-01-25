/*
 * Formance Stack API
 * Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 
 *
 * The version of the OpenAPI document: develop
 * Contact: support@formance.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.formance.formance.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.annotations.SerializedName;

import java.io.IOException;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;

/**
 * Gets or Sets Connectors
 */
@JsonAdapter(Connectors.Adapter.class)
public enum Connectors {
  
  STRIPE("STRIPE"),
  
  DUMMY_PAY("DUMMY-PAY"),
  
  SIE("SIE"),
  
  MODULR("MODULR"),
  
  CURRENCY_CLOUD("CURRENCY-CLOUD"),
  
  BANKING_CIRCLE("BANKING-CIRCLE");

  private String value;

  Connectors(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static Connectors fromValue(String value) {
    for (Connectors b : Connectors.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<Connectors> {
    @Override
    public void write(final JsonWriter jsonWriter, final Connectors enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public Connectors read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return Connectors.fromValue(value);
    }
  }
}

