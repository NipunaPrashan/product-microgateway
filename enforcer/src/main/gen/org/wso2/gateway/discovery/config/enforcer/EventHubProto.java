// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/config/enforcer/event_hub.proto

package org.wso2.gateway.discovery.config.enforcer;

public final class EventHubProto {
  private EventHubProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_wso2_discovery_config_enforcer_EventHub_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_wso2_discovery_config_enforcer_EventHub_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_wso2_discovery_config_enforcer_EventHub_JmsConnectionParametersEntry_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_wso2_discovery_config_enforcer_EventHub_JmsConnectionParametersEntry_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n.wso2/discovery/config/enforcer/event_h" +
      "ub.proto\022\036wso2.discovery.config.enforcer" +
      "\"\327\001\n\010EventHub\022\017\n\007enabled\030\001 \001(\010\022\022\n\nservic" +
      "eUrl\030\002 \001(\t\022f\n\027jmsConnectionParameters\030\003 " +
      "\003(\0132E.wso2.discovery.config.enforcer.Eve" +
      "ntHub.JmsConnectionParametersEntry\032>\n\034Jm" +
      "sConnectionParametersEntry\022\013\n\003key\030\001 \001(\t\022" +
      "\r\n\005value\030\002 \001(\t:\0028\001B\215\001\n*org.wso2.gateway." +
      "discovery.config.enforcerB\rEventHubProto" +
      "P\001ZNgithub.com/envoyproxy/go-control-pla" +
      "ne/wso2/discovery/config/enforcer;enforc" +
      "erb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_wso2_discovery_config_enforcer_EventHub_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_wso2_discovery_config_enforcer_EventHub_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_wso2_discovery_config_enforcer_EventHub_descriptor,
        new java.lang.String[] { "Enabled", "ServiceUrl", "JmsConnectionParameters", });
    internal_static_wso2_discovery_config_enforcer_EventHub_JmsConnectionParametersEntry_descriptor =
      internal_static_wso2_discovery_config_enforcer_EventHub_descriptor.getNestedTypes().get(0);
    internal_static_wso2_discovery_config_enforcer_EventHub_JmsConnectionParametersEntry_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_wso2_discovery_config_enforcer_EventHub_JmsConnectionParametersEntry_descriptor,
        new java.lang.String[] { "Key", "Value", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}