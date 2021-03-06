// Code generated by errors-const v0.2.0 tool DO NOT EDIT.

package ably

// ErrorCode is the type for predefined Ably error codes.
type ErrorCode int

const (
	ErrNotSet                                                                  ErrorCode = 0
	ErrNoError                                                                 ErrorCode = 10000
	ErrBadRequest                                                              ErrorCode = 40000
	ErrInvalidRequestBody                                                      ErrorCode = 40001
	ErrInvalidParameterName                                                    ErrorCode = 40002
	ErrInvalidParameterValue                                                   ErrorCode = 40003
	ErrInvalidHeader                                                           ErrorCode = 40004
	ErrInvalidCredential                                                       ErrorCode = 40005
	ErrInvalidConnectionID                                                     ErrorCode = 40006
	ErrInvalidMessageID                                                        ErrorCode = 40007
	ErrInvalidContentLength                                                    ErrorCode = 40008
	ErrMaximumMessageLengthExceeded                                            ErrorCode = 40009
	ErrInvalidChannelName                                                      ErrorCode = 40010
	ErrStaleRingState                                                          ErrorCode = 40011
	ErrInvalidClientID                                                         ErrorCode = 40012
	ErrInvalidMessageDataOrEncoding                                            ErrorCode = 40013
	ErrResourceDisposed                                                        ErrorCode = 40014
	ErrInvalidDeviceID                                                         ErrorCode = 40015
	ErrBatchError                                                              ErrorCode = 40020
	ErrInvalidPublishRequestUnspecified                                        ErrorCode = 40030
	ErrInvalidPublishRequestInvalidClientSpecifiedID                           ErrorCode = 40031
	ErrUnauthorized                                                            ErrorCode = 40100
	ErrInvalidCredentials                                                      ErrorCode = 40101
	ErrIncompatibleCredentials                                                 ErrorCode = 40102
	ErrInvalidUseOfBasicAuthOverNonTLSTransport                                ErrorCode = 40103
	ErrTimestampNotCurrent                                                     ErrorCode = 40104
	ErrNonceValueReplayed                                                      ErrorCode = 40105
	ErrUnableToObtainCredentialsFromGivenParameters                            ErrorCode = 40106
	ErrAccountDisabled                                                         ErrorCode = 40110
	ErrAccountRestrictedConnectionLimitsExceeded                               ErrorCode = 40111
	ErrAccountBlockedMessageLimitsExceeded                                     ErrorCode = 40112
	ErrAccountBlocked                                                          ErrorCode = 40113
	ErrAccountRestrictedChannelLimitsExceeded                                  ErrorCode = 40114
	ErrApplicationDisabled                                                     ErrorCode = 40120
	ErrKeyErrorUnspecified                                                     ErrorCode = 40130
	ErrKeyRevoked                                                              ErrorCode = 40131
	ErrKeyExpired                                                              ErrorCode = 40132
	ErrKeyDisabled                                                             ErrorCode = 40133
	ErrTokenErrorUnspecified                                                   ErrorCode = 40140
	ErrTokenRevoked                                                            ErrorCode = 40141
	ErrTokenExpired                                                            ErrorCode = 40142
	ErrTokenUnrecognised                                                       ErrorCode = 40143
	ErrInvalidJWTFormat                                                        ErrorCode = 40144
	ErrInvalidTokenFormat                                                      ErrorCode = 40145
	ErrConnectionBlockedLimitsExceeded                                         ErrorCode = 40150
	ErrOperationNotPermittedWithProvidedCapability                             ErrorCode = 40160
	ErrErrorFromClientTokenCallback                                            ErrorCode = 40170
	ErrForbidden                                                               ErrorCode = 40300
	ErrAccountDoesNotPermitTLSConnection                                       ErrorCode = 40310
	ErrOperationRequiresTLSConnection                                          ErrorCode = 40311
	ErrApplicationRequiresAuthentication                                       ErrorCode = 40320
	ErrNotFound                                                                ErrorCode = 40400
	ErrMethodNotAllowed                                                        ErrorCode = 40500
	ErrRateLimitExceededNonfatal                                               ErrorCode = 42910
	ErrMaxPerConnectionPublishRateLimitExceededNonfatal                        ErrorCode = 42911
	ErrRateLimitExceededFatal                                                  ErrorCode = 42920
	ErrMaxPerConnectionPublishRateLimitExceededFatal                           ErrorCode = 42921
	ErrInternalError                                                           ErrorCode = 50000
	ErrInternalChannelError                                                    ErrorCode = 50001
	ErrInternalConnectionError                                                 ErrorCode = 50002
	ErrTimeoutError                                                            ErrorCode = 50003
	ErrRequestFailedDueToOverloadedInstance                                    ErrorCode = 50004
	ErrReactorOperationFailed                                                  ErrorCode = 70000
	ErrReactorOperationFailedPostOperationFailed                               ErrorCode = 70001
	ErrReactorOperationFailedPostOperationReturnedUnexpectedCode               ErrorCode = 70002
	ErrReactorOperationFailedMaximumNumberOfConcurrentInFlightRequestsExceeded ErrorCode = 70003
	ErrExchangeErrorUnspecified                                                ErrorCode = 71000
	ErrForcedReAttachmentDueToPermissionsChange                                ErrorCode = 71001
	ErrExchangePublisherErrorUnspecified                                       ErrorCode = 71100
	ErrNoSuchPublisher                                                         ErrorCode = 71101
	ErrPublisherNotEnabledAsAnExchangePublisher                                ErrorCode = 71102
	ErrExchangeProductErrorUnspecified                                         ErrorCode = 71200
	ErrNoSuchProduct                                                           ErrorCode = 71201
	ErrProductDisabled                                                         ErrorCode = 71202
	ErrNoSuchChannelInThisProduct                                              ErrorCode = 71203
	ErrExchangeSubscriptionErrorUnspecified                                    ErrorCode = 71300
	ErrSubscriptionDisabled                                                    ErrorCode = 71301
	ErrRequesterHasNoSubscriptionToThisProduct                                 ErrorCode = 71302
	ErrConnectionFailed                                                        ErrorCode = 80000
	ErrConnectionFailedNoCompatibleTransport                                   ErrorCode = 80001
	ErrConnectionSuspended                                                     ErrorCode = 80002
	ErrDisconnected                                                            ErrorCode = 80003
	ErrAlreadyConnected                                                        ErrorCode = 80004
	ErrInvalidConnectionIDRemoteNotFound                                       ErrorCode = 80005
	ErrUnableToRecoverConnectionMessagesExpired                                ErrorCode = 80006
	ErrUnableToRecoverConnectionMessageLimitExceeded                           ErrorCode = 80007
	ErrUnableToRecoverConnectionConnectionExpired                              ErrorCode = 80008
	ErrConnectionNotEstablishedNoTransportHandle                               ErrorCode = 80009
	ErrInvalidOperationInvalidTransportHandle                                  ErrorCode = 80010
	ErrUnableToRecoverConnectionIncompatibleAuthParams                         ErrorCode = 80011
	ErrUnableToRecoverConnectionInvalidOrUnspecifiedConnectionSerial           ErrorCode = 80012
	ErrProtocolError                                                           ErrorCode = 80013
	ErrConnectionTimedOut                                                      ErrorCode = 80014
	ErrIncompatibleConnectionParameters                                        ErrorCode = 80015
	ErrOperationOnSupersededTransport                                          ErrorCode = 80016
	ErrConnectionClosed                                                        ErrorCode = 80017
	ErrInvalidConnectionIDInvalidFormat                                        ErrorCode = 80018
	ErrClientConfiguredAuthenticationProviderRequestFailed                     ErrorCode = 80019
	ErrContinuityLossDueToMaximumSubscribeMessageRateExceeded                  ErrorCode = 80020
	ErrClientRestrictionNotSatisfied                                           ErrorCode = 80030
	ErrChannelOperationFailed                                                  ErrorCode = 90000
	ErrChannelOperationFailedInvalidChannelState                               ErrorCode = 90001
	ErrChannelOperationFailedEpochExpiredOrNeverExisted                        ErrorCode = 90002
	ErrUnableToRecoverChannelMessagesExpired                                   ErrorCode = 90003
	ErrUnableToRecoverChannelMessageLimitExceeded                              ErrorCode = 90004
	ErrUnableToRecoverChannelNoMatchingEpoch                                   ErrorCode = 90005
	ErrUnableToRecoverChannelUnboundedRequest                                  ErrorCode = 90006
	ErrChannelOperationFailedNoResponseFromServer                              ErrorCode = 90007
	ErrMaximumNumberOfChannelsPerConnectionExceeded                            ErrorCode = 90010
	ErrUnableToEnterPresenceChannelNoClientID                                  ErrorCode = 91000
	ErrUnableToEnterPresenceChannelInvalidChannelState                         ErrorCode = 91001
	ErrUnableToLeavePresenceChannelThatIsNotEntered                            ErrorCode = 91002
	ErrUnableToEnterPresenceChannelMaximumMemberLimitExceeded                  ErrorCode = 91003
	ErrUnableToAutomaticallyReEnterPresenceChannel                             ErrorCode = 91004
	ErrPresenceStateIsOutOfSync                                                ErrorCode = 91005
	ErrMemberImplicitlyLeftPresenceChannelConnectionClosed                     ErrorCode = 91100
)

