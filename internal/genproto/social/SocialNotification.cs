// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: protosocial/social.notification.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Mobius.Proto.Social {

  /// <summary>Holder for reflection information generated from protosocial/social.notification.proto</summary>
  public static partial class SocialNotificationReflection {

    #region Descriptor
    /// <summary>File descriptor for protosocial/social.notification.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static SocialNotificationReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiVwcm90b3NvY2lhbC9zb2NpYWwubm90aWZpY2F0aW9uLnByb3RvEghwcmVz",
            "ZW5jZRoYcHJvdG9zb2NpYWwvc29jaWFsLnByb3RvIsIBChZTb2NpYWxEYXRh",
            "Tm90aWZpY2F0aW9uEjcKDmZvbGxvd2luZ19saXN0GAEgAygLMh8ucHJlc2Vu",
            "Y2UuUHJvdG9QbGF5ZXJTb2NpYWxJbmZvEjYKDWZvbGxvd2VyX2xpc3QYAiAD",
            "KAsyHy5wcmVzZW5jZS5Qcm90b1BsYXllclNvY2lhbEluZm8SNwoOcmVudGFs",
            "X3BsYXllcnMYAyADKAsyHy5wcmVzZW5jZS5Qcm90b1BsYXllclNvY2lhbElu",
            "Zm8iRAoaUGxheWVyUHJlc2VuY2VOb3RpZmljYXRpb24SEQoJcGxheWVyX2lk",
            "GAEgASgJEhMKC2xhc3Rfb25saW5lGAIgASgDIogBChlQbGF5ZXJMb2Fkb3V0",
            "Tm90aWZpY2F0aW9uEhEKCXBsYXllcl9pZBgBIAEoCRITCgtqb2JfY2FyZF9p",
            "ZBgCIAEoCRIVCg1zdWJfam9iX2luZGV4GAMgASgFEhEKCXdlYXBvbl9pZBgE",
            "IAEoCRIZChFzdWJfd2VhcG9uX3VubG9jaxgFIAEoBSJWChpQbGF5ZXJQb3Np",
            "dGlvbk5vdGlmaWNhdGlvbhIRCglwbGF5ZXJfaWQYASABKAkSEQoJcmVnaW9u",
            "X2lkGAIgASgJEhIKCm5vZGVfaW5kZXgYAyABKAUqzAEKEE5vdGlmaWNhdGlv",
            "blR5cGUSCAoEQkFTRRAAEg8KC1NPQ0lBTF9EQVRBEGQSEwoPUExBWUVSX1BS",
            "RVNFTkNFEGUSEwoPUExBWUVSX1BPU0lUSU9OEGYSEgoOUExBWUVSX0xPQURP",
            "VVQQZxIjCh5TT0NJQUxfQ0hBTExFTkdFX1BMQVlFUl9KT0lORUQQ2QQSHQoY",
            "U09DSUFMX0NIQUxMRU5HRV9TVEFSVEVEENoEEhsKFlNPQ0lBTF9DSEFMTEVO",
            "R0VfRU5ERUQQ2wRCT1o3Z2l0aHViLmNvbS9qdXN0amFjazE1MjEvbWV2aXVt",
            "L3BrZy9nZW5wcm90by9wcm90b3NvY2lhbKoCE01vYml1cy5Qcm90by5Tb2Np",
            "YWxiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Mobius.Proto.Social.SocialReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Mobius.Proto.Social.NotificationType), }, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Mobius.Proto.Social.SocialDataNotification), global::Mobius.Proto.Social.SocialDataNotification.Parser, new[]{ "FollowingList", "FollowerList", "RentalPlayers" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Mobius.Proto.Social.PlayerPresenceNotification), global::Mobius.Proto.Social.PlayerPresenceNotification.Parser, new[]{ "PlayerId", "LastOnline" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Mobius.Proto.Social.PlayerLoadoutNotification), global::Mobius.Proto.Social.PlayerLoadoutNotification.Parser, new[]{ "PlayerId", "JobCardId", "SubJobIndex", "WeaponId", "SubWeaponUnlock" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Mobius.Proto.Social.PlayerPositionNotification), global::Mobius.Proto.Social.PlayerPositionNotification.Parser, new[]{ "PlayerId", "RegionId", "NodeIndex" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Enums
  public enum NotificationType {
    [pbr::OriginalName("BASE")] Base = 0,
    [pbr::OriginalName("SOCIAL_DATA")] SocialData = 100,
    [pbr::OriginalName("PLAYER_PRESENCE")] PlayerPresence = 101,
    [pbr::OriginalName("PLAYER_POSITION")] PlayerPosition = 102,
    [pbr::OriginalName("PLAYER_LOADOUT")] PlayerLoadout = 103,
    [pbr::OriginalName("SOCIAL_CHALLENGE_PLAYER_JOINED")] SocialChallengePlayerJoined = 601,
    [pbr::OriginalName("SOCIAL_CHALLENGE_STARTED")] SocialChallengeStarted = 602,
    [pbr::OriginalName("SOCIAL_CHALLENGE_ENDED")] SocialChallengeEnded = 603,
  }

  #endregion

  #region Messages
  public sealed partial class SocialDataNotification : pb::IMessage<SocialDataNotification>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<SocialDataNotification> _parser = new pb::MessageParser<SocialDataNotification>(() => new SocialDataNotification());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<SocialDataNotification> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Mobius.Proto.Social.SocialNotificationReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SocialDataNotification() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SocialDataNotification(SocialDataNotification other) : this() {
      followingList_ = other.followingList_.Clone();
      followerList_ = other.followerList_.Clone();
      rentalPlayers_ = other.rentalPlayers_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SocialDataNotification Clone() {
      return new SocialDataNotification(this);
    }

    /// <summary>Field number for the "following_list" field.</summary>
    public const int FollowingListFieldNumber = 1;
    private static readonly pb::FieldCodec<global::Mobius.Proto.Social.ProtoPlayerSocialInfo> _repeated_followingList_codec
        = pb::FieldCodec.ForMessage(10, global::Mobius.Proto.Social.ProtoPlayerSocialInfo.Parser);
    private readonly pbc::RepeatedField<global::Mobius.Proto.Social.ProtoPlayerSocialInfo> followingList_ = new pbc::RepeatedField<global::Mobius.Proto.Social.ProtoPlayerSocialInfo>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Mobius.Proto.Social.ProtoPlayerSocialInfo> FollowingList {
      get { return followingList_; }
    }

    /// <summary>Field number for the "follower_list" field.</summary>
    public const int FollowerListFieldNumber = 2;
    private static readonly pb::FieldCodec<global::Mobius.Proto.Social.ProtoPlayerSocialInfo> _repeated_followerList_codec
        = pb::FieldCodec.ForMessage(18, global::Mobius.Proto.Social.ProtoPlayerSocialInfo.Parser);
    private readonly pbc::RepeatedField<global::Mobius.Proto.Social.ProtoPlayerSocialInfo> followerList_ = new pbc::RepeatedField<global::Mobius.Proto.Social.ProtoPlayerSocialInfo>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Mobius.Proto.Social.ProtoPlayerSocialInfo> FollowerList {
      get { return followerList_; }
    }

    /// <summary>Field number for the "rental_players" field.</summary>
    public const int RentalPlayersFieldNumber = 3;
    private static readonly pb::FieldCodec<global::Mobius.Proto.Social.ProtoPlayerSocialInfo> _repeated_rentalPlayers_codec
        = pb::FieldCodec.ForMessage(26, global::Mobius.Proto.Social.ProtoPlayerSocialInfo.Parser);
    private readonly pbc::RepeatedField<global::Mobius.Proto.Social.ProtoPlayerSocialInfo> rentalPlayers_ = new pbc::RepeatedField<global::Mobius.Proto.Social.ProtoPlayerSocialInfo>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Mobius.Proto.Social.ProtoPlayerSocialInfo> RentalPlayers {
      get { return rentalPlayers_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as SocialDataNotification);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(SocialDataNotification other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!followingList_.Equals(other.followingList_)) return false;
      if(!followerList_.Equals(other.followerList_)) return false;
      if(!rentalPlayers_.Equals(other.rentalPlayers_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= followingList_.GetHashCode();
      hash ^= followerList_.GetHashCode();
      hash ^= rentalPlayers_.GetHashCode();
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
      followingList_.WriteTo(output, _repeated_followingList_codec);
      followerList_.WriteTo(output, _repeated_followerList_codec);
      rentalPlayers_.WriteTo(output, _repeated_rentalPlayers_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      followingList_.WriteTo(ref output, _repeated_followingList_codec);
      followerList_.WriteTo(ref output, _repeated_followerList_codec);
      rentalPlayers_.WriteTo(ref output, _repeated_rentalPlayers_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      size += followingList_.CalculateSize(_repeated_followingList_codec);
      size += followerList_.CalculateSize(_repeated_followerList_codec);
      size += rentalPlayers_.CalculateSize(_repeated_rentalPlayers_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(SocialDataNotification other) {
      if (other == null) {
        return;
      }
      followingList_.Add(other.followingList_);
      followerList_.Add(other.followerList_);
      rentalPlayers_.Add(other.rentalPlayers_);
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
            followingList_.AddEntriesFrom(input, _repeated_followingList_codec);
            break;
          }
          case 18: {
            followerList_.AddEntriesFrom(input, _repeated_followerList_codec);
            break;
          }
          case 26: {
            rentalPlayers_.AddEntriesFrom(input, _repeated_rentalPlayers_codec);
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
            followingList_.AddEntriesFrom(ref input, _repeated_followingList_codec);
            break;
          }
          case 18: {
            followerList_.AddEntriesFrom(ref input, _repeated_followerList_codec);
            break;
          }
          case 26: {
            rentalPlayers_.AddEntriesFrom(ref input, _repeated_rentalPlayers_codec);
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class PlayerPresenceNotification : pb::IMessage<PlayerPresenceNotification>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<PlayerPresenceNotification> _parser = new pb::MessageParser<PlayerPresenceNotification>(() => new PlayerPresenceNotification());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<PlayerPresenceNotification> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Mobius.Proto.Social.SocialNotificationReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PlayerPresenceNotification() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PlayerPresenceNotification(PlayerPresenceNotification other) : this() {
      playerId_ = other.playerId_;
      lastOnline_ = other.lastOnline_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PlayerPresenceNotification Clone() {
      return new PlayerPresenceNotification(this);
    }

    /// <summary>Field number for the "player_id" field.</summary>
    public const int PlayerIdFieldNumber = 1;
    private string playerId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string PlayerId {
      get { return playerId_; }
      set {
        playerId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "last_online" field.</summary>
    public const int LastOnlineFieldNumber = 2;
    private long lastOnline_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public long LastOnline {
      get { return lastOnline_; }
      set {
        lastOnline_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as PlayerPresenceNotification);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(PlayerPresenceNotification other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (PlayerId != other.PlayerId) return false;
      if (LastOnline != other.LastOnline) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (PlayerId.Length != 0) hash ^= PlayerId.GetHashCode();
      if (LastOnline != 0L) hash ^= LastOnline.GetHashCode();
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
      if (PlayerId.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(PlayerId);
      }
      if (LastOnline != 0L) {
        output.WriteRawTag(16);
        output.WriteInt64(LastOnline);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (PlayerId.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(PlayerId);
      }
      if (LastOnline != 0L) {
        output.WriteRawTag(16);
        output.WriteInt64(LastOnline);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (PlayerId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(PlayerId);
      }
      if (LastOnline != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(LastOnline);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(PlayerPresenceNotification other) {
      if (other == null) {
        return;
      }
      if (other.PlayerId.Length != 0) {
        PlayerId = other.PlayerId;
      }
      if (other.LastOnline != 0L) {
        LastOnline = other.LastOnline;
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
            PlayerId = input.ReadString();
            break;
          }
          case 16: {
            LastOnline = input.ReadInt64();
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
            PlayerId = input.ReadString();
            break;
          }
          case 16: {
            LastOnline = input.ReadInt64();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class PlayerLoadoutNotification : pb::IMessage<PlayerLoadoutNotification>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<PlayerLoadoutNotification> _parser = new pb::MessageParser<PlayerLoadoutNotification>(() => new PlayerLoadoutNotification());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<PlayerLoadoutNotification> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Mobius.Proto.Social.SocialNotificationReflection.Descriptor.MessageTypes[2]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PlayerLoadoutNotification() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PlayerLoadoutNotification(PlayerLoadoutNotification other) : this() {
      playerId_ = other.playerId_;
      jobCardId_ = other.jobCardId_;
      subJobIndex_ = other.subJobIndex_;
      weaponId_ = other.weaponId_;
      subWeaponUnlock_ = other.subWeaponUnlock_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PlayerLoadoutNotification Clone() {
      return new PlayerLoadoutNotification(this);
    }

    /// <summary>Field number for the "player_id" field.</summary>
    public const int PlayerIdFieldNumber = 1;
    private string playerId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string PlayerId {
      get { return playerId_; }
      set {
        playerId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "job_card_id" field.</summary>
    public const int JobCardIdFieldNumber = 2;
    private string jobCardId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string JobCardId {
      get { return jobCardId_; }
      set {
        jobCardId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "sub_job_index" field.</summary>
    public const int SubJobIndexFieldNumber = 3;
    private int subJobIndex_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int SubJobIndex {
      get { return subJobIndex_; }
      set {
        subJobIndex_ = value;
      }
    }

    /// <summary>Field number for the "weapon_id" field.</summary>
    public const int WeaponIdFieldNumber = 4;
    private string weaponId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string WeaponId {
      get { return weaponId_; }
      set {
        weaponId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "sub_weapon_unlock" field.</summary>
    public const int SubWeaponUnlockFieldNumber = 5;
    private int subWeaponUnlock_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int SubWeaponUnlock {
      get { return subWeaponUnlock_; }
      set {
        subWeaponUnlock_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as PlayerLoadoutNotification);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(PlayerLoadoutNotification other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (PlayerId != other.PlayerId) return false;
      if (JobCardId != other.JobCardId) return false;
      if (SubJobIndex != other.SubJobIndex) return false;
      if (WeaponId != other.WeaponId) return false;
      if (SubWeaponUnlock != other.SubWeaponUnlock) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (PlayerId.Length != 0) hash ^= PlayerId.GetHashCode();
      if (JobCardId.Length != 0) hash ^= JobCardId.GetHashCode();
      if (SubJobIndex != 0) hash ^= SubJobIndex.GetHashCode();
      if (WeaponId.Length != 0) hash ^= WeaponId.GetHashCode();
      if (SubWeaponUnlock != 0) hash ^= SubWeaponUnlock.GetHashCode();
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
      if (PlayerId.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(PlayerId);
      }
      if (JobCardId.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(JobCardId);
      }
      if (SubJobIndex != 0) {
        output.WriteRawTag(24);
        output.WriteInt32(SubJobIndex);
      }
      if (WeaponId.Length != 0) {
        output.WriteRawTag(34);
        output.WriteString(WeaponId);
      }
      if (SubWeaponUnlock != 0) {
        output.WriteRawTag(40);
        output.WriteInt32(SubWeaponUnlock);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (PlayerId.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(PlayerId);
      }
      if (JobCardId.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(JobCardId);
      }
      if (SubJobIndex != 0) {
        output.WriteRawTag(24);
        output.WriteInt32(SubJobIndex);
      }
      if (WeaponId.Length != 0) {
        output.WriteRawTag(34);
        output.WriteString(WeaponId);
      }
      if (SubWeaponUnlock != 0) {
        output.WriteRawTag(40);
        output.WriteInt32(SubWeaponUnlock);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (PlayerId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(PlayerId);
      }
      if (JobCardId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(JobCardId);
      }
      if (SubJobIndex != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(SubJobIndex);
      }
      if (WeaponId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(WeaponId);
      }
      if (SubWeaponUnlock != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(SubWeaponUnlock);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(PlayerLoadoutNotification other) {
      if (other == null) {
        return;
      }
      if (other.PlayerId.Length != 0) {
        PlayerId = other.PlayerId;
      }
      if (other.JobCardId.Length != 0) {
        JobCardId = other.JobCardId;
      }
      if (other.SubJobIndex != 0) {
        SubJobIndex = other.SubJobIndex;
      }
      if (other.WeaponId.Length != 0) {
        WeaponId = other.WeaponId;
      }
      if (other.SubWeaponUnlock != 0) {
        SubWeaponUnlock = other.SubWeaponUnlock;
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
            PlayerId = input.ReadString();
            break;
          }
          case 18: {
            JobCardId = input.ReadString();
            break;
          }
          case 24: {
            SubJobIndex = input.ReadInt32();
            break;
          }
          case 34: {
            WeaponId = input.ReadString();
            break;
          }
          case 40: {
            SubWeaponUnlock = input.ReadInt32();
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
            PlayerId = input.ReadString();
            break;
          }
          case 18: {
            JobCardId = input.ReadString();
            break;
          }
          case 24: {
            SubJobIndex = input.ReadInt32();
            break;
          }
          case 34: {
            WeaponId = input.ReadString();
            break;
          }
          case 40: {
            SubWeaponUnlock = input.ReadInt32();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class PlayerPositionNotification : pb::IMessage<PlayerPositionNotification>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<PlayerPositionNotification> _parser = new pb::MessageParser<PlayerPositionNotification>(() => new PlayerPositionNotification());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<PlayerPositionNotification> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Mobius.Proto.Social.SocialNotificationReflection.Descriptor.MessageTypes[3]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PlayerPositionNotification() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PlayerPositionNotification(PlayerPositionNotification other) : this() {
      playerId_ = other.playerId_;
      regionId_ = other.regionId_;
      nodeIndex_ = other.nodeIndex_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PlayerPositionNotification Clone() {
      return new PlayerPositionNotification(this);
    }

    /// <summary>Field number for the "player_id" field.</summary>
    public const int PlayerIdFieldNumber = 1;
    private string playerId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string PlayerId {
      get { return playerId_; }
      set {
        playerId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "region_id" field.</summary>
    public const int RegionIdFieldNumber = 2;
    private string regionId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string RegionId {
      get { return regionId_; }
      set {
        regionId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "node_index" field.</summary>
    public const int NodeIndexFieldNumber = 3;
    private int nodeIndex_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int NodeIndex {
      get { return nodeIndex_; }
      set {
        nodeIndex_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as PlayerPositionNotification);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(PlayerPositionNotification other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (PlayerId != other.PlayerId) return false;
      if (RegionId != other.RegionId) return false;
      if (NodeIndex != other.NodeIndex) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (PlayerId.Length != 0) hash ^= PlayerId.GetHashCode();
      if (RegionId.Length != 0) hash ^= RegionId.GetHashCode();
      if (NodeIndex != 0) hash ^= NodeIndex.GetHashCode();
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
      if (PlayerId.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(PlayerId);
      }
      if (RegionId.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(RegionId);
      }
      if (NodeIndex != 0) {
        output.WriteRawTag(24);
        output.WriteInt32(NodeIndex);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (PlayerId.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(PlayerId);
      }
      if (RegionId.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(RegionId);
      }
      if (NodeIndex != 0) {
        output.WriteRawTag(24);
        output.WriteInt32(NodeIndex);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (PlayerId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(PlayerId);
      }
      if (RegionId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(RegionId);
      }
      if (NodeIndex != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(NodeIndex);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(PlayerPositionNotification other) {
      if (other == null) {
        return;
      }
      if (other.PlayerId.Length != 0) {
        PlayerId = other.PlayerId;
      }
      if (other.RegionId.Length != 0) {
        RegionId = other.RegionId;
      }
      if (other.NodeIndex != 0) {
        NodeIndex = other.NodeIndex;
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
            PlayerId = input.ReadString();
            break;
          }
          case 18: {
            RegionId = input.ReadString();
            break;
          }
          case 24: {
            NodeIndex = input.ReadInt32();
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
            PlayerId = input.ReadString();
            break;
          }
          case 18: {
            RegionId = input.ReadString();
            break;
          }
          case 24: {
            NodeIndex = input.ReadInt32();
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
