// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/config/enforcer/keystore.proto

package org.wso2.gateway.discovery.config.enforcer;

public interface KeystoreOrBuilder extends
    // @@protoc_insertion_point(interface_extends:wso2.discovery.config.enforcer.Keystore)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * Path to the private-key of the keypair
   * </pre>
   *
   * <code>string privateKeyLocation = 1;</code>
   * @return The privateKeyLocation.
   */
  java.lang.String getPrivateKeyLocation();
  /**
   * <pre>
   * Path to the private-key of the keypair
   * </pre>
   *
   * <code>string privateKeyLocation = 1;</code>
   * @return The bytes for privateKeyLocation.
   */
  com.google.protobuf.ByteString
      getPrivateKeyLocationBytes();

  /**
   * <pre>
   * Path to the public-key of the keypair
   * </pre>
   *
   * <code>string publicKeyLocation = 2;</code>
   * @return The publicKeyLocation.
   */
  java.lang.String getPublicKeyLocation();
  /**
   * <pre>
   * Path to the public-key of the keypair
   * </pre>
   *
   * <code>string publicKeyLocation = 2;</code>
   * @return The bytes for publicKeyLocation.
   */
  com.google.protobuf.ByteString
      getPublicKeyLocationBytes();
}