func (c ErrorCode) String() string {
	switch c {
	case ErrNotSet:
		return "(error code not set)"
	case ErrNoError:
		return "no error"
	case ErrBadRequest:
		return "bad request"
	case ErrInvalidRequestBody:
		return "invalid request body"
	case ErrInvalidParameterName:
		return "invalid parameter name"
	case ErrInvalidParameterValue:
		return "invalid parameter value"
	case ErrInvalidHeader:
		return "invalid header"
	case ErrInvalidCredential:
		return "invalid credential"
	case ErrInvalidConnectionID:
		return "invalid connection id"
	case ErrInvalidMessageID:
		return "invalid message id"
	case ErrInvalidContentLength:
		return "invalid content length"
	case ErrMaximumMessageLengthExceeded:
		return "maximum message length exceeded"
	case ErrInvalidChannelName:
		return "invalid channel name"
	case ErrStaleRingState:
		return "stale ring state"
	case ErrInvalidClientID:
		return "invalid client id"
	case ErrInvalidMessageDataOrEncoding:
		return "Invalid message data or encoding"
	case ErrResourceDisposed:
		return "Resource disposed"
	case ErrInvalidDeviceID:
		return "Invalid device id"
	case ErrBatchError:
		return "Batch error"
	case ErrInvalidPublishRequestUnspecified:
		return "Invalid publish request (unspecified)"
	case ErrInvalidPublishRequestInvalidClientSpecifiedID:
		return "Invalid publish request (invalid client-specified id)"
	case ErrUnauthorized:
		return "unauthorized"
	case ErrInvalidCredentials:
		return "invalid credentials"
	case ErrIncompatibleCredentials:
		return "incompatible credentials"
	case ErrInvalidUseOfBasicAuthOverNonTLSTransport:
		return "invalid use of Basic auth over non-TLS transport"
	case ErrTimestampNotCurrent:
		return "timestamp not current"
	case ErrNonceValueReplayed:
		return "nonce value replayed"
	case ErrUnableToObtainCredentialsFromGivenParameters:
		return "Unable to obtain credentials from given parameters"
	case ErrAccountDisabled:
		return "account disabled"
	case ErrAccountRestrictedConnectionLimitsExceeded:
		return "account restricted (connection limits exceeded)"
	case ErrAccountBlockedMessageLimitsExceeded:
		return "account blocked (message limits exceeded)"
	case ErrAccountBlocked:
		return "account blocked"
	case ErrAccountRestrictedChannelLimitsExceeded:
		return "account restricted (channel limits exceeded)"
	case ErrApplicationDisabled:
		return "application disabled"
	case ErrKeyErrorUnspecified:
		return "key error (unspecified)"
	case ErrKeyRevoked:
		return "key revoked"
	case ErrKeyExpired:
		return "key expired"
	case ErrKeyDisabled:
		return "key disabled"
	case ErrTokenErrorUnspecified:
		return "token error (unspecified)"
	case ErrTokenRevoked:
		return "token revoked"
	case ErrTokenExpired:
		return "token expired"
	case ErrTokenUnrecognised:
		return "token unrecognised"
	case ErrInvalidJWTFormat:
		return "invalid JWT format"
	case ErrInvalidTokenFormat:
		return "invalid token format"
	case ErrConnectionBlockedLimitsExceeded:
		return "connection blocked (limits exceeded)"
	case ErrOperationNotPermittedWithProvidedCapability:
		return "operation not permitted with provided capability"
	case ErrErrorFromClientTokenCallback:
		return "error from client token callback"
	case ErrForbidden:
		return "forbidden"
	case ErrAccountDoesNotPermitTLSConnection:
		return "account does not permit tls connection"
	case ErrOperationRequiresTLSConnection:
		return "operation requires tls connection"
	case ErrApplicationRequiresAuthentication:
		return "application requires authentication"
	case ErrNotFound:
		return "not found"
	case ErrMethodNotAllowed:
		return "method not allowed"
	case ErrRateLimitExceededNonfatal:
		return "rate limit exceeded (nonfatal): request rejected (unspecified)"
	case ErrMaxPerConnectionPublishRateLimitExceededNonfatal:
		return "max per-connection publish rate limit exceeded (nonfatal): unable to publish message"
	case ErrRateLimitExceededFatal:
		return "rate limit exceeded (fatal)"
	case ErrMaxPerConnectionPublishRateLimitExceededFatal:
		return "max per-connection publish rate limit exceeded (fatal); closing connection"
	case ErrInternalError:
		return "internal error"
	case ErrInternalChannelError:
		return "internal channel error"
	case ErrInternalConnectionError:
		return "internal connection error"
	case ErrTimeoutError:
		return "timeout error"
	case ErrRequestFailedDueToOverloadedInstance:
		return "Request failed due to overloaded instance"
	case ErrReactorOperationFailed:
		return "reactor operation failed"
	case ErrReactorOperationFailedPostOperationFailed:
		return "reactor operation failed (post operation failed)"
	case ErrReactorOperationFailedPostOperationReturnedUnexpectedCode:
		return "reactor operation failed (post operation returned unexpected code)"
	case ErrReactorOperationFailedMaximumNumberOfConcurrentInFlightRequestsExceeded:
		return "reactor operation failed (maximum number of concurrent in-flight requests exceeded)"
	case ErrExchangeErrorUnspecified:
		return "Exchange error (unspecified)"
	case ErrForcedReAttachmentDueToPermissionsChange:
		return "Forced re-attachment due to permissions change"
	case ErrExchangePublisherErrorUnspecified:
		return "Exchange publisher error (unspecified)"
	case ErrNoSuchPublisher:
		return "No such publisher"
	case ErrPublisherNotEnabledAsAnExchangePublisher:
		return "Publisher not enabled as an exchange publisher"
	case ErrExchangeProductErrorUnspecified:
		return "Exchange product error (unspecified)"
	case ErrNoSuchProduct:
		return "No such product"
	case ErrProductDisabled:
		return "Product disabled"
	case ErrNoSuchChannelInThisProduct:
		return "No such channel in this product"
	case ErrExchangeSubscriptionErrorUnspecified:
		return "Exchange subscription error (unspecified)"
	case ErrSubscriptionDisabled:
		return "Subscription disabled"
	case ErrRequesterHasNoSubscriptionToThisProduct:
		return "Requester has no subscription to this product"
	case ErrConnectionFailed:
		return "connection failed"
	case ErrConnectionFailedNoCompatibleTransport:
		return "connection failed (no compatible transport)"
	case ErrConnectionSuspended:
		return "connection suspended"
	case ErrDisconnected:
		return "disconnected"
	case ErrAlreadyConnected:
		return "already connected"
	case ErrInvalidConnectionIDRemoteNotFound:
		return "invalid connection id (remote not found)"
	case ErrUnableToRecoverConnectionMessagesExpired:
		return "unable to recover connection (messages expired)"
	case ErrUnableToRecoverConnectionMessageLimitExceeded:
		return "unable to recover connection (message limit exceeded)"
	case ErrUnableToRecoverConnectionConnectionExpired:
		return "unable to recover connection (connection expired)"
	case ErrConnectionNotEstablishedNoTransportHandle:
		return "connection not established (no transport handle)"
	case ErrInvalidOperationInvalidTransportHandle:
		return "invalid operation (invalid transport handle)"
	case ErrUnableToRecoverConnectionIncompatibleAuthParams:
		return "unable to recover connection (incompatible auth params)"
	case ErrUnableToRecoverConnectionInvalidOrUnspecifiedConnectionSerial:
		return "unable to recover connection (invalid or unspecified connection serial)"
	case ErrProtocolError:
		return "protocol error"
	case ErrConnectionTimedOut:
		return "connection timed out"
	case ErrIncompatibleConnectionParameters:
		return "incompatible connection parameters"
	case ErrOperationOnSupersededTransport:
		return "operation on superseded transport"
	case ErrConnectionClosed:
		return "connection closed"
	case ErrInvalidConnectionIDInvalidFormat:
		return "invalid connection id (invalid format)"
	case ErrClientConfiguredAuthenticationProviderRequestFailed:
		return "client configured authentication provider request failed"
	case ErrContinuityLossDueToMaximumSubscribeMessageRateExceeded:
		return "continuity loss due to maximum subscribe message rate exceeded"
	case ErrClientRestrictionNotSatisfied:
		return "client restriction not satisfied"
	case ErrChannelOperationFailed:
		return "channel operation failed"
	case ErrChannelOperationFailedInvalidChannelState:
		return "channel operation failed (invalid channel state)"
	case ErrChannelOperationFailedEpochExpiredOrNeverExisted:
		return "channel operation failed (epoch expired or never existed)"
	case ErrUnableToRecoverChannelMessagesExpired:
		return "unable to recover channel (messages expired)"
	case ErrUnableToRecoverChannelMessageLimitExceeded:
		return "unable to recover channel (message limit exceeded)"
	case ErrUnableToRecoverChannelNoMatchingEpoch:
		return "unable to recover channel (no matching epoch)"
	case ErrUnableToRecoverChannelUnboundedRequest:
		return "unable to recover channel (unbounded request)"
	case ErrChannelOperationFailedNoResponseFromServer:
		return "channel operation failed (no response from server)"
	case ErrMaximumNumberOfChannelsPerConnectionExceeded:
		return "maximum number of channels per connection exceeded"
	case ErrUnableToEnterPresenceChannelNoClientID:
		return "unable to enter presence channel (no clientId)"
	case ErrUnableToEnterPresenceChannelInvalidChannelState:
		return "unable to enter presence channel (invalid channel state)"
	case ErrUnableToLeavePresenceChannelThatIsNotEntered:
		return "unable to leave presence channel that is not entered"
	case ErrUnableToEnterPresenceChannelMaximumMemberLimitExceeded:
		return "unable to enter presence channel (maximum member limit exceeded)"
	case ErrUnableToAutomaticallyReEnterPresenceChannel:
		return "unable to automatically re-enter presence channel"
	case ErrPresenceStateIsOutOfSync:
		return "presence state is out of sync"
	case ErrMemberImplicitlyLeftPresenceChannelConnectionClosed:
		return "member implicitly left presence channel (connection closed)"
	}
	return ""
}
