// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: protoidentity/identity.response.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Mobius.Proto.Identity {

  /// <summary>Holder for reflection information generated from protoidentity/identity.response.proto</summary>
  public static partial class IdentityResponseReflection {

    #region Descriptor
    /// <summary>File descriptor for protoidentity/identity.response.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static IdentityResponseReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiVwcm90b2lkZW50aXR5L2lkZW50aXR5LnJlc3BvbnNlLnByb3RvEghpZGVu",
            "dGl0eRoacHJvdG9pZGVudGl0eS9wbGF5ZXIucHJvdG8iiQEKH0dldFNpbmds",
            "ZVBsYXllcklkZW50aXR5UmVzcG9uc2USLwoIaWRlbnRpdHkYASABKAsyHS5p",
            "ZGVudGl0eS5Qcm90b1BsYXllcklkZW50aXR5EjUKB2xvYWRvdXQYAiABKAsy",
            "JC5pZGVudGl0eS5Qcm90b1BsYXllckxvYWRvdXRJZGVudGl0eSKAAQoeR2V0",
            "U2luZ2xlUGxheWVyTG9hZG91dFJlc3BvbnNlEi8KCGlkZW50aXR5GAEgASgL",
            "Mh0uaWRlbnRpdHkuUHJvdG9QbGF5ZXJJZGVudGl0eRItCgdsb2Fkb3V0GAIg",
            "ASgLMhwuaWRlbnRpdHkuUHJvdG9QbGF5ZXJMb2Fkb3V0In8KHUdldE11bHRp",
            "UGxheWVyTG9hZG91dFJlc3BvbnNlEi8KCGlkZW50aXR5GAEgASgLMh0uaWRl",
            "bnRpdHkuUHJvdG9QbGF5ZXJJZGVudGl0eRItCgdsb2Fkb3V0GAIgASgLMhwu",
            "aWRlbnRpdHkuUHJvdG9QbGF5ZXJMb2Fkb3V0QlNaOWdpdGh1Yi5jb20vanVz",
            "dGphY2sxNTIxL21ldml1bS9wa2cvZ2VucHJvdG8vcHJvdG9pZGVudGl0eaoC",
            "FU1vYml1cy5Qcm90by5JZGVudGl0eWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Mobius.Proto.Identity.PlayerReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Mobius.Proto.Identity.GetSinglePlayerIdentityResponse), global::Mobius.Proto.Identity.GetSinglePlayerIdentityResponse.Parser, new[]{ "Identity", "Loadout" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Mobius.Proto.Identity.GetSinglePlayerLoadoutResponse), global::Mobius.Proto.Identity.GetSinglePlayerLoadoutResponse.Parser, new[]{ "Identity", "Loadout" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Mobius.Proto.Identity.GetMultiPlayerLoadoutResponse), global::Mobius.Proto.Identity.GetMultiPlayerLoadoutResponse.Parser, new[]{ "Identity", "Loadout" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class GetSinglePlayerIdentityResponse : pb::IMessage<GetSinglePlayerIdentityResponse>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<GetSinglePlayerIdentityResponse> _parser = new pb::MessageParser<GetSinglePlayerIdentityResponse>(() => new GetSinglePlayerIdentityResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<GetSinglePlayerIdentityResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Mobius.Proto.Identity.IdentityResponseReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetSinglePlayerIdentityResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetSinglePlayerIdentityResponse(GetSinglePlayerIdentityResponse other) : this() {
      identity_ = other.identity_ != null ? other.identity_.Clone() : null;
      loadout_ = other.loadout_ != null ? other.loadout_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetSinglePlayerIdentityResponse Clone() {
      return new GetSinglePlayerIdentityResponse(this);
    }

    /// <summary>Field number for the "identity" field.</summary>
    public const int IdentityFieldNumber = 1;
    private global::Mobius.Proto.Identity.ProtoPlayerIdentity identity_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Mobius.Proto.Identity.ProtoPlayerIdentity Identity {
      get { return identity_; }
      set {
        identity_ = value;
      }
    }

    /// <summary>Field number for the "loadout" field.</summary>
    public const int LoadoutFieldNumber = 2;
    private global::Mobius.Proto.Identity.ProtoPlayerLoadoutIdentity loadout_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Mobius.Proto.Identity.ProtoPlayerLoadoutIdentity Loadout {
      get { return loadout_; }
      set {
        loadout_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as GetSinglePlayerIdentityResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(GetSinglePlayerIdentityResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(Identity, other.Identity)) return false;
      if (!object.Equals(Loadout, other.Loadout)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (identity_ != null) hash ^= Identity.GetHashCode();
      if (loadout_ != null) hash ^= Loadout.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (identity_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Identity);
      }
      if (loadout_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(Loadout);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (identity_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Identity);
      }
      if (loadout_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(Loadout);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (identity_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Identity);
      }
      if (loadout_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Loadout);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(GetSinglePlayerIdentityResponse other) {
      if (other == null) {
        return;
      }
      if (other.identity_ != null) {
        if (identity_ == null) {
          Identity = new global::Mobius.Proto.Identity.ProtoPlayerIdentity();
        }
        Identity.MergeFrom(other.Identity);
      }
      if (other.loadout_ != null) {
        if (loadout_ == null) {
          Loadout = new global::Mobius.Proto.Identity.ProtoPlayerLoadoutIdentity();
        }
        Loadout.MergeFrom(other.Loadout);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (identity_ == null) {
              Identity = new global::Mobius.Proto.Identity.ProtoPlayerIdentity();
            }
            input.ReadMessage(Identity);
            break;
          }
          case 18: {
            if (loadout_ == null) {
              Loadout = new global::Mobius.Proto.Identity.ProtoPlayerLoadoutIdentity();
            }
            input.ReadMessage(Loadout);
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            if (identity_ == null) {
              Identity = new global::Mobius.Proto.Identity.ProtoPlayerIdentity();
            }
            input.ReadMessage(Identity);
            break;
          }
          case 18: {
            if (loadout_ == null) {
              Loadout = new global::Mobius.Proto.Identity.ProtoPlayerLoadoutIdentity();
            }
            input.ReadMessage(Loadout);
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class GetSinglePlayerLoadoutResponse : pb::IMessage<GetSinglePlayerLoadoutResponse>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<GetSinglePlayerLoadoutResponse> _parser = new pb::MessageParser<GetSinglePlayerLoadoutResponse>(() => new GetSinglePlayerLoadoutResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<GetSinglePlayerLoadoutResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Mobius.Proto.Identity.IdentityResponseReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetSinglePlayerLoadoutResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetSinglePlayerLoadoutResponse(GetSinglePlayerLoadoutResponse other) : this() {
      identity_ = other.identity_ != null ? other.identity_.Clone() : null;
      loadout_ = other.loadout_ != null ? other.loadout_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetSinglePlayerLoadoutResponse Clone() {
      return new GetSinglePlayerLoadoutResponse(this);
    }

    /// <summary>Field number for the "identity" field.</summary>
    public const int IdentityFieldNumber = 1;
    private global::Mobius.Proto.Identity.ProtoPlayerIdentity identity_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Mobius.Proto.Identity.ProtoPlayerIdentity Identity {
      get { return identity_; }
      set {
        identity_ = value;
      }
    }

    /// <summary>Field number for the "loadout" field.</summary>
    public const int LoadoutFieldNumber = 2;
    private global::Mobius.Proto.Identity.ProtoPlayerLoadout loadout_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Mobius.Proto.Identity.ProtoPlayerLoadout Loadout {
      get { return loadout_; }
      set {
        loadout_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as GetSinglePlayerLoadoutResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(GetSinglePlayerLoadoutResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(Identity, other.Identity)) return false;
      if (!object.Equals(Loadout, other.Loadout)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (identity_ != null) hash ^= Identity.GetHashCode();
      if (loadout_ != null) hash ^= Loadout.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (identity_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Identity);
      }
      if (loadout_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(Loadout);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (identity_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Identity);
      }
      if (loadout_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(Loadout);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (identity_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Identity);
      }
      if (loadout_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Loadout);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(GetSinglePlayerLoadoutResponse other) {
      if (other == null) {
        return;
      }
      if (other.identity_ != null) {
        if (identity_ == null) {
          Identity = new global::Mobius.Proto.Identity.ProtoPlayerIdentity();
        }
        Identity.MergeFrom(other.Identity);
      }
      if (other.loadout_ != null) {
        if (loadout_ == null) {
          Loadout = new global::Mobius.Proto.Identity.ProtoPlayerLoadout();
        }
        Loadout.MergeFrom(other.Loadout);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (identity_ == null) {
              Identity = new global::Mobius.Proto.Identity.ProtoPlayerIdentity();
            }
            input.ReadMessage(Identity);
            break;
          }
          case 18: {
            if (loadout_ == null) {
              Loadout = new global::Mobius.Proto.Identity.ProtoPlayerLoadout();
            }
            input.ReadMessage(Loadout);
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            if (identity_ == null) {
              Identity = new global::Mobius.Proto.Identity.ProtoPlayerIdentity();
            }
            input.ReadMessage(Identity);
            break;
          }
          case 18: {
            if (loadout_ == null) {
              Loadout = new global::Mobius.Proto.Identity.ProtoPlayerLoadout();
            }
            input.ReadMessage(Loadout);
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class GetMultiPlayerLoadoutResponse : pb::IMessage<GetMultiPlayerLoadoutResponse>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<GetMultiPlayerLoadoutResponse> _parser = new pb::MessageParser<GetMultiPlayerLoadoutResponse>(() => new GetMultiPlayerLoadoutResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<GetMultiPlayerLoadoutResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Mobius.Proto.Identity.IdentityResponseReflection.Descriptor.MessageTypes[2]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetMultiPlayerLoadoutResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetMultiPlayerLoadoutResponse(GetMultiPlayerLoadoutResponse other) : this() {
      identity_ = other.identity_ != null ? other.identity_.Clone() : null;
      loadout_ = other.loadout_ != null ? other.loadout_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public GetMultiPlayerLoadoutResponse Clone() {
      return new GetMultiPlayerLoadoutResponse(this);
    }

    /// <summary>Field number for the "identity" field.</summary>
    public const int IdentityFieldNumber = 1;
    private global::Mobius.Proto.Identity.ProtoPlayerIdentity identity_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Mobius.Proto.Identity.ProtoPlayerIdentity Identity {
      get { return identity_; }
      set {
        identity_ = value;
      }
    }

    /// <summary>Field number for the "loadout" field.</summary>
    public const int LoadoutFieldNumber = 2;
    private global::Mobius.Proto.Identity.ProtoPlayerLoadout loadout_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Mobius.Proto.Identity.ProtoPlayerLoadout Loadout {
      get { return loadout_; }
      set {
        loadout_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as GetMultiPlayerLoadoutResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(GetMultiPlayerLoadoutResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(Identity, other.Identity)) return false;
      if (!object.Equals(Loadout, other.Loadout)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (identity_ != null) hash ^= Identity.GetHashCode();
      if (loadout_ != null) hash ^= Loadout.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (identity_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Identity);
      }
      if (loadout_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(Loadout);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (identity_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Identity);
      }
      if (loadout_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(Loadout);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (identity_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Identity);
      }
      if (loadout_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Loadout);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(GetMultiPlayerLoadoutResponse other) {
      if (other == null) {
        return;
      }
      if (other.identity_ != null) {
        if (identity_ == null) {
          Identity = new global::Mobius.Proto.Identity.ProtoPlayerIdentity();
        }
        Identity.MergeFrom(other.Identity);
      }
      if (other.loadout_ != null) {
        if (loadout_ == null) {
          Loadout = new global::Mobius.Proto.Identity.ProtoPlayerLoadout();
        }
        Loadout.MergeFrom(other.Loadout);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (identity_ == null) {
              Identity = new global::Mobius.Proto.Identity.ProtoPlayerIdentity();
            }
            input.ReadMessage(Identity);
            break;
          }
          case 18: {
            if (loadout_ == null) {
              Loadout = new global::Mobius.Proto.Identity.ProtoPlayerLoadout();
            }
            input.ReadMessage(Loadout);
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            if (identity_ == null) {
              Identity = new global::Mobius.Proto.Identity.ProtoPlayerIdentity();
            }
            input.ReadMessage(Identity);
            break;
          }
          case 18: {
            if (loadout_ == null) {
              Loadout = new global::Mobius.Proto.Identity.ProtoPlayerLoadout();
            }
            input.ReadMessage(Loadout);
            break;
          }
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code
